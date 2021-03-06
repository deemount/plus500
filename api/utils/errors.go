package utils

import (
	"fmt"
	"strings"
)

// Error represents a Plus500 API error
type Error struct {
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("plus500 error: %s", strings.ToLower(e.Message))
}

// APIErrors
var (
	ErrDefaultResponse      = &Error{"unknown error"}
	ErrUnauthorized         = &Error{"authentication required"}
	ErrCurrencyNotSupported = &Error{"currency does not supported"}
	ErrUserAgentInvalid     = &Error{"invalid user-agent"}
	ErrNotFound             = &Error{"not found"}
)
