package enter

import (
	"encoding/base64"
	"fmt"
	"github.com/scholar-ink/payment/util/http"
	"testing"
)

func TestHjEnter_Create(t *testing.T) {
	Hj := new(HjEnter)

	Hj.InitConfig(&HjConfig{
		MchNo: "888108800008622",
		Key:   "a5095968f9ec42f3889add3297c4e77b",
	})

	ret, err := Hj.Create(&HjCreateConf{
		FilingMerchantId:        "3",
		FullName:                "西安市新城区红运莱商贸部",
		ShortName:               "红运莱生活馆",
		ServicePhone:            "17714500631",
		Category:                "JP2001",
		Address:                 "解放路318号3层",
		Province:                "1200",
		City:                    "7910",
		County:                  "7911",
		MerchantType:            "2",
		LinenceType:             "13",
		LinenceNo:               "92610102MA6WHR2G79",
		CompanyAddress:          "解放路318号3层",
		LicenseValidDate:        "[\"2015-11-17\",\"forever\"]",
		LicensePic:              "OT20191108210907741jn",
		ShopEntrancePic:         "",
		IndoorPic:               "",
		LegalPersonCemType:      "01",
		LegalPersonName:         "李红艳",
		LegalPersonCemNum:       "320830199012082427",
		LegalPersonValidDate:    "[\"2015-11-17\",\"2025-11-17\"]",
		IdNomalPic:              "OT20191108210947992Wc",
		IdReversePic:            "OT20191108211017502gW",
		ContactName:             "李红艳",
		ContactMobile:           "17714500631",
		ContactEmail:            "zhouc@ttouch.com.cn",
		ContactCertfiticateNum:  "320830199012082427",
		ContactCertfiticateType: "01",
		CardNum:                 "6217920471658935",
		UserName:                "李红艳",
		CardNumType:             "0",
		ProductCode:             "wxpay,alipay",
		UseScenes:               "2",
	})
	fmt.Println(err)
	fmt.Printf("%+v", ret)
}

func TestHjEnter_WxApply(t *testing.T) {
	Hj := new(HjEnter)

	Hj.InitConfig(&HjConfig{
		MchNo: "888108800008622",
		Key:   "a5095968f9ec42f3889add3297c4e77b",
	})

	ret, err := Hj.WxApply(&HjWxApplyConf{
		TradeMerchantNo: "777190200042035",
		SubAppid:        "wx29f310b024786339",
	})

	fmt.Println(err)
	fmt.Printf("%+v", ret)
}
func TestHjEnter_WxApply2(t *testing.T) {
	Hj := new(HjEnter)

	Hj.InitConfig(&HjConfig{
		MchNo: "888108800004725",
		Key:   "71a58215acff43dfaba0b64bad06bd75",
	})

	ret, err := Hj.WxApply(&HjWxApplyConf{
		TradeMerchantNo: "777165800051048",
		SubAppid:        "wxdae782e8546f5bb6",
	})

	fmt.Println(err)
	fmt.Printf("%+v", ret)
}

func TestHjEnter_Query(t *testing.T) {
	Hj := new(HjEnter)

	Hj.InitConfig(&HjConfig{
		MchNo: "888108800004725",
		Key:   "71a58215acff43dfaba0b64bad06bd75",
	})

	ret, err := Hj.Query(&HjQueryConf{
		TradeMerchantNo: "777187400086997",
	})

	for _, product := range ret.ProductList {
		fmt.Println(product.ProductStatus)
		fmt.Println(product.ProductCode)
		fmt.Println(product.ProductBizMsg)
	}

	fmt.Println(err)
	fmt.Printf("%+v", ret)
}

func TestHjEnter_Upload(t *testing.T) {

	b, err := http.GetBytes("https://cdn.51shouqianla.com/Fp1X23m6Z42KeZr9ZR7Lw_0ADPLm")

	if err != nil {
		return
	}

	cardPositive := base64.StdEncoding.EncodeToString(b)

	Hj := new(HjEnter)

	Hj.InitConfig(&HjConfig{
		MchNo: "888108800004725",
		Key:   "71a58215acff43dfaba0b64bad06bd75",
	})

	ret, err := Hj.Upload(&HjUploadConf{
		ImgFile: cardPositive,
	})
	fmt.Println(err)
	fmt.Printf("%+v", ret)
}
