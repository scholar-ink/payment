package enter

import (
	"bytes"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"

	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	maps "github.com/scholar-ink/payment/util/map"
	"io/ioutil"
	"strings"
)

const (
	AdaPayEnterUrl = "https://api.adapay.tech/"
)

type AdaPayEnter struct {
	*AdaPayConfig
	Service string `json:"-"`
	Sign    string `json:"-"` //签名
}

type AdaPayConfig struct {
	ApiKey     string `json:"-"`
	PrivateKey string `json:"-"`
	PublicKey  string `json:"-"`
}

type AdaRate struct {
	RateChannel string `json:"rate_channel"`
	FeeRate     string `json:"fee_rate"`
}

type AdaPayCreateConf struct {
	RequestId               string     `json:"request_id"`                  //请求ID
	UsrPhone                string     `json:"usr_phone"`                   //注册手机号
	ContName                string     `json:"cont_name"`                   //联系人姓名
	ContPhone               string     `json:"cont_phone"`                  //联系人手机号码
	CustomerEmail           string     `json:"customer_email"`              //电子邮箱
	MerName                 string     `json:"mer_name"`                    //商户名
	MerShortName            string     `json:"mer_short_name"`              //商户名简称
	LicenseCode             string     `json:"license_code"`                //营业执照编码
	RegAddr                 string     `json:"reg_addr,omitempty"`          //注册地址，企业时必填
	CustAddr                string     `json:"cust_addr"`                   //经营地址，企业时必填
	CustTel                 string     `json:"cust_tel"`                    //商户电话
	MerStartValidDate       string     `json:"mer_start_valid_date"`        //商户有效日期（始），格式 YYYYMMDD
	MerValidDate            string     `json:"mer_valid_date"`              //商户有效日期（至），格式 YYYYMMDD
	LegalName               string     `json:"legal_name"`                  //法人/负责人 姓名
	LegalType               string     `json:"legal_type"`                  //法人证件类型，0-身份证 1-其他
	LegalIdNo               string     `json:"legal_idno"`                  //法人证件号码
	LegalMp                 string     `json:"legal_mp"`                    //法人手机号
	LegalStartCertIdExpires string     `json:"legal_start_cert_id_expires"` //法人身份证有效期（始），格式 YYYYMMDD
	LegalIdExpires          string     `json:"legal_id_expires"`            //法人身份证有效期（至），格式 YYYYMMDD
	CardName                string     `json:"card_name"`                   //结算银行卡开户姓名
	CardIdMask              string     `json:"card_id_mask"`                //结算银行卡号
	BankAcctType            string     `json:"bank_acct_type"`              //结算银行账户类型，1 : 对公， 2 : 对私
	ProvCode                string     `json:"prov_code"`                   //结算银行卡省份编码
	AreaCode                string     `json:"area_code"`                   //结算银行卡地区编码
	BankCode                string     `json:"bank_code"`                   //结算银行卡所属银行code
	RsaPublicKey            string     `json:"rsa_public_key"`              //商户rsa 公钥
	FeeRateList             []*AdaRate `json:"fee_rate_list,omitempty"`     //费率列表
}

type AdaPayCreateReturn struct {
	RequestId string `json:"request_id"`
}

type AdaPayConfigConf struct {
	RequestId      string                 `json:"request_id"`      //请求ID
	SubApiKey      string                 `json:"sub_api_key"`     //商户开户进件返回的API Key
	BankChannelNo  string                 `json:"bank_channel_no"` //通过线下对接人员获取到的渠道号
	FeeType        string                 `json:"fee_type"`        //费率类型：01-标准费率线上，02-标准费率线下
	AppId          string                 `json:"app_id"`          //商户开户进件返回的应用ID
	WxCategory     string                 `json:"wx_category"`     //微信经营类目
	AliPayCategory string                 `json:"alipay_category"` //支付宝经营类目
	ClsId          string                 `json:"cls_id"`          //行业分类
	ModelType      string                 `json:"model_type"`      //入驻模式 入驻模式：1-服务商模式
	MerType        string                 `json:"mer_type"`        //商户种类 1-政府机构,2-国营企业,3-私营企业,4-外资企业,5-个体工商户,7-事业单位
	ProvinceCode   string                 `json:"province_code"`   //省份编码
	CityCode       string                 `json:"city_code"`       //城市编码
	DistrictCode   string                 `json:"district_code"`   //区县编码
	AddValueList   map[string]interface{} `json:"add_value_list"`  //配置信息
}

type AdaPayConfigReturn struct {
	RequestId string `json:"request_id"`
	Status    string `json:"status"`
}

type AdaPayQueryConf struct {
	RequestId string `json:"request_id"` //请求ID
}

type AdaAppId struct {
	AppId   string `json:"app_id"`
	AppName string `json:"app_name"`
}

type AdaPayQueryReturn struct {
	LiveApiKey string      `json:"live_api_key"` //商户编号
	AppIdList  []*AdaAppId `json:"app_id_list"`  //商户编号
}

type AdaPayConfigQueryConf struct {
	RequestId string `json:"request_id"` //请求ID
}

type AdaPayConfigQueryReturn struct {
	AliPayStat map[string]string `json:"alipay_stat"` //商户编号
	WxStat     map[string]string `json:"wx_stat"`     //商户编号
}

func (ada *AdaPayEnter) InitConfig(config *AdaPayConfig) {
	ada.AdaPayConfig = config
}

