package notify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestZgEnterNotify_Handle(t *testing.T) {

	notify := new(ZgEnterNotify)

	pfxData, _ := ioutil.ReadFile("taoqi.pfx")

	notify.InitBaseConfig(&ZgConfig{
		AgentNo:      "b1537337771064377",
		Key:          "GIX4BRVHFT61KBVU",
		PfxData:      pfxData,
		CertPassWord: "Udian888",
	})

	//ret := `{"agentNo":"b1537337771064377","charset":"UTF-8","encryptData":"fg8uj0Kk6c2titvQXWotwCkc2GzgWIs1PPptk5YI1OjG68u46QE0mL/Nb2NARYOIVf/5SukE1aax2ggFkjvfp5YBmORDbe2MzOI68fhYBAwjMIFaIl/9PL9XcOAhayQGt4QN7Sfsg/cZez4lhT6rNeV8vmQlV+LZNfPSoMHtHOOzkyibBpKb+rs9y1YAfw4N/t7bATthIEJb4fN/f4qBx/Aqc7ektscsH2JHV0eWWksM929Sax7OpHmxW6JTR8E/4QfOY1QJ2ponnouzcaCcs4rsYdhUwfj9Iqjcje1UI/ly8d5tc6B5jx71lVS84guS5ZnBW3ozzkhyXi3ysa5MHI25AngK7qD/SiJq/kLkbxiunrIQ3ud/fM/cCtS7RsTH1bR1dik0yYeJIyoDyXxWYylHGnEvYIBKrvnwo7kJOZh3Nmctt4bPSHGbpB8zBD/gASk18nrN1tQN45I+utBGoBfI9EUiJLtQfw8WY415t5Ofra0Ln9z/3n7pkW8JnuFIEEFo/CS+6VXPPQd/JuYjJaYHvITNMfX7Q3XgBzNg9y23637gxWTDKNZbXIXfG7yfYAIOHOZWmiE3sBBzD9+odK+hCGX7t4fatyRY4EjStg1WDy7dfWjg888kWaX9+lSDP5fNPuIfvx1vKOJSxc1e0ItMSI8omw/NYXmpQGwWsYM=","encryptType":"RSA","responseCode":"0000","responseMsg":"接收成功","serviceName":"merchant.enter.create","signData":"83ec9dcd5475b0f8c753d63813a9f90c","version":"3.0"}`
	//ret := `{"agentNo":"b1537337771064377","charset":"UTF-8","encryptData":"SmpUqICLqtgHPDgF3MB8+hrnloYo4RrM/6MOzxh9hk7FoumSUupB1Ykq5uICjlYcTaHL4kXFCes4nwGaq/qNoDQS+PPmUH7uALvCx0zmNCirWR3RK9F4Vp3NuylWbMO5w1fIXaDhPxhrwgSyk4Nf6y/d0ntWhqbVy7ZtXyPEHDb90lJsU//4A/jXcanlBAMoufYwTwvrT0hW8ryXvBN61cCH9lhPszGpPzy1U0KUAYQXfQiU1jtnCpYb4jPbLVTgDWHjC6olIMu9KJ4d4XKbeNqqH3azgxTDWFVAts7cu/N4X0TvaCVraM8kj6k2c30IcX7+DpTWEiFsVqQnYcAR5AbAK4E6AYOyoJ4x2tKyHry65DmlV74KOGKJN/JtRckvulyXHY+VnZr6TsTtRtpnD8mPSyPAf+20SDjuZuZwoErKVhXTz9zWjpG3GqlNP8g9+3q24JHqARjKdwzuW6R4jIzLggKizVICxvuN32n3z+Ig0zjzbhO54Hwv60v8qFc5XMexgsDnvmUJQl3FKgb9BpGGsBvzqBXvhelX+zvUj95PJfgOkEqTfqvOUxsxsQj7H7uShMBA0GA+CnrUAdMAsH/wzI8b+V5I3WRtatefqFonhPWvV3mb7NrDkUni9sKBvqb90QQp1VHU1md8LlgKI6JC8xk6MUJoj4EcAT/OLYXPqn1wUKU85V5sRiLG9qYPfaiO9mUvjB8WcCuZigS8YPqCbpYPT/fwqDIIlz3ZtNjzO2S1DwyujTeW4wLCNvhchlmFviSLxGkwb4wHp3QAsE0dU4jq3o1pu7AF611AA7Xjd2zllUFbOT8u8zdMNy+HSt59ctplUWCac1ookLWlmb9I+JNSZiPf0Ome5MniP2m+zzsbWN3St6Idx2yAIa7sVflZv2S85iO38iGXl0eF0UQB7kN5yJ8/DrhUoQNhH9se02+kUXHkIuOZCHM22ihLXcSi68Xc9PakIGOwrzjWCU8VF3JBYf7vWG96EwD4FiFJExCRHDp7sJ6Tr9BCwEEc","encryptType":"RSA","responseCode":"0000","responseMsg":"接收成功","serviceName":"merchant.sm.products.notify","signData":"f90647dc451e68a0d1833f56db09038d","version":"3.0"}`
	//ret := `{"agentNo":"b1537337771064377","charset":"UTF-8","encryptData":"v8KKlTj10FLcwvtz5USxjMYKKsIy47vJQ89aSlC5YlLJkdPjx5Re3lU1K5SH6O0X2v2MwsJLfLmn98RKp+ie+yw8vPOG34AFe2jJ5+pHvfNFhhlVZBONQgJPNnpv7N9se/JCD7FMu5Myqm1Cq7IszYMcxsrgMImtzTw1srqkML7qOP62AKUoNi3yg6qPj4F7W77G2ZX2llNycDHeESS8mdk82ThiJGVAj29sdQBMiUb5Fi37fVFPsRG4VWa9bSIXm+Ky8FcTdQEz7Ze4sxuooI4MxB+0Pw+sBCQNY4ooUKEou6Y1SSuYEQjQTv7jLCtBo0JSoQTEHOj6DEb++850PXzF4z75fOpg4Zwt93fiBiymaHJiyelHUFB/WEruKECRFz74+XowJ4tyYrY628uVCZkx5Lvu7dspwP9lWdoCgZ9TaYGZLLDWCfkG4c6arYmEBvEPhJNkM0xlQ/PnZMmOe9TUvjOfCXKP3lfD8WjT9+VDEciISDr1kgJEzshALfhDQEkfWcQf3ORs1EOhtU7pOII9qDCT9ocucmId+Eb5Xde3RU7nZO68V4enlrYVah6XJt5tXFVv1UJm0RcEZ+lFI8SQrHln/DffwuWur0ztr2o/OGGB9tiq1yLOtLqnShrn/EQ8Zeaa9gpddsTOxb9t8BWvCIDYezxtculmcm8Yvt4=","encryptType":"RSA","responseCode":"0000","responseMsg":"接收成功","serviceName":"merchant.sm.products.notify","signData":"b7b48e2245991a466e5ae47e763c99fe","version":"3.0"}`
	//ret := `{"agentNo":"b1537337771064377","charset":"UTF-8","encryptData":"eP4/QZnAx53km9anpGeFvgMR4uNuirMI1SNkpOW2tO/9p6+1w7SXIL9KKScJx8UQ2eTPRmmrXyfV/KKzUI2E16SN7NSgcgvj4++fZbZg9+MfuJlH/2sU5c5NNHGOnv8wB8UejcAVC0a5PznVZXre/HZeErNOqgX3xl/TpvvWdFnUt7fHVAGEKELalmMSf18xJku8QrPxYeShKB7vMegMZQbLeCfOHGfziyhxIXAcsTnfKovlqIKTJPjySA9peahf28Al0VqPD4yHZ+I6X+27JGJa+eyoCZu/ozROyDhn/3xD55J6Zuu26l2VBjKGW3DV5KO2Qo/R8nOIqDsJF0/aVcrHwL2XqZiybvJLEaT0HSHyOSHFeEZtxdd2JHWsHwtJh4PBQB9DHQZpaqsPRjt381s9b7MWrdhnb2mOpux17e8749q3AZSJrOSx0xGmGAq8Hmeuq6snvPm2381GdXNYjde3piH5ffWTGYC8XygIe4yI35B3pqBb5UqatZyobY+qkmrzyZTdmNbLbXhtAIF63S+g/eMrJeu0EH17HaPlukADO8jn0ZlbkTmhicsM74fFqI/8w4mPlapSszv/UCcXRASkVl3ceoOl5vXd5JYJxpcgUIuoXC/cKCesmlZWbu0oDiXCyCEjqg9Psc06gRmj2rA11q9XUtoA3D7ERLGJCfg/bFiBEosxqGOu+qMDXYIItT4kZ2bPOFtdDy8Ilp0f8UhXI7FVrecYDOipxtMC8kfqW5D5miHvQIOWsYQbhVUeEppelotiwDl5ttm7p1ZktBeda0rui8eMOyCfvGbHx9ZGaC88wiI3xueOX1JeVrE5mJcki8ZPOVSdv+h+5YKyOzziThzl0jCpRsKqxPr5bSbD/lEYQ5tmbQX4ACohpaPQPVVRGa/E/4YKRITm5eCkBTssv4YE2yd5rYBDMis794pSKt/AbRuQEQivG6jy99d3MNvQGjCTUnznolr7344myxsQwT0+rB8wphIq/ZX6rgiqg7X4oCcVok3E56P91kRb","encryptType":"RSA","responseCode":"0000","responseMsg":"接收成功","serviceName":"merchant.sm.products.notify","signData":"89972fb39820abb9aaf4e678ef5186e1","version":"3.0"}`
	//ret := `{"agentNo":"b1537337771064377","charset":"UTF-8","encryptData":"l8ZyKa1Js9XwjaYEEc7TnqtRYdcEsAwOXQOqFX4qFLActz0wOA0tO36bTjoxMiA8Ae0ECbivvsmPLOZO/nxjMsvf7Gns5OcUT4t84I8LjHCdnySAGFaTXvQaLlWwqtkMMAr11TbK65vkX6HITmxQrQOAKdSTwTyBfJF5W7XgFLeoYGogBtUJ1XiJ9dfLYXtZlTkpUo1THV0emZeDEc6dK5lahoqoViEWwEnwfLNXvNjTGBf5iqAHOT4CxEZqErmyqJqbaui2CVn3r/U+jqQwYrkC1NhSCuyMLnkIUTSgNxNtutlBnbz03u+P9Gp4I2yaaWKPA0wlN5+OXpTc0cvbPlXykGKtzQv6lHNUXGPuweqdimqA/nYfX1qAuA1OTBIr44itcCIAP8g1Kug9vcF1vXZ7OE+sClrLN0uPPmip4D+iiaMxiKEHudsdsWwWTfENi/7kEcb+C+upQBiqCRPrG520exy5PMA3D091cD2d4wwNrjvjFV5I1Bgp0g+xUel/nfzmO62NzDrrOvk+9bLtdBUadhx8vplA9M3rvXg8vEUnhxZUAIM62Zsg0w5WlFEPyD2Vh8rezps3SlkJj5J9w+7jvcIkxipb+6ijQ+ejzO5yIYZgGtYWFOQyWUv+up4QRxz6yeYeVw0Jbcjf1eIdB/4Y3FK8l5e/umDhZ56KRAs=","encryptType":"RSA","responseCode":"0000","responseMsg":"接收成功","serviceName":"merchant.sm.products.notify","signData":"643ceadb8875f12db0f6e4bde68514bf","version":"3.0"}`
	//ret := `{"agentNo":"b1537337771064377","charset":"UTF-8","encryptData":"sHCcavmMJnKnBo5HUJ/vbhAaAVePrj1++e4iqkcBIK0fMytn0r6HRfeDeuPTSBsF1wkK34u8RAb7TCYghKpOchtt2k1ZGYMCYV+WEWU6TI1GugxV8FN1n0T7jKdU3mnr0Mw2uw2OBYI04gXU9of/AqmZ5fW8b+xk5+YJ+Ni7o/EfxVsSpxdfwX/62cu18iHcY5UyCpR6ver7haXPctiboXuGU3SqIgy2I1+FOXHA++/Fqi5MeKIrJnt3PQL0rKVUZ6kfqu1KuZj5FDnUXfZTfmSEbLC0TWiTGUPiCysGbLUMHebeWdaQhfcwFxUr6Efxkm3G2PVPaCrwrc+mwHc4NGcjmYBxlSTAKiHLfUPdyHQaSGpk+Cv0D87YhGsTi0F+jv0TU4wvV9mj4kqjxulQkohzz3AwR64rdqdtNoiJXWk8lpN2U1QacZgMAQDgIabxLnhKjmlhJufS1AWMKWikASmy9Hnx+L9TOlOpG3RXw3cRm2jAa63uasZTJQ5eUJpPxVRhKVhAJYR9/0sXHL8if7fK2Qb3a8GBVsrXMzuOFxwsq04OGkFyAqFNeAXWyiUUgIkYMlE/Xtkuo1C/q69bid36M7ASg2rnK+ztEL0MhuaGr13OwpOVsMUvogmYEoyLj0UIsbGZFp0dvraPYFD7tohO0gAaqCKpWoroQkbzBAU=","encryptType":"RSA","responseCode":"0000","responseMsg":"接收成功","serviceName":"merchant.sm.products.notify","signData":"a55183006f8224ccb976c31c8bcde428","version":"3.0"}`
	//ret := `{"agentNo":"b1537337771064377","charset":"UTF-8","encryptData":"kZG8rZWrJnDLDH80d+LtbSdgBf2uI+IQlrvUKB9P7mtNsmc+taInadHzdOw0JwdNTGDnbWGC0SA8zxlCNQl//UMCRhgKaz98BI6FU/kC0HLM2uOAN9ypJTczRSnPlNMUwHBsXHYONb9AVz1hOE491nwqplF9CIy5REIX18aCeeOCQjcarrWa3IRVX1wKbw2/9fuepc2I4wq+V8pOTM3DPnBij1D741mAcITMjj6C6VVQ8wJ8LLMlqdaYWDZGY0+eRdbGdATaB3YO1LBdJqfYWC6UCVjoMB5lJBa3Ve69+rTredIwarig4ffkdrcZN03zX1WqKXq75Bg2Qi1rJerQJWs+F3Jc8Abpuuz+Cyu/aszrarM8l3BKapYEBYFHd5woaEv9Dnd0vsXQ/ExNI0cvLp2QedC2RQ3XI6m1+z8wV1Pg5dnqQdtRoSSTLmWT+G1PVL+JLSZ7c/jjp/CERiO31IMdoNzmaVTVguL5beN3wAxEZhN08y8rLZhkbd4lneQrfHiVGczVvuxj+dvp+SCrErmPcx5r+QzS+eh8h1JVrT81yywUpv1h4ocsMs7QRbBg+RUuuJON9siBnDQG5zmFQYz1AKlvdCKM+8IbdYyV8UKrcOdNGGQSrsJ7aLyvW4byuLU654Q/JGj/iw8+r2PAaaQ7pH4hNo5KA3SFdIlUo+A=","encryptType":"RSA","responseCode":"0000","responseMsg":"接收成功","serviceName":"merchant.sm.products.notify","signData":"28f97e27dcec8f4c7d72e81a05d0ee7c","version":"3.0"}`
	//ret := `{"agentNo":"b1537337771064377","charset":"UTF-8","encryptData":"kZG8rZWrJnDLDH80d+LtbSdgBf2uI+IQlrvUKB9P7mtNsmc+taInadHzdOw0JwdNTGDnbWGC0SA8zxlCNQl//UMCRhgKaz98BI6FU/kC0HLM2uOAN9ypJTczRSnPlNMUwHBsXHYONb9AVz1hOE491nwqplF9CIy5REIX18aCeeOCQjcarrWa3IRVX1wKbw2/9fuepc2I4wq+V8pOTM3DPnBij1D741mAcITMjj6C6VVQ8wJ8LLMlqdaYWDZGY0+eRdbGdATaB3YO1LBdJqfYWC6UCVjoMB5lJBa3Ve69+rTredIwarig4ffkdrcZN03zX1WqKXq75Bg2Qi1rJerQJWs+F3Jc8Abpuuz+Cyu/aszrarM8l3BKapYEBYFHd5woaEv9Dnd0vsXQ/ExNI0cvLp2QedC2RQ3XI6m1+z8wV1Pg5dnqQdtRoSSTLmWT+G1PVL+JLSZ7c/jjp/CERiO31IMdoNzmaVTVguL5beN3wAxEZhN08y8rLZhkbd4lneQrfHiVGczVvuxj+dvp+SCrErmPcx5r+QzS+eh8h1JVrT81yywUpv1h4ocsMs7QRbBg+RUuuJON9siBnDQG5zmFQYz1AKlvdCKM+8IbdYyV8UKrcOdNGGQSrsJ7aLyvW4byuLU654Q/JGj/iw8+r2PAaaQ7pH4hNo5KA3SFdIlUo+A=","encryptType":"RSA","responseCode":"0000","responseMsg":"接收成功","serviceName":"merchant.sm.products.notify","signData":"28f97e27dcec8f4c7d72e81a05d0ee7c","version":"3.0"}`
	ret := `{"agentNo":"b1537337771064377","charset":"UTF-8","encryptData":"xmm4Y5u245lF6ScOrRAWa2MN+Ol2TyPcVfE+JVTKRhMDZDVOKSB3epRkGor28VP8OQCRWhvdkKfTRUZeFE2f2XeT60ICBmO8S5mEdCITBc8dRqIBVoCsZUqi4nuO2RZyKuoF1BUa8pJuHyRfc606rXauIxnktbyNrdvMqz3U5pGC4BIPvhBq3db1ySnYZsBIioykEHzBuXcV5gx3bxp+oqVgiFeOyC3NDlDkpT7gjWgo3vbFHQgvA6lVKFyY79SQmBB9tyfLG0o8OuHTkDTQygxApiDUcW3oorh2AzBFv3crqCh1M/vRkgKA/CgHI3nNxX+N0VQCvHxySS/YBNAUe7pTvuIGXUycm9oskwqPfGsLH77jBdwS1o4xSS6Arn2wBBfk3FJLM2HkZwGOHRHbiSUA3ooOdUlN2YuBf6x/MFFWtbnW8Kks6MX2jFEWT9Jyr9LGESoW74iEg1MZMfm4r3+euJxL3JWWXfe3ivCg0BMEIk3JTXgBTNibg4YejPXVtsIswvy+xW44rXletuWP+9v03ORyFkpLx7sfC4oN3oZDhL1i3owmH0A0c4JboVmh51XhC0KNNINWHabjdu/OFv6fu81fJWndQp3nD4x0U1yVV+ts0rswCdIv3VclKPI6gQksLuBHAZ801x26sA+pZK/BhjcmwhtnAffK8yAeNFE=","encryptType":"RSA","responseCode":"0000","responseMsg":"接收成功","serviceName":"merchant.enter.create","signData":"b03136336cfb0e954dd1e5d41ccb39df","version":"3.0"}`
	//xml := "<xml><appid><![CDATA[wxf06ac118ca3d9533]]></appid><bank_type><![CDATA[CFT]]></bank_type><cash_fee><![CDATA[1]]></cash_fee><fee_type><![CDATA[CNY]]></fee_type><is_subscribe><![CDATA[N]]></is_subscribe><mch_id><![CDATA[1495589652]]></mch_id><nonce_str><![CDATA[a5ca845515b55b1164cfb3fe0095e943]]></nonce_str><openid><![CDATA[oJb7cwYJW-YC6ynUtvVLFm9UXfgs]]></openid><out_trade_no><![CDATA[2018060615095472882]]></out_trade_no><result_code><![CDATA[SUCCESS]]></result_code><return_code><![CDATA[SUCCESS]]></return_code><sign><![CDATA[A71D51B60BCAAEF4285977F4E4C629DA]]></sign><sub_appid><![CDATA[wxa33cba2b69f869f3]]></sub_appid><sub_is_subscribe><![CDATA[N]]></sub_is_subscribe><sub_mch_id><![CDATA[1495746312]]></sub_mch_id><sub_openid><![CDATA[oyA310LEnY_JW_-BDHVJguSpFyKQ]]></sub_openid><time_end><![CDATA[20180608165743]]></time_end><total_fee>1</total_fee><trade_type><![CDATA[JSAPI]]></trade_type><transaction_id><![CDATA[4200000123201806080945667886]]></transaction_id></xml>"

	result := notify.Handle(ret, func(data *ZgEnterNotifyData) error {
		b, err := json.Marshal(data)

		fmt.Println(string(b))
		fmt.Println(err)
		return nil
	})

	fmt.Println(result)
}
