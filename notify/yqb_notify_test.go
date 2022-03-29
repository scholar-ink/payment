package notify

import (
	"fmt"
	"testing"
)

func TestYqbNotify_Handle(t *testing.T) {

	notify := new(YqbNotify)

	ret := `{"body":{"businessCode":"00","businessMsg":"交易成功","channelOrderNo":"2020021399000011584939860200006","merTradeNo":"200213163818258983073436","merchantId":"900002076077","payTime":"2020-02-13 16:39:51","platformMerchantId":"900002074540","totalAmount":"1","tradeOrderNo":"2002132720375745","tradeType":"02"},"head":{"resultCode":"000000","resultMsg":"成功","signType":"RSA"},"sign":{"signContent":"357AC3009C2CB44EB59DCAC6E494FB5072480A378E98DA57DB713F18D789116B010D1BCB60322A947FA9FB71C1E4323CD9AB823AE016B70A409F6E85515A2678C2F974EFC56938E652AF288211AFCEE1C6161A612C6722C2A437BF5AB87D43ACC775574B6FFBED6E3D11CEFCFE0886690C71AF17F9A03848842D409D8635171A"}}`

	result := notify.Handle(ret, func(merchantNo string) (md5Key, privateKey string, err error) {
		return "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDYD+Zr9F9Y2O08E8QSSVhYCbqZT2rwTGsAp6HyoYVArPPg/WldlNDq98YE52MNJICQ+2xOvWo1qE54GlHs63Lkcf3rYVhqSj80gV6TQPEl5NQ/ZLjkQWQOBuAcW2DXmUB51BrVaH9CNy4oOn1k8KNOfeSvv4Ke9Rx/0e59NrUg4QIDAQAB", "MIICXAIBAAKBgQC0gCKnRQwBVi5vB5D22Hzj5mzPeqkyn0fdQerR0exVYsXonGW8NdQliMPDuenSkNMKLWU61gu+cr16Z9nU3ZEIGafXMrV2ejC534ir4/tN9sJwFN2TBEnPzp64X+fPZ1gukbyMturSENMlEeRy66RDfgOKlQ1andNE0GXPGEaB7QIDAQABAoGAPeLo/rzXvuWTihHoDBpxx5nLmN+Pf4HJG/1ry4DULPpurLBi2UMDd9MKe+MUXlO1Z/nkiWW2JRgmCv/sWeWVwcOblYZQwpo1phRlWVaYjom0ErT6jxXp+KoSzDLD/KWbSTmoCI2hnzHPVn9+n/3ydyxS189xO8Iy7FJg7JZw2SkCQQDi3dp7uNBcNhQFcBB3fuRm0pA1jcNB3ZVGwR1lfnxFJy7IiBPh8NZ0+qDh3zzJpeZ2infAS5pGFAtpH2vaxVofAkEAy64EE2dG0LWI6c1ZI95Pg9Sp1bos8tSyELOKGUjFNGAWlX1FS0Gcx7m0zaQDMx+YZn14P38/kCg9Gko0Zkw6cwJAVh2aqq4bwXTxoqLITWPIdVNmVD3MqXrgCCyT6QvbRjC6bHpi+bf+4Aam/TZND8o+si6biE3qcMF5F11LBo81RwJBAIYAZcQSRV/Yn0wbAtN7VlsIVfg8NYhJcSFVzbyWyM+NjC4rraQ2dZxdTFui1UwUNGLxl2RHJp2EK7cZLOJEsosCQFhPsUobWqDiyNJnTgTACNmpiVqQVeDQrx4VW1sIeRdCcat0u/Zl6r51jrFwXn4ymVRQQp6o5TBbBMb+lhInwkI=", nil
	}, func(data *YqbNotifyData) error {
		fmt.Printf("%+v\n", data)
		return nil
	})

	fmt.Println(result)
}