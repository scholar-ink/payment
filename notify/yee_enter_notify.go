package notify

import (
	"github.com/scholar-ink/payment/helper"

	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

type YeeEnterNotifyData struct {
	MerNo          string `json:"merNo"`
	AgentNo        string `json:"agentNo"`
	ExternalId     string `json:"externalId"`
	Remark         string `json:"remark"`
	MerNetInStatus string `json:"merNetInStatus"`
}
type YeeEnterConfig struct {
	ParentMerchantNo string `json:"parentMerchantNo"`
	PrivateKey       string `json:"-"`
}

type YeeEnterNotify struct {
	*YeeEnterConfig
	notifyData *YeeEnterNotifyData
}

func (yee *YeeEnterNotify) InitBaseConfig(config *YeeEnterConfig) {
	yee.YeeEnterConfig = config
}

func (yee *YeeEnterNotify) getNotifyData(ret string) error {

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

	notify := new(YeeEnterNotifyData)

	err = json.Unmarshal([]byte(originStr), notify)

	if err != nil {
		return errors.New("解析通知数据失败:" + err.Error())
	}

	yee.notifyData = notify

	return nil
}

func (yee *YeeEnterNotify) verifySign() error {
	return nil
}

func (yee *YeeEnterNotify) replyNotify(err error) string {

	if err != nil {
		return err.Error()
	} else {
		return "SUCCESS"
	}
}

type YeeEnterCallBack func(data *YeeEnterNotifyData) error

func (yee *YeeEnterNotify) Handle(ret string, callBack YeeEnterCallBack) string {
	err := yee.getNotifyData(ret)

	if err != nil {
		return yee.replyNotify(err)
	}

	err = callBack(yee.notifyData)

	if err != nil {
		return yee.replyNotify(err)
	}

	return yee.replyNotify(nil)
}

func (yee *YeeEnterNotify) Base64Decode(data string) ([]byte, error) {

	data = strings.ReplaceAll(strings.ReplaceAll(data, "-", "+"), "_", "/")

	switch len(data) % 4 {
	case 2:
		data += "=="
	case 3:
		data += "="
	}

	return base64.StdEncoding.DecodeString(data)
}
