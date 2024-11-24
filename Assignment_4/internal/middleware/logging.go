package middleware

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func InitLogger() {
	logger.SetFormatter(&logrus.JSONFormatter{}) // Use JSON formatting for structured logs
	logger.SetLevel(logrus.InfoLevel)            // Set default log level
	logger.SetReportCaller(true)                 // Include the file and line number in logs
}

// LogEvent logs an application-specific event
func LogEvent(level logrus.Level, message string, fields logrus.Fields) {
	logger.WithFields(fields).Log(level, message)
}

// RequestLoggingMiddleware logs incoming HTTP requests
func RequestLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Use a response wrapper to capture status code
		rr := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(rr, r)

		duration := time.Since(start)
		logger.WithFields(logrus.Fields{
			"method":      r.Method,
			"path":        r.URL.Path,
			"status_code": rr.statusCode,
			"duration_ms": duration.Milliseconds(),
			"remote_addr": r.RemoteAddr,
			"user_agent":  r.UserAgent(),
		}).Info("Handled request")
	})
}

// responseRecorder wraps http.ResponseWriter to capture the status code
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rr *responseRecorder) WriteHeader(code int) {
	rr.statusCode = code
	rr.ResponseWriter.WriteHeader(code)
}
