import pino from 'pino';

interface LoggerConfig {
  service: string;
  level?: string;
}

export class StructuredLogger {
  private logger: pino.Logger;

  constructor(config: LoggerConfig) {
    this.logger = pino({
      name: config.service,
      level: config.level || 'info',
      base: { service: config.service }
    });
  }

  info(msg: string, meta?: Record<string, any>): void {
    this.logger.info(meta, msg);
  }

  error(msg: string, error: Error, meta?: Record<string, any>): void {
    this.logger.error({
      ...meta,
      error_type: error.name,
      error_message: error.message,
      stack: error.stack
    }, msg);
  }

  request(method: string, path: string, status: number, durationMs: number): void {
    this.logger.info({
      http_method: method,
      http_path: path,
      http_status: status,
      duration_ms: durationMs
    }, 'http_request');
  }
}

// Example usage
const log = new StructuredLogger({ service: 'example-service' });
log.info('Starting service', { version: '1.0.0' });
log.request('GET', '/health', 200, 15.2);
