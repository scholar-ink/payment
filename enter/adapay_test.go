package enter

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/scholar-ink/payment/helper"
	"io/ioutil"
	"math/rand"
	"strings"
	"testing"
)

func TestAdaPayEnter_Create(t *testing.T) {
	AdaPay := new(AdaPayEnter)

	AdaPay.InitConfig(&AdaPayConfig{
		ApiKey:     "api_live_c92ccfb0-c22e-49b1-8fa7-248b08c8ad8f",
		PrivateKey: "MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBALSAIqdFDAFWLm8HkPbYfOPmbM96qTKfR91B6tHR7FVixeicZbw11CWIw8O56dKQ0wotZTrWC75yvXpn2dTdkQgZp9cytXZ6MLnfiKvj+032wnAU3ZMESc/Onrhf589nWC6RvIy26tIQ0yUR5HLrpEN+A4qVDVqd00TQZc8YRoHtAgMBAAECgYA94uj+vNe+5ZOKEegMGnHHmcuY349/gckb/WvLgNQs+m6ssGLZQwN30wp74xReU7Vn+eSJZbYlGCYK/+xZ5ZXBw5uVhlDCmjWmFGVZVpiOibQStPqPFen4qhLMMsP8pZtJOagIjaGfMc9Wf36f/fJ3LFLXz3E7wjLsUmDslnDZKQJBAOLd2nu40Fw2FAVwEHd+5GbSkDWNw0HdlUbBHWV+fEUnLsiIE+Hw1nT6oOHfPMml5naKd8BLmkYUC2kfa9rFWh8CQQDLrgQTZ0bQtYjpzVkj3k+D1KnVuizy1LIQs4oZSMU0YBaVfUVLQZzHubTNpAMzH5hmfXg/fz+QKD0aSjRmTDpzAkBWHZqqrhvBdPGioshNY8h1U2ZUPcypeuAILJPpC9tGMLpsemL5t/7gBqb9Nk0Pyj6yLpuITepwwXkXXUsGjzVHAkEAhgBlxBJFX9ifTBsC03tWWwhV+Dw1iElxIVXNvJbIz42MLiutpDZ1nF1MW6LVTBQ0YvGXZEcmnYQrtxks4kSyiwJAWE+xShtaoOLI0mdOBMAI2amJWpBV4NCvHhVbWwh5F0Jxq3S79mXqvnWOsXBefjKZVFBCnqjlMFsExv6WEifCQg==",
		PublicKey:  "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC0gCKnRQwBVi5vB5D22Hzj5mzPeqkyn0fdQerR0exVYsXonGW8NdQliMPDuenSkNMKLWU61gu+cr16Z9nU3ZEIGafXMrV2ejC534ir4/tN9sJwFN2TBEnPzp64X+fPZ1gukbyMturSENMlEeRy66RDfgOKlQ1andNE0GXPGEaB7QIDAQAB",
	})

	userPhone := "1" + fmt.Sprintf("%08d", 201120) + fmt.Sprintf("%02v", rand.Int31n(100))

	requestId := strings.ReplaceAll(uuid.New().String(), "-", "")

	ret, err := AdaPay.Create(&AdaPayCreateConf{
		RequestId:               requestId,
		UsrPhone:                userPhone,
		ContName:                "周永珍",
		ContPhone:               "18014167019",
		CustomerEmail:           userPhone + "@ttouh.com.cn",
		MerName:                 "武汉市江夏区万里香农家乐菜馆",
		MerShortName:            "万里香农家乐菜馆",
		LicenseCode:             "92420115MA4J523P8H",
		RegAddr:                 "贵州省贵安新区湖潮乡广兴村(青山桥)",
		CustAddr:                "贵州省贵安新区湖潮乡广兴村(青山桥)",
		CustTel:                 "18014167019",
		MerStartValidDate:       "20191031",
		MerValidDate:            "20991231",
		LegalName:               "周永珍",
		LegalType:               "0",
		LegalIdNo:               "320830196708062420",
		LegalMp:                 "18014167019",
		LegalStartCertIdExpires: "20180209",
		LegalIdExpires:          "20991231",
		CardName:                "周永珍",
		CardIdMask:              "6217920451370931",
		BankAcctType:            "2",
		ProvCode:                "0032",
		AreaCode:                "3203",
		BankCode:                "03050000",
		FeeRateList: []*AdaRate{
			{
				RateChannel: "wx_pub_offline",
				FeeRate:     "0.003",
			},
			{
				RateChannel: "alipay_qr_offline",
				FeeRate:     "0.003",
			},
		},
	})
	fmt.Println(err)
	fmt.Printf("%+v", ret)
}
func TestAdaPayEnter_Query(t *testing.T) {
	en := new(AdaPayEnter)

	en.InitConfig(&AdaPayConfig{
		ApiKey:     "api_live_c92ccfb0-c22e-49b1-8fa7-248b08c8ad8f",
		PrivateKey: "MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBALSAIqdFDAFWLm8HkPbYfOPmbM96qTKfR91B6tHR7FVixeicZbw11CWIw8O56dKQ0wotZTrWC75yvXpn2dTdkQgZp9cytXZ6MLnfiKvj+032wnAU3ZMESc/Onrhf589nWC6RvIy26tIQ0yUR5HLrpEN+A4qVDVqd00TQZc8YRoHtAgMBAAECgYA94uj+vNe+5ZOKEegMGnHHmcuY349/gckb/WvLgNQs+m6ssGLZQwN30wp74xReU7Vn+eSJZbYlGCYK/+xZ5ZXBw5uVhlDCmjWmFGVZVpiOibQStPqPFen4qhLMMsP8pZtJOagIjaGfMc9Wf36f/fJ3LFLXz3E7wjLsUmDslnDZKQJBAOLd2nu40Fw2FAVwEHd+5GbSkDWNw0HdlUbBHWV+fEUnLsiIE+Hw1nT6oOHfPMml5naKd8BLmkYUC2kfa9rFWh8CQQDLrgQTZ0bQtYjpzVkj3k+D1KnVuizy1LIQs4oZSMU0YBaVfUVLQZzHubTNpAMzH5hmfXg/fz+QKD0aSjRmTDpzAkBWHZqqrhvBdPGioshNY8h1U2ZUPcypeuAILJPpC9tGMLpsemL5t/7gBqb9Nk0Pyj6yLpuITepwwXkXXUsGjzVHAkEAhgBlxBJFX9ifTBsC03tWWwhV+Dw1iElxIVXNvJbIz42MLiutpDZ1nF1MW6LVTBQ0YvGXZEcmnYQrtxks4kSyiwJAWE+xShtaoOLI0mdOBMAI2amJWpBV4NCvHhVbWwh5F0Jxq3S79mXqvnWOsXBefjKZVFBCnqjlMFsExv6WEifCQg==",
	})

	ret, err := en.Query(&AdaPayQueryConf{
		RequestId: "d38fd20104d34ae78a8dfd2088f4ab7f",
	})

	fmt.Printf("%+v", ret)
	fmt.Printf("%+v", ret.AppIdList[0])
	fmt.Println(err)
}

