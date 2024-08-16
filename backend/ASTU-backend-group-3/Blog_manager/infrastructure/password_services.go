package infrastructure

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	mathRand "math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct{}

func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

// HashPassword hashes a password using bcrypt.
func (ps *PasswordService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePasswords compares a hashed password with a plain text password.
func (ps *PasswordService) ComparePasswords(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// GenerateResetToken generates a random reset token.
func (ps *PasswordService) GenerateResetToken() string {
	mathRand.Seed(time.Now().UnixNano())
	token := make([]byte, 20)
	rand.Read(token)
	return hex.EncodeToString(token)
}

// EncodeToken encodes a token using SHA-256.
func (ps *PasswordService) EncodeToken(token string) string {
	hash := sha256.New()
	hash.Write([]byte(token))
	return hex.EncodeToString(hash.Sum(nil))
}
