package sioo

import (
	"_/http"
	"crypto/md5"
	"fmt"
	"strconv"
)

/*
企业代码:	yywl
用户ID:	80918
密码:	yywl66
登录地址:	www.10690221.com



状态码	描述
0		操作成功
-1		签权失败
-2		未检索到被叫号码
-3	    被叫号码过多
-4		内容未签名
-5		内容过长
-6		余额不足
-7		暂停发送
-8		保留
-9		定时发送时间格式错误
-10		发内容为空
-11		账户无效
-12		Ip地址非法
-13		操作频率快
-14		操作失败
-15		拓展码无效
-16		取消定时,seqid错误
-17		未开通报告
-18		暂留
-19		未开通上行
-20     暂留
-21		包含屏蔽词
*/

type User struct {
	api  string
	uid  string
	auth string
}

func Parse(api string, uid string, code string, pass string) (user User) {
	user.api = api
	user.uid = uid
	user.auth = fmt.Sprintf("%x", md5.Sum([]byte(code+pass)))
	return user
}

func (self *User) Send(mobile string, msg string) (code int) {
	code = 1
	var url = fmt.Sprintf("%shy/?uid=%s&auth=%s&mobile=%s&msg=%s&expid=0&encode=utf-8", self.api, self.uid, self.auth, mobile, http.UrlEncode(msg))
	var re = http.Get(url, nil)
	if re.StatusCode == 200 {
		fmt.Sscanf(string(re.Body), "%d", &code)
	}
	return code
}

func (self *User) Get_money() (money int) {
	money = -1
	var url = fmt.Sprintf("%shy/m?uid=%s&auth=%s", self.api, self.uid, self.auth)
	var re = http.Get(url, nil)
	if re.StatusCode == 200 {
		if m, err := strconv.Atoi(string(re.Body)); err == nil {
			money = m
		}
	}
	return money
}
