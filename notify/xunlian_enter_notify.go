package notify

import (
	"encoding/json"
	"errors"
)

type XunLianEnterNotifyResponse struct {
	Timestamp int64  `json:"timestamp"`
	ReqId     string `json:"reqId"`
	InsCode   string `json:"insCode"`
	DataType  int32  `json:"type"`     // 0:受理通知，1:审核通知
	BussType  int32  `json:"bussType"` // 4:新增微信,5:新增支付宝
	MerCode   string `json:"merCode"`
}

type XunLianEnterNotifyData struct {
	Message string                      `json:"message"`
	Status  int32                       `json:"status"`
	Data    *XunLianEnterNotifyResponse `json:"data"`
}

type XunLianEnterNotify struct {
	notifyData *XunLianEnterNotifyData
}

func (xunLian *XunLianEnterNotify) getNotifyData(ret string) error {

	if ret == "" {
		return errors.New("获取通知数据失败")
	}

	notify := new(XunLianEnterNotifyData)

	err := json.Unmarshal([]byte(ret), notify)

	if err != nil {
		return errors.New("解析通知数据失败:" + err.Error())
	}

	xunLian.notifyData = notify

	return nil
}

func (xunLian *XunLianEnterNotify) replyNotify(err error) string {

	if err != nil {
		return "false"
	} else {
		return "true "
	}
}

type XunLianEnterCallBack func(data *XunLianEnterNotifyData) error

func (xunLian *XunLianEnterNotify) Handle(ret string, callBack XunLianEnterCallBack) string {
	err := xunLian.getNotifyData(ret)

	if err != nil {
		return xunLian.replyNotify(err)
	}

	err = callBack(xunLian.notifyData)

	if err != nil {
		return xunLian.replyNotify(err)
	}

	return xunLian.replyNotify(nil)
}
