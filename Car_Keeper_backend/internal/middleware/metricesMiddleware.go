package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path", "method", "status"},
	)

	responseDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_response_duration_seconds",
			Help:    "Histogram of response durations for HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path", "method"},
	)

	statusCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_response_status_total",
			Help: "Total number of HTTP responses by status code",
		},
		[]string{"path", "method", "status_code"},
	)
)

func init() {
	prometheus.MustRegister(requestCounter, responseDuration, statusCounter)
}

// MetricsMiddleware is a Gin middleware that records Prometheus metrics for each request.
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		// After request
		duration := time.Since(start).Seconds()
		statusCode := c.Writer.Status()
		path := c.FullPath() // Better than c.Request.URL.Path for Gin routes with params

		if path == "" {
			// Fallback if route not registered (e.g., 404)
			path = c.Request.URL.Path
		}

		method := c.Request.Method
		statusStr := fmt.Sprintf("%d", statusCode)

		// Update Prometheus metrics
		requestCounter.WithLabelValues(path, method, statusStr).Inc()
		responseDuration.WithLabelValues(path, method).Observe(duration)
		statusCounter.WithLabelValues(path, method, statusStr).Inc()
	}
}
