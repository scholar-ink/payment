package notify

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAdaPayNotify_Handle(t *testing.T) {

	id := "002112020022500354610078055455124787200"
	fmt.Println(id[:32])

	notify := new(AdaPayNotify)

	ret := map[string]interface{}{
		"app_id":       "app_8f2ccfd0-3ae3-4318-8b9d-eaae22819ee3",
		"created_time": "1582562147",
		"data":         "{\"created_time\":\"20200225003547\",\"end_time\":\"20200225003611\",\"expend\":{},\"id\":\"002112020022500354610078055455124787200\",\"order_no\":\"200225003546239388799481\",\"party_order_id\":\"PAYUN00021200225003546m818r7qZDx\",\"pay_amt\":\"0.01\",\"pay_channel\":\"alipay\",\"status\":\"succeeded\"}",
		"id":           "002110078055560567775232",
		"object":       "payment",
		"prod_mode":    "true",
		"sign":         "pTlEpHHEX0XHDdnOs34jRHmtOv8OifcnIqBW2MfkAv5IJ+Y8kNmKUXXBVPKT9Qq4ETbu1u1AH88wwmAohkQpk0dW+l+0TuTT4TSxoPSLGSeN97PpY9Unm3rHHYFvKovDoAmkfHF+kz+GEWoY2HaTErL8nREDeVEeeYdIOMB9Q44=",
		"type":         "payment.succeeded",
	}

	retData := notify.Handle(ret, func(merchantNo string) (md5Key string, err error) {
		return "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCwN6xgd6Ad8v2hIIsQVnbt8a3JituR8o4Tc3B5WlcFR55bz4OMqrG/356Ur3cPbc2Fe8ArNd/0gZbC9q56Eb16JTkVNA/fye4SXznWxdyBPR7+guuJZHc/VW2fKH2lfZ2P3Tt0QkKZZoawYOGSMdIvO+WqK44updyax0ikK6JlNQIDAQAB", nil
	}, func(data *AdaPayNotifyData) error {
		fmt.Printf("%+v\n", data)
		return nil
	})

	fmt.Println(retData)

}

func TestAdaPayNotify_Handle2(t *testing.T) {
	notify := &struct {
		PayExternalId string `json:"pay_external_id"`
	}{}

	err := json.Unmarshal([]byte(`{"agent_mer_no":"8000100100221","attach":"","buyer_pay_amount":"0.10","buyer_user_id":"2088612472076052","channel_type":"UP_ALIPAY","fund_bill_list":"[{\"amount\":\"0.10\",\"fund_channel\":\"ALIPAYACCOUNT\"}]","merchant_no":"8000105202610","org_external_id":"2019112622001476051420837595","out_trade_no":"191126210912438233544500","pay_external_id":"12021199314159905411072","receipt_amount":"0.10","rsp_code":"0000","rsp_msg":"支付成功","sign":"c2fn9SVFwI0a5tE94S3rcVgdOfciRRYa5yiBmvs 0F3TjKKhbP7VdxLzHldgU7kq5FVvmgpH2zOiBoa/VGSYNBopoUo1nWI7yk3Fj1iFWgVliFdok28rW2Ha2eHK6jox6rKAVGphJi1RR//0Yb2yEkTOj0Np h6S8UrIE6UpcyE=","time_end":"2019-11-26 21:09:41","total_fee":"0.1","trx_external_id":"12021199314159901216768","trx_status":"SUCCESS","version":"1.0"}`), notify)

	fmt.Println(err)

	fmt.Printf("%+v", notify)
}
