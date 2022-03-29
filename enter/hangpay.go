package enter

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"
	"github.com/scholar-ink/payment/util/map"
	"io"
	"io/ioutil"
	"mime/multipart"
	"time"
)

const (
	HangPayEnterUrl = "http://api.hangpay.cn/"
)

type HangPayEnter struct {
	*HangPayConfig
	Service     string `json:"-"`
	Version     string `json:"version"`
	Sign        string `json:"sign"`
	EncryKey    string `json:"encryKey"`
	RequestData string `json:"requestData"`
	RequestTime string `json:"requestTime"`
}

type HangPayConfig struct {
	MerchantNo string `json:"merchantNo"`
	Key        string `json:"-"`
	PublicKey  string `json:"-"`
}

type HangBank struct {
	OpenName    string `json:"openName"`    //开户名
	AccountType string `json:"accountType"` //账户类型 1.对公 2.对法人，个人只能对私（对法人）
	OpenBank    string `json:"openBank"`    //开户行
	CardNo      string `json:"cardNo"`      //银行卡号
	OpenBranch  string `json:"openBranch"`  //开户支行
}

type HangPayCreateConf struct {
	NetworkNo           string    `json:"networkNo"`                     //入网流水号
	NotifyUrl           string    `json:"notifyUrl"`                     //通知地址
	MchLevel            string    `json:"mchLevel"`                      //入网类型
	MchType             string    `json:"mchType"`                       //商户类型
	MchName             string    `json:"mchName"`                       //商户名
	BusinessLicense     string    `json:"businessLicense"`               //营业执照号
	LegalPersonName     string    `json:"legalpersonName"`               //法人名
	LegalPersonIdCard   string    `json:"legalpersonIdcard"`             //法人身份证号码
	BusinessModel       string    `json:"businessModel"`                 //经营模式
	ContactName         string    `json:"contactName"`                   //联系人名称
	ContactTel          string    `json:"contactTel"`                    //联系人电话
	Province            string    `json:"province"`                      //商户所在省份
	City                string    `json:"city"`                          //商户所在城市
	District            string    `json:"district"`                      //商户所在地区
	MchAddress          string    `json:"mchAddress"`                    //商户地址
	BankCard            *HangBank `json:"bankCard"`                      //入账信息
	AliPayRate          string    `json:"alipayRate"`                    //支付宝费率
	WxChatRate          string    `json:"wxChatRate"`                    //微信费率
	UnionPayRate        string    `json:"unionPayRate"`                  //银联费率
	WithdrawRate        string    `json:"withdrawRate"`                  //提现手续费
	QuickRate           string    `json:"quickRate,omitempty"`           //快捷费率
	TransferChargesRate string    `json:"transferChargesRate,omitempty"` //转账费率
	ZipFile             []byte    `json:"-"`
}

type HangPayCreateReturn struct {
	ResideMerchantNo string `json:"reside_mer_no"`
}

type HangPayBalanceConf struct {
	SubMerchantNo string `json:"subMerchantNo"` //子商户号
}
type HangPayBalanceReturn struct {
	Amount        string `json:"amount"`        //余额
	SubMerchantNo string `json:"subMerchantNo"` //子商户号
	MerchantNo    string `json:"merchantNo"`    //商户号
}

type HangPayWithdrawConf struct {
	SubMerchantNo string `json:"subMerchantNo"` //子商户号
	Money         string `json:"money"`         //提现金额
	NotifyUrl     string `json:"notifyUrl"`     //通知地址
	WithdrawNo    string `json:"withdrawNo"`    //提现订单号
}
type HangPayWithdrawReturn struct {
	Money string `json:"money"` //余额
}

func (hang *HangPayEnter) InitConfig(config *HangPayConfig) {
	hang.HangPayConfig = config
}

