package gpay

import (
	"encoding/json"
	"errors"
	"github.com/scholar-ink/payment/helper"
	"time"
)

type ConsumeOrderCharge struct {
	*ConsumeOrderChargeConf
	BaseCharge
}

type ConsumeOrderChargeConf struct {
	PaymentType   string `json:"payment_type"`             //支付方式
	NonceStr      string `json:"nonce_str"`                //随机字符串
	OutTradeNo    string `json:"out_trade_no"`             //商户订单号
	AuthCode      string `json:"auth_code"`                //扫码支付授权码
	Subject       string `json:"subject"`                  //订单标题
	TotalAmount   int64  `json:"total_amount,omitempty"`   //订单总金额，单位到分
	TotalFee      int64  `json:"total_fee,omitempty"`      //订单总金额，单位到分
	TransCurrency string `json:"trans_currency,omitempty"` //币种
	FeeType       string `json:"fee_type,omitempty"`       //币种
	MchCreateIp   string `json:"mch_create_ip"`            //订单提交终端IP
	TimeStart     string `json:"time_start"`               //订单生成时间格式：yyyyMMddHHmmss
	TimeExpire    string `json:"time_expire,omitempty"`    //订单失效时间格式：yyyyMMddHHmmss
	NotifyUrl     string `json:"notify_url"`               //订单通知URL
	CallbackUrl   string `json:"callback_url,omitempty"`   //订单回调URL
	DeviceInfo    string `json:"device_info,omitempty"`    //设备号
	Body          string `json:"body"`                     //商品描述
	LimitPay      string `json:"limit_pay,omitempty"`      //支付方式限制，如限制不能用信用卡支付等1-限定不能使用信用卡
	Ext           *Ext   `json:"ext,omitempty"`            //扩展字段，类型为map格式json串
}

type ConsumeOrderChargeReturn struct {
	ResultCode  string `json:"result_code"`
	ErrorCode   string `json:"err_code"`
	ErrCodeDesc string `json:"err_code_des"`
	OutTradeNo  string `json:"out_trade_no"`
	OrderNo     string `json:"order_no"`
	TradeType   string `json:"trade_type"`
	TradeState  string `json:"trade_state"` //交易状态 NOTPAY 未支付 SUCCESS支付成功
	FeeType     string `json:"fee_type"`
	TotalFee    string `json:"total_fee"`
	Attach      string `json:"attach"`
	TimeEnd     string `json:"time_end"` //支付成功时间
}

func (oc *ConsumeOrderCharge) Handle(conf *ConsumeOrderChargeConf) (*ConsumeOrderChargeReturn, error) {
	err := oc.BuildData(conf)
	if err != nil {
		return nil, err
	}
	oc.SetSign()
	ret, err := oc.SendReq(ChargeUrl)
	if err != nil {
		return nil, err
	}
	return oc.RetData(ret)
}

func (oc *ConsumeOrderCharge) RetData(ret []byte) (*ConsumeOrderChargeReturn, error) {

	ret, err := oc.BaseCharge.RetData(ret)

	if err != nil {
		return nil, err
	}

	orderChargeReturn := new(ConsumeOrderChargeReturn)

	err = json.Unmarshal(ret, &orderChargeReturn)

	if err != nil {
		return nil, err
	}

	if orderChargeReturn.ResultCode != "0" {
		return nil, errors.New(orderChargeReturn.ErrCodeDesc)
	}

	return orderChargeReturn, nil
}

func (oc *ConsumeOrderCharge) BuildData(conf *ConsumeOrderChargeConf) error {

	if conf.PaymentType == "" {
		return errors.New("支付方式不能为空")
	}
	if conf.OutTradeNo == "" {
		return errors.New("支付方式不能为空")
	}
	if len(conf.OutTradeNo) > 32 {
		return errors.New("商户订单号不能超过32")
	}
	if conf.PaymentType == "pay-wx-consume" {
		if conf.FeeType == "" {
			return errors.New("币种不能为空")
		}
		if conf.TotalFee <= 0 {
			return errors.New("订单金额需大于0")
		}
	} else if conf.PaymentType == "pay-zfb-consume" {
		if conf.TransCurrency == "" {
			return errors.New("币种不能为空")
		}
		if conf.TotalAmount <= 0 {
			return errors.New("订单金额需大于0")
		}
	}
	if conf.MchCreateIp == "" {
		return errors.New("终端IP不能为空")
	}
	if conf.NotifyUrl == "" {
		return errors.New("订单通知URL不能为空")
	}
	if conf.Subject == "" {
		return errors.New("订单标题不能为空")
	}
	if conf.Body == "" {
		return errors.New("商品描述不能为空")
	}

	conf.TimeStart = time.Now().Format("20060102150405")
	conf.NonceStr = helper.NonceStr()

	if conf.TimeExpire == "" {
		conf.TimeExpire = time.Now().Add(30 * time.Minute).Format("20060102150405")
	}

	b, err := json.Marshal(conf)

	if err != nil {
		return err
	}

	oc.ConsumeOrderChargeConf = conf

	oc.SourceData = string(b)

	oc.ServiceName = "consume.order"

	return nil
}
