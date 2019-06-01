package middlewares

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddlewares(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lrw := NewLoggingResponseWriter(w)
		next.ServeHTTP(lrw, r)
		log.WithFields(log.Fields{
			"host":        r.Host,
			"remote_addr": r.RemoteAddr,
			"proto":       r.Proto,
			"duration_ms": float64(time.Since(start).Nanoseconds()) / (1000000.0),
			"date":        time.Now().Format("2/01/2006-15:04:05.00000"),
			"query":       r.URL.RawQuery,
		}).Info(fmt.Sprintf("%s # %d - %s", r.Method, lrw.statusCode, r.URL.Path))
	})
}
