package enter

import (
	"bytes"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"

	"encoding/json"
	"errors"
	"fmt"
	"github.com/scholar-ink/payment/util/strings"
	"io/ioutil"
)

const (
	HjSharingEnterUrl       = "https://www.joinpay.com/allocFunds"
	HjSharingEnterUploadUrl = "https://upload.joinpay.com/allocFunds"
)

type HjSharingEnter struct {
	*HjSharingConfig
	Method   string      `json:"method"`
	Version  string      `json:"version"`
	Data     interface{} `json:"data"`
	RandStr  string      `json:"rand_str"`
	SignType string      `json:"sign_type"`
	Sign     string      `json:"sign"`
}

type HjSharingReturn struct {
	RespCode string                 `json:"resp_code"`
	RespMsg  string                 `json:"resp_msg"`
	Data     map[string]interface{} `json:"data"`
}

type HjSharingConfig struct {
	MchNo string `json:"mch_no"`
	Key   string `json:"key"`
}

type HjSharingCreateConf struct {
	LoginName           string `json:"login_name"`             //分账方登录名
	AltMchName          string `json:"alt_mch_name"`           //分账方全称
	AltMchShortName     string `json:"alt_mch_short_name"`     //分账方简称
	AltMerchantType     string `json:"alt_merchant_type"`      //分账方商户类型10:个人，11:个体工商户，12:企业
	BusiContactName     string `json:"busi_contact_name"`      //业务联系人姓名
	BusiContactMobileNo string `json:"busi_contact_mobile_no"` //业务联系人手机
	PhoneNo             string `json:"phone_no"`               //法人手机号
	ManageScope         string `json:"manage_scope"`           //经营范围
	ManageAddr          string `json:"manage_addr"`            //经营地址
	LegalPerson         string `json:"legal_person"`           //法人姓名
	IdCardNo            string `json:"id_card_no"`             //身份证号
	IdCardExpiry        string `json:"id_card_expiry"`         //身份证有效期
	LicenseNo           string `json:"license_no"`             //营业执照编号
	LicenseExpiry       string `json:"license_expiry"`         //营业执照有效期
	SettMode            string `json:"sett_mode"`              //结算方式
	SettDateType        string `json:"sett_date_type"`         //结算周期类型
	RiskDay             string `json:"risk_day"`               //结算周期
	BankAccountType     string `json:"bank_account_type"`      //结算账户类型
	BankAccountName     string `json:"bank_account_name"`      //银行账户名称
	BankAccountNo       string `json:"bank_account_no"`        //银行账号
	BankChannelNo       string `json:"bank_channel_no"`        //联行号
	NotifyUrl           string `json:"notify_url"`             //异步通知地址
}
type HjSharingCreateReturn struct {
	AltMchNo    string `json:"alt_mch_no"`
	AltMchName  string `json:"alt_mch_name"`
	OrderStatus string `json:"order_status"`
	BizCode     string `json:"biz_code"`
	BizMsg      string `json:"biz_msg"`
}

type HjSharingSignConf struct {
	AltMchNo   string `json:"alt_mch_no"`  //分账方编号
	SignStatus string `json:"sign_status"` //签约成功:P1000
	SignTime   string `json:"sign_time"`   //签约日期 yyyy-MM-dd hh:mm:ss 格式
}
type HjSharingSignReturn struct {
	AltMchNo   string `json:"alt_mch_no"`
	SignTrxNo  string `json:"sign_trx_no"`
	SignStatus string `json:"sign_status"`
	BizCode    string `json:"biz_code"`
	BizMsg     string `json:"biz_msg"`
}

type HjSharingModifyReturn struct {
	ResultCode  string `json:"resultCode"`
	ErrorCode   string `json:"errorCode"`
	ErrCodeDesc string `json:"errCodeDesc"`
	Status      string `json:"status"`
	Ext         string `json:"ext"`
}

type HjSharingQueryConf struct {
	AccountNo string `json:"accountNo"` //商户登陆账号
}

type HjSharingQueryReturn struct {
	AccountNo    string `json:"accountNo"`
	Status       string `json:"status"`
	AccountType  string `json:"accountType"`
	MerchantName string `json:"merchantName"`
	MerchantMail string `json:"merchantMail"`
	PhoneNo      string `json:"phoneNo"`
	SettleInfo   string `json:"settleInfo"`
	FeeInfo      string `json:"feeInfo"`
	FeeStartTime string `json:"feeStartTime"`
	FeeEndTime   string `json:"feeEndTime"`
}

type HjSharingUploadConf struct {
	AltMchNo           string `json:"alt_mch_no"`           //分账方编号
	CardPositive       string `json:"card_positive"`        //身份证正面
	CardNegative       string `json:"card_negative"`        //身份证反面
	TradeLicence       string `json:"trade_licence"`        //营业执照
	OpenAccountLicence string `json:"open_account_licence"` //开户许可证
}

