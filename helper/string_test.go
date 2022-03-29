package helper

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestRsa1Encrypt(t *testing.T) {

	pfxData, _ := ioutil.ReadFile("1111.pfx")

	b := []byte("1")

	encryptData, err := Rsa1Encrypt(pfxData, b, "584520lym.")

	fmt.Println(err)

	fmt.Println(encryptData)

}

func TestRsa2Encrypt(t *testing.T) {

	//b,err := base64.StdEncoding.DecodeString("c2fn9SVFwI0a5tE94S3rcVgdOfciRRYa5yiBmvs 0F3TjKKhbP7VdxLzHldgU7kq5FVvmgpH2zOiBoa/VGSYNBopoUo1nWI7yk3Fj1iFWgVliFdok28rW2Ha2eHK6jox6rKAVGphJi1RR//0Yb2yEkTOj0Np h6S8UrIE6UpcyE=")
	//
	//fmt.Println(err)
	//
	//fmt.Println(string(b))

	//err := Sha256WithRsaVerify([]byte(`agent_mer_no=8000100100221&buyer_pay_amount=0.10&buyer_user_id=2088612472076052&channel_type=UP_ALIPAY&fund_bill_list=[{"amount":"0.10","fund_channel":"ALIPAYACCOUNT"}]&merchant_no=8000105202610&org_external_id=2019112622001476051420837595&out_trade_no=191126210912438233544500&pay_external_id=12021199314159905411072&receipt_amount=0.10&rsp_code=0000&rsp_msg=支付成功&time_end=2019-11-26 21:09:41&total_fee=0.1&trx_external_id=12021199314159901216768&trx_status=SUCCESS&version=1.0`),"c2fn9SVFwI0a5tE94S3rcVgdOfciRRYa5yiBmvs 0F3TjKKhbP7VdxLzHldgU7kq5FVvmgpH2zOiBoa/VGSYNBopoUo1nWI7yk3Fj1iFWgVliFdok28rW2Ha2eHK6jox6rKAVGphJi1RR//0Yb2yEkTOj0Np h6S8UrIE6UpcyE=","MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCFfVpeFleLYGlalDn6Um3HXKZKBsWokEXTx7a2jdgdzWgq+Y/tUl7l/qfc0UJRwQAtPul9uRWMiG9PDGjlXNIB0g0BPB53SnwSqOLK5AhENTjbUAl2cyihlGB1yb4jY1M1LJ04EJSfx6gvgYlNdkGnkWGQETiV1gUQVE1gW7KA0QIDAQAB")

	//fmt.Println(err)

	ret, err := Sha256WithRsaSign([]byte("1"), "MIIEowIBAAKCAQEAlMqn/KEVzhkdtccidoRCEz3BqXH7AMNt/FfjOJFKQuzwLYRfFP6TmkSVlsTSkIV/Az/lJaieJgKkltU4jkjpmJSC6YS1XDewk1dTmJp0SAnk9MGXpCuUqPAo/j6rz4/0iK4FqWTUxGsjOc0tpDc+tEvDMrzNfWlYybmbMIVH2GIlNkuusol+A7j8kn3t/1yONHmVWZKt7AByYLdd+rKYqYZPOBxXAttHP/9CTLoBHbjLwFCMRamEl/lBL34vQJiAZW81AOIi4CxJg10a3A/Z6tOcJYXSivGasnTvjuyyvqvxYL5OTrHlCT/xQ068xOzU3U8fjDjIbslR375DBl//NwIDAQABAoIBADgypuo7KVIzmE4dDX44DADadXf7bfNm3PbPdynZbnQCq+B1O7hhQvykZN+SLXmaglOG4ZSssDbpDqNNm1PaZChWB3ANyLYw7odoF1HvHHZNDmYHbK/8KeT4+HK21wvJcnHhUJAfXmFlmeNuBIwetZdBelOCjhaNIJTofp3/6RfnvNqt9cOaB/I/1XfVsMv7pdL+L1woZOUYZL5tT6G8nx2rQJVnjf5uCTGCoNwwCpJHiyNpvmk6QOBMInVPU0f76Sgrjz1o5NG0y3WHzgpUyCdl1z7bTahl1BG0qsUpQ4cFR28q744GSGy8D49sjPxf5hNjIQd5KsH/6jTLLQuE7BkCgYEA4lsHJh+rTMk40s+aqbAhojKx72JGMcGUAoLT59NJC6zD0bEajs1V/KS59H3FpoI+lzwDMGPUUmYTozi9N6sPmN8YuoaR0Q3tEdo4WUr5hxkezm10HisKeeC6h0D7NJlkh3/Gp2LScRZDOUAVLjt6T7X/ysAP6ou4uri6tNFaYXMCgYEAqEcoxT7W/7Jo/+OjylOsNk8TxgCmkJtt+sjnVtkBcAFrlRck/jn3rgp7ZmoT8QKmV1ws/V2CEIaBlwpVcZZuvLpcpgIi36sduky+ahILsnOSYM8DRZzqq7WPDFAC7CxS8L8dFVBxINQxiIQrhKkwTgmAaTJGu4nzGk36saSAKi0CgYBAYwhLjeKaOvrQ7IDgF9vZWXZH07qH2LqTZEeGwBEdIw2ojioxyLLW5LyIkWYxkQbg2g9GKn9w2NxpJ3CbbytGnt9X34OG5eEznNE+hRcpmLmsmnHXSwL13Djy1EcglSmFaZFGd9PImz6QAGyF9CE8n1adg7iDTS9E3dsuKAb/hQKBgGKK2Ts4o1q1NXuz6LSQ7yYWhLPMqb3A51SW1bIr/gWDL2btWMJvW7VVehHtSKQ4MwSxe78bRRE8UyMJ8CNGPq7SS3MDiTyFzjDMxC0FSEhGGZALahUX4OyQs6Y4LJ31DtTgdb2Hj9fzqtYQ4BMdqKXqNoJj6Lvl+rCKvaXNeSg5AoGBAMBBf/0/yc0lsUSvmwmTOUsSTeDlQIwOEcr6IrNLGI4GogXc96YOrpeEUZ9KwKmHaEu7P37An5ECw/sYwI9QVc2XPkOFwnaboNPBBAg7+vgufhWzNVekE3f7rut4hX3pRW+DoUxRiEmTnZoZRGo02W6Zph0xDOMoJl4ivkaDmgzd")

	fmt.Println(base64.StdEncoding.EncodeToString(ret))
	fmt.Println(err)
	//
	//err = Sha256WithRsaVerify([]byte("1"),base64.StdEncoding.EncodeToString(ret),"MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCvvmDImcUOnKz5mYRxLaN7QNT3REC4WBi//3Ni+swXj7q5JAU3saLzq6/qcCM3aTrsfIi48fNs5xsaBCW6/1GgMgb97EAUTg/xU3k/6xR76DMP61DxEZchuETntibQ+0Ef6/LtQ4btCGVXmTkAl31iQ1ssKRoKemYmlvfIDNy3+QIDAQAB")
	//
	//fmt.Println(err)

	//ret,err = Rsa2Encrypt2([]byte("1"),"MIICXAIBAAKBgQCvvmDImcUOnKz5mYRxLaN7QNT3REC4WBi//3Ni+swXj7q5JAU3saLzq6/qcCM3aTrsfIi48fNs5xsaBCW6/1GgMgb97EAUTg/xU3k/6xR76DMP61DxEZchuETntibQ+0Ef6/LtQ4btCGVXmTkAl31iQ1ssKRoKemYmlvfIDNy3+QIDAQABAoGAZlXDYcw4tSOCje1Y89aRhang2QNDdJTIBLUpaY+E3ItzPW++IgosSxvEWg1mVFPQXfi+XIN3Lgj8/Q9BMTyPOHO7IRaD1WrRmCAerCxNFSnCHvpLPURCqnzTw0D0IQPo1wcQwC2AuMHZvFukEvkfPW/jjO3U4ZgQSSbLMfm9jykCQQDgmjyI1dDUWwDUnPvK1lvIdw7p01IG0RHS5Hgqf764hZecs1NRzynyshqUpRe0bQ9ozQgO2NcYOtebJcMOm+0/AkEAyE+mSheNJ4YmWeticNBNPOuImn2qGcKmN70kou8y1e0BlPLWV/IHMkmRlhAyKOiX+ze/LKnTiwzOAPTiQGm0xwJATeXwnNzbous1LIiN49nY13xDleGPD4Ivll9bNhI8Sa872ENx4GvjdqNDCM8Bm7g/oe+KneujHmo6ITtFnamC7QJAMWYDGk6IjvC0UISN+EhGY/mp7H+FDWlFWIWanVvj64HRXAwu8+1J/QrLjnhcBl6l7FwpFziiZK45t16s1Tm8TQJBALSf6XP/8heKdUnAjFqxl1oX5ipdpNrrYyMiXw0QrMBRC67QzM/9RQBmcSNclgwVj96Lc+ij3vbZShYiyOV3FNY=")
	////ret,err := Rsa2Encrypt([]byte("1111"),"MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCvvmDImcUOnKz5mYRxLaN7QNT3REC4WBi//3Ni+swXj7q5JAU3saLzq6/qcCM3aTrsfIi48fNs5xsaBCW6/1GgMgb97EAUTg/xU3k/6xR76DMP61DxEZchuETntibQ+0Ef6/LtQ4btCGVXmTkAl31iQ1ssKRoKemYmlvfIDNy3+QIDAQAB")
	//
	//fmt.Println(base64.StdEncoding.EncodeToString(ret))
	//fmt.Println(err)

	//retb,err := Rsa2Decrypt([]byte(ret),"MIICXAIBAAKBgQCvvmDImcUOnKz5mYRxLaN7QNT3REC4WBi//3Ni+swXj7q5JAU3saLzq6/qcCM3aTrsfIi48fNs5xsaBCW6/1GgMgb97EAUTg/xU3k/6xR76DMP61DxEZchuETntibQ+0Ef6/LtQ4btCGVXmTkAl31iQ1ssKRoKemYmlvfIDNy3+QIDAQABAoGAZlXDYcw4tSOCje1Y89aRhang2QNDdJTIBLUpaY+E3ItzPW++IgosSxvEWg1mVFPQXfi+XIN3Lgj8/Q9BMTyPOHO7IRaD1WrRmCAerCxNFSnCHvpLPURCqnzTw0D0IQPo1wcQwC2AuMHZvFukEvkfPW/jjO3U4ZgQSSbLMfm9jykCQQDgmjyI1dDUWwDUnPvK1lvIdw7p01IG0RHS5Hgqf764hZecs1NRzynyshqUpRe0bQ9ozQgO2NcYOtebJcMOm+0/AkEAyE+mSheNJ4YmWeticNBNPOuImn2qGcKmN70kou8y1e0BlPLWV/IHMkmRlhAyKOiX+ze/LKnTiwzOAPTiQGm0xwJATeXwnNzbous1LIiN49nY13xDleGPD4Ivll9bNhI8Sa872ENx4GvjdqNDCM8Bm7g/oe+KneujHmo6ITtFnamC7QJAMWYDGk6IjvC0UISN+EhGY/mp7H+FDWlFWIWanVvj64HRXAwu8+1J/QrLjnhcBl6l7FwpFziiZK45t16s1Tm8TQJBALSf6XP/8heKdUnAjFqxl1oX5ipdpNrrYyMiXw0QrMBRC67QzM/9RQBmcSNclgwVj96Lc+ij3vbZShYiyOV3FNY=")
	//
	//fmt.Println(string(retb))
	//
	//fmt.Println(err)
}

