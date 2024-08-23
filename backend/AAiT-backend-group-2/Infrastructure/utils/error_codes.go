package utils

import (
	domain "AAiT-backend-group-2/Domain"
	"net/http"
)

func GetHttpErrorCodes(err domain.CodedError) int {
	switch err.GetCode() {
	case domain.ERR_FORBIDDEN:
		return http.StatusForbidden
	case domain.ERR_NOT_FOUND:
		return http.StatusNotFound
	case domain.ERR_INTERNAL_SERVER:
		return http.StatusInternalServerError
	case domain.ERR_BAD_REQUEST:
		return http.StatusBadRequest
	case domain.ERR_UNAUTHORIZED:
		return http.StatusUnauthorized
	case domain.ERR_CONFLICT:
		return http.StatusConflict
	case domain.ERR_INVALID_INPUT:
		return http.StatusBadRequest
	case domain.ERR_INVALID_CREDENTIALS:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}