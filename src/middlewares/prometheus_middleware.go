package middlewares

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	TotalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"path"},
	)

	ResponseDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_response_duration_seconds",
			Help:    "Histogram of response durations.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path"},
	)

	LastRequestTime = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "http_last_request_time_seconds",
			Help: "Timestamp of the last HTTP request received.",
		},
		[]string{"path"},
	)

	StatusCodeCounts = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_status_code_counts",
			Help: "Counts of HTTP status codes.",
		},
		[]string{"path", "status"},
	)

	// Mutex to ensure thread safety when updating metrics
	//mutex = &sync.Mutex{}
)

func init() {
	prometheus.MustRegister(TotalRequests)
	prometheus.MustRegister(ResponseDuration)
	prometheus.MustRegister(LastRequestTime)
	prometheus.MustRegister(StatusCodeCounts)
}

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// Call the next handler in the chain
		c.Next()

		// Increment the Prometheus counter with the path label
		TotalRequests.WithLabelValues(c.FullPath()).Inc()

		// Record response duration
		duration := time.Since(startTime).Seconds()
		ResponseDuration.WithLabelValues(c.FullPath()).Observe(duration)

		// Update last request time
		LastRequestTime.WithLabelValues(c.FullPath()).SetToCurrentTime()

		// Update status code counts
		statusCode := c.Writer.Status()
		StatusCodeCounts.WithLabelValues(c.FullPath(), strconv.Itoa(statusCode)).Inc()
	}
}
