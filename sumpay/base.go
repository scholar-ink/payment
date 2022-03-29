package sumpay

import (
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"
	"log"

	"encoding/base64"
	"encoding/json"
	"github.com/scholar-ink/payment/util/map"
	"io/ioutil"
	"time"
)

const (
	ChargeUrl = "https://entrance.sumpay.cn/gateway.htm"
	SUCCESS   = "SUCCESS"
)

type BaseCharge struct {
	*BaseConfig
}

type BaseConfig struct {
	AppId        string `json:"app_id"`
	Service      string `json:"service"`
	Version      string `json:"version"`
	Format       string `json:"format"`
	Timestamp    string `json:"timestamp"`
	TerminalType string `json:"terminal_type"`
	SignType     string `json:"sign_type"`
	Sign         string `json:"sign"`
	AesKey       string `json:"aes_key"`
	PfxData      []byte `json:"-"`
	PublicKey    string `json:"-"`
	CertPassWord string `json:"-"`
}

func (base *BaseCharge) InitBaseConfig(config *BaseConfig) {
	config.Format = "JSON"
	config.Timestamp = time.Now().Format("20060102150405")
	config.TerminalType = "web"
	config.SignType = "CERT"
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

	aesKey := helper.NonceStr()

	aesKey, _ = helper.Rsa2Encrypt([]byte(base64.StdEncoding.EncodeToString([]byte(aesKey))), base.PublicKey)

	params["aes_key"] = aesKey

	sign := helper.CreateLinkString(&params)

	b, _ := helper.Sha256WithRsaSignWithPassWord(base.PfxData, []byte(sign), base.CertPassWord)

	base.Sign = base64.StdEncoding.EncodeToString(b)
}

func (base *BaseCharge) SendReq(reqUrl string, params interface{}) (b []byte, err error) {

	mapData := helper.Struct2Map(params)

	base.makeSign(mapData)

	mapData["sign"] = base.Sign

	values := maps.Map2Values(&mapData)

	log.Println("[charge." + base.Service + "] request:" + values.Encode())

	rsp, err := http.Client.PostForm(reqUrl, values)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	log.Println("[charge.xh." + base.Service + "] response:" + string(b))

	return
}
