package jttp

import (
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"

	"encoding/base64"
	"encoding/json"
	"github.com/scholar-ink/payment/util/map"
	"io/ioutil"
)

const (
	ChargeUrl = "https://mrch.gztyang.com/jttpPay/payProcess"
	SUCCESS   = "SUCCESS"
)

type BaseCharge struct {
	*BaseConfig
	Version    string `json:"version"`
	SignType   string `json:"signType"`   //签名方式
	Charset    string `json:"charset"`    //参数编码字符集
	TxnType    string `json:"txnType"`    //交易类型
	BizContext string `json:"bizContext"` //支付通道
	Sign       string `json:"sign"`
}

type BaseConfig struct {
	SignKey string `json:"-"`
	AesKey  string `json:"-"`
	NodeId  string `json:"nodeId"` //节点号
	OrgId   string `json:"orgId"`  //商户编号
}

func (base *BaseCharge) InitBaseConfig(config *BaseConfig) {
	base.SignType = "RSA2"
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
	signB, _ := json.Marshal(params)
	b, _ := helper.Sha256WithRsaSign(signB, base.SignKey)
	base.Sign = base64.StdEncoding.EncodeToString(b)

	aesKey, _ := base64.StdEncoding.DecodeString(base.AesKey)
	cipherText := helper.PKCS7Padding(signB, 16)
	b, _ = helper.AesEncrypt(cipherText, aesKey) //ECB加密
	base.BizContext = base64.StdEncoding.EncodeToString(b)
}

func (base *BaseCharge) SendReq(reqUrl string, params interface{}) (b []byte, err error) {

	mapData := helper.Struct2Map(params)

	base.makeSign(mapData)

	mapData["version"] = base.Version
	mapData["signType"] = base.SignType
	mapData["charset"] = base.Charset
	mapData["txnType"] = base.TxnType
	mapData["nodeId"] = base.NodeId
	mapData["orgId"] = base.OrgId
	mapData["sign"] = base.Sign
	mapData["bizContext"] = base.BizContext

	values := maps.Map2Values(&mapData)

	rsp, err := http.Client.PostForm(reqUrl, values)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	return
}
