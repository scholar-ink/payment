package notify

import (
	"fmt"
	"testing"
)

func TestHangNotify_Handle(t *testing.T) {

	notify := new(HangNotify)

	ret := map[string]interface{}{
		"amount":        "0.010000",
		"channelNo":     "0559988581",
		"code":          "SUCCESS",
		"merchantNo":    "DL20200323161220297431",
		"message":       "\346\224\257\344\273\230\346\210\220\345\212\237",
		"orderNo":       "200324130938448647792766",
		"sign":          "vcI7qPw2FdQzQENSUiAwrrD4yUPCuXTBvXdUwjcGmvbHjjZDcwAfKAScF6QHw1KqZLt65DLATiLpU3Dzm9ccBTN3nthAGHYqwvOCer43uueqfrSHobrqs2M+pCxDz4VX6SEe1zW5eJ7M3vEufXlkLXWCmnuvr9zFURAORbxHpoU=",
		"subMerchantNo": "IMCH20200324120454268314",
	}

	retData := notify.Handle(ret, func(merchantNo string) (md5Key string, err error) {
		return "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQD2i184JOoIUI6hpx4LzMRo++svvHU7tRwZYq4IZJ8jAP1uswhIhhmPqdgVyBBA/AlvipLdzqZdiyqQTAVinncoOhmMY4sOJnyx2JmzHDJXTBRiQspmK1JuFFX7rWj1qVaw9R4HZrvhB10ORUjOJkc4O2f5/wu73gsNgcJZjuVFNQIDAQAB", nil
	}, func(data *HangNotifyData) error {
		fmt.Printf("%+v\n", data)
		return nil
	})

	fmt.Println(retData)

}
