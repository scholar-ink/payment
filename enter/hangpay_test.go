package enter

import (
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"io/ioutil"
	"testing"
)

func TestHangPayEnter_Create(t *testing.T) {
	enter := new(HangPayEnter)
	enter.InitConfig(&HangPayConfig{
		MerchantNo: "DL20200323161220297431",
		Key:        "MIICeAIBADANBgkqhkiG9w0BAQEFAASCAmIwggJeAgEAAoGBAL76r2mrtxb0MwiYF+hnFSuS2roduVmLFvlczIlpsxOfXCI402uuEggUX5xNsreRi4qUw/lJd6zoW3/+kR+G1aoL7wJ7hUqSFv9K9CgoMVH9UR2HYoRH++yD2hHz/AsVgsdB0wDAOos4LzPmh0Oo5E+m6mgA2A2UffichfKUGvzHAgMBAAECgYBNqO8Pz24CfYcLJZ4DOXfYYj1jlZa7nN3YxS2/ayYRAqqal/URQpW+h1ph2w6jqyHNVrVid8ecnlgm8kPdSB019wxuQQR+rOiOCiwO4b+9CaEfqBaxTXK45+ezXxcGz84sA3F5DtYAWMys4DhknNLp28QznkqSHwoZ5m61v65qoQJBAOEscLcTenUgBZ2z3OpNtEeFEzOVJG7bA0kJgM6IhjtHg8Jb+fgKNuFxAO8wXZ4YKTx7p4ClKugOxvzCP/k3HrkCQQDZH9mz0aojATTI/gpxBokyopZytgS5aMefNfmTPnm9bFoiI1ChRjhDLOHwxqWX6gmSbnNSR77nkRgRJdme8Td/AkEAwteWjQRg0VqtIWISxfgJCF4BWIA0b2w6MofzmkOOi1r9iz/VVClahccnvNFIINXdUMXbEjlZoEWxL+PSQY7NmQJBAJ6+9sDOToKBY0KA2smAadcnoLAF/LZCsZDqOas6RnAERHIpN85yNLiInDkaRAAqEQ2Ky64g3qcYImyHK/FVk6kCQQCI8a2xzQVxrgknVuoZObhBwxweHR93ypPelEec9rw4OotqV8STs9Lkv8i5q8bQFUoNMpZro7zB+QMoeasUsVUZ",
		PublicKey:  "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQD2i184JOoIUI6hpx4LzMRo++svvHU7tRwZYq4IZJ8jAP1uswhIhhmPqdgVyBBA/AlvipLdzqZdiyqQTAVinncoOhmMY4sOJnyx2JmzHDJXTBRiQspmK1JuFFX7rWj1qVaw9R4HZrvhB10ORUjOJkc4O2f5/wu73gsNgcJZjuVFNQIDAQAB",
	})

	b, _ := ioutil.ReadFile("1.zip")

	ret, err := enter.Create(&HangPayCreateConf{
		NetworkNo:         helper.CreateSn(),
		NotifyUrl:         "http://web.udian.me/v1/common/enter-notify",
		MchLevel:          "2",
		MchType:           "6",
		MchName:           "赤壁市考升百货店",
		BusinessLicense:   "92421281MA4DLPXRXB",
		LegalPersonName:   "李红艳",
		LegalPersonIdCard: "320830199012082427",
		BusinessModel:     "2",
		ContactName:       "李红艳",
		ContactTel:        "17714500631",
		Province:          "湖北省",
		City:              "咸宁市",
		District:          "赤壁市",
		MchAddress:        "马港创新聚集区电商产业园A区A0737",
		BankCard: &HangBank{
			OpenName:    "李红艳",
			AccountType: "2",
			OpenBank:    "上海浦东发展银行",
			CardNo:      "6217920471658935",
			OpenBranch:  "上海浦东发展银行股份有限公司淮安分行",
		},
		AliPayRate:   "0.007",
		WxChatRate:   "0.007",
		UnionPayRate: "0.007",
		WithdrawRate: "1.00",
		ZipFile:      b,
	})
	fmt.Println(err)
	fmt.Printf("%+v", ret)
}

