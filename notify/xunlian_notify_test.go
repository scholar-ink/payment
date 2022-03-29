package notify

import (
	"fmt"
	"testing"
)

func TestXunLianNotify_Handle(t *testing.T) {

	notify := new(XunLianNotify)

	ret := `/v1/common/enter-notify?bankType%3DDEBIT%26busicd%3DPAUT%26channelOrderNum%3D2020030622001476050548996279%26charset%3Dutf-8%26chcd%3DALP%26consumerAccount%3D135%2A%2A%2A%2A1080%26consumerId%3D2088612472076052%26errorDetail%3DSUCCESS%26inscd%3D91681888%26mchntid%3D168791153110001%26orderNum%3D200306140936389990067318%26payTime%3D2020-03-06+14%3A11%3A02%26respcd%3D00%26signType%3DSHA256%26terminalid%3Dyunc0001%26transTime%3D2020-03-06+14%3A09%3A36%26txamt%3D000000000001%26txndir%3DA%26version%3D2.3.1%26sign%3D3cae801230f02152391eac77a3868c54cfd031aeb7ee4439095489533e918cf8`

	retData := notify.Handle(ret, func(MchId string) (md5Key string, err error) {
		return "b124ae4174e95d94a929a377c88de833", nil
	}, func(data *XunLianNotifyData) error {
		fmt.Printf("%+v\n", data)
		return nil
	})

	fmt.Println(retData)

}
