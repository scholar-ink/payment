package enter

import (
	"bytes"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"

	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/scholar-ink/payment/util/map"
	"io"
	"io/ioutil"
	"mime/multipart"
	"time"
)

const (
	SumPayEnterUrl = "https://entrance.sumpay.cn/gateway.htm"
)

type SumPayEnter struct {
	*SumPayConfig
	Service      string `json:"service"`
	Version      string `json:"version"`
	Timestamp    string `json:"timestamp"`
	TerminalType string `json:"terminal_type"`
	SignType     string `json:"sign_type"`
	Sign         string `json:"sign"`
}

type SumPayConfig struct {
	AppId        string `json:"app_id"`
	AesKey       string `json:"-"`
	Key          string `json:"-"`
	PfxData      []byte `json:"-"`
	PublicKey    string `json:"-"`
	CertPassWord string `json:"-"`
}

type BaseInfo struct {
	CompanyName            string `json:"company_name"`              //商户名称
	CompanyAbbrName        string `json:"company_abbr_name"`         //商户简称
	ComType                string `json:"com_type"`                  //商户类型
	SocialCreditCode       string `json:"social_credit_code"`        //社会统一信用代码
	LicenseExpiredDate     string `json:"license_expired_date"`      //营业执照截止日期
	CompanyRepresentative  string `json:"company_representative"`    //法人代表姓名
	ComRepIdType           string `json:"com_rep_id_type"`           //法人代表证件类型
	ComRepIdNo             string `json:"com_rep_id_no"`             //法人代表证件号
	ComRepIdExpDate        string `json:"com_rep_id_exp_date"`       //法人代表证件截止日期
	ComProv                string `json:"com_prov"`                  //商户所在省份
	ComCity                string `json:"com_city"`                  //商户所在城市
	Address                string `json:"address"`                   //商户详细地址
	Email                  string `json:"email"`                     //商户邮箱
	PostCode               string `json:"post_code"`                 //邮政编码
	RegisteredCapital      string `json:"registered_capital"`        //注册资本
	RegisteredAddress      string `json:"registered_address"`        //注册地址
	BusinessScope          string `json:"business_scope"`            //经营范围
	ControllingShareholder string `json:"controlling_shareholder"`   //控股股东
	RegisteredTime         string `json:"registered_time"`           //注册时间
	AccessSite             string `json:"access_site"`               //网址或应用链接
	AccessSiteIp           string `json:"access_site_ip"`            //网址或应用链接IP地址
	ParentIndustryCatagory string `json:"parent_industry_catagory"`  //一级行业编码
	SubIndustryCatagory    string `json:"sub_industry_catagory"`     //二级行业编码
	AgentPayIp             string `json:"agent_pay_ip"`              //代付IP地址
	Nature                 string `json:"nature"`                    //企业性质
	AnnualFee              string `json:"annual_fee"`                //系统服务年费
	TradePattern           string `json:"trade_pattern"`             //收单交易类型
	PasswordFreePay        string `json:"password_free_pay"`         //是否免密支付
	ShareBenefitMode       string `json:"share_benefit_mode"`        //分账模式
	AuthCertPic            string `json:"auth_cert_pic"`             //授权书图片 门头照
	BankAccountPermitType  string `json:"bank_account_permit_type"`  //开户许可证类型
	BankAccountPermitPic   string `json:"bank_account_permit_pic"`   //开户许可证照片 银行卡
	LicensePic             string `json:"license_pic"`               //营业执照照片
	ComRepIdFrontPic       string `json:"com_rep_id_front_pic"`      //法人代表身份证正面
	ComRepIdBackPic        string `json:"com_rep_id_back_pic"`       //法人代表身份证反面
	Icp                    string `json:"icp"`                       //ICP照片 营业场所
	SettleCardPic          string `json:"settle_card_pic,omitempty"` //结算卡（卡号面）照片
	BusinessPlacePic       string `json:"business_place_pic"`        //经营场所（含门头照）照片
	BeneficiaryIdentify    string `json:"beneficiary_identify"`      //受益所有人判定
}
type SettleInfo struct {
	AccountType       string `json:"account_type"`           //账户类型 1:对私 2:对公
	RealName          string `json:"realname"`               //收款银行账户名
	BankCode          string `json:"bank_code"`              //收款银行编码
	AccountNo         string `json:"account_no"`             //银行账号
	OpenProv          string `json:"open_prov,omitempty"`    //省份
	OpenCity          string `json:"open_city,omitempty"`    //城市
	BankBranch        string `json:"bank_branch,omitempty"`  //开户支行
	BankLineNo        string `json:"bank_line_no,omitempty"` //联行号
	IdNo              string `json:"id_no,omitempty"`        //开户身份证号
	SettleType        string `json:"settle_type"`            //结算类型
	SettlePeriod      string `json:"settle_period"`          //结算周期
	ReturnWay         string `json:"return_way"`             //返佣模式
	SettleWithdrawWay string `json:"settle_withdraw_way"`    //结算打款方式
}
type ContactInfo struct {
	Contact         string `json:"contact"`           //商户负责人姓名
	ContactIdNo     string `json:"contact_id_no"`     //商户负责人身份证号
	ContactTel      string `json:"contact_tel"`       //商户负责人电话
	ContactEmail    string `json:"contact_email"`     //商户负责人邮箱
	ComRiskName     string `json:"com_risk_name"`     //风控联系人姓名
	ComRiskTel      string `json:"com_risk_tel"`      //风控联系人联系电话
	ComRiskEmail    string `json:"com_risk_email"`    //风控联系人邮箱
	ComOperateName  string `json:"com_operate_name"`  //运营负责人姓名
	ComOperateTel   string `json:"com_operate_tel"`   //运营技术负责人联系电话
	ComOperateEmail string `json:"com_operate_email"` //运营负责人邮箱
}
type ReplenishInfo struct {
	Bond string `json:"bond"` //保证金
	Bd   string `json:"bd"`   //所属BD
}
type ProductInfo struct {
	TradeProCode string `json:"trade_pro_code"` //业务产品码
	PayProCode   string `json:"pay_pro_code"`   //业务产品码
	ChargeWay    string `json:"charge_way"`     //收费方式
	Rate         string `json:"rate"`           //费率
}
type BeneficiaryInfo struct {
	BfyName      string `json:"bfy_name"`        //姓名
	BfyIdType    string `json:"bfy_id_type"`     //证件类型
	BfyIdNo      string `json:"bfy_id_no"`       //证件号
	BfyIdExpDate string `json:"bfy_id_exp_date"` //证件截止日期
	BfyProv      string `json:"bfy_prov"`        //所在省份
	BfyCity      string `json:"bfy_city"`        //所在城市
	BfyAddress   string `json:"bfy_address"`     //详细地址
}

