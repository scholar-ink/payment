package notify

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/scholar-ink/payment/helper"
)

type AdaPayNotifyData struct {
	OrderNo      string `json:"order_no"`
	Id           string `json:"id"`
	PartyOrderId string `json:"party_order_id"`
	Status       string `json:"status"`
}

type AdaPayResponseData struct {
	AppId       string            `json:"app_id"`
	CreatedTime string            `json:"created_time"`
	Data        string            `json:"data"`
	Id          string            `json:"id"`
	Object      string            `json:"object"`
	ProdMode    string            `json:"prod_mode"`
	Sign        string            `json:"sign"`
	TradeType   string            `json:"type"`
	NotifyData  *AdaPayNotifyData `json:"-"`
}

type AdaPayNotify struct {
	Md5Key       string `xml:"-" json:"-"`
	responseData *AdaPayResponseData
}

func (zg *AdaPayNotify) getNotifyData(ret map[string]interface{}) error {

	if len(ret) == 0 {
		return errors.New("获取通知数据失败")
	}

	notify := new(AdaPayResponseData)

	err := helper.DeepCopy(ret, notify)

	if err != nil {
		return errors.New("解析返回数据失败:" + err.Error())
	}

	json.Unmarshal([]byte(notify.Data), &notify.NotifyData)

	zg.responseData = notify

	return nil
}

func (zg *AdaPayNotify) checkNotify() error {
	return zg.verifySign()
}

func (zg *AdaPayNotify) verifySign() error {

	err := helper.Sha1WithRsaVerify([]byte(zg.responseData.Data), zg.responseData.Sign, zg.Md5Key)

	if err != nil {
		return errors.New("返回数据验签失败，可能数据被篡改")
	}

	return nil
}

func (zg *AdaPayNotify) replyNotify(err error) string {

	if err != nil {
		fmt.Println(err)
		return "error"
	} else {
		return "RECV_ORD_ID_" + zg.responseData.NotifyData.Id[len(zg.responseData.NotifyData.Id)-32:]
	}
}

type AdaPayCallBack func(data *AdaPayNotifyData) error

type AdaPayMd5KeyCallBack func(merchantNo string) (md5Key string, err error)

func (zg *AdaPayNotify) Handle(ret map[string]interface{}, md5CallBack AdaPayMd5KeyCallBack, callBack AdaPayCallBack) string {

	err := zg.getNotifyData(ret)

	if err != nil {
		return zg.replyNotify(err)
	}

	md5Key, err := md5CallBack(zg.responseData.AppId)

	if err != nil {
		return zg.replyNotify(err)
	}

	zg.Md5Key = md5Key

	err = zg.checkNotify()

	if err != nil {
		return zg.replyNotify(err)
	}

	err = callBack(zg.responseData.NotifyData)

	if err != nil {
		return zg.replyNotify(err)
	}

	return zg.replyNotify(nil)
}
