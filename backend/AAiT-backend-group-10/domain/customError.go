package domain

type CustomError struct {
	Message    string
	StatusCode int
}

// Error implements error.
func (c *CustomError) Error() string {
	panic("unimplemented")
}

func NewCustomError(message string, statusCode int) *CustomError {
	return &CustomError{
		Message:    message,
		StatusCode: statusCode,
	}
}

// Blog-related errors
var (
	ErrBlogNotFound           = NewCustomError("Blog not found", 404)
	ErrDuplicateBlogID        = NewCustomError("Duplicate blog ID", 409)
	ErrBlogInsertFailed       = NewCustomError("Failed to insert blog", 500)
	ErrBlogUpdateFailed       = NewCustomError("Failed to update blog", 500)
	ErrBlogDeleteFailed       = NewCustomError("Failed to delete blog", 500)
	ErrBlogFetchFailed        = NewCustomError("Failed to fetch blog(s)", 500)
	ErrBlogCursorDecodeFailed = NewCustomError("Failed to decode blog data", 500)
	ErrBlogCountFailed        = NewCustomError("Failed to get blog count for pagination", 500)
)

// User-related errors
var (
	ErrUserNotFound           = NewCustomError("User not found", 404)
	ErrUserEmailExists        = NewCustomError("Email already exists", 409)
	ErrUserCreationFailed     = NewCustomError("Failed to create user", 500)
	ErrUserUpdateFailed       = NewCustomError("Failed to update user", 500)
	ErrUserPromotionFailed    = NewCustomError("Failed to promote user", 500)
	ErrUserFetchFailed        = NewCustomError("Failed to fetch users", 500)
	ErrUserCursorDecodeFailed = NewCustomError("Failed to decode user data", 500)
)

// auth-related errors
var (
	ErrUserTokenUpdateFailed   = NewCustomError("Failed to update user tokens", 500)
	ErrInvalidCredentials      = NewCustomError("Invalid email or password", 401)
	ErrInvalidToken            = NewCustomError("Invalid token", 401)
	ErrInvalidRefreshToken     = NewCustomError("Invalid refresh token", 401)
	ErrInvalidResetCode        = NewCustomError("Invalid reset code", 400)
	ErrUnAuthorized            = NewCustomError("Unauthorized access", 401)
	ErrUnexpectedSigningMethod = NewCustomError("Unexpected signing method", 500)
)

// Comment-related errors
var (
	ErrCommentNotFound       = NewCustomError("Comment not found", 404)
	ErrCommentCreationFailed = NewCustomError("Failed to create comment", 500)
	ErrCommentUpdateFailed   = NewCustomError("Failed to update comment", 500)
	ErrCommentDeletionFailed = NewCustomError("Failed to delete comment", 500)
	ErrCommentFetchFailed    = NewCustomError("Failed to fetch comments", 500)
)

// Like-related errors
var (
	ErrLikeNotFound         = NewCustomError("Like not found", 404)
	ErrLikeCreationFailed   = NewCustomError("Failed to like the blog", 500)
	ErrLikeUpdateFailed     = NewCustomError("Failed to update like status", 500)
	ErrLikeDeletionFailed   = NewCustomError("Failed to delete like", 500)
	ErrLikeCountFetchFailed = NewCustomError("Failed to fetch like count", 500)
)

// JWT-related errors
var (
	ErrTokenGenerationFailed        = NewCustomError("Failed to generate token", 500)
	ErrRefreshTokenGenerationFailed = NewCustomError("Failed to generate refresh token", 500)
	ErrTokenParsingFailed           = NewCustomError("Failed to parse token", 401)
	ErrResetTokenGenerationFailed   = NewCustomError("Failed to generate reset token", 500)
)
var(
	// Password-related errors
	ErrPasswordHashingFailed = NewCustomError("Failed to hash password", 500)

	// Email-related errors
	ErrEmailSendingFailed = NewCustomError("Failed to send email", 500)
)
