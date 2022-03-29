package allinpay

import (
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/map"
)

type OrderCharge struct {
	*OrderChargeConf
	BaseCharge
}

type OrderChargeConf struct {
	Oid        string `json:"oid"`
	Amt        int64  `json:"amt"`
	TrxReserve string `json:"trxreserve"`
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
	CodeUrl string `json:"code_url"`
}

func (oc *OrderCharge) Handle(conf *OrderChargeConf) (*OrderChargeReturn, error) {

	err := oc.BuildData(conf)

	if err != nil {
		return nil, err
	}
	oc.SetSign()

	mapData := helper.Struct2Map(oc)

	values := maps.Map2Values(&mapData)

	values.Del("key")

	orderChargeReturn := new(OrderChargeReturn)

	orderChargeReturn.CodeUrl = ChargeUrl + "?" + values.Encode()

	return orderChargeReturn, nil
}

func (oc *OrderCharge) BuildData(conf *OrderChargeConf) error {

	oc.OrderChargeConf = conf

	return nil
}

func (oc *OrderCharge) SetSign() {

	mapData := helper.Struct2Map(oc)

	signStr := helper.CreateLinkString(&mapData)

	oc.Sign = oc.makeSign(signStr)
}
