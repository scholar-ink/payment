package enter

import (
	"fmt"
	"github.com/scholar-ink/payment/util/http"
	"testing"
)

func TestXinHuiEnter_Create(t *testing.T) {
	xh := new(XinHuiEnter)

	xh.InitConfig(&XinHuiConfig{
		AgentMerNo: "8000100022321",
		Key:        "MIICXAIBAAKBgQCvvmDImcUOnKz5mYRxLaN7QNT3REC4WBi//3Ni+swXj7q5JAU3saLzq6/qcCM3aTrsfIi48fNs5xsaBCW6/1GgMgb97EAUTg/xU3k/6xR76DMP61DxEZchuETntibQ+0Ef6/LtQ4btCGVXmTkAl31iQ1ssKRoKemYmlvfIDNy3+QIDAQABAoGAZlXDYcw4tSOCje1Y89aRhang2QNDdJTIBLUpaY+E3ItzPW++IgosSxvEWg1mVFPQXfi+XIN3Lgj8/Q9BMTyPOHO7IRaD1WrRmCAerCxNFSnCHvpLPURCqnzTw0D0IQPo1wcQwC2AuMHZvFukEvkfPW/jjO3U4ZgQSSbLMfm9jykCQQDgmjyI1dDUWwDUnPvK1lvIdw7p01IG0RHS5Hgqf764hZecs1NRzynyshqUpRe0bQ9ozQgO2NcYOtebJcMOm+0/AkEAyE+mSheNJ4YmWeticNBNPOuImn2qGcKmN70kou8y1e0BlPLWV/IHMkmRlhAyKOiX+ze/LKnTiwzOAPTiQGm0xwJATeXwnNzbous1LIiN49nY13xDleGPD4Ivll9bNhI8Sa872ENx4GvjdqNDCM8Bm7g/oe+KneujHmo6ITtFnamC7QJAMWYDGk6IjvC0UISN+EhGY/mp7H+FDWlFWIWanVvj64HRXAwu8+1J/QrLjnhcBl6l7FwpFziiZK45t16s1Tm8TQJBALSf6XP/8heKdUnAjFqxl1oX5ipdpNrrYyMiXw0QrMBRC67QzM/9RQBmcSNclgwVj96Lc+ij3vbZShYiyOV3FNY=",
	})

	ret, err := xh.Create(&XinHuiCreateConf{
		MerType:          "1",
		MerchantSubName:  "洛云韵网络",
		MerchantName:     "西安市临潼区洛云韵网络工作室",
		CertType:         "2",
		CertNo:           "92610115MA6TUL4R1A",
		CertIndate:       "29990101",
		LegalName:        "徐君兰",
		LegalIdCard:      "140424199605110048",
		Contact:          "徐君兰",
		ContactPhone:     "18878839370",
		ContactEmail:     "zhouc@ttouch.com.cn",
		ProvinceNo:       "610000",
		CityNo:           "610100",
		DistrictNo:       "610115",
		Address:          "纸李街道37号",
		ServicePhone:     "18878839370",
		SettleType:       "1",
		SettleCardNo:     "6217922503487630",
		BankMerName:      "徐君兰",
		BankName:         "上海浦东发展银行",
		BankBranchName:   "上海浦东发展银行股份有限公司长治分行",
		BankCode:         "310164000019",
		BankProvince:     "140000",
		BankCity:         "140400",
		LicensePic:       "oss-cn-beijing.aliyuncs.com/xh-img-service/8000100022321/20200205/8000100022321_07ccb96d3f7d450dbf9e7bc5ae45848a.jpg",
		LegalIdCardFront: "oss-cn-beijing.aliyuncs.com/xh-img-service/8000100022321/20200205/8000100022321_ef8821a01b5147aea2ed04296e8fe11e.jpg",
		LegalIdCardBack:  "oss-cn-beijing.aliyuncs.com/xh-img-service/8000100022321/20200205/8000100022321_689b477edd9945c8bcbaccb5f5804d3c.jpg",
		ProductInfo: []*XinHuiProduct{
			{
				ChannelType: "UP_WX",
				FeeExp:      "0.0025",
			},
			{
				ChannelType: "UP_ALIPAY",
				FeeExp:      "0.0025",
			},
		},
		ChannelInfo: []*XinHuiChannel{
			{
				ChannelType: "UP_WX",
				ChannelId:   "329337467",
				Bussies:     "309",
			},
			{
				ChannelType: "UP_ALIPAY",
				ChannelId:   "2088612472076052",
				Bussies:     "5331",
			},
		},
	})
	fmt.Println(err)
	fmt.Printf("%+v", ret)
}

