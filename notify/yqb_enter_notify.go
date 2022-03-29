package notify

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/scholar-ink/payment/helper"
)

type YqbEnterNotifyData struct {
	ProductNo      string `json:"productNo"`
	ChannelNo      string `json:"channelNo"`
	ChannelBatchNo string `json:"channelBatchNo"`
	MerchantId     string `json:"merchantId"`
	OutMerchantNo  string `json:"outMerchantNo"`
	ApplicationNo  string `json:"applicationNo"`
	ChannelStatus  string `json:"channelStatus"`
	ChannelName    string `json:"channelName"`
	ChannelCode    string `json:"channelCode"`
	Remark         string `json:"remark"`
	AppStatus      string `json:"appStatus"`
	FailureReason  string `json:"failureReason"`
}

type YqbEnterProduct struct {
	Msg         string `json:"msg"`         //异常消息
	ProductType string `json:"productType"` //产品类型
	Status      string `json:"status"`      //状态0:待进件1:进件失败2:进件成功3:受理中，等待异步通知4:进件成功,通知成功
}

type YqbEnterConfig struct {
	AesKey string `json:"-"`
}

type YqbEnterNotify struct {
	*YqbEnterConfig
	notifyData *YqbEnterNotifyData
}

func (yqb *YqbEnterNotify) InitBaseConfig(config *YqbEnterConfig) {
	yqb.YqbEnterConfig = config
}

func (yqb *YqbEnterNotify) getNotifyData(ret string) error {

	if ret == "" {
		return errors.New("获取通知数据失败")
	}

	var data struct {
		Content string `json:"content"`
		Token   string `json:"token"`
	}

	err := json.Unmarshal([]byte(ret), &data)

	if err != nil {
		return errors.New("解析返回数据失败:" + err.Error())
	}

	aesKey, _ := base64.StdEncoding.DecodeString(yqb.AesKey)

	content, _ := base64.StdEncoding.DecodeString(data.Content)

	b, err := helper.AesDecrypt(content, aesKey) //ECB加密

	b = helper.PKCS7UnPadding(b)

	if err != nil {
		return err
	}

	notify := new(YqbEnterNotifyData)

	err = json.Unmarshal(b, notify)

	if err != nil {
		return errors.New("解析通知数据失败:" + err.Error())
	}

	yqb.notifyData = notify

	return nil
}

func (yqb *YqbEnterNotify) checkNotify() error {
	return yqb.verifySign()
}

func (yqb *YqbEnterNotify) verifySign() error {
	return nil
}

func (yqb *YqbEnterNotify) replyNotify(err error) string {
	if err != nil {
		fmt.Println(err)
		return err.Error()
	} else {
		return `{"respCode":"000000","respMsg":"成功"}`
	}
}

type YqbEnterCallBack func(data *YqbEnterNotifyData) error

func (yqb *YqbEnterNotify) Handle(ret string, callBack YqbEnterCallBack) string {
	err := yqb.getNotifyData(ret)

	if err != nil {
		return yqb.replyNotify(err)
	}

	err = callBack(yqb.notifyData)

	if err != nil {
		return yqb.replyNotify(err)
	}

	return yqb.replyNotify(nil)
}
