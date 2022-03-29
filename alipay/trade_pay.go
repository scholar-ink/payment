/**
 * @author dengmeiyu
 * @since 20180711
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

type AliPayTradePayResponse struct {
	AliPayTradePay AliPayTradePay `json:"alipay_trade_pay_response"`
	Sign           string         `json:"sign"`
}
type AliPayTradePay struct {
	Code                string `json:"code"`
	Msg                 string `json:"msg"`
	SubCode             string `json:"sub_code"`
	SubMsg              string `json:"sub_msg"`
	BuyerLogonId        string `json:"buyer_logon_id"`        // 买家支付宝账号
	BuyerPayAmount      string `json:"buyer_pay_amount"`      // 买家实付金额，单位为元，两位小数。
	BuyerUserId         string `json:"buyer_user_id"`         // 买家在支付宝的用户id
	CardBalance         string `json:"card_balance"`          // 支付宝卡余额
	DiscountGoodsDetail string `json:"discount_goods_detail"` // 本次交易支付所使用的单品券优惠的商品优惠信息
	GmtPayment          string `json:"gmt_payment"`
	InvoiceAmount       string `json:"invoice_amount"` // 交易中用户支付的可开具发票的金额，单位为元，两位小数。
	OutTradeNo          string `json:"out_trade_no"`   // 创建交易传入的商户订单号
	TradeNo             string `json:"trade_no"`       // 支付宝交易号
	PointAmount         string `json:"point_amount"`   // 积分支付的金额，单位为元，两位小数。
	ReceiptAmount       string `json:"receipt_amount"` // 实收金额，单位为元，两位小数
	StoreName           string `json:"store_name"`     // 发生支付交易的商户门店名称
	TotalAmount         string `json:"total_amount"`   // 发该笔退款所对应的交易的订单金额
}
type AliConf struct {
	OutTradeNo         string `xml:"out_trade_no" json:"out_trade_no"`
	Scene              string `xml:"scene" json:"scene"`
	AuthCode           string `xml:"auth_code" json:"auth_code"`
	ProductCode        string `xml:"product_code,omitempty" json:"product_code,omitempty"`
	Subject            string `xml:"subject" json:"subject"`
	BuyerId            string `xml:"buyer_id" json:"buyer_id"`
	SellerId           string `xml:"seller_id" json:"seller_id"`
	TotalAmount        string `xml:"total_amount" json:"total_amount"`
	TransCurrency      string `xml:"trans_currency,omitempty" json:"trans_currency,omitempty"`
	SettleCurrency     string `xml:"settle_currency,omitempty" json:"settle_currency,omitempty"`
	DiscountableAmount string `xml:"discountable_amount,omitempty" json:"discountable_amount,omitempty"`
	Body               string `xml:"body" json:"body"`
}

type AliTrade struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	*BaseAliConfig
	*AliConf
}

func (tra *AliTrade) Handle(conf map[string]interface{}, privateKey string, aliPublicKey string) (*AliPayTradePay, error) {

	err := tra.BuildData(conf)
	if err != nil {
		return nil, err
	}
	ret, err := tra.sendReq(ALITRADE, tra, privateKey)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(ret))
	return tra.RetData(ret, privateKey, aliPublicKey)
}

func (tra *AliTrade) RetData(ret []byte, privateKey string, aliPublicKey string) (re *AliPayTradePay, err error) {

	result := new(AliPayTradePayResponse)
	json.Unmarshal(ret, result)

	//b, _ := json.Marshal(result.AliPayTradePay)

	//调用之后验签
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
		return nil, errors.New("延签参数转换失败")
	}

	err = helper.Sha1WithRsaVerify([]byte(content), sign, aliPublicKey)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	var aliTradeRes *AliPayTradePay
	aliTradeRes = &result.AliPayTradePay

	//支付后做查询
	if result.AliPayTradePay.Code != "10000" && result.AliPayTradePay.Msg != "Success" {
		if result.AliPayTradePay.Code == "10003" {

			app := new(AliQuery)

			app.InitBaseConfig(&BaseAliConfig{
				AppId:        tra.AppId,
				AppAuthToken: tra.AppAuthToken,
			})

			ret, err := app.Handle(map[string]interface{}{
				"out_trade_no": result.AliPayTradePay.OutTradeNo,
			}, privateKey, aliPublicKey)

			if err != nil {
				return aliTradeRes, errors.New("调用支付宝查询接口失败：" + err.Error())
			}

			if ret.AliPayTradeQuery.TradeStatus == "TRADE_SUCCESS" {
				return aliTradeRes, nil
			} else {
				return aliTradeRes, errors.New("支付宝条码支付失败：" + ret.AliPayTradeQuery.TradeStatus)
			}
		}
		return aliTradeRes, errors.New("支付宝条码支付失败：" + result.AliPayTradePay.Msg + result.AliPayTradePay.SubMsg)
	}
	return aliTradeRes, nil

}

func (tra *AliTrade) sendReq(reqUrl string, pay interface{}, privateKey string) (b []byte, err error) {

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

func (tra *AliTrade) BuildData(conf map[string]interface{}) error {

	b, _ := json.Marshal(conf)

	countInfo := string(b)

	tra.BaseAliConfig.BizContent = countInfo

	return nil
}

func (tra *AliTrade) InitBaseConfig(config *BaseAliConfig) {

	config.TimeStamp = time.Now().Format("2006-01-02 15:04:05")
	config.Method = "alipay.trade.pay"
	config.Charset = "utf-8"
	config.Version = "1.0"
	config.SignType = "RSA2"
	tra.BaseAliConfig = config
}
