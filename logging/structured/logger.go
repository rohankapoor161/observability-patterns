package logging

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// StructuredLogger provides consistent structured logging
// across services with automatic correlation IDs
type StructuredLogger struct {
	logger zerolog.Logger
}

// NewStructuredLogger creates a logger with service context
func NewStructuredLogger(service string) *StructuredLogger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	
	logger := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Str("service", service).
		Logger()
	
	return &StructuredLogger{logger: logger}
}

// WithTraceID adds trace/correlation ID to logger
func (s *StructuredLogger) WithTraceID(traceID string) *StructuredLogger {
	return &StructuredLogger{
		logger: s.logger.With().Str("trace_id", traceID).Logger(),
	}
}

// Info logs info level message
func (s *StructuredLogger) Info(msg string) {
	s.logger.Info().Msg(msg)
}

// Error logs error with details
func (s *StructuredLogger) Error(err error, msg string) {
	s.logger.Error().
		Err(err).
		Str("error_type", "application").
		Msg(msg)
}

// Request logs HTTP request details
func (s *StructuredLogger) Request(method, path string, status int, duration time.Duration) {
	s.logger.Info().
		Str("http_method", method).
		Str("http_path", path).
		Int("http_status", status).
		Dur("duration_ms", duration).
		Msg("http_request")
}

// Debug logs debug messages
func (s *StructuredLogger) Debug(msg string) {
	s.logger.Debug().Msg(msg)
}

// Fatal logs fatal error and exits
func (s *StructuredLogger) Fatal(err error, msg string) {
	s.logger.Fatal().Err(err).Msg(msg)
}

// Default returns the package-level logger
func Default() zerolog.Logger {
	return log.Logger
}