package enter

import (
	"archive/zip"
	"bytes"
	"github.com/scholar-ink/payment/helper"
	"github.com/scholar-ink/payment/util/http"

	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"strings"
)

const (
	XunLianEnterUrl = "https://www.xunliandata.com/merServPlat/merApi/"
)

type XunLianEnter struct {
	*XunLianConfig
	ReqId   string `json:"reqId"` //入网唯一标识
	Service string `json:"-"`
	Sign    string `json:"-"`
}

type XunLianConfig struct {
	InsCode   string `json:"insCode,omitempty"`   //机构代码
	AgentCode string `json:"agentCode,omitempty"` //代理代码
	GroupCode string `json:"groupCode,omitempty"` //集团商户代码
	Key       string `json:"-"`
}

type XunLianCreateConf struct {
	ReqId                  string `json:"-"`
	WxpChanFlag            string `json:"wxpChanFlag"`                      //微信业务类型
	AlpChanFlag            string `json:"alpChanFlag"`                      //支付宝业务类型
	AlpLevel               string `json:"alpLevel"`                         //入驻等级
	MerType                string `json:"merType"`                          //商户类型
	BigCategoryName        string `json:"bigCategoryName"`                  //商户类目
	SubCategoryCode        string `json:"subCategoryCode"`                  //MCC
	ProvinceName           string `json:"provinceName"`                     //商户所在省份
	CityName               string `json:"cityName"`                         //商户所在城市
	DistrictName           string `json:"districtName"`                     //商户所在区县
	ThreeIntoOne           string `json:"threeIntoOne"`                     //是否选择三证合一
	IsInstitution          string `json:"isInstitution"`                    //是否事业单位
	BizAddress             string `json:"bizAddress"`                       //经营地址
	MerName                string `json:"merName"`                          //注册名称
	ShortName              string `json:"shortName"`                        //经营名称
	LicenseCode            string `json:"licenseCode,omitempty"`            //营业执照注册号
	LicenseEffective       string `json:"licenseEffective,omitempty"`       //营业执照注册日期
	LicenseExpired         string `json:"licenseExpired,omitempty"`         //营业执照有效期
	CardType               string `json:"cardType"`                         //证件类型
	CardHolderType         string `json:"cardHolderType"`                   //证件持有人类型1.法人
	LegalName              string `json:"legalName"`                        //证件持有人姓名
	LegalCard              string `json:"legalCard"`                        //证件号码
	LegalExpiredBegin      string `json:"legalExpiredBegin"`                //证件起始日期
	LegalExpired           string `json:"legalExpired"`                     //证件有效期
	DelayType              string `json:"delayType"`                        //账户类型 对公结算=1/对私结算=0
	BranchProvince         string `json:"branchProvince"`                   //开户银行省份
	BranchCityName         string `json:"branchCityName"`                   //开户银行城市
	BankName               string `json:"bankName"`                         //开户银行支行
	AccountName            string `json:"accountName"`                      //收款人名称
	AccountCode            string `json:"accountCode"`                      //银行账户
	BankPhone              string `json:"bankPhone"`                        //银行预留手机号
	ContactName            string `json:"contactName"`                      //联系人姓名
	ContactMobile          string `json:"contactMobile"`                    //联系手机号码
	ContactFixed           string `json:"contactFixed"`                     //客服电话
	ContactEmail           string `json:"contactEmail"`                     //联系人邮箱
	WxpMerCode             string `json:"wxpMerCode"`                       //微信一级商户号
	BusinessCategoryCode   string `json:"businessCategoryCode"`             //微信类目
	Wxp                    string `json:"wxp"`                              //微信手续费
	AlpFirstMerCode        string `json:"alpFirstMerCode"`                  //支付宝一级商户号
	AlpCategoryCodeV3      string `json:"alpCategoryCodeV3,omitempty"`      //支付宝类目
	ChannelAlpCategoryCode string `json:"channelAlpCategoryCode,omitempty"` //支付宝类目
	Alp                    string `json:"alp"`                              //支付宝手续费
	//PortalPhoto          string `json:"portalPhoto"`                 //经营场所照片(门面)
	//ScenePhoto           string `json:"scenePhoto"`                  //经营场所照片(内景)
	//BizLicense           string `json:"bizLicense,omitempty"`        //营业执照照片
	//LegalIdCardZ         string `json:"legalIdCardZ"`                //法人身份证件(正面)
	//LegalIdCardF         string `json:"legalIdCardF"`                //法人身份证件(反面)
	//HandheldIdCard       string `json:"handheldIdCard,omitempty"`    //手持身份证照片
	//InsCert              string `json:"insCert,omitempty"`           //组织机构代码证照片
	//TaxRegisterCert      string `json:"taxRegisterCert,omitempty"`   //税务登记证照片
	//InAccountBankCard    string `json:"inAccountBankCard,omitempty"` //银行卡照片
	//OpenLicense          string `json:"openLicense,omitempty"`       //开户许可证照片
}

