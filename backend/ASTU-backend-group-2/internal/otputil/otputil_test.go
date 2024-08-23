package otputil

import (
	"testing"
)


func TestGenerateOTP(t *testing.T) {
	otpLength := 8
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Call the GenerateOTP function
	otp, err := GenerateOTP()
	if err != nil {
		t.Errorf("GenerateOTP returned an error: %v", err)
	}

	// Check if the generated OTP is of the correct length
	if len(otp) != otpLength {
		t.Errorf("Expected OTP length %d, got %d", otpLength, len(otp))
	}

	// Check if the generated OTP contains only characters from the allowed charset
	for _, char := range otp {
		if !isCharInCharset(char, charset) {
			t.Errorf("Invalid character '%c' in OTP", char)
		}
	}
}

// isCharInCharset checks if a character is within the allowed charset.
func isCharInCharset(char rune, charset string) bool {
	for _, c := range charset {
		if char == c {
			return true
		}
	}
	return false
}
