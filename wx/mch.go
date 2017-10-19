package wx

import (
	"_/http"
	"fmt"
	"time"
)

func (self app) UnifiedOrder(notify_url, out_trade_no, total_fee, body, spbill_create_ip, trade_type, openid, attach string) (json_string string, err error) {
	var _map = map[string]string{
		"appid":  self.appid,
		"mch_id": self.mch_id,
		"key":    self.key,

		"notify_url": notify_url,
		//*商户订单号
		"out_trade_no": out_trade_no,
		//*随机字符串
		"nonce_str": out_trade_no,
		//*标价金额
		"total_fee": total_fee,
		//*商品描述
		"body": body,
		//*终端IP
		"spbill_create_ip": spbill_create_ip,

		//*交易类型
		"trade_type": trade_type,
		//*用户标识
		"openid": openid,
		//*商品ID
		"product_id": "12235413214070356458058",
	}
	if attach != "" {
		_map["attach"] = attach
	}

	_map["sign"] = sign(_map)

	response := http.Post("https://api.mch.weixin.qq.com/pay/unifiedorder", map2xml(_map), nil)
	err = response.Error
	if response.StatusCode != 200 || response.Error != nil {
		return
	}

	var (
		xml       string
		prepay_id string
		has       bool
	)

	xml = string(response.Body)
	if prepay_id, has = Xml2map(xml)["prepay_id"]; has == false {
		err = fmt.Errorf("%s", xml)
		return
	}

	_map = map[string]string{
		"appId": self.appid,
		"key":   self.key,

		"package":   fmt.Sprintf("prepay_id=%s", prepay_id),
		"signType":  "MD5",
		"timeStamp": fmt.Sprint(time.Now().Unix()),
		"nonceStr":  fmt.Sprint(time.Now().UnixNano()),
	}

	_map["paySign"] = sign(_map)
	json_string = map2json(_map)
	return
}

func (self app) OrderQuery(out_trade_no string) {
	// var data = xml_data(map[string]string{
	// 	"appid":  self.appid,
	// 	"mch_id": self.mch_id,
	// 	"key":    self.key,

	// 	"nonce_str":    fmt.Sprintf("%x", md5.Sum([]byte(out_trade_no))),
	// 	"out_trade_no": out_trade_no,
	// })

	// response := http.Post("https://api.mch.weixin.qq.com/pay/orderquery", data, nil)

	// if response.StatusCode == 200 && response.Error == nil {
	// 	_map := xml2map(string(response.Body))
	// 	fmt.Println(_map)

	// }

}

func (self app) CloseOrder() {

}

func (self app) Refund() {

}

func (self app) RefundQuery() {

}
