package middlewares

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var totalHTTPRequestsCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Subsystem: "das",
	Name:      "http_requests",
	Help:      "Total number of http requests",
})

func PrometheusMetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		totalHTTPRequestsCounter.Inc()
		next.ServeHTTP(w, r)
	})
}
