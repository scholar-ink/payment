package adapay

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"
	"io/ioutil"
)

const (
	ChargeUrl = "https://api.adapay.tech/v1/"
	SUCCESS   = "SUCCESS"
)

type BaseCharge struct {
	Service string `json:"-"`
	Sign    string `json:"-"` //签名
	*BaseConfig
}

type BaseConfig struct {
	ApiKey     string `json:"-"`
	PrivateKey string `json:"-"`
}

func (base *BaseCharge) InitBaseConfig(config *BaseConfig) {
	base.BaseConfig = config
}

func (base *BaseCharge) RetData(ret []byte) (retData []byte, err error) {
	var baseReturn BaseConfig

	err = json.Unmarshal(ret, &baseReturn)

	if err != nil {
		return
	}
	return
}

func (base *BaseCharge) makeSign(reqUrl string, requestJson string) {

	sign := reqUrl + requestJson

	b, err := helper.Sha1WithRsaSignPkcs8([]byte(sign), base.PrivateKey)

	if err != nil {
		fmt.Println(err)
	}

	base.Sign = base64.StdEncoding.EncodeToString(b)
}

func (base *BaseCharge) SendReq(reqUrl string, params interface{}) (b []byte, err error) {

	b, err = json.Marshal(params)

	if err != nil {
		return nil, err
	}

	base.makeSign(reqUrl+base.Service, string(b))

	req := http.NewHttpRequest("POST", reqUrl+base.Service, bytes.NewBuffer(b))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("authorization", base.ApiKey)
	req.Header.Add("sdk_version", "go_v1.0.2")
	req.Header.Add("signature", base.Sign)

	rsp, err := http.Client.Do(req)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	return
}
