package interfaces

import "fmt"

type ErrTokenExpired struct {
	Message string
}

func (e *ErrTokenExpired) Error() string {
	return fmt.Sprintf("token expired: %s", e.Message)
}

type ErrTokenInvalid struct {
	Message string
}

func (e *ErrTokenInvalid) Error() string {
	return fmt.Sprintf("token invalid: %s", e.Message)
}

type ErrTokenMissing struct {
	Message string
	Code    int
}

func (e *ErrTokenMissing) Error() string {
	return fmt.Sprintf("token missing: %s", e.Message)
}
