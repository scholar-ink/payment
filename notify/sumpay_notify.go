package notify

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/scholar-ink/payment/helper"
)

type SumPayNotifyData struct {
	RespCode      string `json:"resp_code"`      //返回码
	RespMsg       string `json:"resp_msg"`       //响应信息
	MerchantNo    string `json:"mer_no"`         //商户编号
	OrderNo       string `json:"order_no"`       //订单号
	OrderTime     string `json:"order_time"`     //订单时间
	TradeNo       string `json:"trade_no"`       //商盟交易流水号
	Status        string `json:"status"`         //交易状态
	SuccessTime   string `json:"success_time"`   //成功时间
	SuccessAmount string `json:"success_amount"` //成功金额
	Version       string `json:"version"`        //版本号
	Sign          string `json:"sign"`           //签名
	SignType      string `json:"sign_type"`      //签名类型
}

type SumPayNotify struct {
	Md5Key     string `xml:"-" json:"-"`
	notifyData *SumPayNotifyData
}

func (zg *SumPayNotify) getNotifyData(ret map[string]interface{}) error {

	if len(ret) == 0 {
		return errors.New("获取通知数据失败")
	}

	notify := new(SumPayNotifyData)

	err := helper.DeepCopy(ret, notify)

	if err != nil {
		return errors.New("解析返回数据失败:" + err.Error())
	}

	zg.notifyData = notify

	return nil
}

func (zg *SumPayNotify) checkNotify(ret map[string]interface{}) error {
	return zg.verifySign(ret)
}

func (zg *SumPayNotify) verifySign(ret map[string]interface{}) error {

	signStr := helper.CreateLinkString(&ret)

	err := helper.Sha256WithRsaVerify([]byte(signStr), zg.notifyData.Sign, zg.Md5Key)

	if err != nil {
		return errors.New("返回数据验签失败，可能数据被篡改")
	}

	return nil
}

func (zg *SumPayNotify) replyNotify(err error) string {

	if err != nil {
		fmt.Println(err)
		return "error"
	} else {
		return "success"
	}
}

type SumPayCallBack func(data *SumPayNotifyData) error

type SumPayMd5KeyCallBack func(merchantNo string) (md5Key string, err error)

func (zg *SumPayNotify) Handle(retJson string, md5CallBack SumPayMd5KeyCallBack, callBack SumPayCallBack) string {

	ret := make(map[string]interface{}, 0)

	json.Unmarshal([]byte(retJson), &ret)

	err := zg.getNotifyData(ret)

	if err != nil {
		return zg.replyNotify(err)
	}

	md5Key, err := md5CallBack(zg.notifyData.MerchantNo)

	if err != nil {
		return zg.replyNotify(err)
	}

	zg.Md5Key = md5Key

	err = zg.checkNotify(ret)

	if err != nil {
		return zg.replyNotify(err)
	}

	err = callBack(zg.notifyData)

	if err != nil {
		return zg.replyNotify(err)
	}

	return zg.replyNotify(nil)
}
