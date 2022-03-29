package xinhui

import (
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"testing"
)

func TestCharge_Handle(t *testing.T) {
	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		AgentMerNo: "8000100022321",
		Key:        "MIICXAIBAAKBgQCvvmDImcUOnKz5mYRxLaN7QNT3REC4WBi//3Ni+swXj7q5JAU3saLzq6/qcCM3aTrsfIi48fNs5xsaBCW6/1GgMgb97EAUTg/xU3k/6xR76DMP61DxEZchuETntibQ+0Ef6/LtQ4btCGVXmTkAl31iQ1ssKRoKemYmlvfIDNy3+QIDAQABAoGAZlXDYcw4tSOCje1Y89aRhang2QNDdJTIBLUpaY+E3ItzPW++IgosSxvEWg1mVFPQXfi+XIN3Lgj8/Q9BMTyPOHO7IRaD1WrRmCAerCxNFSnCHvpLPURCqnzTw0D0IQPo1wcQwC2AuMHZvFukEvkfPW/jjO3U4ZgQSSbLMfm9jykCQQDgmjyI1dDUWwDUnPvK1lvIdw7p01IG0RHS5Hgqf764hZecs1NRzynyshqUpRe0bQ9ozQgO2NcYOtebJcMOm+0/AkEAyE+mSheNJ4YmWeticNBNPOuImn2qGcKmN70kou8y1e0BlPLWV/IHMkmRlhAyKOiX+ze/LKnTiwzOAPTiQGm0xwJATeXwnNzbous1LIiN49nY13xDleGPD4Ivll9bNhI8Sa872ENx4GvjdqNDCM8Bm7g/oe+KneujHmo6ITtFnamC7QJAMWYDGk6IjvC0UISN+EhGY/mp7H+FDWlFWIWanVvj64HRXAwu8+1J/QrLjnhcBl6l7FwpFziiZK45t16s1Tm8TQJBALSf6XP/8heKdUnAjFqxl1oX5ipdpNrrYyMiXw0QrMBRC67QzM/9RQBmcSNclgwVj96Lc+ij3vbZShYiyOV3FNY=",
	})

	ret, err := charge.Handle(&OrderChargeConf{
		MerchantNo:     "8000106176540",
		OpenId:         "",
		ChannelId:      "2088612472076052",
		SubMchId:       "2088300885484170",
		ChannelType:    "UP_ALIPAY",
		TrxType:        "NATIVE",
		OutTradeNo:     helper.CreateSn(),
		TotalFee:       "0.02",
		FeeType:        "0",
		DeviceInfo:     helper.CreateSn(),
		Body:           "投诉QQ:3007638620--订单号:LDED13DBA226957282",
		SpbillCreateIp: "127.0.0.1",
		NotifyUrl:      "http://tq.udian.me/v1/common/enter-notify",
	})

	fmt.Printf("%+v", ret)
	fmt.Println(err)
}