type HjSharingUploadReturn struct {
	AltMchNo   string `json:"alt_mch_no"`
	SignTrxNo  string `json:"sign_trx_no"`
	SignStatus string `json:"sign_status"`
	BizCode    string `json:"biz_code"`
	BizMsg     string `json:"biz_msg"`
}

func (HjSharing *HjSharingEnter) InitConfig(config *HjSharingConfig) {
	HjSharing.Version = "1.1"
	HjSharing.HjSharingConfig = config
}

func (HjSharing *HjSharingEnter) Create(config *HjSharingCreateConf) (*HjSharingCreateReturn, error) {
	HjSharing.Method = "altmch.create"
	HjSharing.buildData(config)
	HjSharing.setSign()

	ret, err := HjSharing.sendReq(HjSharingEnterUrl, HjSharing)
	if err != nil {
		return nil, err
	}

	ret, err = HjSharing.retData(ret)
	if err != nil {
		return nil, err
	}

	enterReturn := new(HjSharingCreateReturn)

	err = json.Unmarshal(ret, &enterReturn)

	if err != nil {
		return nil, err
	}

	return enterReturn, nil
}
func (HjSharing *HjSharingEnter) Signing(config *HjSharingSignConf) (*HjSharingSignReturn, error) {
	HjSharing.Method = "altMchSign.sign"
	HjSharing.buildData(config)
	HjSharing.setSign()

	ret, err := HjSharing.sendReq(HjSharingEnterUrl, HjSharing)
	if err != nil {
		return nil, err
	}

	ret, err = HjSharing.retData(ret)
	if err != nil {
		return nil, err
	}

	signReturn := new(HjSharingSignReturn)

	err = json.Unmarshal(ret, &signReturn)

	if err != nil {
		return nil, err
	}

	return signReturn, nil
}
func (HjSharing *HjSharingEnter) Upload(config *HjSharingUploadConf) (*HjSharingUploadReturn, error) {
	HjSharing.Method = "altMchPics.uploadPic"
	HjSharing.buildData(config)
	HjSharing.setSign()

	ret, err := HjSharing.sendReq(HjSharingEnterUploadUrl, HjSharing)
	if err != nil {
		return nil, err
	}

	ret, err = HjSharing.retData(ret)
	if err != nil {
		return nil, err
	}

	uploadReturn := new(HjSharingUploadReturn)

	err = json.Unmarshal(ret, &uploadReturn)

	if err != nil {
		return nil, err
	}

	return uploadReturn, nil
}
func (HjSharing *HjSharingEnter) Query(config *HjSharingQueryConf) (*HjSharingQueryReturn, error) {
	HjSharing.Method = "merchant.enter.query"
	HjSharing.buildData(config)
	HjSharing.setSign()

	ret, err := HjSharing.sendReq(HjSharingEnterUrl, HjSharing)
	if err != nil {
		return nil, err
	}

	ret, err = HjSharing.retData(ret)
	if err != nil {
		return nil, err
	}

	queryReturn := new(HjSharingQueryReturn)
	err = json.Unmarshal(ret, queryReturn)
	if err != nil {
		return nil, err
	}

	return queryReturn, nil
}
func (HjSharing *HjSharingEnter) buildData(config interface{}) error {
	HjSharing.SignType = "1"
	HjSharing.RandStr = helper.Md5(strings.CreateIdentifyCode())
	HjSharing.Data = config
	return nil
}
func (HjSharing *HjSharingEnter) setSign() {

	b, _ := json.Marshal(HjSharing.Data)

	HjSharing.Sign = helper.Md5(fmt.Sprintf("data=%s&mch_no=%s&method=%s&rand_str=%s&sign_type=%s&version=%s&key=%s", string(b), HjSharing.MchNo, HjSharing.Method, HjSharing.RandStr, HjSharing.SignType, HjSharing.Version, HjSharing.Key))
}

func (HjSharing *HjSharingEnter) sendReq(reqUrl string, params interface{}) (b []byte, err error) {

	buffer := bytes.NewBuffer(b)

	err = json.NewEncoder(buffer).Encode(params)

	if err != nil {
		return nil, err
	}

	rsp, err := http.Client.Post(reqUrl, "application/json;charset=UTF-8", buffer)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	return
}
func (HjSharing *HjSharingEnter) retData(ret []byte) (retData []byte, err error) {
	var baseReturn HjSharingReturn

	err = json.Unmarshal(ret, &baseReturn)

	if err != nil {
		return
	}

	if baseReturn.RespCode != "A1000" {
		bizMsg := baseReturn.Data["biz_msg"].(string)

		err = errors.New(bizMsg)

		return
	}

	return json.Marshal(baseReturn.Data)
}
