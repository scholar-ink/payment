package enter

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"
	"io/ioutil"
	strings2 "strings"
	"time"
)

const (
	FqlEnterUrl = "https://eagateway.fuqianla.net/gateway/api/clientApi/interface"
)

type FqlEnter struct {
	*FqlConfig
	VersionNo string      `json:"versionNo"`
	ServiceId string      `json:"serviceId"`
	SignType  string      `json:"signType"`
	Sign      string      `json:"sign"`
	Timestamp string      `json:"timestamp"`
	SeqNo     string      `json:"seqNo"`
	Data      interface{} `json:"-"`
	DataStr   string      `json:"data"`
}

type FqlReturn struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

type FqlConfig struct {
	MerchId string `json:"merchId"`
	PriKey  string `json:"-"`
}

type FqlCreateConf struct {
	FilingMerchantId        string `json:"filing_merchant_id"`          //商户内部 ID
	FullName                string `json:"full_name"`                   //商户全称
	ShortName               string `json:"short_name"`                  //商户简称
	ServicePhone            string `json:"service_phone"`               //客服电话
	Category                string `json:"category"`                    //经营类目
	Address                 string `json:"address"`                     //商户经营地址
	Province                string `json:"province"`                    //省份
	City                    string `json:"city"`                        //城市
	County                  string `json:"county"`                      //区县
	MerchantType            string `json:"merchant_type"`               //商户主体类型
	LinenceType             string `json:"linence_type"`                //商户证书类型
	LinenceNo               string `json:"linence_no"`                  //证书编号
	CompanyAddress          string `json:"company_address"`             //注册地址
	LicenseValidDate        string `json:"license_valid_date"`          //营业执照有效期
	LicensePic              string `json:"license_pic"`                 //营业执照图片
	MicroType               string `json:"micro_type,omitempty"`        //小微经营类型
	ShopName                string `json:"shop_name,omitempty"`         //门店名称
	ShopAddress             string `json:"shop_address,omitempty"`      //门店地址
	ShopEntrancePic         string `json:"shop_entrance_pic,omitempty"` //门店门头照片
	IndoorPic               string `json:"indoor_pic,omitempty"`        //店内环境照片
	LegalPersonCemType      string `json:"legal_person_cem_type"`       //法人证件类型
	LegalPersonName         string `json:"legal_person_name"`           //法人证件姓名
	LegalPersonCemNum       string `json:"legal_person_cem_num"`        //法人证件号码
	LegalPersonValidDate    string `json:"legal_person_valid_date"`     //法人证件有效日期
	IdNomalPic              string `json:"id_nomal_pic"`                //证件正面照片
	IdReversePic            string `json:"id_reverse_pic"`              //证件反面照片
	ContactName             string `json:"contact_name"`                //联系人
	ContactMobile           string `json:"contact_mobile"`              //联系人手机
	ContactEmail            string `json:"contact_email"`               //联系人邮箱
	ContactCertfiticateNum  string `json:"contact_certfiticate_num"`    //联系人证件号码
	ContactCertfiticateType string `json:"contact_certfiticate_type"`   //联系人证件类型
	CardNum                 string `json:"card_num"`                    //结算银行卡号
	UserName                string `json:"user_name"`                   //结算银行卡开户人
	CardNumType             string `json:"card_num_type"`               //结算卡账户类型
	CardBankNumber          string `json:"card_bank_number"`            //结算卡银行联行号
	ProductCode             string `json:"product_code"`                //拟支付产品编码
	UseScenes               string `json:"use_scenes"`                  //使用场景
}
type FqlCreateReturn struct {
	TradeMerchantNo string `json:"trade_merchant_no"`
	BizCode         string `json:"biz_code"`
	BizMsg          string `json:"biz_msg"`
}

type FqlSendValidateCodeConf struct {
	OrderId string `json:"orderId"`
	PhoneNo string `json:"phoneNo"`
}

type FqlSendValidateCodeReturn struct {
	Status  string `json:"status"`
	SubCode string `json:"subCode"`
	SubMsg  string `json:"subMsg"`
}

type FqlOpenAccAuthEnConf struct {
	OrderId     string `json:"orderId"`
	OrglOrderId string `json:"orglOrderId"`
	ValidType   string `json:"validType"`
	Amount      string `json:"amount"`
	ProvingType string `json:"provingType"`
}

