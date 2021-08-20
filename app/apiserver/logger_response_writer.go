package apiserver

import "net/http"

type LoggerResponseWriter struct {
	http.ResponseWriter
	code int
}

func (w *LoggerResponseWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
