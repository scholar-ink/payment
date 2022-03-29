package enter

import (
	"bytes"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"

	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"strings"
)

const (
	CmEnterUrl = "http://internal.congmingpay.com/internal/"
)

type CmEnter struct {
	*CmConfig
	Service string `json:"-"`
	Sign    string `json:"sign"`
}

type CmConfig struct {
	AgentId string `json:"agent_id"`
	Key     string `json:"-"`
}

type CmCreateConf struct {
	OrderId            string `json:"order_id"`                      //申请商户进件单号
	ShopName           string `json:"shop_name"`                     //商户全称
	ShopNickName       string `json:"shop_nickname"`                 //商户简称
	ShopKeeper         string `json:"shop_keeper"`                   //负责人姓名
	KeeperIdentity     string `json:"keeper_identity"`               //负责人身份证号
	KeeperPhone        string `json:"keeper_phone"`                  //负责人电话
	ShopPhone          string `json:"shop_phone"`                    //客服电话
	Email              string `json:"email"`                         //邮箱
	Province           string `json:"province"`                      //店铺所在省
	City               string `json:"city"`                          //店铺所在市
	Area               string `json:"area"`                          //店铺所在区
	ShopAddress        string `json:"shop_address"`                  //店铺详细地址
	PayType            string `json:"pay_type"`                      //微信渠道
	PayTypeAli         string `json:"pay_type_alipay"`               //支付宝渠道
	BusinessType       string `json:"business_type"`                 //商户类型1.企业2.个体3.小微
	LicenceNo          string `json:"licence_no,omitempty"`          //营业执照号
	LicenceBeginDate   string `json:"licence_begin_date,omitempty"`  //营业执照起始时间
	LicenceExpireDate  string `json:"licence_expire_date,omitempty"` //营业执照截止时间
	ArTifName          string `json:"artif_name"`                    //法人姓名
	ArTifPhone         string `json:"artif_phone,omitempty"`         //法人手机号
	ArTifIdentity      string `json:"artif_identity,omitempty"`      //法人身份证号
	RateWx             string `json:"rate_wx"`                       //微信费率3.8
	RateAliPay         string `json:"rate_alipay"`                   //支付宝费率3.8
	Identity           string `json:"identity"`                      //结算人身份证号
	IdentityStartTime  string `json:"identity_start_time"`           //结算人/法人身份证有 效期起始时间
	IdentityExpireTime string `json:"identity_expire_time"`          //结算人/法人身份证有 效期截止时间
	CardName           string `json:"card_name"`                     //结算人姓名
	Card               string `json:"card"`                          //结算卡号
	CardPhone          string `json:"card_phone"`                    //结算人手机号
	BankName           string `json:"bank_name"`                     //结算银行
	BankAddress        string `json:"bank_address"`                  //结算开户支行
	BankAddNo          string `json:"bank_add_no"`                   //结算卡开户行联行号
	AreaType           string `json:"type"`                          //商户一级经营范围
	Classify           string `json:"classify"`                      //商户二级经营范围
	Description        string `json:"description,omitempty"`         //备注信息，如配置 appid 等
	NotifyUrl          string `json:"notify_url"`                    //商户审核通过回调地址
	AgentId            string `json:"agent_id"`                      //合作商商户号
	SignInType         string `json:"signin_type"`                   //快速进件状态 1.快速进件2.人工
}

type CmCreateReturn struct {
	ShopId  string `json:"shop_id"`
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
}
type CmUploadConf struct {
	ShopId  string            `json:"shop_id"` //商户号
	PicData map[string][]byte `json:"pic_data"`
}

func (cm *CmEnter) InitConfig(config *CmConfig) {
	cm.CmConfig = config
}

func (cm *CmEnter) Create(config *CmCreateConf) (*CmCreateReturn, error) {
	cm.Service = "registernewmerchant"

	ret, err := cm.sendReq(CmEnterUrl, config)
	if err != nil {
		return nil, err
	}

	err = cm.retData(ret)

	if err != nil {
		return nil, err
	}

	enterReturn := new(CmCreateReturn)
	err = json.Unmarshal(ret, &enterReturn)
	if err != nil {
		return nil, err
	}

	return enterReturn, nil
}
func (cm *CmEnter) Upload(config *CmUploadConf) error {

	cm.Service = "uploadmerchantimg"

	ret, err := cm.sendUploadReq(CmEnterUrl, config)

	if err != nil {
		return err
	}

	err = cm.retData(ret)

	if err != nil {
		return err
	}

	var cmPayUploadReturn struct {
		List string `json:"list"`
	}

	err = json.Unmarshal(ret, &cmPayUploadReturn)

	if err != nil {
		return err
	}

	if cmPayUploadReturn.List == "" {
		return errors.New("没有成功上传的图片")
	}

	type cmPayUploadList struct {
		FileType   string `json:"file_type"`
		ResultCode string `json:"result_code"`
		PicUrl     string `json:"pic_url"`
	}

	picList := make([]*cmPayUploadList, 0, 0)

	err = json.Unmarshal([]byte(cmPayUploadReturn.List), &picList)

	if err != nil {
		return err
	}

	if len(picList) == 0 {
		return errors.New("没有成功上传的图片")
	}

	return nil
}
func (cm *CmEnter) buildData(config interface{}) error {
	return nil
}
func (cm *CmEnter) makeSign(params map[string]interface{}) {

	signStr := helper.CreateLinkString(&params)

	fmt.Println(signStr + "&key=" + cm.Key)

	cm.Sign = strings.ToUpper(helper.Md5(signStr + "&key=" + cm.Key))
}

func (cm *CmEnter) sendReq(reqUrl string, params interface{}) (b []byte, err error) {

	mapData := helper.Struct2Map(params)
	mapData["agent_id"] = cm.AgentId

	cm.makeSign(mapData)

	mapData["sign"] = cm.Sign

	buffer := bytes.NewBuffer(b)

	err = json.NewEncoder(buffer).Encode(mapData)

	if err != nil {
		return nil, err
	}

	rsp, err := http.Client.Post(reqUrl+cm.Service+".do", "application/json;charset=UTF-8", buffer)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	return
}
func (cm *CmEnter) sendUploadReq(reqUrl string, params *CmUploadConf) (b []byte, err error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("shop_id", params.ShopId)

	for key, value := range params.PicData {
		//关键的一步操作
		part, err := writer.CreateFormFile(key, "pic_file.png")

		if err != nil {
			return nil, err
		}

		_, err = part.Write(value)
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	contentType := writer.FormDataContentType()

	rsp, err := http.Client.Post(reqUrl+cm.Service+".do", contentType, body)

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

func (cm *CmEnter) retData(ret []byte) (err error) {

	var baseReturn struct {
		ResultCode string `json:"result_code"`
		ErrorMsg   string `json:"error_msg"`
	}

	err = json.Unmarshal(ret, &baseReturn)

	if err != nil {
		return
	}

	if baseReturn.ResultCode != "success" {
		err = errors.New(baseReturn.ErrorMsg)
		return
	}

	return
}
