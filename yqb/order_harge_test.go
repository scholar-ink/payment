package yqb

import (
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"testing"
	"time"
)

func TestCharge_Handle(t *testing.T) {
	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		MerchantId: "900002076110",
		Key:        "MIICXAIBAAKBgQC0gCKnRQwBVi5vB5D22Hzj5mzPeqkyn0fdQerR0exVYsXonGW8NdQliMPDuenSkNMKLWU61gu+cr16Z9nU3ZEIGafXMrV2ejC534ir4/tN9sJwFN2TBEnPzp64X+fPZ1gukbyMturSENMlEeRy66RDfgOKlQ1andNE0GXPGEaB7QIDAQABAoGAPeLo/rzXvuWTihHoDBpxx5nLmN+Pf4HJG/1ry4DULPpurLBi2UMDd9MKe+MUXlO1Z/nkiWW2JRgmCv/sWeWVwcOblYZQwpo1phRlWVaYjom0ErT6jxXp+KoSzDLD/KWbSTmoCI2hnzHPVn9+n/3ydyxS189xO8Iy7FJg7JZw2SkCQQDi3dp7uNBcNhQFcBB3fuRm0pA1jcNB3ZVGwR1lfnxFJy7IiBPh8NZ0+qDh3zzJpeZ2infAS5pGFAtpH2vaxVofAkEAy64EE2dG0LWI6c1ZI95Pg9Sp1bos8tSyELOKGUjFNGAWlX1FS0Gcx7m0zaQDMx+YZn14P38/kCg9Gko0Zkw6cwJAVh2aqq4bwXTxoqLITWPIdVNmVD3MqXrgCCyT6QvbRjC6bHpi+bf+4Aam/TZND8o+si6biE3qcMF5F11LBo81RwJBAIYAZcQSRV/Yn0wbAtN7VlsIVfg8NYhJcSFVzbyWyM+NjC4rraQ2dZxdTFui1UwUNGLxl2RHJp2EK7cZLOJEsosCQFhPsUobWqDiyNJnTgTACNmpiVqQVeDQrx4VW1sIeRdCcat0u/Zl6r51jrFwXn4ymVRQQp6o5TBbBMb+lhInwkI=",
	})

	ret, err := charge.Handle(&OrderChargeConf{
		PayScene:           "01",
		MerTradeNo:         "200212115518020696664518",
		TotalAmount:        "1",
		Subject:            "测试商品",
		OrderDesc:          "测试商品描述2",
		TradeType:          "01",
		PlatformMerchantId: "900002074540",
		BusinessExtendArea: &BusinessExtendArea{
			TerminalId:   "http://www.baidu.com",
			TerminalType: "00",
		},
		OrderExpireTime: time.Now().Add(30 * time.Minute).Format("2006-01-02 15:04:05"),
		NotifyUrl:       "https://web.51shouqianla.com/v1/payment/yqb/notify",
		SubAppId:        "wxdae782e8546f5bb6",
		SubOpenId:       "oNEbm1CDWvqk7nYy7bhhc4hm-8Y8",
	})

	fmt.Println(err)

	fmt.Println(ret)
}

func TestOrderCharge_WxApply(t *testing.T) {
	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		MerchantId: "900002076077",
		Key:        "MIICXAIBAAKBgQC0gCKnRQwBVi5vB5D22Hzj5mzPeqkyn0fdQerR0exVYsXonGW8NdQliMPDuenSkNMKLWU61gu+cr16Z9nU3ZEIGafXMrV2ejC534ir4/tN9sJwFN2TBEnPzp64X+fPZ1gukbyMturSENMlEeRy66RDfgOKlQ1andNE0GXPGEaB7QIDAQABAoGAPeLo/rzXvuWTihHoDBpxx5nLmN+Pf4HJG/1ry4DULPpurLBi2UMDd9MKe+MUXlO1Z/nkiWW2JRgmCv/sWeWVwcOblYZQwpo1phRlWVaYjom0ErT6jxXp+KoSzDLD/KWbSTmoCI2hnzHPVn9+n/3ydyxS189xO8Iy7FJg7JZw2SkCQQDi3dp7uNBcNhQFcBB3fuRm0pA1jcNB3ZVGwR1lfnxFJy7IiBPh8NZ0+qDh3zzJpeZ2infAS5pGFAtpH2vaxVofAkEAy64EE2dG0LWI6c1ZI95Pg9Sp1bos8tSyELOKGUjFNGAWlX1FS0Gcx7m0zaQDMx+YZn14P38/kCg9Gko0Zkw6cwJAVh2aqq4bwXTxoqLITWPIdVNmVD3MqXrgCCyT6QvbRjC6bHpi+bf+4Aam/TZND8o+si6biE3qcMF5F11LBo81RwJBAIYAZcQSRV/Yn0wbAtN7VlsIVfg8NYhJcSFVzbyWyM+NjC4rraQ2dZxdTFui1UwUNGLxl2RHJp2EK7cZLOJEsosCQFhPsUobWqDiyNJnTgTACNmpiVqQVeDQrx4VW1sIeRdCcat0u/Zl6r51jrFwXn4ymVRQQp6o5TBbBMb+lhInwkI=",
	})

	ret, err := charge.WxApply(&WxApplyConf{
		SubAppId:           "wxdae782e8546f5bb6",
		PlatformMerchantId: "900002074540",
	})

	fmt.Println(err)

	fmt.Println(ret)
}

