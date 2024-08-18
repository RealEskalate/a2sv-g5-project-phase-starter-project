package infrastructure

import (
	"time"

	"github.com/dchest/passwordreset"
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

// Handels forgot password
func ForgotPasswordHandler(email string) error {
	secretKey := DotEnvLoader("Reset_Password")

	token := passwordreset.NewToken(email, time.Hour*1, []byte(secretKey), []byte(secretKey))
	if err := sendResetEmail(email, token); err != nil {
		return err
	}

	return nil
}

func VerifyToken(token string) (string, error) {
	secretKey := DotEnvLoader("Reset_Password")
	tokenDataFunc := func(email string) ([]byte, error) {
		return []byte(email), nil
	}

	email, err := passwordreset.VerifyToken(token, tokenDataFunc, []byte(secretKey))
	if err != nil {
		return "", err
	}

	return email, nil
}
