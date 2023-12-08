package gzipper

import (
	"github.com/lev-stas/metricsmonitor.git/internal/logger"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func GzipMiddleware(h http.Handler) http.Handler {
	compressFunc := func(w http.ResponseWriter, r *http.Request) {
		writer := w

		acceptEncoding := r.Header.Get("Accept-Encoding")
		gzipSupport := strings.Contains(acceptEncoding, "gzip")

		if gzipSupport {
			cw := gzipCompressWriter(w)
			writer = cw
			defer cw.Close()
		}

		contentEncoding := r.Header.Get("Content-Encoding")
		sendsGzip := strings.Contains(contentEncoding, "gzip")

		if sendsGzip {
			cr, err := gzipCompressReader(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				logger.Log.Error("Error during reading compressed data", zap.Error(err))
				return
			}

			r.Body = cr
			defer cr.Close()
		}

		h.ServeHTTP(writer, r)
	}

	return http.HandlerFunc(compressFunc)
}
