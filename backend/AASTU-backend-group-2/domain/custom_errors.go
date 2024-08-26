package domain

import (
	"errors"
	"net/http"
)

// AppError is a custom error type that includes an error message and an associated HTTP status code.
type AppError struct {
	message    string
	statusCode int
	err        error
}

// Status returns the HTTP status code associated with the error.
func (e *AppError) Status() int {
	return e.statusCode
}

// Message returns the error message.
func (e *AppError) Message() string {
	return e.message
}

// Unwrap allows the underlying error to be accessible.
func (e *AppError) Unwrap() error {
	return e.err
}

// NewAppError creates a new AppError.
func NewAppError(message string, statusCode int, err error) *AppError {
	return &AppError{
		message:    message,
		statusCode: statusCode,
		err:        err,
	}
}

var (
	ErrBlogNotFound            = NewAppError("blog not found", http.StatusNotFound, errors.New("blog not found"))
	ErrCommentNotFound         = NewAppError("comment not found", http.StatusNotFound, errors.New("comment not found"))
	ErrInvalidObjectID         = NewAppError("invalid object ID", http.StatusBadRequest, errors.New("invalid object ID"))
	ErrPermissionDenied        = NewAppError("permission denied", http.StatusForbidden, errors.New("permission denied"))
	ErrInvalidSortParameter    = NewAppError("invalid sort parameter", http.StatusBadRequest, errors.New("invalid sort parameter"))
	ErrInvalidDirectParameter  = NewAppError("invalid direct parameter", http.StatusBadRequest, errors.New("invalid direct parameter"))
	ErrNoBlogsFound            = NewAppError("no blogs found", http.StatusNotFound, errors.New("no blogs found"))
	ErrCommentInsertionFailed  = NewAppError("failed to insert comment", http.StatusInternalServerError, errors.New("comment insertion failed"))
	ErrBlogUpdateFailed        = NewAppError("failed to update the blog", http.StatusInternalServerError, errors.New("blog update failed"))
	ErrBlogDeletionFailed      = NewAppError("failed to delete the blog", http.StatusInternalServerError, errors.New("blog deletion failed"))
	ErrBlogInsertionFailed     = NewAppError("failed to insert blog", http.StatusInternalServerError, errors.New("blog insertion failed"))
	ErrBlogRetrievalFailed     = NewAppError("failed to retrieve blog", http.StatusInternalServerError, errors.New("blog retrieval failed"))
	ErrBlogDecodingFailed      = NewAppError("failed to decode blog", http.StatusInternalServerError, errors.New("blog decoding failed"))
	ErrBlogAggregationFailed   = NewAppError("failed to aggregate blogs", http.StatusInternalServerError, errors.New("blog aggregation failed"))
	ErrBlogCountFailed         = NewAppError("failed to count blogs", http.StatusInternalServerError, errors.New("blog count failed"))
	ErrSessionStartFailed      = NewAppError("failed to start session", http.StatusInternalServerError, errors.New("session start failed"))
	ErrInternalServerError     = NewAppError("internal server error", http.StatusInternalServerError, errors.New("internal server error"))
	ErrNotFound                = NewAppError("resource not found", http.StatusNotFound, errors.New("resource not found"))
	ErrInvalidInput            = NewAppError("invalid input", http.StatusBadRequest, errors.New("invalid input"))
	ErrInappropriateComment    = NewAppError("the comment contains offensive languages", http.StatusBadRequest, errors.New("offensive languages"))
	ErrInappropriateBlog       = NewAppError("the blog's content contains offensive languages", http.StatusBadRequest, errors.New("offensive languages"))
	ErrGeminiAPIKeyMissing     = NewAppError("Gemini API key is missing or invalid", http.StatusInternalServerError, errors.New("missing or invalid API key"))
	ErrGeminiClientCreation    = NewAppError("failed to create Gemini AI client", http.StatusInternalServerError, errors.New("client creation failed"))
	ErrGeminiModelNotAvailable = NewAppError("the requested generative model is not available", http.StatusServiceUnavailable, errors.New("generative model unavailable"))
	ErrGeminiContentGeneration = NewAppError("failed to generate content using Gemini AI", http.StatusInternalServerError, errors.New("content generation failed"))
	ErrGeminiResponseParsing   = NewAppError("failed to parse response from Gemini AI", http.StatusInternalServerError, errors.New("response parsing failed"))
	ErrGeminiSafetyBlock       = NewAppError("the content was blocked due to safety concerns", http.StatusForbidden, errors.New("safety concerns"))
	ErrOffensiveComment        = NewAppError("the comment contains offensive language", http.StatusBadRequest, errors.New("offensive language"))
	ErrOffensiveBlogContent    = NewAppError("the blog post contains offensive language", http.StatusBadRequest, errors.New("offensive language"))
	// ErrCommentInsertionFailed   = NewAppError("failed to insert comment", http.StatusInternalServerError, errors.New("comment insertion failed"))
	ErrCommentUpdateFailed    = NewAppError("failed to update comment", http.StatusInternalServerError, errors.New("comment update failed"))
	ErrCommentDeletionFailed  = NewAppError("failed to delete comment", http.StatusInternalServerError, errors.New("comment deletion failed"))
	ErrCommentRetrievalFailed = NewAppError("failed to retrieve comments", http.StatusInternalServerError, errors.New("comment retrieval failed"))
	// ErrBlogNotFound             = NewAppError("blog not found", http.StatusNotFound, errors.New("blog not found"))
	// ErrBlogUpdateFailed         = NewAppError("failed to update blog comment count", http.StatusInternalServerError, errors.New("blog update failed"))
	ErrMongoDBConnection = NewAppError("failed to connect to MongoDB", http.StatusInternalServerError, errors.New("mongodb connection failed"))
	// ErrInvalidObjectID              = NewAppError("invalid object ID", http.StatusBadRequest, errors.New("invalid object ID"))
	ErrDislikeInsertionFailed  = NewAppError("failed to insert dislike", http.StatusInternalServerError, errors.New("dislike insertion failed"))
	ErrDislikeUpdateFailed     = NewAppError("failed to update dislike count", http.StatusInternalServerError, errors.New("dislike update failed"))
	ErrDislikeDeletionFailed   = NewAppError("failed to delete dislike", http.StatusInternalServerError, errors.New("dislike deletion failed"))
	ErrDislikeAlreadyExists    = NewAppError("user has already disliked the post", http.StatusBadRequest, errors.New("duplicate dislike"))
	ErrDislikeRetrievalFailed  = NewAppError("failed to retrieve dislikes", http.StatusInternalServerError, errors.New("dislike retrieval failed"))
	ErrLikeRemovalFailed       = NewAppError("failed to remove like when adding dislike", http.StatusInternalServerError, errors.New("like removal failed"))
	ErrDislikeBlogUpdateFailed = NewAppError("failed to update blog dislike count", http.StatusInternalServerError, errors.New("blog dislike count update failed"))
	ErrLikeBlogUpdateFailed    = NewAppError("failed to update blog like count", http.StatusInternalServerError, errors.New("blog like count update failed"))

	ErrLikeInsertionFailed  = NewAppError("failed to insert like", http.StatusInternalServerError, errors.New("like insertion failed"))
	ErrLikeUpdateFailed     = NewAppError("failed to update like count", http.StatusInternalServerError, errors.New("like update failed"))
	ErrLikeDeletionFailed   = NewAppError("failed to delete like", http.StatusInternalServerError, errors.New("like deletion failed"))
	ErrLikeAlreadyExists    = NewAppError("user has already liked the post", http.StatusBadRequest, errors.New("duplicate like"))
	ErrLikeRetrievalFailed  = NewAppError("failed to retrieve likes", http.StatusInternalServerError, errors.New("like retrieval failed"))
	ErrDislikeRemovalFailed = NewAppError("failed to remove dislike when adding like", http.StatusInternalServerError, errors.New("dislike removal failed"))
	// ErrLikeBlogUpdateFailed      = NewAppError("failed to update blog like count", http.StatusInternalServerError, errors.New("blog like count update failed"))

	ErrUserUpdateFailed       = NewAppError("failed to update user details", http.StatusInternalServerError, errors.New("user update failed"))
	ErrUserNotFound           = NewAppError("user not found", http.StatusNotFound, errors.New("user not found"))
	ErrUsernameAlreadyExists  = NewAppError("username already exists", http.StatusConflict, errors.New("username conflict"))
	ErrEmailAlreadyExists     = NewAppError("email already exists", http.StatusConflict, errors.New("email conflict"))
	ErrUserRegistrationFailed = NewAppError("user registration failed", http.StatusInternalServerError, errors.New("user registration failed"))
	ErrEmailNotVerified       = NewAppError("email not verified", http.StatusUnauthorized, errors.New("email not verified"))
	ErrPasswordMismatch       = NewAppError("password does not match", http.StatusUnauthorized, errors.New("password mismatch"))
	ErrTokenGenerationFailed  = NewAppError("token generation failed", http.StatusInternalServerError, errors.New("token generation failed"))
	ErrForgotPasswordFailed   = NewAppError("forgot password operation failed", http.StatusInternalServerError, errors.New("forgot password failed"))
	ErrResetPasswordFailed    = NewAppError("password reset failed", http.StatusInternalServerError, errors.New("reset password failed"))
	ErrLogoutFailed           = NewAppError("logout failed", http.StatusInternalServerError, errors.New("logout failed"))
	ErrInvalidUserID          = NewAppError("invalid user ID", http.StatusBadRequest, errors.New("invalid user ID"))
	ErrInvalidToken           = NewAppError("invalid token", http.StatusUnauthorized, errors.New("token is invalid or expired"))
)
