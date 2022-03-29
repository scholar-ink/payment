package adapay

import (
	"encoding/json"
	"errors"
)

type OrderCharge struct {
	BaseCharge
}

type OrderChargeConf struct {
	AppId      string                 `json:"app_id"`                //渠道
	OrderNo    string                 `json:"order_no"`              //订单号
	PayChannel string                 `json:"pay_channel"`           //订单金额
	TimeExpire string                 `json:"time_expire,omitempty"` //通知地址
	PayAmt     string                 `json:"pay_amt"`               //订单标题
	GoodsTitle string                 `json:"goods_title"`           //通知地址
	GoodsDesc  string                 `json:"goods_desc"`            //通知地址
	Currency   string                 `json:"currency"`              //通知地址
	NotifyUrl  string                 `json:"notify_url,omitempty"`  //通知地址
	Expend     map[string]interface{} `json:"expend,omitempty"`      //通知地址
	SignType   string                 `json:"sign_type"`
}

type OrderRefundConf struct {
	PaymentId     string `json:"payment_id"`      //原交易支付对象ID
	RefundOrderNo string `json:"refund_order_no"` //退款订单号
	RefundAmt     string `json:"refund_amt"`      //退款金额
	Reason        string `json:"reason"`          //退款描述
	SignType      string `json:"sign_type"`
}

type WxApplyConf struct {
	PayScene           string `json:"-"`                  //支付场景 01:服务商微信支付
	SubAppId           string `json:"subAppid"`           //APPID
	PlatformMerchantId string `json:"platformMerchantId"` //平安付聚合平台服务商号，平安付分配
}

type QueryConf struct {
	MerTradeNo         string `json:"merTradeNo,omitempty"`   //商户订单号
	TradeOrderNo       string `json:"tradeOrderNo,omitempty"` //平安付订单号
	PlatformMerchantId string `json:"platformMerchantId"`     //平安付聚合平台服务商号，平安付分配
}

type RefundConf struct {
	PaymentId     string `json:"payment_id"`       //原交易支付对象ID
	RefundOrderNo string `json:"refund_order_no"`  //退款订单号
	RefundAmt     string `json:"refund_amt"`       //退款金额
	Reason        string `json:"reason,omitempty"` //退款描述
	SignType      string `json:"sign_type"`
}

type OrderChargeReturn struct {
	Expend struct {
		PayInfo   string `json:"pay_info"`
		QrCodeUrl string `json:"qrcode_url"`
	} `json:"expend"`
}

type OrderRefundReturn struct {
	Id            string `json:"id"`
	RefundOrderNo string `json:"refund_order_no"`
	PaymentId     string `json:"payment_id"`
	RefundAmt     string `json:"refund_amt"`
	TransState    string `json:"trans_state"`
}

type WxApplyReturn struct {
	BusinessCode string `json:"businessCode"`
	BusinessMsg  string `json:"businessMsg"`
}

type QueryReturn struct {
	BusinessCode       string `json:"businessCode"`
	BusinessMsg        string `json:"businessMsg"`
	TradeOrderNo       string `json:"tradeOrderNo"`       //平安付订单号
	MerTradeNo         string `json:"merTradeNo"`         //商户订单号
	TotalAmount        string `json:"totalAmount"`        //订单总金额（分）
	PayTime            string `json:"payTime"`            //支付时间
	PlatformMerchantId string `json:"platformMerchantId"` //平安付聚合平台服务商号，平安付分配
	MerchantId         string `json:"merchantId"`         //平安付商户号
	ChannelOrderNo     string `json:"channelOrderNo"`     //平安付商户号
}

type RefundReturn struct {
	Id             string `json:"id"`               //平台退款订单号
	RefundOrderNo  string `json:"refund_order_no"`  //商户退款订单号
	PaymentId      string `json:"payment_id"`       //原平台付订单号
	PaymentOrderNo string `json:"payment_order_no"` //原商户订单号
	RefundAmt      string `json:"refund_amt"`       //退款金额
	SucceedTime    string `json:"succeed_time"`     //退款成功时间
}

func (oc *OrderCharge) Handle(conf *OrderChargeConf) (*OrderChargeReturn, error) {
	oc.Service = "payments"
	conf.SignType = "RSA2"
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

func (oc *OrderCharge) Refund(conf *OrderRefundConf) (*OrderRefundReturn, error) {
	oc.Service = "payments/" + conf.PaymentId + "/refunds"
	conf.SignType = "RSA2"
	ret, err := oc.SendReq(ChargeUrl, conf)
	if err != nil {
		return nil, err
	}
	ret, err = oc.RetData(ret)

	if err != nil {
		return nil, err
	}

	orderRefundReturn := new(OrderRefundReturn)
	err = json.Unmarshal(ret, &orderRefundReturn)
	if err != nil {
		return nil, err
	}

	return orderRefundReturn, nil
}

func (oc *OrderCharge) Query(conf *QueryConf) (*QueryReturn, error) {
	oc.Service = "payments"
	ret, err := oc.SendReq(ChargeUrl, conf)
	if err != nil {
		return nil, err
	}
	ret, err = oc.RetData(ret)

	if err != nil {
		return nil, err
	}

	queryReturn := new(QueryReturn)
	err = json.Unmarshal(ret, &queryReturn)
	if err != nil {
		return nil, err
	}
	if queryReturn.BusinessCode != "00" {
		return nil, errors.New(queryReturn.BusinessMsg)
	}
	return queryReturn, nil
}

func (oc *OrderCharge) RetData(ret []byte) (b []byte, err error) {

	type BaseResponse struct {
		Data string `json:"data"`
	}

	baseResponse := new(BaseResponse)

	err = json.Unmarshal(ret, &baseResponse)

	if err != nil {
		return
	}

	type BaseReturn struct {
		ErrorCode string `json:"error_code"`
		ErrorMsg  string `json:"error_msg"`
	}

	baseReturn := new(BaseReturn)

	b = []byte(baseResponse.Data)

	err = json.Unmarshal(b, &baseReturn)

	if err != nil {
		return
	}

	if baseReturn.ErrorCode != "" {
		err = errors.New(baseReturn.ErrorMsg)
		return
	}

	return
}

//
//func (oc *OrderCharge) GetAuthCodeUrl(conf *GetAuthCodeUrlConf) string {
//	mapData := helper.Struct2Map(conf)
//	mapData["app_code"] = oc.AppCode
//	oc.makeSign(mapData)
//	mapData["sign"] = oc.Sign
//
//	values := maps.Map2Values(&mapData)
//
//	return fmt.Sprintf("%s/tool/v1/get_weixin_oauth_code?%s", ChargeUrl, values.Encode())
//}
//
//func (oc *OrderCharge) GetOpenId(conf *GetOpenIdConf) (string, error) {
//	oc.Service = "/tool/v1/get_weixin_openid"
//	ret, err := oc.SendReq(ChargeUrl, conf)
//	if err != nil {
//		return "", err
//	}
//	ret, err = oc.RetData(ret)
//
//	if err != nil {
//		return "", err
//	}
//
//	getOpenIdReturn := new(GetOpenIdReturn)
//	err = json.Unmarshal(ret, &getOpenIdReturn)
//	if err != nil {
//		return "", err
//	}
//	return getOpenIdReturn.Openid, nil
//}
