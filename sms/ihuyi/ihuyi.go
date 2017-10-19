package ihuyi

import (
	"_/http"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/*
用户:	cf_1YYPC
密码:	1YYPPC
登录地址:	http://sms.ihuyi.com/login.html
*/

type User struct {
	api      string
	account  string
	password string
}

func Parse(api string, account string, password string) (user User) {
	fmt.Println("api", api)
	user.api = api
	user.account = account
	user.password = password
	return user
}

func (self *User) Send(mobile string, msg string) (code int) {
	/*
	   <?xml version="1.0" encoding="utf-8"?>
	   <SubmitResult xmlns="http://106.ihuyi.com/">
	   <code>2</code>
	   <msg>提交成功</msg>
	   <smsid>397373856</smsid>
	   </SubmitResult>
	*/
	code = -1
	var url = fmt.Sprintf("%s?method=Submit", self.api)
	var data = fmt.Sprintf("account=%s&password=%s&mobile=%s&content=%s", self.account, self.password, mobile, http.UrlEncode(msg))

	var re = http.PostForm(url, data, nil)
	if re.StatusCode == 200 {
		var tab = regexp.MustCompile(`<code>([-\d]+)</code>`).FindSubmatch(re.Body)
		if len(tab) == 2 {
			if m, err := strconv.Atoi(string(tab[1])); err == nil {
				code = m
			}
		}
	}
	return code
}

func (self *User) Get_money() (money int) {
	/*
	   <?xml version="1.0" encoding="utf-8"?>
	   <GetNumResult xmlns="http://106.ihuyi.com/">
	   <code>2</code>
	   <msg>查询成功</msg>
	   <num>20023</num>
	   </GetNumResult>
	*/
	money = -1
	var url = fmt.Sprintf("%s?method=GetNum", self.api)
	var data = fmt.Sprintf("account=%s&password=%s", self.account, self.password)
	var re = http.PostForm(url, data, nil)
	if re.StatusCode == 200 && strings.Contains(string(re.Body), "<code>2</code>") {
		var tab = regexp.MustCompile(`<num>([-\d]+)</num>`).FindSubmatch(re.Body)
		if len(tab) == 2 {
			if n, err := strconv.Atoi(string(tab[1])); err == nil {
				money = n
			}
		}
	}
	return money
}
