package notify

import (
	"errors"
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"strings"
)

type XinHuiNotifyData struct {
	AgentMerNo    string `json:"agent_mer_no"`    //服务商编号
	MerchantNo    string `json:"merchant_no"`     //商户编号
	OutTradeNo    string `json:"out_trade_no"`    //商户编号
	TrxExternalId string `json:"trx_external_id"` //信汇平台交易流水号
	PayExternalId string `json:"pay_external_id"` //平台支付流水号
	OrgExternalId string `json:"org_external_id"` //信汇上游渠道交易流水号
	TotalFee      string `json:"total_fee"`       //总金额
	ChannelType   string `json:"channel_type"`    //渠道类型
	RspCode       string `json:"rsp_code"`        //请求响应码
	TrxStatus     string `json:"trx_status"`      //交易状态
	RspMsg        string `json:"rsp_msg"`         //响应信息
	TimeEnd       string `json:"time_end"`        //响应信息
	Version       string `json:"version"`         //版本号
	Sign          string `json:"sign"`            //签名
}

type XinHuiNotify struct {
	Md5Key     string `xml:"-" json:"-"`
	notifyData *XinHuiNotifyData
}

func (zg *XinHuiNotify) getNotifyData(ret map[string]interface{}) error {

	if len(ret) == 0 {
		return errors.New("获取通知数据失败")
	}

	notify := new(XinHuiNotifyData)

	err := helper.DeepCopy(ret, notify)

	if err != nil {
		return errors.New("解析返回数据失败:" + err.Error())
	}

	zg.notifyData = notify

	return nil
}

func (zg *XinHuiNotify) checkNotify(ret map[string]interface{}) error {
	return zg.verifySign(ret)
}

func (zg *XinHuiNotify) verifySign(ret map[string]interface{}) error {

	signStr := helper.CreateLinkString(&ret)

	zg.notifyData.Sign = strings.ReplaceAll(zg.notifyData.Sign, " ", "+")

	err := helper.Sha256WithRsaVerify([]byte(signStr), zg.notifyData.Sign, zg.Md5Key)

	if err != nil {
		return errors.New("返回数据验签失败，可能数据被篡改")
	}

	return nil
}

func (zg *XinHuiNotify) replyNotify(err error) string {

	if err != nil {
		fmt.Println(err)
		return "error"
	} else {
		return "success"
	}
}

type XinHuiCallBack func(data *XinHuiNotifyData) error

type XinHuiMd5KeyCallBack func(merchantNo string) (md5Key string, err error)

func (zg *XinHuiNotify) Handle(ret map[string]interface{}, md5CallBack XinHuiMd5KeyCallBack, callBack XinHuiCallBack) string {

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
