package enter

import (
	"encoding/base64"
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"
	"io/ioutil"
	"testing"
)

func TestSumPayEnter_Create(t *testing.T) {
	enter := new(SumPayEnter)

	pfxData, _ := ioutil.ReadFile("shouqianla.pfx")

	fmt.Println(string(pfxData))

	aesKey := helper.NonceStr()

	enter.InitConfig(&SumPayConfig{
		AppId:        "101803663",
		AesKey:       aesKey,
		PfxData:      pfxData,
		CertPassWord: "shouqianla",
		PublicKey:    "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA5iO3Rg9G10gJp9uB9svUDo8TVcMfvd/rCa9mrw4Ai1veh8hEb8Xk5LYQWd9g1DgpSBFjHhe/iO7h40ZWC/STgrmitH4R9K5vEFYBo6mJpVcKyvCimbtp5dcwdLojw7BwRr7Y/7STPSDJ270mVqjNqXPAenTigQ8ay4LChcPtFZUnSWb2DL92d/o6+26gS/eqed3F28TR4WtFdxPxGZy6oTtZ2B8jI9scCFfzFlsF/tqeExCmLnJbGUNaLUDPMQIe02/W+mZOiC2aeAfOkC72Z3leczEm7Wvp4fXKsRV6XmCs6/DrEMY6RjHVWocXCtYxAMY3HOAiL+7voSLZWn1ArwIDAQAB",
	})

	ret, err := enter.Create(&SumPayCreateConf{
		MerNo:       "101803663",
		OpType:      "1",
		ResideMerNo: "101813675",
		BaseInfo: &BaseInfo{
			CompanyName:            "赤壁市考升百货店",
			CompanyAbbrName:        "赤壁市考升百货店",
			ComType:                "2",
			SocialCreditCode:       "92421281MA4DLPXRXB",
			LicenseExpiredDate:     "长期",
			CompanyRepresentative:  enter.EncryptFiled("李红艳"),
			ComRepIdType:           "1",
			ComRepIdNo:             enter.EncryptFiled("320830199012082427"),
			ComRepIdExpDate:        "2025-11-17",
			ComProv:                "420000",
			ComCity:                "421200",
			Address:                "马港创新聚集区电商产业园A区A0737",
			Email:                  "zhouc1@ttouch.com.cn",
			PostCode:               "000000",
			RegisteredCapital:      "0万元",
			RegisteredAddress:      "马港创新聚集区电商产业园A区A0737",
			BusinessScope:          "日用百货、服装、饰品、箱包、鞋帽、皮具、家居用品、家具、玩具、厨房用品、日化用品、体育用品、办公用品零售(含网络销售)(涉及许可经营项目应取得相关部门许可后方可经营)",
			ControllingShareholder: "李红艳",
			RegisteredTime:         "2019-08-29",
			AccessSite:             "/",
			AccessSiteIp:           "0.0.0.0",
			ParentIndustryCatagory: "02",
			SubIndustryCatagory:    "0200",
			AgentPayIp:             "0.0.0.0",
			Nature:                 "05",
			AnnualFee:              "0",
			TradePattern:           "T0002",
			PasswordFreePay:        "0",
			ShareBenefitMode:       "01",
			AuthCertPic:            "crm/101803663/AUTH_CERT_PIC-915ef2e4-c7fd-46dc-ba89-cd62b54e3bc9.png",
			BankAccountPermitType:  "2",
			BankAccountPermitPic:   "crm/101803663/BANK_ACCOUNT_PERMIT_PIC-f1dbe061-da5d-4ed2-9e46-9c4b7c224056.png",
			LicensePic:             "crm/101803663/LICENSE_PIC-4bf6f914-f0aa-427a-841a-32b6dc5f0a26.png",
			ComRepIdFrontPic:       "crm/101803663/ID_PIC-e53247f7-cc4a-4a60-b2a0-f4cc64cb7fdf.png",
			ComRepIdBackPic:        "crm/101803663/ID_PIC-63d3f61a-1931-4a66-834c-74c168e12454.png",
			Icp:                    "crm/101803663/ICP_PIC-b07632a7-9dd1-4b11-bbe4-25450343d813.png",
			BeneficiaryIdentify:    "00",
		},
		SettleInfo: &SettleInfo{
			AccountType:       "1",
			RealName:          enter.EncryptFiled("李红艳"),
			BankCode:          "SPDB",
			AccountNo:         enter.EncryptFiled("6217920471658935"),
			OpenProv:          "320000",
			OpenCity:          "320800",
			IdNo:              enter.EncryptFiled("320830199012082427"),
			SettleType:        "1",
			SettlePeriod:      "0",
			ReturnWay:         "0",
			SettleWithdrawWay: "00",
		},
		ContactInfo: &ContactInfo{
			Contact:         "李红艳",
			ContactIdNo:     enter.EncryptFiled("320830199012082427"),
			ContactTel:      "17714500631",
			ContactEmail:    "zhouc@ttouch.com.cn",
			ComRiskName:     "李红艳",
			ComRiskTel:      "17714500631",
			ComRiskEmail:    "zhouc@ttouch.com.cn",
			ComOperateName:  "李红艳",
			ComOperateTel:   "17714500631",
			ComOperateEmail: "zhouc@ttouch.com.cn",
		},
		ReplenishInfo: &ReplenishInfo{
			Bond: "0",
			Bd:   "0001",
		},
		ProductInfoList: []*ProductInfo{
			{
				TradeProCode: "20301",
				PayProCode:   "90501",
				ChargeWay:    "A",
				Rate:         "0.23",
			},
			{
				TradeProCode: "20301",
				PayProCode:   "90504",
				ChargeWay:    "A",
				Rate:         "0.23",
			},
			{
				TradeProCode: "20301",
				PayProCode:   "90601",
				ChargeWay:    "A",
				Rate:         "0.23",
			},
		},
		BeneficiaryInfoList: []*BeneficiaryInfo{
			{
				BfyName:      enter.EncryptFiled("李红艳"),
				BfyIdType:    "1",
				BfyIdNo:      enter.EncryptFiled("320830199012082427"),
				BfyIdExpDate: "2025-11-17",
				BfyProv:      "320000",
				BfyCity:      "320800",
				BfyAddress:   "马港创新聚集区电商产业园A区A0764",
			},
		},
	})
	fmt.Println(err)
	fmt.Printf("%+v", ret)
}

