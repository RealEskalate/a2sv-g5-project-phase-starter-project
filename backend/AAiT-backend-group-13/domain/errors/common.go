
package er

import "fmt"

// IErr is an interface that should be implemented by all custom errors.
// It provides a method to retrieve the type of the error.
type IErr interface {
	// Type returns the type of the error as a string.
	Type() string
}

const (
	// Validation represents a validation error.
	Validation = "Validation"

	// Conflict represents a conflict error.
	Conflict = "Conflict"

	// Unexpected represents an unexpected server error.
	Unexpected = "ServerError"

	// NotFound represents a resource not found error.
	NotFound = "NotFound"

	// Unauthorized represents an error for unauthorized access.
	Unauthorized = "Unauthorized"
)

// Error represents a custom domain error with a type and message.
// It implements the IErr interface.
type Error struct {
	kind    string // The type of the error (e.g., Validation, Conflict).
	Message string // The error message providing details about the error.
}

// Ensure Error implements the IErr interface.
var _ IErr = &Error{}

// It is used internally to construct errors of specific types.
func new(errType, message string) *Error {
	return &Error{kind: errType, Message: message}
}

// Error returns the string representation of the Error.
func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.kind, e.Message)
}

// Type returns the type of the Error.
func (e Error) Type() string {
	return e.kind
}

// NewValidation creates a new validation error with the given message.
func NewValidation(message string) *Error {
	return new(Validation, message)
}

// NewConflict creates a new conflict error with the given message.
func NewConflict(message string) *Error {
	return new(Conflict, message)
}

// NewUnexpected creates a new unexpected server error with the given message.
func NewUnexpected(message string) *Error {
	return new(Unexpected, message)
}

// NewNotFound creates a new not found error with the given message.
func NewNotFound(message string) *Error {
	return new(NotFound, message)
}

// NewUnauthorized creates a new unauthorized error with the given message.
func NewUnauthorized(message string) *Error {
	return new(Unauthorized, message)
}