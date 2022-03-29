package yqb

import (
	"bytes"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"
	"log"

	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/scholar-ink/payment/util/map"
	"io/ioutil"
)

const (
	ChargeUrl = "https://scancode.1qianbao.com/scancode/service/gateway"
	SUCCESS   = "SUCCESS"
)

type BaseCharge struct {
	ServiceName string `json:"serviceName"` //服务接口名
	Version     string `json:"version"`     //版本号
	SignType    string `json:"signType"`    //签名类型
	Charset     string `json:"charset"`     //编码格式
	Sign        string `json:"-"`           //签名
	PayScene    string `json:"payScene"`    //支付场景 01-微信02-支付宝
	*BaseConfig
}

type BaseConfig struct {
	MerchantId string `json:"merchantId"` //平安付商户号
	Key        string `json:"-"`
}

func (base *BaseCharge) InitBaseConfig(config *BaseConfig) {
	base.SignType = "RSA"
	base.Charset = "UTF-8"
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

	helper.KSort(&params)

	b, _ := json.Marshal(params)

	body := string(b)

	headerMap := maps.Struct2Map(base)

	helper.KSort(&headerMap)

	b, _ = json.Marshal(headerMap)

	header := string(b)

	b, err := helper.Sha1WithRsaSign([]byte(header+body), base.Key)

	if err != nil {
		fmt.Println(err)
	}

	base.Sign = base64.StdEncoding.EncodeToString(b)
}

func (base *BaseCharge) SendReq(reqUrl string, params interface{}) (b []byte, err error) {

	mapData := helper.Struct2Map(params)

	base.makeSign(mapData)

	buffer := bytes.NewBuffer(b)

	body := struct {
		Head *BaseCharge       `json:"head"`
		Body interface{}       `json:"body"`
		Sign map[string]string `json:"sign"`
	}{
		Head: base,
		Body: mapData,
		Sign: map[string]string{
			"signContent": base.Sign,
		},
	}

	err = json.NewEncoder(buffer).Encode(body)

	if err != nil {
		return nil, err
	}

	log.Println("[trade.yqb." + base.ServiceName + "." + base.PayScene + "] request:" + buffer.String())

	rsp, err := http.Client.Post(reqUrl, "application/json;charset=UTF-8", buffer)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	log.Println("[trade.yqb." + base.ServiceName + "." + base.PayScene + "] response:" + string(b))

	return
}
