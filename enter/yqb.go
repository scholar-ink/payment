package enter

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

const (
	YqbEnterUrl = "https://caps.1qianbao.com/caps/request.do"
)

type YqbEnter struct {
	*YqbConfig
	Content         string `json:"-"`
	RequestSerialNo string `json:"reqSerialNo"`
	ChannelBatchNo  string `json:"channelBatchNo"`
	ProductNo       string `json:"productNo"`
	Token           string `json:"-"`
	Version         string `json:"version"`
	HashFunc        string `json:"-"`
	ServiceCode     string `json:"-"`
}

type YqbConfig struct {
	ChannelNo string `json:"channelNo"`
	AesKey    string `json:"-"`
	HashKey   string `json:"-"`
	SystemId  string `json:"-"`
}

//产品信息
type YqbProduct struct {
	ChannelType string `json:"channel_type"` //通道类型
	FeeExp      string `json:"fee_exp"`      //费率
}

//渠道信息
type YqbChannel struct {
	ChannelType string `json:"channel_type"` //通道类型
	ChannelId   string `json:"channel_id"`   //渠道号
	Bussies     string `json:"bussies"`      //经营类目
}

//商户通道信息
type YqbMerchant struct {
	MerRegName  string `json:"mer_reg_name"` //报备商户名称
	SubMchId    string `json:"sub_mch_id"`   //报备商户编号
	ChannelId   string `json:"channel_id"`   //渠道号
	ChannelType string `json:"channel_type"` //通道类型
}

type YqbCreateConf struct {
	OutMerchantNo           string `json:"outMerchantNo"`                    //商户在服务商侧的商户号
	MerchantType            string `json:"merchantType"`                     //商户类型 10 企业 、 11 个体工商户 、12 个人
	BusinessLicenseNo       string `json:"businessLicenseNo,omitempty"`      //营业执照号，企业与个体必填
	MerchantName            string `json:"merchantName"`                     //商户名称
	LegalRepresent          string `json:"legalRepresent"`                   //法人姓名
	BusinessLicensePath     string `json:"businessLicensePath,omitempty"`    //营业执照照片路径，企业与个体必填
	BusinessLicenseEffDate  string `json:"businessLicenseEffDate,omitempty"` //营业执照生效日，企业与个体必填
	BusinessLicenseValDate  string `json:"businessLicenseValDate,omitempty"` //营业执照到期日，企业与个体必填
	FeeRate                 string `json:"feeRate"`                          //费率0~1之间，小数位最多6位
	EstablishDate           string `json:"establishDate,omitempty"`          //成立日期，企业与个体必填
	MerchantShortName       string `json:"merchantShortName"`                //商户简称
	IdentityNo              string `json:"identityNo"`                       //法人身份证号
	IdentityPosPath         string `json:"identityPosPath"`                  //法人身份证正面照路径
	IdentityNegPath         string `json:"identityNegPath"`                  //法人身份证反面照路径
	IdentityEffDate         string `json:"identityEffDate"`                  //法人身份证有效期开始日
	IdentityValDate         string `json:"identityValDate"`                  //法人身份证有效期到期日
	IndustryCode            string `json:"industryCode"`                     //行业编码
	ProvinceCode            string `json:"provinceCode"`                     //省代码
	CityCode                string `json:"cityCode"`                         //市代码
	CountyCode              string `json:"countyCode"`                       //区县代码
	BusinessAddress         string `json:"businessAddress"`                  //商户营业地址
	StoreFacadePath         string `json:"storeFacadePath"`                  //店铺门面照
	StoreBussinessPlacePath string `json:"storeBussinessPlacePath"`          //商户经营场所或仓库照
	StoreWithOwnerPath      string `json:"storeWithOwnerPath,omitempty"`     //经营者手持身份证拍照图片
	ElectronicSignPath      string `json:"electronicSignPath"`               //电子签名照
	BankAccountName         string `json:"bankAccountName"`                  //结算账户名称
	BankAccountNo           string `json:"bankAccountNo"`                    //结算账户账号
	BankName                string `json:"bankName"`                         //开户行名称
	BranchBankName          string `json:"branchBankName,omitempty"`         //支行名称，企业必填
	BankCode                string `json:"bankCode"`                         //银行编码
	SubBankCode             string `json:"subBankCode,omitempty"`            //联行号，企业必填
	BankCardPath            string `json:"bankCardPath"`                     //结算卡照片
	OpenBankPhoneNo         string `json:"openBankPhoneNo,omitempty"`        //银行卡预留手机号，个人或个体时必填
	AdminName               string `json:"adminName"`                        //管理员姓名
	AdminCellPhoneNo        string `json:"adminCellPhoneNo"`                 //管理员手机号
	CallBackUrl             string `json:"callBackUrl"`                      //进件结果回调地址
}

