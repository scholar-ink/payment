package helper

import (
	"bytes"
	"crypto"
	"crypto/hmac"
	"crypto/md5"
	rand2 "crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/scholar-ink/go-pkcs12"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	cipherStr := h.Sum(nil)

	return hex.EncodeToString(cipherStr)
}

func Sha1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	cipherStr := h.Sum(nil)

	return hex.EncodeToString(cipherStr)
}
func Sha256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	cipherStr := h.Sum(nil)

	return hex.EncodeToString(cipherStr)
}

func HMacSha256(origData []byte, key string) []byte {
	h := hmac.New(sha256.New, []byte(key))
	h.Write(origData)
	return h.Sum(nil)
}

//生产orderSn
func CreateSn() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strings.Replace(time.Now().Format("060102150405.000000"), ".", "", -1) + fmt.Sprintf("%06v", rnd.Int31n(1000000))
}

func NonceStr() string {
	return Md5(uuid.New().String())
}

//CurrentTimeStampMS get current time with millisecond
func CurrentTimeStampMS() int64 {
	return time.Now().UnixNano() / time.Millisecond.Nanoseconds()
}

//rsa公钥加密
func Rsa1Encrypt(pfxData, origData []byte, certPassWord string) (string, error) {

	_, cert, err := pkcs12.Decode(pfxData, certPassWord)

	if err != nil {
		return "", err
	}

	pubKey := cert.PublicKey

	var pub = pubKey.(*rsa.PublicKey)

	partLen := pub.N.BitLen()/8 - 11

	chunks := split([]byte(origData), partLen)

	buffer := bytes.NewBufferString("")

	for _, chunk := range chunks {

		b, err := rsa.EncryptPKCS1v15(rand2.Reader, pub, chunk)

		if err != nil {
			return "", err
		}

		buffer.Write(b)

	}
	return base64.StdEncoding.EncodeToString(buffer.Bytes()), nil
}

func Rsa1Encrypt2(origData []byte, publicKey string) (string, error) {

	key := ParsePublicKey(publicKey)
	block, _ := pem.Decode([]byte(key)) //PublicKeyData为私钥文件的字节数组
	if block == nil {
		fmt.Println("Public block空")
		return "", nil
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}

	var pub = pubKey.(*rsa.PublicKey)

	partLen := pub.N.BitLen()/8 - 11

	chunks := split([]byte(origData), partLen)

	buffer := bytes.NewBufferString("")

	for _, chunk := range chunks {

		b, err := rsa.EncryptPKCS1v15(rand2.Reader, pub, chunk)

		if err != nil {
			return "", err
		}

		buffer.Write(b)

	}
	return base64.StdEncoding.EncodeToString(buffer.Bytes()), nil
}

//rsa私钥解密
func Rsa1Decrypt(pfxData []byte, origData string, certPassWord string) ([]byte, error) {

	private, _, err := pkcs12.Decode(pfxData, certPassWord)

	if err != nil {
		return nil, err
	}

	var pri = private.(*rsa.PrivateKey)

	//分段加密
	partLen := pri.N.BitLen() / 8

	raw, err := base64.StdEncoding.DecodeString(origData)

	chunks := split([]byte(raw), partLen)

	buffer := bytes.NewBufferString("")

	for _, chunk := range chunks {

		decrypted, err := rsa.DecryptPKCS1v15(rand2.Reader, pri, chunk)

		if err != nil {
			return nil, err
		}

		buffer.Write(decrypted)

	}

	return buffer.Bytes(), err
}

//rsa私钥解密
func Rsa1Decrypt2(origData string, privateKey string) ([]byte, error) {

	//私钥切片处理
	key := ParsePrivateKey(privateKey)

	block, _ := pem.Decode([]byte(key)) //PiravteKeyData为私钥文件的字节数组
	if block == nil {
		fmt.Println("block空")
		return nil, nil
	}
	//priv即私钥对象,block2.Bytes是私钥的字节流
	privKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)

	priv := privKey.(*rsa.PrivateKey)

	//分段加密
	partLen := priv.N.BitLen() / 8

	raw, err := base64.StdEncoding.DecodeString(origData)

	chunks := split([]byte(raw), partLen)

	buffer := bytes.NewBufferString("")

	for _, chunk := range chunks {

		decrypted, err := rsa.DecryptPKCS1v15(rand2.Reader, priv, chunk)

		if err != nil {
			return nil, err
		}

		buffer.Write(decrypted)

	}

	return buffer.Bytes(), err
}

