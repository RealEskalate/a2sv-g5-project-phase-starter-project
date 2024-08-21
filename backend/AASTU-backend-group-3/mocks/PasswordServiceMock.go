// mocks/password_service.go
package mocks

import (
	"github.com/stretchr/testify/mock"
)

type PasswordService struct {
	mock.Mock
}

func (m *PasswordService) HashPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *PasswordService) CheckPasswordHash(password, hash string) bool {
	args := m.Called(password, hash)
	return args.Bool(0)
}
