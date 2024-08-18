package utils

import (
	"backend-starter-project/domain/interfaces"
	"golang.org/x/crypto/bcrypt"
)

// passwordService implements the PasswordService interface.
type passwordService struct{}

// NewPasswordService creates a new instance of PasswordService.
func NewPasswordService() interfaces.PasswordService {
    return &passwordService{}
}

// HashPassword hashes a plain-text password using bcrypt.
func (s *passwordService) HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

// ComparePassword compares a hashed password with a plain-text password.
func (s *passwordService) ComparePassword(hashedPassword, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