type FqlOpenAccAuthEnReturn struct {
	ConfirmStatus string `json:"confirmStatus"`
	SubCode       string `json:"subCode"`
	SubMsg        string `json:"subMsg"`
}

type FqlOpenAccConf struct {
	OrderId          string `json:"orderId"`          //商户订单号
	ChUserId         string `json:"chUserId"`         //商户用户号
	ParentCustomerId string `json:"parentCustomerId"` //父级客户号
	CustomerFlag     string `json:"customerFlag"`     //开户客户标识 01-代理商，02-交易平台+代理商，03-交易平台，04-小商户
	CorpName         string `json:"corpName"`         //客户名称
	ComType          string `json:"comType"`          //企业主体类型 01企业 02个体工商户
	SettleAccType    string `json:"settleAccType"`    //绑卡类型 1.对公 2.对私
	CertificateNo    string `json:"certificateNo"`    //营业执照号
	ResName          string `json:"resName"`          //法人姓名
	CorpCertNo       string `json:"corpCertNo"`       //法人身份证号
	CorpContactNo    string `json:"corpContactNo"`    //法人手机号
	ContactNo        string `json:"contactNo"`        //企业联系电话
	BankAccNo        string `json:"bankAccNo"`        //银行账号
	BankAccName      string `json:"bankAccName"`      //银行账户名称
	BankAccProvince  string `json:"bankAccProvince"`  //开户行所在省份
	BankAccCity      string `json:"bankAccCity"`      //开户行所在城市
	BankAccAddress   string `json:"bankAccAddress"`   //开户行名称
	BankNo           string `json:"bankNo"`           //银行联行号
	BusLicEcTypeUrl  string `json:"busLicEctypeUrl"`  //营业执照副本url
	PerInfoPageUrl   string `json:"perInfoPageUrl"`   //身份证个人信息页url
	NatEmbPageUrl    string `json:"natEmbPageUrl"`    //身份证国徽页url
	MessageCode      string `json:"messageCode"`      //短信验证码
	Password         string `json:"password"`         //交易密码
}

type FqlOpenAccReturn struct {
	Status     string `json:"status"`
	SubCode    string `json:"subCode"`
	SubMsg     string `json:"subMsg"`
	CustomerId string `json:"customerId"`
}

type FqlOpenAccQueryConf struct {
	OrderId  string `json:"orglOrderId"` //商户订单号
	ChUserId string `json:"chUserId"`    //商户用户号
}

type FqlOpenAccQueryReturn struct {
	Status     string `json:"status"`
	SubCode    string `json:"subCode"`
	SubMsg     string `json:"subMsg"`     //交易状态
	AccDate    string `json:"accDate"`    //开户日期
	ExAccNo    string `json:"exAccNo"`    //虚户账号
	CustomerId string `json:"customerId"` //开户客户号
}

type FqlPersonalOpenConf struct {
	OrderId          string `json:"orderId"`            //商户订单号
	ChUserId         string `json:"chUserId,omitempty"` //商户用户号
	ParentCustomerId string `json:"parentCustomerId"`   //父级客户号
	CustomerFlag     string `json:"customerFlag"`       //开户客户标识
	CertificateNo    string `json:"certificateNo"`      //证件号码
	PhoneNo          string `json:"phoneNo"`            //客户手机号码
	CorpName         string `json:"corpName"`           //客户名称
	NatEmbPageUrl    string `json:"natEmbPageUrl"`      //身份证国徽页url
	PerInfoPageUrl   string `json:"perInfoPageUrl"`     //身份证个人信息页url
	BankAccNo        string `json:"bankAccNo"`          //银行账号
	BankAccName      string `json:"bankAccName"`        //银行账户名称
	BankAccAddress   string `json:"bankAccAddress"`     //开户行名称
	BankNo           string `json:"bankNo"`             //银行联行号
	MessageCode      string `json:"messageCode"`        //短信验证码
	Password         string `json:"password"`           //交易密码
}

type FqlPersonalOpenReturn struct {
	Status        string `json:"status"`
	SubCode       string `json:"subCode"`
	SubMsg        string `json:"subMsg"`
	CustomerId    string `json:"customerId"`    //开户客户号
	ExAccNo       string `json:"exAccNo"`       //虚户账号
	ExBankAccName string `json:"exBankAccName"` //虚户账户名
}

