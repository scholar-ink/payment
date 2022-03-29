package enter

import (
	"bytes"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"

	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

const (
	DuoLaBaoEnterUrl = "https://openapi.duolabao.com"
)

type DuoLaBaoEnter struct {
	*DuoLaBaoConfig
	Token     string `json:"-"`
	Path      string `json:"path"`
	Timestamp string `json:"-"`
}

type DuoLaBaoConfig struct {
	AccessKey string `json:"-"`
	SecretKey string `json:"-"`
}

type DuoLaBaoRate struct {
	RateChannel string `json:"rate_channel"`
	FeeRate     string `json:"fee_rate"`
}

type DuoLaBaoCreateConf struct {
	AgentNum             string `json:"agentNum"`                //代理商编号
	FullName             string `json:"fullName"`                //哆啦宝商户全称
	ShortName            string `json:"shortName"`               //哆啦宝商户简称
	Industry             string `json:"industry"`                //商户所属行业
	Province             string `json:"province"`                //商户所属省份
	City                 string `json:"city"`                    //商户所属城市
	District             string `json:"district"`                //地区
	LinkMan              string `json:"linkMan"`                 //商户联系人
	LinkPhone            string `json:"linkPhone"`               //商户联系电话
	CustomerType         string `json:"customerType"`            //商户类型COMPANY(公司企业)/ PERSON(个人) /INDIVIDUALBISS(个体工商户)
	CertificateType      string `json:"certificateType"`         //证件类型,IDENTIFICATION(身份证)
	CertificateCode      string `json:"certificateCode"`         //证件编号
	CertificateName      string `json:"certificateName"`         //证件人姓名
	CertificateStartDate string `json:"certificateStartDate"`    //法人证件开始日期
	CertificateEndDate   string `json:"certificateEndDate"`      //法人证件结束日期
	ContactPhoneNum      string `json:"contactPhoneNum"`         //联系人手机号-认证使用
	LinkManId            string `json:"linkManId"`               //联系人身份证号-认证使用
	PostalAddress        string `json:"postalAddress,omitempty"` //营业执照注册地址（个人可不传）
}

type DuoLaBaoCreateReturn struct {
	CustomerNum string `json:"customerNum"`
}

type DuoLaBaoBank struct {
	Num  string `json:"num"`  //支付类型编号
	Rate string `json:"rate"` //费率
}

type DuoLaBaoSettleConf struct {
	CustomerNum                 string          `json:"customerNum"`                 //商户编号
	BankAccountName             string          `json:"bankAccountName"`             //银行账户名称
	BankAccountNum              string          `json:"bankAccountNum"`              //银行账户编号
	Province                    string          `json:"province"`                    //商户所属省份
	City                        string          `json:"city"`                        //商户所属城市
	BankBranchName              string          `json:"bankBranchName"`              //银行分行名称
	BankName                    string          `json:"bankName"`                    //银行名称
	SettleAmount                string          `json:"settleAmount"`                //结算金额
	PayBankList                 []*DuoLaBaoBank `json:"payBankList"`                 //结算金额
	AccountType                 string          `json:"accountType"`                 //账户类型PUBLIC(对公)/ PRIVATE(对私)
	Phone                       string          `json:"phone"`                       //银行预留手机号
	SettlerCertificateCode      string          `json:"settlerCertificateCode"`      //结算人身份证号
	SettlerCertificateStartDate string          `json:"settlerCertificateStartDate"` //结算人身份证开始时间
	SettlerCertificateEndDate   string          `json:"settlerCertificateEndDate"`   //结算人身份证结束时间
}

type DuoLaBaoSettleReturn struct {
	SettleNum string `json:"settleNum"`
}

type DuoLaBaoShopConf struct {
	AgentNum     string `json:"agentNum"`               //代理商编号
	CustomerNum  string `json:"customerNum"`            //商户编号
	ShopName     string `json:"shopName"`               //店铺名称
	Address      string `json:"address"`                //店铺地址
	OneIndustry  string `json:"oneIndustry"`            //店铺一级行业
	TwoIndustry  string `json:"twoIndustry"`            //店铺二级行业
	MobilePhone  string `json:"mobilePhone"`            //店铺联系人手机号
	MapLng       string `json:"mapLng"`                 //经度
	MapLat       string `json:"mapLat"`                 //纬度
	MicroBizType string `json:"microBizType,omitempty"` //经营类型 (个人商户必填)
}

type DuoLaBaoShopReturn struct {
	ShopNum string `json:"shopNum"`
}
type DuoLaBaoCompleteConf struct {
	CustomerNum      string `json:"customerNum"`                //商户编号
	LicenseId        string `json:"licenseId,omitempty"`        //营业执照
	LicenseStartTime string `json:"licenseStartTime,omitempty"` //营业执照起始日期
	LicenseEndTime   string `json:"licenseEndTime,omitempty"`   //营业执照结束日期
	CallbackUrl      string `json:"callbackUrl"`                //回调地址
}

type DuoLaBaoCompleteReturn struct {
	CustomerNum   string `json:"customerNum"`
	DeclareStatus string `json:"declareStatus"`
	FullName      string `json:"fullName"`
}

type DuoLaBaoUploadConf struct {
	AttachType  string `json:"attachType"`  //附件类型
	CustomerNum string `json:"customerNum"` //商户编号
	File        string `json:"file"`        //文件base64
}

type DuoLaBaoUploadReturn struct {
	AttachNum string `json:"attachNum"`
	FileName  string `json:"fileName"`
	Url       string `json:"url"`
}

func (base *DuoLaBaoEnter) InitConfig(config *DuoLaBaoConfig) {
	base.Timestamp = strconv.Itoa(int(time.Now().Unix()))
	base.DuoLaBaoConfig = config
}

func (base *DuoLaBaoEnter) Create(config *DuoLaBaoCreateConf) (*DuoLaBaoCreateReturn, error) {
	base.Path = "/v2/agent/declare/customerinfo/create"

	ret, err := base.sendReq(DuoLaBaoEnterUrl, config)
	if err != nil {
		return nil, err
	}

	ret, err = base.RetData(ret)
	if err != nil {
		return nil, err
	}

	enterReturn := new(DuoLaBaoCreateReturn)
	err = json.Unmarshal(ret, &enterReturn)
	if err != nil {
		return nil, err
	}
	return enterReturn, nil
}

func (base *DuoLaBaoEnter) Settle(config *DuoLaBaoSettleConf) (*DuoLaBaoSettleReturn, error) {
	base.Path = "/v2/agent/declare/settleinfo/create"

	ret, err := base.sendReq(DuoLaBaoEnterUrl, config)
	if err != nil {
		return nil, err
	}

	ret, err = base.RetData(ret)
	if err != nil {
		return nil, err
	}

	settleReturn := new(DuoLaBaoSettleReturn)
	err = json.Unmarshal(ret, &settleReturn)
	if err != nil {
		return nil, err
	}
	return settleReturn, nil
}

func (base *DuoLaBaoEnter) Shop(config *DuoLaBaoShopConf) (*DuoLaBaoShopReturn, error) {
	base.Path = "/v1/agent/declare/shopinfo/create"

	ret, err := base.sendReq(DuoLaBaoEnterUrl, config)
	if err != nil {
		return nil, err
	}

	ret, err = base.RetData(ret)
	if err != nil {
		return nil, err
	}

	shopReturn := new(DuoLaBaoShopReturn)
	err = json.Unmarshal(ret, &shopReturn)
	if err != nil {
		return nil, err
	}
	return shopReturn, nil
}
func (base *DuoLaBaoEnter) Complete(config *DuoLaBaoCompleteConf) (*DuoLaBaoCompleteReturn, error) {
	base.Path = "/v2/agent/declare/complete"

	ret, err := base.sendReq(DuoLaBaoEnterUrl, config)
	if err != nil {
		return nil, err
	}

	ret, err = base.RetData(ret)
	if err != nil {
		return nil, err
	}

	completeReturn := new(DuoLaBaoCompleteReturn)
	err = json.Unmarshal(ret, &completeReturn)
	if err != nil {
		return nil, err
	}
	return completeReturn, nil
}
func (base *DuoLaBaoEnter) Upload(config *DuoLaBaoUploadConf) (*DuoLaBaoUploadReturn, error) {
	base.Path = "/v2/agent/declare/attach/upload"

	ret, err := base.sendReq(DuoLaBaoEnterUrl, config)
	if err != nil {
		return nil, err
	}

	ret, err = base.RetData(ret)
	if err != nil {
		return nil, err
	}

	uploadReturn := new(DuoLaBaoUploadReturn)
	err = json.Unmarshal(ret, &uploadReturn)
	if err != nil {
		return nil, err
	}
	return uploadReturn, nil
}

func (base *DuoLaBaoEnter) IndustryList() error {
	base.Path = "/v1/agent/industry/second/list/code/10081614701426099450047"

	ret, err := base.sendGetReq(DuoLaBaoEnterUrl)
	if err != nil {
		return err
	}

	ret, err = base.RetData(ret)
	if err != nil {
		return err
	}
	return nil
}

func (base *DuoLaBaoEnter) makeSign(requestJson string) {

	fmt.Println(fmt.Sprintf("secretKey=%s&timestamp=%s&path=%s&body=%s", base.SecretKey, base.Timestamp, base.Path, requestJson))

	base.Token = strings.ToUpper(helper.Sha1(fmt.Sprintf("secretKey=%s&timestamp=%s&path=%s&body=%s", base.SecretKey, base.Timestamp, base.Path, requestJson)))
}

func (base *DuoLaBaoEnter) sendReq(reqUrl string, params interface{}) (b []byte, err error) {

	b, err = json.Marshal(params)

	if err != nil {
		return nil, err
	}

	base.makeSign(string(b))

	req := http.NewHttpRequest("POST", reqUrl+base.Path, bytes.NewBuffer(b))

	req.Header["token"] = []string{base.Token}
	req.Header["accessKey"] = []string{base.AccessKey}
	req.Header["timestamp"] = []string{base.Timestamp}
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("accessKey", base.AccessKey)
	//req.Header.Set("timestamp", base.Timestamp)
	//req.Header.Set("token", base.Token)

	fmt.Println(req.Header)

	rsp, err := http.Client.Do(req)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	return
}

func (base *DuoLaBaoEnter) makeGetSign() {

	fmt.Println(fmt.Sprintf("secretKey=%s&timestamp=%s&path=%s", base.SecretKey, base.Timestamp, base.Path))

	base.Token = strings.ToUpper(helper.Sha1(fmt.Sprintf("secretKey=%s&timestamp=%s&path=%s", base.SecretKey, base.Timestamp, base.Path)))
}

func (base *DuoLaBaoEnter) sendGetReq(reqUrl string) (b []byte, err error) {

	base.makeGetSign()

	req := http.NewHttpRequest("GET", reqUrl+base.Path, nil)

	req.Header["token"] = []string{base.Token}
	req.Header["accessKey"] = []string{base.AccessKey}
	req.Header["timestamp"] = []string{base.Timestamp}
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("accessKey", base.AccessKey)
	//req.Header.Set("timestamp", base.Timestamp)
	//req.Header.Set("token", base.Token)

	fmt.Println(req.Header)

	rsp, err := http.Client.Do(req)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	return
}

func (base *DuoLaBaoEnter) RetData(ret []byte) (b []byte, err error) {

	result := new(struct {
		Error struct {
			ErrorCode string `json:"errorCode"`
			ErrorMsg  string `json:"errorMsg"`
		} `json:"error"`
		Data   interface{} `json:"data"`
		Result string      `json:"result"`
	})

	err = json.Unmarshal(ret, &result)

	if err != nil {
		return
	}

	if result.Result != "success" {
		return nil, errors.New(result.Error.ErrorMsg)
	}

	return json.Marshal(result.Data)
}
