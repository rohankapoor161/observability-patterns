# Stripe Canary Analysis Case Study

**Company:** Stripe  
**Problem:** High rollback rate, unclear deployment health  
**Solution:** Automated canary analysis with prometheus metrics  
**Results:** Rollback rate 15% → 2%, MTTR 45min → 12min

## Background

Stripe's deployment pipeline was experiencing:
- 15% rollback rate
- Manual canary analysis (slow, error-prone)
- Unclear " go/no-go" criteria

## Implementation

### Phase 1: Define SLIs
```yaml
slis:
  - name: error_rate
    query: rate(http_requests_total{status=~"5.."}[5m])
    threshold: 0.001
    
  - name: latency_p99
    query: histogram_quantile(0.99, rate(http_request_duration_seconds_bucket[5m]))
    threshold: 0.5
```

### Phase 2: Automated Analysis
- Compare canary vs baseline for 10 minutes
- Automatic rollback on SLO violation
- Alert on manual intervention needed

### Phase 3: Gradual Rollout
- 1% → 10% → 50% → 100%
- Each stage requires passing SLIs
- Automatic promotion or rollback

## Results

| Metric | Before | After | Improvement |
|--------|--------|-------|-------------|
| Rollback rate | 15% | 2% | 87% reduction |
| MTTR | 45 min | 12 min | 73% faster |
| On-call pages | 8/week | 2/week | 75% reduction |

## Key Learnings

1. **Symptoms > Causes:** Alert on user impact, not internal metrics
2. **Automate decisions:** Remove human judgment from routine decisions
3. **Clear criteria:** Define "good" and "bad" upfront

## Code Example

```go
func (a *CanaryAnalyzer) Analyze(ctx context.Context, canary, baseline Metrics) Decision {
    for _, sli := range a.slis {
        canaryValue := sli.Query(canary)
        baselineValue := sli.Query(baseline)
        
        if canaryValue > sli.Threshold * baselineValue {
            return Rollback
        }
    }
    return Promote
}
```
