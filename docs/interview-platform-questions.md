# Platform Engineer Interview Questions

## System Design

### Design a deployment system
- How do you ensure safety?
- What metrics matter?
- How do you handle rollbacks?

### Design an observability platform
- What data do you collect?
- How do you ensure low latency?
- How do you handle cardinality?

## Troubleshooting

### Service is slow
1. Check latency distribution (not just average)
2. Identify which endpoint(s) affected
3. Check dependencies (database, caches)
4. Look for resource constraints
5. Review recent deployments

### High error rate
1. Identify error types (4xx vs 5xx)
2. Check upstream/downstream services
3. Review logs for error patterns
4. Correlate with deployments
5. Check for infrastructure issues

## Behavioral

### Tell me about a time you prevented an incident
- What signals did you see?
- What action did you take?
- What was the outcome?

### How do you handle on-call stress?
- Preparation and runbooks
- Escalation procedures
- Post-incident reviews
- Sustainable rotation