func TestSha256WithRsaSignWithPassWord(t *testing.T) {

	pfxData, _ := ioutil.ReadFile("yixuntiankong.pfx")

	b, _ := Sha256WithRsaSignWithPassWord(pfxData, []byte("1"), "sumpay")

	fmt.Println(base64.StdEncoding.EncodeToString(b))

}

func TestSha256WithRsaVerify(t *testing.T) {

	ret := `{"bizType":"03","chUserId":"UD20190320162646","completeTime":"20200107144424","customerId":"d1953fca7baf4438ae07697a79bbadae","orglOrderId":"200106190456953412957081","status":"02"}`

	sign := "Dv4+a+JQY0jSvtwcdFubU8SoPdngNUabL+mAyHU6rBu0zhavGTBt200qXnhJ6qRSjWzwQOvm93Awl7hO1rlTllCEkdKQVJd/AkrCAE2mKkoDi/1QI5uHgsSpi2a/0Pw8JfFNTB4joPvlMN7Bxr/AbUCQ2jyXDO/erbttCGBOSNQV3c2qYnjzYmmL2Q2IITR/Dg7AMo+HZgWZfXYEInJpdTJ4v//n5lJfoDt6X0evtNZCkV4+/1ewiX7XDt9ey1nTf6+45ktjD2cpIPv5sw3QOsD7pJjf9hMayVKq1iNcRQlKhRjyas0eznpZ1RT3g9nG8H3bFI8ziZ8GykuUVCH3Iw=="

	key := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAjsptGNxDkK31cOOwLEKo5cHypVlN3zp/60iRh6zaiG7jGmOn64v7/H0Qh8fCukNAv7CYDG5c7S36n1Sa5M3kuFR0Qk7g+WwXdGMcE9rNKdqGkOecaBHHX6kpKVTBmhTFmhI1aoumnfxrFKuExWm4BY2UO2Uaq9bHB8ZkbcGQ9yq9I0nYGUIkneJN4NZhUudj4fK1SHvtHUx1VO2aCbmiIZi4YaAfvvEhIUSeYtwdBtkJEMn7tlQju46Xob5zVGNUcVKUnQf7SY2xUUghU1DOJmNJ0i1heGrs6zjeOLYv8nyQ0SzWVWmaTb7LYM1QL9KdqMpY0+xTq1v2cJ3bLl3bLwIDAQAB"

	fmt.Println("signInfo=" + sign)
	fmt.Println("data=" + ret)
	fmt.Println("key=" + key)

	//keyB,_ := base64.StdEncoding.DecodeString(key)

	err := Sha1WithRsaVerify([]byte(ret), sign, key)

	fmt.Println(err)
}

