package utils

import (
	"backend-starter-project/domain/interfaces"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type passwordService struct{}

func NewPasswordService() interfaces.PasswordService {
	return &passwordService{}
}

func (service *passwordService) ComparePassword(hashedPassword, password string) error {
	fmt.Println(hashedPassword, password)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}

func (service *passwordService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
