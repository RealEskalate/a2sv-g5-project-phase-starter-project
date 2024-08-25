package controllers

import (
	"errors"
	"regexp"

	"aait.backend.g10/usecases/dto"
)

type UserValidation struct {
}

func ValidateEmail(email string) error {
	// Regular expression for validating an email
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		return errors.New("invalid email address")
	}
	return nil
}

func ValidateFullName(fullName string) error {
	if len(fullName) < 2 || len(fullName) > 50 {
		return errors.New("full name must be between 2 and 50 characters")
	}

	// Regular expression for validating full name
	fullNameRegex := `^[a-zA-Z\s]+$`
	re := regexp.MustCompile(fullNameRegex)
	if !re.MatchString(fullName) {
		return errors.New("full name can only contain letters and spaces")
	}
	return nil
}

func ValidatePassword(password string) error {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	if len(password) >= 8 {
		hasMinLen = true
	}

	for _, char := range password {
		switch {
		case char >= 'A' && char <= 'Z':
			hasUpper = true
		case char >= 'a' && char <= 'z':
			hasLower = true
		case char >= '0' && char <= '9':
			hasNumber = true
		case char == '!' || char == '@' || char == '#' || char == '$' || char == '%' || char == '^' || char == '&' || char == '*' || char == '(' || char == ')' || char == '-' || char == '_':
			hasSpecial = true
		}
	}

	if !hasMinLen || !hasUpper || !hasLower || !hasNumber || !hasSpecial {
		return errors.New("password must be at least 8 characters long and include upper and lower case letters, numbers, and special characters")
	}

	return nil
}

func (v *UserValidation) ValidateUser(user *dto.RegisterUserDTO) error {
	if err := ValidateEmail(user.Email); err != nil {
		return err
	}

	if err := ValidateFullName(user.FullName); err != nil {
		return err
	}

	if err := ValidatePassword(user.Password); err != nil {
		return err
	}

	return nil
}
