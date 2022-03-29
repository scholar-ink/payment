package notify

import (
	"fmt"
	"net/url"
	"testing"
)

func TestHjNotify_Handle(t *testing.T) {

	values, err := url.QueryUnescape("2019-10-26+20%3A42%3A11")

	fmt.Println(values)

	fmt.Println(err)
	//
	return

	notify := new(HjNotify)

	ret := map[string]string{
		"r1_MerchantNo":  "888108800008622",
		"r2_OrderNo":     "191026204118951409596758",
		"r3_Amount":      "0.10",
		"r4_Cur":         "1",
		"r5_Mp":          "",
		"r6_Status":      "100",
		"r7_TrxNo":       "100219102687279951",
		"r8_BankOrderNo": "100219102687279951",
		"r9_BankTrxNo":   "152019102622001476051406515293",
		"ra_PayTime":     "2019-10-26 20:42:11",
		"rb_DealTime":    "2019-10-26 20:42:11",
		"rc_BankCode":    "ALIPAY_NATIVE",
		"hmac":           "d12fa12e6571f1d2d0c8c85e62ed5d2e",
	}

	retData := notify.Handle(ret, func(merchantNo string) (md5Key string, err error) {
		return "a5095968f9ec42f3889add3297c4e77b", nil
	}, func(data *HjNotifyData) error {
		fmt.Printf("%+v\n", data)
		return nil
	})

	fmt.Println(retData)

}
