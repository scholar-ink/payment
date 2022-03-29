package gpay

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/scholar-ink/payment/helper"
	"io/ioutil"
	"strings"
	"time"

	"github.com/scholar-ink/payment/util/http"
)

const (
	ChargeUrl = "https://jh.g-pay.cn/api/order.do"
	SUCCESS   = "SUCCESS"
)

type BaseCharge struct {
	*BaseConfig
}

type BaseConfig struct {
	ServiceName  string `json:"serviceName"`
	Version      string `json:"version"`
	MerchantNo   string `json:"merchantNo"`
	EncryptType  string `json:"encryptType"`
	Md5Key       string `xml:"-" json:"-"`
	SourceData   string `json:"sourceData"`
	SignType     string `json:"signType"`
	SignData     string `json:"signData"`
	RequestId    string `json:"requestId"`
	RequestTime  string `json:"requestTime"`
	ResponseCode string `json:"responseCode,omitempty"`
	ResponseMsg  string `json:"responseMsg,omitempty"`
}

func (base *BaseCharge) InitBaseConfig(config *BaseConfig) {
	config.Version = "2.0"
	config.EncryptType = "NONE"
	config.SignType = "01"
	config.RequestId = uuid.New().String()
	config.RequestTime = time.Now().Format("20060102150405")

	base.BaseConfig = config
}

func (base *BaseCharge) SetSign() {
	mapData := helper.Struct2Map(base)

	signStr := helper.CreateLinkString(&mapData)

	base.SignData = base.makeSign(signStr)
}

func (base *BaseCharge) RetData(ret []byte) (retData []byte, err error) {
	var baseReturn BaseConfig

	err = json.Unmarshal(ret, &baseReturn)

	if err != nil {
		return
	}

	if baseReturn.ResponseCode != "0000" {
		return nil, errors.New(baseReturn.ResponseMsg)
	}

	var mapData map[string]interface{}

	err = json.Unmarshal(ret, &mapData)

	if err != nil {
		return
	}

	delete(mapData, "signData")

	signStr := helper.CreateLinkString(&mapData)

	if baseReturn.SignData != strings.ToUpper(base.makeSign(signStr)) {
		return nil, errors.New("返回结果签名校验失败")
	}

	return []byte(baseReturn.SourceData), nil
}

func (base *BaseCharge) makeSign(sign string) string {
	switch base.SignType {
	case "01":
		sign += "&key=" + base.Md5Key
		sign = helper.Md5(sign)
	}

	return sign
}

func (base *BaseCharge) SendReq(reqUrl string) (b []byte, err error) {

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
