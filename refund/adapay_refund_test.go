package refund

import (
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"testing"
)

func TestAdaRefund_Handle(t *testing.T) {
	refund := new(AdaPayRefund)

	refund.InitConfig(&AdaPayConfig{
		ApiKey:     "api_live_5586a3c7-6c17-4258-82f2-ff6a02f73e21",
		PrivateKey: "MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBALSAIqdFDAFWLm8HkPbYfOPmbM96qTKfR91B6tHR7FVixeicZbw11CWIw8O56dKQ0wotZTrWC75yvXpn2dTdkQgZp9cytXZ6MLnfiKvj+032wnAU3ZMESc/Onrhf589nWC6RvIy26tIQ0yUR5HLrpEN+A4qVDVqd00TQZc8YRoHtAgMBAAECgYA94uj+vNe+5ZOKEegMGnHHmcuY349/gckb/WvLgNQs+m6ssGLZQwN30wp74xReU7Vn+eSJZbYlGCYK/+xZ5ZXBw5uVhlDCmjWmFGVZVpiOibQStPqPFen4qhLMMsP8pZtJOagIjaGfMc9Wf36f/fJ3LFLXz3E7wjLsUmDslnDZKQJBAOLd2nu40Fw2FAVwEHd+5GbSkDWNw0HdlUbBHWV+fEUnLsiIE+Hw1nT6oOHfPMml5naKd8BLmkYUC2kfa9rFWh8CQQDLrgQTZ0bQtYjpzVkj3k+D1KnVuizy1LIQs4oZSMU0YBaVfUVLQZzHubTNpAMzH5hmfXg/fz+QKD0aSjRmTDpzAkBWHZqqrhvBdPGioshNY8h1U2ZUPcypeuAILJPpC9tGMLpsemL5t/7gBqb9Nk0Pyj6yLpuITepwwXkXXUsGjzVHAkEAhgBlxBJFX9ifTBsC03tWWwhV+Dw1iElxIVXNvJbIz42MLiutpDZ1nF1MW6LVTBQ0YvGXZEcmnYQrtxks4kSyiwJAWE+xShtaoOLI0mdOBMAI2amJWpBV4NCvHhVbWwh5F0Jxq3S79mXqvnWOsXBefjKZVFBCnqjlMFsExv6WEifCQg==",
	})

	refundNo, err := refund.Handle(&AdaPayRefundConf{
		PaymentId:     "002112020022502290410078083968434995200",
		RefundOrderNo: helper.CreateSn(),
		RefundAmt:     "0.01",
	})

	fmt.Println(refundNo)

	fmt.Println(err)
}
