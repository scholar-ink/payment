package chinaums

import (
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"testing"
	"time"
)

func TestCharge_Handle(t *testing.T) {
	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		MsgSrc:  "WWW.YTXXKEJI.COM",
		Mid:     "898310153993935",
		Tid:     "20124254",
		InstMid: "QRPAYDEFAULT",
		Md5Key:  "yw4C6MyKSY6ej4nh8jXA8CP43kDmTQxjr4mtWtdixnSP26Qi",
	})

	ret, err := charge.Handle(&OrderChargeConf{
		BillNo:      "4659" + helper.CreateSn(),
		BillDate:    time.Now().Format("2006-01-02"),
		TotalAmount: "1",
		BillDesc:    charge.MsgId + helper.CreateSn() + "-收钱啦",
		NotifyUrl:   "http://tq.udian.me/v1/common/enter-notify",
		ReturnUrl:   "http://www.baidu.com",
	})

	fmt.Printf("%+v", ret)
	fmt.Println(err)
}

func TestCharge_Refund(t *testing.T) {
	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		MsgSrc:  "WWW.JSLBWL.COM",
		Mid:     "898130448161280",
		Tid:     "04938774",
		InstMid: "QRPAYDEFAULT",
		Md5Key:  "Xjw5d8fM587DMmMBHWC8DmfSZZcpENdtr7haFRXJYG2pzx87",
	})

	ret, err := charge.Refund(&OrderRefundConf{
		BillDate:      time.Now().Format("2006-01-02"),
		BillNo:        "5749200305235504283729854850",
		RefundOrderId: "5749" + helper.CreateSn(),
		RefundAmount:  "1",
	})

	fmt.Printf("%+v", ret)
	fmt.Println(err)
}

func TestCharge_Query(t *testing.T) {
	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		MsgSrc:  "WWW.JSLBWL.COM",
		Mid:     "898130448161496",
		Tid:     "04939424",
		InstMid: "QRPAYDEFAULT",
		Md5Key:  "Xjw5d8fM587DMmMBHWC8DmfSZZcpENdtr7haFRXJYG2pzx87",
	})

	ret, err := charge.Query(&OrderQueryConf{
		BillDate: "2020-04-16",
		BillNo:   "5749200417102612325491570164",
	})

	fmt.Printf("%+v", ret)
	fmt.Println(err)
}