func TestAdaPayEnter_Config(t *testing.T) {
	AdaPay := new(AdaPayEnter)

	AdaPay.InitConfig(&AdaPayConfig{
		ApiKey:     "api_live_c92ccfb0-c22e-49b1-8fa7-248b08c8ad8f",
		PrivateKey: "MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBALSAIqdFDAFWLm8HkPbYfOPmbM96qTKfR91B6tHR7FVixeicZbw11CWIw8O56dKQ0wotZTrWC75yvXpn2dTdkQgZp9cytXZ6MLnfiKvj+032wnAU3ZMESc/Onrhf589nWC6RvIy26tIQ0yUR5HLrpEN+A4qVDVqd00TQZc8YRoHtAgMBAAECgYA94uj+vNe+5ZOKEegMGnHHmcuY349/gckb/WvLgNQs+m6ssGLZQwN30wp74xReU7Vn+eSJZbYlGCYK/+xZ5ZXBw5uVhlDCmjWmFGVZVpiOibQStPqPFen4qhLMMsP8pZtJOagIjaGfMc9Wf36f/fJ3LFLXz3E7wjLsUmDslnDZKQJBAOLd2nu40Fw2FAVwEHd+5GbSkDWNw0HdlUbBHWV+fEUnLsiIE+Hw1nT6oOHfPMml5naKd8BLmkYUC2kfa9rFWh8CQQDLrgQTZ0bQtYjpzVkj3k+D1KnVuizy1LIQs4oZSMU0YBaVfUVLQZzHubTNpAMzH5hmfXg/fz+QKD0aSjRmTDpzAkBWHZqqrhvBdPGioshNY8h1U2ZUPcypeuAILJPpC9tGMLpsemL5t/7gBqb9Nk0Pyj6yLpuITepwwXkXXUsGjzVHAkEAhgBlxBJFX9ifTBsC03tWWwhV+Dw1iElxIVXNvJbIz42MLiutpDZ1nF1MW6LVTBQ0YvGXZEcmnYQrtxks4kSyiwJAWE+xShtaoOLI0mdOBMAI2amJWpBV4NCvHhVbWwh5F0Jxq3S79mXqvnWOsXBefjKZVFBCnqjlMFsExv6WEifCQg==",
		PublicKey:  "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC0gCKnRQwBVi5vB5D22Hzj5mzPeqkyn0fdQerR0exVYsXonGW8NdQliMPDuenSkNMKLWU61gu+cr16Z9nU3ZEIGafXMrV2ejC534ir4/tN9sJwFN2TBEnPzp64X+fPZ1gukbyMturSENMlEeRy66RDfgOKlQ1andNE0GXPGEaB7QIDAQAB",
	})

	requestId := strings.ReplaceAll(uuid.New().String(), "-", "")

	ret, err := AdaPay.Config(&AdaPayConfigConf{
		RequestId:      requestId,
		SubApiKey:      "api_live_90f01d3d-c94a-4026-924a-63584c7abd10",
		BankChannelNo:  "00000675",
		FeeType:        "02",
		AppId:          "app_691593e9-029e-49b4-b531-86ab229f8d94",
		WxCategory:     "302",
		AliPayCategory: "2015063000020189",
		ClsId:          "7372",
		ModelType:      "1",
		MerType:        "5",
		ProvinceCode:   "520000",
		CityCode:       "520100",
		DistrictCode:   "520102",
		AddValueList: map[string]interface{}{
			"wx_pub": map[string]interface{}{
				"appid": "wx7e54b6e6740d4360",
				"path":  "https://qr.51shouqianla.com/",
			},
			"alipay_qr": "",
		},
	})
	fmt.Println(err)
	fmt.Printf("%+v", ret)
}