//rsa公钥加密bcd
func Rsa1EncryptBcd(pfxData, origData []byte, certPassWord string) (string, error) {

	_, cert, err := pkcs12.Decode(pfxData, certPassWord)

	if err != nil {
		return "", err
	}

	pubKey := cert.PublicKey

	var pub = pubKey.(*rsa.PublicKey)

	partLen := pub.N.BitLen()/8 - 11

	chunks := split([]byte(origData), partLen)

	buffer := bytes.NewBufferString("")

	for _, chunk := range chunks {

		b, err := rsa.EncryptPKCS1v15(rand2.Reader, pub, chunk)

		if err != nil {
			return "", err
		}

		buffer.Write(b)

	}
	return Bcd2Str(buffer.Bytes()), nil
}

//rsa私钥解密bcd
func Rsa1DecryptBcd(pfxData []byte, origData string, certPassWord string) ([]byte, error) {

	private, _, err := pkcs12.Decode(pfxData, certPassWord)

	if err != nil {
		return nil, err
	}

	var pri = private.(*rsa.PrivateKey)

	//分段加密
	partLen := pri.N.BitLen() / 8

	raw := Bcd2Bytes([]byte(origData))

	chunks := split(raw, partLen)

	buffer := bytes.NewBufferString("")

	for _, chunk := range chunks {

		decrypted, err := rsa.DecryptPKCS1v15(rand2.Reader, pri, chunk)

		if err != nil {
			return nil, err
		}

		buffer.Write(decrypted)

	}

	return buffer.Bytes(), err
}

//rsa公钥加密
func Rsa2Encrypt(origData []byte, publicKey string) (string, error) {

	key := ParsePublicKey(publicKey)
	block, _ := pem.Decode([]byte(key)) //PublicKeyData为私钥文件的字节数组
	if block == nil {
		fmt.Println("Public block空")
		return "", nil
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}

	var pub = pubKey.(*rsa.PublicKey)

	partLen := pub.N.BitLen()/8 - 11

	chunks := split([]byte(origData), partLen)

	buffer := bytes.NewBufferString("")

	for _, chunk := range chunks {

		b, err := rsa.EncryptPKCS1v15(rand2.Reader, pub, chunk)

		if err != nil {
			return "", err
		}

		buffer.Write(b)

	}
	return base64.StdEncoding.EncodeToString(buffer.Bytes()), nil
}
func Rsa2Encrypt3(origData []byte, publicKey string) (string, error) {

	key := ParsePublicKey(publicKey)
	block, _ := pem.Decode([]byte(key)) //PublicKeyData为私钥文件的字节数组
	if block == nil {
		fmt.Println("Public block空")
		return "", nil
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}

	var pub = pubKey.(*rsa.PublicKey)

	b, err := rsa.EncryptPKCS1v15(rand2.Reader, pub, origData)

	return base64.StdEncoding.EncodeToString(b), nil
}

func Rsa2EncryptByCert(origData []byte, publicKey string) (string, error) {

	key := ParsePublicKey(publicKey)
	block, _ := pem.Decode([]byte(key)) //PublicKeyData为私钥文件的字节数组
	if block == nil {
		fmt.Println("Public block空")
		return "", nil
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return "", err
	}

	var pub = cert.PublicKey.(*rsa.PublicKey)

	b, err := rsa.EncryptPKCS1v15(rand2.Reader, pub, origData)

	//partLen := pub.N.BitLen()/8 - 11
	//
	//chunks := split([]byte(origData), partLen)
	//
	//buffer := bytes.NewBufferString("")
	//
	//for _, chunk := range chunks {
	//
	//	b, err := rsa.EncryptPKCS1v15(rand2.Reader, pub, chunk)
	//
	//	if err != nil {
	//		return "", err
	//	}
	//
	//	buffer.Write(b)
	//
	//}
	return base64.StdEncoding.EncodeToString(b), nil
}
func Rsa2DecryptByCert(origData []byte, publicKey string) (string, error) {

	key := ParsePublicKey(publicKey)
	block, _ := pem.Decode([]byte(key)) //PublicKeyData为私钥文件的字节数组
	if block == nil {
		fmt.Println("Public block空")
		return "", nil
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return "", err
	}

	var pub = cert.PublicKey.(*rsa.PublicKey)

	b, err := rsa.EncryptPKCS1v15(rand2.Reader, pub, origData)

	return base64.StdEncoding.EncodeToString(b), nil
}

