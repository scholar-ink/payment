package gpay

import (
	"encoding/json"
	"errors"
	"github.com/scholar-ink/payment/helper"
	"time"
)

type OrderCharge struct {
	*OrderChargeConf
	BaseCharge
}

type OrderChargeConf struct {
	PaymentType string     `json:"payment_type"`           //支付方式
	NonceStr    string     `json:"nonce_str"`              //随机字符串
	OutTradeNo  string     `json:"out_trade_no"`           //商户订单号
	TotalFee    int64      `json:"total_fee"`              //订单金额
	FeeType     string     `json:"fee_type"`               //币种
	MchCreateIp string     `json:"mch_create_ip"`          //订单提交终端IP
	TimeStart   string     `json:"time_start"`             //订单生成时间格式：yyyyMMddHHmmss
	TimeExpire  string     `json:"time_expire,omitempty"`  //订单失效时间格式：yyyyMMddHHmmss
	NotifyUrl   string     `json:"notify_url"`             //订单通知URL
	CallbackUrl string     `json:"callback_url,omitempty"` //订单回调URL
	DeviceInfo  string     `json:"device_info,omitempty"`  //设备号
	Body        string     `json:"body"`                   //商品描述
	Attach      string     `json:"attach,omitempty"`       //商品附加信息
	GoodsTag    string     `json:"goods_tag,omitempty"`    //商品标记，微信平台配置的商品标记，用于优惠券或者满减使用
	ProductId   string     `json:"product_id,omitempty"`   //商品ID，商户在微信平台配置的商品ID，用于在二维码中包含商品ID
	OpUserId    string     `json:"op_user_id,omitempty"`   //操作员ID,微信公众号支付时,值为微信用户openid。 支付宝服务窗支付时，值为买家支付宝用户ID
	SceneInfo   *SceneInfo `json:"scene_info,omitempty"`   //操作员ID,微信公众号支付时,值为微信用户openid。 支付宝服务窗支付时，值为买家支付宝用户ID
	AuthCode    string     `json:"auth_code,omitempty"`    //扫码支付授权码
	LimitPay    string     `json:"limit_pay,omitempty"`    //支付方式限制，如限制不能用信用卡支付等1-限定不能使用信用卡
	GoodsType   string     `json:"goods_type,omitempty"`   //商品类别， 当payment_type=pay-jd-native时，此字段不能为空
	Ext         *Ext       `json:"ext,omitempty"`          //扩展字段，类型为map格式json串
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
	ResultCode  string `json:"result_code"`
	ErrorCode   string `json:"err_code"`
	ErrCodeDesc string `json:"err_code_des"`
	OutTradeNo  string `json:"out_trade_no"`
	OrderNo     string `json:"order_no"`
	CodeUrl     string `json:"code_url"`
	CodeImgUrl  string `json:"code_img_url"`
}

func (oc *OrderCharge) Handle(conf *OrderChargeConf) (*OrderChargeReturn, error) {
	err := oc.BuildData(conf)
	if err != nil {
		return nil, err
	}
	oc.SetSign()
	ret, err := oc.SendReq(ChargeUrl)
	if err != nil {
		return nil, err
	}
	return oc.RetData(ret)
}

func (oc *OrderCharge) RetData(ret []byte) (*OrderChargeReturn, error) {

	ret, err := oc.BaseCharge.RetData(ret)

	if err != nil {
		return nil, err
	}

	orderChargeReturn := new(OrderChargeReturn)

	err = json.Unmarshal(ret, &orderChargeReturn)

	if err != nil {
		return nil, err
	}

	if orderChargeReturn.ResultCode != "0" {
		return nil, errors.New(orderChargeReturn.ErrCodeDesc)
	}

	return orderChargeReturn, nil
}

func (oc *OrderCharge) BuildData(conf *OrderChargeConf) error {

	if conf.PaymentType == "" {
		return errors.New("支付方式不能为空")
	}
	if conf.OutTradeNo == "" {
		return errors.New("支付方式不能为空")
	}
	if len(conf.OutTradeNo) > 32 {
		return errors.New("商户订单号不能超过32")
	}
	if conf.TotalFee <= 0 {
		return errors.New("订单金额需大于0")
	}
	if conf.FeeType == "" {
		return errors.New("币种不能为空")
	}
	if conf.MchCreateIp == "" {
		return errors.New("终端IP不能为空")
	}
	if conf.NotifyUrl == "" {
		return errors.New("订单通知URL不能为空")
	}
	if conf.Body == "" {
		return errors.New("商品描述不能为空")
	}
	if _, ok := map[string]bool{"pay-wx-service": true, "pay-wx-service-online": true, "pay-zfb-service": true}[conf.PaymentType]; ok && conf.OpUserId == "" {
		return errors.New("操作员ID不能为空")
	}

	conf.TimeStart = time.Now().Format("20060102150405")
	conf.NonceStr = helper.NonceStr()

	if conf.TimeExpire == "" {
		conf.TimeExpire = time.Now().Add(30 * time.Minute).Format("20060102150405")
	}

	b, err := json.Marshal(conf)

	if err != nil {
		return err
	}

	oc.OrderChargeConf = conf

	oc.SourceData = string(b)

	oc.ServiceName = "submit.order"

	return nil
}
