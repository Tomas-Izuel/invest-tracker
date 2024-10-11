package errors

import (
	"fmt"
	"net/http"
)

// ValidationError struct for detailed validation error messages
type ValidationError struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Fields  map[string]string `json:"fields"` // Field-specific validation messages
}

// NewValidationError creates a new validation error with field-level details
func NewValidationError(fields map[string]string) *ValidationError {
	return &ValidationError{
		Code:    http.StatusBadRequest,
		Message: "validation error",
		Fields:  fields,
	}
}

// Error implements the error interface for ValidationError
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %v", e.Fields)
}

// Predefined validation errors for common DTO fields
var (
	ErrInvalidName     = "invalid or missing name"
	ErrInvalidPeriod   = "invalid or missing period"
	ErrInvalidEndpoint = "invalid or missing endpoint URL"
	ErrInvalidUserID   = "invalid or missing user ID"
	ErrInvalidStock    = "invalid or missing stock"
	ErrInvalidCode     = "invalid or missing code"
)

// Field-specific validation helper
func NewFieldError(field, message string) map[string]string {
	return map[string]string{field: message}
}
