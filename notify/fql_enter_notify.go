package notify

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/scholar-ink/payment/helper"
)

type FqlEnterNotifyData struct {
	BizType      string `json:"bizType"`      //01-企业开户；02-个人开户；03-企业进件；04-个人进件
	ChUserId     string `json:"chUserId"`     //商户用户号
	CompleteTime string `json:"completeTime"` //完成时间
	CustomerId   string `json:"customerId"`   //客户号
	OrglOrderId  string `json:"orglOrderId"`  //原商户订单号
	Status       string `json:"status"`       //交易成功
}

type FqlConfig struct {
	PlatPubKey string `json:"-"`
}

type FqlEnterNotify struct {
	*FqlConfig
	MerchId    string `json:"merchId"`
	SignType   string `json:"signType"`
	SignInfo   string `json:"signInfo"`
	Data       string `json:"data"`
	notifyData *FqlEnterNotifyData
}

func (fql *FqlEnterNotify) InitBaseConfig(config *FqlConfig) {
	fql.FqlConfig = config
}

func (fql *FqlEnterNotify) getNotifyData(ret string) error {

	if ret == "" {
		return errors.New("获取通知数据失败")
	}

	err := json.Unmarshal([]byte(ret), fql)

	if err != nil {
		return errors.New("解析返回数据失败:" + err.Error())
	}

	notify := new(FqlEnterNotifyData)

	err = json.Unmarshal([]byte(fql.Data), notify)

	if err != nil {
		return errors.New("解析通知数据失败:" + err.Error())
	}

	fql.notifyData = notify

	err = fql.verifySign()

	if err != nil {
		return err
	}

	return nil
}

func (fql *FqlEnterNotify) verifySign() error {

	fmt.Println(fql.Data)

	fmt.Println(fql.SignInfo)

	fmt.Println(fql.PlatPubKey)

	err := helper.Sha1WithRsaVerify([]byte(fql.Data), fql.SignInfo, fql.PlatPubKey)

	if err != nil {
		return errors.New("返回数据验签失败，可能数据被篡改")
	}
	return nil
}

func (fql *FqlEnterNotify) replyNotify(err error) string {

	if err != nil {
		return err.Error()
	} else {
		return "SUCCESS "
	}
}

type FqlEnterCallBack func(data *FqlEnterNotifyData) error

func (fql *FqlEnterNotify) Handle(ret string, callBack FqlEnterCallBack) string {
	err := fql.getNotifyData(ret)

	if err != nil {
		return fql.replyNotify(err)
	}

	err = callBack(fql.notifyData)

	if err != nil {
		return fql.replyNotify(err)
	}

	return fql.replyNotify(nil)
}
