package errors

import (
	"fmt"
	"github.com/go-playground/validator"
)

func ValidateErrToMapString(valErr validator.ValidationErrors) map[string]string {
	result := make(map[string]string, len(valErr))
	for _, err := range valErr {
		result[err.Field()] = fmt.Sprintf("%v: %v", err.Field(), err.ActualTag())
	}
	return result
}
