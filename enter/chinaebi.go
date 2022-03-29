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
	"reflect"
	"strconv"
)

const (
	ChinaEbiEnterUrl = "http://116.228.47.74:18480/merchant_agent_foreign/"
)

type ChinaEbiEnter struct {
	*ChinaEbiConfig
	Service string `json:"-"`
	Sign    string `json:"sign"`
}

type ChinaEbiConfig struct {
	OrgNumber   string `json:"orgNumber"`
	AgentNumber string `json:"agentNumber"`
	PrivateKey  string `json:"-"`
}

type ChinaEbiCreateConf struct {
	SeqNo          string `json:"seqNo"`                    //请求流水号
	MercMbl        string `json:"mercMbl"`                  //商户手机号
	MercCnm        string `json:"mercCnm"`                  //商户名称
	MercAbbr       string `json:"mercAbbr"`                 //商户简称
	MercHotLin     string `json:"mercHotLin"`               //商户电话
	MccCd          string `json:"mccCd"`                    //MCC 码
	MercProv       string `json:"mercProv"`                 //四位银联编号
	MercCity       string `json:"mercCity"`                 //归属市
	MercCounty     string `json:"mercCounty"`               //归属县区
	BusAddr        string `json:"busAddr"`                  //营业地址
	MercAttr       string `json:"mercAttr"`                 //商户性质
	RegId          string `json:"regId"`                    //营业执照号
	RegExpDtD      string `json:"regExpDtD"`                //营业执照过期时间
	CrpIdTyp       string `json:"crpIdTyp"`                 //法人证件类型
	CrpIdNo        string `json:"crpIdNo"`                  //法人证件号码
	CrpNm          string `json:"crpNm"`                    //法人姓名
	CrpExpDtD      string `json:"crpExpDtD"`                //法人证件过期日期
	StlSign        string `json:"stlSign"`                  //结算账号公私标志
	StlWcLbnkNo    string `json:"stlWcLbnkNo"`              //联行行号
	StlOac         string `json:"stlOac"`                   //银行账号
	BnkAcnm        string `json:"bnkAcnm"`                  //银行开户名称
	UsrOprEmail    string `json:"usrOprEmail"`              //商户管理员 EMAIL
	DebitFee       string `json:"debitFee"`                 //借记费率
	DebitFeeLimit  string `json:"debitFeeLimit"`            //借记封顶额
	CreditFee      string `json:"creditFee"`                //贷记费率
	D0Fee          string `json:"d0Fee,omitempty"`          //D0 额外手续费费率
	D0FeeQuota     string `json:"d0FeeQuota,omitempty"`     //D0 额外定额手续费
	UnionCreditFee string `json:"unionCreditFee,omitempty"` //云闪付贷记费率
	UnionDebitFee  string `json:"unionDebitFee,omitempty"`  //云闪付借记费率
	AliFee         string `json:"aliFee"`                   //支付宝费率
	WxFee          string `json:"wxFee"`                    //微信费率
	AliFlg         string `json:"aliFlg"`                   //是否开通支付宝0 开通，1 不开通
	WxFlg          string `json:"wxFlg"`                    //是否开通微信0 开通，1 不开通
	UnionFlg       string `json:"unionFlg"`                 //是否开通微信0 开通，1 不开通
	OutMercId      string `json:"outMercId"`                //外部商户号
	SettType       string `json:"settType"`                 //结算类型 D0、T1(默认)
	ZZ1            []byte `json:"-"`                        //三证合一或营业执照
	SFZ1           []byte `json:"-"`                        //法人身份证正面
	SFZ2           []byte `json:"-"`                        //法人身份证反面
	CDJJ1          []byte `json:"-"`                        //场地街景
	CDMT1          []byte `json:"-"`                        //场地门头
	CDNJ1          []byte `json:"-"`                        //场地内景 1
	YHK            []byte `json:"-"`                        //场地内景 1
}

type ChinaEbiCreateReturn struct {
	MercId  string `json:"mercId"`
	MercSts string `json:"mercSts"`
}

type ChinaEbiQueryConf struct {
	OrgNumber string `json:"orgNumber"`
	DyMchNo   string `json:"dyMchNo"` //电银商户号
}

type ChinaEbiQueryReturn struct {
	DyMchNo    string `json:"dyMchNo"`    //电银商户号
	MercCnm    string `json:"mercCnm"`    //商户名称
	MercAbbr   string `json:"mercAbbr"`   //商户简称
	OutMercId  string `json:"outMercId"`  //外部商户号
	MchStatus  string `json:"mchStatus"`  //商户状态
	MchType    string `json:"mchType"`    //商户类型
	SettType   string `json:"settType"`   //结算类型
	SettStatus string `json:"settStatus"` //结算状态
	AudSts     string `json:"audSts"`     //审核状态
	AudMsg     string `json:"audMsg"`     //审核描述
}

func (ebi *ChinaEbiEnter) InitConfig(config *ChinaEbiConfig) {
	ebi.ChinaEbiConfig = config
}

