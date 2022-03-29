package notify

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/scholar-ink/payment/helper"
)

type JttpNotifyData struct {
	OutTradeNo  string `json:"outTradeNo"`  //商户订单号
	TradeNo     string `json:"tradeNo"`     //平台订单号
	TotalAmount string `json:"totalAmount"` //商户订单金额
	PayAmount   string `json:"payAmount"`   //实际支付金额
	PayTime     string `json:"payTime"`     //交易支付时间
	RetCode     string `json:"retCode"`     //返回码
	RetMsg      string `json:"retMsg"`      //返回码描述
	Sign        string `json:"sign"`        //签名
}

type JttpResponseData struct {
	NodeId     string `json:"nodeId"`     //节点号
	OrgId      string `json:"orgId"`      //商户编号
	BizContext string `json:"bizContext"` //支付通道
	Charset    string `json:"charset"`    //参数编码字符集
	OrderTime  string `json:"orderTime"`  //商户订单提交时间
	Reserve1   string `json:"reserve1"`   //额外数据
	Sign       string `json:"sign"`
	SignType   string `json:"signType"` //签名方式
	TxnType    string `json:"txnType"`  //交易类型
}

type JttpNotify struct {
	Md5Key       string `xml:"-" json:"-"`
	AesKey       string `xml:"-" json:"-"`
	notifyData   *JttpNotifyData
	responseData *JttpResponseData
}

func (zg *JttpNotify) getResponseData(ret map[string]interface{}) error {

	if len(ret) == 0 {
		return errors.New("获取通知数据失败")
	}

	response := new(JttpResponseData)

	err := helper.DeepCopy(ret, response)

	if err != nil {
		return errors.New("解析返回数据失败:" + err.Error())
	}

	zg.responseData = response

	return nil
}

func (zg *JttpNotify) getNotifyData(ret []byte) error {

	if len(ret) == 0 {
		return errors.New("获取通知数据失败")
	}

	notify := new(JttpNotifyData)

	err := json.Unmarshal(ret, notify)

	if err != nil {
		return errors.New("解析返回数据失败:" + err.Error())
	}

	zg.notifyData = notify

	return nil
}

func (zg *JttpNotify) checkNotify(ret map[string]interface{}) error {
	return zg.verifySign(ret)
}

func (zg *JttpNotify) verifySign(ret map[string]interface{}) error {

	b, _ := base64.StdEncoding.DecodeString(zg.responseData.BizContext)

	aesKey, _ := base64.StdEncoding.DecodeString(zg.AesKey)

	origin, err := helper.AesDecrypt(b, aesKey) // ECB解密

	if err != nil {
		return errors.New("返回数据解密失败：" + err.Error())
	}

	origin = helper.PKCS7UnPadding(origin)

	err = helper.Sha256WithRsaVerify(origin, zg.responseData.Sign, zg.Md5Key)

	if err != nil {
		return errors.New("返回数据验签失败:" + err.Error())
	}

	return zg.getNotifyData(origin)
}

func (zg *JttpNotify) replyNotify(err error) string {

	if err != nil {
		fmt.Println(err)
		return "error"
	} else {
		return "success"
	}
}

type JttpCallBack func(data *JttpNotifyData) error

type JttpMd5KeyCallBack func(merchantNo string) (md5Key, aesKey string, err error)

func (zg *JttpNotify) Handle(ret map[string]interface{}, md5CallBack JttpMd5KeyCallBack, callBack JttpCallBack) string {

	err := zg.getResponseData(ret)

	if err != nil {
		return zg.replyNotify(err)
	}

	md5Key, aesKey, err := md5CallBack(zg.responseData.OrgId)

	if err != nil {
		return zg.replyNotify(err)
	}

	zg.Md5Key = md5Key

	zg.AesKey = aesKey

	err = zg.checkNotify(ret)

	if err != nil {
		return zg.replyNotify(err)
	}

	err = callBack(zg.notifyData)

	if err != nil {
		return zg.replyNotify(err)
	}

	return zg.replyNotify(nil)
}