type SumPayCreateConf struct {
	OpType                 string             `json:"op_type"`
	MerNo                  string             `json:"mer_no"`                //商户代码
	ResideMerNo            string             `json:"reside_mer_no"`         //入驻商户代码
	BaseInfoStr            string             `json:"base_info"`             //基本信息
	SettleInfoStr          string             `json:"settle_info"`           //结算信息
	ContactInfoStr         string             `json:"contact_info"`          //联系人信息
	ReplenishInfoStr       string             `json:"replenish_info"`        //补充信息
	ProductInfoListStr     string             `json:"product_info_list"`     //商户签约产品集合
	BeneficiaryInfoListStr string             `json:"beneficiary_info_list"` //受益所有人集合
	BaseInfo               *BaseInfo          `json:"-"`                     //基本信息
	SettleInfo             *SettleInfo        `json:"-"`                     //结算信息
	ContactInfo            *ContactInfo       `json:"-"`                     //联系人信息
	ReplenishInfo          *ReplenishInfo     `json:"-"`                     //补充信息
	ProductInfoList        []*ProductInfo     `json:"-"`                     //商户签约产品集合
	BeneficiaryInfoList    []*BeneficiaryInfo `json:"-"`                     //受益所有人集合
}

type SumPayCreateReturn struct {
	ResideMerchantNo string `json:"reside_mer_no"`
}

