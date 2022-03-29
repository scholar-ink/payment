package notify

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"strings"
)

type ZgNotifyData struct {
	ResultCode     string `json:"result_code"`
	ErrCode        string `json:"err_code"`
	ErrCodeDes     string `json:"err_code_des"`
	PayStatus      string `json:"pay_status"`                 //支付状态码，此字段值最终作为支付是否成功的结果
	PayStatusInfo  string `json:"pay_status_info,omitempty"`  //支付结果描述
	PaymentType    string `json:"payment_type"`               //支付方式代码
	OutTradeNo     string `json:"out_trade_no"`               //商户订单号(唯一编号且长度小于等于32)
	OrderNo        string `json:"order_no"`                   //银通平台订单号
	FeeType        string `json:"fee_type"`                   //币种，默认传值：CNY
	TotalFee       string `json:"total_fee"`                  //订单金额，单位到分
	PaidFee        string `json:"paid_fee"`                   //支付金额，单位到分，支付成功时，需校验支付金额与订单金额是否一致
	CouponFee      string `json:"coupon_fee,omitempty"`       //现金券支付金额，单位到分
	Attach         string `json:"attach,omitempty"`           //商品附加信息
	TimeEnd        string `json:"time_end"`                   //支付完成时间
	Openid         string `json:"openid,omitempty"`           //用户标识,用户在服务商 appid 下的唯一标识
	IsSubscribe    string `json:"is_subscribe"`               //用户是否关注子公众账号，Y-关注，N-未关注
	BankType       string `json:"bank_type,omitempty"`        //付款银行类型
	TransactionId  string `json:"transaction_id,omitempty"`   //收单机构号(微信或支付宝交易号)
	BankBillNo     string `json:"bank_bill_no,omitempty"`     //付款银行订单号，若为微信支付则为空
	SubAppid       string `json:"sub_appid,omitempty"`        //商户公众号 appid
	SubIsSubscribe string `json:"sub_is_subscribe,omitempty"` //是否关注子商户公众账号，Y-关注，N-未关注
	SubOpenid      string `json:"sub_openid,omitempty"`       //用户在商户公众号 appid 下的唯一标识
	Ext            string `json:"ext"`                        //扩展字段，类型为map格式json
}

type ZgNotify struct {
	ServiceName string `json:"serviceName"`
	Version     string `json:"version"`
	MerchantNo  string `json:"merchantNo"`
	Md5Key      string `xml:"-" json:"-"`
	SourceData  string `json:"sourceData"`
	SignData    string `json:"signData"`
	RequestId   string `json:"requestId"`
	RequestTime string `json:"requestTime"`
	notifyData  *ZgNotifyData
}

func (zg *ZgNotify) getNotifyData(ret string) error {

	if ret == "" {
		return errors.New("获取通知数据失败")
	}

	err := json.Unmarshal([]byte(ret), zg)

	if err != nil {
		return errors.New("解析返回数据失败:" + err.Error())
	}

	notify := new(ZgNotifyData)

	err = json.Unmarshal([]byte(zg.SourceData), notify)

	if err != nil {
		return errors.New("解析通知数据失败:" + err.Error())
	}

	zg.notifyData = notify

	return nil
}

func (zg *ZgNotify) checkNotify() error {
	if zg.notifyData.ResultCode != "0" {
		return errors.New("中钢返回错误" + zg.notifyData.ErrCodeDes)
	}
	return zg.verifySign()
}

func (zg *ZgNotify) verifySign() error {

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

func (zg *ZgNotify) replyNotify(err error) string {

	if err != nil {
		fmt.Println(err)
		return "0002"
	} else {
		return "0000"
	}
}

type ZgCallBack func(data *ZgNotifyData) error

type ZgMd5KeyCallBack func(merchantNo string) (md5Key string, err error)

func (zg *ZgNotify) Handle(ret string, md5CallBack ZgMd5KeyCallBack, callBack ZgCallBack) string {

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
