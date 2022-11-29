package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	"io"
	"net/http"
	"shopeeadapterapi/models"
	"strings"
)

var log *logs.BeeLogger

func init() {
	log = logs.GetBeeLogger()
	log.Async()
}

type MyResponseWriter struct {
	http.ResponseWriter
	statusCode int
	buf        *bytes.Buffer
}

func (mrw *MyResponseWriter) Write(p []byte) (int, error) {
	return mrw.buf.Write(p)
}

func (mrw *MyResponseWriter) WriteHeader(statusCode int) {
	mrw.statusCode = statusCode
	mrw.ResponseWriter.WriteHeader(statusCode)
}

func TransformResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Error("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}

		// Work / inspect body. You may even modify it!
		// And now set a new body, which will simulate the same data we read:
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		// Create a response wrapper:
		mrw := &MyResponseWriter{
			ResponseWriter: w,
			buf:            &bytes.Buffer{},
		}

		next.ServeHTTP(mrw, r)

		if err = transform(mrw); err != nil {
			log.Error("cannot transform response body, err: %v", err)
		}
		logger(r, mrw, body)

		// Now inspect response, and finally send it out:
		// (You can also modify it before sending it out!)
		if _, err = io.Copy(w, mrw.buf); err != nil {
			log.Error("Failed to send out response: %v", err)
		}
	})
}

func transform(writer *MyResponseWriter) error {
	resStr := writer.buf.String()
	if !strings.HasPrefix(resStr, "<!DOCTYPE html>") && !strings.HasPrefix(resStr, "{") {
		writer.buf.Reset()
		_, _ = writer.Write(models.NewErrorResponse("error_internal", resStr))
	}

	return nil
}

func logger(r *http.Request, writer *MyResponseWriter, body []byte) {
	logMsg := make(map[string]any)
	reqMsg := map[string]any{
		"method":    r.Method,
		"url":       r.URL.Path,
		"parameter": r.URL.Query(),
	}

	if len(body) > 0 {
		bMap := make(map[string]any)
		if err := json.Unmarshal(body, &bMap); err == nil {
			reqMsg["body"] = bMap
		}
	}
	logMsg["request"] = reqMsg

	// response
	resMsg := map[string]any{
		"status": writer.statusCode,
	}
	if len(body) > 0 {
		bMap := make(map[string]any)
		if err := json.Unmarshal([]byte(writer.buf.String()), &bMap); err == nil {
			resMsg["body"] = bMap
		}
	}
	logMsg["response"] = resMsg

	// print
	var buffer bytes.Buffer
	if err := prettyEncode(logMsg, &buffer); err == nil {
		log.Info(buffer.String())
	}
}

func prettyEncode(data interface{}, out io.Writer) error {
	enc := json.NewEncoder(out)
	enc.SetIndent("", "")
	enc.SetEscapeHTML(false)
	if err := enc.Encode(data); err != nil {
		return err
	}
	return nil
}
