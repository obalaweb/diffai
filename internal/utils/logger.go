package utils

import (
	"log"
	"os"
)

// Logger provides structured logging functionality
type Logger struct {
	*log.Logger
	verbose bool
}

// NewLogger creates a new logger instance
func NewLogger(verbose bool) *Logger {
	return &Logger{
		Logger:  log.New(os.Stdout, "", log.LstdFlags),
		verbose: verbose,
	}
}

// Info logs an info message
func (l *Logger) Info(format string, v ...interface{}) {
	l.Printf("[INFO] "+format, v...)
}

// Debug logs a debug message (only if verbose is enabled)
func (l *Logger) Debug(format string, v ...interface{}) {
	if l.verbose {
		l.Printf("[DEBUG] "+format, v...)
	}
}

// Error logs an error message
func (l *Logger) Error(format string, v ...interface{}) {
	l.Printf("[ERROR] "+format, v...)
}

// Warn logs a warning message
func (l *Logger) Warn(format string, v ...interface{}) {
	l.Printf("[WARN] "+format, v...)
}

// Fatal logs a fatal message and exits
func (l *Logger) Fatal(format string, v ...interface{}) {
	l.Printf("[FATAL] "+format, v...)
	os.Exit(1)
}
