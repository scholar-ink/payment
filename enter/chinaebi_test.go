package enter

import (
	"fmt"
	"github.com/scholar-ink/payment/util/http"
	"strings"
	"testing"
	"time"
)

func TestChinaEbiEnter_Create(t *testing.T) {
	ChinaEbi := new(ChinaEbiEnter)

	ChinaEbi.InitConfig(&ChinaEbiConfig{
		OrgNumber:   "999",
		AgentNumber: "999",
		PrivateKey:  "MIICeQIBADANBgkqhkiG9w0BAQEFAASCAmMwggJfAgEAAoGBAM4/rC99irvS//RV2SMd4KHzhuqknbcpeJWMXC4FvWt+n06f7FP5ZWemx8VAmEL1/n41jOPbLTgSacuFQQmemsCp5H8d2g/XBqYy1q3f/0qPCi8Tt1yop5efcumjma2jcRdAkSyQClc8GFmEeRkEl19chrEpI/6bxtXoBD2dDLpZAgMBAAECgYEAs4btBDGND0ztCuunJFAfdhkaeShtOD/a/KG+ozjP1r/TP4cpGTdfM0gTX/mID9E8gvNt/fCMfeBZQpRtNkhefoGvlij9SSVJMyvCtqaGEbUzvaBhzLkko+d7YdZeeUg6SWAPTnPSSG3nXRwR3pMa7TqFiVpBJN8HyV8mE2fMznkCQQDsredSzLmkKv/sEZaZNngOfQxjCfBME8+Gnoz+LsQwFoXRNxKToGw1x/yt+cjWAkn+FV9Nb98oPkqPjLdKdDsnAkEA3xXVh/S1R42k1gqu0UKB1hHSQ0EUGqg9i8vWqATf0SomC4SB0K7mgmEUNDrka9KNnw1Bc9TL6w/zmPKCcibOfwJBANMjvMan7kCfP5n4gtIBvo6mTcOYnS8xSSQ+E2e6jribjxt6Nu9N4NsFoswNlnYcqqepp1Bsqba8A0YWcXlRQWcCQQDaYYpdg+yttfgF3AFUMlHdWCbH1X4ztjxBjHJ+mf7rx+HkZnuZ6I0YVqYrlvciocQnThejp01Tt5LUR5nw2xJLAkEAzoOMTiGDRQGh1BHGtbfMO0U7abma+T2XPYUUCQA8QH3B4Lr0Zm1JJ1mcPVsM3yVgBDQqiLXYIy4nKJyvn5kLWg==",
	})

	bZZ1, _ := http.GetBytes("https://cdn.51shouqianla.com/Fo8S0jtnr61PSB1R6IVfeF0OfAwx")
	bSFZ1, _ := http.GetBytes("https://cdn.51shouqianla.com/Fgr8zbk1FYCwAHq9Qh6jG6kkbDfp")
	bSFZ2, _ := http.GetBytes("https://cdn.51shouqianla.com/FpuCi_i4PkdJaMcMHdojGZ8tl0kK")
	bCDJJ1, _ := http.GetBytes("http://cdn.51shouqianla.com/FlSfDiuPpvHFJSBbTs00yxffpO2e")
	bCDNJ1, _ := http.GetBytes("https://cdn.51shouqianla.com/FkOLBp_iqJDiBmzyvWmvWphwuKB8")
	bYHK, _ := http.GetBytes("https://cdn.51shouqianla.com/FtrT_WG3WToLDCj-ljik7Py_f3Qf")

	//requestId := strings.ReplaceAll(uuid.New().String(), "-", "")

	ret, err := ChinaEbi.Create(&ChinaEbiCreateConf{
		SeqNo:         strings.Replace(time.Now().Format("060102150405.000000"), ".", "", -1),
		MercMbl:       "13585821080",
		MercCnm:       "青白江朝团清日用品经营部",
		MercAbbr:      "朝团清日用品",
		MercHotLin:    "0591-88871222",
		MccCd:         "5311",
		MercProv:      "6500",
		MercCity:      "6510",
		MercCounty:    "6528",
		BusAddr:       "白江区香岛大道1509号现代物流大厦A区1楼18号",
		MercAttr:      "1",
		RegId:         "92510113MA69R5WW2D",
		RegExpDtD:     "2099-12-12",
		CrpIdTyp:      "0",
		CrpIdNo:       "500234199208102174",
		CrpNm:         "周超",
		CrpExpDtD:     "2023-01-15",
		StlSign:       "1",
		StlWcLbnkNo:   "105290078057",
		StlOac:        "6217001180010001264",
		BnkAcnm:       "周超",
		UsrOprEmail:   "zhouc@ttouch.com.cn",
		DebitFee:      "0.001",
		DebitFeeLimit: "99999999",
		CreditFee:     "0.001",
		D0Fee:         "0",
		D0FeeQuota:    "0",
		AliFee:        "0.0025",
		WxFee:         "0.0025",
		AliFlg:        "1",
		WxFlg:         "1",
		UnionFlg:      "0",
		OutMercId:     "203184",
		SettType:      "T1",
		ZZ1:           bZZ1,
		SFZ1:          bSFZ1,
		SFZ2:          bSFZ2,
		CDJJ1:         bCDJJ1,
		CDMT1:         bCDJJ1,
		CDNJ1:         bCDNJ1,
		YHK:           bYHK,
	})
	fmt.Println(err)
	fmt.Printf("%+v", ret)
}

