package notify

import (
	"errors"
	"fmt"
	"github.com/scholar-ink/payment/helper"
)

type AllInPayNotifyData struct {
	Key         string `json:"key"`
	TrxCode     string `json:"trxcode"`               //交易类型
	AppId       string `json:"appid"`                 //通联分配的appid
	CusId       string `json:"cusid"`                 //商户号
	Timestamp   string `json:"timestamp"`             //调用时间戳
	RandomStr   string `json:"randomstr"`             //随机字符串
	Sign        string `json:"sign"`                  //sign校验码
	BizSeq      string `json:"bizseq"`                //业务流水
	TrxStatus   string `json:"trxstatus"`             //交易结果码
	Amount      string `json:"amount"`                //金额,单位到分
	TrxId       string `json:"trxid"`                 //交易流水号
	SrcTrxId    string `json:"srctrxid"`              //原交易流水
	TrxDay      string `json:"trxday,omitempty"`      //交易请求日期
	PayTime     string `json:"paytime,omitempty"`     //交易完成时间
	TermAuthno  string `json:"termauthno"`            //终端编码
	TermId      string `json:"termid"`                //终端编码
	TermRefNum  string `json:"termrefnum"`            //终端编码
	TermBatchid string `json:"termbatchid,omitempty"` //终端批次号
	TraceNo     string `json:"traceno"`               //终端流水
	TrxReserve  string `json:"trxreserve,omitempty"`  //业务关联内容
	AcctType    string `json:"accttype,omitempty"`    //借贷标志
	Acct        string `json:"acct,omitempty"`        //交易帐号
	Fee         string `json:"fee,omitempty"`         //手续费
	SignType    string `json:"signtype,omitempty"`    //签名类型
}

type AllInPayNotify struct {
	notifyData *AllInPayNotifyData
}

func (zg *AllInPayNotify) getNotifyData(ret map[string]string) error {

	if len(ret) == 0 {
		return errors.New("获取通知数据失败")
	}

	notify := new(AllInPayNotifyData)

	err := helper.DeepCopy(ret, notify)

	if err != nil {
		return errors.New("解析返回数据失败:" + err.Error())
	}

	zg.notifyData = notify

	return nil
}

func (zg *AllInPayNotify) checkNotify() error {

	if zg.notifyData.TrxStatus != "0000" {
		return nil
	}
	return zg.verifySign()
}

func (zg *AllInPayNotify) verifySign() error {

	mapData := helper.Struct2Map(zg.notifyData)

	delete(mapData, "sign")

	signStr := helper.CreateLinkString(&mapData)

	signStr = helper.Md5(signStr)

	if zg.notifyData.Sign != signStr {
		return errors.New("返回数据验签失败，可能数据被篡改")
	}

	return nil
}

func (zg *AllInPayNotify) replyNotify(err error) string {

	if err != nil {
		fmt.Println(err)
		return "error"
	} else {
		return "success"
	}
}

type AllInPayCallBack func(data *AllInPayNotifyData) error

type AllInPayMd5KeyCallBack func(merchantNo string) (md5Key string, err error)

func (zg *AllInPayNotify) Handle(ret map[string]string, md5CallBack AllInPayMd5KeyCallBack, callBack AllInPayCallBack) string {

	err := zg.getNotifyData(ret)

	if err != nil {
		return zg.replyNotify(err)
	}

	md5Key, err := md5CallBack(zg.notifyData.AppId)

	if err != nil {
		return zg.replyNotify(err)
	}

	zg.notifyData.Key = md5Key

	err = zg.checkNotify()

	if err != nil {
		return zg.replyNotify(err)
	}

	err = callBack(zg.notifyData)

	if err != nil {
		return zg.replyNotify(err)
	}

	return zg.replyNotify(nil)
}
