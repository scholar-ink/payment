package qf

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/map"
	"time"
)

type OrderCharge struct {
	BaseCharge
}

type OrderChargeConf struct {
	MchId      string `json:"mchid"`        //子商户号
	PayType    string `json:"pay_type"`     //支付类型
	OutTradeNo string `json:"out_trade_no"` //外部订单号
	TxAmt      string `json:"txamt"`        //订单支付金额，单位分
	TxCurrCd   string `json:"txcurrcd"`     //币种
	TxDtm      string `json:"txdtm"`        //请求交易时间格式为：YYYY-MM-DD HH:MM:SS
	GoodsName  string `json:"goods_name"`   //商品名称或标示
	SubOpenid  string `json:"sub_openid"`   //微信的openid
}

type GetAuthCodeUrlConf struct {
	MchId       string `json:"mchid"`        //子商户号
	RedirectUri string `json:"redirect_uri"` //回跳url
}

type GetOpenIdConf struct {
	MchId string `json:"mchid"` //子商户号
	Code  string `json:"code"`  //微信auth code
}

type OrderChargeReturn struct {
	QrCode    string      `json:"qrcode"`
	PayParams interface{} `json:"pay_params"`
}

type GetOpenIdReturn struct {
	Openid string `json:"openid"`
}

func (oc *OrderCharge) Handle(conf *OrderChargeConf) (*OrderChargeReturn, error) {
	oc.Service = "/trade/v1/payment"
	conf.TxCurrCd = "CNY"
	conf.TxDtm = time.Now().Format("2006-01-02 15:04:05")
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

func (oc *OrderCharge) RetData(ret []byte) (b []byte, err error) {

	type BaseReturn struct {
		RespErr string `json:"resperr"`
		RespCd  string `json:"respcd"`
		RespMsg string `json:"respmsg"`
		SysSn   string `json:"syssn"`
	}

	baseReturn := new(BaseReturn)

	err = json.Unmarshal(ret, &baseReturn)

	if err != nil {
		return
	}

	if baseReturn.RespCd != "0000" {
		err = errors.New(baseReturn.RespErr)
		return
	}

	return ret, nil
}

func (oc *OrderCharge) GetAuthCodeUrl(conf *GetAuthCodeUrlConf) string {
	mapData := helper.Struct2Map(conf)
	mapData["app_code"] = oc.AppCode
	oc.makeSign(mapData)
	mapData["sign"] = oc.Sign

	values := maps.Map2Values(&mapData)

	return fmt.Sprintf("%s/tool/v1/get_weixin_oauth_code?%s", ChargeUrl, values.Encode())
}

func (oc *OrderCharge) GetOpenId(conf *GetOpenIdConf) (string, error) {
	oc.Service = "/tool/v1/get_weixin_openid"
	ret, err := oc.SendReq(ChargeUrl, conf)
	if err != nil {
		return "", err
	}
	ret, err = oc.RetData(ret)

	if err != nil {
		return "", err
	}

	getOpenIdReturn := new(GetOpenIdReturn)
	err = json.Unmarshal(ret, &getOpenIdReturn)
	if err != nil {
		return "", err
	}
	return getOpenIdReturn.Openid, nil
}
