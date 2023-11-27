package logger

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

func RequestResponseLogger(h http.Handler) http.Handler {
	logFunc := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		uri := r.RequestURI
		method := r.Method

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
			zap.Duration("duration", duration),
		)
		Log.Info("Sent response",
			zap.Int("status", responseData.status),
			zap.Int("response size", responseData.size),
		)

	}
	return http.HandlerFunc(logFunc)
}
