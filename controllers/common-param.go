package controllers

import (
	"github.com/gorilla/schema"
	"github.com/pkg/errors"
	"reflect"
	"shopeeadapterapi/shopee"
	"strconv"
	"time"
)

var schemaDecoder *schema.Decoder

func init() {
	timeConverter := func(value string) reflect.Value {
		// timestamp
		if timestamp, err := strconv.ParseInt(value, 10, 64); err == nil {
			return reflect.ValueOf(time.Unix(timestamp, 0))
		}
		// yyyy-MM-ddTHH:mm:ss
		if v, err := time.Parse("2006-01-02T15:04:05", value); err == nil {
			return reflect.ValueOf(v)
		}
		// yyyy-MM-dd HH:mm:ss
		if v, err := time.Parse("2006-01-02 15:04:05", value); err == nil {
			return reflect.ValueOf(v)
		}

		return reflect.ValueOf(time.Time{})
	}

	schemaDecoder = schema.NewDecoder()
	schemaDecoder.RegisterConverter(time.Time{}, timeConverter)
}

type Credential struct {
	AccessToken  string
	RefreshToken string
	PartnerId    string
	PartnerKey   string
}

func createCommonParameter(h map[string][]string) (*shopee.V2CommonParameter, error) {
	param := new(shopee.V2CommonParameter)

	if v, ok := h["Adaptertoken"]; ok {
		param.AccessToken = v[0]
	} else {
		return nil, errors.New("header 'adapterToken' is required")
	}
	if v, ok := h["Adapterkey"]; ok {
		param.PartnerKey = v[0]
	} else {
		return nil, errors.New("header 'adapterKey' is required")
	}
	if v, ok := h["Adaptersecret"]; ok {
		if vi, err := strconv.ParseInt(v[0], 10, 64); err == nil {
			param.PartnerId = vi
		}
	}
	//if v, ok := h["adapterRefresh"]; ok {
	//	param.RefreshToken = v[0]
	//}

	return param, nil
}
