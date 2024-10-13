package lib

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Convierte los errores del validador en un mapa de campo:mensaje
func MapValidationErrors(err error) map[string]string {
	errorMap := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		errorMap[err.Field()] = fmt.Sprintf("Validation failed on '%s' tag", err.Tag())
	}
	return errorMap
}
