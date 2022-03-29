package withdrawal

import (
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"
	"log"

	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"net/url"
	"time"
)

const ZgUrl = "https://pay.g-pay.cn/entrustPay.html"

//中钢委托出款
type ZgWithdrawal struct {
	*ZgConfig
	ServiceName string `json:"service_name"`
	RsaMsg      string `json:"rsa_msg"`
	Md5Msg      string `json:"md_5_msg"`
}

type ZgConfig struct {
	PartnerId    string `json:"partner_id"`
	PfxData      []byte `json:"pfx_data"`
	CertPassWord string `json:"certPassWord"`
	Md5Key       string `json:"md_5_key"`
	Version      string `json:"version"`
}

type ZgHandleConfig struct {
	OutTradeNo      string `json:"out_trade_no"`
	AccountNo       string `json:"account_no"`
	Amount          string `json:"amount"`
	ServerReturnUrl string `json:"server_return_url"`
	Subject         string `json:"subject"`
}
type ZgQueryConfig struct {
	AccountNo  string `json:"account_no"`
	OutTradeNo string `json:"out_trade_no"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
}

type ZgHandleReturn struct {
	Code        string `json:"code"`
	Md5Msg      string `json:"md5msg"`
	Msg         string `json:"msg"`
	PartnerId   string `json:"partner_id"`
	RsaMsg      string `json:"rsamsg"`
	ServiceName string `json:"service_name"`
	Version     string `json:"version"`
	OrderNo     string `json:"orderNo"`
	PayResult   string `json:"payResult"`
	TotalAmount string `json:"totalAmount"`
	ReturnMsg   string `json:"returnMsg"`
	ErrorMsg    string `json:"error_msg"`
}

type ZgQueryReturn struct {
	Code          string `json:"code"`
	Md5Msg        string `json:"md5msg"`
	Msg           string `json:"msg"`
	PartnerId     string `json:"partner_id"`
	RsaMsg        string `json:"rsamsg"`
	ServiceName   string `json:"service_name"`
	Version       string `json:"version"`
	OrderNo       string `json:"order_no"`
	PayResult     string `json:"pay_result"`
	ResultMsg     string `json:"result_msg"`
	SuccessAmount string `json:"success_amount"`
}

type ZgBalanceConfig struct {
	AccountNo string `json:"accountNo"`
}

type ZgAccount struct {
	AccountType   string `json:"accountType"`
	CanUseBalance int64  `json:"canUseBalance"`
	FreezeAmount  int64  `json:"freezeAmount"`
	Status        string `json:"status"`
}

func (zg *ZgWithdrawal) InitConfig(config *ZgConfig) {
	config.Version = "3.0"
	zg.ZgConfig = config
}

func (zg *ZgWithdrawal) Handle(config *ZgHandleConfig) (*ZgHandleReturn, error) {
	zg.ServiceName = "fund_entrust_pay"
	zg.buildData(config)
	zg.setSign()
	ret, err := zg.SendReq(ZgUrl)

	if err != nil {
		return nil, err
	}

	retData, _ := url.QueryUnescape(string(ret))

	log.Println("[ZgWithdrawal.Handle] 返回数据：", retData)

	zgHandleReturn := new(ZgHandleReturn)

	json.Unmarshal([]byte(retData), zgHandleReturn)

	if zgHandleReturn.Code != "0000" {
		return nil, errors.New(zgHandleReturn.Msg)
	}

	b, err := helper.Rsa1Decrypt(zg.PfxData, zgHandleReturn.RsaMsg, zg.CertPassWord)

	if err != nil {
		zgHandleReturn.ErrorMsg = err.Error()

		return zgHandleReturn, nil
	}

	type HandlerResponse struct {
		DivideAccountResponseBody struct {
			OrderNo     string `json:"orderNo"`
			PayResult   string `json:"payResult"`
			ReturnMsg   string `json:"returnMsg"`
			TotalAmount string `json:"totalAmount"`
		} `json:"divideAccountResponseBody"`
	}

	handlerResponse := new(HandlerResponse)

	err = json.Unmarshal(b, handlerResponse)

	if err != nil {

		zgHandleReturn.ErrorMsg = err.Error()

		return zgHandleReturn, nil
	}

	zgHandleReturn.PayResult = handlerResponse.DivideAccountResponseBody.PayResult
	zgHandleReturn.OrderNo = handlerResponse.DivideAccountResponseBody.OrderNo
	zgHandleReturn.TotalAmount = handlerResponse.DivideAccountResponseBody.TotalAmount
	zgHandleReturn.ReturnMsg = handlerResponse.DivideAccountResponseBody.ReturnMsg

	return zgHandleReturn, nil
}

func (zg *ZgWithdrawal) Query(config *ZgQueryConfig) (*ZgQueryReturn, error) {
	zg.ServiceName = "fund_entrust_pay_query"
	zg.buildData(config)
	zg.setSign()
	ret, err := zg.SendReq(ZgUrl)

	if err != nil {
		return nil, err
	}

	retData, _ := url.QueryUnescape(string(ret))

	log.Println("[ZgWithdrawal.Handle] 返回数据：", retData)
	zgQueryReturn := new(ZgQueryReturn)

	json.Unmarshal([]byte(retData), zgQueryReturn)

	if zgQueryReturn.Code != "0000" {
		return nil, errors.New(zgQueryReturn.Msg)
	}

	b, err := helper.Rsa1Decrypt(zg.PfxData, zgQueryReturn.RsaMsg, zg.CertPassWord)

	if err != nil {
		return nil, err
	}

	type HandlerResponse struct {
		EntrustProxyPayQueryResponseBody struct {
			OrderNo       string `json:"order_no"`
			OutTradeNo    string `json:"out_trade_no"`
			PayResult     string `json:"pay_result"`
			ResultMsg     string `json:"result_msg"`
			SuccessAmount string `json:"success_amount"`
		} `json:"entrustProxyPayQueryResponseBody"`
	}

	handlerResponse := new(HandlerResponse)

	err = json.Unmarshal(b, handlerResponse)

	if err != nil {
		return nil, err
	}

	zgQueryReturn.PayResult = handlerResponse.EntrustProxyPayQueryResponseBody.PayResult
	zgQueryReturn.OrderNo = handlerResponse.EntrustProxyPayQueryResponseBody.OrderNo
	zgQueryReturn.SuccessAmount = handlerResponse.EntrustProxyPayQueryResponseBody.SuccessAmount
	zgQueryReturn.ResultMsg = handlerResponse.EntrustProxyPayQueryResponseBody.ResultMsg

	return zgQueryReturn, nil
}

func (zg *ZgWithdrawal) Balance(config *ZgBalanceConfig) ([]*ZgAccount, error) {
	zg.Version = "2.0"
	zg.ServiceName = "account_gpay_query"
	err := zg.buildData(config)

	if err != nil {
		return nil, err
	}

	zg.setSign()
	ret, err := zg.SendReq("https://pay.g-pay.cn/account/queryBalance.html")

	if err != nil {
		return nil, err
	}

	retData, _ := url.QueryUnescape(string(ret))

	zgQueryReturn := new(ZgQueryReturn)

	json.Unmarshal([]byte(retData), zgQueryReturn)

	if zgQueryReturn.Code != "0000" {
		return nil, errors.New(zgQueryReturn.Msg)
	}

	b, err := helper.Rsa1DecryptBcd(zg.PfxData, zgQueryReturn.RsaMsg, zg.CertPassWord)

	log.Println("[ZgWithdrawal.Balance] 返回数据：", string(b))

	if err != nil {
		return nil, err
	}

	type QueryResponse struct {
		Body struct {
			AccountList []*ZgAccount `json:"accountList"`
		} `json:"body"`
	}

	queryResponse := new(QueryResponse)

	json.Unmarshal(b, queryResponse)

	return queryResponse.Body.AccountList, nil
}

func (zg *ZgWithdrawal) buildData(params interface{}) error {

	b, _ := json.Marshal(map[string]interface{}{
		"head": map[string]interface{}{
			"serviceName": zg.ServiceName,
			"traceNo":     uuid.New().String(),
			"version":     zg.Version,
			"charset":     "utf-8",
			"senderId":    zg.PartnerId,
			"sendTime":    time.Now().Format("20060102150405"),
		},
		"body": params,
	})

	data := fmt.Sprintf("partner_id=%s&service_name=%s&data=%s", zg.PartnerId, zg.ServiceName, url.QueryEscape(string(b)))

	var encryptData string

	var err error

	if zg.ServiceName == "account_gpay_query" {
		encryptData, err = helper.Rsa1EncryptBcd(zg.PfxData, []byte(data), zg.CertPassWord)
	} else {
		encryptData, err = helper.Rsa1Encrypt(zg.PfxData, []byte(data), zg.CertPassWord)
	}

	if err != nil {
		return err
	}

	zg.RsaMsg = encryptData
	return nil
}

func (zg *ZgWithdrawal) setSign() error {

	zg.Md5Msg = helper.Md5(zg.RsaMsg + zg.Md5Key)

	return nil
}

func (zg *ZgWithdrawal) SendReq(reqUrl string) (b []byte, err error) {

	reqUrl = fmt.Sprintf("%s?partner_id=%s&service_name=%s&rsamsg=%s&md5msg=%s&version=%s",
		reqUrl, zg.PartnerId, zg.ServiceName, zg.RsaMsg, zg.Md5Msg, zg.Version)

	log.Println("[withdrawal.zg."+zg.ServiceName+"] req start：", reqUrl)

	rsp, err := http.Client.Get(reqUrl)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	return
}
