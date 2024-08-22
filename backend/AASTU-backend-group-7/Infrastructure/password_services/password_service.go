package password_services

import (
	"errors"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

var CompareHashAndPasswordCustom = func(hashedPassword string, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil

}

var GenerateFromPasswordCustom = func(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

var CheckPasswordStrength = func(password string) error {
	// Check password length
	if len(password) < 8 {
		return errors.New("password is too short and must be at least 8 characters long")
	}

	// Check password strength
	hasUpperCase := false
	hasLowerCase := false
	hasDigit := false
	hasSpecialChar := false

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpperCase = true
		case unicode.IsLower(char):
			hasLowerCase = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecialChar = true
		}
	}

	if !hasUpperCase {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !hasLowerCase {
		return errors.New("password must contain at least one lowercase letter")
	}
	if !hasDigit {
		return errors.New("password must contain at least one digit")
	}
	if !hasSpecialChar {
		return errors.New("password must contain at least one special character")
	}

	return nil
}
