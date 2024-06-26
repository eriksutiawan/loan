package errors

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.Code, e.Message)
}

func NewBadRequestError(message string) *Error {
	return &Error{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func NewInternalServerError(message string) *Error {
	return &Error{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func NewEntityNotFound(message string) *Error {
	return &Error{
		Code:    http.StatusNotFound,
		Message: message,
	}
}
