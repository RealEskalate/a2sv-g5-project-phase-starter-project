package userutil

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

// HashPassword hashes the password

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePassword compares the password with the hashed password

func ComparePassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}

// ValidateEmail validates the email

func ValidateEmail(email string) bool {
	// Define a regular expression for validating an email address
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	// Check if the email matches the regex pattern
	return emailRegex.MatchString(email)
}

// ValidatePassword validates the password

func ValidatePassword(password string) bool {
	return len(password) >= 8
}