func (hang *HangPayEnter) Create(config *HangPayCreateConf) (*HangPayCreateReturn, error) {
	hang.Service = "api/access/network"
	hang.Version = "1.0.0"

	ret, err := hang.sendUploadReq(HangPayEnterUrl, config)
	if err != nil {
		return nil, err
	}

	err = hang.retData(ret)

	if err != nil {
		return nil, err
	}

	enterReturn := new(HangPayCreateReturn)
	err = json.Unmarshal(ret, &enterReturn)
	if err != nil {
		return nil, err
	}

	return enterReturn, nil
}
func (hang *HangPayEnter) Balance(config *HangPayBalanceConf) (*HangPayBalanceReturn, error) {
	hang.Service = "api/mch/balance"
	hang.Version = "1.0.0"

	ret, err := hang.sendReq(HangPayEnterUrl, config)
	if err != nil {
		return nil, err
	}

	err = hang.retData(ret)

	if err != nil {
		return nil, err
	}

	balanceReturn := new(HangPayBalanceReturn)
	err = json.Unmarshal(ret, &balanceReturn)
	if err != nil {
		return nil, err
	}

	return balanceReturn, nil
}
func (hang *HangPayEnter) Withdraw(config *HangPayWithdrawConf) (*HangPayWithdrawReturn, error) {
	hang.Service = "api/withdraw/extract"
	hang.Version = "1.0.0"

	ret, err := hang.sendReq(HangPayEnterUrl, config)
	if err != nil {
		return nil, err
	}

	err = hang.retData(ret)

	if err != nil {
		return nil, err
	}

	withdrawReturn := new(HangPayWithdrawReturn)
	err = json.Unmarshal(ret, &withdrawReturn)
	if err != nil {
		return nil, err
	}

	return withdrawReturn, nil
}
func (hang *HangPayEnter) makeSign(params interface{}) {

	aesKey := helper.NonceStr()[0:16]

	b, _ := json.Marshal(params)

	cipherText := helper.PKCS7Padding(b, 16)

	b, _ = helper.AesEncrypt(cipherText, []byte(aesKey)) //ECB加密

	fmt.Println(base64.StdEncoding.EncodeToString(b))

	hang.RequestData = base64.StdEncoding.EncodeToString(b)

	hang.RequestTime = time.Now().Format("20060102150405")

	url := fmt.Sprintf("requestData=%s&requestTime=%s&merchantNo=%s", hang.RequestData, hang.RequestTime, hang.MerchantNo)

	b, _ = helper.Sha1WithRsaSignPkcs8([]byte(url), hang.Key)

	hang.Sign = base64.StdEncoding.EncodeToString(b)

	hang.EncryKey, _ = helper.Rsa2Encrypt3([]byte(base64.StdEncoding.EncodeToString([]byte(aesKey))), hang.PublicKey)

}

func (hang *HangPayEnter) sendReq(reqUrl string, params interface{}) (b []byte, err error) {

	hang.makeSign(params)

	b, err = json.Marshal(params)

	mapData := make(map[string]interface{})

	mapData["requestTime"] = hang.RequestTime
	mapData["version"] = hang.Version
	mapData["merchantNo"] = hang.MerchantNo
	mapData["requestData"] = hang.RequestData
	mapData["sign"] = hang.Sign
	mapData["encryKey"] = hang.EncryKey

	values := maps.Map2Values(&mapData)

	rsp, err := http.Client.PostForm(reqUrl+hang.Service, values)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	return
}
func (hang *HangPayEnter) sendUploadReq(reqUrl string, params *HangPayCreateConf) (b []byte, err error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	hang.makeSign(params)

	mapData := make(map[string]interface{})

	mapData["requestTime"] = hang.RequestTime
	mapData["version"] = hang.Version
	mapData["merchantNo"] = hang.MerchantNo
	mapData["requestData"] = hang.RequestData
	mapData["sign"] = hang.Sign
	mapData["encryKey"] = hang.EncryKey

	for key, value := range mapData {
		writer.WriteField(key, value.(string))
	}

	//关键的一步操作
	part, err := writer.CreateFormFile("file", "file.zip")

	if err != nil {
		return nil, err
	}

	//iocopy
	_, err = io.Copy(part, bytes.NewReader(params.ZipFile))

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	contentType := writer.FormDataContentType()

	rsp, err := http.Client.Post(reqUrl+hang.Service, contentType, body)

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

func (hang *HangPayEnter) retData(ret []byte) (err error) {

	var baseReturn struct {
		Code string `json:"code"`
		Msg  string `json:"msg"`
	}

	err = json.Unmarshal(ret, &baseReturn)

	if err != nil {
		return
	}

	if baseReturn.Code != "SUCCESS" {
		err = errors.New(baseReturn.Msg)
		return
	}

	return
}
