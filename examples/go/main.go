package main

import (
	"context"
	"net/http"
	"time"

	"github.com/rohankapoor161/observability-patterns/logging/structured"
	"github.com/rohankapoor161/observability-patterns/metrics"
)

// ExampleHTTPService demonstrates observability patterns
// in a real HTTP service
type ExampleHTTPService struct {
	logger  *logging.StructuredLogger
	metrics *metrics.SLIMetrics
}

func NewExampleHTTPService() *ExampleHTTPService {
	return &ExampleHTTPService{
		logger:  logging.NewStructuredLogger("example-service"),
		metrics: metrics.NewSLIMetrics("example-service"),
	}
}

func (s *ExampleHTTPService) HandleRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	start := time.Now()
	
	traceID := r.Header.Get("X-Trace-ID")
	if traceID == "" {
		traceID = "gen-" + time.Now().Format("20060102-150405")
	}
	
	logger := s.logger.WithTraceID(traceID)
	logger.Info("handling request")
	
	// Simulate work
	time.Sleep(50 * time.Millisecond)
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok"}`))
	
	duration := time.Since(start)
	logger.Info("request completed")
	s.metrics.RecordRequest(ctx, r.Method, r.URL.Path, "200", duration)
}

func main() {
	svc := NewExampleHTTPService()
	http.HandleFunc("/", svc.HandleRequest)
	
	svc.logger.Info("starting server on :8080")
	http.ListenAndServe(":8080", nil)
}