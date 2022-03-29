package notify

import (
	"fmt"
	"testing"
)

func TestCmNotify_Handle(t *testing.T) {

	notify := new(CmNotify)

	ret := map[string]string{
		"channel_trade_no": "582020021822001476051448179711",
		"device":           "",
		"money":            "0.01",
		"orderId":          "200218111343495217795242",
		"out_trade_no":     "20200218111153994068",
		"result_code":      "SUCCESS",
		"shopId":           "f13db1959ed46e3cb1dfef8ea88bace5",
		"shortKey":         "",
		"sign":             "9288DBF65CBADC7ED31DB2C98813C64B",
		"timeEnd":          "20200218111420",
		"transaction_id":   "1023000023200218111414r8wigtpiDx",
		"type":             "alipay",
	}

	retData := notify.Handle(ret, func(merchantNo string) (md5Key string, err error) {
		return "4EFF68ACE1BC890CEAB068A4E8176503", nil
	}, func(data *CmNotifyData) error {

		fmt.Println(data.ResultCode)
		fmt.Printf("%+v\n", data)
		return nil
	})

	fmt.Println(retData)

}
