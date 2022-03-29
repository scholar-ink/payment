package adapay

import (
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"testing"
)

func TestCharge_Handle(t *testing.T) {

	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		ApiKey:     "api_live_90f01d3d-c94a-4026-924a-63584c7abd10",
		PrivateKey: "MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBALSAIqdFDAFWLm8HkPbYfOPmbM96qTKfR91B6tHR7FVixeicZbw11CWIw8O56dKQ0wotZTrWC75yvXpn2dTdkQgZp9cytXZ6MLnfiKvj+032wnAU3ZMESc/Onrhf589nWC6RvIy26tIQ0yUR5HLrpEN+A4qVDVqd00TQZc8YRoHtAgMBAAECgYA94uj+vNe+5ZOKEegMGnHHmcuY349/gckb/WvLgNQs+m6ssGLZQwN30wp74xReU7Vn+eSJZbYlGCYK/+xZ5ZXBw5uVhlDCmjWmFGVZVpiOibQStPqPFen4qhLMMsP8pZtJOagIjaGfMc9Wf36f/fJ3LFLXz3E7wjLsUmDslnDZKQJBAOLd2nu40Fw2FAVwEHd+5GbSkDWNw0HdlUbBHWV+fEUnLsiIE+Hw1nT6oOHfPMml5naKd8BLmkYUC2kfa9rFWh8CQQDLrgQTZ0bQtYjpzVkj3k+D1KnVuizy1LIQs4oZSMU0YBaVfUVLQZzHubTNpAMzH5hmfXg/fz+QKD0aSjRmTDpzAkBWHZqqrhvBdPGioshNY8h1U2ZUPcypeuAILJPpC9tGMLpsemL5t/7gBqb9Nk0Pyj6yLpuITepwwXkXXUsGjzVHAkEAhgBlxBJFX9ifTBsC03tWWwhV+Dw1iElxIVXNvJbIz42MLiutpDZ1nF1MW6LVTBQ0YvGXZEcmnYQrtxks4kSyiwJAWE+xShtaoOLI0mdOBMAI2amJWpBV4NCvHhVbWwh5F0Jxq3S79mXqvnWOsXBefjKZVFBCnqjlMFsExv6WEifCQg==",
	})

	ret, err := charge.Handle(&OrderChargeConf{
		AppId:      "app_691593e9-029e-49b4-b531-86ab229f8d94",
		OrderNo:    helper.CreateSn(),
		PayChannel: "wx_pub",
		//TimeExpire: "20200225135533",
		PayAmt:     fmt.Sprintf("%.2f", float64(1)/100),
		GoodsTitle: "subject",
		GoodsDesc:  "body",
		Currency:   "cny",
		NotifyUrl:  "http://tq.udian.me/v1/common/enter-notify",
		Expend: map[string]interface{}{
			"open_id": "oNEbm1CDWvqk7nYy7bhhc4hm-8Y8",
		},
	})

	fmt.Println(err)

	fmt.Println(ret)
}

func TestCharge_Refund(t *testing.T) {

	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		ApiKey:     "api_live_114ec65a-065d-4794-b697-7d4dcfd51317",
		PrivateKey: "MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBALSAIqdFDAFWLm8HkPbYfOPmbM96qTKfR91B6tHR7FVixeicZbw11CWIw8O56dKQ0wotZTrWC75yvXpn2dTdkQgZp9cytXZ6MLnfiKvj+032wnAU3ZMESc/Onrhf589nWC6RvIy26tIQ0yUR5HLrpEN+A4qVDVqd00TQZc8YRoHtAgMBAAECgYA94uj+vNe+5ZOKEegMGnHHmcuY349/gckb/WvLgNQs+m6ssGLZQwN30wp74xReU7Vn+eSJZbYlGCYK/+xZ5ZXBw5uVhlDCmjWmFGVZVpiOibQStPqPFen4qhLMMsP8pZtJOagIjaGfMc9Wf36f/fJ3LFLXz3E7wjLsUmDslnDZKQJBAOLd2nu40Fw2FAVwEHd+5GbSkDWNw0HdlUbBHWV+fEUnLsiIE+Hw1nT6oOHfPMml5naKd8BLmkYUC2kfa9rFWh8CQQDLrgQTZ0bQtYjpzVkj3k+D1KnVuizy1LIQs4oZSMU0YBaVfUVLQZzHubTNpAMzH5hmfXg/fz+QKD0aSjRmTDpzAkBWHZqqrhvBdPGioshNY8h1U2ZUPcypeuAILJPpC9tGMLpsemL5t/7gBqb9Nk0Pyj6yLpuITepwwXkXXUsGjzVHAkEAhgBlxBJFX9ifTBsC03tWWwhV+Dw1iElxIVXNvJbIz42MLiutpDZ1nF1MW6LVTBQ0YvGXZEcmnYQrtxks4kSyiwJAWE+xShtaoOLI0mdOBMAI2amJWpBV4NCvHhVbWwh5F0Jxq3S79mXqvnWOsXBefjKZVFBCnqjlMFsExv6WEifCQg==",
	})

	ret, err := charge.Refund(&OrderRefundConf{
		PaymentId:     "002112020051608310410107528482883801088",
		RefundOrderNo: helper.CreateSn(),
		RefundAmt:     "11.74",
		Reason:        "订单退款",
	})

	fmt.Println(err)

	fmt.Println(ret)
}
