package enter

import (
	"bytes"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/array"
	"github.com/scholar-ink/payment/util/http"

	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"mime/multipart"
	http2 "net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	YeeEnterUrl = "https://openapi.yeepay.com/yop-center"
)

type YeeEnter struct {
	*YeeConfig
	Service   string `json:"-"`
	RequestNo string `json:"requestNo"`
	Sign      string `json:"sign"`
}

type YeeConfig struct {
	ParentMerchantNo string `json:"parentMerchantNo"`
	AppKey           string `json:"appKey"`
	PrivateKey       string `json:"-"`
}

//产品信息
type YeeProduct struct {
	ChannelType string `json:"channel_type"` //通道类型
	FeeExp      string `json:"fee_exp"`      //费率
}

//文件信息
type YeeFile struct {
	QuaType string `json:"quaType"` //文件类型
	QuaUrl  string `json:"quaUrl"`  //文件地址
}

//报备信息
type ReportInfo struct {
	IDCardName         string `json:"IDCardName"`                 //法人身份证姓名
	ThreeInOne         string `json:"ThreeInOne"`                 //是否三证合一
	IsInstitution      string `json:"IsInstitution"`              //是否事业单位
	MerTypeX           string `json:"merTypeX"`                   //商户类型
	PicIdCardFront     string `json:"pic_id_card_front"`          //身份证-正面
	PicIdCardBack      string `json:"pic_id_card_back"`           //身份证-反面
	PicMerchantLicense string `json:"pic_merchant_license"`       //营业执照
	PicMerShopPhoto    string `json:"pic_mer_shop_photo"`         //门头照
	PicMerProtocol     string `json:"pic_mer_protocol,omitempty"` //个人用户必填， 此项上传手持银行卡照片
}

type YeeCreateConf struct {
	MerType          int32                  `json:"-"`
	MerFullName      string                 `json:"merFullName"`               //商户全称
	MerShortName     string                 `json:"merShortName"`              //商户简称
	MerCertNo        string                 `json:"merCertNo"`                 //营业执照号
	LegalName        string                 `json:"legalName"`                 //法人姓名
	LegalIdCard      string                 `json:"legalIdCard"`               //法人身份证号
	MerLegalPhone    string                 `json:"merLegalPhone"`             //法人手机号
	MerLevel1No      string                 `json:"merLevel1No"`               //商户一级分类
	MerLevel2No      string                 `json:"merLevel2No"`               //商户二级分类
	MerProvince      string                 `json:"merProvince"`               //商户所在省
	MerCity          string                 `json:"merCity"`                   //商户所在市
	MerDistrict      string                 `json:"merDistrict"`               //商户所在区
	MerAddress       string                 `json:"merAddress"`                //商户所在地址
	BankAccountType  string                 `json:"bankAccountType,omitempty"` //银行账户类型 //PERSONAL CORPORATE
	CardNo           string                 `json:"cardNo"`                    //结算银行卡
	HeadBankCode     string                 `json:"headBankCode,omitempty"`    //商户开户总行编码
	BankCode         string                 `json:"bankCode,omitempty"`        //开户支行编码
	BankProvince     string                 `json:"bankProvince"`              //开户省
	BankCity         string                 `json:"bankCity"`                  //开户市
	ProductInfoStr   string                 `json:"productInfo"`               //开通产品信息
	ProductInfo      map[string]interface{} `json:"-"`                         //开通产品信息
	FileInfoStr      string                 `json:"fileInfo"`                  //文件信息
	FileInfo         []*YeeFile             `json:"-"`                         //文件信息
	NotifyUrl        string                 `json:"notifyUrl"`                 //商户回调地址
	MerAuthorizeType string                 `json:"merAuthorizeType"`          //授权类型	SMS_AUTHORIZE("短信授权")
}

type YeeCreateReturn struct {
	MerchantNo string `json:"merchantNo"`
	ExternalId string `json:"externalId"`
}

type YeeUpdateConf struct {
	MerchantNo     string                 `json:"merchantNo"`
	ProductInfoStr string                 `json:"payProductInfo"` //开通产品信息
	ProductInfo    map[string]interface{} `json:"-"`              //开通产品信息
}

type YeeUpdateReturn struct {
	MerchantNo string `json:"merNo"`
	AllSuccess string `json:"allSuccess"`
}

