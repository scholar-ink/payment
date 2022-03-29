package hj

import (
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"

	"encoding/json"
	"github.com/scholar-ink/payment/util/map"
	"io/ioutil"
)

const (
	ChargeUrl = "https://www.joinpay.com/trade/uniPayApi.action"
	SUCCESS   = "SUCCESS"
)

type BaseCharge struct {
	*BaseConfig
}

type BaseConfig struct {
	Md5Key string `xml:"-" json:"-"`
	Hmac   string `json:"hmac"` //签名
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

func (base *BaseCharge) makeSign(sign string) string {
	sign = helper.Md5(sign)
	return sign
}

func (base *BaseCharge) SendReq(reqUrl string, params interface{}) (b []byte, err error) {

	mapData := helper.Struct2Map(params)

	values := maps.Map2Values(&mapData)

	rsp, err := http.Client.PostForm(reqUrl, values)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	return
}
