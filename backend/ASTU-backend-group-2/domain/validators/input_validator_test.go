package validators

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

//test the models for validation

func TestStrongPassword(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("StrongPassword", StrongPassword)

	tests := []struct {
		password string
		valid    bool
	}{
		{"Password123!", true},  // Valid password
		{"Pass123", false},      // No special character
		{"Password!", false},    // No digit
		{"password123!", false}, // No uppercase letter
		{"PASSWORD123!", false}, // No lowercase letter
		{"P@ssw0rd", true},      // Valid password
		{"123", false},          // Too short
		{"", false},             // Empty string
		{"WeakP@ss", false},     // inValid password
	}

	for _, test := range tests {
		err := validate.Var(test.password, "StrongPassword")
		if (err == nil) != test.valid {
			t.Errorf("Validation failed for password '%s'. Expected valid: %v, got error: %v", test.password, test.valid, err)
		}
	}
}