func TestSha1WithRsaSign(t *testing.T) {

	sign := `https://api.adapay.tech/v1/payments{"app_id":"app_7d87c043-aae3-4357-9b2c-269349a980d6","order_no":"PY_20200224135533626284","pay_channel":"alipay","time_expire":"20200225135533","pay_amt":"0.01","goods_title":"subject","goods_desc":"body","currency":"cny","sign_type":"RSA2"}`

	//sign = "2"

	fmt.Println(sign)

	b, _ := Sha1WithRsaSignPkcs8([]byte(sign), "MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAMQhsygJ2pp4nCiDAXiqnZm6AzKSVAh+C0BgGR6QaeXzt0TdSi9VR0OQ7Qqgm92NREB3ofobXvxxT+wImrDNk6R6lnHPMTuJ/bYpm+sx397rPboRAXpV3kalQmbZ3P7oxtEWOQch0zV5B1bgQnTvxcG3REAsdaUjGs9Xvg0iDS2tAgMBAAECgYAqGFmNdF/4234Yq9V7ApOE1Qmupv1mPTdI/9ckWjaAZkilfSFY+2KqO8bEiygo6xMFCyg2t/0xDVjr/gTFgbn4KRPmYucGG+FzTRLH0nVIqnliG5Ekla6a4gwh9syHfstbOpIvJR4DfldicZ5n7MmcrdEwSmMwXrdinFbIS/P1+QJBAOr6NpFtlxVSGzr6haH5FvBWkAsF7BM0CTAUx6UNHb+RCYYQJbk8g3DLp7/vyio5uiusgCc04gehNHX4laqIdl8CQQDVrckvnYy+NLz+K/RfXEJlqayb0WblrZ1upOdoFyUhu4xqK0BswOh61xjZeS+38R8bOpnYRbLf7eoqb7vGpZ9zAkEAobhdsA99yRW+WgQrzsNxry3Ua1HDHaBVpnrWwNjbHYpDxLn+TJPCXvI7XNU7DX63i/FoLhOucNPZGExjLYBH/wJATHNZQAgGiycjV20yicvgla8XasiJIDP119h4Uu21A1Su8G15J2/9vbWn1mddg1pp3rwgvxhw312oInbHoFMxsQJBAJlyDDu6x05MeZ2nMor8gIokxq2c3+cnm4GYWZgboNgq/BknbIbOMBMoe8dJFj+ji3YNTvi1MSTDdSDqJuN/qS0=")

	//base64.StdEncoding.EncodeToString(b)

	fmt.Println(base64.StdEncoding.EncodeToString(b))

}

