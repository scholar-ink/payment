package sumpay

import (
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"io/ioutil"
	"testing"
	"time"
)

func TestCharge_Handle(t *testing.T) {
	charge := new(OrderCharge)

	pfxData, _ := ioutil.ReadFile("shouqianla.pfx")

	charge.InitBaseConfig(&BaseConfig{
		AppId:        "101803663",
		PfxData:      pfxData,
		CertPassWord: "shouqianla",
		PublicKey:    "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA5iO3Rg9G10gJp9uB9svUDo8TVcMfvd/rCa9mrw4Ai1veh8hEb8Xk5LYQWd9g1DgpSBFjHhe/iO7h40ZWC/STgrmitH4R9K5vEFYBo6mJpVcKyvCimbtp5dcwdLojw7BwRr7Y/7STPSDJ270mVqjNqXPAenTigQ8ay4LChcPtFZUnSWb2DL92d/o6+26gS/eqed3F28TR4WtFdxPxGZy6oTtZ2B8jI9scCFfzFlsF/tqeExCmLnJbGUNaLUDPMQIe02/W+mZOiC2aeAfOkC72Z3leczEm7Wvp4fXKsRV6XmCs6/DrEMY6RjHVWocXCtYxAMY3HOAiL+7voSLZWn1ArwIDAQAB",
	})

	ret, err := charge.Handle(&OrderChargeConf{
		MerNo: "101813675",
		//SubMerNo:     "101813675",
		UserId:       helper.CreateSn(),
		BusinessCode: "09",
		GoodsName:    "测试商品",
		GoodsNum:     "1",
		GoodsType:    "1",
		OrderNo:      helper.CreateSn(),
		OrderTime:    time.Now().Format("20060102150405"),
		OrderAmount:  "0.01",
		NeedNotify:   "1",
		NotifyUrl:    "http://tq.udian.me/v1/common/enter-notify",
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v", ret.PayQrCodePay)
	}

}
