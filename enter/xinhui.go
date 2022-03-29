package enter

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"
	"github.com/scholar-ink/payment/util/map"
	"io/ioutil"
	"strings"
)

const (
	XinHuiEnterUrl = "http://at.xhepay.com/uni/gateway"
)

type XinHuiEnter struct {
	*XinHuiConfig
	ServiceType string `json:"service_type"`
	RequestId   string `json:"request_id"`
	Sign        string `json:"sign"`
	Version     string `json:"version"`
}

type XinHuiConfig struct {
	AgentMerNo string `json:"agent_mer_no"`
	Key        string `json:"-"`
}

//产品信息
type XinHuiProduct struct {
	ChannelType string `json:"channel_type"` //通道类型
	FeeExp      string `json:"fee_exp"`      //费率
}

//渠道信息
type XinHuiChannel struct {
	ChannelType string `json:"channel_type"` //通道类型
	ChannelId   string `json:"channel_id"`   //渠道号
	Bussies     string `json:"bussies"`      //经营类目
}

//商户通道信息
type XinHuiMerchant struct {
	MerRegName  string `json:"mer_reg_name"` //报备商户名称
	SubMchId    string `json:"sub_mch_id"`   //报备商户编号
	ChannelId   string `json:"channel_id"`   //渠道号
	ChannelType string `json:"channel_type"` //通道类型
}

type XinHuiCreateConf struct {
	MerchantNo       string           `json:"merchant_no,omitempty"`       //商户编号
	MerType          string           `json:"mer_type"`                    //商户类型
	MerchantSubName  string           `json:"merchant_sub_name"`           //商户简称
	MerchantName     string           `json:"merchant_name"`               //商户名称
	CertType         string           `json:"cert_type"`                   //证件类型
	CertNo           string           `json:"cert_no"`                     //证件号码
	CertIndate       string           `json:"cert_indate"`                 //证件有效期
	LegalName        string           `json:"legal_name"`                  //法人姓名
	LegalIdCard      string           `json:"legal_idcard"`                //法人身份证号
	Contact          string           `json:"contact"`                     //联系人
	ContactPhone     string           `json:"contact_phone"`               //联系电话
	ContactEmail     string           `json:"contact_email"`               //联系邮箱
	ProvinceNo       string           `json:"province_no"`                 //省份编码
	CityNo           string           `json:"city_no"`                     //城市编码
	DistrictNo       string           `json:"district_no"`                 //区县编码
	Address          string           `json:"address"`                     //地址
	ServicePhone     string           `json:"service_phone"`               //客服电话
	SettleType       string           `json:"settle_type"`                 //结算类型
	SettleCardNo     string           `json:"settle_card_no"`              //结算卡号
	BankMerName      string           `json:"bank_mer_name"`               //开户行商户名称
	BankName         string           `json:"bank_name"`                   //开户行银行名称
	BankProvince     string           `json:"bank_province"`               //开户行省份
	BankCity         string           `json:"bank_city"`                   //开户行城市
	BankBranchName   string           `json:"bank_branch_name"`            //开户行支行名称
	BankCode         string           `json:"bank_code,omitempty"`         //开户联行号
	LicensePic       string           `json:"license_pic"`                 //营业执照
	LegalIdCardFront string           `json:"legal_idcard_front"`          //法人身份证正面
	LegalIdCardBack  string           `json:"legal_idcard_back"`           //法人身份证反面
	ShopEntrancePic  string           `json:"shop_entrance_pic,omitempty"` //门头照
	ProductInfoStr   string           `json:"product_info"`                //开通产品信息
	ProductInfo      []*XinHuiProduct `json:"-"`                           //开通产品信息
	ChannelInfoStr   string           `json:"channel_info"`                //开通产品信息
	ChannelInfo      []*XinHuiChannel `json:"-"`                           //开通产品信息
}

type XinHuiCreateReturn struct {
	MerchantNo string `json:"merchant_no"`
}

