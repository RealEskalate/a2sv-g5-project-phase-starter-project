package utils

import (
	"errors"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

//TODO: API should support developer mode where the errors aren't more technical and verbose and not sugar coated
//TODO: This would help both backend and frontend people debug

// MapErrors maps internal errors to client-facing messages based on their severity.
func MapErrors(err error) string {
	log.Println(err.Error())
	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		return "Resource not found"
	case errors.Is(err, mongo.ErrNilDocument):
		return "Received unexpected nil document from the database"
	case errors.Is(err, mongo.ErrUnacknowledgedWrite):
		return "Write operation not acknowledged by the database"
	case errors.Is(err, http.ErrServerClosed):
		return "Server closed unexpectedly"
	case errors.Is(err, http.ErrAbortHandler):
		return "Request aborted by the handler"
	case errors.Is(err, http.ErrHandlerTimeout):
		return "Handler timed out while processing request"
	case errors.Is(err, http.ErrNotMultipart):
		return "Request is not multipart/form-data"
	case errors.Is(err, http.ErrMissingFile):
		return "Multipart request is missing a file"
	case errors.Is(err, http.ErrMissingBoundary):
		return "Multipart request is missing boundary"
	case errors.Is(err, http.ErrNoCookie):
		return "No cookie found in the request"
	case errors.Is(err, http.ErrNoLocation):
		return "No Location header found in the response"
	default:
		return "An unexpected error occurred, please try again later"
	}
}
