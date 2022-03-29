package enter

import (
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"
	"testing"
)

func TestCmEnter_Create(t *testing.T) {
	enter := new(CmEnter)

	enter.InitConfig(&CmConfig{
		AgentId: "CS_dx_001",
		Key:     "4EFF68ACE1BC890CEAB068A4E8176503",
	})

	ret, err := enter.Create(&CmCreateConf{
		OrderId:            helper.CreateSn(),
		ShopName:           "西安市新城区威汀曼商贸店",
		ShopNickName:       "西安市新城区威汀曼商贸店",
		ShopKeeper:         "李红艳",
		KeeperPhone:        "17714500631",
		ShopPhone:          "17714500631",
		Email:              "zhouc@ttouch.com.cn",
		Province:           "陕西省",
		City:               "西安市",
		Area:               "新城区",
		ShopAddress:        "西安市新城区解放路318号3层",
		PayType:            "huifu",
		BusinessType:       "2",
		LicenceNo:          "92610102MA6WHRJUXK",
		LicenceBeginDate:   "2019-03-13",
		LicenceExpireDate:  "2099-01-01",
		ArTifName:          "李红艳",
		ArTifPhone:         "17714500631",
		ArTifIdentity:      "320830199012082427",
		RateWx:             "2.5",
		RateAliPay:         "2.5",
		Identity:           "320830199012082427",
		IdentityStartTime:  "2015-11-17",
		IdentityExpireTime: "2025-11-17",
		CardName:           "李红艳",
		Card:               "6217920471658935",
		CardPhone:          "17714500631",
		BankName:           "上海浦东发展银行",
		BankAddress:        "上海浦东发展银行股份有限公司淮安分行",
		BankAddNo:          "310308000019",
		AreaType:           "购物",
		Classify:           "综合商场",
		NotifyUrl:          "http://tq.udian.me/v1/common/enter-notify",
	})

	fmt.Println(err)
	fmt.Printf("%+v", ret)

}

func TestCmPayEnter_Upload(t *testing.T) {
	enter := new(CmEnter)

	enter.InitConfig(&CmConfig{
		AgentId: "agent_9088_425012",
		Key:     "419653CDAD57B0321C465C9270B807BB",
	})

	picData := make(map[string][]byte)

	b, _ := http.GetBytes("https://cdn.51shouqianla.com/FvbKfaCd29ejnNjzhSTqL4-4iNFt")

	picData["merchantHead"] = b
	picData["merchantCheck"] = b
	picData["otherPhoto3"] = b

	b, _ = http.GetBytes("https://cdn.51shouqianla.com/FnEJpvm9_kxVEJCT6uMNqxoknHp9")
	picData["identityFace"] = b
	b, _ = http.GetBytes("https://cdn.51shouqianla.com/FlU6HB8-FPvgIzztR2cupUslcVQ_")
	picData["identityBack"] = b
	b, _ = http.GetBytes("https://cdn.51shouqianla.com/FhUuP3bsZLhHzBZgUEnm9bk9UkGr")
	picData["bussinessCard"] = b
	b, _ = http.GetBytes("https://cdn.51shouqianla.com/FnKBPUNAly5lGfEYKlcr1XqyNi18")
	picData["bussiness"] = b
	picData["identityFaceCopy"] = picData["identityFace"]
	picData["identityBackCopy"] = picData["identityBack"]
	picData["cardFace"] = picData["bussinessCard"]

	enter.Upload(&CmUploadConf{
		ShopId:  "d50772950fefbd4969cb86466d9a2cca",
		PicData: picData,
	})
}
