package jttp

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/scholar-ink/payment/helper"
)

type OrderCharge struct {
	*OrderChargeConf
	BaseCharge
}

type OrderChargeConf struct {
	OutTradeNo  string `json:"outTradeNo"`  //商户订单号
	TotalAmount string `json:"totalAmount"` //商户订单金额
	Currency    string `json:"currency"`    //币种
	Body        string `json:"body"`        //交易或商品描述
	NotifyUrl   string `json:"notifyUrl"`   //接受支付结果的通知地址
	OrgCreateIp string `json:"orgCreateIp"` //订单生成的机器 IP
	OrderTime   string `json:"orderTime"`   //商户订单提交时间
}

type OrderChargeReturn struct {
	OutTradeNo  string `json:"outTradeNo"`
	RetCode     string `json:"retCode"`
	RetMsg      string `json:"retMsg"`
	TotalAmount string `json:"totalAmount"`
	PayUrl      string `json:"payUrl"`
	QrCode      string `json:"qrCode"`
	TradeNo     string `json:"tradeNo"`
}

func (oc *OrderCharge) Handle(conf *OrderChargeConf) (*OrderChargeReturn, error) {
	oc.Version = "1.0"
	oc.TxnType = "T20601"
	err := oc.BuildData(conf)
	if err != nil {
		return nil, err
	}
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

	if orderChargeReturn.RetCode != "RC0031" {
		return nil, errors.New(orderChargeReturn.RetMsg)
	}

	return orderChargeReturn, nil
}

func (oc *OrderCharge) RetData(ret []byte) ([]byte, error) {

	type BaseReturn struct {
		BizContext string `json:"bizContext"`
		Msg        string `json:"msg"`
		Code       string `json:"code"`
		Sign       string `json:"sign"`
	}

	baseReturn := new(BaseReturn)

	json.Unmarshal(ret, &baseReturn)

	if baseReturn.Code != "SUCCESS" {
		return nil, errors.New(baseReturn.Msg)
	}

	b, _ := base64.StdEncoding.DecodeString(baseReturn.BizContext)

	aesKey, _ := base64.StdEncoding.DecodeString(oc.AesKey)

	origin, err := helper.AesDecrypt(b, aesKey) // ECB解密

	if err != nil {
		return nil, err
	}

	origin = helper.PKCS7UnPadding(origin)

	return origin, nil
}

func (oc *OrderCharge) BuildData(conf *OrderChargeConf) error {

	oc.OrderChargeConf = conf
	return nil
}
