package notify

import (
	"fmt"
	"testing"
)

func TestZgNotify_Handle(t *testing.T) {
	notify := new(ZgNotify)

	ret := `{"requestTime":"20190717213522","signData":"E6CEC981C8D6F08298288A1CDD249698","requestId":"a83c8fd6-58ae-4ff8-8e06-61830cc26896","serviceName":"notify.pay","sourceData":"{\"divideStatus\":\"S\",\"divideTime\":\"2019-07-17 21:35:23\",\"divideDetailList\":\"[{\\\"divideAmount\\\":\\\"11.00\\\",\\\"divideAccountNo\\\":\\\"b14325090529794048\\\"},{\\\"divideAmount\\\":\\\"0\\\",\\\"divideAccountNo\\\":\\\"b1552470439457613\\\"}]\",\"payOrderNo\":3303289254730752,\"divideAmount\":\"11.00\",\"outTradeNo\":\"0001_201907172135172190639584\"}","merchantNo":100807}`

	result := notify.Handle(ret, func(merchantNo string) (md5Key string, err error) {
		return "b2ab57c9cd3d44d19d6514cfb5755a66", nil
	}, func(data *ZgNotifyData) error {
		fmt.Println(data)
		return nil
	})

	fmt.Println(result)
}
