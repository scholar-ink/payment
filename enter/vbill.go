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
	"time"
)

const (
	VBillEnterUrl = "https://openapi-test.tianquetech.com/"
)

type VBillEnter struct {
	*VBillConfig
	Service   string      `json:"-"`
	ReqId     string      `json:"reqId"`     //请求ID
	Version   string      `json:"version"`   //版本
	Timestamp string      `json:"timestamp"` //请求时间
	SignType  string      `json:"signType"`  //签名类型
	Sign      string      `json:"Sign"`      //签名
	ReqData   interface{} `json:"reqData"`   //业务数据
}

type VBillConfig struct {
	OrgId      string `json:"orgId"`
	PrivateKey string `json:"-"`
	PublicKey  string `json:"-"`
}

type VBillCreateConf struct {
	UsrPhone                string `json:"usr_phone"`                   //注册手机号
	ContName                string `json:"cont_name"`                   //联系人姓名
	ContPhone               string `json:"cont_phone"`                  //联系人手机号码
	CustomerEmail           string `json:"customer_email"`              //电子邮箱
	MerName                 string `json:"mer_name"`                    //商户名
	MerShortName            string `json:"mer_short_name"`              //商户名简称
	LicenseCode             string `json:"license_code"`                //营业执照编码
	RegAddr                 string `json:"reg_addr,omitempty"`          //注册地址，企业时必填
	CustAddr                string `json:"cust_addr"`                   //经营地址，企业时必填
	CustTel                 string `json:"cust_tel"`                    //商户电话
	MerStartValidDate       string `json:"mer_start_valid_date"`        //商户有效日期（始），格式 YYYYMMDD
	MerValidDate            string `json:"mer_valid_date"`              //商户有效日期（至），格式 YYYYMMDD
	LegalName               string `json:"legal_name"`                  //法人/负责人 姓名
	LegalType               string `json:"legal_type"`                  //法人证件类型，0-身份证 1-其他
	LegalIdNo               string `json:"legal_idno"`                  //法人证件号码
	LegalMp                 string `json:"legal_mp"`                    //法人手机号
	LegalStartCertIdExpires string `json:"legal_start_cert_id_expires"` //法人身份证有效期（始），格式 YYYYMMDD
	LegalIdExpires          string `json:"legal_id_expires"`            //法人身份证有效期（至），格式 YYYYMMDD
	CardName                string `json:"card_name"`                   //结算银行卡开户姓名
	CardIdMask              string `json:"card_id_mask"`                //结算银行卡号
	BankAcctType            string `json:"bank_acct_type"`              //结算银行账户类型，1 : 对公， 2 : 对私
	ProvCode                string `json:"prov_code"`                   //结算银行卡省份编码
	AreaCode                string `json:"area_code"`                   //结算银行卡地区编码
	BankCode                string `json:"bank_code"`                   //结算银行卡所属银行code
	RsaPublicKey            string `json:"rsa_public_key"`              //商户rsa 公钥
}

type VBillCreateReturn struct {
	RequestId string `json:"request_id"`
}