type FqlOpenScanConf struct {
	OrderId           string `json:"orderId"`           //商户订单号
	ChUserId          string `json:"chUserId"`          //商户用户号
	CustomerId        string `json:"customerId"`        //商户客户号
	OperatingProvince string `json:"operatingProvince"` //经营地址省
	OperatingCity     string `json:"operatingCity"`     //经营地址市
	OperatingDistrict string `json:"operatingDistrict"` //经营地址区
	OperatingAddress  string `json:"operatingAddress"`  //经营详细地址
	TradeCategory     string `json:"tradeCategory"`     //经营详细地址
	ShopSignUrl       string `json:"shopSignUrl"`       //门头照照片url
	PremisesUrl       string `json:"premisesUrl"`       //营业场所照片url
	PmtType           string `json:"pmtType"`           //支付类型
}

type FqlOpenScanReturn struct {
	Status   string `json:"status"`
	SubCode  string `json:"subCode"`
	SubMsg   string `json:"subMsg"`
	ChUserId string `json:"chUserId"`
	OrderId  string `json:"orderId"`
}

type FqlScanQueryResultConf struct {
	OrderId string `json:"orglOrderId"` //商户请求订单号
}

type FqlScanQueryResultReturn struct {
	Status     string `json:"status"`
	SubCode    string `json:"subCode"`
	SubMsg     string `json:"subMsg"`
	AliPayRate string `json:"alipayRate"` //支付宝费率
	WeChatRate string `json:"wechatRate"` //微信费率
	ExAccNo    string `json:"exAccNo"`    //虚户账号
}

type FqlTransferConf struct {
	OrderId     string `json:"orderId"`     //商户订单号
	CustomerId  string `json:"customerId"`  //商户客户号
	BizType     string `json:"bizType"`     //业务类型
	TxnType     string `json:"txnType"`     //业务类型
	PayExAccNo  string `json:"payExAccNo"`  //付款账号
	RecvExAccNo string `json:"recvExAccNo"` //收款账号
	Amount      string `json:"amount"`      //交易金额
	Password    string `json:"password"`    //交易密码
}

type FqlTransferReturn struct {
	Status    string `json:"status"`
	SubCode   string `json:"subCode"`
	SubMsg    string `json:"subMsg"`
	TxnId     string `json:"txnId"`     //系统单号
	WaybillNo string `json:"waybillNo"` //业务单号
}

type FqlWithdrawConf struct {
	OrderId    string `json:"orderId"`    //商户订单号
	CustomerId string `json:"customerId"` //商户客户号
	ExAccNo    string `json:"exAccNo"`    //虚户账号
	TxnType    string `json:"txnType"`    //交易类型
	BankAccNo  string `json:"bankAccNo"`  //收款银行账号
	Amount     string `json:"amount"`     //交易金额
	Password   string `json:"password"`   //交易密码
}

type FqlWithdrawReturn struct {
	Status    string `json:"status"`
	SubCode   string `json:"subCode"`
	SubMsg    string `json:"subMsg"`
	TxnId     string `json:"txnId"`     //系统单号
	WaybillNo string `json:"waybillNo"` //业务单号
}

type FqlUploadConf struct {
	ImgFile string `json:"img_file"` //图片文件
}

type FqlUploadReturn struct {
	ImageId string `json:"image_id"`
	BizCode string `json:"biz_code"`
	BizMsg  string `json:"biz_msg"`
}

func (fql *FqlEnter) InitConfig(config *FqlConfig) {
	fql.SignType = "rsa2"
	fql.FqlConfig = config
}

func (fql *FqlEnter) Create(config *FqlCreateConf) (*FqlCreateReturn, error) {
	fql.ServiceId = "filing.mch.apply"
	fql.Timestamp = time.Now().Format("20060102150405")
	fql.SeqNo = strings2.ReplaceAll(uuid.New().String(), "-", "")
	fql.VersionNo = "2.00"
	fql.buildData(config)
	fql.setSign()

	ret, err := fql.sendReq(FqlEnterUrl)
	if err != nil {
		return nil, err
	}

	ret, err = fql.retData(ret)
	if err != nil {
		return nil, err
	}

	enterReturn := new(FqlCreateReturn)

	err = json.Unmarshal(ret, &enterReturn)

	if err != nil {
		return nil, err
	}

	if enterReturn.BizCode != "B106014" {
		err = errors.New(enterReturn.BizMsg)
		return nil, err
	}

	return enterReturn, nil
}
func (fql *FqlEnter) OpenAccAuthEn(config *FqlOpenAccAuthEnConf) (*FqlOpenAccAuthEnReturn, error) {
	fql.ServiceId = "openAccAuthen"
	fql.Timestamp = time.Now().Format("20060102150405")
	fql.SeqNo = strings2.ReplaceAll(uuid.New().String(), "-", "")
	fql.VersionNo = "2.00"
	fql.buildData(config)
	fql.setSign()

	ret, err := fql.sendReq(FqlEnterUrl)
	if err != nil {
		return nil, err
	}

	ret, err = fql.retData(ret)
	if err != nil {
		return nil, err
	}

	enterReturn := new(FqlOpenAccAuthEnReturn)

	err = json.Unmarshal(ret, &enterReturn)

	if err != nil {
		return nil, err
	}
	return enterReturn, nil
}