func TestXinHuiUpload_Handle(t *testing.T) {
	xh := new(XinHuiEnter)

	xh.InitConfig(&XinHuiConfig{
		AgentMerNo: "8000100100221",
	})

	b, err := http.GetBytes("https://cdn.51shouqianla.com/FlU6HB8-FPvgIzztR2cupUslcVQ_")

	if err != nil {
		return
	}

	ret, err := xh.Upload(&XinHuiUploadConf{
		AccessKeyId:     "LTAICLbFm1mgvyb2",
		AccessKeySecret: "P2Ckaqks1cN4dQCHsHhMJQJKa6Akey",
		BucketName:      "test-xh-img-service",
		File:            b,
	})

	fmt.Printf("%+v\n", ret)
	fmt.Println(err)
}

func TestXinHuiEnter_Query(t *testing.T) {
	xh := new(XinHuiEnter)

	xh.InitConfig(&XinHuiConfig{
		AgentMerNo: "8000100022321",
		Key:        "MIICXAIBAAKBgQCvvmDImcUOnKz5mYRxLaN7QNT3REC4WBi//3Ni+swXj7q5JAU3saLzq6/qcCM3aTrsfIi48fNs5xsaBCW6/1GgMgb97EAUTg/xU3k/6xR76DMP61DxEZchuETntibQ+0Ef6/LtQ4btCGVXmTkAl31iQ1ssKRoKemYmlvfIDNy3+QIDAQABAoGAZlXDYcw4tSOCje1Y89aRhang2QNDdJTIBLUpaY+E3ItzPW++IgosSxvEWg1mVFPQXfi+XIN3Lgj8/Q9BMTyPOHO7IRaD1WrRmCAerCxNFSnCHvpLPURCqnzTw0D0IQPo1wcQwC2AuMHZvFukEvkfPW/jjO3U4ZgQSSbLMfm9jykCQQDgmjyI1dDUWwDUnPvK1lvIdw7p01IG0RHS5Hgqf764hZecs1NRzynyshqUpRe0bQ9ozQgO2NcYOtebJcMOm+0/AkEAyE+mSheNJ4YmWeticNBNPOuImn2qGcKmN70kou8y1e0BlPLWV/IHMkmRlhAyKOiX+ze/LKnTiwzOAPTiQGm0xwJATeXwnNzbous1LIiN49nY13xDleGPD4Ivll9bNhI8Sa872ENx4GvjdqNDCM8Bm7g/oe+KneujHmo6ITtFnamC7QJAMWYDGk6IjvC0UISN+EhGY/mp7H+FDWlFWIWanVvj64HRXAwu8+1J/QrLjnhcBl6l7FwpFziiZK45t16s1Tm8TQJBALSf6XP/8heKdUnAjFqxl1oX5ipdpNrrYyMiXw0QrMBRC67QzM/9RQBmcSNclgwVj96Lc+ij3vbZShYiyOV3FNY=",
	})

	//ret,err := xh.Query(&XinHuiQueryConf{MerchantNo:"8000105202520"})
	ret, err := xh.Query(&XinHuiQueryConf{MerchantNo: "8000106346040"})

	fmt.Println(err)
	fmt.Printf("%+v", ret)
}

