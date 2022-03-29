package enter

import (
	"encoding/base64"
	"fmt"
	"github.com/scholar-ink/payment/util/http"
	"testing"
)

func TestHjSharingEnter_Create(t *testing.T) {
	HjSharing := new(HjSharingEnter)

	HjSharing.InitConfig(&HjSharingConfig{
		MchNo: "888108800008622",
		Key:   "a5095968f9ec42f3889add3297c4e77b",
	})

	ret, err := HjSharing.Create(&HjSharingCreateConf{
		LoginName:           "1358582108006212@udian.me",
		AltMchName:          "重庆韬启信息科技有限公司",
		AltMchShortName:     "韬启科技",
		AltMerchantType:     "11",
		BusiContactName:     "周超",
		BusiContactMobileNo: "13585821080",
		PhoneNo:             "13585821080",
		ManageScope:         "计算机",
		ManageAddr:          "重庆市江北区渝北三村30号24-11(仅限用于行政办公通讯联旧G",
		LegalPerson:         "周超",
		IdCardNo:            "500234199208102174",
		IdCardExpiry:        "2024-12-11",
		LicenseNo:           "91310109324309418B",
		LicenseExpiry:       "2034-12-14",
		SettMode:            "1",
		SettDateType:        "2",
		RiskDay:             "1",
		BankAccountType:     "1",
		BankAccountName:     "周超",
		BankAccountNo:       "6217001180010001264",
		NotifyUrl:           "http://web.udian.me/v1/common/enter-notify",
	})
	fmt.Println(err)
	fmt.Printf("%+v", ret)
}

func TestHjSharingEnter_Signing(t *testing.T) {
	HjSharing := new(HjSharingEnter)

	HjSharing.InitConfig(&HjSharingConfig{
		MchNo: "888108800008622",
		Key:   "a5095968f9ec42f3889add3297c4e77b",
	})

	ret, err := HjSharing.Signing(&HjSharingSignConf{
		AltMchNo:   "333108900001189",
		SignStatus: "P1000",
		SignTime:   "2019-10-26 19:57:20",
	})

	fmt.Println(err)
	fmt.Printf("%+v", ret)
}

func TestHjSharingEnter_Upload(t *testing.T) {

	b, err := http.GetBytes("https://cdn.51shouqianla.com/FufT1g0H_idiWd9HtamvigYgXX_0")

	if err != nil {
		return
	}

	cardPositive := base64.StdEncoding.EncodeToString(b)

	b, err = http.GetBytes("https://cdn.51shouqianla.com/FhKtNHjpBS7VNwi8V0DK4YkFEtZ9")

	if err != nil {
		return
	}

	cardNegative := base64.StdEncoding.EncodeToString(b)

	b, err = http.GetBytes("https://cdn.51shouqianla.com/Fj5phMwlOoBaP7-6mbto7u8bUvcR")

	if err != nil {
		return
	}

	tradeLicence := base64.StdEncoding.EncodeToString(b)

	HjSharing := new(HjSharingEnter)

	HjSharing.InitConfig(&HjSharingConfig{
		MchNo: "888108800004725",
		Key:   "71a58215acff43dfaba0b64bad06bd75",
	})

	ret, err := HjSharing.Upload(&HjSharingUploadConf{
		AltMchNo:     "333109000005525",
		CardPositive: "data:image/png;base64," + cardPositive,
		CardNegative: "data:image/png;base64," + cardNegative,
		TradeLicence: "data:image/png;base64," + tradeLicence,
	})
	fmt.Println(err)
	fmt.Printf("%+v", ret)
}
