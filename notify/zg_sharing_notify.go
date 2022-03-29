package notify

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"strings"
)

type ZgSharingNotifyData struct {
	DivideStatus     string                `json:"divideStatus"`
	DivideTime       string                `json:"divideTime"`
	PayOrderNo       int64                 `json:"payOrderNo"`
	DivideAmount     string                `json:"divideAmount"`
	OutTradeNo       string                `json:"outTradeNo"`
	DivideDetailData string                `json:"divideDetailList"`   //分账详情
	DivideDetailList []*ZgDivideDetailList `json:"divide_detail_list"` //分账详情
}

type ZgDivideDetailList struct {
	DivideAmount    string `json:"divide_amount"`     //分账金额
	DivideAccountNo string `json:"divide_account_no"` //分账账号
}

type ZgSharingNotify struct {
	ServiceName string `json:"serviceName"`
	Version     string `json:"version"`
	MerchantNo  int32  `json:"merchantNo"`
	Md5Key      string `xml:"-" json:"-"`
	SourceData  string `json:"sourceData"`
	SignData    string `json:"signData"`
	RequestId   string `json:"requestId"`
	RequestTime string `json:"requestTime"`
	notifyData  *ZgSharingNotifyData
}

func (zg *ZgSharingNotify) getNotifyData(ret string) error {

	if ret == "" {
		return errors.New("获取通知数据失败")
	}

	err := json.Unmarshal([]byte(ret), zg)

	if err != nil {
		return errors.New("解析返回数据失败:" + err.Error())
	}

	notify := new(ZgSharingNotifyData)

	err = json.Unmarshal([]byte(zg.SourceData), notify)

	if err != nil {
		return errors.New("解析通知数据失败:" + err.Error())
	}

	divideDetailList := make([]*ZgDivideDetailList, 0, 0)

	err = json.Unmarshal([]byte(notify.DivideDetailData), &divideDetailList)

	if err != nil {
		return errors.New("解析通知数据失败:" + err.Error())
	}

	notify.DivideDetailList = divideDetailList

	zg.notifyData = notify

	return nil
}

func (zg *ZgSharingNotify) checkNotify() error {
	return zg.verifySign()
}

func (zg *ZgSharingNotify) verifySign() error {

	mapData := helper.Struct2Map(zg)

	delete(mapData, "signData")

	signStr := helper.CreateLinkString(&mapData)

	signStr += "&key=" + zg.Md5Key

	signStr = helper.Md5(signStr)

	if zg.SignData != strings.ToUpper(signStr) {
		return errors.New("返回数据验签失败，可能数据被篡改")
	}

	return nil
}

func (zg *ZgSharingNotify) replyNotify(err error) string {
	if err != nil {
		fmt.Println(err)
		return "0002"
	} else {
		return "0000"
	}
}

type ZgSharingCallBack func(data *ZgSharingNotifyData) error

type ZgSharingMd5KeyCallBack func(merchantNo int32) (md5Key string, err error)

func (zg *ZgSharingNotify) Handle(ret string, md5CallBack ZgSharingMd5KeyCallBack, callBack ZgSharingCallBack) string {

	err := zg.getNotifyData(ret)

	if err != nil {
		return zg.replyNotify(err)
	}

	md5Key, err := md5CallBack(zg.MerchantNo)

	if err != nil {
		return zg.replyNotify(err)
	}

	zg.Md5Key = md5Key

	err = zg.checkNotify()

	if err != nil {
		return zg.replyNotify(err)
	}

	err = callBack(zg.notifyData)

	if err != nil {
		return zg.replyNotify(err)
	}

	return zg.replyNotify(nil)
}
