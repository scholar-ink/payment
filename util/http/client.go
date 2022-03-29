package http

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const (
	MaxIdleConnections = 60
	RequestTimeout     = 60
)

var (
	Client *http.Client
)

func init() {
	Client = newHttpClient()
}

const (
	ProxyUserName = "16MVTUSM"
	ProxyPassword = "027366"
	ProxyServer   = "p5.t.16yun.cn:6445"
)

func newHttpClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: false,
			DialContext: (&net.Dialer{
				Timeout:   10 * time.Second, //连接超时时间
				KeepAlive: 30 * time.Second, //连接保持超时时间
			}).DialContext,
			MaxIdleConnsPerHost:   MaxIdleConnections,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 10 * time.Second,
			TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}

	return client
}

func NewTLSHttpClient(certFile, keyFile string) (httpClient *http.Client) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   3 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   3 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			TLSClientConfig:       tlsConfig,
		},
		Timeout: 5 * time.Second,
	}
}

func NewTLSBlockHttpClient(certPEMBlock, keyPEMBlock []byte) (httpClient *http.Client) {
	cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)

	if err != nil {
		return nil
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   3 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   3 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			TLSClientConfig:       tlsConfig,
		},
		Timeout: 5 * time.Second,
	}
}

func NewHttpRequest(method, url string, body io.Reader) *http.Request {

	request, _ := http.NewRequest(method, url, body)

	return request
}

func GetBytes(url string) ([]byte, error) {
	request := NewHttpRequest("GET", url, nil)

	rsp, err := Client.Do(request)

	if err != nil {
		return nil, err
	}

	defer rsp.Body.Close()

	b, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		return nil, err
	}

	return b, nil
}
