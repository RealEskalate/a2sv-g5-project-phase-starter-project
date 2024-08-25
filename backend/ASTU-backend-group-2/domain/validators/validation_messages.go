package validators

import "github.com/go-playground/validator/v10"

func ValidationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Value is too short"
	case "max":
		return "Value is too long"
	case "len":
		return "Value must be of specific length"
	case "StrongPassword":
		return "Password must contain at least one uppercase letter, one lowercase letter, one number, and one special character"
	default:
		return "Invalid value"
	}
}
