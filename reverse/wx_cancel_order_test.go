/**
 * @author dengmeiyu
 * @since 20180618
 */
package reverse

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {

	app := new(OrderReverse)

	app.InitBaseConfig(&BaseConfig{
		AppId:    "wxf06ac118ca3d9533",
		SubAppId: "2018060615095472888",
		MchId:    "1495589652",
		SubMchId: "1495746312",
		Md5Key:   "057177a8459352933f755c535b0ab0ef",
		SignType: "MD5",
	})

	bool := app.Cancel(map[string]interface{}{
		"out_trade_no": "2018061511010705083",
	}, 0)

	fmt.Printf("%+v", bool)
}