type YeeReportConf struct {
	MerchantNo            string      `json:"merchantNo"`
	ChannelNo             string      `json:"channelNo,omitempty"`
	CallBackUrl           string      `json:"callBackUrl"`
	MerchantName          string      `json:"merchantName"`
	ReportMerchantName    string      `json:"reportMerchantName"`
	ReportMerchantAlias   string      `json:"reportMerchantAlias"`
	ReportMerchantComment string      `json:"reportMerchantComment"`
	ServiceTel            string      `json:"serviceTel"`
	ContactName           string      `json:"contactName"`
	ContactPhone          string      `json:"contactPhone"`
	ContactMobile         string      `json:"contactMobile"`
	ContactEmail          string      `json:"contactEmail"`
	MerchantAddress       string      `json:"merchantAddress"`
	MerchantProvince      string      `json:"merchantProvince"`
	MerchantCity          string      `json:"merchantCity"`
	MerchantDistrict      string      `json:"merchantDistrict"`
	MerchantLicenseNo     string      `json:"merchantLicenseNo"`
	CorporateIdCardNo     string      `json:"corporateIdCardNo"`
	ContactType           string      `json:"contactType"`
	ReportInfosJsonStr    string      `json:"reportInfosJsonStr"`
	ReportInfo            *ReportInfo `json:"-"`
	ReportFeeType         string      `json:"reportFeeType"`
}

type YeeReportReturn struct {
	TraceId    string `json:"traceId"`
	DealStatus int32  `json:"dealStatus"`
	BizMsg     string `json:"bizMsg"`
}

type YeeHMacKeyQueryConf struct {
	MerchantNo string `json:"merchantNo"` //子商户号
}
type YeHMacKeyQueryReturn struct {
	MerHMacKey string `json:"merHmacKey"`
}
type YeeBranchInfoConf struct {
	HeadBankCode string `json:"headBankCode"` //商户编号
	ProvinceCode string `json:"provinceCode"` //结算卡号
	CityCode     string `json:"cityCode"`     //开户行银行名称
}

type YeeReportQueryConf struct {
	MerchantNo string `json:"merchantNo"` //子商户号
}

type YeeReportQueryReturn struct {
	MerHMacKey string `json:"merHmacKey"`
}

type YeeRegStatusQueryConf struct {
	ParentMerchantNo string `json:"parentMerchantNo"` //子商户号
	MerchantNo       string `json:"merchantNo"`       //子商户号
}

type YeeRegStatusQueryReturn struct {
	MerHMacKey string `json:"merHmacKey"`
}

type YeeUploadConf struct {
	File []byte
}

func (yee *YeeEnter) InitConfig(config *YeeConfig) {
	yee.YeeConfig = config
}

