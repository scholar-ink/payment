package minpay

import (
	"encoding/json"
	"errors"
)

type OrderCharge struct {
	*OrderChargeConf
	BaseCharge
}

type OrderChargeConf struct {
	MerchantRequestId string `json:"merchantRequestId"`      //商户请求号
	OrderAmount       string `json:"orderAmount"`            //实际支付订单金额(元)
	GoodsName         string `json:"goodsName"`              //商品名称
	GoodsDesc         string `json:"goodsDesc"`              //商品名称
	PayProduct        string `json:"payProduct"`             //支付产品
	CreateTime        string `json:"createTime"`             //订单请求时间yyyyMMddHHmmss
	ValidDate         string `json:"validDate,omitempty"`    //订单有效期
	BuyerLogonId      string `json:"buyerLogonId,omitempty"` //买家支付宝账号
	NonceStr          string `json:"nonceStr,omitempty"`     //买家支付宝账号
	UserIp            string `json:"userIp,omitempty"`       //用户 ip
	MerchantIp        string `json:"merchantIp,omitempty"`   //商户 ip
	CallBackUrl       string `json:"callBackUrl"`            //通知地址
}

type OrderChargeReturn struct {
	ReturnCode  string `json:"returnCode"`
	ReturnMsg   string `json:"returnMessage"`
	AgentNo     string `json:"agent_no"`     //商户订单号
	MerchantNo  string `json:"merchant_no"`  //系统单号
	OutOrderNo  string `json:"out_order_no"` //交易状态
	OrderNo     string `json:"order_no"`     //支付渠道
	ProductNo   string `json:"product_no"`   //支付类型
	TotalAmount int32  `json:"total_amount"` //二维码链接
	PayType     string `json:"pay_type"`     //虚户账号
	TradeStatus int32  `json:"trade_status"` //付款账号
	QrCode      string `json:"qrcode"`
}

func (oc *OrderCharge) Handle(conf *OrderChargeConf) (string, error) {
	oc.ServiceName = "zxsb.code.uscan"
	oc.Version = "V1.1"

	ret, err := oc.SendReq(ChargeUrl, conf)
	if err != nil {
		return "", err
	}

	ret, err = oc.RetData(ret)
	if err != nil {
		return "", err
	}

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
