package enter

import (
	"bytes"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"

	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/url"
)

const (
	HlbUploadUrl = "http://entry.trx.helipay.com/trx/merchantEntry/upload.action"
	HlbEnterUrl  = "http://entry.trx.helipay.com/trx/merchantEntry/interface.action"
)

type HlbEnter struct {
	*HlbConfig
	InterfaceName string `json:"interfaceName"`
	Body          string `json:"body"`
	Sign          string `json:"sign"`
}

type HlbConfig struct {
	MerchantNo string `json:"merchantNo"`
	Key        string `json:"key"`
	SignKey    string `json:"sign_key"`
}

// 图片信息
type HlbPicInfo struct {
	OrganizationCodePic           string `json:"organizationCodePic,omitempty"`        //组织机构代码图片
	TaxNoPic                      string `json:"taxNoPic,omitempty"`                   //税务登记号图片
	LicenseDeadlinePic            string `json:"licenseDeadlinePic,omitempty"`         //营业执照号图片
	OpeningAccountsDeadlinePic    string `json:"openingAccountsDeadlinePic,omitempty"` //开户许可证图片
	BusinessStandingPic           string `json:"businessStandingPic,omitempty"`        //企业信用公示图片
	LegalPersonIdFrontPic         string `json:"legalPersonIdFrontPic"`                //法人身份证正面
	LegalPersonIdOppositePic      string `json:"legalPersonIdOppositePic"`             //法人身份证反面
	LegalPersonBanKCardPic        string `json:"legalPersonBanKCardPic"`               //法人银行卡正面
	OperatorIdDeadlineFrontPic    string `json:"operatorIdDeadlineFrontPic"`           //经办人身份证正面
	OperatorIdDeadlineOppositePic string `json:"operatorIdDeadlineOppositePic"`        //经办人身份证反面
	MerchantDoorHeadPic           string `json:"merchantDoorHeadPic"`                  //商户门牌照片
	MerchantFrontPic              string `json:"merchantFrontPic"`                     //商户门脸照片
	MerchantInsidePic             string `json:"merchantInsidePic"`                    //商户内饰照片
	NoSealAgreement               string `json:"noSealAgreement,omitempty"`            //协议 - 未盖章
	SealAgreement                 string `json:"sealAgreement,omitempty"`              //协议 - 已盖章
	ContractConfirm               string `json:"contractConfirm"`                      //合同确认图片
	OtherPic                      string `json:"otherPic,omitempty"`                   //其他质料
}

// 结算信息
type HlbSettleInfo struct {
	SettleAccountType string `json:"settleAccountType"`         //结算账户类型 1-对公、2-法人对私、3-银通帐号
	OpenBankName      string `json:"openBankName,omitempty"`    //开户银行
	BranchBankName    string `json:"branchBankName,omitempty"`  //支行名称
	AccountName       string `json:"accountName"`               //账户名称
	BankCardNo        string `json:"bankCardNo"`                //银行卡卡号
	SettleMode        string `json:"settleMode,omitempty"`      //结算模式1-委托付款、 2-API（手动出款）
	SettleModeValue   string `json:"settleModeValue,omitempty"` //A1-15:30 A2-16:30 B1- 09:00|13:00|17:00
}

//费率信息
type HlbFeeInfo struct {
	ProductType string `json:"productType"` //产品类型
	FeeRate     string `json:"feeRate"`     //比例收取千分位：1=0.001 固定金额：单位：分
	FeeType     string `json:"feeType"`     //0-比例收取,1-固定金额
	ChargeRole  string `json:"chargeRole"`  //收费角色1-收款方、 2-付款方、 3-平台方
	SettleCycle string `json:"settleCycle"` //结算周期
}

