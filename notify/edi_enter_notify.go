package notify

import (
	"encoding/json"
	"errors"
)

type EdiEnterNotifyData struct {
	AudMsg    string `json:"audMsg"`
	AudSts    string `json:"audSts"`
	BackType  string `json:"backType"`
	MchType   string `json:"mchType"`
	MercId    string `json:"mercId"`
	OutMercId string `json:"outMercId"`
	Status    string `json:"status"`
	UpMerId   string `json:"upMerId"`
}

type EdiEnterNotify struct {
	notifyData *EdiEnterNotifyData
}

func (yqb *EdiEnterNotify) getNotifyData(ret string) error {

	notify := new(EdiEnterNotifyData)

	err := json.Unmarshal([]byte(ret), notify)

	if err != nil {
		return errors.New("解析通知数据失败:" + err.Error())
	}

	yqb.notifyData = notify

	return nil
}

func (yqb *EdiEnterNotify) replyNotify(err error) string {
	if err != nil {
		return `{"result_code":"fail"}`
	} else {
		return `{"result_code":"success"}`
	}
}

type EdiEnterCallBack func(data *EdiEnterNotifyData) error

func (yqb *EdiEnterNotify) Handle(ret string, callBack EdiEnterCallBack) string {
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
