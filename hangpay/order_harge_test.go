package hangpay

import (
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"testing"
	"time"
)

func TestCharge_Handle(t *testing.T) {
	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		MerchantNo: "DL20200323161220297431",
		Key:        "MIICeAIBADANBgkqhkiG9w0BAQEFAASCAmIwggJeAgEAAoGBAL76r2mrtxb0MwiYF+hnFSuS2roduVmLFvlczIlpsxOfXCI402uuEggUX5xNsreRi4qUw/lJd6zoW3/+kR+G1aoL7wJ7hUqSFv9K9CgoMVH9UR2HYoRH++yD2hHz/AsVgsdB0wDAOos4LzPmh0Oo5E+m6mgA2A2UffichfKUGvzHAgMBAAECgYBNqO8Pz24CfYcLJZ4DOXfYYj1jlZa7nN3YxS2/ayYRAqqal/URQpW+h1ph2w6jqyHNVrVid8ecnlgm8kPdSB019wxuQQR+rOiOCiwO4b+9CaEfqBaxTXK45+ezXxcGz84sA3F5DtYAWMys4DhknNLp28QznkqSHwoZ5m61v65qoQJBAOEscLcTenUgBZ2z3OpNtEeFEzOVJG7bA0kJgM6IhjtHg8Jb+fgKNuFxAO8wXZ4YKTx7p4ClKugOxvzCP/k3HrkCQQDZH9mz0aojATTI/gpxBokyopZytgS5aMefNfmTPnm9bFoiI1ChRjhDLOHwxqWX6gmSbnNSR77nkRgRJdme8Td/AkEAwteWjQRg0VqtIWISxfgJCF4BWIA0b2w6MofzmkOOi1r9iz/VVClahccnvNFIINXdUMXbEjlZoEWxL+PSQY7NmQJBAJ6+9sDOToKBY0KA2smAadcnoLAF/LZCsZDqOas6RnAERHIpN85yNLiInDkaRAAqEQ2Ky64g3qcYImyHK/FVk6kCQQCI8a2xzQVxrgknVuoZObhBwxweHR93ypPelEec9rw4OotqV8STs9Lkv8i5q8bQFUoNMpZro7zB+QMoeasUsVUZ",
		PublicKey:  "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQD2i184JOoIUI6hpx4LzMRo++svvHU7tRwZYq4IZJ8jAP1uswhIhhmPqdgVyBBA/AlvipLdzqZdiyqQTAVinncoOhmMY4sOJnyx2JmzHDJXTBRiQspmK1JuFFX7rWj1qVaw9R4HZrvhB10ORUjOJkc4O2f5/wu73gsNgcJZjuVFNQIDAQAB",
	})

	ret, err := charge.Handle(&OrderChargeConf{
		Money:         "0.01",
		SubMerchantNo: "IMCH20200324120454268314",
		OrderNo:       helper.CreateSn(),
		PayType:       "1",
		Remark:        "测试支付",
		TradeTime:     time.Now().Format("2006-01-02 15:04:05"),
		NotifyUrl:     "http://tq.udian.me/v1/common/enter-notify",
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v", ret)
	}
}

func TestCharge_Refund(t *testing.T) {
	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		MerchantNo: "DL20200323161220297431",
		Key:        "MIICeAIBADANBgkqhkiG9w0BAQEFAASCAmIwggJeAgEAAoGBAL76r2mrtxb0MwiYF+hnFSuS2roduVmLFvlczIlpsxOfXCI402uuEggUX5xNsreRi4qUw/lJd6zoW3/+kR+G1aoL7wJ7hUqSFv9K9CgoMVH9UR2HYoRH++yD2hHz/AsVgsdB0wDAOos4LzPmh0Oo5E+m6mgA2A2UffichfKUGvzHAgMBAAECgYBNqO8Pz24CfYcLJZ4DOXfYYj1jlZa7nN3YxS2/ayYRAqqal/URQpW+h1ph2w6jqyHNVrVid8ecnlgm8kPdSB019wxuQQR+rOiOCiwO4b+9CaEfqBaxTXK45+ezXxcGz84sA3F5DtYAWMys4DhknNLp28QznkqSHwoZ5m61v65qoQJBAOEscLcTenUgBZ2z3OpNtEeFEzOVJG7bA0kJgM6IhjtHg8Jb+fgKNuFxAO8wXZ4YKTx7p4ClKugOxvzCP/k3HrkCQQDZH9mz0aojATTI/gpxBokyopZytgS5aMefNfmTPnm9bFoiI1ChRjhDLOHwxqWX6gmSbnNSR77nkRgRJdme8Td/AkEAwteWjQRg0VqtIWISxfgJCF4BWIA0b2w6MofzmkOOi1r9iz/VVClahccnvNFIINXdUMXbEjlZoEWxL+PSQY7NmQJBAJ6+9sDOToKBY0KA2smAadcnoLAF/LZCsZDqOas6RnAERHIpN85yNLiInDkaRAAqEQ2Ky64g3qcYImyHK/FVk6kCQQCI8a2xzQVxrgknVuoZObhBwxweHR93ypPelEec9rw4OotqV8STs9Lkv8i5q8bQFUoNMpZro7zB+QMoeasUsVUZ",
		PublicKey:  "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQD2i184JOoIUI6hpx4LzMRo++svvHU7tRwZYq4IZJ8jAP1uswhIhhmPqdgVyBBA/AlvipLdzqZdiyqQTAVinncoOhmMY4sOJnyx2JmzHDJXTBRiQspmK1JuFFX7rWj1qVaw9R4HZrvhB10ORUjOJkc4O2f5/wu73gsNgcJZjuVFNQIDAQAB",
	})

	err := charge.Refund(&RefundConf{
		OrderNo:       "",
		RefundAmt:     "0.01",
		RefundOrderNo: helper.CreateSn(),
		NotifyUrl:     "http://tq.udian.me/v1/common/enter-notify",
	})

	fmt.Println(err)

}
