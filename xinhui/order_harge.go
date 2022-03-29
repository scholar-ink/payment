package xinhui

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"strings"
)

type OrderCharge struct {
	*OrderChargeConf
	BaseCharge
}

type OrderChargeConf struct {
	MerchantNo     string `json:"merchant_no"`           //商户编号
	AppId          string `json:"appid"`                 //用户标识
	OpenId         string `json:"openid"`                //用户标识
	ChannelId      string `json:"channel_id"`            //渠道商商户号
	SubMchId       string `json:"sub_mch_id"`            //报备商户编号
	ChannelType    string `json:"channel_type"`          //支付通道
	TrxType        string `json:"trx_type"`              //交易类型
	OutTradeNo     string `json:"out_trade_no"`          //商户订单号
	TotalFee       string `json:"total_fee"`             //总金额
	FeeType        string `json:"fee_type"`              //费率类型
	TrxFee         string `json:"trx_fee,omitempty"`     //订单维度手续费
	DeviceInfo     string `json:"device_info"`           //终端设备号
	Body           string `json:"body"`                  //商品描述
	SpbillCreateIp string `json:"spbill_create_ip"`      //终端 IP
	NotifyUrl      string `json:"notify_url"`            //通知地址
	TimeExpire     string `json:"time_expire,omitempty"` //交易结束时间
}

type OrderChargeReturn struct {
	Version   string `json:"version"`
	RspCode   string `json:"rsp_code"`
	RspMsg    string `json:"rsp_msg"`
	Sign      string `json:"sign"`
	WcPayData string `json:"wc_pay_data"`
	AliQrCode string `json:"ali_qr_code"`
}

func (oc *OrderCharge) Handle(conf *OrderChargeConf) (*OrderChargeReturn, error) {
	oc.Version = "1.0"
	oc.ServiceType = "xh.uni.trx.create"
	oc.RequestId = strings.ReplaceAll(uuid.New().String(), "-", "")
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

	if orderChargeReturn.RspCode != "0000" {
		return nil, errors.New(orderChargeReturn.RspMsg)
	}

	return orderChargeReturn, nil
}

func (oc *OrderCharge) BuildData(conf *OrderChargeConf) error {

	oc.OrderChargeConf = conf
	return nil
}
