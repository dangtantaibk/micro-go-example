package logger

import (
	"github.com/google/uuid"
)

// NewRequestID returns new request ID as string.
func NewRequestID() string {
	return uuid.New().String()
}
