package domain

/*
These constants are used to represent the error codes that are used to communicate
errors between the different layers of the API.
*/
const (
	ERR_BAD_REQUEST     = "err_bad_request"
	ERR_NOT_FOUND       = "err_not_found"
	ERR_INTERNAL_SERVER = "err_internal_server"
	ERR_UNAUTHORIZED    = "err_unauthorized"
	ERR_FORBIDDEN       = "err_forbidden"
	ERR_CONFLICT        = "err_conflict"
)

// CodedError is an interface that extends the `error` interface
type CodedError interface {
	error
	GetCode() string
}

// Error is a struct that implements the `CodedError` interface
type Error struct {
	Message string
	Code    string
}

// NewError creates and returns a new `Error` instance
func NewError(message string, code string) *Error {
	return &Error{
		Message: message,
		Code:    code,
	}
}

func (err Error) Error() string {
	return err.Message
}

func (err Error) GetCode() string {
	return err.Code
}