func (ada *AdaPayEnter) Create(config *AdaPayCreateConf) (*AdaPayCreateReturn, error) {
	ada.Service = "v1/batchEntrys/userEntry"

	config.RsaPublicKey = ada.PublicKey

	ret, err := ada.sendReq(AdaPayEnterUrl, config)
	if err != nil {
		return nil, err
	}

	ret, err = ada.RetData(ret)
	if err != nil {
		return nil, err
	}

	enterReturn := new(AdaPayCreateReturn)
	err = json.Unmarshal(ret, &enterReturn)
	if err != nil {
		return nil, err
	}

	return enterReturn, nil
}

func (ada *AdaPayEnter) Query(config *AdaPayQueryConf) (*AdaPayQueryReturn, error) {
	ada.Service = "v1/batchEntrys/userEntry"

	ret, err := ada.sendGetReq(AdaPayEnterUrl, config)
	if err != nil {
		return nil, err
	}

	ret, err = ada.RetData(ret)
	if err != nil {
		return nil, err
	}

	queryReturn := new(AdaPayQueryReturn)

	err = json.Unmarshal(ret, &queryReturn)

	if err != nil {
		return nil, err
	}

	return queryReturn, nil
}

func (ada *AdaPayEnter) Config(config *AdaPayConfigConf) (*AdaPayConfigReturn, error) {
	ada.Service = "v1/batchInput/merConf"

	config.RequestId = strings.ReplaceAll(uuid.New().String(), "-", "")

	ret, err := ada.sendReq(AdaPayEnterUrl, config)
	if err != nil {
		return nil, err
	}

	ret, err = ada.RetData(ret)
	if err != nil {
		return nil, err
	}

	configReturn := new(AdaPayConfigReturn)
	err = json.Unmarshal(ret, &configReturn)
	if err != nil {
		return nil, err
	}

	return configReturn, nil
}
func (ada *AdaPayEnter) ConfigQuery(config *AdaPayConfigQueryConf) (*AdaPayConfigQueryReturn, error) {
	ada.Service = "v1/batchInput/merConf"

	ret, err := ada.sendGetReq(AdaPayEnterUrl, config)
	if err != nil {
		return nil, err
	}

	ret, err = ada.RetData(ret)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(ret))

	queryReturn := new(AdaPayConfigQueryReturn)

	err = json.Unmarshal(ret, &queryReturn)

	if err != nil {
		return nil, err
	}

	return queryReturn, nil
}

func (ada *AdaPayEnter) makeSign(reqUrl string, requestJson string) {

	sign := reqUrl + requestJson

	b, err := helper.Sha1WithRsaSignPkcs8([]byte(sign), ada.PrivateKey)

	if err != nil {
		fmt.Println(err)
	}

	ada.Sign = base64.StdEncoding.EncodeToString(b)
}

func (ada *AdaPayEnter) sendReq(reqUrl string, params interface{}) (b []byte, err error) {

	b, err = json.Marshal(params)

	if err != nil {
		return nil, err
	}

	ada.makeSign(reqUrl+ada.Service, string(b))

	req := http.NewHttpRequest("POST", reqUrl+ada.Service, bytes.NewBuffer(b))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", ada.ApiKey)
	req.Header.Add("Signature", ada.Sign)

	rsp, err := http.Client.Do(req)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	return
}
func (ada *AdaPayEnter) sendGetReq(reqUrl string, params interface{}) (b []byte, err error) {

	mapData := helper.Struct2Map(params)

	values := maps.Map2Values(&mapData)

	ada.makeSign(reqUrl+ada.Service, values.Encode())

	req := http.NewHttpRequest("GET", reqUrl+ada.Service+"?"+values.Encode(), nil)

	req.Header.Set("Content-Type", "text/html")
	req.Header.Add("Authorization", ada.ApiKey)
	req.Header.Add("Signature", ada.Sign)

	rsp, err := http.Client.Do(req)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	return
}
func (ada *AdaPayEnter) RetData(ret []byte) (b []byte, err error) {

	type BaseResponse struct {
		Data string `json:"data"`
	}

	baseResponse := new(BaseResponse)

	err = json.Unmarshal(ret, &baseResponse)

	if err != nil {
		return
	}

	type BaseReturn struct {
		Status   string `json:"status"`
		ErrorMsg string `json:"error_msg"`
	}

	baseReturn := new(BaseReturn)

	b = []byte(baseResponse.Data)

	err = json.Unmarshal(b, &baseReturn)

	if err != nil {
		return
	}

	if baseReturn.Status == "failed" {
		err = errors.New(baseReturn.ErrorMsg)
		return
	}
	return
}

// 图片信息
type AdaPayComposeAgreementConfig struct {
	FingerUrl       string `json:"finger_url"`
	ShopName        string `json:"shop_name"`
	LoginNo         string `json:"login_no"`
	LicenseNo       string `json:"license_no"`
	LicenseDeadline string `json:"license_deadline"`
	Industry2Name   string `json:"industry_2_name"`
	Address         string `json:"address"`
	LicenseArea     string `json:"license_area"`
	IdCardName      string `json:"id_card_name"`
	IdCardNo        string `json:"id_card_no"`
	IdCardDeadline  string `json:"id_card_deadline"`
	ContactName     string `json:"contact_name"`
	ContactMobile   string `json:"contact_mobile"`
	AccountName     string `json:"account_name"`
	BankName        string `json:"bank_name"`
	BankCardNo      string `json:"bank_card_no"`
	AliRate         int32  `json:"ali_rate"`
	WxRate          int32  `json:"wx_rate"`
	BankRate        int32  `json:"bank_rate"`
	SettleCycle     string `json:"settle_cycle"`
}
