package mocks


import(
	"context"
	"github.com/stretchr/testify/mock"
	"aait-backend-group4/Domain"
)
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	args := m.Called(c, email)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(c context.Context, id string, update domain.UserUpdate) (domain.User, error) {
	args := m.Called(c, id, update)
	return args.Get(0).(domain.User), args.Error(1)
}