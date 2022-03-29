package hlb

import (
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/strings"
	"testing"
)

func TestCharge_Handle(t *testing.T) {
	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		Md5Key: "HUflKaIvgRrpKLciHc2RpdUixcnIg4mS",
	})

	ret, err := charge.Handle(&OrderChargeConf{
		P1BizType:        "AppPay",
		P2OrderId:        helper.CreateSn(),
		P3CustomerNumber: "E1801457912",
		P4PayType:        "SCAN",
		P5OrderAmount:    strings.FormatFloat(float64(100000) / 100),
		P6Currency:       "CNY",
		P7Authcode:       "",
		P8AppType:        "ALIPAY",
		P9NotifyUrl:      "http://tq.udian.me/v1/common/enter-notify",
		P11OrderIp:       "127.0.0.1",
		P12GoodsName:     "测试商品",
	})

	fmt.Printf("%+v", ret)
	fmt.Println(err)
}
