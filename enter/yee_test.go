package enter

import (
	"fmt"
	"github.com/scholar-ink/payment/util/http"
	"testing"
)

func TestYeeEnter_Create(t *testing.T) {
	yee := new(YeeEnter)

	yee.InitConfig(&YeeConfig{
		ParentMerchantNo: "10033308281",
		AppKey:           "OPR:10033308281",
		PrivateKey:       "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC5EbYeypHaRfAO1BsJk22tq9qsY223NuHyrhK5UB81vkbual/hoquxxLr4/VhDFKwgdcA1ubV3UB21fD0Nl3XNlDCqQjreQM7RRwqNKWexKyjgwffHhg+Ec19dqnGi0uz24HxZ4x0G9n2+IrVsYn0iUib8bTx0qLnD3Q5cCWZjqYZKVIzw9AoVnjEO/GzH95UgiS9nJFhJ0hi9oZTgLnJ6nRIMea13O76XcqufCa5iu4NMSB68JcjA0V4W67GapMmnFui8vMtBxh/PDlWCvMxAqjzWKFFrZOBuak+oe/PwRym/ZIPehlyokn5X0V6SfxnBGKLTWRlHGa07rAcmIkDTAgMBAAECggEBAKHJRfitBcxXD4EnE2vPojYM4mGJmkRsiiHo4x11WZNWABQ0khVia85AOWOpthsOW1sVFS7iZi0jsJUTQxw6qBVL0y9ELspUxXhWLncxFyDepiG177JRFAeqBFiRxD2fPuCjZpH95UZM/afWF9vkTZhPUR2jMRKa3arH+OGkpgnAkLH8Fu3Q+6Tu0AAWMX4WV9UkEKNqi1FoMRrfBtEGWgiAdPPuUoyZ4hbAM/kA8CpOwyNBVOCtSYVvBHLzCUWtL14exe8dQ1LPjSVNtBljOMIPIBayev53tQmcDXLU4qVb8v/uIDm1mU+igMBxlCvD/Yvp6pwh+WLqolLnGnuM9iECgYEA3WUQvw35UfhmVoAhTolzCF0Y/umpU1ebUT84cIS4BStc1RRc+xjNJWyieb+LArkMGpqRVBq4fMQsXaLVWrYvZkUhbMaZdFs5K0PblaVECZd1pHu2bKTXBRuLikYtw+xnRPzenuHmNhxTbTGp1FYhsXMB3TOGCYsiK+B3f8jNyjcCgYEA1f8bMSD7f8TOPGbT4+eC00/VeEEExl8sD9lr1VbGShGF63ucDRlLxN1gK5ypiyL/FPdmbq5biex8EDoSmlffqJLQqjWS+78FEHnZLgpfY3Y2mu/RojYEpS3wYWfMy7e9ihbBTgd8A9mjiAeG6NDI0ga9rg7edTcUD7Xsh1pLQEUCgYEAyL9YdPTVypbjLLIYAV+ul7na7WGwMKryYbRil9wwBDfygB9rxB8T/UhI9v5QrRJfwEvBVTK5cCBtXiZFMXCbAC/VBA38nT4OU4W/OSzWyZ+1g4TNwCuj8LSuTZ4e51WXuj1UG1sYZJ5Ao3Vx2eCTwpRW711Fv6eSP5jUomDRAPcCgYBShSU/NLDG/GVq9VwQLl5MjiSLUsc8l8at9fGXOvcL6OXzgJ8UjgShzZwgNTFL7hrfQufFFodGEGNm/n3d9rTQlIzSlgYv/eE+ET6meml+OT+lT4VPP/VanPRtG1Hl3NzCOzQfmtM/yNU+x2hGrvxBwUezrxgpqyuZQ1YKe8844QKBgHMXr3BdatErRoFr7lF2MjugQP/uLSXy6tGfVd/f7UpI1tcMs/L89iQC2lx/q00Fjufe60WIRmuRzUUzW/vgoOa5KKPtc+mYjKV1WQjDLX9fH99G8FK4hUeltLwVkRSqUKDsHksHqsVNjLMY62hfLYDcRVVrm9+xaez26RjtU8so",
	})

	ret, err := yee.Create(&YeeCreateConf{
		MerType:          3,
		MerFullName:      "西安市浐灞生态区国瑞菲百货店",
		MerShortName:     "国瑞菲百货5店",
		MerCertNo:        "92610135MA6X3KW299",
		LegalName:        "李红艳",
		LegalIdCard:      "320830199012082427",
		MerLegalPhone:    "17714500631",
		MerLevel1No:      "129",
		MerLevel2No:      "129004",
		MerProvince:      "610000",
		MerCity:          "610100",
		MerDistrict:      "610111",
		MerAddress:       "玄武东路东路天香心苑1号楼3单元3204室",
		BankAccountType:  "PERSONAL",
		CardNo:           "6217920471658935",
		HeadBankCode:     "SPDB",
		BankCode:         "310308000019",
		BankProvince:     "320000",
		BankCity:         "320800",
		NotifyUrl:        "http://tq.udian.me/v1/common/enter-notify",
		MerAuthorizeType: "WEB_AUTHORIZE",
		ProductInfo: map[string]interface{}{
			"payProductMap": map[string]interface{}{
				"USER_SCAN_PAY": map[string]interface{}{
					"dsPayBankMap": map[string]interface{}{
						"WECHAT_ATIVE_SCAN_OFFLINE": map[string]interface{}{
							"rateType": "PERCENTAGE",
							"rate":     "0.25",
						},
						"ALIPAY_OFFLINE": map[string]interface{}{
							"rateType": "PERCENTAGE",
							"rate":     "0.25",
						},
					},
				},
			},
			"payScenarioMap": map[string]interface{}{
				"OFFLINE_STORE": map[string]interface{}{},
			},
		},
		FileInfo: []*YeeFile{
			{
				QuaType: "IDCARD_FRONT", //法人身份证正面
				QuaUrl:  "http://staticres.yeepay.com/jcptb-merchant-netinjt04/2020/02/27/merchant-1582819054397-af24127f8d3947d0896a2f5c5e62602d-KvzWIzCcziOxQTJmWoCN.png",
			},
			{
				QuaType: "IDCARD_BACK", //法人身份证反面
				QuaUrl:  "http://staticres.yeepay.com/jcptb-merchant-netinjt04/2020/02/27/merchant-1582819100736-47360128c3764612b2fb8fa6f2f9ada1-StGkMXMLSumoBXRezciv.png",
			},
			{
				QuaType: "CORP_CODE", //营业执照
				QuaUrl:  "http://staticres.yeepay.com/jcptb-merchant-netinjt04/2020/02/28/merchant-1582820912412-172d474b0e8d4eef95d5000a52e4df51-zNrYkFypMhOhmdqkSXKD.png",
			},
			{
				QuaType: "SETTLE_BANKCARD", //结算银行卡
				QuaUrl:  "http://staticres.yeepay.com/jcptb-merchant-netinjt04/2020/02/27/merchant-1582819143345-a979d247ab0a4ba6b6e4d753a80dc701-NyPCTcVmEcsIUNmkmGSq.png",
			},
			{
				QuaType: "BUSINESS_PLACE", //门头照
				QuaUrl:  "http://staticres.yeepay.com/jcptb-merchant-netinjt04/2020/02/28/merchant-1582821011356-c179c2d44fa94e018f383ed01a578a02-KuekmdVVNpsWAXOhwBLW.png",
			},
			{
				QuaType: "BUSINESS_SITE", //经营场所照
				QuaUrl:  "http://staticres.yeepay.com/jcptb-merchant-netinjt04/2020/02/28/merchant-1582821065003-660ecf76d66a4d128b980e263734e710-MiOcuHLQrcUzQiStGcpE.png",
			},
			{
				QuaType: "CASHIER_SCENE", //收银台场景照
				QuaUrl:  "http://staticres.yeepay.com/jcptb-merchant-netinjt04/2020/02/28/merchant-1582821065003-660ecf76d66a4d128b980e263734e710-MiOcuHLQrcUzQiStGcpE.png",
			},
			{
				QuaType: "HAND_IDCARD", //法人手持营业执照及身份证（个人只手持身份证）
				QuaUrl:  "http://staticres.yeepay.com/jcptb-merchant-netinjt04/2020/02/28/merchant-1582821905448-47096ab31b7942ca96b8bec8619de173-iVQpPMZRzoMGBOYSUJTW.png",
			},
		},
	})
	fmt.Println(err)
	fmt.Printf("%+v", ret)
}
func TestYeeEnter_Update(t *testing.T) {
	yee := new(YeeEnter)

	yee.InitConfig(&YeeConfig{
		ParentMerchantNo: "10033308281",
		AppKey:           "OPR:10033308281",
		PrivateKey:       "MIIEpAIBAAKCAQEAuRG2HsqR2kXwDtQbCZNtravarGNttzbh8q4SuVAfNb5G7mpf4aKrscS6+P1YQxSsIHXANbm1d1AdtXw9DZd1zZQwqkI63kDO0UcKjSlnsSso4MH3x4YPhHNfXapxotLs9uB8WeMdBvZ9viK1bGJ9IlIm/G08dKi5w90OXAlmY6mGSlSM8PQKFZ4xDvxsx/eVIIkvZyRYSdIYvaGU4C5yep0SDHmtdzu+l3KrnwmuYruDTEgevCXIwNFeFuuxmqTJpxbovLzLQcYfzw5VgrzMQKo81ihRa2TgbmpPqHvz8Ecpv2SD3oZcqJJ+V9Fekn8ZwRii01kZRxmtO6wHJiJA0wIDAQABAoIBAQChyUX4rQXMVw+BJxNrz6I2DOJhiZpEbIoh6OMddVmTVgAUNJIVYmvOQDljqbYbDltbFRUu4mYtI7CVE0McOqgVS9MvRC7KVMV4Vi53MRcg3qYhte+yURQHqgRYkcQ9nz7go2aR/eVGTP2n1hfb5E2YT1EdozESmt2qx/jhpKYJwJCx/Bbt0Puk7tAAFjF+FlfVJBCjaotRaDEa3wbRBloIgHTz7lKMmeIWwDP5APAqTsMjQVTgrUmFbwRy8wlFrS9eHsXvHUNSz40lTbQZYzjCDyAWsnr+d7UJnA1y1OKlW/L/7iA5tZlPooDAcZQrw/2L6eqcIfli6qJS5xp7jPYhAoGBAN1lEL8N+VH4ZlaAIU6JcwhdGP7pqVNXm1E/OHCEuAUrXNUUXPsYzSVsonm/iwK5DBqakVQauHzELF2i1Vq2L2ZFIWzGmXRbOStD25WlRAmXdaR7tmyk1wUbi4pGLcPsZ0T83p7h5jYcU20xqdRWIbFzAd0zhgmLIivgd3/Izco3AoGBANX/GzEg+3/Ezjxm0+PngtNP1XhBBMZfLA/Za9VWxkoRhet7nA0ZS8TdYCucqYsi/xT3Zm6uW4nsfBA6EppX36iS0Ko1kvu/BRB52S4KX2N2Nprv0aI2BKUt8GFnzMu3vYoWwU4HfAPZo4gHhujQyNIGva4O3nU3FA+17IdaS0BFAoGBAMi/WHT01cqW4yyyGAFfrpe52u1hsDCq8mG0YpfcMAQ38oAfa8QfE/1ISPb+UK0SX8BLwVUyuXAgbV4mRTFwmwAv1QQN/J0+DlOFvzks1smftYOEzcAro/C0rk2eHudVl7o9VBtbGGSeQKN1cdngk8KUVu9dRb+nkj+Y1KJg0QD3AoGAUoUlPzSwxvxlavVcEC5eTI4ki1LHPJfGrfXxlzr3C+jl84CfFI4Eoc2cIDUxS+4a30LnxRaHRhBjZv593fa00JSM0pYGL/3hPhE+pnppfjk/pU+FTz/1Wpz0bRtR5dzcwjs0H5rTP8jVPsdoRq78QcFHs68YKasrmUNWCnvPOOECgYBzF69wXWrRK0aBa+5RdjI7oED/7i0l8urRn1Xf3+1KSNbXDLPy/PYkAtpcf6tNBY7n3utFiEZrkc1FM1v74KDmuSij7XPpmIyldVkIwy1/Xx/fRvBSuIVHpbS8FZEUqlCg7B5LB6rFTYyzGOtoXy2A3EVVa5vfsWns9ukY7VPLKA==",
	})

	ret, err := yee.Update(&YeeUpdateConf{
		MerchantNo: "10033406581",
		ProductInfo: map[string]interface{}{
			"payProductMap": map[string]interface{}{
				"USER_SCAN_PAY": map[string]interface{}{
					"dsPayBankMap": map[string]interface{}{
						"WECHAT_ATIVE_SCAN_OFFLINE": map[string]interface{}{
							"rateType":   "PERCENTAGE",
							"rate":       "0.25",
							"openSwitch": "true",
						},
						"ALIPAY_OFFLINE": map[string]interface{}{
							"rateType":   "PERCENTAGE",
							"rate":       "0.25",
							"openSwitch": "true",
						},
					},
				},
			},
			"OFFICIAL_ACCOUNT_PAY": map[string]interface{}{
				"WECHAT_OPENID_OFFLINE": map[string]interface{}{
					"rateType":   "PERCENTAGE",
					"rate":       "0.25",
					"openSwitch": "true",
				},
			},
			"payScenarioMap": map[string]interface{}{
				"OFFLINE_STORE": map[string]interface{}{},
			},
			"settlementProductMap": map[string]interface{}{ //结算产品
				"MERCHANT_SETTLEMENT": map[string]interface{}{ //商户结算
					"feeBearParty": "SUB_MERCHANT", //费用承担方：子商户承担SUB_MERCHANT;系统商承担SYSTEM_MERCHANT
					"dsSettlementBankMap": map[string]interface{}{
						"SETTLEMENT_D1": map[string]interface{}{ //T1、D1二选一，D1手续费时，工作日结算费率为0，非工作日结算照此费率扣费。
							"rateType":       "MIXED",
							"fixedRate":      "0.5",
							"percentRate":    "0.02",
							"settlementType": "AUTO_TIMING", //自助SELF_SERVICE，自动AUTO_TIMING
						},
					},
				},
			},
		},
	})
	fmt.Println(err)
	fmt.Printf("%+v", ret)
}
func TestYeeEnter_Report(t *testing.T) {
	yee := new(YeeEnter)

	yee.InitConfig(&YeeConfig{
		ParentMerchantNo: "10033308281",
		AppKey:           "OPR:10033308281",
		PrivateKey:       "MIIEpAIBAAKCAQEAuRG2HsqR2kXwDtQbCZNtravarGNttzbh8q4SuVAfNb5G7mpf4aKrscS6+P1YQxSsIHXANbm1d1AdtXw9DZd1zZQwqkI63kDO0UcKjSlnsSso4MH3x4YPhHNfXapxotLs9uB8WeMdBvZ9viK1bGJ9IlIm/G08dKi5w90OXAlmY6mGSlSM8PQKFZ4xDvxsx/eVIIkvZyRYSdIYvaGU4C5yep0SDHmtdzu+l3KrnwmuYruDTEgevCXIwNFeFuuxmqTJpxbovLzLQcYfzw5VgrzMQKo81ihRa2TgbmpPqHvz8Ecpv2SD3oZcqJJ+V9Fekn8ZwRii01kZRxmtO6wHJiJA0wIDAQABAoIBAQChyUX4rQXMVw+BJxNrz6I2DOJhiZpEbIoh6OMddVmTVgAUNJIVYmvOQDljqbYbDltbFRUu4mYtI7CVE0McOqgVS9MvRC7KVMV4Vi53MRcg3qYhte+yURQHqgRYkcQ9nz7go2aR/eVGTP2n1hfb5E2YT1EdozESmt2qx/jhpKYJwJCx/Bbt0Puk7tAAFjF+FlfVJBCjaotRaDEa3wbRBloIgHTz7lKMmeIWwDP5APAqTsMjQVTgrUmFbwRy8wlFrS9eHsXvHUNSz40lTbQZYzjCDyAWsnr+d7UJnA1y1OKlW/L/7iA5tZlPooDAcZQrw/2L6eqcIfli6qJS5xp7jPYhAoGBAN1lEL8N+VH4ZlaAIU6JcwhdGP7pqVNXm1E/OHCEuAUrXNUUXPsYzSVsonm/iwK5DBqakVQauHzELF2i1Vq2L2ZFIWzGmXRbOStD25WlRAmXdaR7tmyk1wUbi4pGLcPsZ0T83p7h5jYcU20xqdRWIbFzAd0zhgmLIivgd3/Izco3AoGBANX/GzEg+3/Ezjxm0+PngtNP1XhBBMZfLA/Za9VWxkoRhet7nA0ZS8TdYCucqYsi/xT3Zm6uW4nsfBA6EppX36iS0Ko1kvu/BRB52S4KX2N2Nprv0aI2BKUt8GFnzMu3vYoWwU4HfAPZo4gHhujQyNIGva4O3nU3FA+17IdaS0BFAoGBAMi/WHT01cqW4yyyGAFfrpe52u1hsDCq8mG0YpfcMAQ38oAfa8QfE/1ISPb+UK0SX8BLwVUyuXAgbV4mRTFwmwAv1QQN/J0+DlOFvzks1smftYOEzcAro/C0rk2eHudVl7o9VBtbGGSeQKN1cdngk8KUVu9dRb+nkj+Y1KJg0QD3AoGAUoUlPzSwxvxlavVcEC5eTI4ki1LHPJfGrfXxlzr3C+jl84CfFI4Eoc2cIDUxS+4a30LnxRaHRhBjZv593fa00JSM0pYGL/3hPhE+pnppfjk/pU+FTz/1Wpz0bRtR5dzcwjs0H5rTP8jVPsdoRq78QcFHs68YKasrmUNWCnvPOOECgYBzF69wXWrRK0aBa+5RdjI7oED/7i0l8urRn1Xf3+1KSNbXDLPy/PYkAtpcf6tNBY7n3utFiEZrkc1FM1v74KDmuSij7XPpmIyldVkIwy1/Xx/fRvBSuIVHpbS8FZEUqlCg7B5LB6rFTYyzGOtoXy2A3EVVa5vfsWns9ukY7VPLKA==",
	})

	ret, err := yee.Report(&YeeReportConf{
		MerchantNo:            "10033344434",
		CallBackUrl:           "http://web.51shouqianla.com/v1/payment/yee/yc-report-notify",
		MerchantName:          "西安市浐灞生态区国瑞菲百货店",
		ReportMerchantName:    "西安市浐灞生态区国瑞菲百货店",
		ReportMerchantAlias:   "云创信息",
		ReportMerchantComment: "云创信息",
		ServiceTel:            "17714500631",
		ContactName:           "李红艳",
		ContactPhone:          "17714500631",
		ContactMobile:         "17714500631",
		ContactEmail:          "zhou@ttouch.com.cn",
		MerchantAddress:       "玄武东路东路天香心苑1号楼3单元3204室",
		MerchantProvince:      "610000",
		MerchantCity:          "610100",
		MerchantDistrict:      "610111",
		MerchantLicenseNo:     "",
		CorporateIdCardNo:     "320830199012082427",
		ContactType:           "LEGAL_PERSON",
		ReportInfo: &ReportInfo{
			IDCardName:         "李红艳",
			ThreeInOne:         "1",
			IsInstitution:      "0",
			MerTypeX:           "个体",
			PicIdCardFront:     "http://staticres.yeepay.com/jcptb-merchant-netinjt04/2020/02/26/merchant-1582648319448-5feccacac8564d80bd267370309338d0-aHOkxfjYDRbloJxzgbab.png",
			PicIdCardBack:      "http://staticres.yeepay.com/jcptb-merchant-netinjt04/2020/02/26/merchant-1582648379156-1cbf09624b5c43c5999aef65d0ab870a-fQUMQGMkWkamHOsVgZUR.png",
			PicMerchantLicense: "http://staticres.yeepay.com/jcptb-merchant-netinjt04/2020/02/26/merchant-1582648597517-507c886323284a8f8219e49f6bbf7a32-DRrquXMLmUQkTAEwwQcg.png",
			PicMerShopPhoto:    "http://staticres.yeepay.com/jcptb-merchant-netinjt04/2020/02/26/merchant-1582719382926-7f1cf94400704ecdb707741a05b4a54f-fqYmNOiRrOlliWPdBPzm.png",
		},
		ReportFeeType: "XIANXIA",
	})

	fmt.Println(err)
	fmt.Println(ret)
}

