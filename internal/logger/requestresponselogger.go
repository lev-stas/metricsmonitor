package logger

import (
	"bytes"
	"encoding/json"
	"github.com/lev-stas/metricsmonitor.git/internal/datamodels"
	"go.uber.org/zap"
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
			Log.Error("Can not read request body", zap.Error(err))
		}

		r.Body = io.NopCloser(bytes.NewReader(body))

		var requestBody datamodels.Metric
		if err := json.Unmarshal(body, &requestBody); err != nil {
			Log.Error("Error decoding JSON body", zap.Error(err))
		}

		responseData := &responseData{
			status: 0,
			size:   0,
		}

		lw := logResponseWriter{
			ResponseWriter: w,
			responseData:   responseData,
		}

		h.ServeHTTP(&lw, r)

		duration := time.Since(start)

		Log.Info("Incoming request",
			zap.String("URI", uri),
			zap.String("method", method),
			zap.Any("headers", r.Header),
			zap.Any("body", requestBody),
			zap.Duration("duration", duration),
		)
		Log.Info("Sent response",
			zap.Int("status", responseData.status),
			zap.Int("response size", responseData.size),
		)

	}
	return http.HandlerFunc(logFunc)
}
