package notify

import (
	"errors"
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/map"
	"net/url"
	"strings"
)

type XunLianNotifyData struct {
	ChannelOrderNum string `json:"channelOrderNum"` //渠道交易号
	Chcd            string `json:"chcd"`            //支付渠道
	ChcdDiscount    string `json:"chcdDiscount"`    //渠道优惠
	ConsumerAccount string `json:"consumerAccount"` //用户账号
	ConsumerId      string `json:"consumerId"`      //渠道账号 ID
	ErrorDetail     string `json:"errorDetail"`     //错误信息
	Inscd           string `json:"inscd"`           //机构号
	Mchntid         string `json:"mchntid"`         //商户号
	MerDiscount     string `json:"merDiscount"`     //商户优惠
	OrderNum        string `json:"orderNum"`        //订单号
	PayTime         string `json:"payTime"`         //支付时间
	Respcd          string `json:"respcd"`          //交易结果
	Sign            string `json:"sign"`            //签名
	SignType        string `json:"signType"`        //签名类型
	Terminalid      string `json:"terminalid"`      //终端号
	Txamt           string `json:"txamt"`           //订单金额
}

type XunLianNotify struct {
	Md5Key     string `xml:"-" json:"-"`
	notifyData *XunLianNotifyData
	notifyMap  map[string]interface{}
}

func (xunLian *XunLianNotify) getNotifyData(ret string) error {
	if ret == "" {
		return errors.New("获取通知数据失败")
	}

	ret = strings.ReplaceAll(strings.ReplaceAll(ret, "%3D", "="), "%26", "&")

	retUrl, err := url.Parse(ret)

	if err != nil {
		return errors.New("解析返回数据失败:" + err.Error())
	}

	values := retUrl.Query()

	mapData := maps.Values2Map(values)

	notify := new(XunLianNotifyData)

	err = helper.DeepCopy(mapData, notify)

	if err != nil {
		return errors.New("解析通知数据失败:" + err.Error())
	}

	xunLian.notifyData = notify
	xunLian.notifyMap = mapData

	return nil
}

func (xunLian *XunLianNotify) checkNotify() error {
	return xunLian.verifySign()
}

func (xunLian *XunLianNotify) verifySign() error {

	mapData := xunLian.notifyMap

	helper.KSort(&mapData)

	signStr := helper.CreateLinkString(&mapData)

	signStr += xunLian.Md5Key

	fmt.Println(signStr)

	sign := helper.Sha256(signStr)

	if sign != xunLian.notifyData.Sign {
		return errors.New("返回数据验签失败，可能数据被篡改")
	}

	return nil
}

func (xunLian *XunLianNotify) replyNotify(err error) string {

	if err != nil {
		fmt.Println(err)
		return "error"
	} else {
		return "success"
	}
}

type XunLianCallBack func(data *XunLianNotifyData) error

type XunLianMd5KeyCallBack func(merchantNo string) (md5Key string, err error)

func (xunLian *XunLianNotify) Handle(ret string, md5CallBack XunLianMd5KeyCallBack, callBack XunLianCallBack) string {

	err := xunLian.getNotifyData(ret)

	if err != nil {
		return xunLian.replyNotify(err)
	}

	md5Key, err := md5CallBack(xunLian.notifyData.Mchntid)

	if err != nil {
		return xunLian.replyNotify(err)
	}

	xunLian.Md5Key = md5Key

	err = xunLian.checkNotify()

	if err != nil {
		return xunLian.replyNotify(err)
	}

	err = callBack(xunLian.notifyData)

	if err != nil {
		return xunLian.replyNotify(err)
	}

	return xunLian.replyNotify(nil)
}
