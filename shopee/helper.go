package shopee

import (
	"fmt"
	"net/url"
	"strings"
)

func mergeUrlValue(values ...url.Values) url.Values {
	params := url.Values{}
	for _, value := range values {
		for k, v := range value {
			for _, ev := range v {
				params.Add(k, ev)
			}
		}
	}
	return params
}

func encodeParameter(commonKey []string, common, request url.Values) string {
	qList := make([]string, 0)
	// common parameter
	for _, key := range commonKey {
		if value := common.Get(key); value != "" {
			qList = append(qList, fmt.Sprintf("%s=%s", key, url.QueryEscape(value)))
		}
	}
	// request parameter
	for key, values := range request {
		if len(values) > 0 && values[0] != "" {
			qList = append(qList, fmt.Sprintf("%s=%s", key, url.QueryEscape(values[0])))
		}
	}

	return strings.Join(qList, "&")
}
