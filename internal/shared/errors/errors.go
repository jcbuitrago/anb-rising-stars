package errors

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("field '%s': '%s'", e.Field, e.Msg)
}

func NewValidationError(field, msg string) error {
	return &ValidationError{Field: field, Msg: msg}
}

func IsValidationError(err error) bool {
	var e *ValidationError
	return errors.As(err, &e)
}

type NotFoundError struct {
	Resource string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("resource '%s' not found", e.Resource)
}

func NewNotFoundError(resource string) error {
	return &NotFoundError{Resource: resource}
}

func IsNotFoundError(err error) bool {
	var e *NotFoundError
	return errors.As(err, &e)
}
