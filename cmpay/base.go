package cmpay

import (
	"encoding/json"
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"
	"github.com/scholar-ink/payment/util/map"
	"io/ioutil"
	"strings"
)

const (
	ChargeUrl = "http://pay.congmingpay.com/pay/"
	SUCCESS   = "SUCCESS"
)

type BaseCharge struct {
	*BaseConfig
}

type BaseConfig struct {
	Service string `json:"-"`
	Key     string `json:"-"`
	Sign    string `json:"sign"`
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
	signStr := fmt.Sprintf("money=%s&orderId=%s&shopId=%s", params["money"], params["orderId"], params["shopId"])

	fmt.Println(signStr + "&key=" + base.Key)

	base.Sign = strings.ToUpper(helper.Md5(signStr + "&key=" + base.Key))
}

func (base *BaseCharge) GetReq(reqUrl string, params interface{}) string {
	mapData := helper.Struct2Map(params)

	base.makeSign(mapData)

	mapData["sign"] = base.Sign

	values := maps.Map2Values(&mapData)

	return reqUrl + base.Service + "?" + values.Encode()

}

func (base *BaseCharge) SendReq(reqUrl string, params interface{}) (b []byte, err error) {

	mapData := helper.Struct2Map(params)

	base.makeSign(mapData)

	mapData["sign"] = base.Sign

	values := maps.Map2Values(&mapData)

	rsp, err := http.Client.PostForm(reqUrl+base.Service, values)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	return
}
