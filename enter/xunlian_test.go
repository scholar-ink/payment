package enter

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/scholar-ink/payment/util/http"
	"io/ioutil"
	"testing"
)

func TestXunLianEnter_Create(t *testing.T) {
	xh := new(XunLianEnter)

	xh.InitConfig(&XunLianConfig{
		InsCode:   "19940909",
		AgentCode: "19940909",
		Key:       "e143ef60699a4fb0b964d1aa",
	})

	ret, err := xh.Create(&XunLianCreateConf{
		ReqId:                "1f83e21de2c043bc84b36b48a8e76d95",
		WxpChanFlag:          "1",
		AlpChanFlag:          "1",
		AlpLevel:             "01",
		MerType:              "3",
		BigCategoryName:      "01",
		SubCategoryCode:      "5311",
		ProvinceName:         "陕西省",
		CityName:             "西安市",
		DistrictName:         "新城区",
		ThreeIntoOne:         "1",
		IsInstitution:        "0",
		BizAddress:           "西安市新城区解放路318号3层",
		MerName:              "西安市新城区威汀曼商贸店",
		ShortName:            "西安市新城区威汀曼商贸店",
		LicenseCode:          "92610102MA6WHRJUXK",
		LicenseExpired:       "2999-12-31",
		CardType:             "1",
		CardHolderType:       "1",
		LegalName:            "李红艳",
		LegalCard:            "320830199012082427",
		LegalExpired:         "2025-11-17",
		DelayType:            "0",
		BranchProvince:       "江苏省",
		BranchCityName:       "淮安市",
		BankName:             "上海浦东发展银行股份有限公司淮安分行",
		AccountName:          "李红艳",
		AccountCode:          "6217920471658935",
		ContactName:          "李红艳",
		ContactMobile:        "17714500631",
		ContactFixed:         "17714500631",
		ContactEmail:         "zhouc@ttouch.com.cn",
		WxpMerCode:           "0002900F0647976",
		BusinessCategoryCode: "309",
		Wxp:                  "0.25",
		AlpFirstMerCode:      "0002900F0370542",
		AlpCategoryCodeV3:    "2015062600002758",
		Alp:                  "0.25",
	})
	fmt.Println(err)
	fmt.Printf("%+v", ret)
}

func TestXunLianUpload_Handle(t *testing.T) {
	xh := new(XunLianEnter)

	xh.InitConfig(&XunLianConfig{
		InsCode:   "19940909",
		AgentCode: "19940909",
		GroupCode: "1213",
		Key:       "e143ef60699a4fb0b964d1aa",
	})

	pfxData, _ := ioutil.ReadFile("test_19940909_id_rsa")

	fileList := make(map[string][]byte, 0)

	//经营场所照片(门面)
	b, err := http.GetBytes("https://cdn.51shouqianla.com/FvbKfaCd29ejnNjzhSTqL4-4iNFt")
	if err != nil {
		return
	}
	fileList["portalPhoto"] = b

	//经营场所照片(内景)
	b, err = http.GetBytes("https://cdn.51shouqianla.com/FvbKfaCd29ejnNjzhSTqL4-4iNFt")
	if err != nil {
		return
	}
	fileList["scenePhoto"] = b

	//营业执照照片
	b, err = http.GetBytes("https://cdn.51shouqianla.com/FnKBPUNAly5lGfEYKlcr1XqyNi18")
	if err != nil {
		return
	}
	fileList["bizLicense"] = b

	//法人身份证件(正面)
	b, err = http.GetBytes("https://cdn.51shouqianla.com/FnEJpvm9_kxVEJCT6uMNqxoknHp9")
	if err != nil {
		return
	}
	fileList["legalIdCardZ"] = b

	//法人身份证件(反面)
	b, err = http.GetBytes("https://cdn.51shouqianla.com/FlU6HB8-FPvgIzztR2cupUslcVQ_")
	if err != nil {
		return
	}
	fileList["legalIdCardF"] = b

	//银行卡照片
	b, err = http.GetBytes("https://cdn.51shouqianla.com/FhUuP3bsZLhHzBZgUEnm9bk9UkGr")
	if err != nil {
		return
	}
	fileList["inAccountBankCard"] = b

	ret, err := xh.Upload(&XunLianUploadConf{
		User:     "19940909",
		Pem:      pfxData,
		FileList: fileList,
	})

	fmt.Printf("%+v\n", ret)
	fmt.Println(err)
}

func TestXunLianEnter_Zip(t *testing.T) {
	xh := new(XunLianEnter)

	xh.InitConfig(&XunLianConfig{
		InsCode:   "19940909",
		AgentCode: "19940909",
		GroupCode: "1213",
		Key:       "e143ef60699a4fb0b964d1aa",
	})

	//// 创建 zip 包文件
	//fw, err := os.Create("1.zip")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//defer fw.Close()

	//fw := bytes.NewBuffer([]byte{})

	// 实例化新的 zip.Writer
	buf := new(bytes.Buffer)

	zw := zip.NewWriter(buf)

	b, err := http.GetBytes("https://cdn.51shouqianla.com/FhUuP3bsZLhHzBZgUEnm9bk9UkGr")

	f, err := zw.Create("1.png")

	if err != nil {
		t.Fatal(err)
	}
	_, err = f.Write(b)

	if err != nil {
		t.Fatal(err)
	}

	b, err = http.GetBytes("https://cdn.51shouqianla.com/FlU6HB8-FPvgIzztR2cupUslcVQ_")

	f, err = zw.Create("2.png")

	if err != nil {
		t.Fatal(err)
	}
	_, err = f.Write(b)

	if err != nil {
		t.Fatal(err)
	}

	err = zw.Close()
	if err != nil {
		t.Fatal(err)
	}

	//f, err = os.OpenFile("file.zip", os.O_CREATE|os.O_WRONLY, 0666)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//buf.WriteTo(f)

	ioutil.WriteFile("1.zip", buf.Bytes(), 0666)
}
