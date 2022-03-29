package notify

import (
	"fmt"
	"testing"
)

func TestChinaUmsNotify_Handle(t *testing.T) {

	notify := new(ChinaUmsNotify)

	ret := map[string]interface{}{
		"EG":            "kszG",
		"billDate":      "2020-03-05",
		"billDesc":      "200305131400023239572888-\346\224\266\351\222\261\345\225\246",
		"billNo":        "57492003052061420136888532",
		"billPayment":   "{\"buyerUsername\":\"135****1080\",\"payTime\":\"2020-03-05 13:14:24\",\"paySeqId\":\"13361789237N\",\"invoiceAmount\":1,\"settleDate\":\"2020-03-05\",\"buyerId\":\"2088612472076052\",\"receiptAmount\":1,\"totalAmount\":1,\"couponAmount\":0,\"billBizType\":\"bills\",\"buyerPayAmount\":1,\"targetOrderId\":\"2020030522001476050549056120\",\"payDetail\":\"\346\224\257\344\273\230\345\256\235\344\275\231\351\242\235\346\224\257\344\273\2300.01\345\205\203\343\200\202\",\"merOrderId\":\"574920030520614201368885320\",\"status\":\"TRADE_SUCCESS\",\"targetSys\":\"Alipay 2.0\"}",
		"billQRCode":    "https://qr.95516.com/48020000/57492003055431400131756001",
		"billStatus":    "PAID",
		"createTime":    "2020-03-05 13:14:20",
		"instMid":       "QRPAYDEFAULT",
		"mchntUuid":     "e18b904b9cf340aca4ff6e9a776c8cfc",
		"merName":       "\347\233\261\347\234\231\346\200\235\345\215\216\346\224\266\345\221\227\344\277\241\346\201\257\346\212\200\346\234\257\346\234\211\351\231\220\345\205\254\345\217\270",
		"mid":           "898130448161280",
		"notifyId":      "cc99bc64-adf9-4592-a937-4f6ba764cda8",
		"qrCodeId":      "57492003055431400131756001",
		"qrCodeType":    "FIXEDPAY",
		"receiptAmount": "1",
		"seqId":         "13361789237N",
		"sign":          "357550E8A13FDEFBB6F1F2779A8E4754",
		"subInst":       "102800",
		"tid":           "04938774",
		"totalAmount":   "1",
	}

	retData := notify.Handle(ret, func(merchantNo string) (md5Key string, err error) {
		return "Xjw5d8fM587DMmMBHWC8DmfSZZcpENdtr7haFRXJYG2pzx87", nil
	}, func(data *ChinaUmsNotifyData) error {
		fmt.Printf("%+v\n", data)
		fmt.Printf("%+v\n", data.BillPayment.TargetOrderId)
		return nil
	})

	fmt.Println(retData)

}
