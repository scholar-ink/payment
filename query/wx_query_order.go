/**
 * @author dengmeiyu
 * @since 20180618
 */
package query

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"github.com/scholar-ink/payment/helper"
	"io/ioutil"
	"strings"
	"time"
)

const (
	OrderQueryReqUrl = "https://api.mch.weixin.qq.com/pay/orderquery"
	SUCCESS          = "SUCCESS"
)

type OrderQuery struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	*BaseConfig
	*OrderOutTradeNo
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
type QueryReturn struct {
	//DeviceInfo string `xml:"device_info" json:"device_info"`
	OpenId         string `xml:"open_id" json:"open_id"`
	SubIsSubScribe string `xml:"sub_is_sub_scribe" json:"sub_is_sub_scribe"`
	TradeType      string `xml:"trade_type" json:"trade_type"`
	BankType       string `xml:"bank_type" json:"bank_type"`
	TotalFee       int    `xml:"total_fee" json:"total_fee"`
	CashFee        int    `xml:"cash_fee" json:"cash_fee"`
	OutTradeNo     string `xml:"out_trade_no" json:"out_trade_no"`
	TransactionId  string `xml:"transaction_id" json:"transaction_id"`
	TimeEnd        string `xml:"time_end" json:"time_end"`
	TradeStateDesc string `xml:"trade_state_desc" json:"trade_state_desc"`
}

type QueryResult struct {
	Error
	Return
	QueryReturn
}

type OrderOutTradeNo struct {
	OutTradeNo string `xml:"out_trade_no" json:"out_trade_no"`
}

/**
 *
 * 查询订单情况
 * @param string $out_trade_no  商户订单号
 * @param int $succCode         查询订单结果
 * @return 0 订单不成功，1表示订单成功，2表示继续等待
 */
func (que *OrderQuery) Query(data map[string]interface{}, succCode int) (interface{}, int) {

	queryResult, _ := que.OrderQuery(data)

	if queryResult.ReturnCode == "SUCCESS" && queryResult.ResultCode == "SUCCESS" {
		//支付成功
		if queryResult.TradeState == "SUCCESS" {
			succCode = 1
			return queryResult, succCode
		} else if queryResult.TradeState == "USERPAYING" { //用户支付中

			succCode = 2
			return false, succCode
		}
	}

	//如果返回错误码为“此交易订单号不存在”则直接认定失败
	if queryResult.ErrCode == "ORDERNOTEXIST" {
		succCode = 0
	} else {
		//如果是系统错误，则后续继续
		succCode = 2
	}
	return false, succCode
}

func (que *OrderQuery) OrderQuery(data map[string]interface{}) (QueryResult, error) {
	err := que.BuildData(data)
	if err != nil {
		return QueryResult{}, err
	}
	que.SetSign(que)
	ret := que.SendReq(OrderQueryReqUrl, que)
	return que.RetData(ret)
}

func (que *OrderQuery) RetData(ret []byte) (QueryResult, error) {

	result := QueryResult{}
	xml.Unmarshal(ret, &result)
	return result, nil

}

func (que *OrderQuery) BuildData(conf map[string]interface{}) error {

	var OrderOutTradeNo OrderOutTradeNo
	b, _ := json.Marshal(conf)
	json.Unmarshal(b, &OrderOutTradeNo)

	que.OrderOutTradeNo = &OrderOutTradeNo

	return nil
}

func (que *OrderQuery) InitBaseConfig(config *BaseConfig) {

	config.NonceStr = helper.NonceStr()

	que.BaseConfig = config
}

func (que *OrderQuery) SendReq(reqUrl string, pay interface{}) (b []byte) {

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

func (que *OrderQuery) SetSign(pay interface{}) {

	mapData := helper.Struct2Map(pay)

	signStr := helper.CreateLinkString(&mapData)

	que.Sign = que.makeSign(signStr)
}

func (que *OrderQuery) makeSign(sign string) string {

	switch que.SignType {

	case "MD5":
		sign += "&key=" + que.Md5Key

		sign = helper.Md5(sign)
	}

	return strings.ToUpper(sign)
}
