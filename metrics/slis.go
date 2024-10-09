package metrics

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// SLIMetrics tracks Service Level Indicators
// SLIs are the "what" - latency, error rate, throughput
type SLIMetrics struct {
	requestDuration *prometheus.HistogramVec
	requestTotal    *prometheus.CounterVec
	requestErrors   *prometheus.CounterVec
}

// NewSLIMetrics creates SLI metrics for a service
func NewSLIMetrics(service string) *SLIMetrics {
	return &SLIMetrics{
		requestDuration: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "http_request_duration_seconds",
				Help:    "HTTP request duration in seconds",
				Buckets: prometheus.DefBuckets,
			},
			[]string{"service", "endpoint", "method", "status"},
		),
		requestTotal: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_requests_total",
				Help: "Total HTTP requests",
			},
			[]string{"service", "endpoint", "method", "status"},
		),
		requestErrors: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_request_errors_total",
				Help: "Total HTTP request errors",
			},
			[]string{"service", "endpoint", "error_type"},
		),
	}
}

// RecordRequest records an HTTP request
func (s *SLIMetrics) RecordRequest(ctx context.Context, method, endpoint, status string, duration time.Duration) {
	s.requestDuration.WithLabelValues(service, endpoint, method, status).Observe(duration.Seconds())
	s.requestTotal.WithLabelValues(service, endpoint, method, status).Inc()
}

// RecordError records an error
func (s *SLIMetrics) RecordError(endpoint, errorType string) {
	s.requestErrors.WithLabelValues(service, endpoint, errorType).Inc()
}

var service string = "default"