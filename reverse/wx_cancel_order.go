/**
 * @author dengmeiyu
 * @since 20180618
 */
package reverse

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"github.com/scholar-ink/payment/helper"
	"io/ioutil"
	"strings"
	"time"
)

const (
	ReverseReqUrl = "https://api.mch.weixin.qq.com/secapi/pay/reverse"
	SUCCESS       = "SUCCESS"
)

type OrderReverse struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	*BaseConfig
	*OrderOutTradeNo
}

type BaseConfig struct {
	AppId          string        `xml:"appid" json:"appid"`
	MchId          string        `xml:"mch_id" json:"mch_id"`
	SubAppId       string        `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`
	SubMchId       string        `xml:"sub_mch_id,omitempty" json:"sub_mch_id,omitempty"`
	TimeStart      string        `xml:"time_start,omitempty" json:"time_start,omitempty"`
	TimeExpire     string        `xml:"time_expire,omitempty" json:"time_expire,omitempty"`
	NotifyUrl      string        `xml:"notify_url,omitempty" json:"notify_url,omitempty"`
	Md5Key         string        `xml:"-" json:"-"`
	SignType       string        `xml:"sign_type,omitempty" json:"sign_type,omitempty"`
	Sign           string        `xml:"sign" json:"sign"`
	NonceStr       string        `xml:"nonce_str" json:"nonce_str"`
	ExpireDuration time.Duration `xml:"-" json:"-"`
	Cert           string        `xml:"-" json:"-"`
	Key            string        `xml:"-" json:"-"`
}

type OrderOutTradeNo struct {
	OutTradeNo string `xml:"out_trade_no" json:"out_trade_no"`
}
type Error struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	ResultCode string `xml:"result_code"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`
	PrepayId   string `xml:"prepay_id"`
	CodeUrl    string `xml:"code_url"`
	Recall     string `xml:"recall"`
}

type Return struct {
	AppId      string `xml:"appid"`
	SubAppId   string `xml:"sub_appid"`
	MchId      string `xml:"mch_id"`
	DeviceInfo string `xml:"device_info"`
	NonceStr   string `xml:"nonce_str"`
	TradeState string `xml:"trade_state"`
}
type ReverseResult struct {
	Error
	Return
}

func (rev *OrderReverse) Reverse(data map[string]interface{}) (ReverseResult, error) {
	err := rev.BuildData(data)

	if err != nil {
		return ReverseResult{}, err
	}
	rev.SetSign(rev)
	ret := rev.SendReq(ReverseReqUrl, rev)
	return rev.RetData(ret)

}

/**
*
* 撤销订单，如果失败会重复调用10次
* @param string $out_trade_no
* @param 调用深度 $depth
 */
func (rev *OrderReverse) Cancel(data map[string]interface{}, depth int) bool {
	if depth > 10 {
		return false
	}

	result, _ := rev.Reverse(data)

	//接口调用失败
	if result.ReturnCode != "SUCCESS" {
		return false
	}

	//如果结果为success且不需要重新调用撤销，则表示撤销成功
	if result.ReturnCode != "SUCCESS" && result.Recall == "N" {
		return true
	} else if result.Recall == "Y" {
		depth++
		return rev.Cancel(data, depth)
	}
	return false
}

func (rev *OrderReverse) BuildData(conf map[string]interface{}) error {

	var OrderOutTradeNo OrderOutTradeNo
	b, _ := json.Marshal(conf)
	json.Unmarshal(b, &OrderOutTradeNo)

	rev.OrderOutTradeNo = &OrderOutTradeNo

	return nil
}

func (rev *OrderReverse) RetData(ret []byte) (ReverseResult, error) {

	result := ReverseResult{}
	xml.Unmarshal(ret, &result)
	return result, nil

}

func (que *OrderReverse) InitBaseConfig(config *BaseConfig) {

	config.NonceStr = helper.NonceStr()

	que.BaseConfig = config
}

func (que *OrderReverse) SendReq(reqUrl string, pay interface{}) (b []byte) {

	buffer := bytes.NewBuffer(b)

	err := xml.NewEncoder(buffer).Encode(pay)

	if err != nil {
		return nil
	}

	client := helper.NewHttpClient()

	httpResp, err := client.Post(reqUrl, "text/xml; charset=utf-8", buffer)

	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	b, err = ioutil.ReadAll(httpResp.Body)

	return
}

func (que *OrderReverse) SetSign(pay interface{}) {

	mapData := helper.Struct2Map(pay)

	signStr := helper.CreateLinkString(&mapData)

	que.Sign = que.makeSign(signStr)
}

func (que *OrderReverse) makeSign(sign string) string {

	switch que.SignType {

	case "MD5":
		sign += "&key=" + que.Md5Key

		sign = helper.Md5(sign)
	}

	return strings.ToUpper(sign)
}