func (fql *FqlEnter) SendValidateCode(config *FqlSendValidateCodeConf) (*FqlSendValidateCodeReturn, error) {
	fql.ServiceId = "sendValidateCode"
	fql.Timestamp = time.Now().Format("20060102150405")
	fql.SeqNo = strings2.ReplaceAll(uuid.New().String(), "-", "")
	fql.VersionNo = "2.00"
	fql.buildData(config)
	fql.setSign()

	ret, err := fql.sendReq(FqlEnterUrl)
	if err != nil {
		return nil, err
	}

	ret, err = fql.retData(ret)
	if err != nil {
		return nil, err
	}

	enterReturn := new(FqlSendValidateCodeReturn)

	err = json.Unmarshal(ret, &enterReturn)

	if err != nil {
		return nil, err
	}
	return enterReturn, nil
}
func (fql *FqlEnter) OpenAcc(config *FqlOpenAccConf) (*FqlOpenAccReturn, error) {
	fql.ServiceId = "mzOpenAcc"
	fql.Timestamp = time.Now().Format("20060102150405")
	fql.SeqNo = strings2.ReplaceAll(uuid.New().String(), "-", "")
	fql.VersionNo = "2.00"
	fql.buildData(config)
	fql.setSign()

	ret, err := fql.sendReq(FqlEnterUrl)
	if err != nil {
		return nil, err
	}

	ret, err = fql.retData(ret)
	if err != nil {
		return nil, err
	}

	openReturn := new(FqlOpenAccReturn)

	err = json.Unmarshal(ret, &openReturn)

	if err != nil {
		return nil, err
	}
	return openReturn, nil
}
func (fql *FqlEnter) OpenAccQuery(config *FqlOpenAccQueryConf) (*FqlOpenAccQueryReturn, error) {
	fql.ServiceId = "openAccQuery"
	fql.Timestamp = time.Now().Format("20060102150405")
	fql.SeqNo = strings2.ReplaceAll(uuid.New().String(), "-", "")
	fql.VersionNo = "2.00"
	fql.buildData(config)
	fql.setSign()

	ret, err := fql.sendReq(FqlEnterUrl)
	if err != nil {
		return nil, err
	}

	ret, err = fql.retData(ret)
	if err != nil {
		return nil, err
	}

	openReturn := new(FqlOpenAccQueryReturn)

	err = json.Unmarshal(ret, &openReturn)

	if err != nil {
		return nil, err
	}
	return openReturn, nil
}

func (fql *FqlEnter) PersonalOpen(config *FqlPersonalOpenConf) (*FqlPersonalOpenReturn, error) {
	fql.ServiceId = "mzPersonalOpen"
	fql.Timestamp = time.Now().Format("20060102150405")
	fql.SeqNo = strings2.ReplaceAll(uuid.New().String(), "-", "")
	fql.VersionNo = "2.00"
	fql.buildData(config)
	fql.setSign()

	ret, err := fql.sendReq(FqlEnterUrl)
	if err != nil {
		return nil, err
	}

	ret, err = fql.retData(ret)
	if err != nil {
		return nil, err
	}

	openReturn := new(FqlPersonalOpenReturn)

	err = json.Unmarshal(ret, &openReturn)

	if err != nil {
		return nil, err
	}
	if openReturn.SubCode != "0001" && openReturn.SubCode != "0000" && openReturn.SubCode != "E0001" {
		return nil, errors.New(openReturn.SubMsg)
	}
	return openReturn, nil
}
func (fql *FqlEnter) OpenScan(config *FqlOpenScanConf) (*FqlOpenScanReturn, error) {
	fql.ServiceId = "openScan"
	fql.Timestamp = time.Now().Format("20060102150405")
	fql.SeqNo = strings2.ReplaceAll(uuid.New().String(), "-", "")
	fql.VersionNo = "2.00"
	fql.buildData(config)
	fql.setSign()

	ret, err := fql.sendReq(FqlEnterUrl)
	if err != nil {
		return nil, err
	}

	ret, err = fql.retData(ret)
	if err != nil {
		return nil, err
	}

	openReturn := new(FqlOpenScanReturn)

	err = json.Unmarshal(ret, &openReturn)

	if err != nil {
		return nil, err
	}
	return openReturn, nil
}

