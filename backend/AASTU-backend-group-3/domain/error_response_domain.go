package domain

type ErrorResponse struct {
	Message string `json:"message"`
}

type ErrorResponseDomain interface {
	NewErrorResponse(message string) ErrorResponse
}