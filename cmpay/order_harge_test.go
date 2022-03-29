package cmpay

import (
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"testing"
)

func TestCharge_Handle(t *testing.T) {
	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		Key: "4EFF68ACE1BC890CEAB068A4E8176503",
	})

	ret, err := charge.Handle(&OrderChargeConf{
		ShopId:         "f13db1959ed46e3cb1dfef8ea88bace5",
		OrderId:        helper.CreateSn(),
		Money:          "0.01",
		PayType:        "alipay",
		GoodsMsg:       "测试商品" + helper.CreateSn(),
		RedirectUrl:    "http://tq.udian.me/v1/common/enter-notify",
		RedirectNumber: "5",
		ReturnUrl:      "http://www.baidu.com",
		//SubAppid:  "wxdae782e8546f5bb6",
		//SubOpenid: "oNEbm1CDWvqk7nYy7bhhc4hm-8Y8",
	})

	//ret, err := charge.Handle2(&OrderChargeConf{
	//	ShopId:         "f13db1959ed46e3cb1dfef8ea88bace5",
	//	OrderId:        helper.CreateSn(),
	//	Money:          "0.01",
	//	OrderType:      "alipay",
	//	GoodsMsg:       "测试商品" + helper.CreateSn(),
	//	RedirectUrl:    "http://tq.udian.me/v1/common/enter-notify",
	//	RedirectNumber: "5",
	//	//ReturnUrl:      "http://www.baidu.com",
	//	SubAppid:  "wxdae782e8546f5bb6",
	//	SubOpenid: "oNEbm1CDWvqk7nYy7bhhc4hm-8Y8",
	//})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v \n", ret)
	}

}
