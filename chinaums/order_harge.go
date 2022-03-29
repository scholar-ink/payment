package chinaums

import (
	"encoding/json"
	"errors"
)

type OrderCharge struct {
	*OrderChargeConf
	BaseCharge
}

type OrderChargeConf struct {
	BillNo      string `json:"billNo,omitempty"`   //账单号
	BillDate    string `json:"billDate,omitempty"` //账单日期
	BillDesc    string `json:"billDesc"`           //账单描述
	TotalAmount string `json:"totalAmount"`        //支付总金额
	NotifyUrl   string `json:"notifyUrl"`          //通知地址
	ReturnUrl   string `json:"returnUrl"`          //通知地址
}

type OrderChargeReturn struct {
	ErrCode    string `json:"errCode"`
	ErrMsg     string `json:"errMsg"`
	BillQRCode string `json:"billQRCode"` //二维码
}
type OrderQueryConf struct {
	BillDate string `json:"billDate"` //账单日期
	BillNo   string `json:"billNo"`
}
type OrderQueryReturn struct {
	ErrCode     string `json:"errCode"`
	ErrMsg      string `json:"errMsg"`
	BillNo      string `json:"billNo"`
	BillDate    string `json:"billDate"`
	CreateTime  string `json:"createTime"`
	BillStatus  string `json:"billStatus"`
	TotalAmount int32  `json:"totalAmount"`
}
type OrderRefundConf struct {
	BillDate      string `json:"billDate"` //账单日期
	BillNo        string `json:"billNo"`
	RefundOrderId string `json:"refundOrderId"`
	RefundAmount  string `json:"refundAmount"`
}
type OrderRefundReturn struct {
	ErrCode             string `json:"errCode"`
	ErrMsg              string `json:"errMsg"`
	MerOrderId          string `json:"merOrderId"`
	RefundOrderId       string `json:"refundOrderId"`
	RefundTargetOrderId string `json:"refundTargetOrderId"`
}

func (oc *OrderCharge) Handle(conf *OrderChargeConf) (string, error) {
	oc.MsgType = "bills.getQRCode"

	ret, err := oc.SendReq(ChargeUrl, conf)
	if err != nil {
		return "", err
	}

	orderChargeReturn := new(OrderChargeReturn)

	err = json.Unmarshal(ret, &orderChargeReturn)

	if err != nil {
		return "", err
	}

	if orderChargeReturn.ErrCode != "SUCCESS" {
		return "", errors.New(orderChargeReturn.ErrMsg)
	}

	return orderChargeReturn.BillQRCode, nil
}
func (oc *OrderCharge) Refund(conf *OrderRefundConf) (string, error) {
	oc.MsgType = "bills.refund"

	ret, err := oc.SendReq(ChargeUrl, conf)
	if err != nil {
		return "", err
	}

	orderRefundReturn := new(OrderRefundReturn)

	err = json.Unmarshal(ret, &orderRefundReturn)

	if err != nil {
		return "", err
	}

	if orderRefundReturn.ErrCode != "SUCCESS" {
		return "", errors.New(orderRefundReturn.ErrMsg)
	}

	return orderRefundReturn.RefundTargetOrderId, nil
}
func (oc *OrderCharge) Query(conf *OrderQueryConf) (*OrderQueryReturn, error) {
	oc.MsgType = "bills.query"

	ret, err := oc.SendReq(ChargeUrl, conf)
	if err != nil {
		return nil, err
	}

	orderQueryReturn := new(OrderQueryReturn)

	err = json.Unmarshal(ret, &orderQueryReturn)

	if err != nil {
		return nil, err
	}

	if orderQueryReturn.ErrCode != "SUCCESS" {
		return nil, errors.New(orderQueryReturn.ErrMsg)
	}

	return orderQueryReturn, nil
}