type YqbCreateReturn struct {
	MerchantInfos []map[string]interface{} `json:"merchantInfos"`
}

type YqbQueryConf struct {
	*YqbEnter
	MerchantId    string `json:"merchantId,omitempty"`    //平安付商户号
	OutMerchantNo string `json:"outMerchantNo,omitempty"` //服务商侧商户号
	PageNo        string `json:"pageNo"`                  //页码
	PageSize      string `json:"pageSize"`                //每页记录数
}

type YqbQueryReturn struct {
	MerchantNo      string         `json:"merchant_no"`
	MerchantStatus  string         `json:"mer_status"`
	MerchantList    []*YqbMerchant `json:"-"`
	MerchantListStr string         `json:"mer_list"`
}

type YqbUpdateConf struct {
	*YqbEnter
	MerchantId      string `json:"merchantId,omitempty"`      //平安付商户号
	OutMerchantNo   string `json:"outMerchantNo,omitempty"`   //服务商侧商户号
	BankAccountName string `json:"bankAccountName"`           //结算账户名称
	BankAccountNo   string `json:"bankAccountNo"`             //结算账户账号
	BankName        string `json:"bankName"`                  //开户行名称
	BranchBankName  string `json:"branchBankName,omitempty"`  //支行名称，企业必填
	BankCode        string `json:"bankCode"`                  //银行编码
	SubBankCode     string `json:"subBankCode,omitempty"`     //联行号，企业必填
	BankCardPath    string `json:"bankCardPath"`              //结算卡照片
	OpenBankPhoneNo string `json:"openBankPhoneNo,omitempty"` //银行卡预留手机号，个人或个体时必填
	IdentityNo      string `json:"identityNo"`                //法人身份证号
}

type YqbUpdateReturn struct {
}

type YqbUploadConf struct {
	OutMerchantNo string `json:"outMerchantNo"` //商户在服务商侧的商户号
	User          string
	Password      string
	File          []byte
}

func (yqb *YqbEnter) InitConfig(config *YqbConfig) {
	yqb.YqbConfig = config
}

