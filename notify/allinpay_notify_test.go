package notify

import (
	"fmt"
	"testing"
)

func TestAllinPayNotify_Handle(t *testing.T) {

	notify := new(AllInPayNotify)

	ret := map[string]string{
		"acct":       "opn0buPOa_cZleu_dZQvPcPe2ejs",
		"accttype":   "99",
		"amount":     "1",
		"appid":      "00176476",
		"bizseq":     "190920173815592128629675",
		"cusid":      "56045204816B7ZA",
		"fee":        "0",
		"paytime":    "20190920173950",
		"randomstr":  "293303",
		"sign":       "654f2a079d10ab36fbacc76a74acd1e2",
		"signtype":   "MD5",
		"termauthno": "CFT",
		"termid":     "TTWEPMGL",
		"termrefnum": "4200000414201909208615521701",
		"timestamp":  "20190920173950",
		"traceno":    "0",
		"trxcode":    "VSP501",
		"trxday":     "20190920",
		"trxid":      "121930210000039902",
		"trxstatus":  "0000",
	}

	retData := notify.Handle(ret, func(merchantNo string) (md5Key string, err error) {
		return "9689489231d792a260e2559586276916", nil
	}, func(data *AllInPayNotifyData) error {
		fmt.Printf("%+v", data)
		return nil
	})

	fmt.Println(retData)

}
