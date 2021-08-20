package apiserver

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func (s *server) addJsonHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(
			logrus.Fields{
				"remote_addr": r.RemoteAddr,
			})

		logger.Infof("Handle %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &LoggerResponseWriter{
			w, http.StatusOK,
		}
		next.ServeHTTP(rw, r)

		logger.Infof(
			"Completed handler with %d %s in time %v",
			rw.code,
			http.StatusText(rw.code),
			time.Since(start),
		)
	})
}