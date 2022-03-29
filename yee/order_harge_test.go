package yee

import (
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"testing"
)

func TestCharge_Handle(t *testing.T) {
	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		ParentMerchantNo: "10033308281",
		AppKey:           "OPR:10033308281",
		PrivateKey:       "MIIEpAIBAAKCAQEAuRG2HsqR2kXwDtQbCZNtravarGNttzbh8q4SuVAfNb5G7mpf4aKrscS6+P1YQxSsIHXANbm1d1AdtXw9DZd1zZQwqkI63kDO0UcKjSlnsSso4MH3x4YPhHNfXapxotLs9uB8WeMdBvZ9viK1bGJ9IlIm/G08dKi5w90OXAlmY6mGSlSM8PQKFZ4xDvxsx/eVIIkvZyRYSdIYvaGU4C5yep0SDHmtdzu+l3KrnwmuYruDTEgevCXIwNFeFuuxmqTJpxbovLzLQcYfzw5VgrzMQKo81ihRa2TgbmpPqHvz8Ecpv2SD3oZcqJJ+V9Fekn8ZwRii01kZRxmtO6wHJiJA0wIDAQABAoIBAQChyUX4rQXMVw+BJxNrz6I2DOJhiZpEbIoh6OMddVmTVgAUNJIVYmvOQDljqbYbDltbFRUu4mYtI7CVE0McOqgVS9MvRC7KVMV4Vi53MRcg3qYhte+yURQHqgRYkcQ9nz7go2aR/eVGTP2n1hfb5E2YT1EdozESmt2qx/jhpKYJwJCx/Bbt0Puk7tAAFjF+FlfVJBCjaotRaDEa3wbRBloIgHTz7lKMmeIWwDP5APAqTsMjQVTgrUmFbwRy8wlFrS9eHsXvHUNSz40lTbQZYzjCDyAWsnr+d7UJnA1y1OKlW/L/7iA5tZlPooDAcZQrw/2L6eqcIfli6qJS5xp7jPYhAoGBAN1lEL8N+VH4ZlaAIU6JcwhdGP7pqVNXm1E/OHCEuAUrXNUUXPsYzSVsonm/iwK5DBqakVQauHzELF2i1Vq2L2ZFIWzGmXRbOStD25WlRAmXdaR7tmyk1wUbi4pGLcPsZ0T83p7h5jYcU20xqdRWIbFzAd0zhgmLIivgd3/Izco3AoGBANX/GzEg+3/Ezjxm0+PngtNP1XhBBMZfLA/Za9VWxkoRhet7nA0ZS8TdYCucqYsi/xT3Zm6uW4nsfBA6EppX36iS0Ko1kvu/BRB52S4KX2N2Nprv0aI2BKUt8GFnzMu3vYoWwU4HfAPZo4gHhujQyNIGva4O3nU3FA+17IdaS0BFAoGBAMi/WHT01cqW4yyyGAFfrpe52u1hsDCq8mG0YpfcMAQ38oAfa8QfE/1ISPb+UK0SX8BLwVUyuXAgbV4mRTFwmwAv1QQN/J0+DlOFvzks1smftYOEzcAro/C0rk2eHudVl7o9VBtbGGSeQKN1cdngk8KUVu9dRb+nkj+Y1KJg0QD3AoGAUoUlPzSwxvxlavVcEC5eTI4ki1LHPJfGrfXxlzr3C+jl84CfFI4Eoc2cIDUxS+4a30LnxRaHRhBjZv593fa00JSM0pYGL/3hPhE+pnppfjk/pU+FTz/1Wpz0bRtR5dzcwjs0H5rTP8jVPsdoRq78QcFHs68YKasrmUNWCnvPOOECgYBzF69wXWrRK0aBa+5RdjI7oED/7i0l8urRn1Xf3+1KSNbXDLPy/PYkAtpcf6tNBY7n3utFiEZrkc1FM1v74KDmuSij7XPpmIyldVkIwy1/Xx/fRvBSuIVHpbS8FZEUqlCg7B5LB6rFTYyzGOtoXy2A3EVVa5vfsWns9ukY7VPLKA==",
	})

	ret, err := charge.Handle(&OrderChargeConf{
		MerchantNo:     "10033406581",
		OrderId:        helper.CreateSn(),
		OrderAmount:    "0.01",
		PayTool:        "SCCANPAY",
		PayType:        "WECHAT",
		AppId:          "wxdae782e8546f5bb6",
		OpenId:         "oNEbm1CDWvqk7nYy7bhhc4hm-8Y8",
		TimeoutExpress: "86400",
		NotifyUrl:      "http://tq.udian.me/v1/common/enter-notify",
		GoodsParamExt:  &GoodsParamExt{GoodsName: "测试商品", GoodsDesc: "商品描述"},
		HMacKey:        "JAP85h741554A9rk3217o4Xs0rnT9754QM1143ww61PH4426u74Q9e90R65B",
	})

	fmt.Println(err)

	fmt.Println(ret)
}