type VBillConfigConf struct {
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

type VBillConfigReturn struct {
	RequestId string `json:"request_id"`
	Status    string `json:"status"`
}

type VBillQueryConf struct {
	RequestId string `json:"request_id"` //请求ID
}

type VBillQueryReturn struct {
	LiveApiKey string `json:"live_api_key"` //商户编号
}

type VBillConfigQueryConf struct {
	RequestId string `json:"request_id"` //请求ID
}

type VBillConfigQueryReturn struct {
	AliPayStat map[string]string `json:"alipay_stat"` //商户编号
	WxStat     map[string]string `json:"wx_stat"`     //商户编号
}

func (vBill *VBillEnter) InitConfig(config *VBillConfig) {

	vBill.VBillConfig = config
}

func (vBill *VBillEnter) Create(config *VBillCreateConf) (*VBillCreateReturn, error) {
	vBill.Service = "merchant/income"
	vBill.ReqId = strings.ReplaceAll(uuid.New().String(), "-", "")
	vBill.Timestamp = time.Now().Format("20060102150405")
	vBill.SignType = "RSA"
	vBill.Version = "3"
	config.RsaPublicKey = vBill.PublicKey

	ret, err := vBill.sendReq(VBillEnterUrl, config)
	if err != nil {
		return nil, err
	}

	ret, err = vBill.RetData(ret)
	if err != nil {
		return nil, err
	}

	enterReturn := new(VBillCreateReturn)
	err = json.Unmarshal(ret, &enterReturn)
	if err != nil {
		return nil, err
	}

	return enterReturn, nil
}

func (vBill *VBillEnter) Query(config *VBillQueryConf) (*VBillQueryReturn, error) {
	vBill.Service = "v1/batchEntrys/userEntry"

	ret, err := vBill.sendGetReq(VBillEnterUrl, config)
	if err != nil {
		return nil, err
	}

	ret, err = vBill.RetData(ret)
	if err != nil {
		return nil, err
	}

	queryReturn := new(VBillQueryReturn)

	err = json.Unmarshal(ret, &queryReturn)

	if err != nil {
		return nil, err
	}

	return queryReturn, nil
}

func (vBill *VBillEnter) Config(config *VBillConfigConf) (*VBillConfigReturn, error) {
	vBill.Service = "v1/batchInput/merConf"

	config.RequestId = strings.ReplaceAll(uuid.New().String(), "-", "")

	ret, err := vBill.sendReq(VBillEnterUrl, config)
	if err != nil {
		return nil, err
	}

	ret, err = vBill.RetData(ret)
	if err != nil {
		return nil, err
	}

	configReturn := new(VBillConfigReturn)
	err = json.Unmarshal(ret, &configReturn)
	if err != nil {
		return nil, err
	}

	return configReturn, nil
}
func (vBill *VBillEnter) ConfigQuery(config *VBillConfigQueryConf) (*VBillConfigQueryReturn, error) {
	vBill.Service = "v1/batchInput/merConf"

	ret, err := vBill.sendGetReq(VBillEnterUrl, config)
	if err != nil {
		return nil, err
	}

	ret, err = vBill.RetData(ret)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(ret))

	queryReturn := new(VBillConfigQueryReturn)

	err = json.Unmarshal(ret, &queryReturn)

	if err != nil {
		return nil, err
	}

	return queryReturn, nil
}

func (vBill *VBillEnter) makeSign() {

	vBill.Sign = ""

	mapData := helper.Struct2Map(vBill)

	signStr := helper.CreateLinkString(&mapData)

	b, err := helper.Sha1WithRsaSignPkcs8([]byte(signStr), vBill.PrivateKey)

	if err != nil {
		fmt.Println(err)
	}

	vBill.Sign = base64.StdEncoding.EncodeToString(b)
}

func (vBill *VBillEnter) sendReq(reqUrl string, params interface{}) (b []byte, err error) {

	vBill.ReqData = params

	vBill.makeSign()

	b, err = json.Marshal(vBill)

	if err != nil {
		return nil, err
	}

	rsp, err := http.Client.Post(reqUrl+vBill.Service, "application/json;charset=UTF-8", bytes.NewBuffer(b))

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	return
}
func (vBill *VBillEnter) sendGetReq(reqUrl string, params interface{}) (b []byte, err error) {

	mapData := helper.Struct2Map(params)

	values := maps.Map2Values(&mapData)

	vBill.makeSign()

	req := http.NewHttpRequest("GET", reqUrl+vBill.Service+"?"+values.Encode(), nil)

	req.Header.Set("Content-Type", "text/html")
	req.Header.Add("Signature", vBill.Sign)

	rsp, err := http.Client.Do(req)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	return
}
func (vBill *VBillEnter) RetData(ret []byte) (b []byte, err error) {

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
type VBillComposeAgreementConfig struct {
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
