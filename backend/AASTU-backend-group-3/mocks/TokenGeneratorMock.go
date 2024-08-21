// mocks/token_generator.go
package mocks

import (
	"group3-blogApi/domain"

	"github.com/stretchr/testify/mock"
)

type TokenGenerator struct {
	mock.Mock
}

func (m *TokenGenerator) GenerateToken(user domain.User) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}

func (m *TokenGenerator) GenerateRefreshToken(user domain.User) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}

func (m *TokenGenerator) RefreshToken(token string) (string, error) {
	args := m.Called(token)
	return args.String(0), args.Error(1)
}
