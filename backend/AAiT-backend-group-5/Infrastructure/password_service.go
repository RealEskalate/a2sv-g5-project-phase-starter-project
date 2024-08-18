package infrastructure

import (
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	"golang.org/x/crypto/bcrypt"
)

type passwordService struct{}

func NewPasswordService() interfaces.PasswordService {
	return &passwordService{}
}

func (p *passwordService) EncryptPassword(password string) (string, error) {
	cur_pass := []byte(password)
	encryptedPassword, err := bcrypt.GenerateFromPassword(cur_pass, bcrypt.DefaultCost)

	return string(encryptedPassword), err

}

func (p *passwordService) ValidatePassword(password string, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
