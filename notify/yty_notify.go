package notify

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/scholar-ink/payment/helper"
)

type YtyNotifyData struct {
	ReturnCode  string `json:"return_code"`
	ReturnMsg   string `json:"return_msg"`
	AgentNo     string `json:"agent_no"`
	MerchantNo  string `json:"merchant_no"`
	ProductNo   string `json:"product_no"`
	OrderNo     string `json:"order_no"`
	OutOrderNo  string `json:"out_order_no"`
	TotalAmount int32  `json:"total_amount"`
	PayType     string `json:"pay_type"`
	TradeStatus int32  `json:"trade_status"`
	TradeTime   string `json:"trade_time"`
}

type YtyNotify struct {
	Md5Key     string `xml:"-" json:"-"`
	notifyData *YtyNotifyData
}

func (yty *YtyNotify) getNotifyData(ret string) error {
	if ret == "" {
		return errors.New("获取通知数据失败")
	}

	var baseReturn struct {
		Encd string `json:"encd"`
		Sign string `json:"sign"`
		Enck string `json:"enck"`
	}

	err := json.Unmarshal([]byte(ret), &baseReturn)

	if err != nil {
		return errors.New("解析返回数据失败:" + err.Error())
	}

	enck, err := helper.Rsa1Decrypt2(baseReturn.Enck, yty.Md5Key)

	if err != nil {
		return err
	}

	b, _ := base64.StdEncoding.DecodeString(baseReturn.Encd)

	b, err = helper.TripleEcbDesDecrypt(b, enck)

	if err != nil {
		return errors.New("解密通知数据失败:" + err.Error())
	}

	notify := new(YtyNotifyData)

	err = json.Unmarshal(b, notify)

	if err != nil {
		return errors.New("解析通知数据失败:" + err.Error())
	}

	yty.notifyData = notify

	return nil
}

func (yty *YtyNotify) replyNotify(err error) string {

	if err != nil {
		fmt.Println(err)
		return "error"
	} else {
		return "SUCCESS"
	}
}

type YtyCallBack func(data *YtyNotifyData) error

type YtyMd5KeyCallBack func() (md5Key string, err error)

func (yty *YtyNotify) Handle(ret string, md5CallBack YtyMd5KeyCallBack, callBack YtyCallBack) string {

	md5Key, err := md5CallBack()

	if err != nil {
		return yty.replyNotify(err)
	}

	yty.Md5Key = md5Key

	err = yty.getNotifyData(ret)

	if err != nil {
		return yty.replyNotify(err)
	}

	err = callBack(yty.notifyData)

	if err != nil {
		return yty.replyNotify(err)
	}

	return yty.replyNotify(nil)
}
