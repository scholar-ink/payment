package qf

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCharge_Handle(t *testing.T) {
	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		AppCode: "9A02B598D8F64BD789DBE8C6D4805D07",
		Key:     "F3954E8B34F5474AA38BB127D5BA94FE",
	})

	ret, err := charge.Handle(&OrderChargeConf{
		MchId:      "jzG3WuLpjZ",
		PayType:    "800207",
		OutTradeNo: "200103020256635777486116",
		TxAmt:      "1",
		GoodsName:  "订单号：B190925121912999938",
		SubOpenid:  "oNEbm1CDWvqk7nYy7bhhc4hm-8Y8",
	})

	fmt.Println(err)

	b, _ := json.Marshal(ret.PayParams)

	fmt.Println(string(b))

	fmt.Printf("%+v", ret.PayParams)
	fmt.Println(err)
}

func TestOrderCharge_GetAuthCodeUrl(t *testing.T) {
	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		AppCode: "9A02B598D8F64BD789DBE8C6D4805D07",
		Key:     "F3954E8B34F5474AA38BB127D5BA94FE",
	})

	ret := charge.GetAuthCodeUrl(&GetAuthCodeUrlConf{
		MchId:       "vrR0giJVgG",
		RedirectUri: "https://openapi-test.qfpay.com/tools/get_wx_code",
	})

	fmt.Println(ret)
}

func TestOrderCharge_GetOpenId(t *testing.T) {
	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		AppCode: "9A02B598D8F64BD789DBE8C6D4805D07",
		Key:     "F3954E8B34F5474AA38BB127D5BA94FE",
	})

	ret, err := charge.GetOpenId(&GetOpenIdConf{
		MchId: "vrR0giJVgG",
		Code:  "061v6ezu1o3JDh0j7kzu101Zyu1v6ezY",
	})

	fmt.Println(ret)
	fmt.Println(err)
}