func TestSha256WithRsaSignWithPKCS8(t *testing.T) {
	sign := "1"

	sign = "1" + "\n"

	b, _ := Sha256WithRsaSignWithPKCS8([]byte(sign), "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCDCC2oVe6OYd8ZtuhW9AN8wV9bat5wz3rva5H8iPAv99VQkORANnh6l+a7RNVfN9w+Yii6UeavhSsulgicDUngJdCHaPIsuXRWt26ejSsLeHmxXnWPG2AObZcnyYzUzwZ4MiAWJ6RcRrF7BZGpAPkBGK0kLBeZ9e8Ko8SgRUXzVHmPjg8oF5vV0xMNDj92X0oZBVfzt0rOSqlGVWWgRkgIBz6CZKiy9pmLnKOnpG5qOOdiTdth+DsAR7ABK4lzkPeesAsR1VzP4EqW/TKC64YKhMA3N1ovfMC9EpQ2oCPwvairAsQcB/pvXxHBXttF/BTrTw/Ks9tkh2QMRBvZGHpfAgMBAAECggEABr/1GibTEyKXi4uQjGolg9eyQdNPgiAuBQdVjdzAAriRlITiPSyRKD+K8zqogy8teUk1L+PoLkJ95vhzmRZWJ+XKyC7vyr4C8DSizigXf4/FNQ3YoHaYjCW5E6OeTZgcjTSH0pxYKyi5G809o6cZLKVIxgQ/cv7oQXQOPPNUlyQ/aBl1c1cSDAWbyX7BDduqZmk5BPnyud9vtEOuKAQqFwPfy3/ZfkibilUYcvtNqRSUl/7VinZeAisSXPbKre2qk5ll/YXeavxkBZxdq6/JS5O4ivBtrQy9Fnil+7hBe6Qfw4Lt7Fv5NdObJIVwzq7cMTHGxnUaf3MNRpdkHvJsgQKBgQDrPt7vI5BuMvzM6wILXnxQ76quYPFE9nJ1glYPCpCAirKP5kAEjmH/2mJ4IqTi2uT5pgoPb0zGspL7R1tsJcSgGa98qEgyeG1n+6C9M2a+vmht8VTj0nrZeIIigQH2dF16K8c87H1jgg0N5VrjG+pRKG23dQ0rX4O0B+3MoHUN3QKBgQCOl5hCO2OVvillvvk0Wabll3ytWYZZRN/4COWtDaXY10RkpeBRyDZvAUE9Gyi/ZegfvTfZzV5gPnVFtXqbIEY8u0xD4MQSAuncY4V16cv70cvu4u3xGEZKgzgk8TOfPNxInCWUles6lP451x5B3HIAa63Ii1j3Qd0ceuI8iqT7awKBgQCh8M7Q+r+DTPBANItcvjeAE+yATFXqrmjOweFyS0h8ZH5VlyB8wnNuCKz+nIK7dApqXUXRqEHHCskp1850nW9E80md286Ph91w1oSpmkfhiPwkqxxQFOXi7RVQoVRzj1mGL7rhEr+ij7Vi2n99lgrwwY792sMtF3x3o3mtAsxxtQKBgAf5YFFr4tDP9p6zBFqyHMxAIX/MPuAlIuVLEhUQa1LqDvAV+qp4KNsiVdSl/Sxe9ZE40rPCcWGufH5ufLHKJ0NkMgqlujFLqmphwmfqsDaf7+inFilicyPdnLksJ/fivmrtGIjrrWD0ThdL+WwzeMifPPO3Hz2MmGHsWVSLaFiLAoGBANL7mp+y+J9Olx9LPjR14lanOg4PhnhIJ/CQt41WWgkEbSXfign0LaYwJQ2ly6y8KoaVPN/VeICTQ9RXsvIAwUmy0YB4hvRS6kfsdsP+9MWMooecsnsz+fUgY+Ff6pJL3dhnr0cPqiB0J0xH2gMD80i9QFUfaWAmLD7KvB1y3XA4")

	fmt.Println(base64.StdEncoding.EncodeToString(b))
}