//rsa私钥解密
func Rsa2Decrypt(origData []byte, privateKey string) ([]byte, error) {

	//私钥切片处理
	key := ParsePrivateKey(privateKey)

	block, _ := pem.Decode([]byte(key)) //PiravteKeyData为私钥文件的字节数组
	if block == nil {
		fmt.Println("block空")
		return nil, nil
	}

	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	//分段加密
	partLen := pri.N.BitLen() / 8

	raw, err := base64.StdEncoding.DecodeString(string(origData))

	chunks := split([]byte(raw), partLen)

	buffer := bytes.NewBufferString("")

	for _, chunk := range chunks {

		decrypted, err := rsa.DecryptPKCS1v15(rand2.Reader, pri, chunk)

		if err != nil {
			return nil, err
		}

		buffer.Write(decrypted)

	}

	return buffer.Bytes(), err
}
func Rsa2Decrypt2(origData []byte, privateKey string) ([]byte, error) {

	//私钥切片处理
	key := ParsePrivateKey(privateKey)

	block, _ := pem.Decode([]byte(key)) //PiravteKeyData为私钥文件的字节数组
	if block == nil {
		fmt.Println("block空")
		return nil, nil
	}

	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	decrypted, err := rsa.DecryptPKCS1v15(rand2.Reader, pri, origData)

	return decrypted, err
}

//rsa私钥加密
func Rsa2Encrypt2(origData []byte, privateKey string) ([]byte, error) {

	key := ParsePrivateKey(privateKey)
	block, _ := pem.Decode([]byte(key)) //PublicKeyData为私钥文件的字节数组
	if block == nil {
		return nil, errors.New("Public block空")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	signData, err := rsa.SignPKCS1v15(nil, priv, crypto.Hash(0), origData)

	return signData, nil
}

func Md5WithRsaSignWithPKCS8(origData []byte, privateKey string) ([]byte, error) {
	//私钥切片处理
	key := ParsePrivateKey(privateKey)

	block, _ := pem.Decode([]byte(key)) //PiravteKeyData为私钥文件的字节数组
	if block == nil {
		fmt.Println("block空")
		return nil, nil
	}
	//priv即私钥对象,block2.Bytes是私钥的字节流
	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)

	if err != nil {
		fmt.Println(err)
		fmt.Println("无法还原私钥")
		return nil, nil
	}
	h2 := md5.New()
	h2.Write(origData)
	hashed := h2.Sum(nil)
	signature2, err := rsa.SignPKCS1v15(rand2.Reader, priv.(*rsa.PrivateKey), crypto.MD5, hashed) //签名
	return signature2, err
}

//rsa2加密
func Sha256WithRsaSignWithPKCS8(origData []byte, privateKey string) ([]byte, error) {
	//私钥切片处理
	key := ParsePrivateKey(privateKey)

	block, _ := pem.Decode([]byte(key)) //PiravteKeyData为私钥文件的字节数组
	if block == nil {
		fmt.Println("block空")
		return nil, nil
	}
	//priv即私钥对象,block2.Bytes是私钥的字节流
	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)

	if err != nil {
		fmt.Println(err)
		fmt.Println("无法还原私钥")
		return nil, nil
	}
	h2 := sha256.New()
	h2.Write(origData)
	hashed := h2.Sum(nil)
	signature2, err := rsa.SignPKCS1v15(rand2.Reader, priv.(*rsa.PrivateKey), crypto.SHA256, hashed) //签名
	return signature2, err
}

//rsa2加密
func Sha256WithRsaSign(origData []byte, privateKey string) ([]byte, error) {
	//私钥切片处理
	key := ParsePrivateKey(privateKey)

	block, _ := pem.Decode([]byte(key)) //PiravteKeyData为私钥文件的字节数组
	if block == nil {
		fmt.Println("block空")
		return nil, nil
	}
	//priv即私钥对象,block2.Bytes是私钥的字节流
	var priv *rsa.PrivateKey

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	if err != nil {
		fmt.Println(err)
		fmt.Println("无法还原私钥")
		return nil, nil
	}
	h2 := sha256.New()
	h2.Write(origData)
	hashed := h2.Sum(nil)
	signature2, err := rsa.SignPKCS1v15(rand2.Reader, priv, crypto.SHA256, hashed) //签名
	return signature2, err
}

func Sha256WithRsaVerify(src []byte, sign, publicKey string) error {

	signBytes, err := base64.StdEncoding.DecodeString(sign)

	if err != nil {
		return err
	}
	key := ParsePublicKey(publicKey)
	block, _ := pem.Decode([]byte(key)) //PublicKeyData为私钥文件的字节数组
	if block == nil {
		return errors.New("Public block空")
	}

	var pubI interface{}
	pubI, err = x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	h2 := sha256.New()
	h2.Write(src)
	hashed := h2.Sum(nil)
	var pub = pubI.(*rsa.PublicKey)
	err = rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed, signBytes) //验签

	if err != nil {
		return err
	}
	return nil
}

