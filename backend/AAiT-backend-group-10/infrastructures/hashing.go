package infrastructures

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type PwdService struct {
}

func (s *PwdService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (s *PwdService) CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println(err)
	return err == nil
}
