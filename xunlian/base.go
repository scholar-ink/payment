package xunlian

import (
	"bytes"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"
	"log"

	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	ChargeUrl = "https://showmoney.cn/scanpay/unified"
	SUCCESS   = "SUCCESS"
)

type BaseCharge struct {
	BusiCd   string `json:"busicd"`   //交易类型
	Version  string `json:"version"`  //版本号
	SignType string `json:"signType"` //签名类型
	Charset  string `json:"charset"`  //编码格式
	Sign     string `json:"-"`        //签名
	*BaseConfig
}

type BaseConfig struct {
	MerchantId string `json:"mchntid"`    //商户号
	InScd      string `json:"inscd"`      //机构号
	TerminalId string `json:"terminalid"` //机构号
	Key        string `json:"-"`
}

func (base *BaseCharge) InitBaseConfig(config *BaseConfig) {
	base.SignType = "SHA256"
	base.Charset = "utf-8"
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

	signStr := helper.CreateLinkString(&params)

	signStr += base.Key

	fmt.Println(signStr)

	base.Sign = helper.Sha256(signStr)
}

func (base *BaseCharge) SendReq(reqUrl string, params interface{}) (b []byte, err error) {

	mapData := helper.Struct2Map(params)

	mapData["busicd"] = base.BusiCd
	mapData["version"] = base.Version
	mapData["signType"] = base.SignType
	mapData["charset"] = base.Charset
	mapData["mchntid"] = base.MerchantId
	mapData["inscd"] = base.InScd
	mapData["terminalid"] = base.TerminalId

	fmt.Println(mapData)

	base.makeSign(mapData)

	mapData["sign"] = base.Sign

	buffer := bytes.NewBuffer(b)

	err = json.NewEncoder(buffer).Encode(mapData)

	if err != nil {
		return nil, err
	}

	log.Println("[trade.xunlian." + base.BusiCd + "] request:" + buffer.String())

	rsp, err := http.Client.Post(reqUrl, "application/json;charset=UTF-8", buffer)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	log.Println("[trade.xunlian." + base.BusiCd + "] response:" + string(b))

	return
}

func (base *BaseCharge) GetReqUrl(reqUrl string, params interface{}) (string, error) {

	mapData := helper.Struct2Map(params)

	base.makeSign(mapData)

	mapData["sign"] = base.Sign

	buffer := bytes.NewBuffer([]byte{})

	err := json.NewEncoder(buffer).Encode(mapData)

	if err != nil {
		return "", err
	}

	log.Println("[trade.xunlian." + base.BusiCd + "] request:" + buffer.String())

	return reqUrl + "?data=" + base64.StdEncoding.EncodeToString(buffer.Bytes()), nil
}
