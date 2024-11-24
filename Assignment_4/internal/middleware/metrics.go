package middleware

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path"},
	)

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of response time for HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)

	errorCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_errors_total",
			Help: "Total number of HTTP errors",
		},
		[]string{"method", "path", "status_code"},
	)
)

func InitMetrics() {
	prometheus.MustRegister(requestCounter, requestDuration, errorCounter)
}

func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Use response wrapper
		rr := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(rr, r)

		duration := time.Since(start).Seconds()
		requestCounter.WithLabelValues(r.Method, r.URL.Path).Inc()
		requestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)

		if rr.statusCode >= 400 {
			errorCounter.WithLabelValues(r.Method, r.URL.Path, http.StatusText(rr.statusCode)).Inc()
		}
	})
}

func MetricsHandler() http.Handler {
	return promhttp.Handler()
}
