# Observability Patterns

A curated collection of reusable patterns for building observable distributed systems. Born from years of debugging production incidents at Google, Stripe, and Vercel.

## Overview

Most observability implementations are cargo-culted from blog posts written by people selling observability tools. This is different—patterns extracted from real production systems, battle-tested through countless 3 AM incidents.

**What this is:**
- Reusable patterns you can adapt to your stack
- Code examples in Go, Python, and TypeScript
- Guidelines for when (and when not) to use each pattern

**What this is not:**
- A vendor pitch for specific tools
- "Just add tracing" cargo culting
- Theory without practice

---

## Patterns

### 1. Structured Logging

Stop printf debugging. Start query debugging.

```go
// Bad
log.Printf("Request failed: %v", err)

// Good
log.Info().
    Str("trace_id", span.TraceID()).
    Str("service", "payments-api").
    Str("endpoint", "/charge").
    Err(err).
    Dur("latency", time.Since(start)).
    Msg("request failed")
```

**When to use:** Always. Seriously. Structure your logs.

**See:** [`/logging/structured/`](./logging/structured/)

---

### 2. Context Propagation

Your request context should survive network hops.

```go
// Middleware extracts and propagates context
func TraceMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        traceID := r.Header.Get("X-Trace-ID")
        ctx := WithTraceID(r.Context(), traceID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
```

**When to use:** Microservices, async workers, queue consumers.

---

### 3. Health Checks That Matter

`/healthz` returning 200 OK doesn't mean healthy.

```go
type HealthChecker struct {
    checks []Check
}

func (h *HealthChecker) Check(ctx context.Context) HealthStatus {
    // Actually verify dependencies
    // Database connectivity
    // Critical service reachability
    // Resource limits
}
```

**When to use:** Load balancers, Kubernetes probes, circuit breakers.

---

### 4. Meaningful Metrics

Not all numbers should be metrics. Choose wisely.

**Counters:** Events that happen (requests, errors)  
**Gauges:** Point-in-time values (queue depth, memory usage)  
**Histograms:** Latency distributions (request duration, processing time)  
**Summaries:** User-perceived latency (don't use for SLIs)

**See:** [`/metrics/meaningful/`](./metrics/meaningful/)

---

### 5. Distributed Tracing Without Madness

Tracing should answer "what happened," not "what might have happened."

Key principles:
- Sample at ingress, not everywhere
- 100% sampling for errors
- Don't trace health checks
- Annotate spans with business context, not just technical

**When to use:** Request flows > 3 services, async processing, queue-based architectures.

---

### 6. Alerting Rules

Alert fatigue kills incident response.

**Good alerts:**
- User-impacting symptoms (error rate, latency)
- Actionable (you can do something about it)
- Urgency-appropriate (page → slack → dashboard)

**Bad alerts:**
- Thresholds on internal metrics
- "Something might break soon"
- Anything that fires > 3x/day routinely

**See:** [`/alerting/rules/`](./alerting/rules/)

---

### 7. Incident Correlation

Multiple alerts = one incident.

```yaml
# Group related alerts
incident_rules:
  - name: "payment-service-down"
    group_by: [service, region]
    matchers:
      - alertname: "HighErrorRate"
        service: "payments-api"
      - alertname: "HighLatency"
        service: "payments-api"
```

**When to use:** Complex systems, microservices, dependency chains.

---

## Repository Structure

```
observability-patterns/
├── README.md                 # This file
├── logging/
│   ├── structured/         # Structured logging examples
│   ├── correlation/        # Request ID propagation
│   └── sampling/           # Log sampling strategies
├── tracing/
│   ├── propagation/        # Context propagation
│   ├── sampling/           # Head-based and tail-based sampling
│   └── annotation/         # Meaningful span annotations
├── metrics/
│   ├── meaningful/         # Choosing the right metric type
│   ├── cardinality/        # Managing high-cardinality metrics
│   └── slis/              # Service Level Indicators
├── alerting/
│   ├── rules/              # Alerting rule patterns
│   ├── correlation/        # Incident grouping
│   └── fatigue/            # Avoiding alert fatigue
└── examples/
    ├── go/                 # Go implementations
    ├── python/             # Python implementations
    └── typescript/         # TypeScript implementations
```

---

## Contributing

These patterns come from production experience, but every environment is different. PRs welcome for:
- Additional language implementations
- Real-world case studies
- Anti-patterns and lessons learned
- Tool-specific examples (OTel, Prometheus, etc.)

**See [CONTRIBUTING.md](./CONTRIBUTING.md)**

---

## License

Apache 2.0 — See [LICENSE](LICENSE)

---

*"You can't fix what you can't see, and you can't see without the right questions."*

— Rohan Kapoor, after debugging one too many "it's slow" tickets.