type XinHuiUpdateConf struct {
	MerchantNo     string           `json:"merchant_no,omitempty"` //商户编号
	SettleCardNo   string           `json:"settle_card_no"`        //结算卡号
	BankName       string           `json:"bank_name"`             //开户行银行名称
	BankProvince   string           `json:"bank_province"`         //开户行省份
	BankCity       string           `json:"bank_city"`             //开户行城市
	BankBranchName string           `json:"bank_branch_name"`      //开户行支行名称
	BankCode       string           `json:"bank_code,omitempty"`   //开户联行号
	ProductInfoStr string           `json:"product_info"`          //开通产品信息
	ProductInfo    []*XinHuiProduct `json:"-"`                     //开通产品信息
}

type XinHuiUpdateReturn struct {
	MerchantNo string `json:"merchant_no"`
}

type XinHuiQueryConf struct {
	MerchantNo string `json:"merchant_no"` //商户编号
}

type XinHuiQueryReturn struct {
	MerchantNo      string            `json:"merchant_no"`
	MerchantStatus  string            `json:"mer_status"`
	MerchantList    []*XinHuiMerchant `json:"-"`
	MerchantListStr string            `json:"mer_list"`
}
type XinHuiWxApplyConf struct {
	MerchantNo     string `json:"merchant_no"`      //商户编号
	SubMchId       string `json:"sub_mch_id"`       //报备商户编号
	SubAppid       string `json:"sub_appid"`        //绑定APPID
	SubScribeAppId string `json:"sub_scribe_appid"` //推荐关注APPID
	ReceiptAppId   string `json:"receipt_appid"`    //支付凭证推荐小程序appid
	AuthPaths      string `json:"auth_paths"`       //授权目录
}

type XinHuiWxApplyReturn struct {
	MerchantNo     string            `json:"merchant_no"`
	MerchantStatus string            `json:"mer_status"`
	MerchantList   []*XinHuiMerchant `json:"mer_list"`
}
type XinHuiUploadConf struct {
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
	File            []byte
}

func (xh *XinHuiEnter) InitConfig(config *XinHuiConfig) {
	xh.XinHuiConfig = config
}

