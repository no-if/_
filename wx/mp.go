package wx

import (
	"_/http"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func (self *app) Get_access_token() (err error) {
	var response = http.Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", self.appid, self.secret), nil)
	if err = response.Error; err != nil && response.StatusCode != 200 {
		return
	}
	self.access_token = gjson.GetBytes(response.Body, "access_token").String()
	if self.access_token == "" {
		err = fmt.Errorf("%s", response.Body)
	}

	return
}

func (self *app) Get_ticket() (err error) {
	var response = http.Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi", self.access_token), nil)
	if err = response.Error; err != nil && response.StatusCode != 200 {
		return
	}

	self.ticket = gjson.GetBytes(response.Body, "ticket").String()
	return
}

func (self *app) Js_config() string {

	return ""
}

func (self *app) Get_server_ip() (ips []string, err error) {
	var response = http.Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/getcallbackip?access_token=%s", self.access_token), nil)
	if err = response.Error; err != nil && response.StatusCode != 200 {
		return
	}

	for _, ip := range gjson.GetBytes(response.Body, "ip_list").Array() {
		ips = append(ips, ip.String())
	}
	return
}

func (self *app) Get_openid(code string) (openid string, err error) {
	var response = http.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", self.appid, self.secret, code), nil)
	if err = response.Error; err != nil && response.StatusCode != 200 {
		return
	}

	openid = gjson.GetBytes(response.Body, "openid").String()

	return
}

func (self *app) Get_userinfo(openid string) (wx string, err error) {
	var response = http.Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s", self.access_token, openid), nil)
	if err = response.Error; err != nil && response.StatusCode != 200 {
		return
	}

	wx = string(response.Body)
	gjson.Parse(wx).ForEach(func(key, value gjson.Result) bool {
		switch k := key.String(); k {
		case "openid", "nickname", "sex", "city", "headimgurl":
		default:
			wx, _ = sjson.Delete(wx, k)
		}
		return true
	})

	return
}
