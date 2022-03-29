package refund

import (
	"encoding/json"
	"errors"
	"github.com/scholar-ink/payment/gpay"
	"github.com/scholar-ink/payment/helper"
)

type ZgRefund struct {
	*ZgRefundConf
	gpay.BaseCharge
}

type ZgRefundConf struct {
	NonceStr      string `json:"nonce_str"`              //随机字符串
	OrderNo       string `json:"order_no,omitempty"`     //银通平台订单号
	OutTradeNo    string `json:"out_trade_no,omitempty"` //商户订单号
	OutRefundNo   string `json:"out_refund_no"`          //商户退款订单号(唯一编号且长度小于等于32)
	TotalFee      int64  `json:"total_fee"`              //订单金额，单位到分
	RefundFee     int64  `json:"refund_fee"`             //退款金额，单位到分
	RefundChannel string `json:"refund_channel"`         //退款方式
}

type ZgRefundReturn struct {
	ResultCode      string `json:"result_code"`
	ErrCode         string `json:"err_code"`
	ErrCodeDes      string `json:"err_code_des"`
	OutTradeNo      string `json:"out_trade_no"`                //商户订单号(唯一编号且长度小于等于32)
	OrderNo         string `json:"order_no"`                    //银通平台订单号
	OutRefundNo     string `json:"out_refund_no"`               //商户退款订单号
	RefundNo        string `json:"refund_no"`                   //银通平台退款订单号
	RefundStatus    int32  `json:"refund_status"`               //退款状态
	TotalFee        int64  `json:"total_fee"`                   //订单金额，单位到分
	RefundFee       int64  `json:"refund_fee"`                  //退款金额，单位到分
	CouponRefundFee string `json:"coupon_refund_fee,omitempty"` //现金券退款金额，单位到分
	RefundChannel   string `json:"refund_channel"`              //退款方式
	Ext             string `json:"ext"`
}

func (oc *ZgRefund) Handle(conf *ZgRefundConf) (refundNo string, er error) {
	err := oc.BuildData(conf)
	if err != nil {
		return "", err
	}
	oc.SetSign()
	ret, err := oc.SendReq(gpay.ChargeUrl)
	if err != nil {
		return "", err
	}
	return oc.RetData(ret)
}

func (oc *ZgRefund) RetData(ret []byte) (string, error) {

	ret, err := oc.BaseCharge.RetData(ret)

	if err != nil {
		return "", err
	}

	orderChargeReturn := new(ZgRefundReturn)

	err = json.Unmarshal(ret, &orderChargeReturn)

	if err != nil {
		return "", err
	}

	if orderChargeReturn.ResultCode != "0" {
		return "", errors.New(orderChargeReturn.ErrCodeDes)
	}

	return orderChargeReturn.RefundNo, nil
}

func (oc *ZgRefund) BuildData(conf *ZgRefundConf) error {

	conf.NonceStr = helper.NonceStr()

	conf.RefundChannel = "ORIGINAL"

	b, err := json.Marshal(conf)

	if err != nil {
		return err
	}

	oc.ZgRefundConf = conf

	oc.SourceData = string(b)

	oc.ServiceName = "refund.pay"

	return nil
}
