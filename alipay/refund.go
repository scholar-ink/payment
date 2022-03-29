/**
 * @author dengmeiyu
 * @since 20180713
 */
package alipay

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"io/ioutil"
	"net/url"
	"strings"
	"time"
)

type AliRefund struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	*BaseAliConfig
	*AliRefundConf
}

type AliRefundConf struct {
	OutTradeNo     string `xml:"out_trade_no" json:"out_trade_no"`
	TradeNo        string `xml:"trade_no" json:"trade_no"`
	RefundAmount   string `xml:"refund_amount" json:"refund_amount"`
	RefundCurrency string `xml:"refund_currency" json:"refund_currency"`
	RefundReason   string `xml:"refund_reason,omitempty" json:"refund_reason,omitempty"`
}

type AliPayTradeRefundResponse struct {
	AliPayTradeRefund struct {
		Code         string `json:"code"`
		Msg          string `json:"msg"`
		SubCode      string `json:"sub_code"`
		SubMsg       string `json:"sub_msg"`
		TradeNo      string `json:"trade_no"`       // 支付宝交易号
		OutTradeNo   string `json:"out_trade_no"`   // 商户订单号
		BuyerLogonId string `json:"buyer_logon_id"` // 用户的登录id
		BuyerUserId  string `json:"buyer_user_id"`  // 买家在支付宝的用户id
		FundChange   string `json:"fund_change"`    // 本次退款是否发生了资金变化
		RefundFee    string `json:"refund_fee"`     // 退款总金额
		GmtRefundPay string `json:"gmt_refund_pay"` // 退款支付时间
		StoreName    string `json:"store_name"`     // 交易在支付时候的门店名称
	} `json:"alipay_trade_refund_response"`
	Sign string `json:"sign"`
}

func (tra *AliRefund) Handle(conf map[string]interface{}, privateKey string, aliPublicKey string) (string, error) {

	err := tra.BuildData(conf)
	if err != nil {
		return "", err
	}
	ret, err := tra.sendReq(ALITRADE, tra, privateKey)
	if err != nil {
		return "", err
	}
	fmt.Println(string(ret))
	return tra.RetData(ret, aliPublicKey)
}

func (tra *AliRefund) RetData(ret []byte, aliPublicKey string) (re string, err error) {

	result := new(AliPayTradeRefundResponse)
	json.Unmarshal(ret, result)

	//b,_:=json.Marshal(result.AliPayTradeRefund)

	//TODO 调用之后验签 验签方法还要改一下
	dataStr := string(ret)
	var rootNodeName = strings.Replace(tra.Method, ".", "_", -1) + k_RESPONSE_SUFFIX

	var rootIndex = strings.LastIndex(dataStr, rootNodeName)
	var errorIndex = strings.LastIndex(dataStr, k_ERROR_RESPONSE)

	var content string
	var sign string

	if rootIndex > 0 {
		content, sign = tra.ParserJSONSource(dataStr, rootNodeName, rootIndex)
	} else if errorIndex > 0 {
		content, sign = tra.ParserJSONSource(dataStr, k_ERROR_RESPONSE, errorIndex)
	} else {
		return "", errors.New("延签参数转换失败")
	}

	err = helper.Sha1WithRsaVerify([]byte(content), sign, aliPublicKey)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	if result.AliPayTradeRefund.Code != "10000" && result.AliPayTradeRefund.Msg != "Success" {
		return "", errors.New("支付宝退款失败：" + result.AliPayTradeRefund.SubMsg)
	}
	return result.AliPayTradeRefund.TradeNo, nil

}

func (tra *AliRefund) sendReq(reqUrl string, pay interface{}, privateKey string) (b []byte, err error) {

	client := helper.NewHttpClient()
	var data = url.Values{}

	data.Add("app_id", tra.AppId)
	data.Add("method", tra.Method)
	data.Add("charset", tra.Charset)
	data.Add("sign_type", tra.SignType)
	data.Add("timestamp", tra.TimeStamp)
	data.Add("version", tra.Version)
	data.Add("biz_content", tra.BizContent)
	data.Add("app_auth_token", tra.AppAuthToken)
	data.Add("sign", tra.RSASign(data, privateKey))

	httpResp, err := client.PostForm(reqUrl, data)

	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	b, err = ioutil.ReadAll(httpResp.Body)

	return

}

func (tra *AliRefund) BuildData(conf map[string]interface{}) error {

	b, _ := json.Marshal(conf)

	countInfo := string(b)

	tra.BaseAliConfig.BizContent = countInfo

	return nil
}

func (tra *AliRefund) InitBaseConfig(config *BaseAliConfig) {

	config.TimeStamp = time.Now().Format("2006-01-02 15:04:05")
	config.Method = "alipay.trade.refund"
	config.Charset = "utf-8"
	config.Version = "1.0"
	config.SignType = "RSA2"
	tra.BaseAliConfig = config
}
