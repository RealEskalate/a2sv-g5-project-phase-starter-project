package cryptography

import (
	"blog_api/domain"

	"golang.org/x/crypto/bcrypt"
)

// HashingService provides an interface to interact with hashing functions
type HashingService struct{}

// NewHashingService creates a new HashingService
func NewHashingService() *HashingService {
	return &HashingService{}
}

// Accepts a string and hashes it using the bcrypt
func (s *HashingService) HashString(password string) (string, domain.CodedError) {
	hashedPwd, hashErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if hashErr != nil {
		return "", domain.NewError("Internal server error: "+hashErr.Error(), domain.ERR_INTERNAL_SERVER)
	}

	return string(hashedPwd), nil
}

// Accepts a hashed string and a plaintext string and validates the plaintext string
func (s *HashingService) ValidateHashedString(hashedString string, plaintextString string) domain.CodedError {
	compErr := bcrypt.CompareHashAndPassword([]byte(hashedString), []byte(plaintextString))
	if compErr != nil {
		return domain.NewError("Invalid signature", domain.ERR_UNAUTHORIZED)
	}

	return nil
}