func TestHMacSha256(t *testing.T) {
	b := HMacSha256([]byte("a"), "82GJJWlCvJ17xG8P1mz992Y36cI5882m6y6JEguLFK1M2SsxMkEX1190l701")
	fmt.Println(hex.EncodeToString(b))
}

func TestSha256(t *testing.T) {
	fmt.Println(Sha256("a"))
}
func TestMd5(t *testing.T) {
	fmt.Println(Md5(`billDate=2020-03-05&billNo=3194wktest20200305125155&instMid=QRPAYDEFAULT&mid=898340149000005&msgSrc=WWW.TEST.COM&msgType=bills.getQRCode&requestTimestamp=2020-03-05 12:51:55&tid=88880001&totalAmount=1fcAmtnx7MwismjWNhNKdHC44mNXtnEQeJkRrhKJwyrW2ysRR`))
}

func TestMd5WithRsaSignWithPKCS8(t *testing.T) {
	b, _ := Md5WithRsaSignWithPKCS8([]byte("1"), "MIICeQIBADANBgkqhkiG9w0BAQEFAASCAmMwggJfAgEAAoGBAM4/rC99irvS//RV2SMd4KHzhuqknbcpeJWMXC4FvWt+n06f7FP5ZWemx8VAmEL1/n41jOPbLTgSacuFQQmemsCp5H8d2g/XBqYy1q3f/0qPCi8Tt1yop5efcumjma2jcRdAkSyQClc8GFmEeRkEl19chrEpI/6bxtXoBD2dDLpZAgMBAAECgYEAs4btBDGND0ztCuunJFAfdhkaeShtOD/a/KG+ozjP1r/TP4cpGTdfM0gTX/mID9E8gvNt/fCMfeBZQpRtNkhefoGvlij9SSVJMyvCtqaGEbUzvaBhzLkko+d7YdZeeUg6SWAPTnPSSG3nXRwR3pMa7TqFiVpBJN8HyV8mE2fMznkCQQDsredSzLmkKv/sEZaZNngOfQxjCfBME8+Gnoz+LsQwFoXRNxKToGw1x/yt+cjWAkn+FV9Nb98oPkqPjLdKdDsnAkEA3xXVh/S1R42k1gqu0UKB1hHSQ0EUGqg9i8vWqATf0SomC4SB0K7mgmEUNDrka9KNnw1Bc9TL6w/zmPKCcibOfwJBANMjvMan7kCfP5n4gtIBvo6mTcOYnS8xSSQ+E2e6jribjxt6Nu9N4NsFoswNlnYcqqepp1Bsqba8A0YWcXlRQWcCQQDaYYpdg+yttfgF3AFUMlHdWCbH1X4ztjxBjHJ+mf7rx+HkZnuZ6I0YVqYrlvciocQnThejp01Tt5LUR5nw2xJLAkEAzoOMTiGDRQGh1BHGtbfMO0U7abma+T2XPYUUCQA8QH3B4Lr0Zm1JJ1mcPVsM3yVgBDQqiLXYIy4nKJyvn5kLWg==")

	sign := base64.StdEncoding.EncodeToString(b)

	fmt.Println(sign)

}
