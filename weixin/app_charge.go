package weixin

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/scholar-ink/payment/helper"
	"time"
)

type AppCharge struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	*AppConf
	BaseCharge
}

type AppConf struct {
	Openid         string `xml:"openid,omitempty" json:"openid,omitempty"`
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

type AppReturn struct {
	AppId     string `json:"appId"`
	TimeStamp int64  `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}

func (app *AppCharge) Handle(conf map[string]interface{}) (interface{}, error) {
	err := app.BuildData(conf)
	if err != nil {
		return nil, err
	}
	app.SetSign(app)
	ret := app.SendReq(UnifiedorderReqUrl, app)
	return app.RetData(ret)
}

func (app *AppCharge) RetData(ret []byte) (appReturn AppReturn, err error) {
	var result struct {
		Error
		Return
	}

	xml.Unmarshal(ret, &result)

	if result.ReturnCode == SUCCESS && result.ResultCode == SUCCESS {

		if result.SubAppId != "" {
			appReturn.AppId = result.SubAppId
		} else {
			appReturn.AppId = result.AppId
		}

		appReturn.TimeStamp = time.Now().Unix()

		appReturn.NonceStr = helper.NonceStr()

		appReturn.Package = "prepay_id=" + result.PrepayId

		appReturn.SignType = app.SignType

		app.SetSign(appReturn)

		appReturn.PaySign = app.Sign

	} else {

		return appReturn, errors.New(result.ErrCodeDes)
	}

	return appReturn, nil

}

func (app *AppCharge) BuildData(conf map[string]interface{}) error {

	b, _ := json.Marshal(conf)

	var appConf AppConf

	json.Unmarshal(b, &appConf)

	app.AppConf = &appConf

	app.AppConf.TradeType = "APP"

	return nil
}
