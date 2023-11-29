package logger

import "net/http"

type (
	responseData struct {
		status int
		size   int
	}
	logResponseWriter struct {
		http.ResponseWriter
		responseData *responseData
	}
)

func (r *logResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size += size
	return size, err
}

func (r *logResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode
}
