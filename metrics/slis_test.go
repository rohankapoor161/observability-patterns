package metrics

import (
	"testing"
	"time"
)

func TestNewSLIMetrics(t *testing.T) {
	metrics := NewSLIMetrics("test-service")
	if metrics == nil {
		t.Fatal("expected metrics to be initialized")
	}
}

func TestRecordRequest(t *testing.T) {
	metrics := NewSLIMetrics("test-service")
	
	metrics.RecordRequest(nil, "GET", "/health", "200", 50*time.Millisecond)
	
	// In real test, we'd verify prometheus registry
}
