package withdrawal

import (
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"io/ioutil"
	"testing"
)

func TestZgWithdrawal_Handle(t *testing.T) {
	zgWithdrawal := new(ZgWithdrawal)

	pfxData, _ := ioutil.ReadFile("udian2.pfx")

	zgWithdrawal.InitConfig(&ZgConfig{
		PartnerId:    "b1564976173998427",
		PfxData:      pfxData,
		CertPassWord: "Udian888",
		Md5Key:       "AT29WBOWHI02LAPX",
	})
	orderSn := helper.CreateSn()
	zgHandleReturn, err := zgWithdrawal.Handle(&ZgHandleConfig{
		OutTradeNo:      orderSn,
		AccountNo:       "b14382166832945152",
		Amount:          "459730",
		ServerReturnUrl: "http://web.udian.me/v1/common/enter-notify",
		Subject:         "1111体现",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(zgHandleReturn)
}
func TestZgWithdrawal_Query(t *testing.T) {
	zgWithdrawal := new(ZgWithdrawal)

	pfxData, _ := ioutil.ReadFile("udian2.pfx")

	zgWithdrawal.InitConfig(&ZgConfig{
		PartnerId:    "b1537337771064377",
		PfxData:      pfxData,
		CertPassWord: "Udian888",
		Md5Key:       "GIX4BRVHFT61KBVU",
	})
	zgHandleReturn, err := zgWithdrawal.Query(&ZgQueryConfig{
		OutTradeNo: "b14363811134003200",
		AccountNo:  "b14363660774905856",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(zgHandleReturn)
}
func TestZgWithdrawal_Balance(t *testing.T) {
	zgWithdrawal := new(ZgWithdrawal)

	pfxData, _ := ioutil.ReadFile("udian2.pfx")

	zgWithdrawal.InitConfig(&ZgConfig{
		PartnerId:    "100220190806705",
		PfxData:      pfxData,
		CertPassWord: "Udian888",
		Md5Key:       "AT29WBOWHI02LAPX",
	})
	accountList, err := zgWithdrawal.Balance(&ZgBalanceConfig{
		AccountNo: "b14382166832945152",
	})

	if err != nil {
		fmt.Println(err)
	}

	for _, account := range accountList {
		fmt.Println(account.AccountType)
		fmt.Println(account.CanUseBalance)
		fmt.Println(account.Status)
		fmt.Println(account.FreezeAmount)
	}

	fmt.Println(accountList)
}
