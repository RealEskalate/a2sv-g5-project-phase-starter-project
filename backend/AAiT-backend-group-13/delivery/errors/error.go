package errapi

import er "github.com/group13/blog/domain/errors"

// HTTP status codes used in the Error type.
const (
	BadRequest     = 400 // Bad Request
	Conflict       = 409 // Conflict
	ServerError    = 500 // Internal Server Error
	Authentication = 401 // Unauthorized
	Forbidden      = 403 // Forbidden
	NotFound       = 404 // Not Found
)

// Error represents an API error with a status code and message.
type Error struct {
	statusCode int    // HTTP status code for the error
	message    string // Detailed error message
}

// NewBadRequest creates a new Error with a 400 Bad Request status code
// and the provided message.
func NewBadRequest(message string) Error {
	return Error{statusCode: BadRequest, message: message}
}

// NewConflict creates a new Error with a 409 Conflict status code
// and the provided message.
func NewConflict(message string) Error {
	return Error{statusCode: Conflict, message: message}
}

// NewServerError creates a new Error with a 500 Internal Server Error status code
// and the provided message.
func NewServerError(message string) Error {
	return Error{statusCode: ServerError, message: message}
}

// NewAuthentication creates a new Error with a 401 Unauthorized status code
// and the provided message.
func NewAuthentication(message string) Error {
	return Error{statusCode: Authentication, message: message}
}

// NewNotFound creates a new Error with a 404 Not Found status code
// and the provided message.
func NewNotFound(message string) Error {
	return Error{statusCode: NotFound, message: message}
}

// NewForbidden creates a new Error with a 403 Forbidden status code
// and the provided message.
func NewForbidden(message string) Error {
	return Error{statusCode: Forbidden, message: message}
}

// Error returns the error message.
func (e Error) Error() string {
	return e.message
}

// StatusCode returns the HTTP status code for the error.
func (e Error) StatusCode() int {
	return e.statusCode
}

// FromErrDMN converts an errdmn.Error to an errapi.Error.
func FromErrDMN(e *er.Error) Error {
	switch e.Type() {
	case er.Validation:
		return NewBadRequest(e.Message)
	case er.Conflict:
		return NewConflict(e.Message)
	case er.Unexpected:
		return NewServerError(e.Message)
	case er.NotFound:
		return NewNotFound(e.Message)
	case er.Unauthorized:
		return NewAuthentication(e.Message)
	default:
		return NewServerError("unknown error")
	}
}
