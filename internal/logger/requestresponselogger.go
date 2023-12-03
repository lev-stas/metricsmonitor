package logger

import (
	"bytes"
	"encoding/json"
	"github.com/lev-stas/metricsmonitor.git/internal/datamodels"
	"io"
	"net/http"
	"time"
)

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

		lw := logResponseWriter{
			ResponseWriter: w,
			responseData:   &responseData{},
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
			"status", responseData.status,
			"response size", responseData.size,
		)

	}
	return http.HandlerFunc(logFunc)
}
