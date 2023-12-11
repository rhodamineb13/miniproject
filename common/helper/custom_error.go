package helper

import (
	"fmt"
	"net/http"
)

type CustomError struct {
	StatusCode int
	Message    string
}

var (
	ErrRegisterFailed  = NewCustomError(http.StatusBadRequest, "error in registering user")
	ErrParseTimeFormat = NewCustomError(http.StatusBadRequest, "error parsing time format")
	ErrLogin           = NewCustomError(http.StatusBadRequest, "email or password is incorrect")
	ErrAddBook         = NewCustomError(http.StatusBadRequest, "error adding new book")
	ErrBooksEmpty      = NewCustomError(http.StatusOK, "there are no books yet")
	ErrBookNotFound    = NewCustomError(http.StatusNotFound, "book not found")
	ErrCreateToken     = NewCustomError(http.StatusInternalServerError, "create token failed")
)

func NewCustomError(statusCode int, message string) error {
	return &CustomError{
		StatusCode: statusCode,
		Message:    message,
	}
}

func (ce *CustomError) Error() string {
	return fmt.Sprintf("[%d]: %s", ce.StatusCode, ce.Message)
}
