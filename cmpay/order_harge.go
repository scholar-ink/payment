package cmpay

import (
	"encoding/json"
	"errors"
)

type OrderCharge struct {
	*OrderChargeConf
	BaseCharge
}

type OrderChargeConf struct {
	ShopId         string `json:"shopId"`              //商户id
	OrderId        string `json:"orderId"`             //订单号
	Money          string `json:"money"`               //金额
	PayType        string `json:"type,omitempty"`      //支付类型 微信：weixin；支付宝：aipay
	OrderType      string `json:"orderType,omitempty"` //支付类型 微信：weixin；支付宝：aipay
	GoodsMsg       string `json:"goodsMsg"`            //商品信息
	RedirectUrl    string `json:"redirectUrl"`         //回调地址
	RedirectNumber string `json:"redirectNumber"`      //回调次数
	ReturnUrl      string `json:"returnUrl,omitempty"` //完成跳转链接
	SubAppid       string `json:"subAppid,omitempty"`  //用户openid
	SubOpenid      string `json:"subOpenid,omitempty"` //用户openid
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
	PayQrCodePayResponse string        `json:"cmpay_qrcode_pay_response"`
	PayQrCodePay         *PayQrCodePay `json:"-"`
}

func (oc *OrderCharge) Handle(conf *OrderChargeConf) (string, error) {

	oc.Service = "buypay.do"

	err := oc.BuildData(conf)
	if err != nil {
		return "", err
	}
	return oc.GetReq(ChargeUrl, oc), nil
}

func (oc *OrderCharge) Handle2(conf *OrderChargeConf) (*OrderChargeReturn, error) {

	oc.Service = "jsnativepay.do"

	ret, err := oc.SendReq(ChargeUrl, conf)
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

	return orderChargeReturn, nil
}

func (oc *OrderCharge) RetData(ret []byte) ([]byte, error) {

	type BaseReturn struct {
		RespCode string `json:"resp_code"`
		RespMsg  string `json:"resp_msg"`
	}

	baseReturn := new(BaseReturn)

	err := json.Unmarshal(ret, &baseReturn)

	if err != nil {
		return nil, err
	}

	if baseReturn.RespCode != "000000" {
		return nil, errors.New(baseReturn.RespMsg)
	}
	return ret, nil
}

func (oc *OrderCharge) BuildData(conf *OrderChargeConf) error {
	oc.OrderChargeConf = conf
	return nil
}
