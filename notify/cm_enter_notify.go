package notify

import (
	"encoding/json"
	"errors"
)

type CmEnterNotifyData struct {
	ResultCode    string `json:"result_code"`
	ErrorMsg      string `json:"error_msg"`
	OrderId       string `json:"order_id"`
	ShopId        string `json:"shop_id"`
	PayType       string `json:"pay_type"`
	SubMerchantNo string `json:"sub_merchant_no"`
}

type CmEnterNotify struct {
	notifyData *CmEnterNotifyData
}

func (yqb *CmEnterNotify) getNotifyData(ret string) error {

	notify := new(CmEnterNotifyData)

	err := json.Unmarshal([]byte(ret), notify)

	if err != nil {
		return errors.New("解析通知数据失败:" + err.Error())
	}

	yqb.notifyData = notify

	return nil
}

func (yqb *CmEnterNotify) replyNotify(err error) string {
	if err != nil {
		return `{"result_code":"fail"}`
	} else {
		return `{"result_code":"success"}`
	}
}

type CmEnterCallBack func(data *CmEnterNotifyData) error

func (yqb *CmEnterNotify) Handle(ret string, callBack CmEnterCallBack) string {
	err := yqb.getNotifyData(ret)

	if err != nil {
		return yqb.replyNotify(err)
	}

	err = callBack(yqb.notifyData)

	if err != nil {
		return yqb.replyNotify(err)
	}

	return yqb.replyNotify(nil)
}
