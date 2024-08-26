package infrastructure

import (
	"crypto/sha256"
	"fmt"
	"time"
	"unicode"

	"github.com/dchest/passwordreset"
	"golang.org/x/crypto/bcrypt"
)

var hashedEmail [32]byte

// compares the inputted password from the existing hash
func PasswordComparator(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// hashes the password with a SHA-256 encryption
func PasswordHasher(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// Handels forgot password
func ForgotPasswordHandler(email string) error {
	secretKey := DotEnvLoader("Reset_Password")

	// Generate a token valid for 1 hour
	hashedEmail = sha256.Sum256([]byte(email))

	emailConfig, _ := NewEmailConfig()
	emailserv := NewEmailService(emailConfig)

	token := passwordreset.NewToken(email, time.Hour*1, hashedEmail[:], []byte(secretKey))
	if err := emailserv.SendResetEmail(email, token); err != nil {
		return err
	}
	return nil
}

func VerifyToken(token string) (string, error) {
	secretKey := DotEnvLoader("Reset_Password")

	// Verify the token
	email, err := passwordreset.VerifyToken(token, func(email string) ([]byte, error) {
		hashedEmail := sha256.Sum256([]byte(email))
		return hashedEmail[:], nil
	}, []byte(secretKey))

	if err != nil {
		fmt.Printf("Token verification failed")
		return "", err
	}

	return email, nil
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

func UserVerification(email string) error {
	secretKey := DotEnvLoader("Reset_Password")

	// Generate a token valid for 1 hour
	hashedEmail = sha256.Sum256([]byte(email))
	emailConfig, _ := NewEmailConfig()
	emailserv := NewEmailService(emailConfig)
	token := passwordreset.NewToken(email, time.Hour*1, hashedEmail[:], []byte(secretKey))
	if err := emailserv.SendVerificationEmail(email, token); err != nil {
		return err
	}
	return nil
}
