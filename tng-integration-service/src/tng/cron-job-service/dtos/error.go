package dtos

import (
	"tng/common/httpcode"
)

// AppError contains code and message of errors
type AppError struct {
	Meta `json:"meta"`
}

// NewAppError build and returns a new Merchant error.
func NewAppError(code int, messages ...string) AppError {
	msg := httpcode.GetHTTPStatusText(code)
	if len(messages) > 0 {
		msg = messages[0]
	}
	return AppError{
		Meta: Meta{
			Code:    code,
			Message: msg,
		},
	}
}

// Error returns the error as a string.
func (e AppError) Error() string {
	return e.Message
}
