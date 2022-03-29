package notify

import (
	"errors"
	"fmt"
	"github.com/scholar-ink/payment/helper"
)

type HangNotifyData struct {
	Amount        string `json:"amount"`
	ChannelNo     string `json:"channelNo"`
	Code          string `json:"code"`
	MerchantNo    string `json:"merchantNo"`
	Message       string `json:"message"`
	OrderNo       string `json:"orderNo"`
	SubMerchantNo string `json:"subMerchantNo"`
	Sign          string `json:"sign"`
}

type HangNotify struct {
	Md5Key     string `xml:"-" json:"-"`
	notifyData *HangNotifyData
}

func (zg *HangNotify) getNotifyData(ret map[string]interface{}) error {

	if len(ret) == 0 {
		return errors.New("获取通知数据失败")
	}

	notify := new(HangNotifyData)

	err := helper.DeepCopy(ret, notify)

	if err != nil {
		return errors.New("解析返回数据失败:" + err.Error())
	}

	zg.notifyData = notify

	return nil
}

func (zg *HangNotify) checkNotify() error {
	return zg.verifySign()
}

func (zg *HangNotify) verifySign() error {

	sign := fmt.Sprintf("code=%s&message=%s&subMerchantNo=%s&orderNo=%s&merchantNo=%s&channelNo=%s&amount=%s",
		zg.notifyData.Code, zg.notifyData.Message, zg.notifyData.SubMerchantNo, zg.notifyData.OrderNo, zg.notifyData.MerchantNo, zg.notifyData.ChannelNo, zg.notifyData.Amount)

	return helper.Sha1WithRsaVerify([]byte(sign), zg.notifyData.Sign, zg.Md5Key)
}

func (zg *HangNotify) replyNotify(err error) string {

	if err != nil {
		fmt.Println(err)
		return "error"
	} else {
		return "success"
	}
}

type HangCallBack func(data *HangNotifyData) error

type HangMd5KeyCallBack func(merchantNo string) (md5Key string, err error)

func (zg *HangNotify) Handle(ret map[string]interface{}, md5CallBack HangMd5KeyCallBack, callBack HangCallBack) string {

	err := zg.getNotifyData(ret)

	if err != nil {
		return zg.replyNotify(err)
	}

	md5Key, err := md5CallBack(zg.notifyData.SubMerchantNo)

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