type HlbCreateConf struct {
	OrderNo           string `json:"orderNo"`                     //商户订单号
	SignName          string `json:"signName"`                    //子商户签约名
	ShowName          string `json:"showName"`                    //展示名（收银台展示名）
	WebSite           string `json:"webSite,omitempty"`           //网站网址
	AccessUrl         string `json:"accessUrl,omitempty"`         //接入地址
	MerchantType      string `json:"merchantType"`                //子商户类型
	LegalPerson       string `json:"legalPerson"`                 //法人名字
	LegalPersonID     string `json:"legalPersonID"`               //法人身份证号
	OrgNum            string `json:"orgNum"`                      //组织机构代码
	BusinessLicense   string `json:"businessLicense"`             //营业执照号
	Province          string `json:"province,omitempty"`          //子商户所在省份
	City              string `json:"city,omitempty"`              //子商户所在城市
	RegionCode        string `json:"regionCode"`                  //区县编码
	Address           string `json:"address"`                     //通讯地址
	Linkman           string `json:"linkman"`                     //联系人
	LinkPhone         string `json:"linkPhone"`                   //联系电话
	Email             string `json:"email"`                       //联系邮箱
	BindMobile        string `json:"bindMobile,omitempty"`        //绑定手机
	ServicePhone      string `json:"servicePhone"`                //客服联系电话
	BankCode          string `json:"bankCode"`                    //结算卡联行号
	AccountName       string `json:"accountName"`                 //开户名
	AccountNo         string `json:"accountNo"`                   //开户账号
	SettleBankType    string `json:"settleBankType"`              //结算卡类型
	SettlementPeriod  string `json:"settlementPeriod"`            //结算类型
	SettlementMode    string `json:"settlementMode"`              //结算方式
	SettlementRemark  string `json:"settlementRemark,omitempty"`  //结算备注
	MerchantCategory  string `json:"merchantCategory"`            //经营类别
	IndustryTypeCode  string `json:"industryTypeCode,omitempty"`  //行业类型编码
	AuthorizationFlag bool   `json:"authorizationFlag"`           //授权使用平台商秘钥
	UnionPayQrCode    string `json:"unionPayQrCode,omitempty"`    //银联二维码
	NeedPosFunction   bool   `json:"needPosFunction"`             //是否需要开通 POS 功能
	IdCardStartDate   string `json:"idCardStartDate,omitempty"`   //法人身份证开始日期
	IdCardEndDate     string `json:"idCardEndDate,omitempty"`     //法人身份证结束日期
	BusinessDateStart string `json:"businessDateStart,omitempty"` //经营起始日期
	BusinessDateLimit string `json:"businessDateLimit,omitempty"` //经营期限
	AccountIdCard     string `json:"accountIdCard,omitempty"`     //开户人身份证
	Mcc               string `json:"mcc,omitempty"`               //银联 mcc 码
	AgreeProtocol     bool   `json:"agreeProtocol"`               //是否同意协议
	CallbackUrl       string `json:"callbackUrl,omitempty"`       //回调地址
	SettleMode        string `json:"settleMode"`                  //结算模式
	SettlementAuth    string `json:"settlementAuth,omitempty"`    //结算信息鉴权
	PostalAddress     string `json:"postalAddress,omitempty"`     //注册地址
	MicroBizType      string `json:"microBizType,omitempty"`      //小微经营类型
	CertType          string `json:"certType,omitempty"`          //证书类型
	LinkManId         string `json:"linkManId,omitempty"`         //联系人身份证号
	NeedAuthorize     string `json:"needAuthorize,omitempty"`     //是否需要认证
	SpecialSignName   string `json:"specialSignName,omitempty"`   //是否需要特殊处理商户名称
}

type HlbCreateReturn struct {
	EntryStatus string `json:"entryStatus"`
	MerchantNo  string `json:"merchantNo"`
	OrderNo     string `json:"orderNo"`
}

type HlbOpenProductConf struct {
	ProductType          string `json:"productType"`          //产品类型
	FirstClassMerchantNo string `json:"firstClassMerchantNo"` //平台商编号
	MerchantNo           string `json:"merchantNo"`           //子商户编号
	PayType              string `json:"payType,omitempty"`    //支付类型
	AppPayType           string `json:"appPayType"`           //客户端类型
	Value                string `json:"value,omitempty"`      //费率
	MinFee               string `json:"minFee,omitempty"`     //最低费率金额
	AppFeeMode           string `json:"appFeeMode,omitempty"` //费率模式
	FeePurpose           string `json:"feePurpose,omitempty"` //费率类型
}

type HlbModifyConf struct {
	AccountNo    string         `json:"accountNo"`    //商户登陆账号
	FeeInfo      []*HlbFeeInfo  `json:"feeInfo"`      //费率信息
	FeeStartTime string         `json:"feeStartTime"` //费率生效开始时间
	FeeEndTime   string         `json:"feeEndTime"`   //费率生效结束时间
	SettleInfo   *HlbSettleInfo `json:"settle_info"`  //结算信息
}

type HlbModifySettleInfoConf struct {
	AccountNo         string `json:"accountNo"`                 //商户登陆账号
	RequestNo         string `json:"request_no"`                //请求号用于查询修改结果
	SettleAccountType string `json:"settleAccountType"`         //结算账户类型 1-对公、2-法人对私、3-银通帐号
	OpenBankName      string `json:"openBankName,omitempty"`    //开户银行
	BranchBankName    string `json:"branchBankName,omitempty"`  //支行名称
	AccountName       string `json:"accountName"`               //账户名称
	BankCardNo        string `json:"bankCardNo"`                //银行卡卡号
	SettleMode        string `json:"settleMode,omitempty"`      //结算模式1-委托付款、 2-API（手动出款）
	SettleModeValue   string `json:"settleModeValue,omitempty"` //A1-15:30 A2-16:30 B1- 09:00|13:00|17:00
	NotifyUrl         string `json:"notifyUrl"`                 //异步通知地址
}

