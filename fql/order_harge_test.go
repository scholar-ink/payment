package fql

import (
	"fmt"
	"github.com/scholar-ink/payment/helper"
	"testing"
)

func TestCharge_Handle(t *testing.T) {
	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		MerchId: "EW20200427175142",
		PriKey:  "MIIEowIBAAKCAQEAlFBoGPA7mOVPdSqOvWPm1GL8Ln2VUCSb5lfT1DrJys2DfZ3NNjLUJMXJdCtqFKlgyw7aOHCQVUsPOn8Qpgs0eyxjQGl1mknGIqQ2fAnGWFX99UNet4jmR7X11hTTvahCsifTTHaTeVzvIYZFoDRrxiJn7iyDYlZG7egiaBgpDpWP5tQJq64eRhMQbaAcK5U7qpq0pGEjAXLJHDf6pjer4EtKqawzoGX0JgJyPBJzT0+M5ulnfup1mVxEQtjqBSHVXBSAgYGJm1vHm3DP+kt32ia1nhh/hnT7+QEX1hry6XB5eURUgTX+W8zN5Fo6efdL2ovUA4CSn2vN8X9ahNcVdwIDAQABAoIBAGOxzNd+nEEBWzDiA4LxJVd8lhFWH0j44saqINzXC4/EJ3AH48pbzlhNj0YEbNEorcSw3iT0HUEILFtg0Dsc6xEk3C6O9RtaHdJpWap1E5uLaiM0PvXWExz/Bhn6c/5XnUWOGa2bQzRgMOnzDNhMhGlx9TSXPVWbsx/2WzJnkymWeOehvcsHFnRy9RgxHtBHsrtnWOoy18X9DqAbGozvq+AdC7M2n+MphJYz3MxQlyQaVQNfKq0X6Sd23V/MRyp09Qw+j1ksQrG7+mlkiXuyQQD1lNSmidYCu9ohvmbVMmTHT5+zaE64FHB9m8dIEC+zxHk+r3j9fDrL2ZAv2wzMNNECgYEA81yYp948IE4g1GyTmMKc6NFzMLWVVoOgh5oP/KR4sPyzf4n/tdlkgNPZ9NOTm78BUdJLEsfJfddtvM3T/7wqzDEDsUvUEuHNE46BuyXcZquO7q0TlZmWe/yrUOuWKyZ9p9Uz4yoq+lTXQB2/OELS9V/KGyZLr6zPoOqZM/HzH3kCgYEAnAQv/tblG9i1b0T8dP17wUuEmNaYC2j+ua1J07CHE9rpV8UvnT8EDirUzxDdzMepADKS9satFGwDaN53EfuVI0DiFAh2iwX4d0reql4brC6zLBhqt4/34M/4Ew0EJIaLAXASR23VDVtL7f3/pNww/9WwbgAhzSL+VharCzSD8G8CgYACHs1psQL/GQ8qKqQCkbeCQ8joJLzfWxg79jirb9dpW1pJAy/KKtG4WaAthOZlVf2x884ZpMn5obXgDj8JnF8XFlo84jc/zt3VjMlOkx7feiEtrS5dgFASVno7o0HUsGop2y6vjP+A3NpfQaMX32UtJB+Dqwq/xhatBT4+eJ4wqQKBgCkTGSe4n9zy8uF2ghiy27AKF8C9To5CeD2eF7+5y/Y4PhKzgVOe0YyBTkGtqRdOwXmh1vtNVyf3P/aOCz8awyri3CJ2D++OBNozRegImrR8HwPO8dmd956iYidx19Nd8J7SSJ2xltPLY3Qm2PUoW2YfdjGfq0Z8ngVANoveqZXVAoGBAJ6/ahIjr17biAm4oTPQMKL+9H4WZ87z41E/b2N2QzeCNTT4R4qlhyVx1vDRF7oDFZlzV3fj54iu9tQS6uD0X3F2rssO/nnvemz2eOCmZ337jn8D/9dJNF/R4ahpOaTdOAKqfw1nQ6TovsL7hK6kVzBOSqmX7DiwOLE6C6pITh8t",
	})

	ret, err := charge.Handle(&OrderChargeConf{
		OrderId:        helper.CreateSn(),
		ChUserId:       "d23fd8d2981149da8a82d2588ca672f0",
		CustomerId:     "f65f5254a9c14fab8ae836a19a28f429",
		ChannelType:    "WX",
		PmtType:        "59",
		OpenId:         "olWR55GrgjxID5fB5yF2S23oQVfE",
		AuthCode:       "olWR55GrgjxID5fB5yF2S23oQVfE",
		Amount:         "0.01",
		Subject:        "测试商品",
		Body:           helper.CreateSn() + "-李红艳-测试商品",
		SpbillCreateIp: "127.0.0.1",
		TxnType:        "26",
		NotifyUrl:      "https://web.udian.me/v1/payment/fql/notify",
	})

	fmt.Printf("%+v", ret)
	fmt.Println(err)
}
func TestCharge_Query(t *testing.T) {
	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		MerchId: "EW20190320162644",
		PriKey:  "MIIEowIBAAKCAQEAlMqn/KEVzhkdtccidoRCEz3BqXH7AMNt/FfjOJFKQuzwLYRfFP6TmkSVlsTSkIV/Az/lJaieJgKkltU4jkjpmJSC6YS1XDewk1dTmJp0SAnk9MGXpCuUqPAo/j6rz4/0iK4FqWTUxGsjOc0tpDc+tEvDMrzNfWlYybmbMIVH2GIlNkuusol+A7j8kn3t/1yONHmVWZKt7AByYLdd+rKYqYZPOBxXAttHP/9CTLoBHbjLwFCMRamEl/lBL34vQJiAZW81AOIi4CxJg10a3A/Z6tOcJYXSivGasnTvjuyyvqvxYL5OTrHlCT/xQ068xOzU3U8fjDjIbslR375DBl//NwIDAQABAoIBADgypuo7KVIzmE4dDX44DADadXf7bfNm3PbPdynZbnQCq+B1O7hhQvykZN+SLXmaglOG4ZSssDbpDqNNm1PaZChWB3ANyLYw7odoF1HvHHZNDmYHbK/8KeT4+HK21wvJcnHhUJAfXmFlmeNuBIwetZdBelOCjhaNIJTofp3/6RfnvNqt9cOaB/I/1XfVsMv7pdL+L1woZOUYZL5tT6G8nx2rQJVnjf5uCTGCoNwwCpJHiyNpvmk6QOBMInVPU0f76Sgrjz1o5NG0y3WHzgpUyCdl1z7bTahl1BG0qsUpQ4cFR28q744GSGy8D49sjPxf5hNjIQd5KsH/6jTLLQuE7BkCgYEA4lsHJh+rTMk40s+aqbAhojKx72JGMcGUAoLT59NJC6zD0bEajs1V/KS59H3FpoI+lzwDMGPUUmYTozi9N6sPmN8YuoaR0Q3tEdo4WUr5hxkezm10HisKeeC6h0D7NJlkh3/Gp2LScRZDOUAVLjt6T7X/ysAP6ou4uri6tNFaYXMCgYEAqEcoxT7W/7Jo/+OjylOsNk8TxgCmkJtt+sjnVtkBcAFrlRck/jn3rgp7ZmoT8QKmV1ws/V2CEIaBlwpVcZZuvLpcpgIi36sduky+ahILsnOSYM8DRZzqq7WPDFAC7CxS8L8dFVBxINQxiIQrhKkwTgmAaTJGu4nzGk36saSAKi0CgYBAYwhLjeKaOvrQ7IDgF9vZWXZH07qH2LqTZEeGwBEdIw2ojioxyLLW5LyIkWYxkQbg2g9GKn9w2NxpJ3CbbytGnt9X34OG5eEznNE+hRcpmLmsmnHXSwL13Djy1EcglSmFaZFGd9PImz6QAGyF9CE8n1adg7iDTS9E3dsuKAb/hQKBgGKK2Ts4o1q1NXuz6LSQ7yYWhLPMqb3A51SW1bIr/gWDL2btWMJvW7VVehHtSKQ4MwSxe78bRRE8UyMJ8CNGPq7SS3MDiTyFzjDMxC0FSEhGGZALahUX4OyQs6Y4LJ31DtTgdb2Hj9fzqtYQ4BMdqKXqNoJj6Lvl+rCKvaXNeSg5AoGBAMBBf/0/yc0lsUSvmwmTOUsSTeDlQIwOEcr6IrNLGI4GogXc96YOrpeEUZ9KwKmHaEu7P37An5ECw/sYwI9QVc2XPkOFwnaboNPBBAg7+vgufhWzNVekE3f7rut4hX3pRW+DoUxRiEmTnZoZRGo02W6Zph0xDOMoJl4ivkaDmgzd",
	})

	ret, err := charge.Query(&OrderQueryConf{
		OrderId:     "200624190818030101661603",
		OrglOrderId: "",
		ChUserId:    "UD20190320162646",
		CustomerId:  "d1953fca7baf4438ae07697a79bbadae",
	})

	fmt.Printf("%+v", ret)
	fmt.Println(err)
}
func TestOrderCharge_PaySplit(t *testing.T) {
	charge := new(OrderCharge)

	charge.InitBaseConfig(&BaseConfig{
		MerchId: "EW20200427175142",
		PriKey:  "MIIEowIBAAKCAQEAlFBoGPA7mOVPdSqOvWPm1GL8Ln2VUCSb5lfT1DrJys2DfZ3NNjLUJMXJdCtqFKlgyw7aOHCQVUsPOn8Qpgs0eyxjQGl1mknGIqQ2fAnGWFX99UNet4jmR7X11hTTvahCsifTTHaTeVzvIYZFoDRrxiJn7iyDYlZG7egiaBgpDpWP5tQJq64eRhMQbaAcK5U7qpq0pGEjAXLJHDf6pjer4EtKqawzoGX0JgJyPBJzT0+M5ulnfup1mVxEQtjqBSHVXBSAgYGJm1vHm3DP+kt32ia1nhh/hnT7+QEX1hry6XB5eURUgTX+W8zN5Fo6efdL2ovUA4CSn2vN8X9ahNcVdwIDAQABAoIBAGOxzNd+nEEBWzDiA4LxJVd8lhFWH0j44saqINzXC4/EJ3AH48pbzlhNj0YEbNEorcSw3iT0HUEILFtg0Dsc6xEk3C6O9RtaHdJpWap1E5uLaiM0PvXWExz/Bhn6c/5XnUWOGa2bQzRgMOnzDNhMhGlx9TSXPVWbsx/2WzJnkymWeOehvcsHFnRy9RgxHtBHsrtnWOoy18X9DqAbGozvq+AdC7M2n+MphJYz3MxQlyQaVQNfKq0X6Sd23V/MRyp09Qw+j1ksQrG7+mlkiXuyQQD1lNSmidYCu9ohvmbVMmTHT5+zaE64FHB9m8dIEC+zxHk+r3j9fDrL2ZAv2wzMNNECgYEA81yYp948IE4g1GyTmMKc6NFzMLWVVoOgh5oP/KR4sPyzf4n/tdlkgNPZ9NOTm78BUdJLEsfJfddtvM3T/7wqzDEDsUvUEuHNE46BuyXcZquO7q0TlZmWe/yrUOuWKyZ9p9Uz4yoq+lTXQB2/OELS9V/KGyZLr6zPoOqZM/HzH3kCgYEAnAQv/tblG9i1b0T8dP17wUuEmNaYC2j+ua1J07CHE9rpV8UvnT8EDirUzxDdzMepADKS9satFGwDaN53EfuVI0DiFAh2iwX4d0reql4brC6zLBhqt4/34M/4Ew0EJIaLAXASR23VDVtL7f3/pNww/9WwbgAhzSL+VharCzSD8G8CgYACHs1psQL/GQ8qKqQCkbeCQ8joJLzfWxg79jirb9dpW1pJAy/KKtG4WaAthOZlVf2x884ZpMn5obXgDj8JnF8XFlo84jc/zt3VjMlOkx7feiEtrS5dgFASVno7o0HUsGop2y6vjP+A3NpfQaMX32UtJB+Dqwq/xhatBT4+eJ4wqQKBgCkTGSe4n9zy8uF2ghiy27AKF8C9To5CeD2eF7+5y/Y4PhKzgVOe0YyBTkGtqRdOwXmh1vtNVyf3P/aOCz8awyri3CJ2D++OBNozRegImrR8HwPO8dmd956iYidx19Nd8J7SSJ2xltPLY3Qm2PUoW2YfdjGfq0Z8ngVANoveqZXVAoGBAJ6/ahIjr17biAm4oTPQMKL+9H4WZ87z41E/b2N2QzeCNTT4R4qlhyVx1vDRF7oDFZlzV3fj54iu9tQS6uD0X3F2rssO/nnvemz2eOCmZ337jn8D/9dJNF/R4ahpOaTdOAKqfw1nQ6TovsL7hK6kVzBOSqmX7DiwOLE6C6pITh8t",
	})

	ret, err := charge.PaySplit(&OrderPaySplitConf{
		OrderId:     "200611144545752643134294",
		OrglOrderId: "200611144545752643134293",
		BizType:     "01",
		CustomerId:  "f65f5254a9c14fab8ae836a19a28f429",
		List: []*OrderPaySplitItem{
			{
				Amount:      "0.13",
				RecvExAccNo: "0011562100000024429092301",
			},
			{
				Amount:      "0.71",
				RecvExAccNo: "0011562100000024429092303",
			},
		},
	})

	fmt.Printf("%+v", ret)
	fmt.Println(err)
}
