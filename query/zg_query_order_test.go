package query

import (
	"fmt"
	"github.com/scholar-ink/payment/gpay"
	"testing"
)

func TestZgOrderQuery_Handle(t *testing.T) {
	for {
		query := new(ZgOrderQuery)

		query.InitBaseConfig(&gpay.BaseConfig{
			MerchantNo: "653068",
			Md5Key:     "2b8fa9a37cc74a01972046e3b22386f4",
		})
		ret, err := query.Handle(&ZgOrderQueryConf{
			OutTradeNo: "190725172109458619243815",
		})

		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(ret)

		//fmt.Println(ret.PayStatus)

	}

}
