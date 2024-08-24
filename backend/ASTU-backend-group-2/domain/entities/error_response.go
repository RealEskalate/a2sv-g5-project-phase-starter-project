package entities

import "errors"

// General Errors
var (
	ErrInvalidToken               = errors.New("token is invalid")
	ErrUserNotFound               = errors.New("user not found")
	ErrUserNotVerified            = errors.New("user is not verified")
	ErrAlreadyVerified            = errors.New("user is already verified")
	ErrUserAlreadyExists          = errors.New("user already exists with the given email")
	ErrInvalidPasswordLength      = errors.New("password must be at least 8 characters long")
	ErrErrorBindingRequest        = errors.New("error binding request")
	ErrErrorEncryptingPassword    = errors.New("error encrypting password")
	ErrErrorCreatingUser          = errors.New("error creating user")
	ErrErrorSendingVerificationEmail = errors.New("error sending verification email")
	ErrInvalidEmailFormat         = errors.New("email is not valid")
)

// Token Errors
var (
	ErrStateExpired               = errors.New("state has expired")
	ErrTokenBlacklisted           = errors.New("token is blacklisted and is no longer valid")
	ErrTokenNotFound              = errors.New("token not found")
)

// User Errors
var (
	ErrInvalidUsernameLength     = errors.New("username must be between 3 and 30 characters")
	ErrInvalidUsernameChar       = errors.New("username can only contain alphanumeric characters, underscores, or hyphens")
	ErrInvalidEmailLength        = errors.New("email must be between 3 and 320 characters")
	ErrPasswordNoUppercase       = errors.New("password must contain at least one uppercase letter")
	ErrPasswordNoLowercase       = errors.New("password must contain at least one lowercase letter")
	ErrPasswordNoNumber          = errors.New("password must contain at least one number")
	ErrPasswordNoSpecialChar     = errors.New("password must contain at least one special character")
)

// Additional Errors
var (
	ErrErrorExtractingClaims     = errors.New("error extracting user claims from token")
	ErrErrorActivatingUser       = errors.New("error activating user")
	ErrErrorCheckingOwnership    = errors.New("error checking user ownership")
	ErrErrorCreatingVerificationToken = errors.New("error creating verification token")
)
