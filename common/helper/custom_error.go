package helper

import (
	"fmt"
)

type CustomError struct {
	StatusCode int
	Message    string
}

func NewCustomError(statusCode int, message string) error {
	return &CustomError{
		StatusCode: statusCode,
		Message:    message,
	}
}

func (ce *CustomError) Error() string {
	return fmt.Sprintf("[%d]: %s", ce.StatusCode, ce.Message)
}
