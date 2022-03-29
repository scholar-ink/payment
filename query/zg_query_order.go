package query

import (
	"encoding/json"
	"errors"
	"github.com/scholar-ink/payment/gpay"
	"github.com/scholar-ink/payment/helper"
)

type ZgOrderQuery struct {
	*ZgOrderQueryConf
	gpay.BaseCharge
}

type ZgOrderQueryConf struct {
	NonceStr   string `json:"nonce_str"`    //随机字符串
	OutTradeNo string `json:"out_trade_no"` //商户订单号
}

type ZgOrderQueryReturn struct {
	ResultCode     string `json:"result_code"`
	ErrCode        string `json:"err_code"`
	ErrCodeDes     string `json:"err_code_des"`
	PayStatus      string `json:"pay_status"`                 //支付状态码，此字段值最终作为支付是否成功的结果
	PaymentType    string `json:"payment_type"`               //支付方式代码
	OutTradeNo     string `json:"out_trade_no"`               //商户订单号(唯一编号且长度小于等于32)
	OrderNo        string `json:"order_no"`                   //银通平台订单号
	FeeType        string `json:"fee_type"`                   //币种，默认传值：CNY
	TotalFee       string `json:"total_fee"`                  //订单金额，单位到分
	PaidFee        string `json:"paid_fee"`                   //支付金额，单位到分，支付成功时，需校验支付金额与订单金额是否一致
	CouponFee      string `json:"coupon_fee,omitempty"`       //现金券支付金额，单位到分
	Attach         string `json:"attach,omitempty"`           //商品附加信息
	TimeEnd        string `json:"time_end"`                   //支付完成时间
	Openid         string `json:"openid,omitempty"`           //用户标识,用户在服务商 appid 下的唯一标识
	IsSubscribe    string `json:"is_subscribe"`               //用户是否关注子公众账号，Y-关注，N-未关注
	BankType       string `json:"bank_type,omitempty"`        //付款银行类型
	TransactionId  string `json:"transaction_id,omitempty"`   //收单机构号(微信或支付宝交易号)
	BankBillNo     string `json:"bank_bill_no,omitempty"`     //付款银行订单号，若为微信支付则为空
	SubAppid       string `json:"sub_appid,omitempty"`        //商户公众号 appid
	SubIsSubscribe string `json:"sub_is_subscribe,omitempty"` //是否关注子商户公众账号，Y-关注，N-未关注
	SubOpenid      string `json:"sub_openid,omitempty"`       //用户在商户公众号 appid 下的唯一标识
	Ext            string `json:"ext"`
}

func (oc *ZgOrderQuery) Handle(conf *ZgOrderQueryConf) (*ZgOrderQueryReturn, error) {
	err := oc.BuildData(conf)
	if err != nil {
		return nil, err
	}
	oc.SetSign()
	ret, err := oc.SendReq(gpay.ChargeUrl)
	if err != nil {
		return nil, err
	}

	return oc.RetData(ret)
}

func (oc *ZgOrderQuery) RetData(ret []byte) (*ZgOrderQueryReturn, error) {

	ret, err := oc.BaseCharge.RetData(ret)

	if err != nil {
		return nil, err
	}

	orderChargeReturn := new(ZgOrderQueryReturn)

	err = json.Unmarshal(ret, &orderChargeReturn)

	if err != nil {
		return nil, err
	}

	if orderChargeReturn.ResultCode != "0" && orderChargeReturn.ErrCode != "" {
		return nil, errors.New(orderChargeReturn.ErrCodeDes)
	}

	return orderChargeReturn, nil
}

func (oc *ZgOrderQuery) BuildData(conf *ZgOrderQueryConf) error {

	conf.NonceStr = helper.NonceStr()

	b, err := json.Marshal(conf)

	if err != nil {
		return err
	}

	oc.ZgOrderQueryConf = conf

	oc.SourceData = string(b)

	oc.ServiceName = "query.pay"

	return nil
}
