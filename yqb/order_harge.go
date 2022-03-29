package yqb

import (
	"encoding/json"
	"errors"
	"time"
)

type OrderCharge struct {
	BaseCharge
}

type BusinessExtendArea struct {
	TerminalId   string `json:"terminalId"`   //壹钱包的码URL
	TerminalType string `json:"terminalType"` //终端类型
}

type OrderChargeConf struct {
	PayScene            string              `json:"-"`                             //支付场景 01-微信02-支付宝
	MerTradeNo          string              `json:"merTradeNo"`                    //商户订单号
	CurrType            string              `json:"currType"`                      //币种，目前只支持： CNY(人民币)
	TotalAmount         string              `json:"totalAmount"`                   //订单总金额（分）
	Subject             string              `json:"subject"`                       //订单标题
	TradeTime           string              `json:"tradeTime"`                     //交易时间，格式："yyyy-MM-dd HH:mm:ss"
	PlatformMerchantId  string              `json:"platformMerchantId"`            //平安付聚合平台服务商号，平安付分配
	AggregateMerchantId string              `json:"aggregateMerchantId,omitempty"` //平安付聚合服务商，平安付分配号
	ChannelMerchantId   string              `json:"channelMerchantId,omitempty"`   //渠道商商户号
	TradeType           string              `json:"tradeType"`                     //交易类型 	01-公众号支付 02-扫码支付
	SubAppId            string              `json:"subAppid,omitempty"`            //子商户公众账号ID
	SubOpenId           string              `json:"subOpenid,omitempty"`           //用户子标识
	OrderDesc           string              `json:"orderDesc"`                     //订单描述
	OrderStartTime      string              `json:"orderStartTime"`                //交易起始时间
	OrderExpireTime     string              `json:"orderExpireTime"`               //交易结束时间
	NotifyUrl           string              `json:"notifyUrl"`                     //通知地址
	BusinessExtendArea  *BusinessExtendArea `json:"businessExtendArea"`            //受理终端位置信息
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
	OrigMerTradeNo     string `json:"origMerTradeNo,omitempty"`     //原商户订单号
	OrigTradeOrderNo   string `json:"origTradeOrderNo,omitempty"`   //原平安付订单号
	OrigChannelOrderNo string `json:"origChannelOrderNo,omitempty"` //原渠道商户订单号
	PlatformMerchantId string `json:"platformMerchantId"`           //平安付聚合平台服务商号，平安付分配
	MerRefundTradeNo   string `json:"merRefundTradeNo"`             //商户退款订单号
	RefundAmount       string `json:"refundAmount"`                 //退款金额
	RefundTransTime    string `json:"refundTransTime"`              //退款交易请求时间
}

type OrderChargeReturn struct {
	BusinessCode string      `json:"businessCode"`
	BusinessMsg  string      `json:"businessMsg"`
	CodeUrl      string      `json:"codeUrl"`
	WcPayData    interface{} `json:"wcPayData"`
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
	BusinessCode       string `json:"businessCode"`
	BusinessMsg        string `json:"businessMsg"`
	OrigMerTradeNo     string `json:"origMerTradeNo,omitempty"`   //原商户订单号
	OrigTradeOrderNo   string `json:"origTradeOrderNo,omitempty"` //原平安付订单号
	TradeOrderNo       string `json:"tradeOrderNo"`               //平安付退款订单号
	MerchantId         string `json:"merchantId"`                 //平安付商户号
	RefundAmount       string `json:"refundAmount"`               //退款金额
	MerRefundTradeNo   string `json:"merRefundTradeNo"`           //商户退款订单号
	PlatformMerchantId string `json:"platformMerchantId"`         //平安付聚合平台服务商号，平安付分配
	RefundSuccTime     string `json:"refundSuccTime"`             //退款成功时间
	ApplyRefundAmount  string `json:"applyRefundAmount"`          //申请退款金额
	FactRefundAmount   string `json:"factRefundAmount"`           //实际退款金额
}

func (oc *OrderCharge) Handle(conf *OrderChargeConf) (*OrderChargeReturn, error) {
	oc.BaseCharge.ServiceName = "aggregate.trade.placeOrder"
	oc.BaseCharge.Version = "V1.0"
	oc.BaseCharge.PayScene = conf.PayScene
	conf.CurrType = "CNY"
	conf.TradeTime = time.Now().Format("2006-01-02 15:04:05")
	conf.OrderStartTime = time.Now().Format("2006-01-02 15:04:05")
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

	if orderChargeReturn.BusinessCode != "00" {
		return nil, errors.New(orderChargeReturn.BusinessMsg)
	}

	return orderChargeReturn, nil
}

func (oc *OrderCharge) Query(conf *QueryConf) (*QueryReturn, error) {
	oc.BaseCharge.ServiceName = "aggregate.trade.query"
	oc.BaseCharge.Version = "V1.0"
	oc.BaseCharge.PayScene = "01"
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

func (oc *OrderCharge) Refund(conf *RefundConf) (*RefundReturn, error) {
	oc.BaseCharge.ServiceName = "aggregate.trade.refund"
	oc.BaseCharge.Version = "V1.0"
	oc.BaseCharge.PayScene = "01"
	ret, err := oc.SendReq(ChargeUrl, conf)
	if err != nil {
		return nil, err
	}
	ret, err = oc.RetData(ret)

	if err != nil {
		return nil, err
	}

	refundReturn := new(RefundReturn)
	err = json.Unmarshal(ret, &refundReturn)
	if err != nil {
		return nil, err
	}
	if refundReturn.BusinessCode != "00" {
		return nil, errors.New(refundReturn.BusinessMsg)
	}
	return refundReturn, nil
}

func (oc *OrderCharge) WxApply(conf *WxApplyConf) (*WxApplyReturn, error) {
	oc.BaseCharge.ServiceName = "aggregate.trade.bindWechatAppid"
	oc.BaseCharge.Version = "V1.0"
	oc.BaseCharge.PayScene = "01"
	ret, err := oc.SendReq(ChargeUrl, conf)
	if err != nil {
		return nil, err
	}
	ret, err = oc.RetData(ret)

	if err != nil {
		return nil, err
	}

	wxApplyReturn := new(WxApplyReturn)
	err = json.Unmarshal(ret, &wxApplyReturn)
	if err != nil {
		return nil, err
	}

	if wxApplyReturn.BusinessCode != "00" {
		return nil, errors.New(wxApplyReturn.BusinessMsg)
	}
	return wxApplyReturn, nil
}

func (oc *OrderCharge) RetData(ret []byte) (b []byte, err error) {

	type BaseReturn struct {
		Head struct {
			ResultCode string `json:"resultCode"`
			ResultMsg  string `json:"resultMsg"`
			SignType   string `json:"signType"`
		} `json:"head"`
		Body interface{} `json:"body"`
	}

	baseReturn := new(BaseReturn)

	err = json.Unmarshal(ret, &baseReturn)

	if err != nil {
		return
	}

	if baseReturn.Head.ResultCode != "000000" {
		err = errors.New(baseReturn.Head.ResultMsg)
		return
	}

	b, _ = json.Marshal(baseReturn.Body)

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
