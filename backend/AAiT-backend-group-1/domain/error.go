package domain

type Error interface {
	StatusCode() int
	Error() string
}

type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) StatusCode() int {
	return e.Code
}

func (e *CustomError) Error() string {
	return e.Message
}
