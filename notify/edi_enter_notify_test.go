package notify

import (
	"fmt"
	"testing"
)

func TestEdiEnterNotify_Handle(t *testing.T) {
	notify := new(EdiEnterNotify)

	ret := `{"audMsg":"录入审核通过","audSts":"1","backType":"0","mchType":"1","mercId":"999652853110005","outMercId":"203184","seqNo":"200523152220805743","sign":"TUdRZt+JSHk5SyWEgsuGUu1YtrSIhbatlR+I4KSBSEtBcGmyfuhPKs8+U5gAcLFIXbOWDE51EUvOfWkDb9qbWubp+GLvV//7IOC2JeTd2KX7aPd1r8oWxtg+9amC0VWXQBRWlmo76o3dooIwUAMzJFPienPi/E5tosj88bn0Xo0=","status":"0","upMerId":"872652853110002"}`

	retData := notify.Handle(ret, func(data *EdiEnterNotifyData) error {
		fmt.Println(data)
		return nil
	})

	fmt.Println(retData)
}
