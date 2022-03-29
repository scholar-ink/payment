package minpay

import (
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"io/ioutil"
	"testing"
	"time"
)

func TestCharge_Handle(t *testing.T) {
	charge := new(OrderCharge)

	key, _ := ioutil.ReadFile("minfu2.pfx")

	charge.InitBaseConfig(&BaseConfig{
		MerchantNo:   "872320879935001",
		Key:          key,
		PubKey:       "MIIE2zCCA8OgAwIBAgIFQhcmdHUwDQYJKoZIhvcNAQELBQAwWDELMAkGA1UEBhMCQ04xMDAuBgNVBAoMJ0NoaW5hIEZpbmFuY2lhbCBDZXJ0aWZpY2F0aW9uIEF1dGhvcml0eTEXMBUGA1UEAwwOQ0ZDQSBBQ1MgT0NBMzEwHhcNMTkxMDE2MTMwNDA5WhcNMjQxMDE2MTMwNDA5WjCBiTELMAkGA1UEBhMCQ04xFzAVBgNVBAoMDkNGQ0EgQUNTIE9DQTMxMRAwDgYDVQQLDAdmYXN0UGF5MRkwFwYDVQQLDBBPcmdhbml6YXRpb25hbC0xMTQwMgYDVQQDDCtmYXN0UGF5QElQU+acjeWKoeWZqOetvuWQjeivgeS5pkBaTVBTMDAwMUAxMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsnYjIh3O1hPv0GFf5hxBojnDnvoT4zYO1NEMceZrcBu4TqaSYuN7XFu3AKHwOwnyxbcSJOqGzo/i3CIo3gbssN/cQSGR8tBnOTRSwOPJQiS8BzOLt+qxAT1RjIypueGrFfHPH30Se18H9PJhbmcVlbkXlqypn87B3gxdG9CcW4Gj0AnV71ttg2bEQ0plcVCJlPPzCMOUNmBUPoU2uVPwmLiVP2mMcuiAaFH29hjIXsivUG+0Az+/vsIboqZhZMBY1ElsFgITYl/XQK1HqM/ESrRTYw//bUV+EdebG3rAztuAu4rO9hMSwp2ooT6JNGOApIjdYzRsO/iBAPsG3ifzmQIDAQABo4IBeDCCAXQwbAYIKwYBBQUHAQEEYDBeMCgGCCsGAQUFBzABhhxodHRwOi8vb2NzcC5jZmNhLmNvbS5jbi9vY3NwMDIGCCsGAQUFBzAChiZodHRwOi8vY3JsLmNmY2EuY29tLmNuL29jYTMxL29jYTMxLmNlcjAfBgNVHSMEGDAWgBTitAnLzWGhc0p5f/GKgwvdtH6MHTAMBgNVHRMBAf8EAjAAMEgGA1UdIARBMD8wPQYIYIEchu8qAQQwMTAvBggrBgEFBQcCARYjaHR0cDovL3d3dy5jZmNhLmNvbS5jbi91cy91cy0xNC5odG0wPQYDVR0fBDYwNDAyoDCgLoYsaHR0cDovL2NybC5jZmNhLmNvbS5jbi9vY2EzMS9SU0EvY3JsMjAzMi5jcmwwDgYDVR0PAQH/BAQDAgbAMB0GA1UdDgQWBBRQfnYHki0LNHuaiLoo7jQcWdQJ6TAdBgNVHSUEFjAUBggrBgEFBQcDAgYIKwYBBQUHAwQwDQYJKoZIhvcNAQELBQADggEBALkM+QidmpzWFENa6DCKS9pdKpvggWBzCa2mD1/QCt3H48Q0DV8/wsoN6srsaYzg5pOkS9+hULKdnCOKzhdy+6JA7kabaqCqrqT+Nz1kSs9yrQMR5M/2K+lj///+qzM44M6d/rqBerwzYP9xgsrpJ+NIT/WwtdpXjE4i419/11D7XayUIcd0XlxIrDF44VDB4osxtkjtc5ropTExZKgRRmt3Wz/7tvsezdAhdJXnrov3FHFDHNNEkjRzH0VVx68fzwCnuNtl/6w7LCGvdGPN0NW1P7UL29IuLxInCW62KcPUiGh0beYGE0NSDs4aU++3oiIhkjeOTIjLVN7n02DbLQI=",
		CertPassword: "584520",
	})

	ret, err := charge.Handle(&OrderChargeConf{
		MerchantRequestId: helper.CreateSn(),
		OrderAmount:       "0.01",
		GoodsName:         "测试商品",
		GoodsDesc:         "测试商品",
		PayProduct:        "GATHER_USCAN",
		CreateTime:        time.Now().Format("20060102150405"),
		//BuyerLogonId:      "13585821080",
		//NonceStr:          helper.NonceStr(),
		//UserIp:            "219.152.110.53",
		//MerchantIp:        "47.102.138.87",
		CallBackUrl: "http://tq.udian.me/v1/common/enter-notify",
	})

	fmt.Printf("%+v", ret)
	fmt.Println(err)
}
