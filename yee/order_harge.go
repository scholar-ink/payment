package yee

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/scholar-ink/payment/helper"
	"strings"
)

type OrderCharge struct {
	BaseCharge
}

type GoodsParamExt struct {
	GoodsName string `json:"goodsName"`
	GoodsDesc string `json:"GoodsDesc"`
}

type OrderChargeConf struct {
	MerchantNo        string         `json:"merchantNo"`              //子商户号
	OrderId           string         `json:"orderId"`                 //商户订单号，请自定义
	OrderAmount       string         `json:"orderAmount"`             //订单支付金额，单位元
	TimeoutExpress    string         `json:"timeoutExpress"`          //订单有效期
	NotifyUrl         string         `json:"notifyUrl"`               //支付成功服务器回调地址
	GoodsParamExtInfo string         `json:"goodsParamExt,omitempty"` //商品拓展信息
	GoodsParamExt     *GoodsParamExt `json:"-"`                       //商品拓展信息
	HMac              string         `json:"hmac"`
	HMacKey           string         `json:"-"`
	PayTool           string         `json:"-"` //支付工具
	PayType           string         `json:"-"` //支付类型
	AppId             string         `json:"-"` //支付类型
	OpenId            string         `json:"-"` //支付类型
}

type OrderRefundConf struct {
	MerchantNo      string `json:"merchantNo"`              //子商户号
	OrderId         string `json:"orderId"`                 //商户订单号
	RefundRequestId string `json:"refundRequestId"`         //退款订单号
	UniqueOrderNo   string `json:"uniqueOrderNo,omitempty"` //统一订单号
	RefundAmount    string `json:"refundAmount"`            //退款金额
	HMac            string `json:"hmac"`
	HMacKey         string `json:"-"`
}

type OrderToken struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type OrderChargeReturn struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	ResultData string `json:"resultData"`
}

type OrderRefundReturn struct {
	Code            string `json:"code"`
	Message         string `json:"message"`
	OrderId         string `json:"orderId"`
	RefundRequestId string `json:"refundRequestId"`
	UniqueRefundNo  string `json:"uniqueRefundNo"`
	RefundAmount    string `json:"refundAmount"`
}

func (oc *OrderCharge) Handle(conf *OrderChargeConf) (*OrderChargeReturn, error) {
	oc.Service = "/rest/v1.0/sys/trade/order"
	oc.RequestNo = strings.ReplaceAll(uuid.New().String(), "-", "")
	if conf.GoodsParamExt != nil {
		b, _ := json.Marshal(conf.GoodsParamExt)

		conf.GoodsParamExtInfo = string(b)
	}

	hMac := helper.HMacSha256([]byte(fmt.Sprintf("parentMerchantNo=%s&merchantNo=%s&orderId=%s&orderAmount=%s&notifyUrl=%s",
		oc.ParentMerchantNo, conf.MerchantNo, conf.OrderId, conf.OrderAmount, conf.NotifyUrl,
	)), conf.HMacKey)

	conf.HMac = hex.EncodeToString(hMac)

	ret, err := oc.SendReq(ChargeUrl, conf)

	if err != nil {
		return nil, err
	}

	ret, err = oc.retData(ret)

	if err != nil {
		return nil, err
	}

	orderToken := new(OrderToken)

	err = json.Unmarshal(ret, &orderToken)

	if err != nil {
		return nil, err
	}

	if orderToken.Code != "OPR00000" {
		return nil, errors.New(orderToken.Message)
	}

	return oc.apiOrder(orderToken.Token, conf.PayTool, conf.PayType, conf.AppId, conf.OpenId)
}
func (oc *OrderCharge) Refund(conf *OrderRefundConf) (string, error) {
	oc.Service = "/rest/v1.0/sys/trade/refund"
	oc.RequestNo = strings.ReplaceAll(uuid.New().String(), "-", "")

	hMac := helper.HMacSha256([]byte(fmt.Sprintf("parentMerchantNo=%s&merchantNo=%s&orderId=%s&uniqueOrderNo=%s&refundRequestId=%s&refundAmount=%s",
		oc.ParentMerchantNo, conf.MerchantNo, conf.OrderId, conf.UniqueOrderNo, conf.RefundRequestId, conf.RefundAmount,
	)), conf.HMacKey)

	conf.HMac = hex.EncodeToString(hMac)

	ret, err := oc.SendReq(ChargeUrl, conf)

	if err != nil {
		return "", err
	}

	ret, err = oc.retData(ret)

	if err != nil {
		return "", err
	}

	refundReturn := new(OrderRefundReturn)

	err = json.Unmarshal(ret, &refundReturn)

	if err != nil {
		return "", err
	}

	if refundReturn.Code != "OPR00000" {
		return "", errors.New(refundReturn.Message)
	}

	return refundReturn.UniqueRefundNo, nil
}

func (oc *OrderCharge) apiOrder(token, payTool, payType, appId, openId string) (*OrderChargeReturn, error) {
	oc.Service = "/rest/v1.0/nccashierapi/api/pay"
	oc.RequestNo = strings.ReplaceAll(uuid.New().String(), "-", "")

	params := &struct {
		Token       string `json:"token"`            //订单 token
		PayTool     string `json:"payTool"`          //支付工具
		PayType     string `json:"payType"`          //支付类型
		AppId       string `json:"appId,omitempty"`  //商家公众号ID
		OpenId      string `json:"openId,omitempty"` //用户OPENID
		UserNo      string `json:"userNo"`           //用户标识
		UserType    string `json:"userType"`         //用户标识类型
		UserIp      string `json:"userIp"`
		Version     string `json:"version"`
		ExtParamMap string `json:"extParamMap"`
	}{
		Token:       token,
		PayTool:     payTool,
		PayType:     payType,
		AppId:       appId,
		OpenId:      openId,
		UserNo:      strings.ReplaceAll(uuid.New().String(), "-", ""),
		UserType:    "IMEI",
		UserIp:      "127.0.0.1",
		Version:     "1.0",
		ExtParamMap: `{"reportFee":"XIANXIA"}`,
	}

	ret, err := oc.SendReq(ChargeUrl, params)

	if err != nil {
		return nil, err
	}
	ret, err = oc.retData(ret)

	if err != nil {
		return nil, err
	}

	orderReturn := new(OrderChargeReturn)
	err = json.Unmarshal(ret, &orderReturn)
	if err != nil {
		return nil, err
	}

	if orderReturn.Code != "CAS00000" {
		return nil, errors.New(orderReturn.Message)
	}

	return orderReturn, nil
}
