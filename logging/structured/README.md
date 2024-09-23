# Structured Logging Pattern

Structured logging converts ad-hoc log lines into queryable data. Instead of parsing strings, you query fields.

## Problem

Traditional logging:
```json
{"msg": "Request failed: connection timeout after 5s to database-db-01", "ts": "2024-11-02T15:04:05Z"}
```

Parsing this requires regex. Regex is fragile.

## Solution

Structured logging:
```json
{
  "msg": "database connection failed",
  "ts": "2024-11-02T15:04:05Z",
  "trace_id": "abc123",
  "service": "payments-api",
  "error_type": "database",
  "error": "connection timeout",
  "host": "database-db-01",
  "timeout_ms": 5000
}
```

Now you can query: `error_type="database" AND service="payments-api"`

## Implementation

### Go (zerolog)

See [logger.go](./logger.go) for full implementation.

Quick start:
```go
import "github.com/rohankapoor161/observability-patterns/logging/structured"

logger := logging.NewStructuredLogger("my-service")
logger.WithTraceID(traceID).Info("processing request")
```

### Python (structlog)

```python
import structlog

logger = structlog.get_logger()
logger.info("processing_request", trace_id=trace_id, user_id=user_id)
```

### TypeScript (pino)

```typescript
import pino from 'pino';

const logger = pino({ name: 'my-service' });
logger.info({ traceId, userId }, 'processing request');
```

## Best Practices

1. **Log at the right level**
   - Debug: Detailed flow (user lookups, cache hits)
   - Info: Business events (request handled, job completed)
   - Warn: Recovery scenarios (retry attempted)
   - Error: Failures (request failed, job errored)
   - Fatal: Can't continue (database unavailable, out of memory)

2. **Include context**
   - trace_id for request correlation
   - user_id for user-specific debugging
   - service name for multi-service logs

3. **Avoid logging sensitive data**
   - No passwords, tokens, PII
   - Hash or mask user emails

4. **Consistent field names**
   - snake_case: `trace_id`, `user_id`
   - Not camelCase: `traceId`, `userId`
   - Not PascalCase: `TraceId`, `UserId`

## See Also

- [Log Sampling](../sampling/README.md) — When to drop logs
- [Correlation IDs](../correlation/README.md) — Threading context through requests