func (yee *YeeEnter) Create(config *YeeCreateConf) (*YeeCreateReturn, error) {
	if config.MerType == 3 {
		yee.Service = "/rest/v1.0/sys/merchant/personreginfoadd"
	} else if config.MerType == 2 {
		yee.Service = "/rest/v1.0/sys/merchant/individualreginfoadd"
	} else {
		yee.Service = "/rest/v1.0/sys/merchant/individualreginfoadd"
	}
	yee.RequestNo = strings.ReplaceAll(uuid.New().String(), "-", "")

	if len(config.ProductInfo) == 0 {
		config.ProductInfo = map[string]interface{}{}
	}
	b, _ := json.Marshal(config.ProductInfo)
	config.ProductInfoStr = string(b)

	if len(config.FileInfo) == 0 {
		config.FileInfo = make([]*YeeFile, 0, 0)
	}
	b, _ = json.Marshal(config.FileInfo)
	config.FileInfoStr = string(b)

	ret, err := yee.sendReq(YeeEnterUrl, config)
	if err != nil {
		return nil, err
	}

	b, err = yee.retData(ret)

	if err != nil {
		return nil, err
	}

	enterReturn := new(YeeCreateReturn)
	err = json.Unmarshal(b, &enterReturn)
	if err != nil {
		return nil, err
	}

	return enterReturn, nil
}
func (yee *YeeEnter) Update(config *YeeUpdateConf) (*YeeUpdateReturn, error) {
	yee.Service = "/rest/v1.0/router/modify/pay-product-info"

	yee.RequestNo = strings.ReplaceAll(uuid.New().String(), "-", "")

	if len(config.ProductInfo) == 0 {
		config.ProductInfo = map[string]interface{}{}
	}
	b, _ := json.Marshal(config.ProductInfo)
	config.ProductInfoStr = string(b)

	ret, err := yee.sendReq(YeeEnterUrl, config)
	if err != nil {
		return nil, err
	}

	b, err = yee.retData(ret)

	if err != nil {
		return nil, err
	}

	updateReturn := new(YeeUpdateReturn)
	err = json.Unmarshal(b, &updateReturn)
	if err != nil {
		return nil, err
	}

	return updateReturn, nil
}
func (yee *YeeEnter) Report(config *YeeReportConf) (*YeeReportReturn, error) {
	yee.Service = "/rest/v1.0/router/open-pay-async-report/report"
	yee.RequestNo = strings.ReplaceAll(uuid.New().String(), "-", "")

	if config.ReportInfo == nil {
		config.ReportInfo = new(ReportInfo)
	}
	b, _ := json.Marshal(config.ReportInfo)
	config.ReportInfosJsonStr = string(b)

	ret, err := yee.sendReq(YeeEnterUrl, config)
	if err != nil {
		return nil, err
	}

	b, err = yee.retData(ret)

	if err != nil {
		return nil, err
	}

	reportReturn := new(YeeReportReturn)
	err = json.Unmarshal(b, &reportReturn)

	if err != nil {
		return nil, err
	}

	return reportReturn, nil
}
func (yee *YeeEnter) HMacKeyQuery(config *YeeHMacKeyQueryConf) (string, error) {
	yee.Service = "/rest/v1.0/sys/merchant/hmackeyquery"
	yee.RequestNo = strings.ReplaceAll(uuid.New().String(), "-", "")

	ret, err := yee.sendReq(YeeEnterUrl, config)
	if err != nil {
		return "", err
	}

	ret, err = yee.retData(ret)

	if err != nil {
		return "", err
	}

	fmt.Println(string(ret))

	hMacKeyQueryReturn := new(YeHMacKeyQueryReturn)

	err = json.Unmarshal(ret, &hMacKeyQueryReturn)

	if err != nil {
		return "", err
	}
	return hMacKeyQueryReturn.MerHMacKey, nil
}
func (yee *YeeEnter) BranchInfo(config *YeeBranchInfoConf) (*YeeCreateReturn, error) {
	yee.Service = "/rest/v1.0/sys/merchant/bankbranchInfo"
	yee.RequestNo = strings.ReplaceAll(uuid.New().String(), "-", "")

	ret, err := yee.sendReq(YeeEnterUrl, config)
	if err != nil {
		return nil, err
	}

	b, err := yee.retData(ret)

	if err != nil {
		return nil, err
	}

	enterReturn := new(YeeCreateReturn)
	err = json.Unmarshal(b, &enterReturn)
	if err != nil {
		return nil, err
	}

	return enterReturn, nil
}
func (yee *YeeEnter) ReportQuery(config *YeeReportQueryConf) (*YeeReportQueryReturn, error) {
	yee.Service = "/rest/v1.0/router/open-pay-report/query"
	yee.RequestNo = strings.ReplaceAll(uuid.New().String(), "-", "")

	ret, err := yee.sendReq(YeeEnterUrl, config)
	if err != nil {
		return nil, err
	}

	b, err := yee.retData(ret)

	if err != nil {
		return nil, err
	}

	reportQueryReturn := new(YeeReportQueryReturn)
	err = json.Unmarshal(b, &reportQueryReturn)
	if err != nil {
		return nil, err
	}

	return reportQueryReturn, nil
}
func (yee *YeeEnter) RegStatusQuery(config *YeeRegStatusQueryConf) (*YeeRegStatusQueryReturn, error) {
	yee.Service = "/rest/v1.0/sys/merchant/regstatusquery"
	yee.RequestNo = strings.ReplaceAll(uuid.New().String(), "-", "")

	ret, err := yee.sendReq(YeeEnterUrl, config)
	if err != nil {
		return nil, err
	}

	b, err := yee.retData(ret)

	if err != nil {
		return nil, err
	}

	regStatusQueryReturn := new(YeeRegStatusQueryReturn)
	err = json.Unmarshal(b, &regStatusQueryReturn)
	if err != nil {
		return nil, err
	}

	return regStatusQueryReturn, nil
}
func (yee *YeeEnter) Upload(config *YeeUploadConf) (string, error) {
	yee.Service = "/yos/v1.0/sys/merchant/qual/upload"
	yee.RequestNo = strings.ReplaceAll(uuid.New().String(), "-", "")

	ret, err := yee.sendUploadReq(YeeEnterUrl, config)

	if err != nil {
		return "", err
	}

	b, err := yee.retData(ret)

	if err != nil {
		return "", err
	}

	var uploadReturn struct {
		MerQualUrl string `json:"merQualUrl"`
	}

	json.Unmarshal(b, &uploadReturn)

	return uploadReturn.MerQualUrl, nil
}
func (yee *YeeEnter) buildData(config interface{}) error {
	return nil
}
func (yee *YeeEnter) makeSign(params map[string]interface{}, header http2.Header) {

	protocolVersion := "yop-auth-v2"
	expiredSeconds := "1800"
	timestamp := time.Now().Format("2006-01-02T15:04:05+0800")
	canonicalURI := yee.Service
	canonicalQueryString := yee.GetSortQuery(params)
	canonicalHeader := yee.GetHeader(header)
	canonicalRequest := protocolVersion + "/" + yee.AppKey + "/" + timestamp + "/" + expiredSeconds + "\n" + "POST" + "\n" + canonicalURI + "\n" + canonicalQueryString + "\n" + canonicalHeader

	fmt.Println(canonicalRequest)

	b, _ := helper.Sha256WithRsaSign([]byte(canonicalRequest), yee.PrivateKey)

	sign := yee.Base64Encode(b, false)

	yee.Sign = "YOP-RSA2048-SHA256 " + protocolVersion + "/" + yee.AppKey + "/" + timestamp + "/" + expiredSeconds + "/" + yee.GetSignedHeader() + "/" + sign + "$SHA256"

	fmt.Println(yee.Sign)

}

