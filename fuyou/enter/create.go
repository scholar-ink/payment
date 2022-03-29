package enter

import (
	"encoding/json"
	"github.com/scholar-ink/payment/helper"
)

type CreateEnter struct {
	*CreateConf
	BaseCharge
}

type CreateConf struct {
	LoginNo                  string      `json:"loginNo"`                            //商户登陆账号
	AccountType              string      `json:"accountType"`                        //开户类型
	MerchantName             string      `json:"merchantName"`                       //商户名称
	MerchantType             string      `json:"merchantType"`                       //商户类型
	CompanyIdType            string      `json:"companyIdType,omitempty"`            //公司证件类型
	MccCode                  string      `json:"mccCode"`                            //mcc码
	MerchantMail             string      `json:"merchantMail"`                       //商户邮箱
	SendMail                 string      `json:"sendMail,omitempty"`                 //是否发送邮件
	PhoneNo                  string      `json:"phoneNo"`                            //联系人手机号
	TaxNo                    string      `json:"taxNo,omitempty"`                    //税务登记号
	TaxNoDeadline            string      `json:"taxNoDeadline,omitempty"`            //税务登记号期限
	OrganizationCode         string      `json:"organizationCode,omitempty"`         //组织机构代码
	OrganizationCodeDeadline string      `json:"organizationCodeDeadline,omitempty"` //组织机构代码有效期
	LicenseNo                string      `json:"licenseNo,omitempty"`                //营业执照号
	LicenseDeadline          string      `json:"licenseDeadline,omitempty"`          //营业执照号有效期
	LicenseProvince          string      `json:"licenseProvince,omitempty"`          //营业执照号登记省
	LicenseCity              string      `json:"licenseCity,omitempty"`              //营业执照号登记市
	LicenseDistrict          string      `json:"licenseDistrict,omitempty"`          //营业执照号登记区
	ManagementArea           string      `json:"managementArea,omitempty"`           //商户经营范围
	OpeningAccountsDeadline  string      `json:"openingAccountsDeadline,omitempty"`  //开户许可证发证日期
	LegalPersonName          string      `json:"legalPersonName"`                    //法人名称
	LegalPersonId            string      `json:"legalPersonId"`                      //法人身份证号
	LegalPersonIdDeadline    string      `json:"legalPersonIdDeadline"`              //法人身份证有效期
	OperatorName             string      `json:"operatorName"`                       //经办人姓名
	OperatorId               string      `json:"operatorId"`                         //经办身份证号
	OperatorIdDeadline       string      `json:"operatorIdDeadline"`                 //经办人身份证有效期
	MerchantContactsName     string      `json:"merchantContactsName"`               //商户联系人
	IcpNo                    string      `json:"icpNo,omitempty"`                    //ICP备案号
	DomainName               string      `json:"domainName,omitempty"`               //域名/APP/小程序/公众号
	MerchantAddress          string      `json:"merchantAddress"`                    //商户地址
	ServicePhone             string      `json:"servicePhone"`                       //客服电话
	MerchantShortName        string      `json:"merchantShortName,omitempty"`        //商户简称
	PicInfo                  *PicInfo    `json:"picInfo"`                            //图片信息
	SettleInfo               *SettleInfo `json:"settleInfo"`                         //结算信息
	FeeInfo                  []*FeeInfo  `json:"feeInfo,omitempty"`                  //费率信息
	FeeStartTime             string      `json:"feeStartTime"`                       //费率生效开始时间
	FeeEndTime               string      `json:"feeEndTime"`                         //费率生效结束时间
	NotifyUrl                string      `json:"notifyUrl"`                          //异步通知地址
}

// 图片信息
type PicInfo struct {
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
type SettleInfo struct {
	SettleAccountType string `json:"settleAccountType"`         //结算账户类型 1-对公、2-法人对私、3-银通帐号
	OpenBankName      string `json:"openBankName,omitempty"`    //开户银行
	BranchBankName    string `json:"branchBankName,omitempty"`  //支行名称
	AccountName       string `json:"accountName"`               //账户名称
	BankCardNo        string `json:"bankCardNo"`                //银行卡卡号
	SettleMode        string `json:"settleMode,omitempty"`      //结算模式1-委托付款、 2-API（手动出款）
	SettleModeValue   string `json:"settleModeValue,omitempty"` //A1-15:30 A2-16:30 B1- 09:00|13:00|17:00
}

//费率信息
type FeeInfo struct {
	ProductType string `json:"productType"` //产品类型
	FeeRate     string `json:"feeRate"`     //比例收取千分位：1=0.001 固定金额：单位：分
	FeeType     string `json:"feeType"`     //0-比例收取,1-固定金额
	ChargeRole  string `json:"chargeRole"`  //收费角色1-收款方、 2-付款方、 3-平台方
	SettleCycle string `json:"settleCycle"` //结算周期
}

type CreateReturn struct {
	ResultCode  string `json:"resultCode"`
	ErrorCode   string `json:"errorCode"`
	ErrCodeDesc string `json:"errCodeDesc"`
	AccountNo   string `json:"accountNo"`
	LoginNo     string `json:"loginNo"`
	Status      string `json:"status"`
}

func (en *CreateEnter) Handle(conf *CreateConf) (*CreateReturn, error) {
	err := en.BuildData(conf)
	if err != nil {
		return nil, err
	}
	en.SetSign()
	ret, err := en.SendReq(EnterUrl, en)
	if err != nil {
		return nil, err
	}
	return en.RetData(ret)
}

func (en *CreateEnter) RetData(ret []byte) (*CreateReturn, error) {

	ret, err := en.BaseCharge.RetData(ret)

	if err != nil {
		return nil, err
	}

	enterReturn := new(CreateReturn)

	err = json.Unmarshal(ret, &enterReturn)

	if err != nil {
		return nil, err
	}

	return enterReturn, nil
}

func (en *CreateEnter) BuildData(conf *CreateConf) error {

	b, err := json.Marshal(conf)

	if err != nil {
		return err
	}

	en.CreateConf = conf

	en.ServiceName = "merchant.enter.create"

	encryptData, err := helper.Rsa1Encrypt(en.PfxData, b, en.CertPassWord)

	if err != nil {
		return err
	}

	en.EncryptData = encryptData

	return nil
}
