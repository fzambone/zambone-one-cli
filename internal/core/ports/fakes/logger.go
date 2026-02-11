package fakes

import "strings"

// CapturingLogger implements ports.Logger by storing messages in slices.
// Used in tests to verify that services logged expected messages.
type CapturingLogger struct {
	Infos  []string
	Errors []string
	Debugs []string
}

// NewCapturingLogger creates a new capturing logger.
func NewCapturingLogger() *CapturingLogger {
	return &CapturingLogger{
		Infos:  make([]string, 0),
		Errors: make([]string, 0),
		Debugs: make([]string, 0),
	}
}

// Info appends an info message
func (l *CapturingLogger) Info(msg string, args ...interface{}) {
	l.Infos = append(l.Infos, msg)
}

// Error appends an error message
func (l *CapturingLogger) Error(msg string, args ...interface{}) {
	l.Errors = append(l.Errors, msg)
}

// Debug appends a debug message
func (l *CapturingLogger) Debug(msg string, args ...interface{}) {
	l.Debugs = append(l.Debugs, msg)
}

// HasInfo checks if any Info message contains the given substring
func (l *CapturingLogger) HasInfo(substring string) bool {
	for _, msg := range l.Infos {
		if strings.Contains(msg, substring) {
			return true
		}
	}

	return false
}

// HasError checks if any Error message contains the given substring
func (l *CapturingLogger) HasError(substring string) bool {
	for _, msg := range l.Errors {
		if strings.Contains(msg, substring) {
			return true
		}
	}

	return false
}

// HasDebug checks if any Debug message contains the given substring
func (l *CapturingLogger) HasDebug(substring string) bool {
	for _, msg := range l.Debugs {
		if strings.Contains(msg, substring) {
			return true
		}
	}

	return false
}
