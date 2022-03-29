package enter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"
	"github.com/scholar-ink/payment/util/map"
	"io"
	"io/ioutil"
	"mime/multipart"
	"reflect"
	"strconv"
)

const (
	QfPayEnterUrl = "https://openapi.qfpay.com"
)

type QfPayEnter struct {
	*QfPayConfig
	Service string `json:"-"`
	Sign    string `json:"sign"`
}

type QfPayConfig struct {
	AppCode string `json:"-"`
	Key     string `json:"key"`
}

type QfPayCreateConf struct {
	UserName              string `json:"username"`                           //用户名（注册手机号）
	PassWord              string `json:"password,omitempty"`                 //密码
	BankType              string `json:"banktype,omitempty"`                 //账户类型，1为对私，2为对公
	UserType              string `json:"usertype,omitempty"`                 //1小微 2个体 3企业
	IdNumber              string `json:"idnumber"`                           //法人身份证号
	IdStatDate            string `json:"idstatdate"`                         //身份证有效开始日期
	IdendDate             string `json:"idenddate"`                          //身份证有效截止日期
	LegalPerson           string `json:"legalperson"`                        //法人姓名
	Name                  string `json:"orgNum"`                             //店主姓名/公司名称
	ShopName              string `json:"shopname"`                           //店铺名称
	LicenseNumber         string `json:"licensenumber,omitempty"`            //营业执照号
	ShopTypeId            string `json:"shoptype_id"`                        //商户类别id
	SalesmanMobile        string `json:"salesman_mobile"`                    //商户绑定的业务员
	BankUser              string `json:"bankuser"`                           //银行开户名
	BankAccount           string `json:"bankaccount"`                        //银行开户号
	HeadBankName          string `json:"headbankname"`                       //所属总行名称
	BankName              string `json:"bankname"`                           //支行名称
	BankCode              string `json:"bankcode"`                           //联行号
	BankProvince          string `json:"bankprovince"`                       //支行所在省份
	BankCity              string `json:"bankcity"`                           //支行所属城市
	BankMobile            string `json:"bankmobile"`                         //开户银行预留手机号
	TenPayRatio           string `json:"tenpay_ratio"`                       //微信支付费率
	AliPayRatio           string `json:"alipay_ratio"`                       //支付宝费率
	UnionPayRatioLe100000 string `json:"unionpay_ratio_le_100000,omitempty"` //云闪付一千元以下的费率
	UnionPayRatioGt100000 string `json:"unionpay_ratio_gt_100000,omitempty"` //云闪付一千元以上的费率
	Province              string `json:"province"`                           //所在省份
	City                  string `json:"city"`                               //商户所在市
	Address               string `json:"address"`                            //详细地址
	IdCardFront           []byte `json:"-"`                                  //身份证正面
	IdCardBack            []byte `json:"-"`                                  //身份证背面
	IdCardInHand          []byte `json:"-"`                                  //实际收款人手持身份证正面店内照
	ShopPhoto             []byte `json:"-"`                                  //经营场所
	GoodsPhoto            []byte `json:"-"`                                  //经营场所内景照片
	LicensePhoto          []byte `json:"-"`                                  //营业执照
	OpenLicense           []byte `json:"-"`                                  //开户许可证
	AuthBankCardFront     []byte `json:"-"`                                  //收款银行卡正面照片
}

type QfPayCreateReturn struct {
	UserName string `json:"username"`
	MchId    string `json:"mchid"`
}

type QfPayQueryConf struct {
	UserName       string `json:"username"`        //用户名（注册手机号）
	SalesManMobile string `json:"salesman_mobile"` //绑定业务员手机号
}

type QfPayQueryReturn struct {
	MchNtId int    `json:"mchnt_id"` //商户编号
	State   int    `json:"state"`    //状态对应码：（10：自动审核成功）（8：审核失败）（7：审核拒绝）（9：等待复审）（5：审核通过）（4：等待审核）
	Memo    string `json:"memo"`     //驳回原因描述
	Desc    string `json:"desc"`     //审核状态描述
}

func (qf *QfPayEnter) InitConfig(config *QfPayConfig) {
	qf.QfPayConfig = config
}

func (qf *QfPayEnter) Create(config *QfPayCreateConf) (*QfPayCreateReturn, error) {
	qf.Service = "/mch/v1/signup"

	qf.buildData(config)

	ret, err := qf.sendUploadReq(QfPayEnterUrl, config)
	if err != nil {
		return nil, err
	}

	ret, err = qf.retData(ret)
	if err != nil {
		return nil, err
	}

	enterReturn := new(QfPayCreateReturn)
	err = json.Unmarshal(ret, &enterReturn)
	if err != nil {
		return nil, err
	}

	return enterReturn, nil
}

func (qf *QfPayEnter) Query(config *QfPayQueryConf) (*QfPayQueryReturn, error) {
	qf.Service = "/mch/v1/apply_info"
	qf.buildData(config)

	ret, err := qf.sendReq(QfPayEnterUrl, config)
	if err != nil {
		return nil, err
	}

	ret, err = qf.retData(ret)
	if err != nil {
		return nil, err
	}

	queryReturn := make([]*QfPayQueryReturn, 0, 0)

	err = json.Unmarshal(ret, &queryReturn)

	if err != nil {
		return nil, err
	}

	if len(queryReturn) == 0 {
		return nil, errors.New("商户账号不存在")
	}

	return queryReturn[0], nil
}

