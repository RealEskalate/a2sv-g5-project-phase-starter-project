package controller

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func passwordValidation(fl validator.FieldLevel) bool {
    password := fl.Field().String()

    // Check for minimum length, uppercase, lowercase, and special character
    hasMinLen := len(password) >= 8
    hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
    hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
    hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
    hasSpecial := regexp.MustCompile(`[!@#\$%\^&\*\(\)\-_\+=]`).MatchString(password)

    return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}