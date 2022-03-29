package strings

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/scholar-ink/payment/helper"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestImplode(t *testing.T) {
	fmt.Println(uuid.New().String())
}

func BenchmarkCreateSn(b *testing.B) {
	orders := make(map[string]bool, 0)
	for { //use b.N for looping
		orderSn := helper.CreateSn()
		fmt.Println(orders[orderSn])
		if _, ok := orders[orderSn]; ok {
			b.Fatal("订单号重复：" + orderSn)
		}
		orders[orderSn] = true
		fmt.Println(orderSn+"总数：", len(orders))
	}
}

func TestMd5(t *testing.T) {
	for {
		fmt.Println(time.Now().Format("060102150405") + fmt.Sprintf("%03d", rand.Int31n(1000)))
	}
}

func TestIsMobileNumber(t *testing.T) {

	a := strconv.FormatFloat(100000.100876, 'f', -1, 64)

	fmt.Println(a)

	fmt.Println(fmt.Sprintf("%.2f", 0.1))
}

func TestCreateIdentifyCode(t *testing.T) {

	version := "1.1.1"

	version = strings.Replace(version, ".", "", -1)

	fmt.Println(strconv.Atoi(version))
}
