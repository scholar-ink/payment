package xunlian

import (
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"testing"
)

func TestCharge_Handle(t *testing.T) {
	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		MerchantId: "676366148160010",
		InScd:      "91681888",
		TerminalId: "",
		Key:        "a429257e08683c814a620a3f43b9213c",
	})

	ret, err := charge.Handle(&OrderChargeConf{
		OrderNum: helper.CreateSn(),
		ChCd:     "ALP",
		TxAmt:    fmt.Sprintf("%012d", 1),
		Subject:  "测试商品",
		//PayLimit:   "credit",
		BackUrl: "http://tq.udian.me/v1/common/enter-notify",
		//FrontUrl:    "http://www.cardinfolink.com/",
		//TimeStart:  "20170909120000",
		//TimeExpire: "20170909130000",
	})

	fmt.Println(err)

	fmt.Println(ret)
}

func TestOrderCharge_Query(t *testing.T) {
	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		MerchantId: "100000000000203",
		InScd:      "10130001",
		TerminalId: "00000001",
		Key:        "zsdfyreuoyamdphhaweyrjbvzkgfdycs",
	})

	ret, err := charge.Query(&QueryConf{
		OrigOrderNum: "200303211334631929516436",
	})

	fmt.Println(err)

	fmt.Println(ret)
}

func TestOrderCharge_Refund(t *testing.T) {
	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		MerchantId: "168652650720001",
		InScd:      "91681888",
		TerminalId: "yunc0001",
		Key:        "4b6b1d37707cc839e7c08deac27858fe",
	})

	ret, err := charge.Refund(&RefundConf{
		OrigOrderNum: "200307130514500350262074",
		OrderNum:     helper.CreateSn(),
		TxAmt:        fmt.Sprintf("%012d", 1),
	})

	fmt.Println(err)

	fmt.Println(ret)
}
