package allinpay

import (
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"testing"
)

func TestCharge_Handle(t *testing.T) {
	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		AppId: "00185542",
		C:     "TTW0SaNR",
		Key:   "9689489231d792a260e2559586276916",
	})

	ret, err := charge.Handle(&OrderChargeConf{
		Oid:        helper.CreateSn(),
		Amt:        2,
		TrxReserve: "05|S1#" + helper.CreateSn() + "|S2#",
	})

	fmt.Printf("%+v", ret)
	fmt.Println(err)
}