func (ebi *ChinaEbiEnter) Create(config *ChinaEbiCreateConf) (*ChinaEbiCreateReturn, error) {
	ebi.Service = "/rest/standardMerchant/inComing"

	ret, err := ebi.sendUploadReq(ChinaEbiEnterUrl, config)
	if err != nil {
		return nil, err
	}

	ret, err = ebi.retData(ret)
	if err != nil {
		return nil, err
	}

	enterReturn := new(ChinaEbiCreateReturn)
	err = json.Unmarshal(ret, &enterReturn)
	if err != nil {
		return nil, err
	}

	return enterReturn, nil
}

func (ebi *ChinaEbiEnter) Query(config *ChinaEbiQueryConf) (*ChinaEbiQueryReturn, error) {
	ebi.Service = "/rest/merchantInfo/query"
	config.OrgNumber = ebi.OrgNumber
	ret, err := ebi.sendReq(ChinaEbiEnterUrl, config)
	if err != nil {
		return nil, err
	}

	ret, err = ebi.retData(ret)
	if err != nil {
		return nil, err
	}

	queryReturn := new(ChinaEbiQueryReturn)

	err = json.Unmarshal(ret, &queryReturn)

	if err != nil {
		return nil, err
	}

	return queryReturn, nil
}

func (ebi *ChinaEbiEnter) makeSign(params map[string]interface{}) {

	b, _ := json.Marshal(params)

	fmt.Println(string(b))

	b, err := helper.Md5WithRsaSignWithPKCS8(b, ebi.PrivateKey)

	if err != nil {
		fmt.Println(err)
	}

	ebi.Sign = base64.StdEncoding.EncodeToString(b)

	fmt.Println(ebi.Sign)
}

func (ebi *ChinaEbiEnter) sendReq(reqUrl string, params interface{}) (b []byte, err error) {

	mapData := helper.Struct2Map(params)

	ebi.makeSign(mapData)

	mapData["sign"] = ebi.Sign

	buffer := bytes.NewBuffer(b)

	err = json.NewEncoder(buffer).Encode(mapData)

	if err != nil {
		return nil, err
	}

	rsp, err := http.Client.Post(reqUrl+ebi.Service, "application/json;charset=UTF-8", buffer)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	return
}
func (ebi *ChinaEbiEnter) sendUploadReq(reqUrl string, params *ChinaEbiCreateConf) (b []byte, err error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	mapData := helper.Struct2Map(params)
	mapData["orgNumber"] = ebi.OrgNumber
	mapData["agentNumber"] = ebi.AgentNumber

	ebi.makeSign(mapData)

	mapData["sign"] = ebi.Sign

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

	if len(params.ZZ1) > 0 {
		part, err := writer.CreateFormFile("ZZ1", "ZZ1.png")

		if err != nil {
			return nil, err
		}

		//iocopy
		_, err = io.Copy(part, bytes.NewReader(params.ZZ1))

		if err != nil {
			return nil, err
		}
	}

	if len(params.SFZ1) > 0 {
		part, err := writer.CreateFormFile("SFZ1", "SFZ1.png")

		if err != nil {
			return nil, err
		}

		//iocopy
		_, err = io.Copy(part, bytes.NewReader(params.SFZ1))

		if err != nil {
			return nil, err
		}
	}

	if len(params.SFZ2) > 0 {
		part, err := writer.CreateFormFile("SFZ2", "SFZ2.png")

		if err != nil {
			return nil, err
		}

		//iocopy
		_, err = io.Copy(part, bytes.NewReader(params.SFZ2))

		if err != nil {
			return nil, err
		}
	}

	if len(params.CDJJ1) > 0 {
		part, err := writer.CreateFormFile("CDJJ1", "CDJJ1.jpg")

		if err != nil {
			return nil, err
		}

		//iocopy
		_, err = io.Copy(part, bytes.NewReader(params.CDJJ1))

		if err != nil {
			return nil, err
		}
	}

	if len(params.CDMT1) > 0 {
		part, err := writer.CreateFormFile("CDMT1", "CDMT1.jpg")

		if err != nil {
			return nil, err
		}

		//iocopy
		_, err = io.Copy(part, bytes.NewReader(params.CDMT1))

		if err != nil {
			return nil, err
		}
	}

	if len(params.CDNJ1) > 0 {
		part, err := writer.CreateFormFile("CDNJ1", "CDNJ1.jpg")

		if err != nil {
			return nil, err
		}

		//iocopy
		_, err = io.Copy(part, bytes.NewReader(params.CDNJ1))

		if err != nil {
			return nil, err
		}
	}

	if len(params.YHK) > 0 {
		part, err := writer.CreateFormFile("YHK", "YHK.jpg")

		if err != nil {
			return nil, err
		}

		//iocopy
		_, err = io.Copy(part, bytes.NewReader(params.YHK))

		if err != nil {
			return nil, err
		}
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	contentType := writer.FormDataContentType()

	req := http.NewHttpRequest("POST", reqUrl+ebi.Service, body)

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
func (ebi *ChinaEbiEnter) retData(ret []byte) (retData []byte, err error) {

	var baseReturn struct {
		Code string `json:"code"`
		Msg  string `json:"msg"`
	}

	err = json.Unmarshal(ret, &baseReturn)

	if err != nil {
		return
	}

	if baseReturn.Code != "000000" {
		err = errors.New(baseReturn.Msg)
		return
	}

	if err != nil {
		return
	}

	retData = ret
	return
}

// 图片信息
type ChinaEbiComposeAgreementConfig struct {
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
