package yty

import (
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"testing"
	"time"
)

func TestCharge_Handle(t *testing.T) {
	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		AgentNo:    "88000193155608",
		MerchantNo: "66000193303313",
		Key:        "MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBALSAIqdFDAFWLm8HkPbYfOPmbM96qTKfR91B6tHR7FVixeicZbw11CWIw8O56dKQ0wotZTrWC75yvXpn2dTdkQgZp9cytXZ6MLnfiKvj+032wnAU3ZMESc/Onrhf589nWC6RvIy26tIQ0yUR5HLrpEN+A4qVDVqd00TQZc8YRoHtAgMBAAECgYA94uj+vNe+5ZOKEegMGnHHmcuY349/gckb/WvLgNQs+m6ssGLZQwN30wp74xReU7Vn+eSJZbYlGCYK/+xZ5ZXBw5uVhlDCmjWmFGVZVpiOibQStPqPFen4qhLMMsP8pZtJOagIjaGfMc9Wf36f/fJ3LFLXz3E7wjLsUmDslnDZKQJBAOLd2nu40Fw2FAVwEHd+5GbSkDWNw0HdlUbBHWV+fEUnLsiIE+Hw1nT6oOHfPMml5naKd8BLmkYUC2kfa9rFWh8CQQDLrgQTZ0bQtYjpzVkj3k+D1KnVuizy1LIQs4oZSMU0YBaVfUVLQZzHubTNpAMzH5hmfXg/fz+QKD0aSjRmTDpzAkBWHZqqrhvBdPGioshNY8h1U2ZUPcypeuAILJPpC9tGMLpsemL5t/7gBqb9Nk0Pyj6yLpuITepwwXkXXUsGjzVHAkEAhgBlxBJFX9ifTBsC03tWWwhV+Dw1iElxIVXNvJbIz42MLiutpDZ1nF1MW6LVTBQ0YvGXZEcmnYQrtxks4kSyiwJAWE+xShtaoOLI0mdOBMAI2amJWpBV4NCvHhVbWwh5F0Jxq3S79mXqvnWOsXBefjKZVFBCnqjlMFsExv6WEifCQg==",
		PubKey:     "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCCyhAPSfMB25xMqKyqzKDjhLaUMSkd47W3XUPHq/ercW8K27GOzgG7ifULsuaBpLbq+cEz/MEb2CiIFPkD1xrR34wkWB709MvPuPdIxxQoy6cVezfpmaTeEFczKctmdJHcD2foMv062il85V5+BJfMa36w14r0v1v+sZZz2ZbEGwIDAQAB",
	})

	ret, err := charge.Handle(&OrderChargeConf{
		OutOrderNo:  helper.CreateSn(),
		TotalAmount: "1",
		PayType:     "WE_CHAT",
		OrderBody:   "测试商品",
		PayTime:     time.Now().Format("20060102150405"),
		NotifyUrl:   "http://tq.udian.me/v1/common/enter-notify",
	})

	fmt.Printf("%+v", ret)
	fmt.Println(err)
}
