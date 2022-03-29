package enter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"io/ioutil"
)

const (
	UploadUrl = "http://mgateway.g-pay.cn/image/upload.do"
	EnterUrl  = "http://mgateway.g-pay.cn/merchant/enter/api.do"
	SUCCESS   = "SUCCESS"
)

type BaseCharge struct {
	*BaseConfig
}

type BaseConfig struct {
	ServiceName  string `json:"serviceName"`
	Charset      string `json:"charset"`
	Version      string `json:"version"`
	AgentNo      string `json:"agentNo"`
	Key          string `json:"key"`
	PfxData      []byte `json:"pfx_data"`
	CertPassWord string `json:"certPassWord"`
	EncryptType  string `json:"encryptType"`
	EncryptData  string `json:"encryptData"`
	SignData     string `json:"signData"`
	ResponseCode string `xml:"responseCode" json:"responseCode"`
	ResponseMsg  string `xml:"responseMsg" json:"responseMsg"`
}

func (base *BaseCharge) InitBaseConfig(config *BaseConfig) {
	config.Version = "3.0"
	config.Charset = "UTF-8"
	config.EncryptType = "RSA"
	base.BaseConfig = config
}

func (base *BaseCharge) SetSign() {
	base.SignData = helper.Md5(base.EncryptData + base.Key)
}

func (base *BaseCharge) RetData(ret []byte) (retData []byte, err error) {
	var baseReturn BaseConfig

	err = json.Unmarshal(ret, &baseReturn)

	if err != nil {
		return
	}

	if baseReturn.ResponseCode != "" {
		err = errors.New(baseReturn.ResponseMsg)
		return
	}

	//MD5值验证
	stringMd5 := helper.Md5(baseReturn.EncryptData + base.Key)

	if stringMd5 != baseReturn.SignData {
		fmt.Print("MD5值验证不通过，远程MD5值：" + baseReturn.SignData + ",计算出MD5值：" + stringMd5)
		return nil, errors.New("MD5值验证不通过")
	}

	retData, err = helper.Rsa1Decrypt(base.PfxData, baseReturn.EncryptData, base.CertPassWord)

	return
}

func (base *BaseCharge) makeSign(sign string) string {

	//switch base.SignType {
	//
	//case "MD5":
	//	sign += "&key=" + base.Md5Key
	//
	//	sign = helper.Md5(sign)
	//}

	//return strings.ToUpper(sign)
	return ""
}

func (base *BaseCharge) SendReq(reqUrl string, pay interface{}) (b []byte, err error) {

	buffer := bytes.NewBuffer(b)

	err = json.NewEncoder(buffer).Encode(pay)

	if err != nil {
		return nil, err
	}

	client := helper.NewHttpClient()

	rsp, err := client.Post(reqUrl, "application/json;charset=UTF-8", buffer)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	return
}
