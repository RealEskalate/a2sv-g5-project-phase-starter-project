// Infrastructure/password_service.go
package infrastructure

import "golang.org/x/crypto/bcrypt"

// PasswordService defines the interface for password hashing and comparison.
type PasswordService interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) error
}

// bcryptPasswordService is the implementation of PasswordService using bcrypt.
type bcryptPasswordService struct{}

// NewPasswordService creates a new instance of bcryptPasswordService.
func NewPasswordService() PasswordService {
	return &bcryptPasswordService{}
}

// HashPassword hashes the given password using bcrypt.
func (s *bcryptPasswordService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares a bcrypt hashed password with its possible plaintext equivalent.
func (s *bcryptPasswordService) CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