func (qf *QfPayEnter) buildData(config interface{}) error {
	return nil
}
func (qf *QfPayEnter) makeSign(params map[string]interface{}) {

	signStr := helper.CreateLinkString(&params)

	fmt.Println(signStr + qf.Key)

	qf.Sign = helper.Md5(signStr + qf.Key)
}

func (qf *QfPayEnter) sendReq(reqUrl string, params interface{}) (b []byte, err error) {

	mapData := helper.Struct2Map(params)

	qf.makeSign(mapData)

	values := maps.Map2Values(&mapData)

	req := http.NewHttpRequest("GET", reqUrl+qf.Service+"?"+values.Encode(), nil)

	req.Header["X-QF-APPCODE"] = []string{qf.AppCode}
	req.Header["X-QF-SIGN"] = []string{qf.Sign}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	fmt.Println(req.Header)

	rsp, err := http.Client.Do(req)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	return
}
func (qf *QfPayEnter) sendUploadReq(reqUrl string, params *QfPayCreateConf) (b []byte, err error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	b, err = json.Marshal(params)

	mapData := helper.Struct2Map(params)

	qf.makeSign(mapData)

	for k, v := range mapData {
		rt := reflect.TypeOf(v)

		var value string

		switch rt.Kind() {
		case reflect.Int:
			value = strconv.Itoa(v.(int))
		case reflect.Int64:
			value = strconv.FormatInt(v.(int64), 10)
		case reflect.Float64:
			value = strconv.Itoa(int(v.(float64)))
		case reflect.String:
			value = v.(string)
		}
		writer.WriteField(k, value)
	}

	if len(params.IdCardFront) > 0 {
		part, err := writer.CreateFormFile("idcardfront", "idcardfront.png")

		if err != nil {
			return nil, err
		}

		//iocopy
		_, err = io.Copy(part, bytes.NewReader(params.IdCardFront))

		if err != nil {
			return nil, err
		}
	}

	if len(params.IdCardBack) > 0 {
		part, err := writer.CreateFormFile("idcardback", "idcardback.png")

		if err != nil {
			return nil, err
		}

		//iocopy
		_, err = io.Copy(part, bytes.NewReader(params.IdCardBack))

		if err != nil {
			return nil, err
		}
	}

	if len(params.IdCardInHand) > 0 {
		part, err := writer.CreateFormFile("idcardinhand", "idcardinhand.png")

		if err != nil {
			return nil, err
		}

		//iocopy
		_, err = io.Copy(part, bytes.NewReader(params.IdCardInHand))

		if err != nil {
			return nil, err
		}
	}

	if len(params.ShopPhoto) > 0 {
		part, err := writer.CreateFormFile("shopphoto", "shopphoto.png")

		if err != nil {
			return nil, err
		}

		//iocopy
		_, err = io.Copy(part, bytes.NewReader(params.ShopPhoto))

		if err != nil {
			return nil, err
		}
	}

	if len(params.GoodsPhoto) > 0 {
		part, err := writer.CreateFormFile("goodsphoto", "goodsphoto.png")

		if err != nil {
			return nil, err
		}

		//iocopy
		_, err = io.Copy(part, bytes.NewReader(params.GoodsPhoto))

		if err != nil {
			return nil, err
		}
	}

	if len(params.LicensePhoto) > 0 {
		part, err := writer.CreateFormFile("licensephoto", "licensephoto.png")

		if err != nil {
			return nil, err
		}

		//iocopy
		_, err = io.Copy(part, bytes.NewReader(params.LicensePhoto))

		if err != nil {
			return nil, err
		}
	}

	if len(params.OpenLicense) > 0 {
		part, err := writer.CreateFormFile("openlicense", "openlicense.png")

		if err != nil {
			return nil, err
		}

		//iocopy
		_, err = io.Copy(part, bytes.NewReader(params.OpenLicense))

		if err != nil {
			return nil, err
		}
	}

	if len(params.AuthBankCardFront) > 0 {
		part, err := writer.CreateFormFile("authbankcardfront", "authbankcardfront.png")

		if err != nil {
			return nil, err
		}

		//iocopy
		_, err = io.Copy(part, bytes.NewReader(params.AuthBankCardFront))

		if err != nil {
			return nil, err
		}
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	contentType := writer.FormDataContentType()

	req := http.NewHttpRequest("POST", reqUrl+qf.Service, body)

	req.Header["X-QF-APPCODE"] = []string{qf.AppCode}
	req.Header["X-QF-SIGN"] = []string{qf.Sign}
	req.Header.Set("Content-Type", contentType)

	rsp, err := http.Client.Do(req)

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
func (qf *QfPayEnter) retData(ret []byte) (retData []byte, err error) {

	var baseReturn struct {
		RespErr string      `json:"resperr"`
		RespCd  string      `json:"respcd"`
		RespMsg string      `json:"respmsg"`
		Data    interface{} `json:"data"`
	}

	err = json.Unmarshal(ret, &baseReturn)

	if err != nil {
		return
	}

	if baseReturn.RespCd != "0000" {
		fmt.Println(baseReturn.RespMsg)
		err = errors.New(baseReturn.RespErr)
		return
	}

	if err != nil {
		return
	}

	retData, err = json.Marshal(baseReturn.Data)

	return
}

// 图片信息
type QfPayComposeAgreementConfig struct {
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
