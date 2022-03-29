package duolabao

import (
	"bytes"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"

	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

const (
	ChargeUrl = "https://openapi.duolabao.com"
	SUCCESS   = "SUCCESS"
)

type BaseCharge struct {
	*BaseConfig
}

type BaseConfig struct {
	Timestamp string `json:"-"`
	AccessKey string `json:"-"`
	SecretKey string `json:"-"`
	Token     string `json:"-"`
	Path      string `json:"path"`
}

func (base *BaseCharge) InitBaseConfig(config *BaseConfig) {
	config.Timestamp = strconv.Itoa(int(time.Now().Unix()))
	base.BaseConfig = config
}

func (base *BaseCharge) RetData(ret []byte) (b []byte, err error) {

	result := new(struct {
		Error struct {
			ErrorCode string `json:"errorCode"`
			ErrorMsg  string `json:"errorMsg"`
		} `json:"error"`
		Data   interface{} `json:"data"`
		Result string      `json:"result"`
	})

	err = json.Unmarshal(ret, &result)

	if err != nil {
		return
	}

	if result.Result != "success" {
		return nil, errors.New(result.Error.ErrorMsg)
	}

	return json.Marshal(result.Data)
}

func (base *BaseCharge) makeSign(requestJson string) {

	fmt.Println(fmt.Sprintf("secretKey=%s&timestamp=%s&path=%s&body=%s", base.SecretKey, base.Timestamp, base.Path, requestJson))

	base.Token = strings.ToUpper(helper.Sha1(fmt.Sprintf("secretKey=%s&timestamp=%s&path=%s&body=%s", base.SecretKey, base.Timestamp, base.Path, requestJson)))
}

func (base *BaseCharge) sendReq(reqUrl string, params interface{}) (b []byte, err error) {

	b, err = json.Marshal(params)

	if err != nil {
		return nil, err
	}

	base.makeSign(string(b))

	req := http.NewHttpRequest("POST", reqUrl+base.Path, bytes.NewBuffer(b))

	req.Header["token"] = []string{base.Token}
	req.Header["accessKey"] = []string{base.AccessKey}
	req.Header["timestamp"] = []string{base.Timestamp}
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("accessKey", base.AccessKey)
	//req.Header.Set("timestamp", base.Timestamp)
	//req.Header.Set("token", base.Token)

	fmt.Println(req.Header)

	rsp, err := http.Client.Do(req)

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	return
}
