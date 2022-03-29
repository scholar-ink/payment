package notify

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"strings"
)

type QfNotifyData struct {
	Status     string `json:"status"`
	PayType    string `json:"pay_type"`
	SysDtm     string `json:"sysdtm"`
	PayDtm     string `json:"paydtm"` //用户支付时间
	TxDtm      string `json:"txdtm"`
	TxAmt      string `json:"txamt"` //订单支付金额，单位分；
	MchId      string `json:"mchid"`
	OutTradeNo string `json:"out_trade_no"` //外部订单号，开发者定义订单号；
	SysSn      string `json:"syssn"`        //交易流水号
	RespCd     string `json:"respcd"`
}

type QfNotify struct {
	Md5Key     string `json:"-"`
	Sign       string `json:"-"`
	notifyData *QfNotifyData
}

func (Qf *QfNotify) getNotifyData(ret string) error {

	if ret == "" {
		return errors.New("获取通知数据失败")
	}

	notify := new(QfNotifyData)

	err := json.Unmarshal([]byte(ret), notify)

	if err != nil {
		return errors.New("解析返回数据失败:" + err.Error())
	}

	Qf.notifyData = notify

	return nil
}

func (Qf *QfNotify) checkNotify(ret string) error {
	if Qf.notifyData.RespCd != "0000" {
		return errors.New("钱方返回错误")
	}
	return Qf.verifySign(ret)
}

func (Qf *QfNotify) verifySign(signStr string) error {

	signStr += Qf.Md5Key

	fmt.Println(signStr)

	signStr = strings.ToUpper(helper.Md5(signStr))

	fmt.Println(signStr)

	fmt.Println(Qf.Sign)

	if Qf.Sign != strings.ToUpper(signStr) {
		return errors.New("返回数据验签失败，可能数据被篡改")
	}

	return nil
}

func (Qf *QfNotify) replyNotify(err error) string {

	if err != nil {
		fmt.Println(err)
		return "ERROR"
	} else {
		return "SUCCESS"
	}
}

type QfCallBack func(data *QfNotifyData) error

type QfMd5KeyCallBack func(MchId string) (md5Key string, err error)

func (Qf *QfNotify) Handle(ret, sign string, md5CallBack QfMd5KeyCallBack, callBack QfCallBack) string {

	Qf.Sign = sign

	err := Qf.getNotifyData(ret)

	if err != nil {
		return Qf.replyNotify(err)
	}

	md5Key, err := md5CallBack(Qf.notifyData.MchId)

	if err != nil {
		return Qf.replyNotify(err)
	}

	Qf.Md5Key = md5Key

	err = Qf.checkNotify(ret)

	if err != nil {
		return Qf.replyNotify(err)
	}

	err = callBack(Qf.notifyData)

	if err != nil {
		return Qf.replyNotify(err)
	}

	return Qf.replyNotify(nil)
}
