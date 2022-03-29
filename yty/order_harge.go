package yty

import (
	"encoding/json"
	"errors"
	"fmt"
)

type OrderCharge struct {
	*OrderChargeConf
	BaseCharge
}

type OrderChargeConf struct {
	OutOrderNo  string `json:"out_order_no"` //商户系统订单号
	TotalAmount string `json:"total_amount"` //交易金额
	PayType     string `json:"pay_type"`     //支付类型
	OrderBody   string `json:"order_body"`   //交易金额
	PayTime     string `json:"pay_time"`     //交易结束时间
	NotifyUrl   string `json:"notify_url"`   //通知地址
}

type OrderChargeReturn struct {
	ReturnCode  string `json:"return_code"`
	ReturnMsg   string `json:"return_msg"`
	AgentNo     string `json:"agent_no"`     //商户订单号
	MerchantNo  string `json:"merchant_no"`  //系统单号
	OutOrderNo  string `json:"out_order_no"` //交易状态
	OrderNo     string `json:"order_no"`     //支付渠道
	ProductNo   string `json:"product_no"`   //支付类型
	TotalAmount int32  `json:"total_amount"` //二维码链接
	PayType     string `json:"pay_type"`     //虚户账号
	TradeStatus int32  `json:"trade_status"` //付款账号
	QrCode      string `json:"qrcode"`       //小程序支付、公众号支付：当status为02时返回
}

func (oc *OrderCharge) Handle(conf *OrderChargeConf) (string, error) {
	oc.Service = "qrpay"
	oc.ProductNo = "P3002"

	ret, err := oc.SendReq(ChargeUrl, conf)
	if err != nil {
		return "", err
	}

	ret, err = oc.RetData(ret)
	if err != nil {
		return "", err
	}

	fmt.Println(string(ret))

	orderChargeReturn := new(OrderChargeReturn)

	err = json.Unmarshal(ret, &orderChargeReturn)

	if err != nil {
		return "", err
	}

	if orderChargeReturn.ReturnCode != "0000" {
		return "", errors.New(orderChargeReturn.ReturnMsg)
	}

	return orderChargeReturn.QrCode, nil
}
