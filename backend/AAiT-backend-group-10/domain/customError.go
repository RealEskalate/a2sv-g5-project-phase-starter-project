package domain

import "net/http"

type CustomError struct {
	Message    string
	StatusCode int
}

// Error implements error.
func (c *CustomError) Error() string {
	return c.Message
}

func NewCustomError(message string, statusCode int) *CustomError {
	return &CustomError{
		Message:    message,
		StatusCode: statusCode,
	}
}

// Blog-related errors
var (
	ErrBlogNotFound           = NewCustomError("Blog not found", http.StatusNotFound)
	ErrDuplicateBlogID        = NewCustomError("Duplicate blog ID", http.StatusConflict)
	ErrBlogInsertFailed       = NewCustomError("Failed to insert blog", http.StatusInternalServerError)
	ErrBlogUpdateFailed       = NewCustomError("Failed to update blog", http.StatusInternalServerError)
	ErrBlogDeleteFailed       = NewCustomError("Failed to delete blog", http.StatusInternalServerError)
	ErrBlogFetchFailed        = NewCustomError("Failed to fetch blog(s)", http.StatusInternalServerError)
	ErrBlogCursorDecodeFailed = NewCustomError("Failed to decode blog data", http.StatusInternalServerError)
	ErrBlogCountFailed        = NewCustomError("Failed to get blog count for pagination", http.StatusInternalServerError)
)

// User-related errors
var (
	ErrUserNotFound           = NewCustomError("User not found", http.StatusNotFound)
	ErrUserEmailExists        = NewCustomError("Email already exists", 409)
	ErrUserCreationFailed     = NewCustomError("Failed to create user", http.StatusInternalServerError)
	ErrUserUpdateFailed       = NewCustomError("Failed to update user", http.StatusInternalServerError)
	ErrUserPromotionFailed    = NewCustomError("Failed to promote user", http.StatusInternalServerError)
	ErrUserFetchFailed        = NewCustomError("Failed to fetch users", http.StatusInternalServerError)
	ErrUserCursorDecodeFailed = NewCustomError("Failed to decode user data", http.StatusInternalServerError)
)

// auth-related errors
var (
	ErrUserTokenUpdateFailed   = NewCustomError("Failed to update user tokens", http.StatusInternalServerError)
	ErrInvalidCredentials      = NewCustomError("Invalid email or password", http.StatusUnauthorized)
	ErrInvalidToken            = NewCustomError("Invalid token", http.StatusUnauthorized)
	ErrInvalidRefreshToken     = NewCustomError("Invalid refresh token", http.StatusUnauthorized)
	ErrInvalidResetCode        = NewCustomError("Invalid reset code", http.StatusBadRequest)
	ErrUnAuthorized            = NewCustomError("Unauthorized access", http.StatusUnauthorized)
	ErrUnexpectedSigningMethod = NewCustomError("Unexpected signing method", http.StatusInternalServerError)
)

// Comment-related errors
var (
	ErrCommentNotFound       = NewCustomError("Comment not found", http.StatusNotFound)
	ErrCommentCreationFailed = NewCustomError("Failed to create comment", http.StatusInternalServerError)
	ErrCommentUpdateFailed   = NewCustomError("Failed to update comment", http.StatusInternalServerError)
	ErrCommentDeletionFailed = NewCustomError("Failed to delete comment", http.StatusInternalServerError)
	ErrCommentFetchFailed    = NewCustomError("Failed to fetch comments", http.StatusInternalServerError)
)

// Like-related errors
var (
	ErrLikeNotFound         = NewCustomError("Like not found", http.StatusNotFound)
	ErrLikeCreationFailed   = NewCustomError("Failed to like the blog", http.StatusInternalServerError)
	ErrLikeUpdateFailed     = NewCustomError("Failed to update like status", http.StatusInternalServerError)
	ErrLikeDeletionFailed   = NewCustomError("Failed to delete like", http.StatusInternalServerError)
	ErrLikeCountFetchFailed = NewCustomError("Failed to fetch like count", http.StatusInternalServerError)
)

// JWT-related errors
var (
	ErrTokenGenerationFailed        = NewCustomError("Failed to generate token", http.StatusInternalServerError)
	ErrRefreshTokenGenerationFailed = NewCustomError("Failed to generate refresh token", http.StatusInternalServerError)
	ErrTokenParsingFailed           = NewCustomError("Invalid token", http.StatusUnauthorized)
	ErrResetTokenGenerationFailed   = NewCustomError("Failed to generate reset token", http.StatusInternalServerError)
)
var (
	// Password-related errors
	ErrPasswordHashingFailed = NewCustomError("Failed to hash password", http.StatusInternalServerError)

	// Email-related errors
	ErrEmailSendingFailed = NewCustomError("Failed to send email", http.StatusInternalServerError)
)

// Ai-related errors
var (
	ErrAiContentGenerationFailed = NewCustomError("Failed to generate content", http.StatusInternalServerError)
)