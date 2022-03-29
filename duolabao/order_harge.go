package duolabao

import (
	"encoding/json"
)

type OrderCharge struct {
	BaseCharge
}

type OrderChargeConf struct {
	AgentNum    string `json:"agentNum,omitempty"`
	CustomerNum string `json:"customerNum"`
	ShopNum     string `json:"shopNum"`
	BankType    string `json:"bankType,omitempty"` //支付方式
	RequestNum  string `json:"requestNum"`         //流水号
	Amount      string `json:"amount"`             //订单金额
	Source      string `json:"source,omitempty"`   //API
	CallbackUrl string `json:"callbackUrl"`        //订单通知URL
	AuthId      string `json:"authId,omitempty"`   //操作员ID,微信公众号支付时,值为微信用户openid。 支付宝服务窗支付时，值为买家支付宝用户ID
}

type OrderChargeReturn struct {
	ResultCode  string `json:"result_code"`
	ErrorCode   string `json:"err_code"`
	ErrCodeDesc string `json:"err_code_des"`
	OutTradeNo  string `json:"out_trade_no"`
	OrderNo     string `json:"order_no"`
	CodeUrl     string `json:"code_url"`
	CodeImgUrl  string `json:"code_img_url"`
	Url         string `json:"url"`
}

func (oc *OrderCharge) Handle(conf *OrderChargeConf) (*OrderChargeReturn, error) {

	if conf.AgentNum != "" {
		if conf.BankType == "WX_XCX" {
			oc.Path = "/v1/agent/order/pay/create"
		} else {
			oc.Path = "/v1/agent/order/payurl/create"
		}
	} else {
		if conf.BankType == "WX_XCX" {
			oc.Path = "/v1/customer/order/pay/create"
		} else {
			oc.Path = "/v1/customer/order/payurl/create"
		}
	}

	ret, err := oc.sendReq(ChargeUrl, conf)
	if err != nil {
		return nil, err
	}

	ret, err = oc.RetData(ret)
	if err != nil {
		return nil, err
	}

	orderReturn := new(OrderChargeReturn)
	err = json.Unmarshal(ret, &orderReturn)
	if err != nil {
		return nil, err
	}
	return orderReturn, nil
}
