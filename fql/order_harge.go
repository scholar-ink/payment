package fql

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"strings"
	"time"
)

type OrderCharge struct {
	*OrderChargeConf
	BaseCharge
}

type OrderChargeConf struct {
	OrderId        string `json:"orderId"`               //商户订单号
	ChUserId       string `json:"chUserId"`              //商户用户号
	CustomerId     string `json:"customerId"`            //商户客户号
	ChannelType    string `json:"channelType"`           //支付渠道
	PmtType        string `json:"pmtType"`               //支付类型
	OpenId         string `json:"openid,omitempty"`      //用户标识
	Amount         string `json:"amount"`                //交易金额
	Subject        string `json:"subject"`               //商品标题
	SpbillCreateIp string `json:"spbillCreateIp"`        //终端IP
	AuthCode       string `json:"authCode,omitempty"`    //支付授权码
	Body           string `json:"body"`                  //商品描述
	TxnType        string `json:"txnType"`               //交易类型04-虚户充值、24-消费即时到账、25-消费担保交易、26-普通分账、27-批量担保
	TimeExpire     string `json:"time_expire,omitempty"` //交易结束时间
	NotifyUrl      string `json:"notify_url"`            //通知地址
}

type OrderChargeReturn struct {
	SubCode     string `json:"subCode"`
	SubMsg      string `json:"subMsg"`
	OrderId     string `json:"orderId"`     //商户订单号
	TxnId       string `json:"txnId"`       //系统单号
	Status      string `json:"status"`      //交易状态
	ChannelType string `json:"channelType"` //支付渠道
	PmtType     string `json:"pmtType"`     //支付类型
	CodeUrl     string `json:"codeUrl"`     //二维码链接
	ExAccNo     string `json:"exAccNo"`     //虚户账号
	PayExAccNo  string `json:"payExAccNo"`  //付款账号
	PrepayId    string `json:"prepayId"`    //小程序支付、公众号支付：当status为02时返回
	PayInfo     string `json:"payInfo"`     //交易流水号
	WaybillNo   string `json:"waybillNo"`   //业务单号
}

type OrderPaySplitItem struct {
	Amount      string `json:"amount"`
	RecvExAccNo string `json:"recvExAccNo"`
}

type OrderPaySplitConf struct {
	OrderId     string               `json:"orderId"`     //商户订单号
	OrglOrderId string               `json:"orglOrderId"` //原交易订单号
	BizType     string               `json:"bizType"`     //业务类型
	CustomerId  string               `json:"customerId"`  //商户客户号
	List        []*OrderPaySplitItem `json:"list"`
}

type OrderPaySplitReturn struct {
	SubCode string `json:"subCode"`
	SubMsg  string `json:"subMsg"`
	OrderId string `json:"orderId"` //商户订单号
	TxnId   string `json:"txnId"`   //系统单号
	Status  string `json:"status"`  //交易状态
}

type OrderQueryConf struct {
	OrderId     string `json:"orderId"`     //商户订单号
	OrglOrderId string `json:"orglOrderId"` //原交易订单号
	ChUserId    string `json:"chUserId"`    //商户用户号
	CustomerId  string `json:"customerId"`  //商户客户号
	QueryType   string `json:"queryType"`   //查询类型
}

type OrderQueryReturn struct {
	SubCode     string `json:"subCode"`
	SubMsg      string `json:"subMsg"`
	OrderId     string `json:"orderId"`     //商户订单号
	TxnId       string `json:"txnId"`       //系统单号
	Status      string `json:"status"`      //交易状态
	ChannelType string `json:"channelType"` //支付渠道
	PmtType     string `json:"pmtType"`     //支付类型
	Amount      string `json:"amount"`      //交易金额
}

func (oc *OrderCharge) Handle(conf *OrderChargeConf) (*OrderChargeReturn, error) {
	oc.ServiceId = "scanPay"
	oc.Timestamp = time.Now().Format("20060102150405")
	oc.SeqNo = strings.ReplaceAll(uuid.New().String(), "-", "")
	oc.VersionNo = "2.00"
	oc.buildData(conf)
	oc.makeSign()

	ret, err := oc.SendReq(ChargeUrl, oc)
	if err != nil {
		return nil, err
	}

	ret, err = oc.RetData(ret)
	if err != nil {
		return nil, err
	}

	orderChargeReturn := new(OrderChargeReturn)

	err = json.Unmarshal(ret, &orderChargeReturn)

	if err != nil {
		return nil, err
	}

	if orderChargeReturn.SubCode != "0001" && orderChargeReturn.SubCode != "0000" && orderChargeReturn.SubCode != "E0001" {
		return nil, errors.New(orderChargeReturn.SubMsg)
	}

	return orderChargeReturn, nil
}
func (oc *OrderCharge) PaySplit(conf *OrderPaySplitConf) (*OrderPaySplitReturn, error) {
	oc.ServiceId = "bulkPaySplit"
	oc.Timestamp = time.Now().Format("20060102150405")
	oc.SeqNo = strings.ReplaceAll(uuid.New().String(), "-", "")
	oc.VersionNo = "2.00"
	oc.buildData(conf)
	oc.makeSign()

	ret, err := oc.SendReq(ChargeUrl, oc)
	if err != nil {
		return nil, err
	}

	ret, err = oc.RetData(ret)
	if err != nil {
		return nil, err
	}

	orderPaySplitReturn := new(OrderPaySplitReturn)

	err = json.Unmarshal(ret, &orderPaySplitReturn)

	if err != nil {
		return nil, err
	}

	if orderPaySplitReturn.SubCode != "0001" && orderPaySplitReturn.SubCode != "0000" && orderPaySplitReturn.SubCode != "E0001" {
		return nil, errors.New(orderPaySplitReturn.SubMsg)
	}

	return orderPaySplitReturn, nil
}
func (oc *OrderCharge) Query(conf *OrderQueryConf) (*OrderQueryReturn, error) {
	oc.ServiceId = "tradeQuery"
	oc.Timestamp = time.Now().Format("20060102150405")
	oc.SeqNo = strings.ReplaceAll(uuid.New().String(), "-", "")
	oc.VersionNo = "2.00"
	oc.buildData(conf)
	oc.makeSign()

	ret, err := oc.SendReq(ChargeUrl, oc)
	if err != nil {
		return nil, err
	}

	ret, err = oc.RetData(ret)
	if err != nil {
		return nil, err
	}

	orderQueryReturn := new(OrderQueryReturn)

	err = json.Unmarshal(ret, &orderQueryReturn)

	if err != nil {
		return nil, err
	}

	if orderQueryReturn.SubCode != "0001" && orderQueryReturn.SubCode != "0000" {
		return nil, errors.New(orderQueryReturn.SubMsg)
	}

	return orderQueryReturn, nil
}
