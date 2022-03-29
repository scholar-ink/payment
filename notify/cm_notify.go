package notify

import (
	"errors"
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"strings"
)

type CmNotifyData struct {
	ChannelTradeNo string `json:"channel_trade_no"` //微信/支付宝/三方订单号
	Money          string `json:"money"`            //金额
	OrderId        string `json:"orderId"`          //订单号
	OutTradeNo     string `json:"out_trade_no"`     //聪明付订单号
	ResultCode     string `json:"result_code"`      //支付状态
	ShopId         string `json:"shopId"`           //商户id
	Sign           string `json:"sign"`             //sign
	TimeEnd        string `json:"timeEnd"`          //支付成功时间
	TransactionId  string `json:"transaction_id"`   //上游订单号
	PayType        string `json:"type"`             //支付类型
}

type CmNotify struct {
	Md5Key     string `xml:"-" json:"-"`
	notifyData *CmNotifyData
}

func (cm *CmNotify) getNotifyData(ret map[string]string) error {

	if len(ret) == 0 {
		return errors.New("获取通知数据失败")
	}

	notify := new(CmNotifyData)

	err := helper.DeepCopy(ret, notify)

	if err != nil {
		return errors.New("解析返回数据失败:" + err.Error())
	}

	cm.notifyData = notify

	return nil
}

func (cm *CmNotify) checkNotify() error {

	return cm.verifySign()
}

func (cm *CmNotify) verifySign() error {

	signStr := fmt.Sprintf("money=%s&orderId=%s&result_code=%s&shopId=%s", cm.notifyData.Money, cm.notifyData.OrderId, cm.notifyData.ResultCode, cm.notifyData.ShopId)

	signStr = strings.ToUpper(helper.Md5(signStr + "&key=" + cm.Md5Key))

	if cm.notifyData.Sign != signStr {
		return errors.New("返回数据验签失败，可能数据被篡改")
	}

	return nil
}

func (cm *CmNotify) replyNotify(err error) string {

	if err != nil {
		fmt.Println(err)
		return "error"
	} else {
		return "success"
	}
}

type CmCallBack func(data *CmNotifyData) error

type CmMd5KeyCallBack func(merchantNo string) (md5Key string, err error)

func (cm *CmNotify) Handle(ret map[string]string, md5CallBack CmMd5KeyCallBack, callBack CmCallBack) string {

	err := cm.getNotifyData(ret)

	if err != nil {
		return cm.replyNotify(err)
	}

	md5Key, err := md5CallBack(cm.notifyData.ShopId)

	if err != nil {
		return cm.replyNotify(err)
	}

	cm.Md5Key = md5Key

	err = cm.checkNotify()

	if err != nil {
		return cm.replyNotify(err)
	}

	err = callBack(cm.notifyData)

	if err != nil {
		return cm.replyNotify(err)
	}

	return cm.replyNotify(nil)
}
