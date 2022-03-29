package enter

import (
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"io/ioutil"
	"testing"
)

func TestEnter_Handle(t *testing.T) {
	en := new(CreateEnter)

	pfxData, _ := ioutil.ReadFile("1.pfx")

	en.InitBaseConfig(&BaseConfig{
		AgentNo:      "b1555484944055448",
		Key:          "DA726Z2OE2ZBNW0V",
		PfxData:      pfxData,
		CertPassWord: "584520lym.",
	})

	picInfo := new(PicInfo)
	//picInfo.LicenseDeadlinePic = "/b1537337771064377/2019-04-18/190418102220180800555342.png"       //营业执照号图片
	picInfo.LegalPersonIdFrontPic = "/b1537337771064377/2019-06-17/190617163751917222602737.png"    //法人身份正面照
	picInfo.LegalPersonIdOppositePic = "/b1537337771064377/2019-06-17/190617163732658441426829.png" //法人证件号反面
	picInfo.LegalPersonBanKCardPic = "/b1537337771064377/2019-06-17/190617162024759021422869.png"   //法人银行卡图片
	picInfo.OperatorIdDeadlineFrontPic = picInfo.LegalPersonIdFrontPic                              //经办人身份证正面
	picInfo.OperatorIdDeadlineOppositePic = picInfo.LegalPersonIdOppositePic                        //经办人身份证反面
	picInfo.MerchantDoorHeadPic = "/b1537337771064377/2019-06-17/190617164839769292477754.png"      //商户门头照片
	picInfo.MerchantFrontPic = "/b1537337771064377/2019-06-17/190617164839769292477754.png"         //商户门脸照片
	picInfo.MerchantInsidePic = "/b1537337771064377/2019-06-17/190617162242768106684871.png"        //商户内饰照片
	//picInfo.NoSealAgreement="/zl@cs.sh.cn/2018-07-25/1.png" //协议 - 未盖章
	//picInfo.SealAgreement="/zl@cs.sh.cn/2018-07-25/1.png" //协议 - 已盖章
	picInfo.ContractConfirm = "/b1552470439457613/2019-03-27/190327114024674450794181.png" //合同确认图片

	settleInfo := new(SettleInfo)
	settleInfo.SettleAccountType = "2"
	settleInfo.AccountName = "周超"
	settleInfo.BankCardNo = "6217001180010001264"
	settleInfo.SettleMode = "1"
	settleInfo.SettleModeValue = "B1"

	feeInfoList := []*FeeInfo{
		{
			FeeRate:     "2.5",
			FeeType:     "0",
			ProductType: "20010002-1001-106",
			SettleCycle: "T24",
			ChargeRole:  "1",
		},
		{
			FeeRate:     "2.5",
			FeeType:     "0",
			ProductType: "20010002-1001-107",
			SettleCycle: "T24",
			ChargeRole:  "1",
		},
		{
			FeeRate:     "2.5",
			FeeType:     "0",
			ProductType: "20010002-1001-108",
			SettleCycle: "T24",
			ChargeRole:  "1",
		},
		{
			FeeRate:     "2.5",
			FeeType:     "0",
			ProductType: "20010002-1001-203",
			SettleCycle: "T24",
			ChargeRole:  "1",
		},
		{
			FeeRate:     "2.5",
			FeeType:     "0",
			ProductType: "20010002-1001-204",
			SettleCycle: "T24",
			ChargeRole:  "1",
		},
		{
			FeeRate:     "2.5",
			FeeType:     "0",
			ProductType: "20010002-1001-206",
			SettleCycle: "T24",
			ChargeRole:  "1",
		},
	}

	ret, err := en.Handle(&CreateConf{
		LoginNo:      "135858210800621",
		AccountType:  "1",            //开户类型
		MerchantName: "重庆韬启信息科技有限公司", //商户名称
		//MerchantType: "8",            //个体工商户
		MerchantType: "9",                   //小微商户
		MccCode:      "5331",                //mcc码
		MerchantMail: "zhouc@ttouch.com.cn", //商户邮箱
		SendMail:     "1",                   //是否发送邮件
		PhoneNo:      "13585820180",         //联系人手机号
		//个体工商户
		//CompanyIdType:   "2",
		//LicenseNo:       "91310109324309418B",
		//LicenseDeadline: "2034-12-14",
		LicenseProvince: "500000",
		LicenseCity:     "500100",
		LicenseDistrict: "500105",

		LegalPersonName:       "周超",                               //法人名称
		LegalPersonId:         "500234199208102174",               //法人身份证号
		LegalPersonIdDeadline: "2023-01-15",                       //法人身份证有效期
		OperatorName:          "周超",                               //经办人姓名
		OperatorId:            "500234199208102174",               //经办身份证号
		OperatorIdDeadline:    "2024-12-11",                       //经办人身份证有效期
		MerchantContactsName:  "周超",                               //商户联系人
		MerchantAddress:       "重庆市江北区渝北三村30号24-11(仅限用于行政办公通讯联旧G", //商户地址
		ServicePhone:          "13585821080",                      //客服电话
		MerchantShortName:     "韬启便利店",                            //商户简称
		PicInfo:               picInfo,                            //图片信息
		SettleInfo:            settleInfo,                         //结算信息
		FeeInfo:               feeInfoList,                        //费率信息
		FeeStartTime:          "2019-04-15 00:00:00",              //费率生效开始时间
		FeeEndTime:            "2020-04-15 00:00:00",              //费率生效结束时间
		NotifyUrl:             "http://web.udian.me/v1/common/enter-notify",
	})
	fmt.Println(err)
	fmt.Printf("%+v", ret)
}
func TestEnter_BuildData(t *testing.T) {

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
