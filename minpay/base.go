package minpay

import (
	"bytes"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"

	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	ChargeUrl = "https://api.minpayment.com/pay.do"
	SUCCESS   = "SUCCESS"
)

type BaseCharge struct {
	*BaseConfig
	ServiceName string      `json:"serviceName"`
	Version     string      `json:"version"`
	MsgBody     interface{} `json:"msgBody"`
	Signature   string      `json:"signature,omitempty"`
	Body        string      `json:"-"`
}

type BaseConfig struct {
	MerchantNo   string `json:"merchantNo"` //商户编号
	Key          []byte `json:"-"`
	PubKey       string `json:"-"`
	CertPassword string `json:"-"`
}

type BaseReturn struct {
	Status  int32  `json:"status"`
	Message string `json:"message"`
	Sign    string `json:"sign"`
}

func (base *BaseCharge) InitBaseConfig(config *BaseConfig) {
	base.BaseConfig = config
}

func (base *BaseCharge) RetData(ret []byte) (retData []byte, err error) {

	b, _ := base64.StdEncoding.DecodeString(string(ret))

	msgSign := base.getCipherMsg(string(b))

	dgt := base.getDgt(string(b))

	b, _ = helper.Rsa1Decrypt(base.Key, dgt, base.CertPassword)

	aes := string(b)

	msgSignB, _ := base64.StdEncoding.DecodeString(msgSign)

	b, _ = helper.AesDecrypt(msgSignB, []byte(aes))

	b = helper.PKCS7UnPadding(b)

	var baseReturn struct {
		MsgBody interface{} `json:"msgBody"`
	}

	_ = json.Unmarshal(b, &baseReturn)

	return json.Marshal(baseReturn.MsgBody)
}

func (base *BaseCharge) makeSign() {
	b, _ := json.Marshal(base)

	b, _ = helper.Sha256WithRsaSignWithPassWord(base.Key, b, base.CertPassword)

	base.Signature = base64.StdEncoding.EncodeToString(b)

	aesKey := "GheGgmsrxEnlafurzmexsdsbksaAglfA"

	dgt, _ := helper.Rsa2EncryptByCert([]byte(aesKey), base.PubKey)

	b, _ = json.Marshal(base)

	cipherText := helper.PKCS7Padding(b, 16)

	b, _ = helper.AesEncrypt(cipherText, []byte(aesKey)) //ECB加密

	fullMsg := fmt.Sprintf("{%s}{dgtlenvlp:%s}", base64.StdEncoding.EncodeToString(b), dgt)

	base.Body = base64.StdEncoding.EncodeToString([]byte(fullMsg))
}

func (base *BaseCharge) SendReq(reqUrl string, params interface{}) (b []byte, err error) {

	base.MsgBody = params
	base.makeSign()

	buffer := bytes.NewBufferString(base.Body)

	rsp, err := http.Client.Post(reqUrl, "application/json;charset=UTF-8", buffer)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	return
}

func (base *BaseCharge) getCipherMsg(msgFull string) string {
	start := strings.Index(msgFull, "{")
	end := strings.Index(msgFull, "}")

	return msgFull[start+1 : end]
}

func (base *BaseCharge) getDgt(msgFull string) string {
	flag := "dgtlenvlp:"
	start := strings.Index(msgFull, flag)
	end := strings.LastIndex(msgFull, "}")

	return msgFull[start+len(flag) : end]
}
