package infrastructure

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func ComparePasswords(hashedPassword, plainPassword string) error {
	if plainPassword == "" {
		return errors.New("password cannot be empty")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		return err
	}

	return nil
}

func HashPassword(plainPassword string) (string, error) {
	if plainPassword == "" {
		return "", errors.New("password cannot be empty")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil

}
