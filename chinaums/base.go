package chinaums

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"
	"io/ioutil"
	"strings"
	"time"
)

const (
	ChargeUrl = "https://qr.chinaums.com/netpay-route-server/api/"
	SUCCESS   = "SUCCESS"
)

type BaseCharge struct {
	*BaseConfig
	RequestTimestamp string `json:"requestTimestamp"`
	SignType         string `json:"signType"`
	MsgType          string `json:"msgType"`
	Sign             string `json:"sign"`
}

type BaseConfig struct {
	MsgId   string `json:"-"`
	MsgSrc  string `json:"msgSrc"`
	Mid     string `json:"mid"`
	Tid     string `json:"tid"`
	InstMid string `json:"instMid"`
	Md5Key  string `json:"-"`
}

type BaseReturn struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func (base *BaseCharge) InitBaseConfig(config *BaseConfig) {
	base.RequestTimestamp = time.Now().Format("2006-01-02 15:04:05")
	base.SignType = "MD5"
	base.BaseConfig = config
}

func (base *BaseCharge) makeSign(params map[string]interface{}) {
	helper.KSort(&params)

	signStr := helper.CreateLinkString(&params)

	signStr += base.Md5Key

	base.Sign = strings.ToUpper(helper.Md5(signStr))

	fmt.Println(base.Sign)
}

func (base *BaseCharge) SendReq(reqUrl string, params interface{}) (b []byte, err error) {

	mapData := helper.Struct2Map(params)

	mapData["requestTimestamp"] = base.RequestTimestamp
	mapData["signType"] = base.SignType
	mapData["msgType"] = base.MsgType
	mapData["msgSrc"] = base.MsgSrc
	mapData["mid"] = base.Mid
	mapData["tid"] = base.Tid
	mapData["instMid"] = base.InstMid

	base.makeSign(mapData)

	mapData["sign"] = base.Sign

	buffer := bytes.NewBuffer(b)

	err = json.NewEncoder(buffer).Encode(mapData)

	if err != nil {
		return nil, err
	}

	rsp, err := http.Client.Post(reqUrl, "application/json;charset=UTF-8", buffer)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	return
}