func TestSumPayEnter_Query(t *testing.T) {
	enter := new(SumPayEnter)

	pfxData, _ := ioutil.ReadFile("shouqianla.pfx")

	aesKey := helper.NonceStr()

	enter.InitConfig(&SumPayConfig{
		AppId:        "101803663",
		AesKey:       aesKey,
		PfxData:      pfxData,
		CertPassWord: "shouqianla",
		PublicKey:    "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA5iO3Rg9G10gJp9uB9svUDo8TVcMfvd/rCa9mrw4Ai1veh8hEb8Xk5LYQWd9g1DgpSBFjHhe/iO7h40ZWC/STgrmitH4R9K5vEFYBo6mJpVcKyvCimbtp5dcwdLojw7BwRr7Y/7STPSDJ270mVqjNqXPAenTigQ8ay4LChcPtFZUnSWb2DL92d/o6+26gS/eqed3F28TR4WtFdxPxGZy6oTtZ2B8jI9scCFfzFlsF/tqeExCmLnJbGUNaLUDPMQIe02/W+mZOiC2aeAfOkC72Z3leczEm7Wvp4fXKsRV6XmCs6/DrEMY6RjHVWocXCtYxAMY3HOAiL+7voSLZWn1ArwIDAQAB",
	})

	ret, err := enter.Query(&SumPayQueryConf{
		MerNo:            "101803663",
		ResideMerchantNo: "101973662",
	})
	fmt.Println(err)
	fmt.Printf("%+v", ret)
}
func TestSumValidate(t *testing.T) {
	enter := new(SumPayEnter)

	pfxData, _ := ioutil.ReadFile("shouqianla.pfx")

	aesKey := helper.NonceStr()

	enter.InitConfig(&SumPayConfig{
		AppId:        "101803663",
		AesKey:       aesKey,
		PfxData:      pfxData,
		CertPassWord: "shouqianla",
		PublicKey:    "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA5iO3Rg9G10gJp9uB9svUDo8TVcMfvd/rCa9mrw4Ai1veh8hEb8Xk5LYQWd9g1DgpSBFjHhe/iO7h40ZWC/STgrmitH4R9K5vEFYBo6mJpVcKyvCimbtp5dcwdLojw7BwRr7Y/7STPSDJ270mVqjNqXPAenTigQ8ay4LChcPtFZUnSWb2DL92d/o6+26gS/eqed3F28TR4WtFdxPxGZy6oTtZ2B8jI9scCFfzFlsF/tqeExCmLnJbGUNaLUDPMQIe02/W+mZOiC2aeAfOkC72Z3leczEm7Wvp4fXKsRV6XmCs6/DrEMY6RjHVWocXCtYxAMY3HOAiL+7voSLZWn1ArwIDAQAB",
	})

	ret, err := enter.Validate(&SumValidateConf{
		MerNo:            "101803663",
		ResideMerchantNo: "101893707",
		RcvAmount:        "0.72",
	})
	fmt.Println(err)
	fmt.Printf("%+v", ret)
}
func TestSumContract(t *testing.T) {
	enter := new(SumPayEnter)

	pfxData, _ := ioutil.ReadFile("shouqianla.pfx")

	aesKey := helper.NonceStr()

	enter.InitConfig(&SumPayConfig{
		AppId:        "101803663",
		AesKey:       aesKey,
		PfxData:      pfxData,
		CertPassWord: "shouqianla",
		PublicKey:    "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA5iO3Rg9G10gJp9uB9svUDo8TVcMfvd/rCa9mrw4Ai1veh8hEb8Xk5LYQWd9g1DgpSBFjHhe/iO7h40ZWC/STgrmitH4R9K5vEFYBo6mJpVcKyvCimbtp5dcwdLojw7BwRr7Y/7STPSDJ270mVqjNqXPAenTigQ8ay4LChcPtFZUnSWb2DL92d/o6+26gS/eqed3F28TR4WtFdxPxGZy6oTtZ2B8jI9scCFfzFlsF/tqeExCmLnJbGUNaLUDPMQIe02/W+mZOiC2aeAfOkC72Z3leczEm7Wvp4fXKsRV6XmCs6/DrEMY6RjHVWocXCtYxAMY3HOAiL+7voSLZWn1ArwIDAQAB",
	})

	ret, err := enter.Contract(&SumContractConf{
		MerNo:            "101803663",
		ResideMerchantNo: "101893707",
	})
	fmt.Println(err)
	fmt.Printf("%+v", ret)
}

