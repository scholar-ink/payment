package hlb

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type OrderCharge struct {
	*OrderChargeConf
	BaseCharge
}

type OrderChargeConf struct {
	P1BizType         string `json:"P1_bizType"`                   //交易类型
	P2OrderId         string `json:"P2_orderId"`                   //商户订单号
	P3CustomerNumber  string `json:"P3_customerNumber"`            //商户编号
	P4PayType         string `json:"P4_payType"`                   //支付类型
	P5OrderAmount     string `json:"P5_orderAmount"`               //交易金额
	P6Currency        string `json:"P6_currency"`                  //币种类型
	P7Authcode        string `json:"P7_authcode"`                  //授权码
	P8AppType         string `json:"P8_appType"`                   //客户端类型
	P9NotifyUrl       string `json:"P9_notifyUrl"`                 //通知回调地址
	P10SuccessToUrl   string `json:"P10_successToUrl"`             //成功跳转URL
	P11OrderIp        string `json:"P11_orderIp"`                  //商户IP
	P12GoodsName      string `json:"P12_goodsName"`                //商品名称
	P13GoodsDetail    string `json:"P13_goodsDetail"`              //商品详情
	P14Desc           string `json:"P14_desc"`                     //备注
	P16AppId          string `json:"P16_appId,omitempty"`          //公众号appId
	P17LimitCreditPay string `json:"P17_limitCreditPay,omitempty"` //是否限制借贷记
	P18GoodsTag       string `json:"P18_goodsTag,omitempty"`       //商品标记
	P19Guid           string `json:"P19_guid,omitempty"`           //微信进件时上送的唯一号
	P20MarketingRule  string `json:"P20_marketingRule,omitempty"`  //营销参数规则
	P21Identity       string `json:"P21_identity,omitempty"`       //实名参数
	SplitBillType     string `json:"splitBillType,omitempty"`      //分账类型
	RuleJson          string `json:"ruleJson,omitempty"`           //分账规则串
	HbfqNum           string `json:"hbfqNum,omitempty"`            //花呗分期数
	DeviceInfo        string `json:"deviceInfo,omitempty"`         //终端号
	TimeExpire        string `json:"timeExpire,omitempty"`         //超时时间
	Sign              string `json:"sign"`                         //签名
}

type SceneInfo struct {
	SceneType      string `json:"scene_type"`
	SceneBizType   string `json:"scene_biz_type"`
	AppName        string `json:"app_name"`
	AppPackageName string `json:"app_package_name"`
	WapName        string `json:"wap_name"`
	WapUrl         string `json:"wap_url"`
}

type Ext struct {
	SharingParams    []*SharingParam `json:"sharingParams,omitempty"`
	SharingNotifyUrl string          `json:"sharingNotifyUrl,omitempty"`
}

type SharingParam struct {
	FeeValue       string `json:"FeeValue"` //1：按比例 2：按固定金额
	FeeType        string `json:"FeeType"`  //按比例和不能大于1,按金额不能大于订单金额
	AccountType    string `json:"AccountType"`
	SharingAccount string `json:"SharingAccount"`
}

type OrderChargeReturn struct {
	RetCode string `json:"rt2_retCode"`
	RetMsg  string `json:"rt3_retMsg"`
	QrCode  string `json:"rt8_qrcode"`
}

func (oc *OrderCharge) Handle(conf *OrderChargeConf) (*OrderChargeReturn, error) {
	err := oc.BuildData(conf)
	if err != nil {
		return nil, err
	}
	oc.SetSign()
	ret, err := oc.SendReq(ChargeUrl, oc)
	if err != nil {
		return nil, err
	}
	return oc.RetData(ret)
}

func (oc *OrderCharge) RetData(ret []byte) (*OrderChargeReturn, error) {

	orderChargeReturn := new(OrderChargeReturn)

	err := json.Unmarshal(ret, &orderChargeReturn)

	if err != nil {
		return nil, err
	}

	if orderChargeReturn.RetCode != "0000" {
		return nil, errors.New(orderChargeReturn.RetMsg)
	}

	return orderChargeReturn, nil
}

func (oc *OrderCharge) BuildData(conf *OrderChargeConf) error {

	if conf.P4PayType == "SCAN" {
		conf.P7Authcode = "SCAN"
	}

	oc.OrderChargeConf = conf

	return nil
}

func (oc *OrderCharge) SetSign() {

	mapData := helper.Struct2Map(oc)

	signStr := CreateLinkString(&mapData)

	fmt.Println("&" + signStr + "&" + oc.Md5Key)

	oc.Sign = oc.makeSign("&" + signStr + "&" + oc.Md5Key)

	fmt.Println(oc.Sign)
}

type StringSort struct {
	data []string
}

func (s *StringSort) Len() int {
	return len(s.data)
}

func (s *StringSort) Swap(i, j int) {
	s.data[i], s.data[j] = s.data[j], s.data[i]
}

func (s *StringSort) Less(i, j int) bool {

	data := strings.Split(s.data[i], "_")

	numi, _ := strconv.Atoi(strings.Replace(data[0], "P", "", 1))

	data = strings.Split(s.data[j], "_")

	numj, _ := strconv.Atoi(strings.Replace(data[0], "P", "", 1))

	return numi < numj
}

func CreateLinkString(inputs *map[string]interface{}) string {

	var buf bytes.Buffer

	var keys []string

	for k := range *inputs {
		keys = append(keys, k)
	}

	ss := &StringSort{
		data: keys,
	}

	sort.Sort(ss)

	fmt.Println(ss.data)

	for _, k := range ss.data {

		if k != "sign" && k != "paySign" && k != "timeExpire" {

			v := (*inputs)[k]

			//if v == reflect.Zero(reflect.TypeOf(v)).Interface() {
			//	continue
			//}

			if buf.Len() > 0 {
				buf.WriteByte('&')
			}

			rt := reflect.TypeOf(v)

			switch rt.Kind() {
			case reflect.Int:
				buf.WriteString(strconv.Itoa(v.(int)))
			case reflect.Float64:
				buf.WriteString(strconv.Itoa(int(v.(float64))))
			case reflect.String:
				buf.WriteString(v.(string))
			}
		}
	}
	return buf.String()
}
