package config

import (
	"errors"
	"net/http"
)

var (
	ErrInvalidToken          = errors.New("token is invalid")
	ErrInvalidUsernameLength = errors.New("username must be between 3 and 30 characters")
	ErrInvalidUsernameChar   = errors.New("username can only contain alphanumeric characters, underscores, or hyphens")
	ErrInvalidEmailLength    = errors.New("email must be between 3 and 320 characters")
	ErrInvalidEmailFormat    = errors.New("invalid email address")
	ErrInvalidPasswordLength = errors.New("password must be at least 8 characters")
	ErrPasswordNoUppercase   = errors.New("password must contain at least one uppercase letter")
	ErrPasswordNoLowercase   = errors.New("password must contain at least one lowercase letter")
	ErrPasswordNoNumber      = errors.New("password must contain at least one number")
	ErrPasswordNoSpecialChar = errors.New("password must contain at least one special character")
	ErrUserNotVerified       = errors.New("user is not verified")
	ErrIncorrectPassword     = errors.New("incorrect username or password")
	ErrUserCantPromote       = errors.New("only root user can promote pr demote other users")
	ErrAlreadyAdmin          = errors.New("user is already an admin")
	ErrAlreadyUser           = errors.New("user is already a regular user")
	ErrUpdateRole            = errors.New("role cannot be updated")
	ErrUpdateJoined          = errors.New("joined date cannot be updated")
	ErrAlreadyVerified       = errors.New("user is already verified")
	ErrRootAlreadyExists     = errors.New("root user already exists")
	ErrUsernameEmailExists   = errors.New("username or email already exists")
	ErrUserNotFound          = errors.New("user not found")
	ErrTokenNotFound         = errors.New("token not found")
	ErrTokenBlacklisted      = errors.New("token is blacklisted and is no longer valid")
	ErrUserNotLoggedIn       = errors.New("user is not logged in")
	ErrStateExpired          = errors.New("state has expired")
	ErrBlogNotFound          = errors.New("blog not found")
	ErrLikeNotFound          = errors.New("like not found")
	ErrCommentNotFound       = errors.New("comment not found")
	ErrBlogOrLikeNotFound    = errors.New("blog or like not found")
	ErrOnlyAuthorOrAdminDel  = errors.New("only author or admin can delete the blog")
	ErrOnlyAuthorUpdates     = errors.New("only author can update the blog")
	ErrUserCantBePromoted    = errors.New("user cannot be promoted")
	ErrSamePassword          = errors.New("old password and new password cannot be the same")
	ErrUserAlreadyVerified   = errors.New("user is already verified")
)

func GetStatusCode(err error) int {
	switch err {
	case nil:
		return http.StatusOK
	case ErrInvalidToken, ErrUserNotVerified, ErrIncorrectPassword, ErrTokenBlacklisted, ErrStateExpired:
		return http.StatusUnauthorized
	case ErrInvalidUsernameLength, ErrInvalidUsernameChar, ErrInvalidEmailLength, ErrInvalidEmailFormat, ErrInvalidPasswordLength, ErrPasswordNoUppercase, ErrPasswordNoLowercase, ErrPasswordNoNumber, ErrPasswordNoSpecialChar:
		return http.StatusBadRequest
	case ErrUserCantPromote, ErrAlreadyAdmin, ErrAlreadyUser, ErrUpdateRole, ErrUpdateJoined, ErrOnlyAuthorOrAdminDel, ErrOnlyAuthorUpdates:
		return http.StatusForbidden
	case ErrAlreadyVerified, ErrRootAlreadyExists, ErrUsernameEmailExists, ErrUserNotLoggedIn, ErrSamePassword, ErrUserAlreadyVerified:
		return http.StatusConflict
	case ErrUserNotFound, ErrTokenNotFound, ErrBlogNotFound, ErrLikeNotFound, ErrCommentNotFound, ErrBlogOrLikeNotFound:
		return http.StatusNotFound
	case ErrUserCantBePromoted:
		return http.StatusUnprocessableEntity
	default:
		return http.StatusInternalServerError
	}
}