type SumPayQueryConf struct {
	MerNo            string `json:"mer_no"`        //商户代码
	ResideMerchantNo string `json:"reside_mer_no"` //入驻商户代码
}
type SumPayQueryReturn struct {
	Response    string `json:"sumpay_merchant_reside_query_response"`
	ResideMerNo string `json:"reside_mer_no"` //入驻商户代码
	Status      string `json:"status"`        //入驻状态
}
type SumValidateConf struct {
	MerNo            string `json:"mer_no"`        //商户代码
	ResideMerchantNo string `json:"reside_mer_no"` //入驻商户代码
	RcvAmount        string `json:"rcv_amount"`    //验资金额，保留两位小数，如：50.01
}
type SumValidateReturn struct {
	ResideMerchantNo string `json:"reside_mer_no"` //入驻商户代码
}
type SumContractConf struct {
	MerNo            string `json:"mer_no"`        //商户代码
	ResideMerchantNo string `json:"reside_mer_no"` //入驻商户代码
}
type SumContractReturn struct {
	ResideMerchantNo string `json:"reside_mer_no"` //入驻商户代码
}
type SumPayUploadConf struct {
	MerNo   string `json:"mer_no"`   //商户编号
	PicType string `json:"pic_type"` //图片类型
	PicUse  string `json:"pic_use"`  //图片类型 1:营业执照照片 2:法人代表身份证 3:ICP许可证 4:开户许可证照片 5:授权书图片 6:其它图片（图片类型必须为zip压缩包）
	PicFile []byte `json:"-"`
}

func (sum *SumPayEnter) InitConfig(config *SumPayConfig) {
	sum.Timestamp = time.Now().Format("20060102150405")
	sum.TerminalType = "api"
	sum.SignType = "CERT"
	sum.SumPayConfig = config
}

func (sum *SumPayEnter) Create(config *SumPayCreateConf) (*SumPayCreateReturn, error) {
	sum.Service = "fosun.sumpay.api.merchant.reside.apply"
	sum.Version = "1.0"

	b, _ := json.Marshal(config.BaseInfo)
	config.BaseInfoStr = base64.StdEncoding.EncodeToString(b)

	b, _ = json.Marshal(config.SettleInfo)
	config.SettleInfoStr = base64.StdEncoding.EncodeToString(b)

	b, _ = json.Marshal(config.ContactInfo)
	config.ContactInfoStr = base64.StdEncoding.EncodeToString(b)

	b, _ = json.Marshal(config.ReplenishInfo)
	config.ReplenishInfoStr = base64.StdEncoding.EncodeToString(b)

	if len(config.ProductInfoList) == 0 {
		config.ProductInfoList = make([]*ProductInfo, 0, 0)
	}
	b, _ = json.Marshal(config.ProductInfoList)
	config.ProductInfoListStr = base64.StdEncoding.EncodeToString(b)

	if len(config.BeneficiaryInfoList) == 0 {
		config.BeneficiaryInfoList = make([]*BeneficiaryInfo, 0, 0)
	}
	b, _ = json.Marshal(config.BeneficiaryInfoList)
	config.BeneficiaryInfoListStr = base64.StdEncoding.EncodeToString(b)

	ret, err := sum.sendReq(SumPayEnterUrl, config)
	if err != nil {
		return nil, err
	}

	err = sum.retData(ret)

	if err != nil {
		return nil, err
	}

	enterReturn := new(SumPayCreateReturn)
	err = json.Unmarshal(ret, &enterReturn)
	if err != nil {
		return nil, err
	}

	return enterReturn, nil
}
func (sum *SumPayEnter) Query(config *SumPayQueryConf) (*SumPayQueryReturn, error) {
	sum.Service = "fosun.sumpay.api.merchant.reside.query"
	sum.Version = "1.0"
	ret, err := sum.sendReq(SumPayEnterUrl, config)
	if err != nil {
		return nil, err
	}

	err = sum.retData(ret)

	if err != nil {
		return nil, err
	}

	queryReturn := new(SumPayQueryReturn)
	err = json.Unmarshal(ret, &queryReturn)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(queryReturn.Response), &queryReturn)
	if err != nil {
		return nil, err
	}

	return queryReturn, nil
}
func (sum *SumPayEnter) Upload(config *SumPayUploadConf) (string, error) {

	sum.Service = "fosun.sumpay.api.merchant.reside.pic.upload"
	sum.Version = "1.0"

	ret, err := sum.sendUploadReq(SumPayEnterUrl, config)

	if err != nil {
		return "", err
	}

	err = sum.retData(ret)

	if err != nil {
		return "", err
	}

	var sumPayUploadReturn struct {
		PicName string `json:"pic_name"`
	}

	json.Unmarshal(ret, &sumPayUploadReturn)

	return sumPayUploadReturn.PicName, nil
}
func (sum *SumPayEnter) Validate(config *SumValidateConf) (*SumValidateReturn, error) {
	sum.Service = "fosun.sumpay.api.merchant.reside.remitt.validate"
	sum.Version = "1.0"
	ret, err := sum.sendReq(SumPayEnterUrl, config)
	if err != nil {
		return nil, err
	}

	err = sum.retData(ret)

	if err != nil {
		return nil, err
	}

	validateReturn := new(SumValidateReturn)
	err = json.Unmarshal(ret, &validateReturn)
	if err != nil {
		return nil, err
	}

	return validateReturn, nil
}
func (sum *SumPayEnter) Contract(config *SumContractConf) (*SumContractReturn, error) {
	sum.Service = "fosun.sumpay.api.merchant.reside.sign.contract"
	sum.Version = "1.0"
	ret, err := sum.sendReq(SumPayEnterUrl, config)
	if err != nil {
		return nil, err
	}

	err = sum.retData(ret)

	if err != nil {
		return nil, err
	}

	contractReturn := new(SumContractReturn)
	err = json.Unmarshal(ret, &contractReturn)
	if err != nil {
		return nil, err
	}

	return contractReturn, nil
}
func (sum *SumPayEnter) buildData(config interface{}) error {
	return nil
}
func (sum *SumPayEnter) makeSign(params map[string]interface{}) {

	aesKey, _ := helper.Rsa2Encrypt([]byte(base64.StdEncoding.EncodeToString([]byte(sum.AesKey))), sum.PublicKey)

	params["aes_key"] = aesKey

	sign := helper.CreateLinkString(&params)

	b, err := helper.Sha256WithRsaSignWithPassWord(sum.PfxData, []byte(sign), sum.CertPassWord)

	fmt.Println(err)

	sum.Sign = base64.StdEncoding.EncodeToString(b)
}