func (yee *YeeEnter) sendReq(reqUrl string, params interface{}) (b []byte, err error) {

	mapData := helper.Struct2Map(params)

	mapData["parentMerchantNo"] = yee.ParentMerchantNo
	mapData["requestNo"] = yee.RequestNo

	postBody := yee.GetPostBody(mapData)

	req := http.NewHttpRequest("POST", reqUrl+yee.Service, strings.NewReader(postBody))

	req.Header["x-yop-appkey"] = []string{yee.AppKey}
	req.Header["x-yop-request-id"] = []string{yee.RequestNo}

	yee.makeSign(mapData, req.Header)

	req.Header["Authorization"] = []string{yee.Sign}
	req.Header["x-yop-sdk-langs"] = []string{"go"}
	req.Header["x-yop-sdk-version"] = []string{"3.0.0"}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	rsp, err := http.Client.Do(req)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	return
}
func (yee *YeeEnter) sendUploadReq(reqUrl string, params *YeeUploadConf) (b []byte, err error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	//关键的一步操作
	part, err := writer.CreateFormFile("merQual", "merQual.png")

	if err != nil {
		return nil, err
	}

	//iocopy
	_, err = io.Copy(part, bytes.NewReader(params.File))

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	contentType := writer.FormDataContentType()

	req := http.NewHttpRequest("POST", reqUrl+yee.Service, body)

	req.Header["x-yop-appkey"] = []string{yee.AppKey}
	req.Header["x-yop-request-id"] = []string{yee.RequestNo}

	mapData := map[string]interface{}{}

	yee.makeSign(mapData, req.Header)

	req.Header["Authorization"] = []string{yee.Sign}
	req.Header["x-yop-sdk-langs"] = []string{"go"}
	req.Header["x-yop-sdk-version"] = []string{"3.0.0"}
	req.Header.Set("Content-Type", contentType)

	rsp, err := http.Client.Do(req)

	if err != nil {
		return nil, err
	}

	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	if err != nil {
		return nil, err
	}

	return b, nil
}
func (yee *YeeEnter) retData(ret []byte) (b []byte, err error) {

	var baseReturn struct {
		Code       string      `json:"code"`
		Message    string      `json:"message"`
		SubCode    string      `json:"subCode"`
		SubMessage string      `json:"subMessage"`
		Result     interface{} `json:"Result"`
	}

	var result struct {
		ReturnMsg  string `json:"returnMsg"`
		ReturnCode string `json:"returnCode"`
	}

	err = json.Unmarshal(ret, &baseReturn)

	if err != nil {
		return
	}

	helper.DeepCopy(baseReturn.Result, &result)

	if result.ReturnCode != "REG00000" && result.ReturnCode != "" {
		if baseReturn.SubMessage == "" {
			err = errors.New(result.ReturnMsg)
		} else {
			err = errors.New(baseReturn.SubMessage)
		}
		return
	}

	b, _ = json.Marshal(baseReturn.Result)

	return
}

func (yee *YeeEnter) GetSortQuery(arr map[string]interface{}) string {
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
func (yee *YeeEnter) GetHeader(arr http2.Header) string {
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

func (yee *YeeEnter) GetSignedHeader() string {
	headers := []string{"x-yop-request-id"}
	str := strings.Join(headers, ";")
	return strings.ToLower(str)
}

func (yee *YeeEnter) Base64Encode(data []byte, usePadding bool) string {
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
func (yee *YeeEnter) GetPostBody(params map[string]interface{}) string {
	var list []string
	query := ""
	for key, val := range params {
		query = key + "=" + url.QueryEscape(val.(string))
		list = append(list, query)
	}
	sort.Strings(list)
	return strings.Join(list, "&")
}
