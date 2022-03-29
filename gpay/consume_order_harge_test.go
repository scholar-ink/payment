package gpay

import (
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"testing"
)

func TestConsumeOrderCharge_Handle(t *testing.T) {
	charge := new(ConsumeOrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		//MerchantNo: "490401", //呦兔
		//Md5Key:     "0de473e7cbbf4482ad3cf01f0aa6333f",
		//MerchantNo: "100869",  //韬启下商户
		//Md5Key:     "9bde62e2b7004dc8b821518bb39ef27f",
		MerchantNo: "100807", //韬启
		Md5Key:     "b2ab57c9cd3d44d19d6514cfb5755a66",
	})

	//sharingParams := []*SharingParam{
	//	{
	//		FeeType:        "2",
	//		FeeValue:      divideFeeStr,
	//		AccountType:    "2",
	//		SharingAccount: "b3144456307545088",
	//	},
	//}
	//
	//if udFee > 0 {
	//	udFeeFloat := fmt.Sprintf("%.4f", float64(udFee)/100)
	//
	//	sharingParams = append(sharingParams, &SharingParam{
	//		FeeType:        "2",
	//		FeeValue:       udFeeFloat,
	//		AccountType:    "1",
	//		SharingAccount: "b1552470439457613",
	//	})
	//}
	//
	ret, err := charge.Handle(&ConsumeOrderChargeConf{
		PaymentType: "pay-wx-consume",
		OutTradeNo:  helper.CreateSn(),
		TotalFee:    int64(1),
		FeeType:     "CNY",
		MchCreateIp: "127.0.0.1",
		NotifyUrl:   "http://tq.udian.me/v1/common/enter-notify",
		Subject:     "测试商品",
		Body:        "测试商品",
		AuthCode:    "134887854856338669",
		//Ext: &Ext{
		//	SharingParams: []*SharingParam{
		//		{
		//			FeeType:        "2",
		//			FeeValue:       divideFeeStr,
		//			AccountType:    "2",
		//			SharingAccount: "b1537337771064377",
		//		},
		//		{
		//			FeeType:        "2",
		//			FeeValue:       udFeeStr,
		//			AccountType:    "1",
		//			SharingAccount: "b1552470439457613",
		//		},
		//	},
		//	SharingNotifyUrl: "http://tq.udian.me/v1/common/enter-notify",
		//},
	})

	fmt.Printf("%+v", ret)
	fmt.Println(err)
}
