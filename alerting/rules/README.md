# Alerting Rules That Don't Suck

Alert fatigue is real. If your alerts fire every day, they're not alerts—they're noise.

## Principles

### 1. Alert on Symptoms, Not Causes

**Bad:** "Database CPU > 90%"  
**Good:** "Payment API latency P99 > 500ms"

Users don't care about database CPU. They care that payments are slow.

### 2. The Three Tiers

- **Critical (Page):** Service is down, users can't complete actions
- **Warning (Slack):** Degradation expected, but service functional
- **Info (Dashboard):** Trends, capacity planning, not actionable now

### 3. Actionability

Every alert must answer: **"What do I do now?"**

If the answer is "wait and see," it's not an alert.

## Example Rules

```yaml
groups:
  - name: service-health
    rules:
      # Critical: High error rate
      - alert: HighErrorRate
        expr: rate(http_requests_total{status=~"5.."}[5m]) > 0.1
        for: 2m
        labels:
          severity: critical
        annotations:
          summary: "High error rate on {{ $labels.service }}"
          description: "Error rate above 10% for 2 minutes"
          runbook_url: "https://wiki/runbooks/high-error-rate"
      
      # Warning: Elevated latency
      - alert: ElevatedLatency
        expr: histogram_quantile(0.99, rate(http_request_duration_seconds_bucket[5m])) > 0.5
        for: 5m
        labels:
          severity: warning
        annotations:
          summary: "Elevated latency on {{ $labels.service }}"
          description: "P99 latency above 500ms for 5 minutes"
```

## Anti-Patterns

❌ **Thresholds on utilization** — "CPU > 80%" will fire constantly  
✅ **Thresholds on impact** — "Request latency > target SLO"

❌ **Alert on every error** — You'll be overwhelmed  
✅ **Alert on error rate** — Context matters

❌ **Multiple alerts for same incident** — Correlation needed  
✅ **One alert per incident** — Group related symptoms

## Rate Limiting

- Max 5 pages/week (if you page more, your thresholds are wrong)
- Max 20 warnings/week
- Info alerts don't notify, they trend
