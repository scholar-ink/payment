package notify

import (
	"fmt"
	"testing"
)

func TestSumPayNotify_Handle(t *testing.T) {

	notify := new(SumPayNotify)

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

	ret := `{"order_no":"191213102513048271797200","mer_no":"s100000040","offset_amount":"0","success_amount":"0.01","sign":"NWFOdN5IC/lrjgBZpnqKUB4N5LlFp77rZd4vAIDjzsBDM14eNE7RD+vA+3zL76TOPt5X+M01rMYKOS0ra7rS1V135MqnbsQ9Mwwq1UgcLepzdif+zrDSQUug+2aZFCHiiCuiUTHwvK5lbchvl/R45D7VHubpuEcA4twesd1Zlb4dmWLQ5IWpj0VKOTM/jdjfJiy/XnHsdhdlRvIsQa/5euZqO6xL4BgHJk1Q8hPm42j8u0xINkBrwaV6igxEtBc+FD53ndcysuFdz1p5dXrQFDwhlEx3nUjf/mWkT+SEJts8wiAe7GAb91Asp347p+zB5gQC3zWwaGIPK/8SpTZCdQ==","order_time":"20191213102513","success_time":"20191213103047","resp_code":"000000","paid_amount":"0.01","trade_no":"T10831547","resp_msg":"success","sign_type":"CERT","status":"1"}`

	retData := notify.Handle(ret, func(merchantNo string) (md5Key string, err error) {
		return "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAmmIZecTy9IknC9qITcyLEL9Avwu4IegkiPZDLuyCQ2sAc4Z41WtNvLZ/vADmtLY89Lc35SXCZ/DLlR+OOYT8/MuURbQ82h47NiX+7qN8xWpUzNVWBNH2magAjoXKU4st0r2DGicHzBi9VvcNYi96JIzXjbOjfMWCT9BbAj+4tnCKCgajDT/waqHkjfOxgsNmtAdEA+TS9A0z8fiLweu+ndBz4C2F2SqYe4UWIqErvTnpPwRSLMEKGG8qKmJtHxKOPKvjXSTpv1KtSodOUWjAZRZOh6wrq9PwrSOj5KeBNgTz6Mikm3nqIAeElK4Eg6u66ni9nIBtwPHlD/DltJwpwQIDAQAB", nil
	}, func(data *SumPayNotifyData) error {
		fmt.Printf("%+v\n", data)
		return nil
	})

	fmt.Println(retData)

}
