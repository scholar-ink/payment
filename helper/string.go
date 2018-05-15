package helper

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/chanxuehong/rand"
)

func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	cipherStr := h.Sum(nil)

	return hex.EncodeToString(cipherStr)
}

func NonceStr() string {
	return string(rand.NewHex())
}
