// utility/prometheus.go
package utility

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal   *prometheus.CounterVec
	httpRequestDuration *prometheus.HistogramVec
	fileUploadsTotal    prometheus.Counter
	fileSizeBytes       *prometheus.HistogramVec
)

// InitMetrics initializes Prometheus metrics
func InitMetrics() {
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)

	fileUploadsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "file_uploads_total",
			Help: "Total number of file uploads",
		},
	)

	fileSizeBytes = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "file_size_bytes",
			Help:    "Size of uploaded files in bytes",
			Buckets: []float64{1024, 1024 * 1024, 10 * 1024 * 1024, 100 * 1024 * 1024},
		},
		[]string{"file_type"},
	)

	// Register metrics
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
	prometheus.MustRegister(fileUploadsTotal)
	prometheus.MustRegister(fileSizeBytes)
}

// PrometheusHandler returns a handler for the /metrics endpoint
func PrometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// RecordHTTPRequest records metrics for an HTTP request
func RecordHTTPRequest(method, endpoint string, status int, duration float64) {
	httpRequestsTotal.WithLabelValues(method, endpoint, fmt.Sprintf("%d", status)).Inc()
	httpRequestDuration.WithLabelValues(method, endpoint).Observe(duration)
}

// RecordFileUpload records metrics for a file upload
func RecordFileUpload(fileType string, sizeBytes float64) {
	fileUploadsTotal.Inc()
	fileSizeBytes.WithLabelValues(fileType).Observe(sizeBytes)
}
