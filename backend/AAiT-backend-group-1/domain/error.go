package domain

type CustomError interface {
	Code() int
	Message() string
}

type Error struct {
	code    int
	message string
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Message() string {
	return e.message
}