package hash

import (
	"github.com/group13/blog/domain/i_hash"
	"golang.org/x/crypto/bcrypt"
)

type Service struct{}

// Hash implements ihash.Service.
func (s *Service) Hash(word string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(word), 13)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// Match implements ihash.Service.
func (s *Service) Match(hashedWord string, plainWord string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedWord), []byte(plainWord))
	if err != nil {
		return false, nil
	}
	return true, nil
}

var _ ihash.Service = &Service{}
