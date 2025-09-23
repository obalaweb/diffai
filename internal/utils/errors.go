package utils

import (
	"fmt"
	"runtime"
)

// Error represents a custom error with context
type Error struct {
	Message   string
	Operation string
	Err       error
	File      string
	Line      int
}

// Error implements the error interface
func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (at %s:%d)", e.Operation, e.Message, e.File, e.Line)
	}
	return fmt.Sprintf("%s: %s (at %s:%d)", e.Operation, e.Message, e.File, e.Line)
}

// Unwrap returns the underlying error
func (e *Error) Unwrap() error {
	return e.Err
}

// NewError creates a new error with context
func NewError(operation, message string, err error) *Error {
	_, file, line, _ := runtime.Caller(1)
	return &Error{
		Message:   message,
		Operation: operation,
		Err:       err,
		File:      file,
		Line:      line,
	}
}

// WrapError wraps an existing error with context
func WrapError(operation string, err error) *Error {
	_, file, line, _ := runtime.Caller(1)
	return &Error{
		Message:   err.Error(),
		Operation: operation,
		Err:       err,
		File:      file,
		Line:      line,
	}
}
