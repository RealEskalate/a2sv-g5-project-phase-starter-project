package infrastructure

import (
	"fmt"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

// compares the inputted password from the existing hash
func PasswordComparator(hash string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil
}

// hashes the password with a SHA-256 encryption
func PasswordHasher(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
func PasswordValidator(password string) error {

	var (
		hasMinLen      = false
		hasUpper       = false
		hasLower       = false
		hasNumber      = false
		hasSpecialChar = false
		minLength      = 8
	)

	if len(password) >= minLength {
		hasMinLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecialChar = true
		}
	}

	if !hasMinLen {
		return fmt.Errorf("password must be at least %d characters long", minLength)
	}
	if !hasUpper {
		return fmt.Errorf("password must have at least one uppercase letter")
	}
	if !hasLower {
		return fmt.Errorf("password must have at least one lowercase letter")
	}
	if !hasNumber {
		return fmt.Errorf("password must have at least one digit")
	}
	if !hasSpecialChar {
		return fmt.Errorf("password must have at least one special character")
	}

	return nil

}
