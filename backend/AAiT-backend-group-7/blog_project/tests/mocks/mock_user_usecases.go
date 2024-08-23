package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"blog_project/domain"
)

// MockUserUsecase is a mock type for the UserUsecase interface
type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.User), args.Error(1)
}

func (m *MockUserUsecase) GetUserByID(ctx context.Context, id int) (domain.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserUsecase) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserUsecase) UpdateUser(ctx context.Context, id int, user domain.User) (domain.User, error) {
	args := m.Called(ctx, id, user)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserUsecase) DeleteUser(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockUserUsecase) AddBlog(ctx context.Context, userID int, blog domain.Blog) (domain.User, error) {
	args := m.Called(ctx, userID, blog)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserUsecase) DeleteBlog(ctx context.Context, userID int, blogID int) (domain.User, error) {
	args := m.Called(ctx, userID, blogID)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserUsecase) Login(ctx context.Context, username, password string) (string, string, error) {
	args := m.Called(ctx, username, password)
	return args.String(0), args.String(1), args.Error(2)
}

func (m *MockUserUsecase) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	args := m.Called(ctx, refreshToken)
	return args.String(0), args.Error(1)
}

func (m *MockUserUsecase) ForgetPassword(ctx context.Context, email string) error {
	args := m.Called(ctx, email)
	return args.Error(0)
}

func (m *MockUserUsecase) ResetPassword(ctx context.Context, username, password string) error {
	args := m.Called(ctx, username, password)
	return args.Error(0)
}

func (m *MockUserUsecase) PromoteUser(ctx context.Context, userID int) (domain.User, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserUsecase) DemoteUser(ctx context.Context, userID int) (domain.User, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserUsecase) Logout(ctx context.Context, token string) error {
	args := m.Called(ctx, token)
	return args.Error(0)
}