func TestChinaEbiEnter_Query(t *testing.T) {
	ChinaEbi := new(ChinaEbiEnter)

	ChinaEbi.InitConfig(&ChinaEbiConfig{
		OrgNumber:   "999",
		AgentNumber: "999",
		PrivateKey:  "MIICeQIBADANBgkqhkiG9w0BAQEFAASCAmMwggJfAgEAAoGBAM4/rC99irvS//RV2SMd4KHzhuqknbcpeJWMXC4FvWt+n06f7FP5ZWemx8VAmEL1/n41jOPbLTgSacuFQQmemsCp5H8d2g/XBqYy1q3f/0qPCi8Tt1yop5efcumjma2jcRdAkSyQClc8GFmEeRkEl19chrEpI/6bxtXoBD2dDLpZAgMBAAECgYEAs4btBDGND0ztCuunJFAfdhkaeShtOD/a/KG+ozjP1r/TP4cpGTdfM0gTX/mID9E8gvNt/fCMfeBZQpRtNkhefoGvlij9SSVJMyvCtqaGEbUzvaBhzLkko+d7YdZeeUg6SWAPTnPSSG3nXRwR3pMa7TqFiVpBJN8HyV8mE2fMznkCQQDsredSzLmkKv/sEZaZNngOfQxjCfBME8+Gnoz+LsQwFoXRNxKToGw1x/yt+cjWAkn+FV9Nb98oPkqPjLdKdDsnAkEA3xXVh/S1R42k1gqu0UKB1hHSQ0EUGqg9i8vWqATf0SomC4SB0K7mgmEUNDrka9KNnw1Bc9TL6w/zmPKCcibOfwJBANMjvMan7kCfP5n4gtIBvo6mTcOYnS8xSSQ+E2e6jribjxt6Nu9N4NsFoswNlnYcqqepp1Bsqba8A0YWcXlRQWcCQQDaYYpdg+yttfgF3AFUMlHdWCbH1X4ztjxBjHJ+mf7rx+HkZnuZ6I0YVqYrlvciocQnThejp01Tt5LUR5nw2xJLAkEAzoOMTiGDRQGh1BHGtbfMO0U7abma+T2XPYUUCQA8QH3B4Lr0Zm1JJ1mcPVsM3yVgBDQqiLXYIy4nKJyvn5kLWg==",
	})

	ret, err := ChinaEbi.Query(&ChinaEbiQueryConf{
		DyMchNo: "999652853110002",
	})

	fmt.Println(ret)
	fmt.Println(err)
}
