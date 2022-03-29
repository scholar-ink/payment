package enter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"
	"io"
	"io/ioutil"
	"mime/multipart"
)

const (
	ZgUploadUrl = "http://mgateway.g-pay.cn/image/upload.do"
	ZgEnterUrl  = "http://mgateway.g-pay.cn/merchant/enter/api.do"
)

type ZgEnter struct {
	*ZgConfig
	ServiceName  string `json:"serviceName"`
	Charset      string `json:"charset"`
	Version      string `json:"version"`
	EncryptType  string `json:"encryptType"`
	EncryptData  string `json:"encryptData"`
	SignData     string `json:"signData"`
	ResponseCode string `xml:"responseCode" json:"responseCode"`
	ResponseMsg  string `xml:"responseMsg" json:"responseMsg"`
}

type ZgConfig struct {
	AgentNo      string `json:"agentNo"`
	Key          string `json:"key"`
	PfxData      []byte `json:"pfx_data"`
	CertPassWord string `json:"certPassWord"`
}

// 图片信息
type ZgPicInfo struct {
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
type ZgSettleInfo struct {
	SettleAccountType string `json:"settleAccountType"`         //结算账户类型 1-对公、2-法人对私、3-银通帐号
	OpenBankName      string `json:"openBankName,omitempty"`    //开户银行
	BranchBankName    string `json:"branchBankName,omitempty"`  //支行名称
	AccountName       string `json:"accountName"`               //账户名称
	BankCardNo        string `json:"bankCardNo"`                //银行卡卡号
	SettleMode        string `json:"settleMode,omitempty"`      //结算模式1-委托付款、 2-API（手动出款）
	SettleModeValue   string `json:"settleModeValue,omitempty"` //A1-15:30 A2-16:30 B1- 09:00|13:00|17:00
}

//费率信息
type ZgFeeInfo struct {
	ProductType string `json:"productType"` //产品类型
	FeeRate     string `json:"feeRate"`     //比例收取千分位：1=0.001 固定金额：单位：分
	FeeType     string `json:"feeType"`     //0-比例收取,1-固定金额
	ChargeRole  string `json:"chargeRole"`  //收费角色1-收款方、 2-付款方、 3-平台方
	SettleCycle string `json:"settleCycle"` //结算周期
}

type ZgCreateConf struct {
	LoginNo                  string        `json:"loginNo"`                            //商户登陆账号
	AccountType              string        `json:"accountType"`                        //开户类型
	MerchantName             string        `json:"merchantName"`                       //商户名称
	MerchantType             string        `json:"merchantType"`                       //商户类型
	CompanyIdType            string        `json:"companyIdType,omitempty"`            //公司证件类型
	MccCode                  string        `json:"mccCode"`                            //mcc码
	MerchantMail             string        `json:"merchantMail"`                       //商户邮箱
	SendMail                 string        `json:"sendMail,omitempty"`                 //是否发送邮件
	PhoneNo                  string        `json:"phoneNo"`                            //联系人手机号
	TaxNo                    string        `json:"taxNo,omitempty"`                    //税务登记号
	TaxNoDeadline            string        `json:"taxNoDeadline,omitempty"`            //税务登记号期限
	OrganizationCode         string        `json:"organizationCode,omitempty"`         //组织机构代码
	OrganizationCodeDeadline string        `json:"organizationCodeDeadline,omitempty"` //组织机构代码有效期
	LicenseNo                string        `json:"licenseNo,omitempty"`                //营业执照号
	LicenseDeadline          string        `json:"licenseDeadline,omitempty"`          //营业执照号有效期
	LicenseProvince          string        `json:"licenseProvince,omitempty"`          //营业执照号登记省
	LicenseCity              string        `json:"licenseCity,omitempty"`              //营业执照号登记市
	LicenseDistrict          string        `json:"licenseDistrict,omitempty"`          //营业执照号登记区
	ManagementArea           string        `json:"managementArea,omitempty"`           //商户经营范围
	OpeningAccountsDeadline  string        `json:"openingAccountsDeadline,omitempty"`  //开户许可证发证日期
	LegalPersonName          string        `json:"legalPersonName"`                    //法人名称
	LegalPersonId            string        `json:"legalPersonId"`                      //法人身份证号
	LegalPersonIdDeadline    string        `json:"legalPersonIdDeadline"`              //法人身份证有效期
	OperatorName             string        `json:"operatorName"`                       //经办人姓名
	OperatorId               string        `json:"operatorId"`                         //经办身份证号
	OperatorIdDeadline       string        `json:"operatorIdDeadline"`                 //经办人身份证有效期
	MerchantContactsName     string        `json:"merchantContactsName"`               //商户联系人
	IcpNo                    string        `json:"icpNo,omitempty"`                    //ICP备案号
	DomainName               string        `json:"domainName,omitempty"`               //域名/APP/小程序/公众号
	MerchantAddress          string        `json:"merchantAddress"`                    //商户地址
	ServicePhone             string        `json:"servicePhone"`                       //客服电话
	MerchantShortName        string        `json:"merchantShortName,omitempty"`        //商户简称
	PicInfo                  *ZgPicInfo    `json:"picInfo"`                            //图片信息
	SettleInfo               *ZgSettleInfo `json:"settleInfo"`                         //结算信息
	FeeInfo                  []*ZgFeeInfo  `json:"feeInfo,omitempty"`                  //费率信息
	FeeStartTime             string        `json:"feeStartTime"`                       //费率生效开始时间
	FeeEndTime               string        `json:"feeEndTime"`                         //费率生效结束时间
	NotifyUrl                string        `json:"notifyUrl"`                          //异步通知地址
}

type ZgCreateReturn struct {
	ResultCode  string `json:"resultCode"`
	ErrorCode   string `json:"errorCode"`
	ErrCodeDesc string `json:"errCodeDesc"`
	AccountNo   string `json:"accountNo"`
	LoginNo     string `json:"loginNo"`
	Status      string `json:"status"`
}

type ZgModifyConf struct {
	AccountNo    string        `json:"accountNo"`    //商户登陆账号
	FeeInfo      []*ZgFeeInfo  `json:"feeInfo"`      //费率信息
	FeeStartTime string        `json:"feeStartTime"` //费率生效开始时间
	FeeEndTime   string        `json:"feeEndTime"`   //费率生效结束时间
	SettleInfo   *ZgSettleInfo `json:"settle_info"`  //结算信息
}

type ZgModifySettleInfoConf struct {
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

type ZgModifyFeeInfoConf struct {
	AccountNo    string       `json:"accountNo"`         //商户登陆账号
	RequestNo    string       `json:"request_no"`        //请求号用于查询修改结果
	FeeInfo      []*ZgFeeInfo `json:"feeInfo,omitempty"` //费率信息
	FeeStartTime string       `json:"feeStartTime"`      //费率生效开始时间
	FeeEndTime   string       `json:"feeEndTime"`        //费率生效结束时间
	NotifyUrl    string       `json:"notifyUrl"`         //异步通知地址
}

type ZgModifyReturn struct {
	ResultCode  string `json:"resultCode"`
	ErrorCode   string `json:"errorCode"`
	ErrCodeDesc string `json:"errCodeDesc"`
	Status      string `json:"status"`
	Ext         string `json:"ext"`
}

type ZgQueryConf struct {
	AccountNo string `json:"accountNo"` //商户登陆账号
}

type ZgQueryReturn struct {
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

type ZgUploadConf struct {
	ImageData []byte `json:"image_data"`
	ImageNo   string `json:"image_no"`
}

type ZgUploadReturn struct {
	ResultCode  string `json:"resultCode"`
	ErrorCode   string `json:"errorCode"`
	ErrCodeDesc string `json:"errCodeDesc"`
	ImageNo     string `json:"imageNo"`
	ImagePath   string `json:"imagePath"`
	Ext         string `json:"ext"`
}

func (zg *ZgEnter) InitConfig(config *ZgConfig) {
	zg.Version = "3.0"
	zg.Charset = "UTF-8"
	zg.EncryptType = "RSA"
	zg.ZgConfig = config
}

func (zg *ZgEnter) Create(config *ZgCreateConf) (*ZgCreateReturn, error) {
	zg.ServiceName = "merchant.enter.create"
	zg.buildData(config)
	zg.setSign()

	ret, err := zg.sendReq(ZgEnterUrl, zg)
	if err != nil {
		return nil, err
	}

	ret, err = zg.retData(ret)
	if err != nil {
		return nil, err
	}

	enterReturn := new(ZgCreateReturn)
	err = json.Unmarshal(ret, &enterReturn)
	if err != nil {
		return nil, err
	}

	return enterReturn, nil
}

func (zg *ZgEnter) Modify(config *ZgModifyConf) (*ZgModifyReturn, error) {
	zg.ServiceName = "merchant.enter.modify"
	zg.buildData(config)
	zg.setSign()

	ret, err := zg.sendReq(ZgEnterUrl, zg)
	if err != nil {
		return nil, err
	}

	ret, err = zg.retData(ret)
	if err != nil {
		return nil, err
	}

	modifyReturn := new(ZgModifyReturn)
	err = json.Unmarshal(ret, &modifyReturn)
	if err != nil {
		return nil, err
	}

	return modifyReturn, nil
}
func (zg *ZgEnter) ModifySettleInfo(config *ZgModifySettleInfoConf) (*ZgModifyReturn, error) {
	zg.ServiceName = "merchant.enter.modify.settleInfo"
	zg.buildData(config)
	zg.setSign()

	ret, err := zg.sendReq(ZgEnterUrl, zg)
	if err != nil {
		return nil, err
	}

	ret, err = zg.retData(ret)
	if err != nil {
		return nil, err
	}

	modifyReturn := new(ZgModifyReturn)
	err = json.Unmarshal(ret, &modifyReturn)
	if err != nil {
		return nil, err
	}

	return modifyReturn, nil
}
func (zg *ZgEnter) ModifyFeeInfo(config *ZgModifyFeeInfoConf) (*ZgModifyReturn, error) {
	zg.ServiceName = "merchant.enter.modify.feeInfo"
	zg.buildData(config)
	zg.setSign()

	ret, err := zg.sendReq(ZgEnterUrl, zg)
	if err != nil {
		return nil, err
	}

	ret, err = zg.retData(ret)
	if err != nil {
		return nil, err
	}

	modifyReturn := new(ZgModifyReturn)
	err = json.Unmarshal(ret, &modifyReturn)
	if err != nil {
		return nil, err
	}

	return modifyReturn, nil
}
func (zg *ZgEnter) Query(config *ZgQueryConf) (*ZgQueryReturn, error) {
	zg.ServiceName = "merchant.enter.query"
	zg.buildData(config)
	zg.setSign()

	ret, err := zg.sendReq(ZgEnterUrl, zg)
	if err != nil {
		return nil, err
	}

	ret, err = zg.retData(ret)
	if err != nil {
		return nil, err
	}

	queryReturn := new(ZgQueryReturn)
	err = json.Unmarshal(ret, queryReturn)
	if err != nil {
		return nil, err
	}

	return queryReturn, nil
}

func (zg *ZgEnter) Upload(config *ZgUploadConf) (*ZgUploadReturn, error) {
	zg.ServiceName = "merchant.enter.upload"
	zg.buildData(config)
	zg.setUploadSign(config.ImageNo)

	ret, err := zg.sendUploadReq(ZgUploadUrl, config)
	if err != nil {
		return nil, err
	}

	var baseReturn ZgEnter
	err = json.Unmarshal(ret, &baseReturn)
	if err != nil {
		return nil, err
	}

	uploadReturn := new(ZgUploadReturn)
	err = json.Unmarshal([]byte(baseReturn.EncryptData), &uploadReturn)
	if err != nil {
		return nil, err
	}

	return uploadReturn, nil
}
func (zg *ZgEnter) buildData(config interface{}) error {
	b, err := json.Marshal(config)

	if err != nil {
		return err
	}

	if zg.PfxData != nil {
		encryptData, err := helper.Rsa1Encrypt(zg.PfxData, b, zg.CertPassWord)

		if err != nil {
			return err
		}

		zg.EncryptData = encryptData
	}
	return nil
}
func (zg *ZgEnter) setSign() {
	zg.SignData = helper.Md5(zg.EncryptData + zg.Key)
}

func (zg *ZgEnter) setUploadSign(imageNo string) {
	zg.SignData = helper.Md5(zg.AgentNo + zg.Version + imageNo + zg.Key)
}

func (zg *ZgEnter) sendReq(reqUrl string, params interface{}) (b []byte, err error) {

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
func (zg *ZgEnter) sendUploadReq(reqUrl string, params *ZgUploadConf) (b []byte, err error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("agentNo", zg.AgentNo)
	writer.WriteField("serviceName", zg.ServiceName)
	writer.WriteField("version", zg.Version)
	writer.WriteField("imageNo", params.ImageNo)
	writer.WriteField("signData", zg.SignData)

	//关键的一步操作
	part, err := writer.CreateFormFile("imageData", "file.png")

	if err != nil {
		return nil, err
	}

	//iocopy
	_, err = io.Copy(part, bytes.NewReader(params.ImageData))

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
func (zg *ZgEnter) retData(ret []byte) (retData []byte, err error) {
	var baseReturn ZgEnter

	err = json.Unmarshal(ret, &baseReturn)

	if err != nil {
		return
	}

	if baseReturn.ResponseCode != "" {
		err = errors.New(baseReturn.ResponseMsg)
		return
	}

	//MD5值验证
	stringMd5 := helper.Md5(baseReturn.EncryptData + zg.Key)

	if stringMd5 != baseReturn.SignData {
		fmt.Print("MD5值验证不通过，远程MD5值：" + baseReturn.SignData + ",计算出MD5值：" + stringMd5)
		return nil, errors.New("MD5值验证不通过")
	}

	retData, err = helper.Rsa1Decrypt(zg.PfxData, baseReturn.EncryptData, zg.CertPassWord)

	return
}

// 图片信息
type ZgComposeAgreementConfig struct {
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
