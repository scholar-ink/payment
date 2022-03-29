package duolabao

import (
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"testing"
)

func TestCharge_Handle(t *testing.T) {
	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		AccessKey: "ecd1319dd3da43268ac96599a4ad232c1836f78b",
		SecretKey: "1f4b4ee723054facb34c8c16d188a4a2b3b53c1c",
	})

	ret, err := charge.Handle(&OrderChargeConf{
		AgentNum:    "10001015856358316531057",
		CustomerNum: "10001115872265337539380",
		ShopNum:     "10001215872266565541729",
		BankType:    "WX",
		RequestNum:  helper.CreateSn(),
		Amount:      "1",
		Source:      "API",
		CallbackUrl: "http://tq.udian.me/v1/common/enter-notify",
	})

	fmt.Printf("%+v", ret)
	fmt.Println(err)
}
