import structlog
import time
from typing import Dict, Any

logger = structlog.get_logger()

class StructuredLogger:
    """Python structured logging example."""
    
    def __init__(self, service: str):
        self.service = service
        self.logger = logger.bind(service=service)
    
    def info(self, msg: str, **kwargs):
        self.logger.info(msg, **kwargs)
    
    def error(self, msg: str, error: Exception, **kwargs):
        self.logger.error(
            msg,
            error_type=type(error).__name__,
            error=str(error),
            **kwargs
        )
    
    def request(self, method: str, path: str, status: int, duration_ms: float):
        self.logger.info(
            "http_request",
            http_method=method,
            http_path=path,
            http_status=status,
            duration_ms=duration_ms
        )

# Example usage
if __name__ == "__main__":
    log = StructuredLogger("example-service")
    log.info("Starting service", version="1.0.0")
    log.request("GET", "/health", 200, 15.2)
