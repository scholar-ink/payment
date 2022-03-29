package hangpay

import (
	"encoding/json"
	"errors"
)

type OrderCharge struct {
	BaseCharge
}

type OrderChargeConf struct {
	Money         string `json:"money"`             //交易金额
	SubMerchantNo string `json:"subMerchantNo"`     //子商户号
	OrderNo       string `json:"orderNo"`           //订单号
	PayType       string `json:"payType"`           //支付类型 1.支付宝交易 2.微信交易3.云闪付
	Remark        string `json:"remark"`            //备注
	TradeTime     string `json:"tradeTime"`         //交易时间
	SucPage       string `json:"sucPage,omitempty"` //成功地址
	NotifyUrl     string `json:"notifyUrl"`         //通知地址
}

type OrderChargeReturn struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	OrderNo string `json:"orderNo"`
	QrCode  string `json:"qrCode"`
}

type RefundConf struct {
	RefundAmt     string `json:"refundAmt"`     //退款金额
	OrderNo       string `json:"orderNo"`       //交易订单号
	NotifyUrl     string `json:"notifyUrl"`     //通知地址
	RefundOrderNo string `json:"refundOrderNo"` //退款订单号
}

type RefundReturn struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func (oc *OrderCharge) Handle(conf *OrderChargeConf) (string, error) {
	oc.Version = "1.0.0"
	if conf.PayType == "1" {
		oc.Service = "api/alipay/launch"
	} else if conf.PayType == "2" {
		oc.Service = "api/weChat/launch"
	} else {
		return "", errors.New("暂不支持该支付方式")
	}
	ret, err := oc.SendReq(ChargeUrl, conf)
	if err != nil {
		return "", err
	}

	orderChargeReturn := new(OrderChargeReturn)

	err = json.Unmarshal(ret, &orderChargeReturn)

	if err != nil {
		return "", err
	}

	if orderChargeReturn.Code != "SUCCESS" {
		return "", errors.New(orderChargeReturn.Msg)
	}

	return orderChargeReturn.QrCode, nil
}

func (oc *OrderCharge) Refund(conf *RefundConf) error {
	oc.Version = "1.0.0"
	oc.Service = "api/refund"
	ret, err := oc.SendReq(ChargeUrl, conf)
	if err != nil {
		return err
	}

	refundReturn := new(RefundReturn)

	err = json.Unmarshal(ret, &refundReturn)

	if err != nil {
		return err
	}

	if refundReturn.Code != "SUCCESS" {
		return errors.New(refundReturn.Msg)
	}

	return nil
}