//rsa1加密
func Sha1WithRsaSign(origData []byte, privateKey string) ([]byte, error) {
	//私钥切片处理
	key := ParsePrivateKey(privateKey)

	block, _ := pem.Decode([]byte(key)) //PiravteKeyData为私钥文件的字节数组
	if block == nil {
		fmt.Println("block空")
		return nil, nil
	}
	//priv即私钥对象,block2.Bytes是私钥的字节流
	var priv *rsa.PrivateKey

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	if err != nil {
		fmt.Println(err)
		fmt.Println("无法还原私钥")
		return nil, nil
	}
	h2 := sha1.New()
	//h2.Write(origData)
	h2.Write([]byte([]byte(origData)))
	hashed := h2.Sum(nil)
	signature2, err := rsa.SignPKCS1v15(rand2.Reader, priv, crypto.SHA1, hashed) //签名
	return signature2, err
}
func Sha1WithRsaSignPkcs8(origData []byte, privateKey string) ([]byte, error) {
	//私钥切片处理
	key := ParsePrivateKey(privateKey)

	block, _ := pem.Decode([]byte(key)) //PiravteKeyData为私钥文件的字节数组
	if block == nil {
		fmt.Println("block空")
		return nil, nil
	}
	//priv即私钥对象,block2.Bytes是私钥的字节流
	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)

	if err != nil {
		fmt.Println(err)
		fmt.Println("无法还原私钥")
		return nil, nil
	}
	h2 := sha1.New()
	//h2.Write(origData)
	h2.Write([]byte([]byte(origData)))
	hashed := h2.Sum(nil)
	signature2, err := rsa.SignPKCS1v15(rand2.Reader, priv.(*rsa.PrivateKey), crypto.SHA1, hashed[:]) //签名
	return signature2, err
}

func Sha1WithRsaVerify(src []byte, sign, publicKey string) error {

	signBytes, err := base64.StdEncoding.DecodeString(sign)

	if err != nil {
		return err
	}
	key := ParsePublicKey(publicKey)
	block, _ := pem.Decode([]byte(key)) //PublicKeyData为私钥文件的字节数组
	if block == nil {
		return errors.New("Public block空")
	}

	var pubI interface{}
	pubI, err = x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	h2 := sha1.New()
	h2.Write(src)
	hashed := h2.Sum(nil)
	var pub = pubI.(*rsa.PublicKey)
	err = rsa.VerifyPKCS1v15(pub, crypto.SHA1, hashed, signBytes) //验签

	if err != nil {
		return err
	}
	return nil
}

func Sha1WithRsaVerifyHex(src []byte, sign, publicKey string) error {

	signBytes, err := hex.DecodeString(sign)

	if err != nil {
		return err
	}
	key := ParsePublicKey(publicKey)
	block, _ := pem.Decode([]byte(key)) //PublicKeyData为私钥文件的字节数组
	if block == nil {
		return errors.New("Public block空")
	}

	var pubI interface{}
	pubI, err = x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	h2 := sha1.New()
	h2.Write(src)
	hashed := h2.Sum(nil)
	var pub = pubI.(*rsa.PublicKey)
	err = rsa.VerifyPKCS1v15(pub, crypto.SHA1, hashed, signBytes) //验签

	if err != nil {
		return err
	}
	return nil
}

//rsa2加密
func Sha256WithRsaSignWithPassWord(pfxData, origData []byte, certPassWord string) ([]byte, error) {

	key, _, err := pkcs12.Decode(pfxData, certPassWord)

	if err != nil {
		return nil, err
	}

	//priv即私钥对象
	var priv = key.(*rsa.PrivateKey)

	if err != nil {
		fmt.Println("无法还原私钥" + err.Error())
		return nil, nil
	}
	h2 := sha256.New()
	h2.Write(origData)
	hashed := h2.Sum(nil)
	signature2, err := rsa.SignPKCS1v15(rand2.Reader, priv, crypto.SHA256, hashed) //签名
	return signature2, err
}

func split(buf []byte, lim int) [][]byte {

	var chunk []byte

	chunks := make([][]byte, 0, len(buf)/lim+1)

	for len(buf) >= lim {

		chunk, buf = buf[:lim], buf[lim:]

		chunks = append(chunks, chunk)

	}
	if len(buf) > 0 {

		chunks = append(chunks, buf[:])

	}

	return chunks
}

func RawUrlEncode(str string) string {
	return strings.Replace(url.QueryEscape(str), "+", "%20", -1)
}

/**
* @Title: DeepCopy
* @Description: todo(深度拷贝)
* @author zhouchao
 */
func DeepCopy(model interface{}, rpc interface{}) (err error) {
	b, err := json.Marshal(model)

	if err != nil {
		return
	}

	err = json.Unmarshal(b, rpc)
	return
}
