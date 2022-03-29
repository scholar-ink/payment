package hj

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
	P0Version          string `json:"p0_Version"`                    //版本号
	P1MerchantNo       string `json:"p1_MerchantNo"`                 //商户编号
	P2OrderNo          string `json:"p2_OrderNo"`                    //商户订单号
	P3Amount           string `json:"p3_Amount"`                     //订单金额
	P4Cur              string `json:"p4_Cur"`                        //交易币种
	P5ProductName      string `json:"p5_ProductName"`                //商品名称
	P6ProductDesc      string `json:"p6_ProductDesc,omitempty"`      //商品描述
	P7Mp               string `json:"p7_Mp,omitempty"`               //公用回传参数
	P8ReturnUrl        string `json:"p8_ReturnUrl,omitempty"`        //商户页面通知地址
	P9NotifyUrl        string `json:"p9_NotifyUrl"`                  //服务器异步通知地址
	Q1FrpCode          string `json:"q1_FrpCode"`                    //交易类型
	Q2MerchantBankCode string `json:"q2_MerchantBankCode,omitempty"` //银行商户编码
	Q4IsShowPic        string `json:"q4_IsShowPic,omitempty"`        //是否展示图片
	Q5OpenId           string `json:"q5_OpenId,omitempty"`           //微信 Openid
	Q6AuthCode         string `json:"q6_AuthCode,omitempty"`         //付款码数字
	Q7AppId            string `json:"q7_AppId,omitempty"`            //APPID
	Q8TerminalNo       string `json:"q8_TerminalNo,omitempty"`       //终端号
	Q9TransactionModel string `json:"q9_TransactionModel,omitempty"` //微信/支付宝 H5 模式
	QaTradeMerchantNo  string `json:"qa_TradeMerchantNo,omitempty"`  //交易商户号
	QbBuyerId          string `json:"qb_buyerId,omitempty"`          //买家的支付宝 唯一用户号
	QcIsAlt            string `json:"qc_IsAlt,omitempty"`            //是否分账交易
	QdAltType          string `json:"qd_AltType,omitempty"`          //分账类型
	QeAltInfo          string `json:"qe_AltInfo,omitempty"`          //分账信息
	QfAltUrl           string `json:"qf_AltUrl,omitempty"`           //分账通知地址
	QgMarketingAmount  string `json:"qg_MarketingAmount,omitempty"`  //营销金额
}

type SharingParam struct {
	AltMchNo  string `json:"altMchNo"`
	AltAmount string `json:"altAmount"` //1：按比例 2：按固定金额
	IsGuar    string `json:"isGuar"`    //按比例和不能大于1,按金额不能大于订单金额
}

type OrderChargeReturn struct {
	RaCode    int64  `json:"ra_Code"`
	RbCodeMsg string `json:"rb_CodeMsg"`
	RcResult  string `json:"rc_Result"`
}

func (oc *OrderCharge) Handle(conf *OrderChargeConf) (*OrderChargeReturn, error) {
	conf.P0Version = "1.0"
	conf.P4Cur = "1"
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

	if orderChargeReturn.RaCode != 100 {
		return nil, errors.New(orderChargeReturn.RbCodeMsg)
	}

	return orderChargeReturn, nil
}

func (oc *OrderCharge) BuildData(conf *OrderChargeConf) error {

	oc.OrderChargeConf = conf
	return nil
}

func (oc *OrderCharge) SetSign() {

	mapData := helper.Struct2Map(oc)

	signStr := CreateLinkString(&mapData)

	oc.Hmac = oc.makeSign(signStr + oc.Md5Key)

	fmt.Println(oc.Hmac)
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

	numi := strings.Replace(strings.Replace(data[0], "p", "1", 1), "q", "2", 1)

	data = strings.Split(s.data[j], "_")

	numj := strings.Replace(strings.Replace(data[0], "p", "1", 1), "q", "2", 1)

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

		if k != "hmac" {

			v := (*inputs)[k]

			//if v == reflect.Zero(reflect.TypeOf(v)).Interface() {
			//	continue
			//}

			//if buf.Len() > 0 {
			//	buf.WriteByte(' ')
			//}

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
