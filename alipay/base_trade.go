/**
 * @author dengmeiyu
 * @since 20180713
 */
package alipay

import (
	"github.com/scholar-ink/payment/helper"
	"net/url"
	"sort"
	"strings"
)

const (
	ALITRADE = "https://openapi.alipay.com/gateway.do?charset=utf-8"
	SUCCESS  = "SUCCESS"
	FAIL     = "FAIL"

	k_RESPONSE_SUFFIX = "_response"
	k_ERROR_RESPONSE  = "error_response"
	k_SIGN_NODE_NAME  = "sign"
)

type BaseAliConfig struct {
	AppId        string `xml:"app_id" json:"app_id"`
	Method       string `xml:"method" json:"method"`
	SignType     string `xml:"sign_type" json:"sign_type"`
	Sign         string `xml:"sign" json:"sign"`
	TimeStamp    string `xml:"time_stamp" json:"time_stamp"`
	Version      string `xml:"version" json:"version"`
	BizContent   string `xml:"-" json:"-"`
	Charset      string `xml:"-" json:"-"`
	NotifyUrl    string `xml:"-" json:"-"`
	AppAuthToken string `xml:"app_auth_token" json:"app_auth_token"`
}

//RSA256签名
func (base *BaseAliConfig) RSASign(m url.Values, privateKey string) string {
	//对url.values进行排序
	if m == nil {
		m = make(url.Values, 0)
	}

	var pList = make([]string, 0, 0)
	for key := range m {
		var value = strings.TrimSpace(m.Get(key))
		if len(value) > 0 {
			pList = append(pList, key+"="+value)
		}
	}
	sort.Strings(pList)
	var src = strings.Join(pList, "&")
	//对排序后的数据进行rsa2加密，获得sign
	encodeStr, _ := helper.Rsa2Encrypt([]byte(src), privateKey)

	return encodeStr
}

//获取支付返回的需要验签的内容
func (base *BaseAliConfig) ParserJSONSource(rawData string, nodeName string, nodeIndex int) (content string, sign string) {
	var dataStartIndex = nodeIndex + len(nodeName) + 2
	var signIndex = strings.LastIndex(rawData, "\""+k_SIGN_NODE_NAME+"\"")
	var dataEndIndex = signIndex - 1

	var indexLen = dataEndIndex - dataStartIndex
	if indexLen < 0 {
		return "", ""
	}
	content = rawData[dataStartIndex:dataEndIndex]

	var signStartIndex = signIndex + len(k_SIGN_NODE_NAME) + 4
	sign = rawData[signStartIndex:]
	var signEndIndex = strings.LastIndex(sign, "\"}")
	sign = sign[:signEndIndex]

	return content, sign
}
