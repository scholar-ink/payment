package strings

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//生成四位短信验证码
func CreateIdentifyCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	return fmt.Sprintf("%04v", rnd.Int31n(10000))
}

func FormatFloat(floatNum float64) string {
	return strconv.FormatFloat(floatNum, 'f', -1, 64)
}
