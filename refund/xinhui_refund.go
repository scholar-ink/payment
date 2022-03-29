package refund

import (
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"
	"log"

	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/scholar-ink/payment/util/map"
	"io/ioutil"
	"strings"
)

const (
	XinHuiRefundUrl = "http://at.xhepay.com/uni/gateway"
)

type XinHuiConfig struct {
	AgentMerNo string `json:"agent_mer_no"`
	Key        string `json:"-"`
}

type XinHuiRefund struct {
	*XinHuiConfig
	ServiceType string `json:"service_type"`
	RequestId   string `json:"request_id"`
	Sign        string `json:"sign"`
	Version     string `json:"version"`
}

type XinHuiRefundConf struct {
	MerchantNo  string `json:"merchant_no"`            //商户编号
	OutRefundNo string `json:"out_refund_no"`          //商户退款订单号(唯一编号且长度小于等于32)
	OutTradeNo  string `json:"out_trade_no,omitempty"` //商户订单号
	RefundFee   string `json:"refund_fee"`             //退款金额，单位到分
}

type XinHuiRefundReturn struct {
	MerchantNo       string `json:"merchant_no"`
	OutRefundNo      string `json:"out_refund_no"`   //商户退款订单号
	OutTradeNo       string `json:"out_trade_no"`    //商户订单号(唯一编号且长度小于等于32)
	RefundExternalId string `json:"rfd_external_id"` //信汇平台退款流水号	M
}

func (xh *XinHuiRefund) InitConfig(config *XinHuiConfig) {
	xh.XinHuiConfig = config
}

func (xh *XinHuiRefund) Handle(config *XinHuiRefundConf) (refundNo string, er error) {
	xh.ServiceType = "xh.uni.trx.refund"
	xh.Version = "1.0"
	xh.RequestId = strings.ReplaceAll(uuid.New().String(), "-", "")

	ret, err := xh.sendReq(XinHuiRefundUrl, config)
	if err != nil {
		return "", err
	}

	err = xh.retData(ret)

	if err != nil {
		return "", err
	}

	refundReturn := new(XinHuiRefundReturn)
	err = json.Unmarshal(ret, &refundReturn)
	if err != nil {
		return "", err
	}

	return refundReturn.RefundExternalId, nil
}

func (xh *XinHuiRefund) buildData(config interface{}) error {
	return nil
}
func (xh *XinHuiRefund) setSign(params map[string]interface{}) {

	signStr := helper.CreateLinkString(&params)

	b, err := helper.Sha256WithRsaSign([]byte(signStr), xh.Key)

	fmt.Println(err)

	xh.Sign = base64.StdEncoding.EncodeToString(b)

	fmt.Println(xh.Sign)
}

func (xh *XinHuiRefund) sendReq(reqUrl string, params interface{}) (b []byte, err error) {

	mapData := helper.Struct2Map(params)

	mapData["service_type"] = xh.ServiceType
	mapData["request_id"] = xh.RequestId
	mapData["version"] = xh.Version
	mapData["agent_mer_no"] = xh.AgentMerNo

	xh.setSign(mapData)

	if cardNo, ok := mapData["settle_card_no"]; ok {

		b, _ := helper.Rsa2Encrypt2([]byte(cardNo.(string)), xh.Key)

		mapData["settle_card_no"] = base64.StdEncoding.EncodeToString(b)
	}

	mapData["sign"] = xh.Sign

	values := maps.Map2Values(&mapData)

	rsp, err := http.Client.PostForm(reqUrl, values)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	log.Println("[enter.xh." + xh.ServiceType + "] response:" + string(b))

	return
}
func (xh *XinHuiRefund) retData(ret []byte) (err error) {

	var baseReturn struct {
		Version string `json:"version"`
		RspCode string `json:"rsp_code"`
		RspMsg  string `json:"rsp_msg"`
		Sign    string `json:"sign"`
	}

	err = json.Unmarshal(ret, &baseReturn)

	if err != nil {
		return
	}

	if baseReturn.RspCode != "0000" {
		err = errors.New(baseReturn.RspMsg)
		return
	}
	return
}
