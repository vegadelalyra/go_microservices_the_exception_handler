package models

import "fmt"

const (
	ErrorTypeFatal        = "Fatal"
	ErrorTypeError        = "Error"
	ErrorTypeValidation   = "Validation Error"
	ErrorTypeInfo         = "Info"
	ErrorTypeDebug        = "Debug"
	ErrorTypeUnauthorized = "Unauthorized"
)

type ErrorDetail struct {
	ErrorType    string
	ErrorMessage string
}

func (err *ErrorDetail) Error() string {
	return fmt.Sprintf("ErrorType: %s, Error Message: %s", err.ErrorType, err.ErrorMessage)
}

type Response struct {
	Data    interface{}   `json:"data,omitempty"`
	Status  int           `json:"status"`
	Error   []ErrorDetail `json:"error,omitempty"`
	Message string        `json:"message"`
}
