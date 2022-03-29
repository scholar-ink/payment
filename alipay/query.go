/**
 * @author dengmeiyu
 * @since 20180716
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

type AliQuery struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	*BaseAliConfig
	*AliPayTradeQuery
}

type AliPayTradeQuery struct {
	OutTradeNo string `json:"out_trade_no,omitempty"` // 订单支付时传入的商户订单号, 与 TradeNo 二选一
	TradeNo    string `json:"trade_no,omitempty"`     // 支付宝交易号
}

type AliPayTradeQueryResponse struct {
	AliPayTradeQuery struct {
		Code             string `json:"code"`
		Msg              string `json:"msg"`
		SubCode          string `json:"sub_code"`
		SubMsg           string `json:"sub_msg"`
		AuthTradePayMode string `json:"auth_trade_pay_mode"` // 预授权支付模式，该参数仅在信用预授权支付场景下返回。信用预授权支付：CREDIT_PREAUTH_PAY
		BuyerLogonId     string `json:"buyer_logon_id"`      // 买家支付宝账号
		BuyerPayAmount   string `json:"buyer_pay_amount"`    // 买家实付金额，单位为元，两位小数。
		BuyerUserId      string `json:"buyer_user_id"`       // 买家在支付宝的用户id
		BuyerUserType    string `json:"buyer_user_type"`     // 买家用户类型。CORPORATE:企业用户；PRIVATE:个人用户。
		InvoiceAmount    string `json:"invoice_amount"`      // 交易中用户支付的可开具发票的金额，单位为元，两位小数。
		OutTradeNo       string `json:"out_trade_no"`        // 商家订单号
		PointAmount      string `json:"point_amount"`        // 积分支付的金额，单位为元，两位小数。
		ReceiptAmount    string `json:"receipt_amount"`      // 实收金额，单位为元，两位小数
		SendPayDate      string `json:"send_pay_date"`       // 本次交易打款给卖家的时间
		TotalAmount      string `json:"total_amount"`        // 交易的订单金额
		TradeNo          string `json:"trade_no"`            // 支付宝交易号
		TradeStatus      string `json:"trade_status"`        // 交易状态

		DiscountAmount      string `json:"discount_amount"`       // 平台优惠金额
		MdiscountAmount     string `json:"mdiscount_amount"`      // 商家优惠金额
		PayAmount           string `json:"pay_amount"`            // 支付币种订单金额
		PayCurrency         string `json:"pay_currency"`          // 订单支付币种
		SettleAmount        string `json:"settle_amount"`         // 结算币种订单金额
		SettleCurrency      string `json:"settle_currency"`       // 订单结算币种
		SettleTransRate     string `json:"settle_trans_rate"`     // 结算币种兑换标价币种汇率
		StoreId             string `json:"store_id"`              // 商户门店编号
		StoreName           string `json:"store_name"`            // 请求交易支付中的商户店铺的名称
		TerminalId          string `json:"terminal_id"`           // 商户机具终端编号
		TransCurrency       string `json:"trans_currency"`        // 标价币种
		TransPayRate        string `json:"trans_pay_rate"`        // 标价币种兑换支付币种汇率
		DiscountGoodsDetail string `json:"discount_goods_detail"` // 本次交易支付所使用的单品券优惠的商品优惠信息
		IndustrySepcDetail  string `json:"industry_sepc_detail"`  // 行业特殊信息（例如在医保卡支付业务中，向用户返回医疗信息）。
	} `json:"alipay_trade_query_response"`
	Sign string `json:"sign"`
}

func (tra *AliQuery) Handle(conf map[string]interface{}, privateKey string, aliPublicKey string) (*AliPayTradeQueryResponse, error) {

	err := tra.BuildData(conf)
	if err != nil {
		return nil, err
	}
	ret, err := tra.sendReq(ALITRADE, tra, privateKey)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(ret))
	return tra.RetData(ret, aliPublicKey)
}

func (tra *AliQuery) RetData(ret []byte, aliPublicKey string) (re *AliPayTradeQueryResponse, err error) {

	result := new(AliPayTradeQueryResponse)
	json.Unmarshal(ret, result)

	//b,_:=json.Marshal(result.AliPayTradeQuery)

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
		return nil, errors.New("延签参数转换失败")
	}

	err = helper.Sha1WithRsaVerify([]byte(content), sign, aliPublicKey)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	if result.AliPayTradeQuery.Code != "10000" && result.AliPayTradeQuery.Msg != "Success" {
		return nil, errors.New("支付宝支付查询失败：" + result.AliPayTradeQuery.SubMsg)
	}
	return result, nil

}

func (tra *AliQuery) sendReq(reqUrl string, pay interface{}, privateKey string) (b []byte, err error) {

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

func (tra *AliQuery) BuildData(conf map[string]interface{}) error {

	b, _ := json.Marshal(conf)

	countInfo := string(b)

	tra.BaseAliConfig.BizContent = countInfo

	return nil
}

func (tra *AliQuery) InitBaseConfig(config *BaseAliConfig) {

	config.TimeStamp = time.Now().Format("2006-01-02 15:04:05")
	config.Method = "alipay.trade.query"
	config.Charset = "utf-8"
	config.Version = "1.0"
	config.SignType = "RSA2"
	tra.BaseAliConfig = config
}
