/**
 * @author dengmeiyu
 * @since 20180618
 */
package query

import (
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {

	app := new(OrderQuery)

	app.InitBaseConfig(&BaseConfig{
		AppId:    "wxf06ac118ca3d9533",
		SubAppId: "wxa33cba2b69f869f3",
		MchId:    "1495589652",
		SubMchId: "1495746312",
		Md5Key:   "057177a8459352933f755c535b0ab0ef",
		SignType: "MD5",
	})

	ret, err := app.OrderQuery(map[string]interface{}{
		"out_trade_no": "2018061511000063151",
	})

	fmt.Printf("%+v", ret)

	fmt.Println(err)
}