func TestXinHuiEnter_Update(t *testing.T) {
	xh := new(XinHuiEnter)

	xh.InitConfig(&XinHuiConfig{
		AgentMerNo: "8000100022311",
		Key:        "MIICXAIBAAKBgQCvvmDImcUOnKz5mYRxLaN7QNT3REC4WBi//3Ni+swXj7q5JAU3saLzq6/qcCM3aTrsfIi48fNs5xsaBCW6/1GgMgb97EAUTg/xU3k/6xR76DMP61DxEZchuETntibQ+0Ef6/LtQ4btCGVXmTkAl31iQ1ssKRoKemYmlvfIDNy3+QIDAQABAoGAZlXDYcw4tSOCje1Y89aRhang2QNDdJTIBLUpaY+E3ItzPW++IgosSxvEWg1mVFPQXfi+XIN3Lgj8/Q9BMTyPOHO7IRaD1WrRmCAerCxNFSnCHvpLPURCqnzTw0D0IQPo1wcQwC2AuMHZvFukEvkfPW/jjO3U4ZgQSSbLMfm9jykCQQDgmjyI1dDUWwDUnPvK1lvIdw7p01IG0RHS5Hgqf764hZecs1NRzynyshqUpRe0bQ9ozQgO2NcYOtebJcMOm+0/AkEAyE+mSheNJ4YmWeticNBNPOuImn2qGcKmN70kou8y1e0BlPLWV/IHMkmRlhAyKOiX+ze/LKnTiwzOAPTiQGm0xwJATeXwnNzbous1LIiN49nY13xDleGPD4Ivll9bNhI8Sa872ENx4GvjdqNDCM8Bm7g/oe+KneujHmo6ITtFnamC7QJAMWYDGk6IjvC0UISN+EhGY/mp7H+FDWlFWIWanVvj64HRXAwu8+1J/QrLjnhcBl6l7FwpFziiZK45t16s1Tm8TQJBALSf6XP/8heKdUnAjFqxl1oX5ipdpNrrYyMiXw0QrMBRC67QzM/9RQBmcSNclgwVj96Lc+ij3vbZShYiyOV3FNY=",
	})

	ret, err := xh.Update(&XinHuiUpdateConf{
		MerchantNo:     "8000106040390",
		SettleCardNo:   "62302900484548841",
		BankName:       "中国银行",
		BankBranchName: "中国银行武汉市湖大支行",
		BankProvince:   "420000",
		BankCity:       "420100",
		BankCode:       "104521004154",
	})
	fmt.Println(err)
	fmt.Printf("%+v", ret)
}

func TestXinHuiEnter_WxApply(t *testing.T) {

	xh := new(XinHuiEnter)

	xh.InitConfig(&XinHuiConfig{
		AgentMerNo: "8000100100221",
		Key:        "MIICXAIBAAKBgQCvvmDImcUOnKz5mYRxLaN7QNT3REC4WBi//3Ni+swXj7q5JAU3saLzq6/qcCM3aTrsfIi48fNs5xsaBCW6/1GgMgb97EAUTg/xU3k/6xR76DMP61DxEZchuETntibQ+0Ef6/LtQ4btCGVXmTkAl31iQ1ssKRoKemYmlvfIDNy3+QIDAQABAoGAZlXDYcw4tSOCje1Y89aRhang2QNDdJTIBLUpaY+E3ItzPW++IgosSxvEWg1mVFPQXfi+XIN3Lgj8/Q9BMTyPOHO7IRaD1WrRmCAerCxNFSnCHvpLPURCqnzTw0D0IQPo1wcQwC2AuMHZvFukEvkfPW/jjO3U4ZgQSSbLMfm9jykCQQDgmjyI1dDUWwDUnPvK1lvIdw7p01IG0RHS5Hgqf764hZecs1NRzynyshqUpRe0bQ9ozQgO2NcYOtebJcMOm+0/AkEAyE+mSheNJ4YmWeticNBNPOuImn2qGcKmN70kou8y1e0BlPLWV/IHMkmRlhAyKOiX+ze/LKnTiwzOAPTiQGm0xwJATeXwnNzbous1LIiN49nY13xDleGPD4Ivll9bNhI8Sa872ENx4GvjdqNDCM8Bm7g/oe+KneujHmo6ITtFnamC7QJAMWYDGk6IjvC0UISN+EhGY/mp7H+FDWlFWIWanVvj64HRXAwu8+1J/QrLjnhcBl6l7FwpFziiZK45t16s1Tm8TQJBALSf6XP/8heKdUnAjFqxl1oX5ipdpNrrYyMiXw0QrMBRC67QzM/9RQBmcSNclgwVj96Lc+ij3vbZShYiyOV3FNY=",
	})

	ret, err := xh.WxApply(&XinHuiWxApplyConf{
		MerchantNo:     "8000105768310",
		SubMchId:       "333362454",
		SubAppid:       "wxdae782e8546f5bb6",
		SubScribeAppId: "wxdae782e8546f5bb6",
		AuthPaths:      "/index",
	})

	fmt.Println(err)
	fmt.Printf("%+v", ret)
}
