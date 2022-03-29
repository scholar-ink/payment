package notify

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type HjNotifyData struct {
	R1MerchantNo  string `json:"r1_MerchantNo"`   //商户编号
	R2OrderNo     string `json:"r2_OrderNo"`      //商户订单号
	R3Amount      string `json:"r3_Amount"`       //支付金额
	R4Cur         string `json:"r4_Cur"`          //交易币种
	R5Mp          string `json:"r5_Mp,omitempty"` //公用回传参 数
	R6Status      string `json:"r6_Status"`       //支付状态
	R7TrxNo       string `json:"r7_TrxNo"`        //交易流水号
	R8BankOrderNo string `json:"r8_BankOrderNo"`  //银行订单号
	R9BankTrxNo   string `json:"r9_BankTrxNo"`    //银行流水号
	RaPayTime     string `json:"ra_PayTime"`      //支付时间
	RbDealTime    string `json:"rb_DealTime"`     //交易结果通 知时间
	RcBankCode    string `json:"rc_BankCode"`     //银行编码
	Hmac          string `json:"hmac"`            //签名数据
}

type HjNotify struct {
	Md5Key     string `xml:"-" json:"-"`
	notifyData *HjNotifyData
}

func (zg *HjNotify) getNotifyData(ret map[string]string) error {

	if len(ret) == 0 {
		return errors.New("获取通知数据失败")
	}

	notify := new(HjNotifyData)

	err := helper.DeepCopy(ret, notify)

	if err != nil {
		return errors.New("解析返回数据失败:" + err.Error())
	}

	zg.notifyData = notify

	return nil
}

func (zg *HjNotify) checkNotify() error {
	return zg.verifySign()
}

func (zg *HjNotify) verifySign() error {

	mapData := helper.Struct2Map(zg.notifyData)

	signStr := zg.CreateLinkString(&mapData)

	signStr = helper.Md5(signStr + zg.Md5Key)

	if zg.notifyData.Hmac != signStr {
		return errors.New("返回数据验签失败，可能数据被篡改")
	}

	return nil
}

func (zg *HjNotify) replyNotify(err error) string {

	if err != nil {
		fmt.Println(err)
		return "error"
	} else {
		return "success"
	}
}

type HjCallBack func(data *HjNotifyData) error

type HjMd5KeyCallBack func(merchantNo string) (md5Key string, err error)

func (zg *HjNotify) Handle(ret map[string]string, md5CallBack HjMd5KeyCallBack, callBack HjCallBack) string {

	err := zg.getNotifyData(ret)

	if err != nil {
		return zg.replyNotify(err)
	}

	md5Key, err := md5CallBack(zg.notifyData.R1MerchantNo)

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

type HjStringSort struct {
	data []string
}

func (s *HjStringSort) Len() int {
	return len(s.data)
}

func (s *HjStringSort) Swap(i, j int) {
	s.data[i], s.data[j] = s.data[j], s.data[i]
}

func (s *HjStringSort) Less(i, j int) bool {

	data := strings.Split(s.data[i], "_")

	numi := strings.Replace(data[0], "r", "1", 1)

	data = strings.Split(s.data[j], "_")

	numj := strings.Replace(data[0], "r", "1", 1)

	return numi < numj
}

func (zg *HjNotify) CreateLinkString(inputs *map[string]interface{}) string {

	var buf bytes.Buffer

	var keys []string

	for k := range *inputs {
		keys = append(keys, k)
	}

	ss := &HjStringSort{
		data: keys,
	}

	sort.Sort(ss)

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
