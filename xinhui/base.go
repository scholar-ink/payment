package xinhui

import (
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"
	"log"

	"encoding/base64"
	"encoding/json"
	"github.com/scholar-ink/payment/util/map"
	"io/ioutil"
)

const (
	ChargeUrl = "http://at.xhepay.com/uni/gateway"
	SUCCESS   = "SUCCESS"
)

type BaseCharge struct {
	*BaseConfig
}

type BaseConfig struct {
	AgentMerNo  string `json:"agent_mer_no"`
	Key         string `json:"-"`
	ServiceType string `json:"service_type"`
	RequestId   string `json:"request_id"`
	Sign        string `json:"sign"`
	Version     string `json:"version"`
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

func (base *BaseCharge) makeSign(params map[string]interface{}) {
	sign := helper.CreateLinkString(&params)
	b, _ := helper.Sha256WithRsaSign([]byte(sign), base.Key)
	base.Sign = base64.StdEncoding.EncodeToString(b)
}

func (base *BaseCharge) SendReq(reqUrl string, params interface{}) (b []byte, err error) {

	mapData := helper.Struct2Map(params)

	base.makeSign(mapData)

	mapData["sign"] = base.Sign

	values := maps.Map2Values(&mapData)

	log.Println("[enter." + base.ServiceType + "] request:" + values.Encode())

	rsp, err := http.Client.PostForm(reqUrl, values)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	log.Println("[enter.xh." + base.ServiceType + "] response:" + string(b))

	return
}
