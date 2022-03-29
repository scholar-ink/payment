package notify

import (
	"fmt"
	"testing"
)

func TestFqlEnterNotify_Handle(t *testing.T) {
	notify := new(FqlEnterNotify)

	notify.InitBaseConfig(&FqlConfig{
		PlatPubKey: "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAjsptGNxDkK31cOOwLEKo5cHypVlN3zp/60iRh6zaiG7jGmOn64v7/H0Qh8fCukNAv7CYDG5c7S36n1Sa5M3kuFR0Qk7g+WwXdGMcE9rNKdqGkOecaBHHX6kpKVTBmhTFmhI1aoumnfxrFKuExWm4BY2UO2Uaq9bHB8ZkbcGQ9yq9I0nYGUIkneJN4NZhUudj4fK1SHvtHUx1VO2aCbmiIZi4YaAfvvEhIUSeYtwdBtkJEMn7tlQju46Xob5zVGNUcVKUnQf7SY2xUUghU1DOJmNJ0i1heGrs6zjeOLYv8nyQ0SzWVWmaTb7LYM1QL9KdqMpY0+xTq1v2cJ3bLl3bLwIDAQAB",
	})

	ret := `{"data":"{\"bizType\":\"01\",\"chUserId\":\"201129\",\"completeTime\":\"20200401173339\",\"customerId\":\"888031125bfa419f88cb51d78cdeb606\",\"orglOrderId\":\"200401165039161482194295\",\"status\":\"03\"}","merchId":"EW20190320162644","signInfo":"KLO0YSZxjJPGZ97kM3NXk4+LnkQoBP02Ix2jxY2m/BlMb3l2+wtr1OL546XjcjUrSCd2TEahJ8GTAkMgrkz+bmO2fWfbX41A4NWE6jysBqzr2cDM+B2DM3tyFmQTZJT4U7gNS1+HyBGkDLQXuZrk9aUQKLWHvzL6amPSbqCWYDr9mQtIAyMIPNtRUYkzYWoKqpZstA0RSjlEoDi17I5hAly3pkkWtJhJsAp/9vTqYU4aBvHUiLtTm1DWQbevZD5Qy8ZtPug/g7zDZ47AT/nYeclYEyyp3krE5Dupf8OXqRBxw7/x+8a0Ipx2ZB/9FRCVgT1ukRx/HrDALpwUOzgCew==","signType":"RSA"}`

	retData := notify.Handle(ret, func(data *FqlEnterNotifyData) error {
		fmt.Printf("%+v", data)
		return nil
	})

	fmt.Println(retData)
}
