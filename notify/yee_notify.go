package notify

import (
	"github.com/scholar-ink/payment/helper"

	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

type YeeNotifyData struct {
	MerchantNo    string `json:"merchantNo"`
	OrderId       string `json:"orderId"`
	UniqueOrderNo string `json:"uniqueOrderNo"`
	BankTrxId     string `json:"bankTrxId"`
	Status        string `json:"status"`
	OrderAmount   string `json:"orderAmount"`
	PayAmount     string `json:"payAmount"`
}

type YeeNotify struct {
	ParentMerchantNo string `json:"parentMerchantNo"`
	PrivateKey       string `json:"-"`
	notifyData       *YeeNotifyData
}

func (yee *YeeNotify) getNotifyData(ret string) error {

	if ret == "" {
		return errors.New("获取通知数据失败")
	}

	args := strings.Split(ret, "$")
	if 4 != len(args) {
		return errors.New("response has wrong args")
	}

	randomKey, err := yee.Base64Decode(args[0])
	if err != nil {
		return errors.New("Base64Decode args[0] fail," + err.Error())
	}

	encryptedData, err := yee.Base64Decode(args[1])

	if err != nil {
		return errors.New("Base64Decode args[1] fail," + err.Error())
	}

	b, err := helper.Rsa2Decrypt2(randomKey, yee.PrivateKey)

	origin, err := helper.AesDecrypt(encryptedData, b)

	if err != nil {
		return err
	}

	originStr := string(origin)

	originStr = originStr[:strings.LastIndex(originStr, "$")]

	notify := new(YeeNotifyData)

	err = json.Unmarshal([]byte(originStr), notify)

	if err != nil {
		return errors.New("解析通知数据失败:" + err.Error())
	}

	yee.notifyData = notify

	return nil
}

func (yee *YeeNotify) verifySign() error {
	return nil
}

func (yee *YeeNotify) replyNotify(err error) string {

	if err != nil {
		return err.Error()
	} else {
		return "SUCCESS"
	}
}

type YeeCallBack func(data *YeeNotifyData) error
type YeeMd5KeyCallBack func(merchantNo string) (privateKey string, err error)

func (yee *YeeNotify) Handle(parentMerchantNo, ret string, md5CallBack YeeMd5KeyCallBack, callBack YeeCallBack) string {

	privateKey, err := md5CallBack(parentMerchantNo)

	if err != nil {
		return yee.replyNotify(err)
	}

	yee.PrivateKey = privateKey

	err = yee.getNotifyData(ret)

	if err != nil {
		return yee.replyNotify(err)
	}

	err = callBack(yee.notifyData)

	if err != nil {
		return yee.replyNotify(err)
	}

	return yee.replyNotify(nil)
}

func (yee *YeeNotify) Base64Decode(data string) ([]byte, error) {

	data = strings.ReplaceAll(strings.ReplaceAll(data, "-", "+"), "_", "/")

	switch len(data) % 4 {
	case 2:
		data += "=="
	case 3:
		data += "="
	}

	return base64.StdEncoding.DecodeString(data)
}
