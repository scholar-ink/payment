package notify

import (
	"fmt"
	"testing"
)

func TestQfNotify_Handle(t *testing.T) {

	notify := new(QfNotify)

	ret := `{"status": "1", "pay_type": "800101", "sysdtm": "2019-12-27 18:07:53", "paydtm": "2019-12-27 18:08:12", "goods_name": "\u6d4b\u8bd5\u5546\u54c1", "txcurrcd": "CNY", "txdtm": "2019-12-27 18:07:52", "mchid": "ZajNRulk9l", "txamt": "1", "chnlsn2": "", "out_trade_no": "191227180752992176881193", "syssn": "20191227000300020061834157", "cash_fee_type": "", "cancel": "0", "respcd": "0000", "goods_info": "", "cash_fee": "0", "notify_type": "payment", "chnlsn": "932019122722001476051433464997", "cardcd": ""}`

	retData := notify.Handle(ret, "CCAF44894A57E125A1C78CC4EEE90841", func(MchId string) (md5Key string, err error) {
		return "F3954E8B34F5474AA38BB127D5BA94FE", nil
	}, func(data *QfNotifyData) error {
		fmt.Printf("%+v\n", data)
		return nil
	})

	fmt.Println(retData)

}
