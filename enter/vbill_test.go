package enter

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestVBillEnter_Create(t *testing.T) {
	vBill := new(VBillEnter)

	vBill.InitConfig(&VBillConfig{
		PrivateKey: "MIICdgIBVBillNBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBALSAIqdFDAFWLm8HkPbYfOPmbM96qTKfR91B6tHR7FVixeicZbw11CWIw8O56dKQ0wotZTrWC75yvXpn2dTdkQgZp9cytXZ6MLnfiKvj+032wnAU3ZMESc/Onrhf589nWC6RvIy26tIQ0yUR5HLrpEN+A4qVDVqd00TQZc8YRoHtAgMBAAECgYA94uj+vNe+5ZOKEegMGnHHmcuY349/gckb/WvLgNQs+m6ssGLZQwN30wp74xReU7Vn+eSJZbYlGCYK/+xZ5ZXBw5uVhlDCmjWmFGVZVpiOibQStPqPFen4qhLMMsP8pZtJOagIjaGfMc9Wf36f/fJ3LFLXz3E7wjLsUmDslnDZKQJBAOLd2nu40Fw2FAVwEHd+5GbSkDWNw0HdlUbBHWV+fEUnLsiIE+Hw1nT6oOHfPMml5naKd8BLmkYUC2kfa9rFWh8CQQDLrgQTZ0bQtYjpzVkj3k+D1KnVuizy1LIQs4oZSMU0YBaVfUVLQZzHubTNpAMzH5hmfXg/fz+QKD0aSjRmTDpzAkBWHZqqrhvBdPGioshNY8h1U2ZUPcypeuAILJPpC9tGMLpsemL5t/7gBqb9Nk0Pyj6yLpuITepwwXkXXUsGjzVHAkEAhgBlxBJFX9ifTBsC03tWWwhV+Dw1iElxIVXNvJbIz42MLiutpDZ1nF1MW6LVTBQ0YvGXZEcmnYQrtxks4kSyiwJAWE+xShtaoOLI0mdOBMAI2amJWpBV4NCvHhVbWwh5F0Jxq3S79mXqvnWOsXBefjKZVFBCnqjlMFsExv6WEifCQg==",
		PublicKey:  "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC0gCKnRQwBVi5vB5D22Hzj5mzPeqkyn0fdQerR0exVYsXonGW8NdQliMPDuenSkNMKLWU61gu+cr16Z9nU3ZEIGafXMrV2ejC534ir4/tN9sJwFN2TBEnPzp64X+fPZ1gukbyMturSENMlEeRy66RDfgOKlQ1andNE0GXPGEaB7QIDAQAB",
	})

	userPhone := "1" + fmt.Sprintf("%08d", 201120) + fmt.Sprintf("%02v", rand.Int31n(100))

	ret, err := vBill.Create(&VBillCreateConf{
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
	})
	fmt.Println(err)
	fmt.Printf("%+v", ret)
}
