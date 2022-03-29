package weixin

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/scholar-ink/payment/helper"
	"time"
)

type PubCharge struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	*PubConf
	BaseCharge
}

type PubConf struct {
	Openid         string `xml:"openid,omitempty" json:"openid,omitempty"`
	SubOpenid      string `xml:"sub_openid,omitempty" json:"sub_openid,omitempty"`
	Body           string `xml:"body" json:"body"`
	Detail         string `xml:"detail,omitempty" json:"detail,omitempty"`
	Attach         string `xml:"attach,omitempty" json:"attach,omitempty"`
	OutTradeNo     string `xml:"out_trade_no" json:"out_trade_no"`
	FeeType        string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`
	TotalFee       int64  `xml:"total_fee" json:"total_fee"`
	SpbillCreateIp string `xml:"spbill_create_ip" json:"spbill_create_ip"`
	GoodsTag       string `xml:"goods_tag,omitempty" json:"goods_tag,omitempty" `
	TradeType      string `xml:"trade_type" json:"trade_type"`
}

type PubReturn struct {
	AppId     string `json:"appId"`
	TimeStamp int64  `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}

func (app *PubCharge) Handle(conf map[string]interface{}) (interface{}, error) {
	err := app.BuildData(conf)

	if err != nil {
		return nil, err
	}
	app.SetSign(app)
	ret := app.SendReq(UnifiedorderReqUrl, app)
	return app.RetData(ret)
}

func (app *PubCharge) RetData(ret []byte) (pubReturn PubReturn, err error) {
	var result struct {
		Error
		Return
	}

	xml.Unmarshal(ret, &result)

	if result.ReturnCode == SUCCESS && result.ResultCode == SUCCESS {

		if result.SubAppId != "" {
			pubReturn.AppId = result.SubAppId
		} else {
			pubReturn.AppId = result.AppId
		}

		pubReturn.TimeStamp = time.Now().Unix()

		pubReturn.NonceStr = helper.NonceStr()

		pubReturn.Package = "prepay_id=" + result.PrepayId

		pubReturn.SignType = app.SignType

		app.SetSign(pubReturn)

		pubReturn.PaySign = app.Sign

	} else {

		return pubReturn, errors.New(result.ReturnMsg + result.ErrCodeDes)
	}

	return pubReturn, nil

}

func (app *PubCharge) BuildData(conf map[string]interface{}) error {

	b, _ := json.Marshal(conf)

	var pubConf PubConf

	json.Unmarshal(b, &pubConf)

	if pubConf.SpbillCreateIp == "" {
		pubConf.SpbillCreateIp = "127.0.0.1"
	}

	if pubConf.FeeType == "" {
		pubConf.FeeType = "CNY"
	}

	app.PubConf = &pubConf

	app.PubConf.TradeType = "JSAPI"

	return nil
}
