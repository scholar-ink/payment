package hangpay

import (
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"

	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/scholar-ink/payment/util/map"
	"io/ioutil"
	"time"
)

const (
	ChargeUrl = "http://api.hangpay.cn/"
	SUCCESS   = "SUCCESS"
)

type BaseCharge struct {
	*BaseConfig
	Service     string `json:"-"`
	Version     string `json:"version"`
	Sign        string `json:"sign"`
	EncryKey    string `json:"encryKey"`
	RequestData string `json:"requestData"`
	RequestTime string `json:"requestTime"`
}

type BaseConfig struct {
	MerchantNo string `json:"merchantNo"`
	Key        string `json:"-"`
	PublicKey  string `json:"-"`
}

func (base *BaseCharge) InitBaseConfig(config *BaseConfig) {
	base.BaseConfig = config
	base.Version = "1.0.0"
}

func (base *BaseCharge) RetData(ret []byte) (retData []byte, err error) {
	var baseReturn BaseConfig

	err = json.Unmarshal(ret, &baseReturn)

	if err != nil {
		return
	}
	return
}

func (base *BaseCharge) makeSign(params interface{}) {

	aesKey := helper.NonceStr()[0:16]

	b, _ := json.Marshal(params)

	cipherText := helper.PKCS7Padding(b, 16)

	b, _ = helper.AesEncrypt(cipherText, []byte(aesKey)) //ECB加密

	base.RequestData = base64.StdEncoding.EncodeToString(b)

	base.RequestTime = time.Now().Format("20060102150405")

	url := fmt.Sprintf("requestData=%s&requestTime=%s&merchantNo=%s", base.RequestData, base.RequestTime, base.MerchantNo)

	b, _ = helper.Sha1WithRsaSignPkcs8([]byte(url), base.Key)

	base.Sign = base64.StdEncoding.EncodeToString(b)

	base.EncryKey, _ = helper.Rsa2Encrypt3([]byte(base64.StdEncoding.EncodeToString([]byte(aesKey))), base.PublicKey)
}

func (base *BaseCharge) SendReq(reqUrl string, params interface{}) (b []byte, err error) {

	base.makeSign(params)

	b, err = json.Marshal(params)

	mapData := make(map[string]interface{})

	mapData["requestTime"] = base.RequestTime
	mapData["version"] = base.Version
	mapData["merchantNo"] = base.MerchantNo
	mapData["requestData"] = base.RequestData
	mapData["sign"] = base.Sign
	mapData["encryKey"] = base.EncryKey

	values := maps.Map2Values(&mapData)

	rsp, err := http.Client.PostForm(reqUrl+base.Service, values)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	return
}