func (xh *XinHuiEnter) Create(config *XinHuiCreateConf) (*XinHuiCreateReturn, error) {
	xh.ServiceType = "xh.uni.merchant.reg"
	xh.Version = "1.0"
	xh.RequestId = strings.ReplaceAll(uuid.New().String(), "-", "")

	if len(config.ProductInfo) == 0 {
		config.ProductInfo = make([]*XinHuiProduct, 0, 0)
	}
	b, _ := json.Marshal(config.ProductInfo)
	config.ProductInfoStr = string(b)

	if len(config.ChannelInfo) == 0 {
		config.ChannelInfo = make([]*XinHuiChannel, 0, 0)
	}
	b, _ = json.Marshal(config.ChannelInfo)
	config.ChannelInfoStr = string(b)

	ret, err := xh.sendReq(XinHuiEnterUrl, config)
	if err != nil {
		return nil, err
	}

	err = xh.retData(ret)

	if err != nil {
		return nil, err
	}

	enterReturn := new(XinHuiCreateReturn)
	err = json.Unmarshal(ret, &enterReturn)
	if err != nil {
		return nil, err
	}

	return enterReturn, nil
}
func (xh *XinHuiEnter) Update(config *XinHuiUpdateConf) (*XinHuiUpdateReturn, error) {
	xh.ServiceType = "xh.uni.merchant.update"
	xh.Version = "1.0"
	xh.RequestId = strings.ReplaceAll(uuid.New().String(), "-", "")

	if len(config.ProductInfo) == 0 {
		config.ProductInfo = make([]*XinHuiProduct, 0, 0)
	}
	b, _ := json.Marshal(config.ProductInfo)
	config.ProductInfoStr = string(b)

	ret, err := xh.sendReq(XinHuiEnterUrl, config)
	if err != nil {
		return nil, err
	}

	err = xh.retData(ret)

	if err != nil {
		return nil, err
	}

	updateReturn := new(XinHuiUpdateReturn)
	err = json.Unmarshal(ret, &updateReturn)
	if err != nil {
		return nil, err
	}

	return updateReturn, nil
}
func (xh *XinHuiEnter) Query(config *XinHuiQueryConf) (*XinHuiQueryReturn, error) {
	xh.ServiceType = "xh.uni.merchant.qry"
	xh.Version = "1.0"
	xh.RequestId = strings.ReplaceAll(uuid.New().String(), "-", "")

	ret, err := xh.sendReq(XinHuiEnterUrl, config)
	if err != nil {
		return nil, err
	}

	err = xh.retData(ret)

	if err != nil {
		return nil, err
	}

	queryReturn := new(XinHuiQueryReturn)
	err = json.Unmarshal(ret, &queryReturn)

	if err != nil {
		return nil, err
	}

	merchantList := make([]*XinHuiMerchant, 0, 0)
	json.Unmarshal([]byte(queryReturn.MerchantListStr), &merchantList)
	queryReturn.MerchantList = merchantList

	return queryReturn, nil
}
func (xh *XinHuiEnter) WxApply(config *XinHuiWxApplyConf) (*XinHuiWxApplyReturn, error) {
	xh.ServiceType = "xh.uni.merchant.wx.config"
	xh.Version = "1.0"
	xh.RequestId = strings.ReplaceAll(uuid.New().String(), "-", "")

	ret, err := xh.sendReq(XinHuiEnterUrl, config)
	if err != nil {
		return nil, err
	}

	err = xh.retData(ret)

	if err != nil {
		return nil, err
	}

	WxApplyReturn := new(XinHuiWxApplyReturn)
	err = json.Unmarshal(ret, &WxApplyReturn)
	if err != nil {
		return nil, err
	}

	return WxApplyReturn, nil
}
func (xh *XinHuiEnter) Upload(config *XinHuiUploadConf) (string, error) {
	return "", nil
}
func (xh *XinHuiEnter) buildData(config interface{}) error {
	return nil
}
func (xh *XinHuiEnter) setSign(params map[string]interface{}) {

	signStr := helper.CreateLinkString(&params)

	b, err := helper.Sha256WithRsaSign([]byte(signStr), xh.Key)

	fmt.Println(err)

	xh.Sign = base64.StdEncoding.EncodeToString(b)

	fmt.Println(xh.Sign)
}

func (xh *XinHuiEnter) sendReq(reqUrl string, params interface{}) (b []byte, err error) {

	mapData := helper.Struct2Map(params)

	mapData["service_type"] = xh.ServiceType
	mapData["request_id"] = xh.RequestId
	mapData["version"] = xh.Version
	mapData["agent_mer_no"] = xh.AgentMerNo

	xh.setSign(mapData)

	if cardNo, ok := mapData["settle_card_no"]; ok {

		b, _ := helper.Rsa2Encrypt2([]byte(cardNo.(string)), xh.Key)

		mapData["settle_card_no"] = base64.StdEncoding.EncodeToString(b)
	}

	mapData["sign"] = xh.Sign

	values := maps.Map2Values(&mapData)

	rsp, err := http.Client.PostForm(reqUrl, values)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	return
}
func (xh *XinHuiEnter) retData(ret []byte) (err error) {

	var baseReturn struct {
		Version string `json:"version"`
		RspCode string `json:"rsp_code"`
		RspMsg  string `json:"rsp_msg"`
		Sign    string `json:"sign"`
	}

	err = json.Unmarshal(ret, &baseReturn)

	if err != nil {
		return
	}

	if baseReturn.RspCode != "0000" {
		err = errors.New(baseReturn.RspMsg)
		return
	}
	return
}
