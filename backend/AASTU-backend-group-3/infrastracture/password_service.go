package infrastracture

import (
	"golang.org/x/crypto/bcrypt"
)

// PasswordServiceImpl is a struct that implements the PasswordService interface.
type PasswordServiceImpl struct{}

// NewPasswordService creates a new instance of PasswordServiceImpl.
func NewPasswordService() *PasswordServiceImpl {
    return &PasswordServiceImpl{}
}

// HashPassword hashes the given password using bcrypt.
func (p *PasswordServiceImpl) HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

// CheckPasswordHash compares the given password with the stored hash.
func (p *PasswordServiceImpl) CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
