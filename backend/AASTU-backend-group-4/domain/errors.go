package domain

import "errors"

// Define domain-specific errors
var (
	ErrInvalidCredentials = errors.New("invalid username or password")
)
