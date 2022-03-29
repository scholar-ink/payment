package notify

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/scholar-ink/payment/helper"
)

type ZgEnterNotifyData struct {
	AccountNo    string            `json:"accountNo"`
	MerchantName string            `json:"merchantName"`
	SmMd5Key     string            `json:"smMd5Key"`
	SmMerchantNo string            `json:"smMerchantNo"`
	Status       string            `json:"status"`   //00:审核中、01:审核通过、02:审核拒绝
	Products     []*ZgEnterProduct `json:"products"` //扫码产品
}

type ZgEnterProduct struct {
	Msg         string `json:"msg"`         //异常消息
	ProductType string `json:"productType"` //产品类型
	Status      string `json:"status"`      //状态0:待进件1:进件失败2:进件成功3:受理中，等待异步通知4:进件成功,通知成功
}

type ZgConfig struct {
	AgentNo      string `json:"agentNo"`
	Key          string `json:"key"`
	PfxData      []byte `json:"pfx_data"`
	CertPassWord string `json:"certPassWord"`
}

type ZgEnterNotify struct {
	*ZgConfig
	EncryptType  string `json:"encryptType"`
	EncryptData  string `json:"encryptData"`
	SignData     string `json:"signData"`
	ResponseCode string `xml:"responseCode" json:"responseCode"`
	ResponseMsg  string `xml:"responseMsg" json:"responseMsg"`
	notifyData   *ZgEnterNotifyData
}

func (zg *ZgEnterNotify) InitBaseConfig(config *ZgConfig) {
	zg.ZgConfig = config
}

func (zg *ZgEnterNotify) getNotifyData(ret string) error {

	if ret == "" {
		return errors.New("获取通知数据失败")
	}

	err := json.Unmarshal([]byte(ret), zg)

	if err != nil {
		return errors.New("解析返回数据失败:" + err.Error())
	}

	err = zg.checkNotify()

	if err != nil {
		return err
	}

	retData, err := helper.Rsa1Decrypt(zg.PfxData, zg.EncryptData, zg.CertPassWord)

	if err != nil {
		return err
	}

	notify := new(ZgEnterNotifyData)

	err = json.Unmarshal(retData, notify)

	if err != nil {
		return errors.New("解析通知数据失败:" + err.Error())
	}

	zg.notifyData = notify

	return nil
}

func (zg *ZgEnterNotify) checkNotify() error {
	if zg.ResponseCode != "0000" {
		return errors.New("中钢返回错误" + zg.ResponseMsg)
	}
	return zg.verifySign()
}

func (zg *ZgEnterNotify) verifySign() error {
	return nil
}

func (zg *ZgEnterNotify) replyNotify(err error) bool {

	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}
}

type ZgEnterCallBack func(data *ZgEnterNotifyData) error

func (zg *ZgEnterNotify) Handle(ret string, callBack ZgEnterCallBack) bool {
	err := zg.getNotifyData(ret)

	if err != nil {
		return zg.replyNotify(err)
	}

	err = callBack(zg.notifyData)

	if err != nil {
		return zg.replyNotify(err)
	}

	return zg.replyNotify(nil)
}
