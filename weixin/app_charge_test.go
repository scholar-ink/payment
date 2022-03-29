package weixin

import (
	"fmt"
	"testing"
	"time"
)

func TestHandle(t *testing.T) {

	app := new(PubCharge)

	app.InitBaseConfig(&BaseConfig{
		AppId:          "wxa33cba2b69f869f3",
		MchId:          "1491561542",
		Md5Key:         "9689489231d792a260e2559586276916",
		SignType:       "MD5",
		ExpireDuration: time.Second * 60,
		NotifyUrl:      "http://api.store.udian.me/v1/payment/notify",
	})

	//////服务商发起支付
	//app.InitBaseConfig(&BaseConfig{
	//	AppId:     "wxf06ac118ca3d9533",
	//	MchId:     "1495589652",
	//	SubAppId:  "wxa33cba2b69f869f3",
	//	SubMchId:  "1495746312",
	//	Md5Key:    "057177a8459352933f755c535b0ab0ef",
	//	SignType:  "MD5",
	//	NotifyUrl: "http://api.store.udian.me/v1/payment/notify",
	//})

	ret, err := app.Handle(map[string]interface{}{
		"device_info":      "WEB",
		"body":             "腾讯充值中心-QQ会员充",
		"detail":           "商品详细描述",
		"attach":           "1111",
		"out_trade_no":     "20150806125346212221121",
		"fee_type":         "CNY",
		"total_fee":        1,
		"spbill_create_ip": "123.12.12.123",
		"goods_tag":        "",
		//"openid":           "oyA310LEnY_JW_-BDHVJguSpFyKQ",
		"sub_openid": "oyA310LEnY_JW_-BDHVJguSpFyKQ",
	})

	fmt.Printf("%+v", ret)

	fmt.Println(err)
}