func TestYeeEnter_HMacKeyQuery(t *testing.T) {
	yee := new(YeeEnter)

	yee.InitConfig(&YeeConfig{
		ParentMerchantNo: "10033308281",
		AppKey:           "OPR:10033308281",
		PrivateKey:       "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC5EbYeypHaRfAO1BsJk22tq9qsY223NuHyrhK5UB81vkbual/hoquxxLr4/VhDFKwgdcA1ubV3UB21fD0Nl3XNlDCqQjreQM7RRwqNKWexKyjgwffHhg+Ec19dqnGi0uz24HxZ4x0G9n2+IrVsYn0iUib8bTx0qLnD3Q5cCWZjqYZKVIzw9AoVnjEO/GzH95UgiS9nJFhJ0hi9oZTgLnJ6nRIMea13O76XcqufCa5iu4NMSB68JcjA0V4W67GapMmnFui8vMtBxh/PDlWCvMxAqjzWKFFrZOBuak+oe/PwRym/ZIPehlyokn5X0V6SfxnBGKLTWRlHGa07rAcmIkDTAgMBAAECggEBAKHJRfitBcxXD4EnE2vPojYM4mGJmkRsiiHo4x11WZNWABQ0khVia85AOWOpthsOW1sVFS7iZi0jsJUTQxw6qBVL0y9ELspUxXhWLncxFyDepiG177JRFAeqBFiRxD2fPuCjZpH95UZM/afWF9vkTZhPUR2jMRKa3arH+OGkpgnAkLH8Fu3Q+6Tu0AAWMX4WV9UkEKNqi1FoMRrfBtEGWgiAdPPuUoyZ4hbAM/kA8CpOwyNBVOCtSYVvBHLzCUWtL14exe8dQ1LPjSVNtBljOMIPIBayev53tQmcDXLU4qVb8v/uIDm1mU+igMBxlCvD/Yvp6pwh+WLqolLnGnuM9iECgYEA3WUQvw35UfhmVoAhTolzCF0Y/umpU1ebUT84cIS4BStc1RRc+xjNJWyieb+LArkMGpqRVBq4fMQsXaLVWrYvZkUhbMaZdFs5K0PblaVECZd1pHu2bKTXBRuLikYtw+xnRPzenuHmNhxTbTGp1FYhsXMB3TOGCYsiK+B3f8jNyjcCgYEA1f8bMSD7f8TOPGbT4+eC00/VeEEExl8sD9lr1VbGShGF63ucDRlLxN1gK5ypiyL/FPdmbq5biex8EDoSmlffqJLQqjWS+78FEHnZLgpfY3Y2mu/RojYEpS3wYWfMy7e9ihbBTgd8A9mjiAeG6NDI0ga9rg7edTcUD7Xsh1pLQEUCgYEAyL9YdPTVypbjLLIYAV+ul7na7WGwMKryYbRil9wwBDfygB9rxB8T/UhI9v5QrRJfwEvBVTK5cCBtXiZFMXCbAC/VBA38nT4OU4W/OSzWyZ+1g4TNwCuj8LSuTZ4e51WXuj1UG1sYZJ5Ao3Vx2eCTwpRW711Fv6eSP5jUomDRAPcCgYBShSU/NLDG/GVq9VwQLl5MjiSLUsc8l8at9fGXOvcL6OXzgJ8UjgShzZwgNTFL7hrfQufFFodGEGNm/n3d9rTQlIzSlgYv/eE+ET6meml+OT+lT4VPP/VanPRtG1Hl3NzCOzQfmtM/yNU+x2hGrvxBwUezrxgpqyuZQ1YKe8844QKBgHMXr3BdatErRoFr7lF2MjugQP/uLSXy6tGfVd/f7UpI1tcMs/L89iQC2lx/q00Fjufe60WIRmuRzUUzW/vgoOa5KKPtc+mYjKV1WQjDLX9fH99G8FK4hUeltLwVkRSqUKDsHksHqsVNjLMY62hfLYDcRVVrm9+xaez26RjtU8so",
	})

	ret, err := yee.HMacKeyQuery(&YeeHMacKeyQueryConf{
		MerchantNo: "10033343784",
	})

	fmt.Println(err)
	fmt.Println(ret)
}
func TestYeeEnter_BranchInfo(t *testing.T) {
	yee := new(YeeEnter)

	yee.InitConfig(&YeeConfig{
		ParentMerchantNo: "10033308281",
		AppKey:           "OPR:10014929805",
		PrivateKey:       "MIIEpAIBAAKCAQEAuRG2HsqR2kXwDtQbCZNtravarGNttzbh8q4SuVAfNb5G7mpf4aKrscS6+P1YQxSsIHXANbm1d1AdtXw9DZd1zZQwqkI63kDO0UcKjSlnsSso4MH3x4YPhHNfXapxotLs9uB8WeMdBvZ9viK1bGJ9IlIm/G08dKi5w90OXAlmY6mGSlSM8PQKFZ4xDvxsx/eVIIkvZyRYSdIYvaGU4C5yep0SDHmtdzu+l3KrnwmuYruDTEgevCXIwNFeFuuxmqTJpxbovLzLQcYfzw5VgrzMQKo81ihRa2TgbmpPqHvz8Ecpv2SD3oZcqJJ+V9Fekn8ZwRii01kZRxmtO6wHJiJA0wIDAQABAoIBAQChyUX4rQXMVw+BJxNrz6I2DOJhiZpEbIoh6OMddVmTVgAUNJIVYmvOQDljqbYbDltbFRUu4mYtI7CVE0McOqgVS9MvRC7KVMV4Vi53MRcg3qYhte+yURQHqgRYkcQ9nz7go2aR/eVGTP2n1hfb5E2YT1EdozESmt2qx/jhpKYJwJCx/Bbt0Puk7tAAFjF+FlfVJBCjaotRaDEa3wbRBloIgHTz7lKMmeIWwDP5APAqTsMjQVTgrUmFbwRy8wlFrS9eHsXvHUNSz40lTbQZYzjCDyAWsnr+d7UJnA1y1OKlW/L/7iA5tZlPooDAcZQrw/2L6eqcIfli6qJS5xp7jPYhAoGBAN1lEL8N+VH4ZlaAIU6JcwhdGP7pqVNXm1E/OHCEuAUrXNUUXPsYzSVsonm/iwK5DBqakVQauHzELF2i1Vq2L2ZFIWzGmXRbOStD25WlRAmXdaR7tmyk1wUbi4pGLcPsZ0T83p7h5jYcU20xqdRWIbFzAd0zhgmLIivgd3/Izco3AoGBANX/GzEg+3/Ezjxm0+PngtNP1XhBBMZfLA/Za9VWxkoRhet7nA0ZS8TdYCucqYsi/xT3Zm6uW4nsfBA6EppX36iS0Ko1kvu/BRB52S4KX2N2Nprv0aI2BKUt8GFnzMu3vYoWwU4HfAPZo4gHhujQyNIGva4O3nU3FA+17IdaS0BFAoGBAMi/WHT01cqW4yyyGAFfrpe52u1hsDCq8mG0YpfcMAQ38oAfa8QfE/1ISPb+UK0SX8BLwVUyuXAgbV4mRTFwmwAv1QQN/J0+DlOFvzks1smftYOEzcAro/C0rk2eHudVl7o9VBtbGGSeQKN1cdngk8KUVu9dRb+nkj+Y1KJg0QD3AoGAUoUlPzSwxvxlavVcEC5eTI4ki1LHPJfGrfXxlzr3C+jl84CfFI4Eoc2cIDUxS+4a30LnxRaHRhBjZv593fa00JSM0pYGL/3hPhE+pnppfjk/pU+FTz/1Wpz0bRtR5dzcwjs0H5rTP8jVPsdoRq78QcFHs68YKasrmUNWCnvPOOECgYBzF69wXWrRK0aBa+5RdjI7oED/7i0l8urRn1Xf3+1KSNbXDLPy/PYkAtpcf6tNBY7n3utFiEZrkc1FM1v74KDmuSij7XPpmIyldVkIwy1/Xx/fRvBSuIVHpbS8FZEUqlCg7B5LB6rFTYyzGOtoXy2A3EVVa5vfsWns9ukY7VPLKA==",
	})

	yee.BranchInfo(&YeeBranchInfoConf{
		HeadBankCode: "ABC",
		ProvinceCode: "110000",
		CityCode:     "110100",
	})
}
func TestYeeEnter_ReportQuery(t *testing.T) {
	yee := new(YeeEnter)

	yee.InitConfig(&YeeConfig{
		ParentMerchantNo: "10033308281",
		AppKey:           "OPR:10033308281",
		PrivateKey:       "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC5EbYeypHaRfAO1BsJk22tq9qsY223NuHyrhK5UB81vkbual/hoquxxLr4/VhDFKwgdcA1ubV3UB21fD0Nl3XNlDCqQjreQM7RRwqNKWexKyjgwffHhg+Ec19dqnGi0uz24HxZ4x0G9n2+IrVsYn0iUib8bTx0qLnD3Q5cCWZjqYZKVIzw9AoVnjEO/GzH95UgiS9nJFhJ0hi9oZTgLnJ6nRIMea13O76XcqufCa5iu4NMSB68JcjA0V4W67GapMmnFui8vMtBxh/PDlWCvMxAqjzWKFFrZOBuak+oe/PwRym/ZIPehlyokn5X0V6SfxnBGKLTWRlHGa07rAcmIkDTAgMBAAECggEBAKHJRfitBcxXD4EnE2vPojYM4mGJmkRsiiHo4x11WZNWABQ0khVia85AOWOpthsOW1sVFS7iZi0jsJUTQxw6qBVL0y9ELspUxXhWLncxFyDepiG177JRFAeqBFiRxD2fPuCjZpH95UZM/afWF9vkTZhPUR2jMRKa3arH+OGkpgnAkLH8Fu3Q+6Tu0AAWMX4WV9UkEKNqi1FoMRrfBtEGWgiAdPPuUoyZ4hbAM/kA8CpOwyNBVOCtSYVvBHLzCUWtL14exe8dQ1LPjSVNtBljOMIPIBayev53tQmcDXLU4qVb8v/uIDm1mU+igMBxlCvD/Yvp6pwh+WLqolLnGnuM9iECgYEA3WUQvw35UfhmVoAhTolzCF0Y/umpU1ebUT84cIS4BStc1RRc+xjNJWyieb+LArkMGpqRVBq4fMQsXaLVWrYvZkUhbMaZdFs5K0PblaVECZd1pHu2bKTXBRuLikYtw+xnRPzenuHmNhxTbTGp1FYhsXMB3TOGCYsiK+B3f8jNyjcCgYEA1f8bMSD7f8TOPGbT4+eC00/VeEEExl8sD9lr1VbGShGF63ucDRlLxN1gK5ypiyL/FPdmbq5biex8EDoSmlffqJLQqjWS+78FEHnZLgpfY3Y2mu/RojYEpS3wYWfMy7e9ihbBTgd8A9mjiAeG6NDI0ga9rg7edTcUD7Xsh1pLQEUCgYEAyL9YdPTVypbjLLIYAV+ul7na7WGwMKryYbRil9wwBDfygB9rxB8T/UhI9v5QrRJfwEvBVTK5cCBtXiZFMXCbAC/VBA38nT4OU4W/OSzWyZ+1g4TNwCuj8LSuTZ4e51WXuj1UG1sYZJ5Ao3Vx2eCTwpRW711Fv6eSP5jUomDRAPcCgYBShSU/NLDG/GVq9VwQLl5MjiSLUsc8l8at9fGXOvcL6OXzgJ8UjgShzZwgNTFL7hrfQufFFodGEGNm/n3d9rTQlIzSlgYv/eE+ET6meml+OT+lT4VPP/VanPRtG1Hl3NzCOzQfmtM/yNU+x2hGrvxBwUezrxgpqyuZQ1YKe8844QKBgHMXr3BdatErRoFr7lF2MjugQP/uLSXy6tGfVd/f7UpI1tcMs/L89iQC2lx/q00Fjufe60WIRmuRzUUzW/vgoOa5KKPtc+mYjKV1WQjDLX9fH99G8FK4hUeltLwVkRSqUKDsHksHqsVNjLMY62hfLYDcRVVrm9+xaez26RjtU8so",
	})

	yee.ReportQuery(&YeeReportQueryConf{
		MerchantNo: "10033343784",
	})
}
func TestYeeEnter_RegStatusQuery(t *testing.T) {
	yee := new(YeeEnter)

	yee.InitConfig(&YeeConfig{
		ParentMerchantNo: "10033308281",
		AppKey:           "OPR:10033308281",
		PrivateKey:       "MIIEpAIBAAKCAQEAuRG2HsqR2kXwDtQbCZNtravarGNttzbh8q4SuVAfNb5G7mpf4aKrscS6+P1YQxSsIHXANbm1d1AdtXw9DZd1zZQwqkI63kDO0UcKjSlnsSso4MH3x4YPhHNfXapxotLs9uB8WeMdBvZ9viK1bGJ9IlIm/G08dKi5w90OXAlmY6mGSlSM8PQKFZ4xDvxsx/eVIIkvZyRYSdIYvaGU4C5yep0SDHmtdzu+l3KrnwmuYruDTEgevCXIwNFeFuuxmqTJpxbovLzLQcYfzw5VgrzMQKo81ihRa2TgbmpPqHvz8Ecpv2SD3oZcqJJ+V9Fekn8ZwRii01kZRxmtO6wHJiJA0wIDAQABAoIBAQChyUX4rQXMVw+BJxNrz6I2DOJhiZpEbIoh6OMddVmTVgAUNJIVYmvOQDljqbYbDltbFRUu4mYtI7CVE0McOqgVS9MvRC7KVMV4Vi53MRcg3qYhte+yURQHqgRYkcQ9nz7go2aR/eVGTP2n1hfb5E2YT1EdozESmt2qx/jhpKYJwJCx/Bbt0Puk7tAAFjF+FlfVJBCjaotRaDEa3wbRBloIgHTz7lKMmeIWwDP5APAqTsMjQVTgrUmFbwRy8wlFrS9eHsXvHUNSz40lTbQZYzjCDyAWsnr+d7UJnA1y1OKlW/L/7iA5tZlPooDAcZQrw/2L6eqcIfli6qJS5xp7jPYhAoGBAN1lEL8N+VH4ZlaAIU6JcwhdGP7pqVNXm1E/OHCEuAUrXNUUXPsYzSVsonm/iwK5DBqakVQauHzELF2i1Vq2L2ZFIWzGmXRbOStD25WlRAmXdaR7tmyk1wUbi4pGLcPsZ0T83p7h5jYcU20xqdRWIbFzAd0zhgmLIivgd3/Izco3AoGBANX/GzEg+3/Ezjxm0+PngtNP1XhBBMZfLA/Za9VWxkoRhet7nA0ZS8TdYCucqYsi/xT3Zm6uW4nsfBA6EppX36iS0Ko1kvu/BRB52S4KX2N2Nprv0aI2BKUt8GFnzMu3vYoWwU4HfAPZo4gHhujQyNIGva4O3nU3FA+17IdaS0BFAoGBAMi/WHT01cqW4yyyGAFfrpe52u1hsDCq8mG0YpfcMAQ38oAfa8QfE/1ISPb+UK0SX8BLwVUyuXAgbV4mRTFwmwAv1QQN/J0+DlOFvzks1smftYOEzcAro/C0rk2eHudVl7o9VBtbGGSeQKN1cdngk8KUVu9dRb+nkj+Y1KJg0QD3AoGAUoUlPzSwxvxlavVcEC5eTI4ki1LHPJfGrfXxlzr3C+jl84CfFI4Eoc2cIDUxS+4a30LnxRaHRhBjZv593fa00JSM0pYGL/3hPhE+pnppfjk/pU+FTz/1Wpz0bRtR5dzcwjs0H5rTP8jVPsdoRq78QcFHs68YKasrmUNWCnvPOOECgYBzF69wXWrRK0aBa+5RdjI7oED/7i0l8urRn1Xf3+1KSNbXDLPy/PYkAtpcf6tNBY7n3utFiEZrkc1FM1v74KDmuSij7XPpmIyldVkIwy1/Xx/fRvBSuIVHpbS8FZEUqlCg7B5LB6rFTYyzGOtoXy2A3EVVa5vfsWns9ukY7VPLKA==",
	})

	yee.RegStatusQuery(&YeeRegStatusQueryConf{
		ParentMerchantNo: "10033308281",
		MerchantNo:       "10033344427",
	})
}
func TestYeeUpload_Handle(t *testing.T) {
	yee := new(YeeEnter)

	yee.InitConfig(&YeeConfig{
		ParentMerchantNo: "10033308281",
		AppKey:           "OPR:10033308281",
		PrivateKey:       "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC5EbYeypHaRfAO1BsJk22tq9qsY223NuHyrhK5UB81vkbual/hoquxxLr4/VhDFKwgdcA1ubV3UB21fD0Nl3XNlDCqQjreQM7RRwqNKWexKyjgwffHhg+Ec19dqnGi0uz24HxZ4x0G9n2+IrVsYn0iUib8bTx0qLnD3Q5cCWZjqYZKVIzw9AoVnjEO/GzH95UgiS9nJFhJ0hi9oZTgLnJ6nRIMea13O76XcqufCa5iu4NMSB68JcjA0V4W67GapMmnFui8vMtBxh/PDlWCvMxAqjzWKFFrZOBuak+oe/PwRym/ZIPehlyokn5X0V6SfxnBGKLTWRlHGa07rAcmIkDTAgMBAAECggEBAKHJRfitBcxXD4EnE2vPojYM4mGJmkRsiiHo4x11WZNWABQ0khVia85AOWOpthsOW1sVFS7iZi0jsJUTQxw6qBVL0y9ELspUxXhWLncxFyDepiG177JRFAeqBFiRxD2fPuCjZpH95UZM/afWF9vkTZhPUR2jMRKa3arH+OGkpgnAkLH8Fu3Q+6Tu0AAWMX4WV9UkEKNqi1FoMRrfBtEGWgiAdPPuUoyZ4hbAM/kA8CpOwyNBVOCtSYVvBHLzCUWtL14exe8dQ1LPjSVNtBljOMIPIBayev53tQmcDXLU4qVb8v/uIDm1mU+igMBxlCvD/Yvp6pwh+WLqolLnGnuM9iECgYEA3WUQvw35UfhmVoAhTolzCF0Y/umpU1ebUT84cIS4BStc1RRc+xjNJWyieb+LArkMGpqRVBq4fMQsXaLVWrYvZkUhbMaZdFs5K0PblaVECZd1pHu2bKTXBRuLikYtw+xnRPzenuHmNhxTbTGp1FYhsXMB3TOGCYsiK+B3f8jNyjcCgYEA1f8bMSD7f8TOPGbT4+eC00/VeEEExl8sD9lr1VbGShGF63ucDRlLxN1gK5ypiyL/FPdmbq5biex8EDoSmlffqJLQqjWS+78FEHnZLgpfY3Y2mu/RojYEpS3wYWfMy7e9ihbBTgd8A9mjiAeG6NDI0ga9rg7edTcUD7Xsh1pLQEUCgYEAyL9YdPTVypbjLLIYAV+ul7na7WGwMKryYbRil9wwBDfygB9rxB8T/UhI9v5QrRJfwEvBVTK5cCBtXiZFMXCbAC/VBA38nT4OU4W/OSzWyZ+1g4TNwCuj8LSuTZ4e51WXuj1UG1sYZJ5Ao3Vx2eCTwpRW711Fv6eSP5jUomDRAPcCgYBShSU/NLDG/GVq9VwQLl5MjiSLUsc8l8at9fGXOvcL6OXzgJ8UjgShzZwgNTFL7hrfQufFFodGEGNm/n3d9rTQlIzSlgYv/eE+ET6meml+OT+lT4VPP/VanPRtG1Hl3NzCOzQfmtM/yNU+x2hGrvxBwUezrxgpqyuZQ1YKe8844QKBgHMXr3BdatErRoFr7lF2MjugQP/uLSXy6tGfVd/f7UpI1tcMs/L89iQC2lx/q00Fjufe60WIRmuRzUUzW/vgoOa5KKPtc+mYjKV1WQjDLX9fH99G8FK4hUeltLwVkRSqUKDsHksHqsVNjLMY62hfLYDcRVVrm9+xaez26RjtU8so",
	})

	b, err := http.GetBytes("http://cdn.51shouqianla.com/FjvBWqLV1sZ0eevVKFwUgkht4Nxz")

	if err != nil {
		return
	}

	ret, err := yee.Upload(&YeeUploadConf{
		File: b,
	})

	fmt.Printf("%+v\n", ret)
	fmt.Println(err)
}
