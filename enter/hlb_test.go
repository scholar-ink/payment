package enter

import (
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"io/ioutil"
	"testing"
)

func TestHlbEnter_Create(t *testing.T) {
	hlb := new(HlbEnter)

	hlb.InitConfig(&HlbConfig{
		MerchantNo: "C1801372489",
		Key:        "r8g71ZarxNqS0UWWW4Yj4CXi",
		SignKey:    "e5dxYLrhOQweTvR2NiPeKd6gEehGlBOE",
	})

	ret, err := hlb.Create(&HlbCreateConf{
		OrderNo:           helper.CreateSn(),
		SignName:          "赤壁市豆蓝百货店",
		ShowName:          "赤壁市豆蓝百货店",
		MerchantType:      "PERSON",
		LegalPerson:       "李红艳",
		LegalPersonID:     "320830199012082427",
		OrgNum:            "320830199012082427",
		BusinessLicense:   "320830199012082427",
		RegionCode:        "171106",
		Address:           "马港创新聚集区电商产业园A区A0764",
		Linkman:           "李红艳",
		LinkPhone:         "17714500631",
		Email:             "17714500632@udian.me",
		BankCode:          "310308000019",
		AccountName:       "李红艳",
		AccountNo:         "6217920471658935",
		SettleBankType:    "TOPRIVATE",
		SettlementPeriod:  "D1",
		SettlementMode:    "AUTO",
		MerchantCategory:  "OFFLINE_RETAIL",
		IndustryTypeCode:  "486",
		AuthorizationFlag: true,
		NeedPosFunction:   false,
		IdCardStartDate:   "20151117",
		IdCardEndDate:     "20351117",
		BusinessDateStart: "",
		BusinessDateLimit: "",
		AccountIdCard:     "",
		Mcc:               "",
		AgreeProtocol:     true,
		SettleMode:        "DEFAULT",
	})
	fmt.Println(err)
	fmt.Printf("%+v", ret)
}

func TestHlbUpload_Handle(t *testing.T) {
	hlb := new(HlbEnter)

	hlb.InitConfig(&HlbConfig{
		MerchantNo: "C1801372489",
		Key:        "r8g71ZarxNqS0UWWW4Yj4CXi",
		SignKey:    "e5dxYLrhOQweTvR2NiPeKd6gEehGlBOE",
	})

	b, _ := ioutil.ReadFile("inside.png")

	ret, err := hlb.Upload(&HlbUploadConf{
		MerchantNo:     "E1801435392",
		OrderNo:        helper.CreateSn(),
		CredentialType: "INTERIOR_PHOTO",
		FileSign:       helper.Md5(string(b)),
		File:           b,
	})

	fmt.Printf("%+v\n", ret)
	fmt.Println(err)
}

func TestHlbEnter_OpenProduct(t *testing.T) {
	hlb := new(HlbEnter)

	hlb.InitConfig(&HlbConfig{
		MerchantNo: "C1801372489",
		Key:        "r8g71ZarxNqS0UWWW4Yj4CXi",
		SignKey:    "e5dxYLrhOQweTvR2NiPeKd6gEehGlBOE",
	})
	err := hlb.OpenProduct(&HlbOpenProductConf{
		ProductType: "APPPAY",
		MerchantNo:  "E1801448490",
		PayType:     "SCAN",
		AppPayType:  "QQPAY",
		Value:       "0.25",
	})

	fmt.Println(err)
}

func TestHlbEnter_Modify(t *testing.T) {
	en := new(HlbEnter)

	en.InitConfig(&HlbConfig{
		MerchantNo: "b1564976173998427",
		Key:        "YC4QSNWEls84sJzwhPKqPtz7fm0BPA07",
	})

	feeInfoList := []*HlbFeeInfo{
		{
			FeeRate:     "0",
			FeeType:     "1",
			ProductType: "30040004-3001-131",
			SettleCycle: "D0",
			ChargeRole:  "1",
		},
		{
			FeeRate:     "4.2",
			FeeType:     "0",
			ProductType: "20010002-1001-106",
			SettleCycle: "D0",
			ChargeRole:  "1",
		},
		{
			FeeRate:     "4.2",
			FeeType:     "0",
			ProductType: "20010002-1001-107",
			SettleCycle: "D0",
			ChargeRole:  "1",
		},
		{
			FeeRate:     "4.2",
			FeeType:     "0",
			ProductType: "20010002-1001-108",
			SettleCycle: "D0",
			ChargeRole:  "1",
		},
		{
			FeeRate:     "4.2",
			FeeType:     "0",
			ProductType: "20010002-1001-203",
			SettleCycle: "D0",
			ChargeRole:  "1",
		},
		{
			FeeRate:     "4.2",
			FeeType:     "0",
			ProductType: "20010002-1001-204",
			SettleCycle: "D0",
			ChargeRole:  "1",
		},
		{
			FeeRate:     "4.2",
			FeeType:     "0",
			ProductType: "20010002-1001-206",
			SettleCycle: "D0",
			ChargeRole:  "1",
		},
	}

	settleInfo := new(HlbSettleInfo)
	settleInfo.SettleAccountType = "2"
	settleInfo.AccountName = "周丹"
	settleInfo.BankCardNo = "6230200168071087"
	settleInfo.SettleMode = "2"
	//settleInfo.SettleModeValue = "A1"

	ret, err := en.Modify(&HlbModifyConf{
		AccountNo:    "b14385747570726912",
		FeeInfo:      feeInfoList,
		SettleInfo:   settleInfo,
		FeeStartTime: "2019-08-19 00:00:00", //费率生效开始时间
		FeeEndTime:   "2020-08-19 00:00:00", //费率生效结束时间
	})

	fmt.Printf("%+v", ret)
	fmt.Println(err)
}

