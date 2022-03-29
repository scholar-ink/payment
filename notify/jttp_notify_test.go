package notify

import (
	"fmt"
	"testing"
)

func TestJttpNotify_Handle(t *testing.T) {

	notify := new(JttpNotify)

	ret := map[string]interface{}{
		"bizContext": "ppdXNXClFO2sCQghxzCyaYKdhnd6Q+2T00+hmH7C7CBFz7Z2hIJs+CpYWfj5ODKYnmRn3vZcBlxQ6J6EtBwlX8NoOtyFOyrfp6fQYu7N1oQdGHhUUh+cnYdsDVhdgoc8z4CKjgwCC7cKKxRcKTvjJKbcaIop988AQqT57C0aylGysToUkotz3OjQE2UFCNtNaC/IHKpFou64OHac7ystZA==",
		"charset":    "UTF-8",
		"nodeId":     "10090119",
		"orderTime":  "20200218150309",
		"orgId":      "1009300000000120",
		"reserve1":   "200218150309861701056739",
		"sign":       "djpxSPKY0jp1/Ag/3pfgQ5VOLJ+ZTdr7OFHebyNI4MBzt0pZ8O2SrTL3NouqNDL4dR3LLJJG94M2eR2nK2K19R++yMyTXBcRo2WtXG1/cNfpbHKr/WFxV+mcJuXGUrlE7eqYCPVCqFvZL2B8hQV0dYHiXpJlXeXaRvzSo1SZ0Dpsi4Ou4T9iuyIqNZ994Tmbg4CjA+mu7xlCKe1fOrYrfZ0DDgG91uadXJYlPXIKWOPG9Kxr0uUdE/OS5INSc1u/c3eS98/87pBSdzzYAKVnC4c1XE73z27P2bjDr46mAJgQGGULRqhvRuwkESbMV+osQSNdik34hJQNwAWz/AB+EA==",
		"signType":   "RSA2",
		"txnType":    "T20601",
	}

	retData := notify.Handle(ret, func(merchantNo string) (md5Key, aesKey string, err error) {
		return "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAps9QQZRfsRHfg+diBaJcfKcSPhtXZ2I9OG/oDcCRr0ZL+9V/yAXXfAUlnXd+gt9PD9L6yraAFxVxnbWOI/K5RlzpUY2tSEB5J8di+s/nXIq296LMoLTDCJ2G5Du5rVQnexoRRUpiJVBcgEgeX144Z8sb3uoqWGAHzf9D0uj5EEh90ei+XLrKk09z1wFE9qZf0U8SAzIXJnKB0sB35gdsoZlUz9+WyrZBR+mG86E8aWVj4OvI7HGnTyC4gegozqxLbYIE2X/D4eAPShpahElc/glMfaKZh2SL3r7lovqJEIkvPX1dFAlMUQBsq3vk3qJ+vy73h19dnURiN4dWWyhgkwIDAQAB", "lnfbmUz/rwUWYMYTl+CVRw==", nil
	}, func(data *JttpNotifyData) error {
		fmt.Printf("%+v\n", data)
		return nil
	})

	fmt.Println(retData)

}
