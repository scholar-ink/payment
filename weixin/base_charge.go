package weixin

import (
	"bytes"
	"encoding/xml"
	"github.com/scholar-ink/payment/helper"
	"io/ioutil"
	"strings"
	"time"
)

const (
	UnifiedorderReqUrl = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	MicropayReqUrl     = "https://api.mch.weixin.qq.com/pay/micropay"
	SUCCESS            = "SUCCESS"
)

type Error struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	ResultCode string `xml:"result_code"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`
	PrepayId   string `xml:"prepay_id"`
	CodeUrl    string `xml:"code_url"`
	Recall     string `xml:"recall"`
}

type Return struct {
	AppId      string `xml:"appid"`
	SubAppId   string `xml:"sub_appid"`
	MchId      string `xml:"mch_id"`
	DeviceInfo string `xml:"device_info"`
	NonceStr   string `xml:"nonce_str"`
	TradeState string `xml:"trade_state"`
}

type BaseCharge struct {
	*BaseConfig
}

type BaseConfig struct {
	AppId          string        `xml:"appid" json:"appid"`
	MchId          string        `xml:"mch_id" json:"mch_id"`
	SubAppId       string        `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`
	SubMchId       string        `xml:"sub_mch_id,omitempty" json:"sub_mch_id,omitempty"`
	TimeStart      string        `xml:"time_start,omitempty" json:"time_start,omitempty"`
	TimeExpire     string        `xml:"time_expire,omitempty" json:"time_expire,omitempty"`
	NotifyUrl      string        `xml:"notify_url,omitempty" json:"notify_url,omitempty"`
	Md5Key         string        `xml:"-" json:"-"`
	SignType       string        `xml:"sign_type,omitempty" json:"sign_type,omitempty"`
	Sign           string        `xml:"sign" json:"sign"`
	NonceStr       string        `xml:"nonce_str" json:"nonce_str"`
	ExpireDuration time.Duration `xml:"-" json:"-"`
	Cert           string        `xml:"-" json:"-"`
	Key            string        `xml:"-" json:"-"`
}

func (base *BaseCharge) SendReq(reqUrl string, pay interface{}) (b []byte) {

	buffer := bytes.NewBuffer(b)

	err := xml.NewEncoder(buffer).Encode(pay)

	if err != nil {
		return nil
	}

	client := helper.NewHttpClient()

	httpResp, err := client.Post(reqUrl, "text/xml; charset=utf-8", buffer)

	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	b, err = ioutil.ReadAll(httpResp.Body)

	return
}

func (base *BaseCharge) SetSign(pay interface{}) {

	mapData := helper.Struct2Map(pay)

	signStr := helper.CreateLinkString(&mapData)

	base.Sign = base.makeSign(signStr)
}

func (base *BaseCharge) makeSign(sign string) string {

	switch base.SignType {

	case "MD5":
		sign += "&key=" + base.Md5Key

		sign = helper.Md5(sign)
	}

	return strings.ToUpper(sign)
}

func (base *BaseCharge) InitBaseConfig(config *BaseConfig) {
	config.NonceStr = helper.NonceStr()

	if config.ExpireDuration != 0 {
		config.TimeExpire = time.Now().Add(config.ExpireDuration).Format("20060102150405")
	} else {
		config.TimeExpire = ""
	}
	config.TimeStart = time.Now().Format("20060102150405")
	base.BaseConfig = config
}