func TestHangPayEnter_Balance(t *testing.T) {
	enter := new(HangPayEnter)
	enter.InitConfig(&HangPayConfig{
		MerchantNo: "DL20200323161220297431",
		Key:        "MIICeAIBADANBgkqhkiG9w0BAQEFAASCAmIwggJeAgEAAoGBAL76r2mrtxb0MwiYF+hnFSuS2roduVmLFvlczIlpsxOfXCI402uuEggUX5xNsreRi4qUw/lJd6zoW3/+kR+G1aoL7wJ7hUqSFv9K9CgoMVH9UR2HYoRH++yD2hHz/AsVgsdB0wDAOos4LzPmh0Oo5E+m6mgA2A2UffichfKUGvzHAgMBAAECgYBNqO8Pz24CfYcLJZ4DOXfYYj1jlZa7nN3YxS2/ayYRAqqal/URQpW+h1ph2w6jqyHNVrVid8ecnlgm8kPdSB019wxuQQR+rOiOCiwO4b+9CaEfqBaxTXK45+ezXxcGz84sA3F5DtYAWMys4DhknNLp28QznkqSHwoZ5m61v65qoQJBAOEscLcTenUgBZ2z3OpNtEeFEzOVJG7bA0kJgM6IhjtHg8Jb+fgKNuFxAO8wXZ4YKTx7p4ClKugOxvzCP/k3HrkCQQDZH9mz0aojATTI/gpxBokyopZytgS5aMefNfmTPnm9bFoiI1ChRjhDLOHwxqWX6gmSbnNSR77nkRgRJdme8Td/AkEAwteWjQRg0VqtIWISxfgJCF4BWIA0b2w6MofzmkOOi1r9iz/VVClahccnvNFIINXdUMXbEjlZoEWxL+PSQY7NmQJBAJ6+9sDOToKBY0KA2smAadcnoLAF/LZCsZDqOas6RnAERHIpN85yNLiInDkaRAAqEQ2Ky64g3qcYImyHK/FVk6kCQQCI8a2xzQVxrgknVuoZObhBwxweHR93ypPelEec9rw4OotqV8STs9Lkv8i5q8bQFUoNMpZro7zB+QMoeasUsVUZ",
		PublicKey:  "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQD2i184JOoIUI6hpx4LzMRo++svvHU7tRwZYq4IZJ8jAP1uswhIhhmPqdgVyBBA/AlvipLdzqZdiyqQTAVinncoOhmMY4sOJnyx2JmzHDJXTBRiQspmK1JuFFX7rWj1qVaw9R4HZrvhB10ORUjOJkc4O2f5/wu73gsNgcJZjuVFNQIDAQAB",
	})

	ret, err := enter.Balance(&HangPayBalanceConf{
		SubMerchantNo: "IMCH20200324120454268314",
	})
	fmt.Println(err)
	fmt.Printf("%+v", ret)
}

func TestHangPayEnter_Withdraw(t *testing.T) {
	enter := new(HangPayEnter)
	enter.InitConfig(&HangPayConfig{
		MerchantNo: "DL20200323161220297431",
		Key:        "MIICeAIBADANBgkqhkiG9w0BAQEFAASCAmIwggJeAgEAAoGBAL76r2mrtxb0MwiYF+hnFSuS2roduVmLFvlczIlpsxOfXCI402uuEggUX5xNsreRi4qUw/lJd6zoW3/+kR+G1aoL7wJ7hUqSFv9K9CgoMVH9UR2HYoRH++yD2hHz/AsVgsdB0wDAOos4LzPmh0Oo5E+m6mgA2A2UffichfKUGvzHAgMBAAECgYBNqO8Pz24CfYcLJZ4DOXfYYj1jlZa7nN3YxS2/ayYRAqqal/URQpW+h1ph2w6jqyHNVrVid8ecnlgm8kPdSB019wxuQQR+rOiOCiwO4b+9CaEfqBaxTXK45+ezXxcGz84sA3F5DtYAWMys4DhknNLp28QznkqSHwoZ5m61v65qoQJBAOEscLcTenUgBZ2z3OpNtEeFEzOVJG7bA0kJgM6IhjtHg8Jb+fgKNuFxAO8wXZ4YKTx7p4ClKugOxvzCP/k3HrkCQQDZH9mz0aojATTI/gpxBokyopZytgS5aMefNfmTPnm9bFoiI1ChRjhDLOHwxqWX6gmSbnNSR77nkRgRJdme8Td/AkEAwteWjQRg0VqtIWISxfgJCF4BWIA0b2w6MofzmkOOi1r9iz/VVClahccnvNFIINXdUMXbEjlZoEWxL+PSQY7NmQJBAJ6+9sDOToKBY0KA2smAadcnoLAF/LZCsZDqOas6RnAERHIpN85yNLiInDkaRAAqEQ2Ky64g3qcYImyHK/FVk6kCQQCI8a2xzQVxrgknVuoZObhBwxweHR93ypPelEec9rw4OotqV8STs9Lkv8i5q8bQFUoNMpZro7zB+QMoeasUsVUZ",
		PublicKey:  "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQD2i184JOoIUI6hpx4LzMRo++svvHU7tRwZYq4IZJ8jAP1uswhIhhmPqdgVyBBA/AlvipLdzqZdiyqQTAVinncoOhmMY4sOJnyx2JmzHDJXTBRiQspmK1JuFFX7rWj1qVaw9R4HZrvhB10ORUjOJkc4O2f5/wu73gsNgcJZjuVFNQIDAQAB",
	})

	ret, err := enter.Withdraw(&HangPayWithdrawConf{
		SubMerchantNo: "IMCH20200324120454268314",
		Money:         "144181.90",
		NotifyUrl:     "http://web.udian.me/v1/common/enter-notify",
		WithdrawNo:    helper.CreateSn(),
	})
	fmt.Println(err)
	fmt.Printf("%+v", ret)
}
