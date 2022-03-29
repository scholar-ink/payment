package notify

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/map"
)

type YqbNotifyData struct {
	BusinessCode       string `json:"businessCode"`       //业务结果
	BusinessMsg        string `json:"businessMsg"`        //业务结果描述
	TradeOrderNo       string `json:"tradeOrderNo"`       //平安付订单号
	ChannelOrderNo     string `json:"channelOrderNo"`     //渠道商户订单号
	MerTradeNo         string `json:"merTradeNo"`         //商户订单号
	MerchantId         string `json:"merchantId"`         //平安付商户号
	TotalAmount        string `json:"totalAmount"`        //订单总金额，1代表一分
	PayTime            string `json:"payTime"`            //支付时间，交易成功时返回
	PlatformMerchantId string `json:"platformMerchantId"` //平安付聚合平台服务商号
	TradeType          string `json:"tradeType"`          //交易类型，01-公众号支付 02-扫码支付
}

type YqbNotifyHeaderData struct {
	ResultCode string `json:"resultCode"`
	ResultMsg  string `json:"resultMsg"`
	SignType   string `json:"signType"`
}

type YqbNotifySignData struct {
	SignContent string `json:"signContent"`
}

type YqbNotify struct {
	Md5Key           string               `json:"-"`
	PrivateKey       string               `json:"-"`
	Sign             *YqbNotifySignData   `json:"sign"`
	NotifyData       *YqbNotifyData       `json:"body"`
	NotifyHeaderData *YqbNotifyHeaderData `json:"head"`
}

func (yqb *YqbNotify) getNotifyData(ret string) error {

	if ret == "" {
		return errors.New("获取通知数据失败")
	}

	err := json.Unmarshal([]byte(ret), &yqb)

	if err != nil {
		return errors.New("解析返回数据失败:" + err.Error())
	}

	return nil
}

func (yqb *YqbNotify) verifySign() error {

	bodyMap := helper.Struct2Map(yqb.NotifyData)

	helper.KSort(&bodyMap)

	b, _ := json.Marshal(bodyMap)

	bodyData := string(b)

	headerMap := maps.Struct2Map(yqb.NotifyHeaderData)

	helper.KSort(&headerMap)

	b, _ = json.Marshal(headerMap)

	headerData := string(b)

	return helper.Sha1WithRsaVerifyHex([]byte(headerData+bodyData), yqb.Sign.SignContent, yqb.Md5Key)
}

func (yqb *YqbNotify) replyNotify(err error) string {

	type ResponseData struct {
		BusinessCode string `json:"businessCode"` //业务结果
		BusinessMsg  string `json:"businessMsg"`  //业务结果描述
		NeedRetry    string `json:"needRetry"`    //是否需要重发，0:不需要重发 1:需要重发
	}

	type Response struct {
		Sign               *YqbNotifySignData   `json:"sign"`
		ResponseData       *ResponseData        `json:"body"`
		ResponseHeaderData *YqbNotifyHeaderData `json:"head"`
	}

	responseSignData := new(YqbNotifySignData)
	responseData := new(ResponseData)
	responseHeaderData := new(YqbNotifyHeaderData)

	responseHeaderData.ResultCode = "000000"
	responseHeaderData.ResultMsg = "成功"
	responseHeaderData.SignType = "RSA"

	if err != nil {
		responseData.BusinessCode = "2000"
		responseData.BusinessMsg = err.Error()
		responseData.NeedRetry = "1"
	} else {
		responseData.BusinessCode = "1000"
		responseData.BusinessMsg = "通知成功"
		responseData.NeedRetry = "0"
	}

	responseDataMap := helper.Struct2Map(responseData)

	helper.KSort(&responseDataMap)

	b, _ := json.Marshal(responseDataMap)

	body := string(b)

	headerMap := maps.Struct2Map(responseHeaderData)

	helper.KSort(&headerMap)

	b, _ = json.Marshal(headerMap)

	header := string(b)

	b, err = helper.Sha1WithRsaSign([]byte(header+body), yqb.PrivateKey)

	if err != nil {
		return err.Error()
	}

	responseSignData.SignContent = base64.StdEncoding.EncodeToString(b)

	response := &Response{
		Sign:               responseSignData,
		ResponseData:       responseData,
		ResponseHeaderData: responseHeaderData,
	}

	b, _ = json.Marshal(response)

	return string(b)
}

type YqbCallBack func(data *YqbNotifyData) error

type YqbMd5KeyCallBack func(merchantNo string) (md5Key, privateKey string, err error)

func (yqb *YqbNotify) Handle(ret string, md5CallBack YqbMd5KeyCallBack, callBack YqbCallBack) string {

	err := yqb.getNotifyData(ret)

	if err != nil {
		return yqb.replyNotify(err)
	}

	md5Key, privateKey, err := md5CallBack(yqb.NotifyData.MerchantId)

	if err != nil {
		return yqb.replyNotify(err)
	}

	yqb.Md5Key = md5Key
	yqb.PrivateKey = privateKey

	err = yqb.verifySign()

	if err != nil {
		return yqb.replyNotify(err)
	}

	err = callBack(yqb.NotifyData)

	if err != nil {
		return yqb.replyNotify(err)
	}

	return yqb.replyNotify(nil)
}