func (yqb *YqbEnter) Create(config *YqbCreateConf) (*YqbCreateReturn, error) {
	yqb.ServiceCode = "R0867"
	yqb.ProductNo = "P0000119"
	yqb.Version = "1.0"

	yqb.RequestSerialNo = strings.ReplaceAll(uuid.New().String(), "-", "")
	//yqb.ChannelBatchNo = strings.ReplaceAll(uuid.New().String(), "-", "")
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	yqb.ChannelBatchNo = fmt.Sprintf("%06v", rnd.Int31n(1000000))

	merchantInfo := make([]*YqbCreateConf, 0, 1)
	merchantInfo = append(merchantInfo, config)

	params := map[string]interface{}{
		"version":        yqb.Version,
		"reqSerialNo":    yqb.RequestSerialNo,
		"channelBatchNo": yqb.ChannelBatchNo,
		"productNo":      yqb.ProductNo,
		"channelNo":      yqb.ChannelNo,
		"merchantInfos":  merchantInfo,
	}

	ret, err := yqb.sendReq(YqbEnterUrl, params)
	if err != nil {
		return nil, err
	}

	ret, err = yqb.retData(ret)

	if err != nil {
		return nil, err
	}

	enterReturn := new(YqbCreateReturn)
	err = json.Unmarshal(ret, &enterReturn)
	if err != nil {
		return nil, err
	}

	return enterReturn, nil
}
func (yqb *YqbEnter) Query(config *YqbQueryConf) (*YqbQueryReturn, error) {
	yqb.ServiceCode = "R0868"
	yqb.ProductNo = "P0000119"
	yqb.Version = "1.0"
	config.YqbEnter = yqb

	yqb.RequestSerialNo = strings.ReplaceAll(uuid.New().String(), "-", "")
	yqb.ChannelBatchNo = strings.ReplaceAll(uuid.New().String(), "-", "")

	ret, err := yqb.sendReq(YqbEnterUrl, config)
	if err != nil {
		return nil, err
	}

	ret, err = yqb.retData(ret)

	if err != nil {
		return nil, err
	}

	queryReturn := new(YqbQueryReturn)
	err = json.Unmarshal(ret, &queryReturn)
	if err != nil {
		return nil, err
	}

	return queryReturn, nil
}
func (yqb *YqbEnter) Update(config *YqbUpdateConf) (*YqbCreateReturn, error) {
	yqb.ServiceCode = "R0869"
	yqb.ProductNo = "P0000119"
	yqb.Version = "1.0"
	config.YqbEnter = yqb

	yqb.RequestSerialNo = strings.ReplaceAll(uuid.New().String(), "-", "")
	yqb.ChannelBatchNo = strings.ReplaceAll(uuid.New().String(), "-", "")

	ret, err := yqb.sendReq(YqbEnterUrl, config)
	if err != nil {
		return nil, err
	}

	ret, err = yqb.retData(ret)

	if err != nil {
		return nil, err
	}

	enterReturn := new(YqbCreateReturn)
	err = json.Unmarshal(ret, &enterReturn)
	if err != nil {
		return nil, err
	}

	return enterReturn, nil
}
func (yqb *YqbEnter) Upload(config *YqbUploadConf) (string, error) {
	return "", nil
}
func (yqb *YqbEnter) buildData(config interface{}) error {
	return nil
}
func (yqb *YqbEnter) setSign(content []byte) {

	aesKey, _ := base64.StdEncoding.DecodeString(yqb.AesKey)

	cipherText := helper.PKCS7Padding([]byte(content), 16)

	b, _ := helper.AesEncrypt(cipherText, aesKey) //ECB加密

	yqb.Content = base64.StdEncoding.EncodeToString(b)

	fmt.Println(string(content) + yqb.HashKey)

	yqb.Token = helper.Sha1(string(content) + yqb.HashKey)
}

func (yqb *YqbEnter) sendReq(reqUrl string, params interface{}) (b []byte, err error) {

	yqb.HashFunc = "SHA-1"

	buffer := bytes.NewBuffer(b)

	err = json.NewEncoder(buffer).Encode(params)

	if err != nil {
		return nil, err
	}

	yqb.setSign(buffer.Bytes())

	params = map[string]interface{}{
		"content":     yqb.Content,
		"token":       yqb.Token,
		"hashFunc":    yqb.HashFunc,
		"serviceCode": yqb.ServiceCode,
		"systemId":    yqb.SystemId,
	}

	buffer = bytes.NewBuffer(b)

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
func (yqb *YqbEnter) retData(ret []byte) (origin []byte, err error) {

	var baseReturn struct {
		Code     string `json:"code"`
		HashFunc string `json:"hashFunc"`
		Memo     string `json:"memo"`
		Content  string `json:"content"`
	}

	err = json.Unmarshal(ret, &baseReturn)

	if err != nil {
		return
	}

	if baseReturn.Code != "000000" {
		err = errors.New(baseReturn.Memo)
		return
	}

	b, _ := base64.StdEncoding.DecodeString(baseReturn.Content)

	aesKey, _ := base64.StdEncoding.DecodeString(yqb.AesKey)

	origin, err = helper.AesDecrypt(b, aesKey) // ECB解密

	if err != nil {
		return
	}

	origin = helper.PKCS7UnPadding(origin)

	var rspReturn struct {
		RespCode string `json:"respCode"`
		RespMsg  string `json:"respMsg"`
	}

	json.Unmarshal(ret, &rspReturn)

	if rspReturn.RespCode != "000000" {
		err = errors.New(rspReturn.RespMsg)
		return
	}
	return
}
