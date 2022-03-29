package notify

import (
	"fmt"
	"testing"
)

func TestHlbNotify_Handle(t *testing.T) {

	notify := new(HlbNotify)

	ret := map[string]string{
		"channelSettlementAmount":    "0.01",
		"productFee":                 "0.00",
		"realCreditAmount":           "0.01",
		"rt10_openId":                "oGex90rV1USKe2HeUuF0iEBx5LCM",
		"rt11_channelOrderNum":       "4200000374201909212817917049",
		"rt12_orderCompleteDate":     "2019-09-21 14:42:58",
		"rt13_onlineCardType":        "CREDIT",
		"rt14_cashFee":               "0.01",
		"rt15_couponFee":             "0.00",
		"rt17_outTransactionOrderId": "4200000374201909212817917049",
		"rt18_bankType":              "GDB_CREDIT",
		"rt1_customerNumber":         "E1801436315",
		"rt20_orderAttribute":        "UNDIRECT_DEFAULT",
		"rt23_paymentAmount":         "0.01",
		"rt24_creditAmount":          "0.01",
		"rt26_appPayType":            "WXPAY",
		"rt27_payType":               "SWIPE",
		"rt2_orderId":                "190921144247733693005420",
		"rt3_systemSerial":           "2196901692",
		"rt4_status":                 "SUCCESS",
		"rt5_orderAmount":            "0.01",
		"rt6_currency":               "CNY",
		"rt7_timestamp":              "1569048178178",
		"rt8_desc":                   "",
		"sign":                       "196ead884238bc14e0ec0f75011be1aa",
		"tradeType":                  "MICROPAY",
	}

	retData := notify.Handle(ret, func(merchantNo string) (md5Key string, err error) {
		return "HUflKaIvgRrpKLciHc2RpdUixcnIg4mS", nil
	}, func(data *HlbNotifyData) error {
		fmt.Printf("%+v", data)
		return nil
	})

	fmt.Println(retData)

}
