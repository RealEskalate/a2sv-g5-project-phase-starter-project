package domain


var (
	ERR_FORBIDDEN = "err_forbidden"
	ERR_NOT_FOUND = "err_not_found"
	ERR_INTERNAL_SERVER = "err_internal_server"
	ERR_BAD_REQUEST = "err_bad_request"
	ERR_UNAUTHORIZED = "err_unauthorized"
	ERR_CONFLICT = "err_conflict"
	ERR_INVALID_INPUT = "err_invalid_input"
	ERR_INVALID_CREDENTIALS = "err_invalid_credentials"
)

// an interface that extends the error interface
type CodedError interface {
	error
	GetCode() string
}


type Error struct {
	Message string
	Code    string
}


func NewError(message, code string) *Error {
	return &Error {
		Message: message,
		Code:    code,
	}
}


func (err *Error) Error() string {
	return err.Message
}

func (err *Error) GetCode() string {
	return err.Code
}