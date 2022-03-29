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

	pfxData, _ := ioutil.ReadFile("zg.pfx")

	encrptData, err := helper.Rsa1Encrypt(pfxData, []byte("1111"), "Udian888")
	//
	fmt.Println(encrptData)
	fmt.Println(err)

	//encrptData := "s+UtXL6//g7FSICUKBshy5qNXnOnMXNVHlVefW4OcDDr3+i0RpIuU/ZeO9Fr3kEbVHM9uw69HQfXdgsE2F7wPlB+SD1h5dAAzMhVEaFHsbscqtepTeaZAUYkOm80aFypYqv+8dhKSVZtomcd1KqUdrDR/KJnEcpe2Cr8RQhg16wOu5Fz8EfTkThkeMwDFDwlM3OxLg8wC1qGItSBOPuvrkpIM4fwEVkNGV8JnJsYHFBLprHGwcjfEnEWZL0eNTnMePIUCnkCD8qtzVMeQMX1O5SjBwZPNXfTvvwVicR6IWnrLGeg9U8RzUND48N7M5YURK+aRmPJRNRo6gDln83Uxg=="
	//
	decrptData, err := helper.Rsa1Decrypt(pfxData, encrptData, "Taoqi888")

	fmt.Println(string(decrptData))
	fmt.Println(err)
}