func (sum *SumPayEnter) sendReq(reqUrl string, params interface{}) (b []byte, err error) {

	b, err = json.Marshal(params)

	mapData := helper.Struct2Map(params)

	mapData["service"] = sum.Service
	mapData["version"] = sum.Version
	mapData["app_id"] = sum.AppId
	mapData["timestamp"] = sum.Timestamp
	mapData["terminal_type"] = sum.TerminalType
	mapData["sign_type"] = sum.SignType

	sum.makeSign(mapData)

	mapData["sign"] = sum.Sign

	values := maps.Map2Values(&mapData)

	rsp, err := http.Client.PostForm(reqUrl, values)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	return
}
func (sum *SumPayEnter) sendUploadReq(reqUrl string, params *SumPayUploadConf) (b []byte, err error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	mapData := helper.Struct2Map(params)

	mapData["service"] = sum.Service
	mapData["version"] = sum.Version
	mapData["app_id"] = sum.AppId
	mapData["timestamp"] = sum.Timestamp
	mapData["terminal_type"] = sum.TerminalType
	mapData["sign_type"] = sum.SignType

	sum.makeSign(mapData)

	mapData["sign"] = sum.Sign

	b, _ = json.Marshal(mapData)

	for key, value := range mapData {
		writer.WriteField(key, value.(string))
	}

	//关键的一步操作
	part, err := writer.CreateFormFile("pic_file", "pic_file.png")

	if err != nil {
		return nil, err
	}

	//iocopy
	_, err = io.Copy(part, bytes.NewReader(params.PicFile))

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

func (sum *SumPayEnter) retData(ret []byte) (err error) {

	var baseReturn struct {
		RespCode string `json:"resp_code"`
		RespMsg  string `json:"resp_msg"`
	}

	err = json.Unmarshal(ret, &baseReturn)

	if err != nil {
		return
	}

	if baseReturn.RespCode != "000000" {
		err = errors.New(baseReturn.RespMsg)
		return
	}

	return
}

func (sum *SumPayEnter) EncryptFiled(origin string) string {

	cipherText := helper.PKCS7Padding([]byte(origin), 16)

	encrypted, _ := helper.AesEncrypt(cipherText, []byte(sum.AesKey))

	return base64.StdEncoding.EncodeToString(encrypted)
}
