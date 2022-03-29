package notify

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestXinHuiNotify_Handle(t *testing.T) {

	notify := new(XinHuiNotify)

	//ret := map[string]interface{}{
	//	"agent_mer_no":"8000100100221",
	//	"attach":"",
	//	"buyer_pay_amount":"0.02",
	//	"buyer_user_id":"2088612472076052",
	//	"channel_type":"UP_ALIPAY",
	//	"fund_bill_list":`[{"amount":"0.02","fund_channel":"ALIPAYACCOUNT"}]`,
	//	"merchant_no":"8000105202540",
	//	"org_external_id":"2019112622001476051421051188",
	//	"out_trade_no":"191126171631458247147781",
	//	"pay_external_id":"13071199255603843829761",
	//	"receipt_amount":"0.02",
	//	"rsp_code":"0000",
	//	"rsp_msg":"支付成功",
	//	"sign":"Ae7ki/criXfbk2syh7qM+8RNiFRk934fTWlGEk64LB0ddLek7IoOO/B04OUTHPhN8d2Q61UlSii3gRmtiPGlv7CXRK0FbhCMg5I3whcpEFKiIcPGUyflmdHm4EzAXUxkvnXTShRZwb0cHs+8uaj4neD84WpVps2esW5Q4zuS/dw=",
	//	"time_end":"2019-11-26 17:18:45",
	//	"total_fee":"0.02",
	//	"trx_external_id":"13071199255603843829760",
	//	"trx_status":"SUCCESS",
	//	"version":"1.0",
	//}

	ret := map[string]interface{}{
		"agent_mer_no":         "8000100022311",
		"app_id":               "wxcf7675d03cae0701",
		"attach":               "",
		"bank_type":            "OTHERS",
		"cash_fee":             "0.5",
		"channel_type":         "UP_WX",
		"coupon_fee":           "0.0",
		"merchant_no":          "8000106469540",
		"openid":               "oNEbm1NzJeS-j4UprwmR5MAkrnfY",
		"org_external_id":      "4200000480202002146983902985",
		"out_trade_no":         "200214200119356239047148",
		"pay_external_id":      "10031228288105603846145",
		"rsp_code":             "0000",
		"rsp_msg":              "支付成功",
		"settlement_total_fee": "0.5",
		"sign":                 "etj8hy242v5LX3b2NxddSOamoFhv2pkiVQDgCbs6dLoewFXy5ZZhKhUB6rddbbY8GRMB8eRLZYeFM6alUnCNECkkp3NYcYMVugvRN6RXBXJ/nf51Byh3AIWcEAR2/uC OSgo568YBthr4fpMPRvgm6dOMX2QoaZyLihV TtK2 A=",
		"time_end":             "2020-02-14 20:01:22",
		"total_fee":            "0.5",
		"trx_external_id":      "10031228288105603846144",
		"trx_status":           "SUCCESS",
		"version":              "1.0",
		//"buyer_pay_amount":"0.10",
		//"buyer_user_id":"2088612472076052",
		//"fund_bill_list":`[{"amount":"0.10","fund_channel":"ALIPAYACCOUNT"}]`,
		//"merchant_no":"8000105202610",
		//"org_external_id":"2019112622001476051420837595",
		//"out_trade_no":"191126210912438233544500",
		//"pay_external_id":"12021199314159905411072",
		//"receipt_amount":"0.10",
		//"rsp_code":"0000",
		//"rsp_msg":"支付成功",
		//"sign":"c2fn9SVFwI0a5tE94S3rcVgdOfciRRYa5yiBmvs 0F3TjKKhbP7VdxLzHldgU7kq5FVvmgpH2zOiBoa/VGSYNBopoUo1nWI7yk3Fj1iFWgVliFdok28rW2Ha2eHK6jox6rKAVGphJi1RR//0Yb2yEkTOj0Np h6S8UrIE6UpcyE=",
		//"time_end":"2019-11-26 21:09:41",
		//"total_fee":"0.1",
		//"trx_external_id":"12021199314159901216768",
		//"trx_status":"SUCCESS",
		//"version":"1.0",
	}

	retData := notify.Handle(ret, func(merchantNo string) (md5Key string, err error) {
		return "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCG1S0FLarWe9F4RGWKQud+WpFD2GD+3trgL9FIkCOsWB2LC0azGbA53kdpaDRTiLQNvzJGgyQxNhPNJk78r08xiENscwXwayGjGwDMva09OUxhPvSy5ha/T3I0JRiedXAYJPgTc+WgvZcMvWGC7TdqYGXgs5uxZ2qSaHaRBaZPDQIDAQAB", nil
	}, func(data *XinHuiNotifyData) error {
		fmt.Printf("%+v\n", data)
		return nil
	})

	fmt.Println(retData)

}

func TestXinHuiNotify_Handle2(t *testing.T) {
	notify := &struct {
		PayExternalId string `json:"pay_external_id"`
	}{}

	err := json.Unmarshal([]byte(`{"agent_mer_no":"8000100100221","attach":"","buyer_pay_amount":"0.10","buyer_user_id":"2088612472076052","channel_type":"UP_ALIPAY","fund_bill_list":"[{\"amount\":\"0.10\",\"fund_channel\":\"ALIPAYACCOUNT\"}]","merchant_no":"8000105202610","org_external_id":"2019112622001476051420837595","out_trade_no":"191126210912438233544500","pay_external_id":"12021199314159905411072","receipt_amount":"0.10","rsp_code":"0000","rsp_msg":"支付成功","sign":"c2fn9SVFwI0a5tE94S3rcVgdOfciRRYa5yiBmvs 0F3TjKKhbP7VdxLzHldgU7kq5FVvmgpH2zOiBoa/VGSYNBopoUo1nWI7yk3Fj1iFWgVliFdok28rW2Ha2eHK6jox6rKAVGphJi1RR//0Yb2yEkTOj0Np h6S8UrIE6UpcyE=","time_end":"2019-11-26 21:09:41","total_fee":"0.1","trx_external_id":"12021199314159901216768","trx_status":"SUCCESS","version":"1.0"}`), notify)

	fmt.Println(err)

	fmt.Printf("%+v", notify)
}