func TestHlbEnter_ModifySettleInfo(t *testing.T) {
	en := new(HlbEnter)

	en.InitConfig(&HlbConfig{
		MerchantNo: "b1564976173998427",
		Key:        "AT29WBOWHI02LAPX",
	})

	ret, err := en.ModifySettleInfo(&HlbModifySettleInfoConf{
		AccountNo:         "b14385747570726912",
		RequestNo:         helper.CreateSn(),
		SettleAccountType: "2",
		AccountName:       "周丹",
		BankCardNo:        "6230200168071087",
		SettleMode:        "2",
		NotifyUrl:         "http://web.yunc.udian.me/v1/common/enter-notify",
	})

	fmt.Printf("%+v", ret)
	fmt.Println(err)
}
func TestHlbEnter_ModifyFeeInfo(t *testing.T) {
	en := new(HlbEnter)

	en.InitConfig(&HlbConfig{
		MerchantNo: "b1564976173998427",
		Key:        "AT29WBOWHI02LAPX",
	})

	feeInfoList := []*HlbFeeInfo{
		{
			FeeRate:     "0",
			FeeType:     "1",
			ProductType: "30040004-3001-131",
			SettleCycle: "D0",
			ChargeRole:  "1",
		},
		{
			FeeRate:     "4.2",
			FeeType:     "0",
			ProductType: "20010002-1001-106",
			SettleCycle: "D0",
			ChargeRole:  "1",
		},
		{
			FeeRate:     "4.2",
			FeeType:     "0",
			ProductType: "20010002-1001-107",
			SettleCycle: "D0",
			ChargeRole:  "1",
		},
		{
			FeeRate:     "4.2",
			FeeType:     "0",
			ProductType: "20010002-1001-108",
			SettleCycle: "D0",
			ChargeRole:  "1",
		},
		{
			FeeRate:     "4.2",
			FeeType:     "0",
			ProductType: "20010002-1001-203",
			SettleCycle: "D0",
			ChargeRole:  "1",
		},
		{
			FeeRate:     "4.2",
			FeeType:     "0",
			ProductType: "20010002-1001-204",
			SettleCycle: "D0",
			ChargeRole:  "1",
		},
		{
			FeeRate:     "4.2",
			FeeType:     "0",
			ProductType: "20010002-1001-206",
			SettleCycle: "D0",
			ChargeRole:  "1",
		},
	}

	ret, err := en.ModifyFeeInfo(&HlbModifyFeeInfoConf{
		AccountNo:    "b14380425205159936",
		RequestNo:    helper.CreateSn(),
		FeeInfo:      feeInfoList,           //费率信息
		FeeStartTime: "2019-04-15 00:00:00", //费率生效开始时间
		FeeEndTime:   "2020-04-15 00:00:00", //费率生效结束时间
		NotifyUrl:    "http://web.yunc.udian.me/v1/payment/hlb/d0-enter-notify",
	})

	fmt.Printf("%+v", ret)
	fmt.Println(err)
}

func TestHlbEnter_Query(t *testing.T) {
	en := new(HlbEnter)

	en.InitConfig(&HlbConfig{
		MerchantNo: "b1564976173998427",
		Key:        "AT29WBOWHI02LAPX",
	})

	ret, err := en.Query(&HlbQueryConf{
		AccountNo: "b14386708932286464",
	})

	fmt.Println(ret)
	fmt.Println(err)
}

func TestHlbEnter_Query2(t *testing.T) {
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
