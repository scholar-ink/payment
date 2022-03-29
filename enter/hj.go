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
	HjEnterUrl = "https://www.joinpay.com/filing/"
)

type HjEnter struct {
	*HjConfig
	Method   string      `json:"method"`
	Version  string      `json:"version"`
	Data     interface{} `json:"data"`
	RandStr  string      `json:"rand_str"`
	SignType string      `json:"sign_type"`
	Sign     string      `json:"sign"`
}

type HjReturn struct {
	RespCode string                 `json:"resp_code"`
	RespMsg  string                 `json:"resp_msg"`
	Data     map[string]interface{} `json:"data"`
}

type HjConfig struct {
	MchNo string `json:"mch_no"`
	Key   string `json:"key"`
}

type HjCreateConf struct {
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
type HjCreateReturn struct {
	TradeMerchantNo string `json:"trade_merchant_no"`
	BizCode         string `json:"biz_code"`
	BizMsg          string `json:"biz_msg"`
}

type HjWxApplyConf struct {
	TradeMerchantNo string `json:"trade_merchant_no"` //交易商户号
	SubAppid        string `json:"sub_appid"`         //关联 APPID
	JsapiPath       string `json:"jsapi_path"`        //授权目录
	SubscribeAppid  string `json:"subscribe_appid"`   //推荐关注 APPID
	ReceiptAppid    string `json:"receipt_appid"`     //支付凭证推荐小程序 appid
}
type HjWxApplyReturn struct {
	TradeMchId string `json:"trade_mch_id"`
	BizCode    string `json:"biz_code"`
	BizMsg     string `json:"biz_msg"`
}

type HjModifyReturn struct {
	ResultCode  string `json:"resultCode"`
	ErrorCode   string `json:"errorCode"`
	ErrCodeDesc string `json:"errCodeDesc"`
	Status      string `json:"status"`
	Ext         string `json:"ext"`
}

type HjQueryConf struct {
	TradeMerchantNo string `json:"trade_merchant_no"` //交易商户号
}

type HjQueryReturn struct {
	ProductList []*ProductList `json:"product_list"`
}

type ProductList struct {
	ProductCode   string `json:"product_code"`
	ProductStatus string `json:"product_status"`
	ProductBizMsg string `json:"product_biz_msg"`
}

type HjUploadConf struct {
	ImgFile string `json:"img_file"` //图片文件
}

type HjUploadReturn struct {
	ImageId string `json:"image_id"`
	BizCode string `json:"biz_code"`
	BizMsg  string `json:"biz_msg"`
}

func (Hj *HjEnter) InitConfig(config *HjConfig) {
	Hj.Version = "1.0"
	Hj.HjConfig = config
}

func (Hj *HjEnter) Create(config *HjCreateConf) (*HjCreateReturn, error) {
	Hj.Method = "filing.mch.apply"
	Hj.Version = "2.0"
	Hj.buildData(config)
	Hj.setSign()

	ret, err := Hj.sendReq(HjEnterUrl, Hj)
	if err != nil {
		return nil, err
	}

	ret, err = Hj.retData(ret)
	if err != nil {
		return nil, err
	}

	enterReturn := new(HjCreateReturn)

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
func (Hj *HjEnter) WxApply(config *HjWxApplyConf) (*HjWxApplyReturn, error) {
	Hj.Method = "filing.weixin.apply"
	Hj.Version = "1.0"
	Hj.buildData(config)
	Hj.setSign()

	ret, err := Hj.sendReq(HjEnterUrl, Hj)
	if err != nil {
		return nil, err
	}

	ret, err = Hj.retData(ret)
	if err != nil {
		return nil, err
	}

	signReturn := new(HjWxApplyReturn)

	err = json.Unmarshal(ret, &signReturn)

	if err != nil {
		return nil, err
	}

	if signReturn.BizCode != "B106014" {
		err = errors.New(signReturn.BizMsg)
		return nil, err
	}

	return signReturn, nil
}
func (Hj *HjEnter) Upload(config *HjUploadConf) (*HjUploadReturn, error) {
	Hj.Method = "filing.mch.upload"
	Hj.buildData(config)
	Hj.setSign()

	ret, err := Hj.sendReq(HjEnterUrl, Hj)
	if err != nil {
		return nil, err
	}

	ret, err = Hj.retData(ret)
	if err != nil {
		return nil, err
	}

	uploadReturn := new(HjUploadReturn)

	err = json.Unmarshal(ret, &uploadReturn)

	if err != nil {
		return nil, err
	}

	return uploadReturn, nil
}
func (Hj *HjEnter) Query(config *HjQueryConf) (*HjQueryReturn, error) {
	Hj.Method = "filing.mch.query"
	Hj.Version = "2.0"
	Hj.buildData(config)
	Hj.setSign()

	ret, err := Hj.sendReq(HjEnterUrl, Hj)
	if err != nil {
		return nil, err
	}

	ret, err = Hj.retData(ret)
	if err != nil {
		return nil, err
	}

	queryReturn := new(HjQueryReturn)
	err = json.Unmarshal(ret, queryReturn)
	if err != nil {
		return nil, err
	}

	return queryReturn, nil
}
func (Hj *HjEnter) buildData(config interface{}) error {
	Hj.SignType = "1"
	Hj.RandStr = helper.Md5(strings.CreateIdentifyCode())
	Hj.Data = config
	return nil
}
func (Hj *HjEnter) setSign() {

	b, _ := json.Marshal(Hj.Data)

	Hj.Sign = helper.Md5(fmt.Sprintf("data=%s&mch_no=%s&method=%s&rand_str=%s&sign_type=%s&version=%s&key=%s", string(b), Hj.MchNo, Hj.Method, Hj.RandStr, Hj.SignType, Hj.Version, Hj.Key))
}

func (Hj *HjEnter) sendReq(reqUrl string, params interface{}) (b []byte, err error) {

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
func (Hj *HjEnter) retData(ret []byte) (retData []byte, err error) {
	var baseReturn HjReturn

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