func (fql *FqlEnter) ScanQueryResult(config *FqlScanQueryResultConf) (*FqlScanQueryResultReturn, error) {
	fql.ServiceId = "scanQueryResult"
	fql.Timestamp = time.Now().Format("20060102150405")
	fql.SeqNo = strings2.ReplaceAll(uuid.New().String(), "-", "")
	fql.VersionNo = "2.00"
	fql.buildData(config)
	fql.setSign()

	ret, err := fql.sendReq(FqlEnterUrl)
	if err != nil {
		return nil, err
	}

	ret, err = fql.retData(ret)
	if err != nil {
		return nil, err
	}

	openReturn := new(FqlScanQueryResultReturn)

	err = json.Unmarshal(ret, &openReturn)

	if err != nil {
		return nil, err
	}
	return openReturn, nil
}

func (fql *FqlEnter) Transfer(config *FqlTransferConf) (*FqlTransferReturn, error) {
	fql.ServiceId = "transfer"
	fql.Timestamp = time.Now().Format("20060102150405")
	fql.SeqNo = strings2.ReplaceAll(uuid.New().String(), "-", "")
	fql.VersionNo = "2.00"
	fql.buildData(config)
	fql.setSign()

	ret, err := fql.sendReq(FqlEnterUrl)
	if err != nil {
		return nil, err
	}

	ret, err = fql.retData(ret)
	if err != nil {
		return nil, err
	}

	transferReturn := new(FqlTransferReturn)

	err = json.Unmarshal(ret, &transferReturn)

	if err != nil {
		return nil, err
	}
	if transferReturn.SubCode != "0001" && transferReturn.SubCode != "0000" && transferReturn.SubCode != "E0001" {
		return nil, errors.New(transferReturn.SubMsg)
	}
	return transferReturn, nil
}

func (fql *FqlEnter) Withdraw(config *FqlWithdrawConf) (*FqlWithdrawReturn, error) {
	fql.ServiceId = "withdraw"
	fql.Timestamp = time.Now().Format("20060102150405")
	fql.SeqNo = strings2.ReplaceAll(uuid.New().String(), "-", "")
	fql.VersionNo = "2.00"
	fql.buildData(config)
	fql.setSign()

	ret, err := fql.sendReq(FqlEnterUrl)
	if err != nil {
		return nil, err
	}

	ret, err = fql.retData(ret)
	if err != nil {
		return nil, err
	}

	withdrawReturn := new(FqlWithdrawReturn)

	err = json.Unmarshal(ret, &withdrawReturn)

	if err != nil {
		return nil, err
	}
	if withdrawReturn.SubCode != "0001" && withdrawReturn.SubCode != "0000" && withdrawReturn.SubCode != "E0001" {
		return nil, errors.New(withdrawReturn.SubMsg)
	}
	return withdrawReturn, nil
}

func (fql *FqlEnter) buildData(config interface{}) error {
	fql.Data = config
	b, _ := json.Marshal(config)

	fql.DataStr = string(b)

	return nil
}
func (fql *FqlEnter) setSign() {
	b, _ := helper.Sha256WithRsaSign([]byte(fql.SeqNo+fql.DataStr), fql.PriKey)

	fql.Sign = base64.StdEncoding.EncodeToString(b)
}

func (fql *FqlEnter) sendReq(reqUrl string) (b []byte, err error) {

	buffer := bytes.NewBuffer(b)

	err = json.NewEncoder(buffer).Encode(fql)

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
func (fql *FqlEnter) retData(ret []byte) (retData []byte, err error) {
	var baseReturn FqlReturn

	err = json.Unmarshal(ret, &baseReturn)

	if err != nil {
		return
	}

	if baseReturn.Code != "0000" {
		err = errors.New(baseReturn.Msg)

		return
	}

	return []byte(baseReturn.Data), nil
}
