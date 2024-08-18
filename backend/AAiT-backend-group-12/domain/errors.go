package domain

const (
	ERR_BAD_REQUEST     = "err_bad_request"
	ERR_NOT_FOUND       = "err_not_found"
	ERR_INTERNAL_SERVER = "err_internal_server"
	ERR_UNAUTHORIZED    = "err_unauthorized"
	ERR_FORBIDDEN       = "err_forbidden"
	ERR_CONFLICT        = "err_conflict"
)

type CodedError interface {
	error
	GetCode() string
}

type Error struct {
	Message string
	Code    string
}

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
