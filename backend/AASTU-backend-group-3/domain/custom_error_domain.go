package domain

import "net/http"

type CustomError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func New(message string, statusCode int) *CustomError {
	return &CustomError{
		Message:    message,
		StatusCode: statusCode,
	}
}

func (e *CustomError) Error() string {
	return e.Message
}

var (
	// General Errors
	ErrNotFound       = New("resource not found", http.StatusNotFound)
	ErrInternalServer = New("internal server error", http.StatusInternalServerError)
	ErrBadRequest     = New("bad request", http.StatusBadRequest)
	ErrUnauthorized   = New("unauthorized", http.StatusUnauthorized)
	ErrForbidden      = New("forbidden", http.StatusForbidden)

	// Auth-specific Errors
	ErrUserNotFound         = New("user not found", http.StatusNotFound)
	ErrRefreshTokenNotFound = New("refresh token not found", http.StatusNotFound)
	ErrInvalidRefreshToken  = New("invalid refresh token", http.StatusUnauthorized)
	ErrInvalidAccessToken   = New("invalid access token", http.StatusUnauthorized)
	ErrExpiredAccessToken   = New("expired access token", http.StatusUnauthorized)
	ErrExpiredRefreshToken  = New("expired refresh token", http.StatusUnauthorized)
	ErrInvalidToken         = New("invalid token", http.StatusUnauthorized)
	ErrInvalidRole          = New("invalid role", http.StatusUnauthorized)
	ErrInvalidUserID        = New("invalid user id", http.StatusBadRequest)
	ErrInvalidDeviceID      = New("invalid device id", http.StatusBadRequest)
	ErrDeviceNotFound       = New("device not found", http.StatusNotFound)
	ErrInvalidEmail         = New("invalid email", http.StatusBadRequest)
	ErrInvalidPassword      = New("invalid password", http.StatusBadRequest)

	ErrFailedToUpdateUser    = New("failed to update user", http.StatusInternalServerError)
	ErrMissingRequiredFields = New("missing required fields", http.StatusBadRequest)
	ErrInvalidUpdateRequest  = New("invalid update request", http.StatusBadRequest)
	ErrFailedToSendEmail     = New("failed to send email", http.StatusInternalServerError)
	ErrActivationFailed      = New("account activation failed", http.StatusInternalServerError)
	ErrFailedToDeleteUser    = New("failed to delete user", http.StatusInternalServerError)
	ErrFailedToDeleteAccount = New("failed to delete account", http.StatusInternalServerError)
	ErrFailedToUploadImage   = New("failed to upload image", http.StatusInternalServerError)
	ErrFailedToUpdateProfile = New("failed to update profile", http.StatusInternalServerError)

	// Oauth-specific Errors
	ErrFailedToFindOrCreateUser = New("failed to find or create user", http.StatusInternalServerError)
	ErrFailedToGenerateToken    = New("failed to generate token", http.StatusInternalServerError)

	// Blog-specific Errors
	ErrFailedToCreateBlog        = New("failed to create blog", http.StatusInternalServerError)
	ErrBlogNotFound              = New("blog not found", http.StatusNotFound)
	ErrPostNotFound              = New("blog post not found", http.StatusNotFound)
	ErrCommentNotFound           = New("comment not found", http.StatusNotFound)
	ErrFailedToDeleteBlog        = New("failed to delete blog", http.StatusInternalServerError)
	ErrFailedToUpdateBlog        = New("failed to update blog", http.StatusInternalServerError)
	ErrFailedToRetrieveBlogs     = New("failed to retrieve blogs", http.StatusInternalServerError)
	ErrFailedToRetrieveUserBlogs = New("failed to retrieve user blogs", http.StatusInternalServerError)
	ErrReplyNotFound             = New("reply not found", http.StatusNotFound)
	ErrLikeAlreadyExists         = New("like already exists", http.StatusConflict)
	ErrLikeNotFound              = New("like not found", http.StatusNotFound)
	ErrUserAlreadyExists         = New("user already exists", http.StatusConflict)
	ErrInvalidCredentials        = New("invalid credentials", http.StatusUnauthorized)
	ErrInsufficientRights        = New("insufficient rights", http.StatusForbidden)
	ErrAdminRoleRequired         = New("admin role required", http.StatusForbidden)
	ErrUserRoleRequired          = New("user role required", http.StatusForbidden)
	ErrFailedToGetUser 		 = New("failed to get user", http.StatusInternalServerError)
	ErrUserNotVerified		   = New("user not verified", http.StatusForbidden)
	// Comment-specific Errors
	ErrFailedToCreateComment       = New("failed to create comment", http.StatusInternalServerError)
	ErrFailedToUpdateComment       = New("failed to update comment", http.StatusInternalServerError)
	ErrFailedToDeleteComment       = New("failed to delete comment", http.StatusInternalServerError)
	ErrFailedToGetComment          = New("failed to get comment", http.StatusInternalServerError)
	ErrInvalidCommentID            = New("invalid comment ID", http.StatusBadRequest)
	ErrInvalidPaginationParameters = New("invalid pagination parameters", http.StatusBadRequest)
	ErrFailedToCreateReply         = New("failed to create reply", http.StatusInternalServerError)
	ErrFailedToUpdateReply         = New("failed to update reply", http.StatusInternalServerError)
	ErrFailedToDeleteReply         = New("failed to delete reply", http.StatusInternalServerError)
	ErrFailedToGetReplies          = New("failed to get replies", http.StatusInternalServerError)
	ErrInvalidReplyID              = New("invalid reply ID", http.StatusBadRequest)
	ErrFailedToLikeComment         = New("failed to like comment", http.StatusInternalServerError)
	ErrFailedToUnlikeComment       = New("failed to unlike comment", http.StatusInternalServerError)
	ErrFailedToLikeReply           = New("failed to like reply", http.StatusInternalServerError)
	ErrFailedToUnlikeReply         = New("failed to unlike reply", http.StatusInternalServerError)
	ErrFailedToGetComments         = New("failed to get comments", http.StatusInternalServerError)

	// Like-specific Errors
	ErrFailedToLikePost   = New("failed to like post", http.StatusInternalServerError)
	ErrFailedToUnlikePost = New("failed to unlike post", http.StatusInternalServerError)
)
