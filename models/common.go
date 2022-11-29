package models

import (
	"encoding/json"
	"strings"
)

const (
	ErrorParameter = "error_param"
	ErrorInternal  = "error_internal"
)

type ErrorResponse struct {
	Error     string `json:"error"`
	Message   string `json:"message"`
	RequestId string `json:"request_id"`
}

func NewErrorResponse(errCode string, msg string) []byte {
	model := &ErrorResponse{
		Error:   errCode,
		Message: msg,
	}
	b, _ := json.Marshal(&model)
	return b
}

func NewErrorResponseString(errCode string, msg string) string {
	if strings.HasPrefix(msg, "{") {
		return msg
	}
	return string(NewErrorResponse(errCode, msg))
}
