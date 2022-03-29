package qf

import (
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"

	"encoding/json"
	"fmt"
	"github.com/scholar-ink/payment/util/map"
	"io/ioutil"
	"strings"
)

const (
	ChargeUrl = "https://openapi.qfpay.com"
	SUCCESS   = "SUCCESS"
)

type BaseCharge struct {
	*BaseConfig
}

type BaseConfig struct {
	AppCode string `json:"-"`
	Key     string `json:"-"`
	Service string `json:"-"`
	Sign    string `json:"-"`
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

	signStr := helper.CreateLinkString(&params)

	fmt.Println(signStr + base.Key)

	base.Sign = strings.ToUpper(helper.Md5(signStr + base.Key))
}

func (base *BaseCharge) SendReq(reqUrl string, params interface{}) (b []byte, err error) {

	mapData := helper.Struct2Map(params)

	base.makeSign(mapData)

	values := maps.Map2Values(&mapData)

	req := http.NewHttpRequest("POST", reqUrl+base.Service, strings.NewReader(values.Encode()))

	req.Header["X-QF-APPCODE"] = []string{base.AppCode}
	req.Header["X-QF-SIGN"] = []string{base.Sign}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rsp, err := http.Client.Do(req)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	return
}