type XunLianCreateReturn struct {
	Message string `json:"message"`
	Status  int32  `json:"status"`
}

type XunLianQueryConf struct {
	MerchantId    string `json:"merchantId,omitempty"`    //平安付商户号
	OutMerchantNo string `json:"outMerchantNo,omitempty"` //服务商侧商户号
	PageNo        string `json:"pageNo"`                  //页码
	PageSize      string `json:"pageSize"`                //每页记录数
}

type XunLianQueryReturn struct {
	MerchantNo      string `json:"merchant_no"`
	MerchantStatus  string `json:"mer_status"`
	MerchantListStr string `json:"mer_list"`
}

type XunLianUploadConf struct {
	User     string
	Pem      []byte
	FileList map[string][]byte
}

func (xunLian *XunLianEnter) InitConfig(config *XunLianConfig) {
	xunLian.XunLianConfig = config
}

func (xunLian *XunLianEnter) Create(config *XunLianCreateConf) (*XunLianCreateReturn, error) {
	xunLian.Service = "acceptMer.json"

	xunLian.ReqId = config.ReqId

	ret, err := xunLian.sendReq(XunLianEnterUrl, config)
	if err != nil {
		return nil, err
	}
	enterReturn := new(XunLianCreateReturn)
	err = json.Unmarshal(ret, &enterReturn)
	if err != nil {
		return nil, err
	}

	return enterReturn, nil
}
func (xunLian *XunLianEnter) Query(config *XunLianQueryConf) (*XunLianQueryReturn, error) {
	xunLian.Service = "queryMerParam.json"

	ret, err := xunLian.sendReq(XunLianEnterUrl, config)
	if err != nil {
		return nil, err
	}

	queryReturn := new(XunLianQueryReturn)
	err = json.Unmarshal(ret, &queryReturn)
	if err != nil {
		return nil, err
	}

	return queryReturn, nil
}
func (xunLian *XunLianEnter) Upload(config *XunLianUploadConf) (string, error) {

	reqId := strings.ReplaceAll(uuid.New().String(), "-", "")

	_, err := xunLian.Zip(config.FileList)

	if err != nil {
		return "", err
	}

	return reqId, nil
}
func (xunLian *XunLianEnter) buildData(config interface{}) error {
	return nil
}
func (xunLian *XunLianEnter) makeSign(content []byte) {
	fmt.Println(string(content) + xunLian.Key)
	xunLian.Sign = helper.Sha256(string(content) + xunLian.Key)
	fmt.Println(xunLian.Sign)
}

func (xunLian *XunLianEnter) sendReq(reqUrl string, params interface{}) (b []byte, err error) {

	mapData := helper.Struct2Map(params)

	mapData["insCode"] = xunLian.InsCode
	mapData["agentCode"] = xunLian.AgentCode
	mapData["groupCode"] = xunLian.GroupCode
	mapData["reqId"] = xunLian.ReqId

	b, _ = json.Marshal(mapData)

	xunLian.makeSign(b)

	req := http.NewHttpRequest("POST", reqUrl+xunLian.Service, bytes.NewBuffer(b))

	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Set("Signature", xunLian.Sign)

	rsp, err := http.Client.Do(req)

	defer rsp.Body.Close()

	b, err = ioutil.ReadAll(rsp.Body)

	return
}

func (xunLian *XunLianEnter) Zip(fileList map[string][]byte) ([]byte, error) {

	// 实例化新的 zip.Writer
	buf := new(bytes.Buffer)

	zw := zip.NewWriter(buf)

	for fileName, file := range fileList {

		f, err := zw.Create(fileName + ".png")

		if err != nil {
			return nil, err
		}
		_, err = f.Write(file)

		if err != nil {
			return nil, err
		}
	}

	err := zw.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
