package services

import (
    "errors"
    "regexp"
)

type IValidationService interface {
    ValidatePassword(password string) (bool, error)
    ValidateEmail(email string) (bool, error)
}

type ValidationService struct{}

func NewValidationService() IValidationService {
    return &ValidationService{}
}

func (v *ValidationService) ValidatePassword(password string) (bool, error) {
    if len(password) < 8 {
        return false, errors.New("password must be at least 8 characters long")
    }

    hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString
    hasLowercase := regexp.MustCompile(`[a-z]`).MatchString
    hasDigit := regexp.MustCompile(`\d`).MatchString
    hasSpecial := regexp.MustCompile(`[!@#~$%^&*(),.?":{}|<>]`).MatchString

    if !hasUppercase(password) {
        return false, errors.New("password must contain at least one uppercase letter")
    }
    if !hasLowercase(password) {
        return false, errors.New("password must contain at least one lowercase letter")
    }
    if !hasDigit(password) {
        return false, errors.New("password must contain at least one digit")
    }
    if !hasSpecial(password) {
        return false, errors.New("password must contain at least one special character")
    }

    return true, nil
}

func (v *ValidationService) ValidateEmail(email string) (bool, error) {
    emailRegex := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
    if !emailRegex.MatchString(email) {
        return false, errors.New("invalid email format")
    }

    return true, nil
}

