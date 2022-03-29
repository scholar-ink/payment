package fql

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"
	"io/ioutil"
)

const (
	ChargeUrl = "https://eagateway.fuqianla.net/gateway/api/clientApi/interface"
	SUCCESS   = "SUCCESS"
)

type BaseCharge struct {
	*BaseConfig
	VersionNo string      `json:"versionNo"`
	ServiceId string      `json:"serviceId"`
	SignType  string      `json:"signType"`
	Sign      string      `json:"sign"`
	Timestamp string      `json:"timestamp"`
	SeqNo     string      `json:"seqNo"`
	Data      interface{} `json:"-"`
	DataStr   string      `json:"data"`
}

type BaseConfig struct {
	MerchId string `json:"merchId"`
	PriKey  string `json:"-"`
}

type BaseReturn struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func (base *BaseCharge) InitBaseConfig(config *BaseConfig) {
	base.SignType = "RSA2"
	base.BaseConfig = config
}

func (base *BaseCharge) RetData(ret []byte) (retData []byte, err error) {
	var baseReturn BaseReturn

	err = json.Unmarshal(ret, &baseReturn)

	if err != nil {
		return
	}

	if baseReturn.Code != "0000" {
		err = errors.New(baseReturn.Msg)

		return
	}

	return []byte(baseReturn.Data), nil
}

func (base *BaseCharge) makeSign() {
	b, _ := helper.Sha256WithRsaSign([]byte(base.SeqNo+base.DataStr), base.PriKey)

	base.Sign = base64.StdEncoding.EncodeToString(b)
}

func (base *BaseCharge) SendReq(reqUrl string, params interface{}) (b []byte, err error) {

	buffer := bytes.NewBuffer(b)

	err = json.NewEncoder(buffer).Encode(base)

	if err != nil {
		return nil, err
	}

	rsp, err := http.Client.Post(reqUrl, "application/json;charset=UTF-8", buffer)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	return
}

func (base *BaseCharge) buildData(config interface{}) error {
	base.Data = config
	b, _ := json.Marshal(config)

	base.DataStr = string(b)

	return nil
}
