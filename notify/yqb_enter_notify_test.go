package notify

import (
	"fmt"
	"testing"
)

func TestYqbEnterNotify_Handle(t *testing.T) {

	notify := new(YqbEnterNotify)

	notify.InitBaseConfig(&YqbEnterConfig{
		AesKey: "s9DZBvxB2omgBeo0R6rVwg==",
	})

	//ret := `{"content":"VZyQV7FOu66+O4o3H2JJrSlCNIs1AZqV/GH1jMghgVLinMWql+zzWOWmoXstYGWjkE9FNj7dtG/T\n4H64JetYcQzvmhrgw05qKe6W61ABJx4wGAUlPL8w/l5UB2Y2mJmebmMeWGUpfxxWh6WVB+dtkVte\n3qLGxkvJ+9ZrY5euiFa6ddp4tvTUk7ROFfYGwTNz","token":"0e16fb98e2c49e676030082a2c71d66f947755ba"}`
	//ret := `{"content":"VZyQV7FOu66+O4o3H2JJrSlCNIs1AZqV/GH1jMghgVLinMWql+zzWOWmoXstYGWjkE9FNj7dtG/T\n4H64JetYcQzvmhrgw05qKe6W61ABJx4wGAUlPL8w/l5UB2Y2mJmeaaON2C+LCNkqQ/Ss+ExUxB++\n9RNQIQ343R3+zpW6YzSMiT2PO7BUmOsLDncD6jHY0rpJqBEImiwVZyUi+ZFMQDdrkYp479/CHpF8\nL8fGpldp8AACXnDC8HzZxkfF2ZeP80WyDamAxPjMJ4Y9WhdYCdlkMWMo/113xapJOD9Nc6zbM8t2\n6ekdptrLib3Ei6JnRw1PKvwYMJaKsg+u2L+Rgmyo9W9IIuqr5OMod8s6/zs3GIjw+6Qklt9uCMM3\n9UFe","token":"6c382044402e6b464c8c96cebffbb775015036ae"}`
	//ret := `{"content":"VZyQV7FOu66+O4o3H2JJrSlCNIs1AZqV/GH1jMghgVLinMWql+zzWOWmoXstYGWjkE9FNj7dtG/T\n4H64JetYcQzvmhrgw05qKe6W61ABJx4wGAUlPL8w/l5UB2Y2mJmeaaON2C+LCNkqQ/Ss+ExUxB++\n9RNQIQ343R3+zpW6YzSMiT2PO7BUmOsLDncD6jHY0rpJqBEImiwVZyUi+ZFMQDdrkYp479/CHpF8\nL8fGpldp8AACXnDC8HzZxkfF2ZeP80WyDamAxPjMJ4Y9WhdYCdlkMWMo/113xapJOD9Nc6zbM8t2\n6ekdptrLib3Ei6JnXbOvZigGD4XPx4Yd/YwTdhWH1Cu42N4lTEkeP012Mo03cBO5sDa9SNWTuUUs\nY+Sv","token":"6494b57cc1b3062eb962c16b837f88d1abdc2baf"}`
	//ret := `{"content":"VZyQV7FOu66+O4o3H2JJrSlCNIs1AZqV/GH1jMghgVLinMWql+zzWOWmoXstYGWjkE9FNj7dtG/T\n4H64JetYcQzvmhrgw05qKe6W61ABJx7JjEVVfHkLhbeu9SlTq+ypSlr4RAsNDgGA7ln+vRnGyvsK\njashKhvBhDm2A9TP3RPxi42Vd9LYmNga4150OsY8RsZuuvrKHHxPAxYAou9Szg\u003d\u003d","token":"211ae9c4a2568e51e9bb9aa6742e4deda73a9afb"}`
	//ret := `{"content":"VZyQV7FOu66+O4o3H2JJrSlCNIs1AZqV/GH1jMghgVLinMWql+zzWOWmoXstYGWjkE9FNj7dtG/T\n4H64JetYcQzvmhrgw05qKe6W61ABJx7JjEVVfHkLhbeu9SlTq+yp1BdvLCd6latjKuqsP71WonD5\nvs9baNIQpsL2AVhwweqMiT2PO7BUmOsLDncD6jHYlI5Z+JsED7WMkU1Ytauhs39Z7X7fw/Qtr+6v\nyZRUJ0NRHgiLvOF7JGYwbJ1fgfvO1wX7EJTy2XGkASgpFmd0Zb/2Dk8uHXDuBAdYyQBaefj+ooIK\nr3VAxxBScRb0RFYbONfhNCww8/Qxzqi8/mpJU8gUW63RUIWk0usrqP5pQf5WAX9StcRRN+LWFxZ3\nO8ntegeA8myuIlWdd57i0RKU9Q\u003d\u003d","token":"1d658c4db999cd2ce9f894ed6219ef782d186219"}`
	//ret := `{"content":"VZyQV7FOu66+O4o3H2JJrSlCNIs1AZqV/GH1jMghgVLinMWql+zzWOWmoXstYGWjkE9FNj7dtG/T\n4H64JetYcQzvmhrgw05qKe6W61ABJx7JjEVVfHkLhbeu9SlTq+yp1BdvLCd6latjKuqsP71WonD5\nvs9baNIQpsL2AVhwweqMiT2PO7BUmOsLDncD6jHYlI5Z+JsED7WMkU1Ytauhs39Z7X7fw/Qtr+6v\nyZRUJ0NRHgiLvOF7JGYwbJ1fgfvO1wX7EJTy2XGkASgpFmd0Zb/2Dk8uHXDuBAdYyQBaefj+ooIK\nr3VAxxBScRb0RFYbN8B3+zDebdz9EEXdrfFbRVm0H6SRHwAZPzfsPPkJCef/LPLyq/ljh5Z8lvlb\n4Uk1","token":"55c4939cfcccef2e11d4e6b48473e47756bac888"}`
	//ret := `{"content":"VZyQV7FOu66+O4o3H2JJrSlCNIs1AZqV/GH1jMghgVLinMWql+zzWOWmoXstYGWjkE9FNj7dtG/T\n4H64JetYcQzvmhrgw05qKe6W61ABJx5t3R6ycggxTeEIw+S6gW81Tblv5McxzZrhmC5IO26O0ADv\niOrkkX1qNG7eplQsCOg8DW1veXFxeHkCP1NwEDeT3/sHSDg0hwWmCktokMLtdGkXOKYQ/aS76yXM\nCuLU1Trk/q8QtHkMaTWbcENhatF9fGwXAKbkhFS+htUackRh/3x2PFqwK292ju59U55IlzLPZaXi\nbRvTaec0A0yrpPMw","token":"d7aa889c32ee53bf80299af43fe2e5d943a1e6f4"}`
	ret := `{"content":"VZyQV7FOu66+O4o3H2JJrSlCNIs1AZqV/GH1jMghgVLinMWql+zzWOWmoXstYGWjkE9FNj7dtG/T\n4H64JetYcQzvmhrgw05qKe6W61ABJx5+siLNYdYZI5h5IDIla3VqlhV2FhGUeU0AOg3ByoJ1Y004\n/JqWEpoddIaBi6chA6aMiT2PO7BUmOsLDncD6jHYK//7mXdKuS+jMaXajTU2ln9Z7X7fw/Qtr+6v\nyZRUJ0Ma1NW8ucs/gQJbweOERjuWffWS6uLGOrr/jflXR6pXV0u3X3OBLmwUrVgiKFBfOQ9+O2YH\nHI5p4DdBdE/th0MrONfhNCww8/Qxzqi8/mpJU8gUW63RUIWk0usrqP5pQf5WAX9StcRRN+LWFxZ3\nO8ntegeA8myuIlWdd57i0RKU9Q\u003d\u003d","token":"442aaf7539452f2ddba7dd28c0c5db8e28059bc0"}`

	result := notify.Handle(ret, func(data *YqbEnterNotifyData) error {
		fmt.Printf("%+v\n", data)
		return nil
	})

	fmt.Println(result)
}
