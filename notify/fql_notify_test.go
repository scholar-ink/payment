package notify

import (
	"fmt"
	"testing"
)

func TestFqlNotify_Handle(t *testing.T) {
	notify := new(FqlNotify)

	ret := `{"data":"{\"amount\":\"0.84\",\"channelType\":\"WX\",\"completeTime\":\"20200612123125\",\"exAccNo\":\"0011562100000024429092321\",\"feeAmtAmount\":\"0.01\",\"feeAmtRate\":\"0.0065\",\"orderId\":\"200612123119079827920731\",\"payExAccNo\":\"\",\"pmtType\":\"59\",\"retCode\":\"0000\",\"retMsg\":\"业务处理成功\",\"status\":\"02\",\"txnId\":\"S20061212311923661804W1qm0qBvvAK\"}","merchId":"EW20200427175142","signInfo":"ebdSZQOjlpUAFY9merLOQJg89Yegtb9hpBj8FBiZLH/hJl3X7flGdgaQyJf82q7QnWn8WdxIDbI+DK0FRps1RRqxsszwf/3BA52j4828+tl0OJY3GUVkJxHJxkONwWQ9wGr32MQr3EZi8w1oYasvoloVWSl9HMEbFZ4KmJ8+QSyv+XEHR7CqBWGlx8ryZJGKpXPQGkJTg4UlBDn0NecPzY2ug/QpX13lPmLcKpJC8PYL8WcIUwswpCJits3ICnhly4u7NPaeBO8OulCU6NBl3+hWQGOM6zpxMYFEspU/dEPaRBydpinkA/nOBkly5D3cyxrd7mhFdSIG7lmUnZLZKg==","signType":"RSA"}`

	retData := notify.Handle(ret, func(merchantNo string) (md5Key string, err error) {
		fmt.Println(merchantNo)
		return "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAj5TuDhWkB0t3wYbau7ZLhwreqlBN6lDjRG/hn5165UvwpVl+l0nsr8CiIBXMXsmjKbs6Kgzj+2Zg73CO5EDoXRfT7lnsr14hJCN4mhaZQcd3rUs+ubEhYXoqZHYpTS2tUACTR4eq3R/4fDcRnqgYg7vI+yHUWMRXHbv8ldVSx2izrHGOuQdFAt+r5mqVsmlk422elTe99HR59jrQAK5Qb+mdzMLdLZQQl0Hc3fVNv90m7g0lvXx6UmMXXMD8JajLmwI24zgIhGF2iOxtqSA7MKkzeyz6KTTFR70QE+7sM0Kuw0P/6xis0wK1mhR4tlN3f4I3rdjSZfavOYCwqTKJHQIDAQAB", nil
	}, func(data *FqlNotifyData) error {
		fmt.Printf("%+v", data)
		return nil
	})

	fmt.Println(retData)
}
