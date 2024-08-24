package infrastructure_mocks


import (

	"github.com/stretchr/testify/mock"
)


type MockPasswordService struct {
	mock.Mock
}

func (m *MockPasswordService) HashPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}
