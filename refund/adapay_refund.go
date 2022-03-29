package refund

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
	AdaPayRefundUrl = "https://api.adapay.tech/v1/"
)

type AdaPayConfig struct {
	ApiKey     string `json:"-"`
	PrivateKey string `json:"-"`
}

type AdaPayRefund struct {
	*AdaPayConfig
	Service string `json:"-"`
	Sign    string `json:"sign"`
}

type AdaPayRefundConf struct {
	PaymentId     string `json:"payment_id"`       //原交易支付对象ID
	RefundOrderNo string `json:"refund_order_no"`  //退款订单号
	RefundAmt     string `json:"refund_amt"`       //退款金额
	Reason        string `json:"reason,omitempty"` //退款描述
	SignType      string `json:"sign_type"`
}

type AdaPayRefundReturn struct {
	Id             string `json:"id"`               //平台退款订单号
	RefundOrderNo  string `json:"refund_order_no"`  //商户退款订单号
	PaymentId      string `json:"payment_id"`       //原平台付订单号
	PaymentOrderNo string `json:"payment_order_no"` //原商户订单号
	RefundAmt      string `json:"refund_amt"`       //退款金额
	SucceedTime    string `json:"succeed_time"`     //退款成功时间
}

func (ada *AdaPayRefund) InitConfig(config *AdaPayConfig) {
	ada.AdaPayConfig = config
}

func (ada *AdaPayRefund) Handle(config *AdaPayRefundConf) (refundNo string, er error) {
	ada.Service = "payments" + "/" + config.PaymentId + "/refunds"
	config.SignType = "RSA2"
	ret, err := ada.sendReq(AdaPayRefundUrl, config)
	if err != nil {
		return "", err
	}

	ret, err = ada.retData(ret)

	if err != nil {
		return "", err
	}

	refundReturn := new(AdaPayRefundReturn)
	err = json.Unmarshal(ret, &refundReturn)
	if err != nil {
		return "", err
	}

	return refundReturn.Id, nil
}

func (ada *AdaPayRefund) buildData(config interface{}) error {
	return nil
}
func (ada *AdaPayRefund) setSign(reqUrl string, requestJson string) {

	sign := reqUrl + requestJson

	b, err := helper.Sha1WithRsaSignPkcs8([]byte(sign), ada.PrivateKey)

	if err != nil {
		fmt.Println(err)
	}

	ada.Sign = base64.StdEncoding.EncodeToString(b)
}

func (ada *AdaPayRefund) sendReq(reqUrl string, params interface{}) (b []byte, err error) {

	b, err = json.Marshal(params)

	if err != nil {
		return nil, err
	}

	ada.setSign(reqUrl+ada.Service, string(b))

	req := http.NewHttpRequest("POST", reqUrl+ada.Service, bytes.NewBuffer(b))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("authorization", ada.ApiKey)
	req.Header.Add("sdk_version", "go_v1.0.2")
	req.Header.Add("signature", ada.Sign)

	rsp, err := http.Client.Do(req)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	log.Println("[trade.adapay." + ada.Service + "] response:" + string(b))

	return
}
func (ada *AdaPayRefund) retData(ret []byte) (b []byte, err error) {

	type BaseResponse struct {
		Data string `json:"data"`
	}

	baseResponse := new(BaseResponse)

	err = json.Unmarshal(ret, &baseResponse)

	if err != nil {
		return
	}

	type BaseReturn struct {
		ErrorCode string `json:"error_code"`
		ErrorMsg  string `json:"error_msg"`
	}

	baseReturn := new(BaseReturn)

	b = []byte(baseResponse.Data)

	err = json.Unmarshal(b, &baseReturn)

	if err != nil {
		return
	}

	if baseReturn.ErrorCode != "" {
		err = errors.New(baseReturn.ErrorMsg)
		return
	}
	return
}
