package gpay

import (
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"testing"
)

func TestCharge_Handle(t *testing.T) {
	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		//MerchantNo: "594107", //呦兔
		//Md5Key:     "e2527f1abfe64003a6842e0aaddbd680",
		//MerchantNo:"100869",//韬启下商户
		//Md5Key:"9bde62e2b7004dc8b821518bb39ef27f",
		MerchantNo: "736971", //韬启
		Md5Key:     "63d623cbb482405f979085d2f6100ae4",
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
	ret, err := charge.Handle(&OrderChargeConf{
		PaymentType: "pay-zfb-native",
		OutTradeNo:  helper.CreateSn(),
		TotalFee:    int64(1),
		FeeType:     "CNY",
		MchCreateIp: "127.0.0.1",
		NotifyUrl:   "http://tq.udian.me/v1/common/enter-notify",
		Body:        "测试商品",
		Attach:      "测试商品附加信息",
		OpUserId:    "orxYXv1vuNJWJYKhZKpf8Wc2KxRY",
		AuthCode:    "134545271646929162",
		SceneInfo: &SceneInfo{
			SceneType:    "H5",
			SceneBizType: "WAP",
		},
		//Ext: &Ext{
		//	SharingParams: []*SharingParam{
		//		{
		//			FeeType:        "2",
		//			FeeValue:       divideFeeStr,
		//			AccountType:    "2",
		//			SharingAccount: "b1552470439457613",
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
