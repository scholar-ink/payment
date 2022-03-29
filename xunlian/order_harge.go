package xunlian

import (
	"encoding/json"
	"errors"
)

type OrderCharge struct {
	BaseCharge
	*OrderChargeConf
}

type OrderChargeConf struct {
	OrderNum   string `json:"orderNum"`             //订单号
	ChCd       string `json:"chcd"`                 //渠道
	TxAmt      string `json:"txamt"`                //订单金额
	Subject    string `json:"subject"`              //订单标题
	BackUrl    string `json:"backUrl"`              //通知地址
	FrontUrl   string `json:"frontUrl,omitempty"`   //通知地址
	PayLimit   string `json:"paylimit,omitempty"`   //通知地址
	TimeExpire string `json:"timeExpire,omitempty"` //通知地址
	TimeStart  string `json:"timeStart,omitemptyt"` //通知地址
}

type QueryConf struct {
	OrigOrderNum    string `json:"origOrderNum,omitempty"`    //商户订单号
	ChannelOrderNum string `json:"channelOrderNum,omitempty"` //平安付订单号
	Terminalid      string `json:"terminalid"`                //平安付聚合平台服务商号，平安付分配
	Mchntid         string `json:"mchntid"`                   //平安付聚合平台服务商号，平安付分配
}

type RefundConf struct {
	OrigOrderNum string `json:"origOrderNum,omitempty"` //原商户订单号
	OrderNum     string `json:"orderNum"`               //商户退款订单号
	TxAmt        string `json:"txamt"`                  //退款金额
}

type OrderChargeReturn struct {
	RespCd      string `json:"respcd"`
	ErrorDetail string `json:"errorDetail"`
	QrCode      string `json:"qrcode"`
}

type QueryReturn struct {
	BusinessCode       string `json:"businessCode"`
	BusinessMsg        string `json:"businessMsg"`
	TradeOrderNo       string `json:"tradeOrderNo"`       //平安付订单号
	MerTradeNo         string `json:"merTradeNo"`         //商户订单号
	TotalAmount        string `json:"totalAmount"`        //订单总金额（分）
	PayTime            string `json:"payTime"`            //支付时间
	PlatformMerchantId string `json:"platformMerchantId"` //平安付聚合平台服务商号，平安付分配
	MerchantId         string `json:"merchantId"`         //平安付商户号
	ChannelOrderNo     string `json:"channelOrderNo"`     //平安付商户号
}

type RefundReturn struct {
	RespCd          string `json:"respcd"`
	ErrorDetail     string `json:"errorDetail"`
	ChannelOrderNum string `json:"channelOrderNum"`
	OrderNum        string `json:"orderNum"`
	OrigOrderNum    string `json:"origOrderNum"`
}

func (oc *OrderCharge) Handle(conf *OrderChargeConf) (string, error) {
	oc.BaseCharge.Version = "2.3.1"
	oc.BusiCd = "PAUT"
	ret, err := oc.SendReq(ChargeUrl, conf)
	if err != nil {
		return "", err
	}

	orderChargeReturn := new(OrderChargeReturn)
	err = json.Unmarshal(ret, &orderChargeReturn)
	if err != nil {
		return "", err
	}

	if orderChargeReturn.RespCd != "09" {
		return "", errors.New(orderChargeReturn.ErrorDetail)
	}

	return orderChargeReturn.QrCode, nil
}

func (oc *OrderCharge) Query(conf *QueryConf) (*QueryReturn, error) {
	oc.BaseCharge.Version = "2.3.1"
	oc.BusiCd = "INQY"
	ret, err := oc.SendReq(ChargeUrl, conf)
	if err != nil {
		return nil, err
	}

	queryReturn := new(QueryReturn)
	err = json.Unmarshal(ret, &queryReturn)
	if err != nil {
		return nil, err
	}
	if queryReturn.BusinessCode != "00" {
		return nil, errors.New(queryReturn.BusinessMsg)
	}
	return queryReturn, nil
}

func (oc *OrderCharge) Refund(conf *RefundConf) (string, error) {
	oc.BaseCharge.Version = "2.3.1"
	oc.BusiCd = "REFD"
	ret, err := oc.SendReq(ChargeUrl, conf)
	if err != nil {
		return "", err
	}

	refundReturn := new(RefundReturn)
	err = json.Unmarshal(ret, &refundReturn)
	if err != nil {
		return "", err
	}

	if refundReturn.RespCd != "00" {
		return "", errors.New(refundReturn.ErrorDetail)
	}

	return refundReturn.ChannelOrderNum, nil
}
