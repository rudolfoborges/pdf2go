package pdf2go

import "log"

// Logger is the interface that wraps the Debugf method.
type Logger interface {
	// Debugf logs a debug message.
	Debugf(format string, args ...interface{})

	// Errorf logs an error message.
	Errorf(format string, args ...interface{})
}

// LogLevelDegub is the debug log level.
const (
	LogLevelError = "error"
	LogLevelDegub = "debug"
)

// DefaultLogger is the default logger.
type DefaultLogger struct {
	logLevel string
}

// NewDefaultLogger creates a new DefaultLogger.
func NewDefaultLogger(logLevel string) *DefaultLogger {
	return &DefaultLogger{logLevel: logLevel}
}

// Debugf logs a debug message.
// It logs only if the level is "debug".
func (l *DefaultLogger) Debugf(format string, args ...interface{}) {
	if l.logLevel == LogLevelDegub {
		log.Printf(format, args...)
	}
}

// Errorf logs an error message.
func (l *DefaultLogger) Errorf(format string, args ...interface{}) {
	if l.logLevel == LogLevelDegub || l.logLevel == LogLevelError {
		log.Printf(format, args...)
	}
}
