package tracing

import (
	context
	fmt
)

// TraceContext holds distributed trace information
type TraceContext struct {
	TraceID  string
	SpanID   string
	ParentID string
	Sampled  bool
}

// ContextKey for storing trace context
type ContextKey struct{}

var traceContextKey = ContextKey{}

// WithTraceContext adds trace context to context
func WithTraceContext(ctx context.Context, tc *TraceContext) context.Context {
	return context.WithValue(ctx, traceContextKey, tc)
}

// FromContext extracts trace context
func FromContext(ctx context.Context) (*TraceContext, bool) {
	tc, ok := ctx.Value(traceContextKey).(*TraceContext)
	return tc, ok
}

// Inject serializes trace context for propagation
func (tc *TraceContext) Inject() map[string]string {
	return map[string]string{
		trace-id:  tc.TraceID,
		span-id:   tc.SpanID,
		sampled:   fmt.Sprintf(%t, tc.Sampled),
	}
}
