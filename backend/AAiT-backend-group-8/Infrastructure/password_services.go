package infrastructure

import (
	"golang.org/x/crypto/bcrypt"
)

func (infrastructure *Infrastructure) CompareHashAndPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (infrastructure *Infrastructure) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
