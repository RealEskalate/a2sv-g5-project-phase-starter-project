package mocks

import (
	"blog_project/domain"
	"context"

	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock type for the IUserRepository interface
type MockUserRepository struct {
	mock.Mock
}

// GetAllUsers mocks the GetAllUsers method of the userRepository
func (m *MockUserRepository) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.User), args.Error(1)
}

// GetUserByID mocks the GetUserByID method of the userRepository
func (m *MockUserRepository) GetUserByID(ctx context.Context, id int) (domain.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.User), args.Error(1)
}

// CreateUser mocks the CreateUser method of the userRepository
func (m *MockUserRepository) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(domain.User), args.Error(1)
}

// UpdateUser mocks the UpdateUser method of the userRepository
func (m *MockUserRepository) UpdateUser(ctx context.Context, id int, user domain.User) (domain.User, error) {
	args := m.Called(ctx, id, user)
	return args.Get(0).(domain.User), args.Error(1)
}

// DeleteUser mocks the DeleteUser method of the userRepository
func (m *MockUserRepository) DeleteUser(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// SearchByUsername mocks the SearchByUsername method of the userRepository
func (m *MockUserRepository) SearchByUsername(ctx context.Context, username string) (domain.User, error) {
	args := m.Called(ctx, username)
	return args.Get(0).(domain.User), args.Error(1)
}

// SearchByEmail mocks the SearchByEmail method of the userRepository
func (m *MockUserRepository) SearchByEmail(ctx context.Context, email string) (domain.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(domain.User), args.Error(1)
}

// AddBlog mocks the AddBlog method of the userRepository
func (m *MockUserRepository) AddBlog(ctx context.Context, userID int, blog domain.Blog) (domain.User, error) {
	args := m.Called(ctx, userID, blog)
	return args.Get(0).(domain.User), args.Error(1)
}

// StoreRefreshToken mocks the StoreRefreshToken method of the userRepository
func (m *MockUserRepository) StoreRefreshToken(ctx context.Context, userID int, refreshToken string) error {
	args := m.Called(ctx, userID, refreshToken)
	return args.Error(0)
}

// ValidateRefreshToken mocks the ValidateRefreshToken method of the userRepository
func (m *MockUserRepository) ValidateRefreshToken(ctx context.Context, userID int, refreshToken string) (bool, error) {
	args := m.Called(ctx, userID, refreshToken)
	return args.Bool(0), args.Error(1)
}

// GetRefreshToken mocks the GetRefreshToken method of the userRepository
func (m *MockUserRepository) GetRefreshToken(ctx context.Context, userID int) (string, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(string), args.Error(1)
}
