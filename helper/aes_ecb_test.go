package helper

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestAesEncrypt(t *testing.T) {

	content := "1"
	key, _ := base64.StdEncoding.DecodeString("U4YkiyL5NayGoeqwgvGW7g==")

	ciphertext := PKCS7Padding([]byte(content), 16)

	crypted, err := AesEncrypt(ciphertext, key) //ECB加密
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("crypted: ", crypted, base64.StdEncoding.EncodeToString(crypted))
	//

	origin, err := AesDecrypt(crypted, key) // ECB解密
	if err != nil {
		fmt.Println(err)
		return
	}

	origin = PKCS7UnPadding(origin)
	fmt.Println("decrypted: ", origin, string(origin))

	//b,err := AesEncrypt([]byte("1"),[]byte("4B8202A3BB3B4D650306BEE15DF35E13F70D2D704A4E4F8C"))
	//
	//data := base64.StdEncoding.EncodeToString(b)
	//
	//fmt.Println(data)
	//fmt.Println(err)
}

func TestAesEncrypt2(t *testing.T) {
	aesKey := NonceStr()

	origin := "1"

	fmt.Println(aesKey)

	cipherText := PKCS7Padding([]byte(origin), 16)

	encrypted, _ := AesEncrypt(cipherText, []byte(aesKey))

	fmt.Println(base64.StdEncoding.EncodeToString(encrypted))
}
