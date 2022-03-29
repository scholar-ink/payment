package notify

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"strings"
)

type BillPayment struct {
	TargetOrderId string `json:"targetOrderId"`
}

type ChinaUmsNotifyData struct {
	BillNo          string       `json:"billNo"`
	BillStatus      string       `json:"billStatus"`
	BillPaymentInfo string       `json:"billPayment"`
	BillPayment     *BillPayment `json:"-"`
	CreateTime      string       `json:"createTime"`
	Mid             string       `json:"mid"`
	SeqId           string       `json:"seqId"`
	TotalAmount     string       `json:"totalAmount"`
	Sign            string       `json:"sign"`
}

type ChinaUmsNotify struct {
	Md5Key     string `xml:"-" json:"-"`
	notifyData *ChinaUmsNotifyData
}

func (cu *ChinaUmsNotify) getNotifyData(ret map[string]interface{}) error {

	if len(ret) == 0 {
		return errors.New("获取通知数据失败")
	}

	notify := new(ChinaUmsNotifyData)

	err := helper.DeepCopy(ret, notify)

	if err != nil {
		return errors.New("解析返回数据失败:" + err.Error())
	}

	billPayment := new(BillPayment)

	err = json.Unmarshal([]byte(notify.BillPaymentInfo), billPayment)

	if err != nil {
		return errors.New("解析返回数据失败:" + err.Error())
	}

	notify.BillPayment = billPayment

	cu.notifyData = notify

	return nil
}

func (cu *ChinaUmsNotify) verifySign(ret map[string]interface{}) error {

	helper.KSort(&ret)

	signStr := helper.CreateLinkString(&ret)

	signStr += cu.Md5Key

	signStr = strings.ToUpper(helper.Md5(signStr))

	if signStr != cu.notifyData.Sign {
		return errors.New("返回数据验签失败，可能数据被篡改")
	}

	return nil
}

func (cu *ChinaUmsNotify) replyNotify(err error) string {

	if err != nil {
		fmt.Println(err)
		return "FAILED"
	} else {
		return "SUCCESS"
	}
}

type ChinaUmsCallBack func(data *ChinaUmsNotifyData) error

type ChinaUmsMd5KeyCallBack func(merchantNo string) (md5Key string, err error)

func (cu *ChinaUmsNotify) Handle(ret map[string]interface{}, md5CallBack ChinaUmsMd5KeyCallBack, callBack ChinaUmsCallBack) string {

	err := cu.getNotifyData(ret)

	if err != nil {
		return cu.replyNotify(err)
	}

	md5Key, err := md5CallBack(cu.notifyData.Mid)

	if err != nil {
		return cu.replyNotify(err)
	}

	cu.Md5Key = md5Key

	err = cu.verifySign(ret)

	if err != nil {
		return cu.replyNotify(err)
	}

	err = callBack(cu.notifyData)

	if err != nil {
		return cu.replyNotify(err)
	}

	return cu.replyNotify(nil)
}
