package sumpay

import (
	"encoding/base64"
	"encoding/json"
	"errors"
)

type OrderCharge struct {
	*OrderChargeConf
	BaseCharge
}

type OrderChargeConf struct {
	MerNo        string `json:"mer_no"`               //商户代码
	SubMerNo     string `json:"sub_mer_no,omitempty"` //二级商户代码
	UserId       string `json:"user_id"`              //付方用户
	BusinessCode string `json:"business_code"`        //支付代码
	TradeCode    string `json:"trade_code,omitempty"` //交易码
	GoodsName    string `json:"goods_name"`           //商品名称
	GoodsNum     string `json:"goods_num"`            //商品数量
	GoodsType    string `json:"goods_type"`           //商品类型
	OrderNo      string `json:"order_no"`             //订单号
	OrderTime    string `json:"order_time"`           //订单时间
	OrderAmount  string `json:"order_amount"`         //订单金额
	NeedNotify   string `json:"need_notify"`          //是否需要商户后台通知
	NotifyUrl    string `json:"notify_url,omitempty"` //商户后台通知URL
}

type PayQrCodePay struct {
	OrderNo string `json:"order_no"`
	CodeUrl string `json:"code_url"`
}

type OrderChargeReturn struct {
	RespCode             string        `json:"resp_code"`
	RespMsg              string        `json:"resp_msg"`
	SignType             string        `json:"sign_type"`
	Sign                 string        `json:"sign"`
	PayQrCodePayResponse string        `json:"sumpay_qrcode_pay_response"`
	PayQrCodePay         *PayQrCodePay `json:"-"`
}

func (oc *OrderCharge) Handle(conf *OrderChargeConf) (*OrderChargeReturn, error) {
	oc.Version = "1.0"

	oc.Service = "fosun.sumpay.api.pay.qrcode.trade.apply"

	if conf.BusinessCode == "09" {
		oc.Service = "fosun.sumpay.api.pay.new.qrcode.trade.apply"
	}

	err := oc.BuildData(conf)
	if err != nil {
		return nil, err
	}
	ret, err := oc.SendReq(ChargeUrl, oc)
	if err != nil {
		return nil, err
	}
	return oc.RetData(ret)
}

func (oc *OrderCharge) RetData(ret []byte) (*OrderChargeReturn, error) {

	orderChargeReturn := new(OrderChargeReturn)

	err := json.Unmarshal(ret, &orderChargeReturn)

	if err != nil {
		return nil, err
	}

	if orderChargeReturn.RespCode != "000000" {
		return nil, errors.New(orderChargeReturn.RespMsg)
	}

	json.Unmarshal([]byte(orderChargeReturn.PayQrCodePayResponse), &orderChargeReturn.PayQrCodePay)

	return orderChargeReturn, nil
}

func (oc *OrderCharge) BuildData(conf *OrderChargeConf) error {

	conf.GoodsName = base64.StdEncoding.EncodeToString([]byte(conf.GoodsName))

	conf.NotifyUrl = base64.StdEncoding.EncodeToString([]byte(conf.NotifyUrl))

	oc.OrderChargeConf = conf
	return nil
}
