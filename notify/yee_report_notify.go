package notify

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/scholar-ink/payment/helper"
	"strings"
)

type YeeReportInfo struct {
	TraceId    string              `json:"traceId"`
	BizMsg     string              `json:"bizMsg"`
	BizCode    string              `json:"bizCode"`
	DealStatus int32               `json:"dealStatus"`
	Products   []*YeeReportProduct `json:"openPayReportDTOList"` //扫码产品
}

type YeeReportNotifyData struct {
	MerNo         string         `json:"merNo"`
	AgentNo       string         `json:"agentNo"`
	ExternalId    string         `json:"externalId"`
	ReportInfoStr string         `json:"reportInfo"`
	ReportInfo    *YeeReportInfo `json:"-"`
}

type YeeReportProduct struct {
	BankCode         string `json:"bankCode"`         //银行编码
	BankType         string `json:"bankType"`         //通道类型	ALIPAY：支付宝 WECHAT：微信
	SceneType        string `json:"sceneType"`        //场景类型ACTIVE：主扫 PASSIVE：被扫 H5：H5支付 JSAPI：公众号
	ChannelNo        string `json:"channelNo"`        //通道渠道号
	MerchantNo       string `json:"merchantNo"`       //子商编
	ReportMerchantNo string `json:"reportMerchantNo"` //通道侧商编
	DealStatus       string `json:"dealStatus"`       //SUCCESS：报备成功 ，FAIL：失败
	ErrMsg           string `json:"errMsg"`           //错误信息
}

type YeeReportConfig struct {
	ParentMerchantNo string `json:"parentMerchantNo"`
	PrivateKey       string `json:"-"`
}

type YeeReportNotify struct {
	*YeeReportConfig
	notifyData *YeeReportNotifyData
}

func (yee *YeeReportNotify) InitBaseConfig(config *YeeReportConfig) {
	yee.YeeReportConfig = config
}

func (yee *YeeReportNotify) getNotifyData(ret string) error {

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

	notify := new(YeeReportNotifyData)

	err = json.Unmarshal([]byte(originStr), notify)

	if err != nil {
		return errors.New("解析通知数据失败:" + err.Error())
	}

	reportInfo := new(YeeReportInfo)

	err = json.Unmarshal([]byte(notify.ReportInfoStr), reportInfo)

	if err != nil {
		return errors.New("解析reportInfo数据失败:" + err.Error())
	}

	notify.ReportInfo = reportInfo

	yee.notifyData = notify

	return nil
}

func (yee *YeeReportNotify) verifySign() error {
	return nil
}

func (yee *YeeReportNotify) replyNotify(err error) string {

	if err != nil {
		return err.Error()
	} else {
		return "SUCCESS"
	}
}

type YeeReportCallBack func(data *YeeReportNotifyData) error

func (yee *YeeReportNotify) Handle(ret string, callBack YeeReportCallBack) string {
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

func (yee *YeeReportNotify) Base64Decode(data string) ([]byte, error) {

	data = strings.ReplaceAll(strings.ReplaceAll(data, "-", "+"), "_", "/")

	switch len(data) % 4 {
	case 2:
		data += "=="
	case 3:
		data += "="
	}

	return base64.StdEncoding.DecodeString(data)
}
