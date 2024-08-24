package infrastructures

import (
	"aait.backend.g10/domain"
	"golang.org/x/crypto/bcrypt"
)

type HashingService struct {
}

func (s *HashingService) HashPassword(password string) (string, *domain.CustomError) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", domain.ErrPasswordHashingFailed
	}
	return string(bytes), nil
}

func (s *HashingService) CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
