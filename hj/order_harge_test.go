package hj

import (
	"encoding/json"
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"testing"
)

func TestCharge_Handle(t *testing.T) {
	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		Md5Key: "a5095968f9ec42f3889add3297c4e77b",
	})

	sharingParams := []*SharingParam{
		{
			AltAmount: "0.01",
			AltMchNo:  "333108900001294",
			IsGuar:    "12",
		},
	}

	b, _ := json.Marshal(sharingParams)

	ret, err := charge.Handle(&OrderChargeConf{
		P1MerchantNo:  "888108800008622",
		P2OrderNo:     helper.CreateSn(),
		P3Amount:      "0.02",
		P5ProductName: "测试商品",
		P9NotifyUrl:   "http://tq.udian.me/v1/common/enter-notify",
		Q1FrpCode:     "WEIXIN_NATIVE",
		//Q7AppId:"wx6c4cc850defdb695",
		QcIsAlt:           "1",
		QdAltType:         "11",
		QeAltInfo:         string(b),
		QaTradeMerchantNo: "777107200047200",
	})

	fmt.Printf("%+v", ret)
	fmt.Println(err)
}
