package notify

import (
	"encoding/xml"
	"errors"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/weixin"
	"strings"
)

type WxNotifyData struct {
	Appid          string `xml:"appid" json:"appid"`
	BankType       string `xml:"bank_type" json:"bank_type"`
	CashFee        string `xml:"cash_fee" json:"cash_fee"`
	FeeType        string `xml:"fee_type" json:"fee_type"`
	IsSubscribe    string `xml:"is_subscribe" json:"is_subscribe"`
	MchId          string `xml:"mch_id" json:"mch_id"`
	NonceStr       string `xml:"nonce_str" json:"nonce_str"`
	Openid         string `xml:"openid" json:"openid"`
	OutTrade_no    string `xml:"out_trade_no" json:"out_trade_no"`
	ResultCode     string `xml:"result_code" json:"result_code"`
	ReturnCode     string `xml:"return_code" json:"return_code"`
	Sign           string `xml:"sign" json:"sign"`
	SubAppid       string `xml:"sub_appid" json:"sub_appid"`
	SubIsSubscribe string `xml:"sub_is_subscribe" json:"sub_is_subscribe"`
	SubMchId       string `xml:"sub_mch_id" json:"sub_mch_id"`
	SubOpenid      string `xml:"sub_openid" json:"sub_openid"`
	TimeEnd        string `xml:"time_end" json:"time_end"`
	TotalFee       string `xml:"total_fee" json:"total_fee"`
	TradeType      string `xml:"trade_type" json:"trade_type"`
	TransactionId  string `xml:"transaction_id" json:"transaction_id"`
}

type WxNotify struct {
	*weixin.BaseConfig
	notifyData *WxNotifyData
}

func (wx *WxNotify) InitBaseConfig(config *weixin.BaseConfig) {
	wx.BaseConfig = config
}

func (wx *WxNotify) getNotifyData(notifyXml string) error {

	if notifyXml == "" {
		return errors.New("获取通知数据失败")
	}

	notify := new(WxNotifyData)

	err := xml.Unmarshal([]byte(notifyXml), notify)

	if err != nil {
		return errors.New("获取通知数据失败:" + err.Error())
	}

	wx.notifyData = notify

	return nil
}

func (wx *WxNotify) checkNotify() error {
	if wx.notifyData.ResultCode != "SUCCESS" || wx.notifyData.ReturnCode != "SUCCESS" {
		return errors.New("微信返回失败")
	}

	return wx.verifySign()
}

func (wx *WxNotify) verifySign() error {

	mapData := helper.Struct2Map(wx.notifyData)

	signStr := helper.CreateLinkString(&mapData)

	switch wx.SignType {

	case "MD5":
		signStr += "&key=" + wx.Md5Key

		signStr = helper.Md5(signStr)
	}

	if wx.notifyData.Sign != strings.ToUpper(signStr) {
		return errors.New("返回数据验签失败，可能数据被篡改")
	}

	return nil
}

func (wx *WxNotify) replyNotify(flag bool, msg string) string {

	type ReturnResult struct {
		XMLName    xml.Name `xml:"xml" json:"-"`
		ReturnCode string   `xml:"return_code"`
		ReturnMsg  string   `xml:"return_msg"`
	}

	result := &ReturnResult{
		ReturnCode: "SUCCESS",
		ReturnMsg:  "OK",
	}
	if !flag {
		result = &ReturnResult{
			ReturnCode: "FAIL",
			ReturnMsg:  msg,
		}
	}

	b, err := xml.Marshal(result)

	if err != nil {
		return ""
	}

	return string(b)
}

type WxCallBack func(data *WxNotifyData) error

func (wx *WxNotify) Handle(xml string, callBack WxCallBack) string {
	err := wx.getNotifyData(xml)

	if err != nil {
		return wx.replyNotify(false, err.Error())
	}

	err = wx.checkNotify()

	if err != nil {
		return wx.replyNotify(false, err.Error())
	}

	err = callBack(wx.notifyData)

	if err != nil {
		return wx.replyNotify(false, err.Error())
	}

	return wx.replyNotify(true, "")
}
