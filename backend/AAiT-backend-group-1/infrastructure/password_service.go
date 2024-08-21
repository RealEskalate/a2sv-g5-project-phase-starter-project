package infrastructure

import (
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"golang.org/x/crypto/bcrypt"
)

type passwordService struct {
}

func NewPasswordService() domain.PasswordService {
	return &passwordService{}
}

func (service *passwordService) HashPassword(password string) (string, error) {
	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if errHash != nil {
		return "", errHash
	}
	return string(hashedPassword), nil
}

func (service *passwordService) ComparePassword(hashedPassword, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
