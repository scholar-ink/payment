package notify

import (
	"fmt"
	"testing"
)

func TestCmEnterNotify_Handle(t *testing.T) {
	notify := new(CmEnterNotify)

	ret := `{"shop_id":"f13db1959ed46e3cb1dfef8ea88bace5","error_msg":"商户负责人身份证号不能为空，请查证","result_code":"fail","order_id":"200217195846912743068828"}`

	retData := notify.Handle(ret, func(data *CmEnterNotifyData) error {
		fmt.Println(data)
		return nil
	})

	fmt.Println(retData)
}
