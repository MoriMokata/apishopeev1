package shopee

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/pkg/errors"
)

const (
	ParamPartnerIdKey   = "partner_id"
	ParamTimestampKey   = "timestamp"
	ParamAccessTokenKey = "access_token"
	ParamShopIdKey      = "shop_id"
	ParamSignKey        = "sign"
)

var defaultCommonKey = []string{
	ParamPartnerIdKey,
	ParamShopIdKey,
	ParamTimestampKey,
	ParamAccessTokenKey,
	ParamSignKey,
}

type V2Response[T any] struct {
	RequestId string `json:"request_id"`
	Error     string `json:"error"`
	Message   string `json:"message"`
	Response  T      `json:"response"`
}

var logger = logs.GetBeeLogger()

type V2CommonParameter struct {
	PartnerKey  string
	PartnerId   int64
	AccessToken string
	ShopId      int64
	TimeStamp   int64
}

type ShopApiV2[T any] struct {
	baseUrl   string
	method    string
	path      string
	parameter url.Values
	body      []byte
	commonKey []string
}

func (r *ShopApiV2[T]) Get() *ShopApiV2[T] {
	r.method = http.MethodGet
	return r
}

func (r *ShopApiV2[T]) Post() *ShopApiV2[T] {
	r.method = http.MethodPost
	return r
}

func (r *ShopApiV2[T]) Path(path string) *ShopApiV2[T] {
	r.path = path
	return r
}

func (r *ShopApiV2[T]) SetParameter(key string, value any) {
	switch value.(type) {
	case string:
		r.parameter.Add(key, value.(string))
	case int, int8, int16, int32, int64:
		r.parameter.Add(key, strconv.FormatInt(value.(int64), 10))
	case bool:
		r.parameter.Add(key, strconv.FormatBool(value.(bool)))
	case float32, float64:
		r.parameter.Add(key, strconv.FormatFloat(value.(float64), 'f', 2, 64))
	}
}

func (r *ShopApiV2[T]) BindParameter() error {
	return nil
}

func (r *ShopApiV2[T]) Body(body any) {
	b, err := json.Marshal(body)
	if err != nil {
		r.body = b
	}
}

func (r *ShopApiV2[T]) sign(cParam *V2CommonParameter) (string, error) {
	s := fmt.Sprintf("%v%v%v%v%v", cParam.PartnerId, r.path, cParam.TimeStamp, cParam.AccessToken, cParam.ShopId)

	h := hmac.New(sha256.New, []byte(cParam.PartnerKey))
	_, err := h.Write([]byte(s))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func (r *ShopApiV2[T]) createRequest(cParam *V2CommonParameter) (*http.Response, error) {
	if r.baseUrl == "" {
		baseUrl, err := beego.AppConfig.String("shopee::SMP_SHOPEE_BASEURL")
		if err != nil {
			return nil, errors.New("cannot get base url from config 'shopee::base-url'")
		}
		r.baseUrl = baseUrl
	}

	if cParam.TimeStamp == 0 {
		cParam.TimeStamp = time.Now().Unix()
	}

	reqUrl, err := url.JoinPath(r.baseUrl, r.path)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	var req *http.Request
	if r.method == http.MethodPost && r.body != nil {
		req, err = http.NewRequest(r.method, reqUrl, bytes.NewReader(r.body))
	} else {
		req, err = http.NewRequest(r.method, reqUrl, nil)
	}
	if err != nil {
		return nil, err
	}

	sign, err := r.sign(cParam)
	if err != nil {
		return nil, err
	}

	// merge parameter
	common := url.Values{}
	if cParam.PartnerId > 0 {
		common.Add(ParamPartnerIdKey, strconv.FormatInt(cParam.PartnerId, 10))
	}
	if cParam.AccessToken != "" {
		common.Add(ParamAccessTokenKey, cParam.AccessToken)
	}
	if cParam.ShopId > 0 {
		common.Add(ParamShopIdKey, strconv.FormatInt(cParam.ShopId, 10))
	}
	if cParam.TimeStamp > 0 {
		common.Add(ParamTimestampKey, strconv.FormatInt(cParam.TimeStamp, 10))
	}
	if sign != "" {
		common.Add(ParamSignKey, sign)
	}

	//params := mergeUrlValue(common, r.parameter)
	commonKey := r.commonKey
	if commonKey == nil || len(commonKey) == 0 {
		commonKey = defaultCommonKey
	}
	req.URL.RawQuery = encodeParameter(commonKey, common, r.parameter)

	requestId := time.Now().UnixNano()
	logger.Info("[SHOPEE][%v][REQ][%s] %s", requestId, r.method, req.URL.RequestURI())
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	logger.Info("[SHOPEE][%v][RES] %v", requestId, res.Status)

	return res, nil
}

func (r *ShopApiV2[T]) Do(cParam *V2CommonParameter) (*V2Response[T], error) {
	res, err := r.createRequest(cParam)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resBody := new(V2Response[T])
	if err = json.Unmarshal(resBytes, resBody); err != nil {
		return nil, err
	}

	if resBody.Error != "" {
		return nil, errors.Errorf("%s", resBytes)
	}

	return resBody, nil
}
