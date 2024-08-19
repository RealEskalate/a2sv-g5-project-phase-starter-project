package infrastructure

import "golang.org/x/crypto/bcrypt"

type PasswprdService struct {
}

func (service *PasswprdService) HashPassword(password string) (string, error) {
	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if errHash != nil {
		return "", errHash
	}
	return string(hashedPassword), nil
}

func (service *PasswprdService) ComparePassword(hashedPassword, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
