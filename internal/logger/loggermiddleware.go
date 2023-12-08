package logger

import (
	"bytes"
	"encoding/json"
	"github.com/lev-stas/metricsmonitor.git/internal/datamodels"
	"io"
	"net/http"
	"time"
)

type (
	respData struct {
		status int
		size   int
		body   []byte
	}
	logResponseWriter struct {
		http.ResponseWriter
		responseData *respData
	}
)

func (r *logResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size += size
	r.responseData.body = b
	return size, err
}

func (r *logResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode
}

func RequestResponseLogger(h http.Handler) http.Handler {
	logFunc := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		uri := r.RequestURI
		method := r.Method

		body, err := io.ReadAll(r.Body)
		if err != nil {
			Log.Errorw("Can not read request body", "error", err)
		}

		r.Body = io.NopCloser(bytes.NewReader(body))

		var requestBody datamodels.Metric
		if err = json.Unmarshal(body, &requestBody); err != nil {
			Log.Errorw("Error decoding JSON body", "error", err)
		}

		//responseData := &responseData{
		//	status: 0,
		//	size:   0,
		//}

		lw := logResponseWriter{
			ResponseWriter: w,
			responseData:   &respData{},
		}

		h.ServeHTTP(&lw, r)

		duration := time.Since(start)

		Log.Infow("Incoming request",
			"URI", uri,
			"method", method,
			"headers", r.Header,
			"body", requestBody,
			"duration", duration,
		)
		Log.Infow("Sent response",
			"status", lw.responseData.status,
			"response size", lw.responseData.size,
			"response body", string(lw.responseData.body),
		)

	}
	return http.HandlerFunc(logFunc)
}