func TestSumPayUpload_Handle(t *testing.T) {

	enter := new(SumPayEnter)

	pfxData, _ := ioutil.ReadFile("shouqianla.pfx")

	aesKey := helper.NonceStr()

	enter.InitConfig(&SumPayConfig{
		AppId:        "101803663",
		AesKey:       aesKey,
		PfxData:      pfxData,
		CertPassWord: "shouqianla",
		PublicKey:    "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA5iO3Rg9G10gJp9uB9svUDo8TVcMfvd/rCa9mrw4Ai1veh8hEb8Xk5LYQWd9g1DgpSBFjHhe/iO7h40ZWC/STgrmitH4R9K5vEFYBo6mJpVcKyvCimbtp5dcwdLojw7BwRr7Y/7STPSDJ270mVqjNqXPAenTigQ8ay4LChcPtFZUnSWb2DL92d/o6+26gS/eqed3F28TR4WtFdxPxGZy6oTtZ2B8jI9scCFfzFlsF/tqeExCmLnJbGUNaLUDPMQIe02/W+mZOiC2aeAfOkC72Z3leczEm7Wvp4fXKsRV6XmCs6/DrEMY6RjHVWocXCtYxAMY3HOAiL+7voSLZWn1ArwIDAQAB",
	})

	b, err := http.GetBytes("https://cdn.51shouqianla.com/FkBloSC57pXj07uv-5femehyn5Dg")

	if err != nil {
		fmt.Println(err)
		return
	}

	ret, err := enter.Upload(&SumPayUploadConf{
		MerNo:   "101803663",
		PicType: "1",
		PicUse:  "3",
		PicFile: b,
	})

	fmt.Printf("%+v\n", ret)
	fmt.Println(err)
}
func TestZgComposeAutograph2(t *testing.T) {
	b, _ := base64.StdEncoding.DecodeString(`eyJjb21wYW55X25hbWUiOiLnm7HnnJnotqPnjqnnvZHnu5znp5HmioDmnInpmZDlhazlj7giLCJjb21wYW55X2FiYnJfbmFtZSI6Iui2o+eOqee9kee7nCIsImNvbV90eXBlIjoiMiIsInNvY2lhbF9jcmVkaXRfY29kZSI6IjkxMzIwODMwTUEyMFlZVEc0USIsImxpY2Vuc2VfZXhwaXJlZF9kYXRlIjoi6ZW/5pyfIiwiY29tcGFueV9yZXByZXNlbnRhdGl2ZSI6ImdXVDdsSXNVMFpRczZuZVozdmxKQ3c9PSIsImNvbV9yZXBfaWRfdHlwZSI6IjEiLCJjb21fcmVwX2lkX25vIjoiajBtOEFYalNGOHVRdUJrSVVzUkJodWRCbkl6QlZsSnQwMUxEam00bGJiQT0iLCJjb21fcmVwX2lkX2V4cF9kYXRlIjoiMjAyMi0wNy0yMyIsImNvbV9wcm92IjoiMzIwMDAwIiwiY29tX2NpdHkiOiIzMjA4MDAiLCJhZGRyZXNzIjoi55ux5Z+O5L2w5Lq/6LSi5a+M5bm/5Zy65LiA5pyfMzQxLTM0NeWPtyIsImVtYWlsIjoic2hvcF8yMDI3MjJAdHRvdWNoLmNvbS5jbiIsInBvc3RfY29kZSI6IjAwMDAwMCIsInJlZ2lzdGVyZWRfY2FwaXRhbCI6IjEwMOS4h+WFgyIsInJlZ2lzdGVyZWRfYWRkcmVzcyI6IuebseWfjuS9sOS6v+i0ouWvjOW5v+WcuuS4gOacnzM0MS0zNDXlj7ciLCJidXNpbmVzc19zY29wZSI6IuS6kuiBlOe9keS4iue9keacjeWKoeesrOS6jOexu+WinuWAvOeUteS/oeS4muWKoee9kee7nOaWh+WMlue7j+iQpeS6kuiBlOe9keS/oeaBr+acjeWKoSjkvp3ms5Xpobvnu4/mibnlh4bnmoTpobnnm67nu4/nm7jlhbPpg6jpl6jmibkiLCJjb250cm9sbGluZ19zaGFyZWhvbGRlciI6Iue9l+aXuiIsInJlZ2lzdGVyZWRfdGltZSI6IjIwMjAtMDMtMTEiLCJhY2Nlc3Nfc2l0ZSI6Ind3dy43MWthLmNvbSIsImFjY2Vzc19zaXRlX2lwIjoiMC4wLjAuMCIsInBhcmVudF9pbmR1c3RyeV9jYXRhZ29yeSI6IjAyIiwic3ViX2luZHVzdHJ5X2NhdGFnb3J5IjoiMDIwMSIsImFnZW50X3BheV9pcCI6IjAuMC4wLjAiLCJuYXR1cmUiOiIwMSIsImFubnVhbF9mZWUiOiIwIiwidHJhZGVfcGF0dGVybiI6IlQwMDAyIiwicGFzc3dvcmRfZnJlZV9wYXkiOiIwIiwic2hhcmVfYmVuZWZpdF9tb2RlIjoiMDEiLCJhdXRoX2NlcnRfcGljIjoiY3JtLzEwMTgwMzY2My9BVVRIX0NFUlRfUElDLTFiZGYzMTc1LTkxYmQtNGYyOS1iODNhLWIyMTgxZGFiZDI0NS5wbmciLCJiYW5rX2FjY291bnRfcGVybWl0X3R5cGUiOiIxIiwiYmFua19hY2NvdW50X3Blcm1pdF9waWMiOiJjcm0vMTAxODAzNjYzL0JBTktfQUNDT1VOVF9QRVJNSVRfUElDLWNkYWQyYTIxLTZiZTktNDdhZC05YTk0LTczNWM2ZjFiMzdhMi5wbmciLCJsaWNlbnNlX3BpYyI6ImNybS8xMDE4MDM2NjMvTElDRU5TRV9QSUMtZDBkMmVlNmItNjc5Mi00ZTgzLThmNDMtOWI1MWJlMzRmZjI1LnBuZyIsImNvbV9yZXBfaWRfZnJvbnRfcGljIjoiY3JtLzEwMTgwMzY2My9JRF9QSUMtNDQ1NWM3NzktOTc5Mi00YTI5LWJmNWUtODY5OTU3MTEyOTIwLnBuZyIsImNvbV9yZXBfaWRfYmFja19waWMiOiJjcm0vMTAxODAzNjYzL0lEX1BJQy02NWU3OTQ0NS01ZDEyLTRjODUtOWY1Yy0zYWNhNzVjOTFiMzgucG5nIiwiaWNwIjoiY3JtLzEwMTgwMzY2My9JQ1BfUElDLTcyMDg4ODliLTg2OWMtNGI3Yi1iODIzLTRlNDdhYmNjMjYxZS5wbmciLCJidXNpbmVzc19wbGFjZV9waWMiOiJjcm0vMTAxODAzNjYzL1NFVFRMRV9DQVJEX1BJQy0wNDc5YjI1Yy02Nzg2LTQ5OTQtOGUzZS1iYjZkZGM2NmQ0YmMucG5nIiwiYmVuZWZpY2lhcnlfaWRlbnRpZnkiOiIwMCJ9&beneficiary_info_list=W3siYmZ5X25hbWUiOiJnV1Q3bElzVTBaUXM2bmVaM3ZsSkN3PT0iLCJiZnlfaWRfdHlwZSI6IjEiLCJiZnlfaWRfbm8iOiJqMG04QVhqU0Y4dVF1QmtJVXNSQmh1ZEJuSXpCVmxKdDAxTERqbTRsYmJBPSIsImJmeV9pZF9leHBfZGF0ZSI6IjIwMjItMDctMjMiLCJiZnlfcHJvdiI6IjMyMDAwMCIsImJmeV9jaXR5IjoiMzIwODAwIiwiYmZ5X2FkZHJlc3MiOiLnm7Hln47kvbDkur/otKLlr4zlub/lnLrkuIDmnJ8zNDEtMzQ15Y+3In1d&contact_info=eyJjb250YWN0Ijoi572X5pe6IiwiY29udGFjdF9pZF9ubyI6ImowbThBWGpTRjh1UXVCa0lVc1JCaHVkQm5JekJWbEp0MDFMRGptNGxiYkE9IiwiY29udGFjdF90ZWwiOiIxNTk1MjM3OTc4NCIsImNvbnRhY3RfZW1haWwiOiJ6aG91Y0B0dG91Y2guY29tLmNuIiwiY29tX3Jpc2tfbmFtZSI6Iue9l+aXuiIsImNvbV9yaXNrX3RlbCI6IjE1OTUyMzc5Nzg0IiwiY29tX3Jpc2tfZW1haWwiOiJ6aG91Y0B0dG91Y2guY29tLmNuIiwiY29tX29wZXJhdGVfbmFtZSI6Iue9l+aXuiIsImNvbV9vcGVyYXRlX3RlbCI6IjE1OTUyMzc5Nzg0IiwiY29tX29wZXJhdGVfZW1haWwiOiJ6aG91Y0B0dG91Y2guY29tLmNuIn0=`)

	fmt.Println(string(b))

}