type HlbModifyFeeInfoConf struct {
	AccountNo    string        `json:"accountNo"`         //商户登陆账号
	RequestNo    string        `json:"request_no"`        //请求号用于查询修改结果
	FeeInfo      []*HlbFeeInfo `json:"feeInfo,omitempty"` //费率信息
	FeeStartTime string        `json:"feeStartTime"`      //费率生效开始时间
	FeeEndTime   string        `json:"feeEndTime"`        //费率生效结束时间
	NotifyUrl    string        `json:"notifyUrl"`         //异步通知地址
}

type HlbModifyReturn struct {
	ResultCode  string `json:"resultCode"`
	ErrorCode   string `json:"errorCode"`
	ErrCodeDesc string `json:"errCodeDesc"`
	Status      string `json:"status"`
	Ext         string `json:"ext"`
}

type HlbQueryConf struct {
	AccountNo string `json:"accountNo"` //商户登陆账号
}

type HlbQueryReturn struct {
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

type HlbUploadConf struct {
	MerchantNo     string `json:"merchantNo"`
	OrderNo        string `json:"orderNo"`
	CredentialType string `json:"credentialType"`
	FileSign       string `json:"fileSign"`
	File           []byte
}

type HlbUploadReturn struct {
	MerchantNo       string `json:"merchantNo"`
	OrderNo          string `json:"orderNo"`
	CredentialType   string `json:"orderNo"`
	CredentialStatus string `json:"credentialStatus"`
}

func (hlb *HlbEnter) InitConfig(config *HlbConfig) {
	hlb.HlbConfig = config
}

func (hlb *HlbEnter) Create(config *HlbCreateConf) (*HlbCreateReturn, error) {
	hlb.InterfaceName = "register"

	if config.MerchantType == "PERSON" {
		config.SignName = "商户_" + config.LegalPerson
	}
	hlb.buildData(config)
	hlb.setSign()

	ret, err := hlb.sendReq(HlbEnterUrl)
	if err != nil {
		return nil, err
	}

	ret, err = hlb.retData(ret)
	if err != nil {
		return nil, err
	}

	enterReturn := new(HlbCreateReturn)
	err = json.Unmarshal(ret, &enterReturn)
	if err != nil {
		return nil, err
	}

	return enterReturn, nil
}

func (hlb *HlbEnter) OpenProduct(config *HlbOpenProductConf) error {
	hlb.InterfaceName = "openProduct"
	config.FirstClassMerchantNo = hlb.MerchantNo
	hlb.buildData(config)
	hlb.setSign()

	ret, err := hlb.sendReq(HlbEnterUrl)
	if err != nil {
		return err
	}

	_, err = hlb.retData(ret)
	if err != nil {
		return err
	}
	return nil
}

func (hlb *HlbEnter) Modify(config *HlbModifyConf) (*HlbModifyReturn, error) {
	hlb.InterfaceName = "merchant.enter.modify"
	hlb.buildData(config)
	hlb.setSign()

	ret, err := hlb.sendReq(HlbEnterUrl)
	if err != nil {
		return nil, err
	}

	ret, err = hlb.retData(ret)
	if err != nil {
		return nil, err
	}

	modifyReturn := new(HlbModifyReturn)
	err = json.Unmarshal(ret, &modifyReturn)
	if err != nil {
		return nil, err
	}

	return modifyReturn, nil
}
func (hlb *HlbEnter) ModifySettleInfo(config *HlbModifySettleInfoConf) (*HlbModifyReturn, error) {
	hlb.InterfaceName = "merchant.enter.modify.settleInfo"
	hlb.buildData(config)
	hlb.setSign()

	ret, err := hlb.sendReq(HlbEnterUrl)
	if err != nil {
		return nil, err
	}

	ret, err = hlb.retData(ret)
	if err != nil {
		return nil, err
	}

	modifyReturn := new(HlbModifyReturn)
	err = json.Unmarshal(ret, &modifyReturn)
	if err != nil {
		return nil, err
	}

	return modifyReturn, nil
}
func (hlb *HlbEnter) ModifyFeeInfo(config *HlbModifyFeeInfoConf) (*HlbModifyReturn, error) {
	hlb.InterfaceName = "merchant.enter.modify.feeInfo"
	hlb.buildData(config)
	hlb.setSign()

	ret, err := hlb.sendReq(HlbEnterUrl)
	if err != nil {
		return nil, err
	}

	ret, err = hlb.retData(ret)
	if err != nil {
		return nil, err
	}

	modifyReturn := new(HlbModifyReturn)
	err = json.Unmarshal(ret, &modifyReturn)
	if err != nil {
		return nil, err
	}

	return modifyReturn, nil
}
func (hlb *HlbEnter) Query(config *HlbQueryConf) (*HlbQueryReturn, error) {
	hlb.InterfaceName = "merchant.enter.query"
	hlb.buildData(config)
	hlb.setSign()

	ret, err := hlb.sendReq(HlbEnterUrl)
	if err != nil {
		return nil, err
	}

	ret, err = hlb.retData(ret)
	if err != nil {
		return nil, err
	}

	queryReturn := new(HlbQueryReturn)
	err = json.Unmarshal(ret, queryReturn)
	if err != nil {
		return nil, err
	}

	return queryReturn, nil
}

func (hlb *HlbEnter) Upload(config *HlbUploadConf) (*HlbUploadReturn, error) {
	hlb.InterfaceName = "uploadCredential"
	hlb.buildData(config)
	hlb.setSign()

	ret, err := hlb.sendUploadReq(HlbUploadUrl, config)
	if err != nil {
		return nil, err
	}

	ret, err = hlb.retData(ret)
	if err != nil {
		return nil, err
	}

	uploadReturn := new(HlbUploadReturn)

	err = json.Unmarshal(ret, &uploadReturn)
	if err != nil {
		return nil, err
	}

	return uploadReturn, nil
}
func (hlb *HlbEnter) buildData(config interface{}) error {
	b, err := json.Marshal(config)

	if err != nil {
		return err
	}

	encryptData, _ := helper.TripleEcbDesEncrypt([]byte(string(b)), []byte(hlb.Key))

	hlb.Body = base64.StdEncoding.EncodeToString(encryptData)

	return nil
}
func (hlb *HlbEnter) setSign() {

	hlb.Sign = helper.Md5(hlb.Body + "&" + hlb.MerchantNo + "&" + hlb.SignKey)

	fmt.Println(hlb.Body + "&" + hlb.MerchantNo + "&" + hlb.SignKey)
	fmt.Println(hlb.Sign)
}

func (hlb *HlbEnter) setUploadSign(imageNo string) {
}

func (hlb *HlbEnter) sendReq(reqUrl string) (b []byte, err error) {

	values := url.Values{}

	values.Add("body", hlb.Body)
	values.Add("merchantNo", hlb.MerchantNo)
	values.Add("sign", hlb.Sign)
	values.Add("interfaceName", hlb.InterfaceName)

	rsp, err := http.Client.PostForm(reqUrl, values)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	return
}
func (hlb *HlbEnter) sendUploadReq(reqUrl string, params *HlbUploadConf) (b []byte, err error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("body", hlb.Body)
	writer.WriteField("merchantNo", hlb.MerchantNo)
	writer.WriteField("sign", hlb.Sign)
	writer.WriteField("interfaceName", hlb.InterfaceName)

	//关键的一步操作
	part, err := writer.CreateFormFile("file", "file.png")

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

	rsp, err := http.Client.Post(reqUrl, contentType, body)

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
func (hlb *HlbEnter) retData(ret []byte) (retData []byte, err error) {

	var baseReturn struct {
		Success bool   `json:"success"`
		Code    string `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
		Sign    string `json:"sign"`
	}

	err = json.Unmarshal(ret, &baseReturn)

	if err != nil {
		return
	}

	if !baseReturn.Success {
		err = errors.New(baseReturn.Message)
		return
	}

	if baseReturn.Code != "0000" {
		err = errors.New(baseReturn.Message)
		return
	}

	//MD5值验证
	stringMd5 := helper.Md5(baseReturn.Data + "&" + hlb.SignKey)

	if stringMd5 != baseReturn.Sign {
		fmt.Print("MD5值验证不通过，远程MD5值：" + baseReturn.Sign + ",计算出MD5值：" + stringMd5)
		return nil, errors.New("MD5值验证不通过")
	}

	b, _ := base64.StdEncoding.DecodeString(baseReturn.Data)

	retData, err = helper.TripleEcbDesDecrypt(b, []byte(hlb.Key))

	if err != nil {
		return
	}

	return
}

// 图片信息
type HlbComposeAgreementConfig struct {
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
