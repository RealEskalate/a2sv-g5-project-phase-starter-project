package custom_error

import (
	"errors"
	"net/http"
)

type ErrorResponse struct {
	Error      ErrorMessage              `json:"error,omitempty"`
	Validation []ValidationErrorResponse `json:"validation,omitempty"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func ErrMessage(message error) ErrorResponse {
	return ErrorResponse{
		Error: ErrorMessage{
			Message: message.Error(),
		},
	}
}

func ErrValidation(validation []ValidationErrorResponse) ErrorResponse {
	return ErrorResponse{
		Validation: validation,
	}
}

// ValidationErrorResponse represents a detailed validation error response
type ValidationErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// General Errors
var (
	ErrInvalidToken      = errors.New("Invalid token")
	ErrUserNotFound      = errors.New("User not found")
	ErrUserNotVerified   = errors.New("User is not verified")
	ErrAlreadyVerified   = errors.New("User is already verified")
	ErrUserAlreadyExists = errors.New("User already exists with the given email")

	ErrUserNotActive                 = errors.New("User is not active")
	ErrInvalidPasswordLength         = errors.New("Password must be at least 8 characters long")
	ErrErrorBindingRequest           = errors.New("Error binding request")
	ErrErrorEncryptingPassword       = errors.New("Error encrypting password")
	ErrErrorCreatingUser             = errors.New("Error creating user")
	ErrErrorSendingVerificationEmail = errors.New("Error sending verification email")
	ErrInvalidEmailFormat            = errors.New("Invalid email format")
	EreInvalidRequestBody            = errors.New("Request body cannot be empty")
)

// Token Errors
var (
	ErrStateExpired     = errors.New("State has expired")
	ErrTokenBlacklisted = errors.New("Token is no longer valid")
	ErrTokenNotFound    = errors.New("Token not found")
)

// User Errors
var (
	ErrInvalidUsernameLength = errors.New("Username must be between 3 and 30 characters")
	ErrInvalidUsernameChar   = errors.New("Username can only contain alphanumeric characters, underscores, or hyphens")
	ErrInvalidEmailLength    = errors.New("Email must be between 3 and 320 characters")
	ErrPasswordNoUppercase   = errors.New("Password must contain at least one uppercase letter")
	ErrPasswordNoLowercase   = errors.New("Password must contain at least one lowercase letter")
	ErrPasswordNoNumber      = errors.New("Password must contain at least one number")
	ErrPasswordNoSpecialChar = errors.New("Password must contain at least one special character")
	ErrCredentialsNotValid   = errors.New("No user found with the given credentials")

	ErrErrorUpdatingPassword = errors.New("Error updating password")

	ErrFilteringUsers = errors.New("Error filtering users")

	ErrErrorUpdatingUser = errors.New("Error updating user")
)

// Additional Errors
var (
	ErrErrorExtractingClaims          = errors.New("Error extracting user claims from token")
	ErrErrorActivatingUser            = errors.New("Error activating user")
	ErrErrorCheckingOwnership         = errors.New("Error checking user ownership")
	ErrErrorCreatingVerificationToken = errors.New("Error creating verification token")

	ErrErrorSavingOtp   = errors.New("Error saving OTP")
	ErrErrorGettingOtp  = errors.New("Error getting OTP")
	ErrErrorDeletingOtp = errors.New("Error deleting OTP")
)

// Blog Errors
var (
	ErrBlogNotFound = errors.New("Blog not found")
	ErrIDNotFound   = errors.New("ID not found")
	ErrInvalidID    = errors.New("Invalid ID")

	ErrCreatingBlogs = errors.New("Error creating blog")

	// Comment Errors

	ErrCreatingComment = errors.New("Error creating comment")
	ErrGettingComments = errors.New("Error fetching comments")
	ErrGettingComment  = errors.New("Error fetching comment")
	ErrUpdatingComment = errors.New("Error updating comment")
	ErrDeletingComment = errors.New("Error deleting comment")
)

// Mongo Errors
var (
	ErrErrorFindingUser          = errors.New("Error finding user")
	ErrErrorFindingBlog          = errors.New("Error finding blog")
	ErrErrorFindingUserPosts     = errors.New("Error finding user blogs")
	ErrErrorCountingBlogLikes    = errors.New("Error counting blog likes")
	ErrErrorCountingBlogDisLikes = errors.New("Error counting blog dislikes")
	ErrErrorCountingBlogViews    = errors.New("Error counting blog views")
	ErrErrorCountingUserPosts    = errors.New("Error counting user blogs")
	ErrFilteringBlogs            = errors.New("Error filtering blogs")
	ErrFilteringComments         = errors.New("Error filtering comments")
	ErrUpdatingLikeCount         = errors.New("Error updating like count")
)

var errorStatusMap = map[error]int{
	// General Errors
	ErrInvalidToken:                  http.StatusUnauthorized,
	ErrUserNotFound:                  http.StatusNotFound,
	ErrUserNotVerified:               http.StatusForbidden,
	ErrAlreadyVerified:               http.StatusBadRequest,
	ErrUserAlreadyExists:             http.StatusConflict,
	ErrInvalidPasswordLength:         http.StatusBadRequest,
	ErrErrorBindingRequest:           http.StatusBadRequest,
	ErrErrorEncryptingPassword:       http.StatusInternalServerError,
	ErrErrorCreatingUser:             http.StatusInternalServerError,
	ErrErrorSendingVerificationEmail: http.StatusInternalServerError,
	ErrInvalidEmailFormat:            http.StatusBadRequest,

	// Token Errors
	ErrStateExpired:     http.StatusUnauthorized,
	ErrTokenBlacklisted: http.StatusUnauthorized,
	ErrTokenNotFound:    http.StatusNotFound,

	// User Errors
	ErrInvalidUsernameLength: http.StatusBadRequest,
	ErrInvalidUsernameChar:   http.StatusBadRequest,
	ErrInvalidEmailLength:    http.StatusBadRequest,
	ErrPasswordNoUppercase:   http.StatusBadRequest,
	ErrPasswordNoLowercase:   http.StatusBadRequest,
	ErrPasswordNoNumber:      http.StatusBadRequest,
	ErrPasswordNoSpecialChar: http.StatusBadRequest,
	ErrCredentialsNotValid:   http.StatusUnauthorized,

	ErrErrorUpdatingPassword: http.StatusInternalServerError,
	ErrFilteringUsers:        http.StatusInternalServerError,
	ErrErrorUpdatingUser:     http.StatusInternalServerError,

	// Additional Errors
	ErrErrorExtractingClaims:          http.StatusInternalServerError,
	ErrErrorActivatingUser:            http.StatusInternalServerError,
	ErrErrorCheckingOwnership:         http.StatusInternalServerError,
	ErrErrorCreatingVerificationToken: http.StatusInternalServerError,

	ErrErrorSavingOtp:   http.StatusInternalServerError,
	ErrErrorGettingOtp:  http.StatusInternalServerError,
	ErrErrorDeletingOtp: http.StatusInternalServerError,

	// Blog Errors
	ErrBlogNotFound: http.StatusNotFound,
	ErrIDNotFound:   http.StatusNotFound,
	ErrInvalidID:    http.StatusBadRequest,

	ErrCreatingBlogs: http.StatusInternalServerError,

	// Comment Errors
	ErrCreatingComment: http.StatusInternalServerError,
	ErrGettingComments: http.StatusInternalServerError,
	ErrGettingComment:  http.StatusInternalServerError,
	ErrUpdatingComment: http.StatusInternalServerError,
	ErrDeletingComment: http.StatusInternalServerError,

	// Mongo Errors
	ErrErrorFindingUser:          http.StatusInternalServerError,
	ErrErrorFindingBlog:          http.StatusInternalServerError,
	ErrErrorFindingUserPosts:     http.StatusInternalServerError,
	ErrErrorCountingBlogLikes:    http.StatusInternalServerError,
	ErrErrorCountingBlogDisLikes: http.StatusInternalServerError,
	ErrErrorCountingBlogViews:    http.StatusInternalServerError,
	ErrErrorCountingUserPosts:    http.StatusInternalServerError,
	ErrFilteringBlogs:            http.StatusInternalServerError,
	ErrFilteringComments:         http.StatusInternalServerError,
	ErrUpdatingLikeCount:         http.StatusInternalServerError,
}

func MapErrorToStatusCode(err error) int {
	if statusCode, exists := errorStatusMap[err]; exists {
		return statusCode
	}
	return http.StatusInternalServerError // Default to 500 if error not mapped
}