func TestOrderCharge_Refund(t *testing.T) {
	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		ParentMerchantNo: "10033308281",
		AppKey:           "OPR:10033308281",
		PrivateKey:       "MIIEpAIBAAKCAQEAuRG2HsqR2kXwDtQbCZNtravarGNttzbh8q4SuVAfNb5G7mpf4aKrscS6+P1YQxSsIHXANbm1d1AdtXw9DZd1zZQwqkI63kDO0UcKjSlnsSso4MH3x4YPhHNfXapxotLs9uB8WeMdBvZ9viK1bGJ9IlIm/G08dKi5w90OXAlmY6mGSlSM8PQKFZ4xDvxsx/eVIIkvZyRYSdIYvaGU4C5yep0SDHmtdzu+l3KrnwmuYruDTEgevCXIwNFeFuuxmqTJpxbovLzLQcYfzw5VgrzMQKo81ihRa2TgbmpPqHvz8Ecpv2SD3oZcqJJ+V9Fekn8ZwRii01kZRxmtO6wHJiJA0wIDAQABAoIBAQChyUX4rQXMVw+BJxNrz6I2DOJhiZpEbIoh6OMddVmTVgAUNJIVYmvOQDljqbYbDltbFRUu4mYtI7CVE0McOqgVS9MvRC7KVMV4Vi53MRcg3qYhte+yURQHqgRYkcQ9nz7go2aR/eVGTP2n1hfb5E2YT1EdozESmt2qx/jhpKYJwJCx/Bbt0Puk7tAAFjF+FlfVJBCjaotRaDEa3wbRBloIgHTz7lKMmeIWwDP5APAqTsMjQVTgrUmFbwRy8wlFrS9eHsXvHUNSz40lTbQZYzjCDyAWsnr+d7UJnA1y1OKlW/L/7iA5tZlPooDAcZQrw/2L6eqcIfli6qJS5xp7jPYhAoGBAN1lEL8N+VH4ZlaAIU6JcwhdGP7pqVNXm1E/OHCEuAUrXNUUXPsYzSVsonm/iwK5DBqakVQauHzELF2i1Vq2L2ZFIWzGmXRbOStD25WlRAmXdaR7tmyk1wUbi4pGLcPsZ0T83p7h5jYcU20xqdRWIbFzAd0zhgmLIivgd3/Izco3AoGBANX/GzEg+3/Ezjxm0+PngtNP1XhBBMZfLA/Za9VWxkoRhet7nA0ZS8TdYCucqYsi/xT3Zm6uW4nsfBA6EppX36iS0Ko1kvu/BRB52S4KX2N2Nprv0aI2BKUt8GFnzMu3vYoWwU4HfAPZo4gHhujQyNIGva4O3nU3FA+17IdaS0BFAoGBAMi/WHT01cqW4yyyGAFfrpe52u1hsDCq8mG0YpfcMAQ38oAfa8QfE/1ISPb+UK0SX8BLwVUyuXAgbV4mRTFwmwAv1QQN/J0+DlOFvzks1smftYOEzcAro/C0rk2eHudVl7o9VBtbGGSeQKN1cdngk8KUVu9dRb+nkj+Y1KJg0QD3AoGAUoUlPzSwxvxlavVcEC5eTI4ki1LHPJfGrfXxlzr3C+jl84CfFI4Eoc2cIDUxS+4a30LnxRaHRhBjZv593fa00JSM0pYGL/3hPhE+pnppfjk/pU+FTz/1Wpz0bRtR5dzcwjs0H5rTP8jVPsdoRq78QcFHs68YKasrmUNWCnvPOOECgYBzF69wXWrRK0aBa+5RdjI7oED/7i0l8urRn1Xf3+1KSNbXDLPy/PYkAtpcf6tNBY7n3utFiEZrkc1FM1v74KDmuSij7XPpmIyldVkIwy1/Xx/fRvBSuIVHpbS8FZEUqlCg7B5LB6rFTYyzGOtoXy2A3EVVa5vfsWns9ukY7VPLKA==",
	})

	ret, err := charge.Refund(&OrderRefundConf{
		MerchantNo:      "10033343784",
		OrderId:         "200305221526827369904340",
		RefundRequestId: helper.CreateSn(),
		UniqueOrderNo:   "1001202003050000001459886237",
		RefundAmount:    "0.01",
		HMacKey:         "d5F41523xwDT6Sp1gp68EdT1x7Z6f36V81q906b7D32pF98Xa903e0i4q63j",
	})

	fmt.Println(err)

	fmt.Println(ret)

}
