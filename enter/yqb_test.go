package enter

import (
	"fmt"
	"github.com/scholar-ink/payment/util/http"
	"testing"
)

func TestYqbEnter_Create(t *testing.T) {
	enter := new(YqbEnter)

	enter.InitConfig(&YqbConfig{
		AesKey:    "s9DZBvxB2omgBeo0R6rVwg==",
		HashKey:   "deb5e1b5",
		ChannelNo: "CN00BX0002164341442020012117100809600",
		SystemId:  "BZXYYC",
	})

	ret, err := enter.Create(&YqbCreateConf{
		OutMerchantNo:           "02",
		MerchantType:            "11",
		BusinessLicenseNo:       "92610102MA6WHRJUXK",
		MerchantName:            "西安市新城区威汀曼商贸店",
		MerchantShortName:       "西安市新城区威汀曼商贸店",
		LegalRepresent:          "李红艳",
		BusinessLicensePath:     "/Application/Picture/20200210_203911.jpg",
		BusinessLicenseEffDate:  "2019-03-13",
		BusinessLicenseValDate:  "2999-01-01",
		FeeRate:                 "0.0035",
		EstablishDate:           "2019-03-13",
		IdentityNo:              "320830199012082427",
		IdentityPosPath:         "/Application/Picture/20200210_824000.jpg",
		IdentityNegPath:         "/Application/Picture/20200210_804053.jpg",
		IdentityEffDate:         "2015-11-17",
		IdentityValDate:         "2025-11-17",
		IndustryCode:            "L7299",
		ProvinceCode:            "610000",
		CityCode:                "610100",
		CountyCode:              "610102",
		BusinessAddress:         "西安市新城区解放路318号3层",
		StoreFacadePath:         "/Application/Picture/20200210_657264.jpg",
		StoreBussinessPlacePath: "/Application/Picture/20200210_908975.jpg",
		ElectronicSignPath:      "/Application/Picture/20200210_840097.png",
		BankAccountName:         "李红艳",
		BankAccountNo:           "6217920471658935",
		BankName:                "上海浦东发展银行",
		BranchBankName:          "上海浦东发展银行股份有限公司淮安分行",
		BankCode:                "SPDB",
		BankCardPath:            "/Application/Picture/20200210_597303.jpg",
		OpenBankPhoneNo:         "17714500631",
		AdminName:               "周超",
		AdminCellPhoneNo:        "13585821080",
		CallBackUrl:             "https://web.51shouqianla.com/v1/payment/yqb/enter-notify",
	})

	fmt.Println(err)
	fmt.Printf("%+v", ret)

}
func TestYqbEnter_Query(t *testing.T) {
	enter := new(YqbEnter)

	enter.InitConfig(&YqbConfig{
		AesKey:    "s9DZBvxB2omgBeo0R6rVwg==",
		HashKey:   "deb5e1b5",
		ChannelNo: "CN00BX0002164341442020012117100809600",
		SystemId:  "BZXYYC",
	})

	ret, err := enter.Query(&YqbQueryConf{
		OutMerchantNo: "202480",
		PageNo:        "1",
		PageSize:      "20",
	})

	fmt.Println(err)
	fmt.Printf("%+v", ret)
}

func TestYqbEnter_Update(t *testing.T) {
	enter := new(YqbEnter)

	enter.InitConfig(&YqbConfig{
		AesKey:    "s9DZBvxB2omgBeo0R6rVwg==",
		HashKey:   "deb5e1b5",
		ChannelNo: "CN00BX0002164341442020012117100809600",
		SystemId:  "BZXYYC",
	})

	ret, err := enter.Update(&YqbUpdateConf{
		OutMerchantNo:   "02",
		MerchantId:      "900002075962",
		BankAccountName: "李红艳",
		BankAccountNo:   "6217920471658935",
		BankName:        "上海浦东发展银行",
		BranchBankName:  "上海浦东发展银行股份有限公司淮安分行",
		BankCode:        "SPDB",
		BankCardPath:    "/Application/Picture/20200210_597303.jpg",
		OpenBankPhoneNo: "17714500631",
		IdentityNo:      "320830199012082427",
	})

	fmt.Println(err)
	fmt.Printf("%+v", ret)
}

func TestYqbEnter_Upload(t *testing.T) {
	enter := new(YqbEnter)

	enter.InitConfig(&YqbConfig{
		AesKey:    "s9DZBvxB2omgBeo0R6rVwg==",
		HashKey:   "deb5e1b5",
		ChannelNo: "CN00BX0002164341442020012117100809600",
		SystemId:  "BZXYYC",
	})

	b, err := http.GetBytes("https://cdn.51shouqianla.com/FnEJpvm9_kxVEJCT6uMNqxoknHp9")

	if err != nil {
		return
	}

	ret, err := enter.Upload(&YqbUploadConf{
		OutMerchantNo: "001",
		User:          "BZXYYC_074540",
		Password:      "ejow+817+FPW",
		File:          b,
	})

	fmt.Println(err)
	fmt.Println(ret)
}
