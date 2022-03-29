package notify

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/scholar-ink/payment/helper"
)

type FqlNotifyData struct {
	Amount       string `json:"amount"`       //交易金额
	ChannelType  string `json:"channelType"`  //支付渠道
	CompleteTime string `json:"completeTime"` //完成时间
	ExAccNo      string `json:"exAccNo"`      //虚户账号
	FeeAmtAmount string `json:"feeAmtAmount"` //手续费金额
	FeeAmtRate   string `json:"feeAmtRate"`   //手续费费率
	OrderId      string `json:"orderId"`      //商户订单号
	PayExAccNo   string `json:"payExAccNo"`   //付款账号
	PmtType      string `json:"pmtType"`      //支付类型
	Status       string `json:"status"`       //交易状态
	TxnId        string `json:"txnId"`        //系统单号 02 交易成功
}

type FqlNotify struct {
	MerchId    string `json:"merchId"`
	SignType   string `json:"signType"`
	SignInfo   string `json:"signInfo"`
	Data       string `json:"data"`
	Md5Key     string `xml:"-" json:"-"`
	notifyData *FqlNotifyData
}

func (fql *FqlNotify) getNotifyData(ret string) error {
	if ret == "" {
		return errors.New("获取通知数据失败")
	}

	err := json.Unmarshal([]byte(ret), fql)

	if err != nil {
		return errors.New("解析返回数据失败:" + err.Error())
	}

	notify := new(FqlNotifyData)

	err = json.Unmarshal([]byte(fql.Data), notify)

	if err != nil {
		return errors.New("解析通知数据失败:" + err.Error())
	}

	fmt.Println(notify)

	fql.notifyData = notify

	return nil

	return nil
}

func (fql *FqlNotify) checkNotify() error {
	return fql.verifySign()
}

func (fql *FqlNotify) verifySign() error {

	err := helper.Sha1WithRsaVerify([]byte(fql.Data), fql.SignInfo, fql.Md5Key)

	fmt.Println(err)

	if err != nil {
		return errors.New("返回数据验签失败，可能数据被篡改")
	}
	return nil

}

func (fql *FqlNotify) replyNotify(err error) string {

	if err != nil {
		fmt.Println(err)
		return "error"
	} else {
		return "success"
	}
}

type FqlCallBack func(data *FqlNotifyData) error

type FqlMd5KeyCallBack func(merchantNo string) (md5Key string, err error)

func (fql *FqlNotify) Handle(ret string, md5CallBack FqlMd5KeyCallBack, callBack FqlCallBack) string {

	err := fql.getNotifyData(ret)

	if err != nil {
		return fql.replyNotify(err)
	}

	md5Key, err := md5CallBack(fql.MerchId)

	if err != nil {
		return fql.replyNotify(err)
	}

	fql.Md5Key = md5Key

	err = fql.checkNotify()

	if err != nil {
		return fql.replyNotify(err)
	}

	err = callBack(fql.notifyData)

	if err != nil {
		return fql.replyNotify(err)
	}

	return fql.replyNotify(nil)
}
