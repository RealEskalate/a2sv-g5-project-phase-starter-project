package domain

type Error interface {
	Code() int
	Message() string
}

type CustomError struct {
	code    int
	message string
}

func (e *CustomError) Code() int {
	return e.code
}

func (e *CustomError) Message() string {
	return e.message
}