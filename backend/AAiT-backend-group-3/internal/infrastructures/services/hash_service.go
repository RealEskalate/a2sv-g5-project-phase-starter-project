package services

import "golang.org/x/crypto/bcrypt"

type IHashService interface{
	HashPassword(password string) (string, error)
	CompareHash(hash, password string) bool
}

type PasswordService struct {}

func NewPasswordService(password string) IHashService{
	return &PasswordService{}
}

func (p_service *PasswordService) HashPassword(password string) (string, error){
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (p_service *PasswordService) CompareHash(hash, password string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}