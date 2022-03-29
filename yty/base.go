package yty

import (
	"bytes"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"
	"log"

	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

const (
	ChargeUrl = "http://101.200.137.12:19000/office/api/"
	SUCCESS   = "SUCCESS"
)

type BaseCharge struct {
	*BaseConfig
	Enck      string `json:"-"`
	Encd      string `json:"-"`
	Service   string `json:"-"`
	ProductNo string `json:"product_no"`
	Sign      string `json:"-"`
}

type BaseConfig struct {
	AgentNo    string `json:"agent_no"`
	MerchantNo string `json:"merchant_no"`
	Key        string `json:"-"`
	PubKey     string `json:"-"`
}

type BaseReturn struct {
	Status  int32  `json:"status"`
	Message string `json:"message"`
	Encd    string `json:"encd"`
	Sign    string `json:"sign"`
	Enck    string `json:"enck"`
}

func (base *BaseCharge) InitBaseConfig(config *BaseConfig) {
	base.BaseConfig = config
}

func (base *BaseCharge) RetData(ret []byte) (retData []byte, err error) {

	var baseReturn BaseReturn

	err = json.Unmarshal(ret, &baseReturn)

	if err != nil {
		return
	}

	if baseReturn.Status == 500 {
		return nil, errors.New(baseReturn.Message)
	}

	fmt.Println(baseReturn.Enck)
	fmt.Println(base.Key)

	enck, err := helper.Rsa1Decrypt2(baseReturn.Enck, base.Key)

	if err != nil {
		return
	}

	b, _ := base64.StdEncoding.DecodeString(baseReturn.Encd)

	b, err = helper.TripleEcbDesDecrypt(b, enck)

	return b, err
}

func (base *BaseCharge) makeSign(params interface{}) {
	b, _ := json.Marshal(params)

	b, err := helper.TripleEcbDesEncrypt(b, []byte(base.Enck))

	if err != nil {
		fmt.Println(err)
		base.Sign = ""
	} else {
		data := base64.StdEncoding.EncodeToString(b)

		base.Encd = data

		base.Enck, _ = helper.Rsa1Encrypt2([]byte(base.Enck), base.PubKey)

		sign := fmt.Sprintf("enck=%s&encd=%s", base.Enck, base.Encd)

		b, _ := helper.Sha256WithRsaSignWithPKCS8([]byte(sign), base.Key)

		base.Sign = base64.StdEncoding.EncodeToString(b)
	}
}

func (base *BaseCharge) SendReq(reqUrl string, params interface{}) (b []byte, err error) {

	mapData := helper.Struct2Map(params)

	mapData["product_no"] = base.ProductNo
	mapData["agent_no"] = base.AgentNo
	mapData["merchant_no"] = base.MerchantNo

	base.makeSign(mapData)

	reqData := make(map[string]string, 0)

	reqData["enck"] = base.Enck
	reqData["encd"] = base.Encd
	reqData["sign"] = base.Sign

	buffer := bytes.NewBuffer(b)

	err = json.NewEncoder(buffer).Encode(reqData)

	if err != nil {
		return nil, err
	}

	log.Println("[trade.yitong." + base.Service + "] request:" + buffer.String())

	rsp, err := http.Client.Post(reqUrl+base.Service, "application/json;charset=UTF-8", buffer)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	log.Println("[trade.yitong." + base.Service + "] response:" + string(b))

	return
}
