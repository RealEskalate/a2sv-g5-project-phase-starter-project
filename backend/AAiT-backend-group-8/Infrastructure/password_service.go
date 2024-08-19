package infrastructure

import (
	"AAiT-backend-group-8/Domain"

	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct {
}

func NewPasswordService() Domain.IPasswordService {
	return &PasswordService{}
}

func (ps *PasswordService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func (ps *PasswordService) VerifyPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
