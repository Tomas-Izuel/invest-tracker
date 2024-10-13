package errors

import (
	"fmt"
	"net/http"
	"runtime"
)

// CustomError is a generic struct for application errors
type CustomError struct {
	Code    int    `json:"code"`    // HTTP status code
	Message string `json:"message"` // Error message
	Err     error  `json:"-"`       // Internal error
	File    string `json:"-"`       // File where the error occurred
	Line    int    `json:"-"`       // Line number of the error
	Func    string `json:"-"`       // Function name where the error occurred
}

// New creates a new custom error with file, line, and function info
func New(code int, message string, err error) *CustomError {
	file, line, fn := getCallerInfo(2)
	return &CustomError{
		Code:    code,
		Message: message,
		Err:     err,
		File:    file,
		Line:    line,
		Func:    fn,
	}
}

// Wrap wraps an existing error with additional context, file, line, and function
func Wrap(code int, message string, err error) *CustomError {
	file, line, fn := getCallerInfo(2)
	return &CustomError{
		Code:    code,
		Message: message,
		Err:     err,
		File:    file,
		Line:    line,
		Func:    fn,
	}
}

// Error implements the error interface
func (e *CustomError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s (in %s, %s:%d): %v", e.Message, e.Func, e.File, e.Line, e.Err)
	}
	return fmt.Sprintf("%s (in %s, %s:%d)", e.Message, e.Func, e.File, e.Line)
}

// getCallerInfo returns the file, line, and function name where the error occurred
func getCallerInfo(skip int) (file string, line int, fn string) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "unknown", 0, "unknown"
	}
	fnName := runtime.FuncForPC(pc).Name()
	return file, line, fnName
}

// Predefined common errors
var (
	ErrNotFound            = New(http.StatusNotFound, "resource not found", nil)
	ErrBadRequest          = New(http.StatusBadRequest, "bad request", nil)
	ErrUnauthorized        = New(http.StatusUnauthorized, "unauthorized", nil)
	ErrForbidden           = New(http.StatusForbidden, "forbidden", nil)
	ErrInternalServerError = New(http.StatusInternalServerError, "internal server error", nil)
)

// Is checks if the error is of a specific type
func Is(err, target error) bool {
	if e, ok := err.(*CustomError); ok {
		return e.Message == target.Error()
	}
	return err.Error() == target.Error()
}
