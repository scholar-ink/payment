package jttp

import (
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"testing"
	"time"
)

func TestCharge_Handle(t *testing.T) {
	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		NodeId:  "10090119",
		OrgId:   "1009300000000120",
		AesKey:  "lnfbmUz/rwUWYMYTl+CVRw==",
		SignKey: "MIIEowIBAAKCAQEA05fdu8rSS2REQ4jyomiuKUsmAPX25VxubE8A+4q0WdoLG4JgGtgQar+BD+2s12fwbR0A+Yiv3U45R43Xp7b2ZE5mAHemIMc9SoHkl90kS/XYOeWUUu9kuoPtUoIoIAvWdOltTmUNo8YkY6RhoR5Zh6YKAiRaIHiC4cOJGgxcBET9NEmmZ6HVkDw/LtaO8TymRJ2ctQSOf5kRdGnGlByK5FmE8m+6PpDZXSArVSX4V1gAHCDOdQ0+0bG7g/E8QBkCcIQIaueDSP9ckgJhBuaV6oFY7tA+BnyFB/oTPs6wgJIaaI4bcAfXnwUehfRu7K7DXTbFyGsNtlrPc1Z4yeJxUwIDAQABAoIBACKHqwjFbZIeF8HJrIbyfFGC0P3hJdyCPAby0Z87IHl1StC/cv2OewdgnyhBSi6Q6Lx1uj3+n2yHInPZ4vKcuhLedGjpHbRFr2vkTLC3xv+abCHupEjwhkil1OndEb1BEUtc+JiNhy9N24xjPAd2E3g+kW2ODyMzZ3B+X3Zfw/hyGcHDTXhT1jUPjp/kaS/bKgzAdfGtE3JTknEyYMwtA9nulbC0a9vo1jk5AS3qJF13nT+02Buxba5aFfCZ7TybYLKSGh0GFKLRiRPfb+tFR+AMWWG8A3hSgTrks8CN+T97drmXQA0WtnzK5XuQQS73TB7uvaKu1xft/cXpKhpSpxkCgYEA7Wc7BmcV/YAfdFaBnYxawA8ijgRNEjXNIWcJ3KYdFRIGX/qAjlIBAO2tKi2dNicXX49ltrX3oLmzeka8kccErYrwVNLJcEXNuPyrlo3QY7H3HKnrZsEO5XTNteGlMBbn/qb9ezy0p7ym2OzY3cOqGFRyP60fTfGhLrjUchE8aL8CgYEA5CsN1Fwkws9SfUGglIOuIIUg/e5444pUKlfMG0X3/c4+wnaURYESGrCfZZ4VZvnksbRX55jw3b70jbkzFwkpdc+DK0X2ihk8OKxs46MdQQF0+DdDvnDr2c176fXVzPl1G7T5mHX+V7PNf3D/HNkPIFEz3MtQHGJL6REWYkQ0KG0CgYEA4n0l9snKVPszBw6wwdnximHmGY9I6CKj/UYMRpHEFSIJKvAWTbE2e+hE2ISEA/Hvfp+T7mhUQuZnsDRkGr/AWsC+4G3o+E/gIKgOG6hYM4TJuHLzvihZSdfRmcAYlHSGCJDQLA6SW6TDmRj9HTVaxbNq/AzyXK41lBmJtPl0pIkCgYBAdPE5Nraj2hHBlNKaYgDE6xcA5Wd9UEyqkZb1dXSzXJpaMUl9wRKuO4ssF9aP+rRih0H2CTyySAmqJ9GJBmuR/oddqCwXoz1h/UPdouzWumSi4mne2OOn6ebBl1NIzogIxb1lFqA9gmvhPrizG5asWIRAMad1/UbYlp0uMXpSmQKBgC6Hu+BD1Qb/jkFL6foXRgtgtbiOkJQ1rjyociRdrQw/ExzSuzd9Jhi8iPZSee1x5kSpB459dwo8cUs5appVl767MKJPI5mabWZZ2euoUgzxPN6AVgzwNjBCXE2l6uJjebr492ZARbyW4RqRg6/cEOLzkNWfXMAk9QePIiQ9mdqq",
	})

	ret, err := charge.Handle(&OrderChargeConf{
		OutTradeNo:  helper.CreateSn(),
		TotalAmount: "0.01",
		Currency:    "CNY",
		Body:        "投诉QQ:3007638620--订单号:LDED13DBA226957282",
		NotifyUrl:   "http://tq.udian.me/v1/common/enter-notify",
		OrgCreateIp: "127.0.0.1",
		OrderTime:   time.Now().Format("20060102150405"),
	})

	fmt.Printf("%+v", ret)
	fmt.Println(err)
}
