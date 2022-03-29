package notify

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/array"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type HlbNotifyData struct {
	CustomerNumber          string `json:"rt1_customerNumber"`                //商户号
	OrderId                 string `json:"rt2_orderId"`                       //商户订单号
	SystemSerial            string `json:"rt3_systemSerial"`                  //平台流水号
	Status                  string `json:"rt4_status"`                        //订单状态
	OrderAmount             string `json:"rt5_orderAmount"`                   //订单金额
	Currency                string `json:"rt6_currency"`                      //币种
	Timestamp               string `json:"rt7_timestamp"`                     //通知时间
	Desc                    string `json:"rt8_desc"`                          //备注
	OpenId                  string `json:"rt10_openId,omitempty"`             //用户openId
	ChannelOrderNum         string `json:"rt11_channelOrderNum,omitempty"`    //第三方平台订单号
	OrderCompleteDate       string `json:"rt12_orderCompleteDate,omitempty"`  //订单完成时间
	OnlineCardType          string `json:"rt13_onlineCardType"`               //支付卡类型
	CashFee                 string `json:"rt14_cashFee,omitempty"`            //上游返回 :现金支付金额
	CouponFee               string `json:"rt15_couponFee,omitempty"`          //上游返回:现金券金额
	FundBillList            string `json:"rt16_fundBillList,omitempty"`       //支付宝使用的资金渠道和优惠信息
	OutTransactionOrderId   string `json:"rt17_outTransactionOrderId"`        //微信支付宝交易订单号
	BankType                string `json:"rt18_bankType,omitempty"`           //用户付款银
	SubOpenId               string `json:"rt19_subOpenId,omitempty"`          //
	OrderAttribute          string `json:"rt20_orderAttribute,omitempty"`     //通道订单属性
	MarketingRule           string `json:"rt21_marketingRule,omitempty"`      //营销参数规则
	PromotionDetail         string `json:"rt22_promotionDetail,omitempty"`    //优惠信息详情
	PaymentAmount           string `json:"rt23_paymentAmount,omitempty"`      //实际支付金额
	CreditAmount            string `json:"rt24_creditAmount,omitempty"`       //入账面额
	AppId                   string `json:"rt25_appId,omitempty,omitempty"`    //子商户公众号sub_appid
	AppPayType              string `json:"rt26_appPayType,omitempty"`         //客户端类型
	PayType                 string `json:"rt27_payType,omitempty"`            //支付类型
	RuleJson                string `json:"ruleJson,omitempty"`                //分账规则及状态
	ProductFee              string `json:"productFee,omitempty"`              //交易手续费
	ChannelSettlementAmount string `json:"channelSettlementAmount,omitempty"` //渠道结算金额
	RealCreditAmount        string `json:"realCreditAmount,omitempty"`        //商户实际入账发生额
	TradeType               string `json:"tradeType,omitempty"`               //微信交易类型
	ChargeFlag              string `json:"chargeFlag,omitempty"`              //渠道支付宝费率活动标识
	UpAddData               string `json:"upAddData,omitempty"`               //银联二维码返回的付款方附加数据
	ResvData                string `json:"resvData,omitempty"`                //银联二维码返回的保留数据
	Sign                    string `json:"sign"`                              //银联二维码返回的保留数据
}

type HlbNotify struct {
	Md5Key     string `xml:"-" json:"-"`
	notifyData *HlbNotifyData
}

func (zg *HlbNotify) getNotifyData(ret map[string]string) error {

	if len(ret) == 0 {
		return errors.New("获取通知数据失败")
	}

	notify := new(HlbNotifyData)

	err := helper.DeepCopy(ret, notify)

	if err != nil {
		return errors.New("解析返回数据失败:" + err.Error())
	}

	zg.notifyData = notify

	return nil
}

func (zg *HlbNotify) checkNotify() error {
	return zg.verifySign()
}

func (zg *HlbNotify) verifySign() error {

	mapData := helper.Struct2Map(zg.notifyData)

	signStr := CreateLinkString(&mapData)

	fmt.Println("&" + signStr + "&" + zg.Md5Key)

	signStr = helper.Md5("&" + signStr + "&" + zg.Md5Key)

	fmt.Println(signStr)
	fmt.Println(zg.notifyData.Sign)

	if zg.notifyData.Sign != signStr {
		return errors.New("返回数据验签失败，可能数据被篡改")
	}

	return nil
}

func (zg *HlbNotify) replyNotify(err error) string {

	if err != nil {
		fmt.Println(err)
		return "error"
	} else {
		return "success"
	}
}

type HlbCallBack func(data *HlbNotifyData) error

type HlbMd5KeyCallBack func(merchantNo string) (md5Key string, err error)

func (zg *HlbNotify) Handle(ret map[string]string, md5CallBack HlbMd5KeyCallBack, callBack HlbCallBack) string {

	err := zg.getNotifyData(ret)

	if err != nil {
		return zg.replyNotify(err)
	}

	md5Key, err := md5CallBack(zg.notifyData.CustomerNumber)

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

	numi, _ := strconv.Atoi(strings.Replace(data[0], "rt", "", 1))

	data = strings.Split(s.data[j], "_")

	numj, _ := strconv.Atoi(strings.Replace(data[0], "rt", "", 1))

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

	exceptSign := []string{
		"sign",
		"rt10_openId",
		"rt11_channelOrderNum",
		"rt12_orderCompleteDate",
		"rt13_onlineCardType",
		"rt14_cashFee",
		"rt15_couponFee",
		"rt16_fundBillList",
		"rt17_outTransactionOrderId",
		"rt18_bankType",
		"rt19_subOpenId",
		"rt20_orderAttribute",
		"rt21_marketingRule",
		"rt22_promotionDetail",
		"rt23_paymentAmount",
		"rt24_creditAmount",
		"rt25_appId",
		"rt26_appPayType",
		"rt27_payType",
		"ruleJson",
		"productFee",
		"channelSettlementAmount",
		"realCreditAmount",
		"tradeType",
		"chargeFlag",
		"upAddData",
		"resvData",
	}

	for _, k := range ss.data {

		if !array.InArray(k, exceptSign) {

			v := (*inputs)[k]

			//if v == reflect.Zero(reflect.TypeOf(v)).Interface() {
			//	continue
			//}

			//prefix := k + "="

			if buf.Len() > 0 {
				buf.WriteByte('&')
			}

			rt := reflect.TypeOf(v)

			//buf.WriteString(prefix)

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
