package validators

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// since the custom validation is not in the default validator, we need to register it
// before the application starts
func init() {
	if ginValidator, ok := binding.Validator.Engine().(*validator.Validate); ok {
		ginValidator.RegisterValidation("StrongPassword", StrongPassword)
	}
}

type UserValidator struct {
	Validate *validator.Validate
}

func StrongPassword(f1 validator.FieldLevel) bool {
	password := f1.Field().String()
	if len(password) < 6 {
		return false
	}
	var (
		hasUpper   = regexp.MustCompile(`[A-Z]`).MatchString(password)
		hasLower   = regexp.MustCompile(`[a-z]`).MatchString(password)
		hasDigit   = regexp.MustCompile(`\d`).MatchString(password)
		hasSpecial = regexp.MustCompile(`[^A-Za-z0-9]`).MatchString(password)
	)
	return hasDigit && hasLower && hasUpper && hasSpecial
}
