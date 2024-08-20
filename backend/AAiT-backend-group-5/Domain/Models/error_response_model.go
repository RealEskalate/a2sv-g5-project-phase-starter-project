package models

import (
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Code    int
	Message string
}

func (e *ErrorResponse) Error() string {

	return fmt.Errorf("message: %s", e.Message).Error()
}

func BadRequest(msg string) *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: msg,
	}
}

func ErrTokenNotFound(msg string) *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusUnauthorized,
		Message: msg,
	}
}

func Unauthorized(msg string) *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusUnauthorized,
		Message: msg,
	}
}

func Forbidden(msg string) *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusForbidden,
		Message: msg,
	}
}

func NotFound(msg string) *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusNotFound,
		Message: msg,
	}
}

func InternalServerError(msg string) *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: msg,
	}
}

func Conflict(msg string) *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusConflict,
		Message: msg,
	}
}

func UnprocessableEntity(msg string) *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusUnprocessableEntity,
		Message: msg,
	}
}


func Nil() *ErrorResponse{
	return nil
}
