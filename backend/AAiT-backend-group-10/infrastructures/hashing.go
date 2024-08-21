package infrastructures

import (
	"fmt"

	"aait.backend.g10/domain"
	"golang.org/x/crypto/bcrypt"
)

type PwdService struct {
}

func (s *PwdService) HashPassword(password string) (string, *domain.CustomError) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", domain.ErrPasswordHashingFailed
	}
	return string(bytes), nil
}

func (s *PwdService) CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println(err)
	return err == nil
}