func TestAdaPayEnter_ConfigQuery(t *testing.T) {
	AdaPay := new(AdaPayEnter)

	AdaPay.InitConfig(&AdaPayConfig{
		ApiKey:     "api_live_c92ccfb0-c22e-49b1-8fa7-248b08c8ad8f",
		PrivateKey: "MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBALSAIqdFDAFWLm8HkPbYfOPmbM96qTKfR91B6tHR7FVixeicZbw11CWIw8O56dKQ0wotZTrWC75yvXpn2dTdkQgZp9cytXZ6MLnfiKvj+032wnAU3ZMESc/Onrhf589nWC6RvIy26tIQ0yUR5HLrpEN+A4qVDVqd00TQZc8YRoHtAgMBAAECgYA94uj+vNe+5ZOKEegMGnHHmcuY349/gckb/WvLgNQs+m6ssGLZQwN30wp74xReU7Vn+eSJZbYlGCYK/+xZ5ZXBw5uVhlDCmjWmFGVZVpiOibQStPqPFen4qhLMMsP8pZtJOagIjaGfMc9Wf36f/fJ3LFLXz3E7wjLsUmDslnDZKQJBAOLd2nu40Fw2FAVwEHd+5GbSkDWNw0HdlUbBHWV+fEUnLsiIE+Hw1nT6oOHfPMml5naKd8BLmkYUC2kfa9rFWh8CQQDLrgQTZ0bQtYjpzVkj3k+D1KnVuizy1LIQs4oZSMU0YBaVfUVLQZzHubTNpAMzH5hmfXg/fz+QKD0aSjRmTDpzAkBWHZqqrhvBdPGioshNY8h1U2ZUPcypeuAILJPpC9tGMLpsemL5t/7gBqb9Nk0Pyj6yLpuITepwwXkXXUsGjzVHAkEAhgBlxBJFX9ifTBsC03tWWwhV+Dw1iElxIVXNvJbIz42MLiutpDZ1nF1MW6LVTBQ0YvGXZEcmnYQrtxks4kSyiwJAWE+xShtaoOLI0mdOBMAI2amJWpBV4NCvHhVbWwh5F0Jxq3S79mXqvnWOsXBefjKZVFBCnqjlMFsExv6WEifCQg==",
		PublicKey:  "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC0gCKnRQwBVi5vB5D22Hzj5mzPeqkyn0fdQerR0exVYsXonGW8NdQliMPDuenSkNMKLWU61gu+cr16Z9nU3ZEIGafXMrV2ejC534ir4/tN9sJwFN2TBEnPzp64X+fPZ1gukbyMturSENMlEeRy66RDfgOKlQ1andNE0GXPGEaB7QIDAQAB",
	})

	ret, err := AdaPay.ConfigQuery(&AdaPayConfigQueryConf{
		RequestId: "f23cb8954493406589ce865389a870f4",
	})
	fmt.Println(err)
	fmt.Printf("%+v", ret)
}

func TestAdaPayEnter_Query2(t *testing.T) {
	pfxData, _ := ioutil.ReadFile("udian2.pfx")

	//encrptData, err := helper.Rsa1Encrypt(pfxData, []byte("1111"), "Udian888")
	////
	//fmt.Println(encrptData)
	//fmt.Println(err)

	encrptData := "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"
	//
	decrptData, err := helper.Rsa1Decrypt(pfxData, encrptData, "Udian888")

	fmt.Println(string(decrptData))
	fmt.Println(err)
}
