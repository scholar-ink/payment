package allinpay

import (
	"encoding/json"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/map"
)

const (
	ChargeUrl = "https://syb.allinpay.com/sappweb/usertrans/cuspay"
	SUCCESS   = "SUCCESS"
)

type BaseCharge struct {
	*BaseConfig
}

type BaseConfig struct {
	Orgid string `json:"orgid,omitempty"`
	AppId string `json:"appid"`
	C     string `json:"c"`
	Key   string `json:"key"`
	Sign  string `json:"sign"`
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

	values.Del("key")

	return
}
