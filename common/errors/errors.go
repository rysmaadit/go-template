package errors

import (
	"errors"
	"fmt"
	"strings"
)

type BadRequestError struct {
	message string
}

func (e *BadRequestError) Error() string {
	return e.message
}

func NewBadRequestError(err error) *BadRequestError {
	return &BadRequestError{message: err.Error()}
}

type InternalError struct {
	message string
}

func (e *InternalError) Error() string {
	return e.message
}

func NewInternalError(err error, message string) *InternalError {
	return &InternalError{message: fmt.Sprintf("%s: %s", message, err.Error())}
}

type ExternalError struct {
	message string
}

func (e *ExternalError) Error() string {
	return e.message
}

type UnauthorizedError struct {
	message string
}

func (e *UnauthorizedError) Error() string {
	return e.message
}

func NewUnauthorizedError(message string) *UnauthorizedError {
	return &UnauthorizedError{message: message}
}

type ValidationError struct {
	ValidationErrors map[string]string `json:"validation_errors"`
}

func NewValidationError(validationErrors map[string]string) *ValidationError {
	return &ValidationError{ValidationErrors: validationErrors}
}

func (e *ValidationError) Error() string {
	message := make([]string, 0)
	for _, msg := range e.ValidationErrors {
		message = append(message, fmt.Sprintf("%s", msg))
	}
	return strings.Join(message, ", ")
}

func New(message string) error {
	return errors.New(message)
}
