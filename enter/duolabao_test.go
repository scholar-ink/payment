package enter

import (
	"encoding/base64"
	"fmt"
	"github.com/scholar-ink/payment/util/http"
	"testing"
)

func TestDuoLaBaoEnter_Create(t *testing.T) {
	DuoLaBao := new(DuoLaBaoEnter)

	DuoLaBao.InitConfig(&DuoLaBaoConfig{
		AccessKey: "ecd1319dd3da43268ac96599a4ad232c1836f78b",
		SecretKey: "1f4b4ee723054facb34c8c16d188a4a2b3b53c1c",
	})

	ret, err := DuoLaBao.Create(&DuoLaBaoCreateConf{
		AgentNum:             "10001015856358316531057",
		FullName:             "赤壁市豆蓝百货店",
		ShortName:            "收钱啦收银",
		Industry:             "生活方式",
		Province:             "湖北",
		City:                 "咸宁",
		District:             "赤壁市",
		LinkMan:              "李红艳",
		LinkPhone:            "17714500630",
		CustomerType:         "INDIVIDUALBISS",
		CertificateType:      "IDENTIFICATION",
		CertificateCode:      "320830199012082427",
		CertificateName:      "李红艳",
		CertificateStartDate: "2015-11-17",
		CertificateEndDate:   "2025-11-17",
		ContactPhoneNum:      "17714500631",
		LinkManId:            "320830199012082427",
		PostalAddress:        "马港创新聚集区电商产业园A区A0764",
	})
	fmt.Println(err)
	fmt.Printf("%+v", ret)
}

func TestDuoLaBaoEnter_Settle(t *testing.T) {
	DuoLaBao := new(DuoLaBaoEnter)

	DuoLaBao.InitConfig(&DuoLaBaoConfig{
		AccessKey: "ecd1319dd3da43268ac96599a4ad232c1836f78b",
		SecretKey: "1f4b4ee723054facb34c8c16d188a4a2b3b53c1c",
	})

	ret, err := DuoLaBao.Settle(&DuoLaBaoSettleConf{
		CustomerNum:     "10001115872265337539380",
		BankAccountName: "李红艳",
		BankAccountNum:  "6217920471658935",
		Province:        "江苏",
		City:            "淮安",
		BankBranchName:  "上海浦东发展银行股份有限公司淮安分行",
		BankName:        "上海浦东发展银行",
		SettleAmount:    "1.00",
		PayBankList: []*DuoLaBaoBank{
			{
				Num:  "10031414639876930831004",
				Rate: "0.25",
			},
			{
				Num:  "10031414639876930831005",
				Rate: "0.25",
			},
		},
		AccountType:                 "PRIVATE",
		Phone:                       "17714500631",
		SettlerCertificateCode:      "320830199012082427",
		SettlerCertificateStartDate: "2015-11-17",
		SettlerCertificateEndDate:   "2025-11-17",
	})
	fmt.Println(err)
	fmt.Printf("%+v", ret)
}

func TestDuoLaBaoEnter_Shop(t *testing.T) {
	DuoLaBao := new(DuoLaBaoEnter)

	DuoLaBao.InitConfig(&DuoLaBaoConfig{
		AccessKey: "ecd1319dd3da43268ac96599a4ad232c1836f78b",
		SecretKey: "1f4b4ee723054facb34c8c16d188a4a2b3b53c1c",
	})

	ret, err := DuoLaBao.Shop(&DuoLaBaoShopConf{
		AgentNum:    "10001015856358316531057",
		CustomerNum: "10001115872265337539380",
		ShopName:    "赤壁市豆蓝百货店",
		Address:     "马港创新聚集区电商产业园A区A0764",
		OneIndustry: "购物",
		TwoIndustry: "批发市场",
		MobilePhone: "17714500631",
		MapLat:      "29.735429",
		MapLng:      "113.87867",
	})
	fmt.Println(err)
	fmt.Printf("%+v", ret)
}

func TestDuoLaBaoEnter_Complete(t *testing.T) {
	DuoLaBao := new(DuoLaBaoEnter)

	DuoLaBao.InitConfig(&DuoLaBaoConfig{
		AccessKey: "ecd1319dd3da43268ac96599a4ad232c1836f78b",
		SecretKey: "1f4b4ee723054facb34c8c16d188a4a2b3b53c1c",
	})

	ret, err := DuoLaBao.Complete(&DuoLaBaoCompleteConf{
		CustomerNum:      "10001115872265337539380",
		LicenseId:        "92421281MA4DLPXX9H",
		LicenseStartTime: "2019-09-29",
		LicenseEndTime:   "2099-12-30",
		CallbackUrl:      "http://tq.udian.me/v1/common/enter-notify",
	})
	fmt.Println(err)
	fmt.Printf("%+v", ret)
}

func TestDuoLaBaoEnter_Upload(t *testing.T) {
	DuoLaBao := new(DuoLaBaoEnter)

	DuoLaBao.InitConfig(&DuoLaBaoConfig{
		AccessKey: "ecd1319dd3da43268ac96599a4ad232c1836f78b",
		SecretKey: "1f4b4ee723054facb34c8c16d188a4a2b3b53c1c",
	})

	b, _ := http.GetBytes("https://cdn.51shouqianla.com/Fi-1BKl0t41qQYPP4WDXegn-ozbA!640compress")

	ret, err := DuoLaBao.Upload(&DuoLaBaoUploadConf{
		AttachType:  "LICENCE",
		CustomerNum: "10001115872265337539380",
		File:        base64.StdEncoding.EncodeToString(b),
	})
	fmt.Println(err)
	fmt.Printf("%+v", ret)
}

func TestDuoLaBaoEnter_IndustryList(t *testing.T) {
	DuoLaBao := new(DuoLaBaoEnter)

	DuoLaBao.InitConfig(&DuoLaBaoConfig{
		AccessKey: "ecd1319dd3da43268ac96599a4ad232c1836f78b",
		SecretKey: "1f4b4ee723054facb34c8c16d188a4a2b3b53c1c",
	})

	err := DuoLaBao.IndustryList()
	fmt.Println(err)
}