func TestOrderCharge_Query(t *testing.T) {
	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		MerchantId: "900002076118",
		Key:        "MIICXAIBAAKBgQC0gCKnRQwBVi5vB5D22Hzj5mzPeqkyn0fdQerR0exVYsXonGW8NdQliMPDuenSkNMKLWU61gu+cr16Z9nU3ZEIGafXMrV2ejC534ir4/tN9sJwFN2TBEnPzp64X+fPZ1gukbyMturSENMlEeRy66RDfgOKlQ1andNE0GXPGEaB7QIDAQABAoGAPeLo/rzXvuWTihHoDBpxx5nLmN+Pf4HJG/1ry4DULPpurLBi2UMDd9MKe+MUXlO1Z/nkiWW2JRgmCv/sWeWVwcOblYZQwpo1phRlWVaYjom0ErT6jxXp+KoSzDLD/KWbSTmoCI2hnzHPVn9+n/3ydyxS189xO8Iy7FJg7JZw2SkCQQDi3dp7uNBcNhQFcBB3fuRm0pA1jcNB3ZVGwR1lfnxFJy7IiBPh8NZ0+qDh3zzJpeZ2infAS5pGFAtpH2vaxVofAkEAy64EE2dG0LWI6c1ZI95Pg9Sp1bos8tSyELOKGUjFNGAWlX1FS0Gcx7m0zaQDMx+YZn14P38/kCg9Gko0Zkw6cwJAVh2aqq4bwXTxoqLITWPIdVNmVD3MqXrgCCyT6QvbRjC6bHpi+bf+4Aam/TZND8o+si6biE3qcMF5F11LBo81RwJBAIYAZcQSRV/Yn0wbAtN7VlsIVfg8NYhJcSFVzbyWyM+NjC4rraQ2dZxdTFui1UwUNGLxl2RHJp2EK7cZLOJEsosCQFhPsUobWqDiyNJnTgTACNmpiVqQVeDQrx4VW1sIeRdCcat0u/Zl6r51jrFwXn4ymVRQQp6o5TBbBMb+lhInwkI=",
	})

	ret, err := charge.Query(&QueryConf{
		MerTradeNo: "200215014630472938017827",
		//TradeOrderNo:       "200215014630472938017827",
		PlatformMerchantId: "900002074540",
	})

	fmt.Println(err)

	fmt.Println(ret)
}

func TestOrderCharge_Refund(t *testing.T) {
	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		MerchantId: "900002076077",
		Key:        "MIICXAIBAAKBgQC0gCKnRQwBVi5vB5D22Hzj5mzPeqkyn0fdQerR0exVYsXonGW8NdQliMPDuenSkNMKLWU61gu+cr16Z9nU3ZEIGafXMrV2ejC534ir4/tN9sJwFN2TBEnPzp64X+fPZ1gukbyMturSENMlEeRy66RDfgOKlQ1andNE0GXPGEaB7QIDAQABAoGAPeLo/rzXvuWTihHoDBpxx5nLmN+Pf4HJG/1ry4DULPpurLBi2UMDd9MKe+MUXlO1Z/nkiWW2JRgmCv/sWeWVwcOblYZQwpo1phRlWVaYjom0ErT6jxXp+KoSzDLD/KWbSTmoCI2hnzHPVn9+n/3ydyxS189xO8Iy7FJg7JZw2SkCQQDi3dp7uNBcNhQFcBB3fuRm0pA1jcNB3ZVGwR1lfnxFJy7IiBPh8NZ0+qDh3zzJpeZ2infAS5pGFAtpH2vaxVofAkEAy64EE2dG0LWI6c1ZI95Pg9Sp1bos8tSyELOKGUjFNGAWlX1FS0Gcx7m0zaQDMx+YZn14P38/kCg9Gko0Zkw6cwJAVh2aqq4bwXTxoqLITWPIdVNmVD3MqXrgCCyT6QvbRjC6bHpi+bf+4Aam/TZND8o+si6biE3qcMF5F11LBo81RwJBAIYAZcQSRV/Yn0wbAtN7VlsIVfg8NYhJcSFVzbyWyM+NjC4rraQ2dZxdTFui1UwUNGLxl2RHJp2EK7cZLOJEsosCQFhPsUobWqDiyNJnTgTACNmpiVqQVeDQrx4VW1sIeRdCcat0u/Zl6r51jrFwXn4ymVRQQp6o5TBbBMb+lhInwkI=",
	})

	ret, err := charge.Refund(&RefundConf{
		OrigMerTradeNo:     "200215014630472938017827",
		PlatformMerchantId: "900002074540",
		MerRefundTradeNo:   helper.CreateSn(),
		RefundAmount:       "1800",
		RefundTransTime:    time.Now().Format("2006-01-02 15:04:05"),
	})

	fmt.Println(err)

	fmt.Println(ret)
}

//
//func TestOrderCharge_GetAuthCodeUrl(t *testing.T) {
//	charge := new(OrderCharge)
//
//	charge.InitBaseConfig(&BaseConfig{
//		AppCode: "9A02B598D8F64BD789DBE8C6D4805D07",
//		Key:     "F3954E8B34F5474AA38BB127D5BA94FE",
//	})
//
//	ret := charge.GetAuthCodeUrl(&GetAuthCodeUrlConf{
//		MchId:"vrR0giJVgG",
//		RedirectUri:"https://openapi-test.qfpay.com/tools/get_wx_code",
//	})
//
//	fmt.Println(ret)
//}
//
//func TestOrderCharge_GetOpenId(t *testing.T) {
//	charge := new(OrderCharge)
//
//	charge.InitBaseConfig(&BaseConfig{
//		AppCode: "9A02B598D8F64BD789DBE8C6D4805D07",
//		Key:     "F3954E8B34F5474AA38BB127D5BA94FE",
//	})
//
//	ret,err := charge.GetOpenId(&GetOpenIdConf{
//		MchId:"vrR0giJVgG",
//		Code:"061v6ezu1o3JDh0j7kzu101Zyu1v6ezY",
//	})
//
//	fmt.Println(ret)
//	fmt.Println(err)
//}
