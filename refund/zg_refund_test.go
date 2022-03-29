package refund

import (
	"fmt"
	"github.com/scholar-ink/payment/gpay"
	"testing"
)

func TestZgRefund_Handle(t *testing.T) {
	refund := new(ZgRefund)

	refund.InitBaseConfig(&gpay.BaseConfig{
		MerchantNo: "751627",
		Md5Key:     "76c8c058cceb46b793cbd31cf9bed007",
	})

	refundNo, err := refund.Handle(&ZgRefundConf{
		//OrderNo:"190819120755834951672954",
		OutTradeNo:  "190819120755834951672954",
		OutRefundNo: "190328192530960395452694120",
		TotalFee:    100,
		RefundFee:   10,
	})

	fmt.Println(refundNo)

	fmt.Println(err)
}
