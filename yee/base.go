package yee

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/array"
	"github.com/scholar-ink/payment/util/http"
	"io/ioutil"
	"log"
	http2 "net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	ChargeUrl = "https://openapi.yeepay.com/yop-center"
	SUCCESS   = "SUCCESS"
)

type BaseCharge struct {
	*BaseConfig
	RequestNo string `json:"requestNo"`
	Sign      string `json:"-"`
	Service   string `json:"-"`
}

type BaseConfig struct {
	ParentMerchantNo string `json:"parentMerchantNo"`
	AppKey           string `json:"appKey"`
	PrivateKey       string `json:"-"`
}

func (base *BaseCharge) InitBaseConfig(config *BaseConfig) {
	base.BaseConfig = config
}

func (base *BaseCharge) makeSign(params map[string]interface{}, header http2.Header) {

	protocolVersion := "yop-auth-v2"
	expiredSeconds := "1800"
	timestamp := time.Now().Format("2006-01-02T15:04:05+0800")
	canonicalURI := base.Service
	canonicalQueryString := base.GetSortQuery(params)
	canonicalHeader := base.GetHeader(header)
	canonicalRequest := protocolVersion + "/" + base.AppKey + "/" + timestamp + "/" + expiredSeconds + "\n" + "POST" + "\n" + canonicalURI + "\n" + canonicalQueryString + "\n" + canonicalHeader

	fmt.Println(canonicalRequest)

	b, _ := helper.Sha256WithRsaSign([]byte(canonicalRequest), base.PrivateKey)

	sign := base.Base64Encode(b, false)

	base.Sign = "YOP-RSA2048-SHA256 " + protocolVersion + "/" + base.AppKey + "/" + timestamp + "/" + expiredSeconds + "/" + base.GetSignedHeader() + "/" + sign + "$SHA256"
}

func (base *BaseCharge) SendReq(reqUrl string, params interface{}) (b []byte, err error) {

	mapData := helper.Struct2Map(params)

	mapData["parentMerchantNo"] = base.ParentMerchantNo
	//mapData["requestNo"] = base.RequestNo

	postBody := base.GetPostBody(mapData)

	req := http.NewHttpRequest("POST", reqUrl+base.Service, strings.NewReader(postBody))

	req.Header["x-yop-appkey"] = []string{base.AppKey}
	req.Header["x-yop-request-id"] = []string{base.RequestNo}

	base.makeSign(mapData, req.Header)

	req.Header["Authorization"] = []string{base.Sign}
	req.Header["x-yop-sdk-langs"] = []string{"go"}
	req.Header["x-yop-sdk-version"] = []string{"3.0.0"}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	fmt.Println(req.Header)

	rsp, err := http.Client.Do(req)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	log.Println("[enter.base." + base.Service + "] response:" + string(b))

	return
}
func (base *BaseCharge) retData(ret []byte) (b []byte, err error) {

	var baseReturn struct {
		Result map[string]interface{} `json:"result"`
	}

	err = json.Unmarshal(ret, &baseReturn)

	b, _ = json.Marshal(baseReturn.Result)

	return
}

func (base *BaseCharge) GetSortQuery(arr map[string]interface{}) string {
	var params []string
	var query string
	for key, val := range arr {
		if "Authorization" == key {
			continue
		}
		query = key + "=" + helper.RawUrlEncode(val.(string))
		params = append(params, query)
	}
	sort.Strings(params)
	return strings.Join(params, "&")
}
func (base *BaseCharge) GetHeader(arr http2.Header) string {
	allowed := []string{"x-yop-request-id"}

	var headers []string
	header := ""
	for key, val := range arr {
		if array.InArray(key, allowed) {
			header = helper.RawUrlEncode(strings.ToLower(key)) + ":" + helper.RawUrlEncode(val[0])
			headers = append(headers, header)
		}
	}
	sort.Strings(headers)
	return strings.Join(headers, "\n")
}

func (base *BaseCharge) GetSignedHeader() string {
	headers := []string{"x-yop-request-id"}
	str := strings.Join(headers, ";")
	return strings.ToLower(str)
}

func (base *BaseCharge) Base64Encode(data []byte, usePadding bool) string {
	if 0 == len(data) {
		return ""
	}

	encoded := strings.ReplaceAll(strings.ReplaceAll(base64.StdEncoding.EncodeToString(data), "+", "-"), "/", "_")

	if usePadding {
		return encoded
	}
	str := strings.TrimRight(encoded, "=")
	return str
}
func (base *BaseCharge) GetPostBody(params map[string]interface{}) string {
	var list []string
	query := ""
	for key, val := range params {
		query = key + "=" + url.QueryEscape(val.(string))
		list = append(list, query)
	}
	sort.Strings(list)
	return strings.Join(list, "&")
}
