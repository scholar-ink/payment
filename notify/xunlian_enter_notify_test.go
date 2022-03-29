package notify

import (
	"errors"
	"fmt"
	"testing"
)

func TestXunLianEnterNotify_Handle(t *testing.T) {
	notify := new(XunLianEnterNotify)

	ret := `{"message":"[\"集团商户代码，未找到对应集团商户\"]","status":2,"data":{"timestamp":1583220154667,"reqId":"4894522b52bd4d948d8c8c81de042d92","insCode":"19940909","type":0,"bussType":0}}`

	retData := notify.Handle(ret, func(data *XunLianEnterNotifyData) error {
		fmt.Printf("%+v\n", data.Data)
		return errors.New("111")
	})

	fmt.Println(retData)
}
