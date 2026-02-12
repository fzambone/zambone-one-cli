// Package ports defines interfaces for external dependencies.
package ports

// Logger defines operations for logging messages at different levels.
type Logger interface {
	// Info logs an informational message.
	Info(msg string, args ...interface{})

	// Error logs an error message.
	Error(msg string, args ...interface{})

	// Debug logs a debug message.
	Debug(msg string, args ...interface{})
}
