package checker

import (
	"errors"
	"unicode"
)

// isValidPassword checks if the password is valid
func IsValidPassword(password string) error {
	var hasUpper, hasLower, hasNumber, hasSpecial bool
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return errors.New("password must include at least one uppercase letter")
	}
	if !hasLower {
		return errors.New("password must include at least one lowercase letter")
	}
	if !hasNumber {
		return errors.New("password must include at least one number")
	}
	if !hasSpecial {
		return errors.New("password must include at least one special character")
	}

	return nil
}
