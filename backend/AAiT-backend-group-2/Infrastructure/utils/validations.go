package utils

import (
	domain "AAiT-backend-group-2/Domain"
	"net/mail"
	"regexp"
)



func ValidateUser(user domain.User) domain.CodedError {
	if err := ValidateEmail(user.Email); err != nil {
		return err
	}

	if err := validatePassword(user.Password); err != nil {
		return err
	}

	if err := ValidateUsername(user.Username); err != nil {
		return err
	}

	return nil
}

func validatePassword(password string) domain.CodedError {
	if len(password) < 8 {
		return domain.NewError("Password must be at least 8 characters long", domain.ERR_BAD_REQUEST)
	}

	return nil
}

func ValidateEmail(email string) domain.CodedError {
	if email == "" {
		return domain.NewError("Email is required", domain.ERR_BAD_REQUEST)
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return domain.NewError("Invalid email", domain.ERR_BAD_REQUEST)
	}

	return nil
}

func ValidateUsername(username string) domain.CodedError {
	if len(username) < 3 {
		return domain.NewError("Username must be at least 3 characters long", domain.ERR_BAD_REQUEST)
	}

	if len(username) > 20 {
		return domain.NewError("Username must be at most 20 characters long", domain.ERR_BAD_REQUEST)
	}

	re := regexp.MustCompile("^[a-zA-Z0-9_]*$")
	if !re.MatchString(username) {
		return domain.NewError("Username can only contain letters, numbers and underscores", domain.ERR_BAD_REQUEST)
	}

	return nil
}