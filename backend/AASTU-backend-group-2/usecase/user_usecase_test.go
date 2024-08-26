package usecase_test

import (
	"blog_g2/domain"
	"blog_g2/mocks"
	"blog_g2/usecase"
	"context"
	"time"

	"github.com/stretchr/testify/suite"
)

// UserUseCasetestSuite struct to hold any shared resources or setup for the tests
type UserUseCasetestSuite struct {
	suite.Suite
	mockUserRepository *mocks.UserRepository
	UserUsecase        domain.UserUsecase
}

// SetupTest runs before each test case
func (s *UserUseCasetestSuite) SetupTest() {
	s.mockUserRepository = new(mocks.UserRepository)
	s.UserUsecase = usecase.NewUserUsecase(s.mockUserRepository, time.Second*2)
}

// TearDownTest runs after each test case
func (s *UserUseCasetestSuite) TearDownTest() {
	// Clean up resources if needed
}

// TestRegisterUser tests the RegisterUser method
func (s *UserUseCasetestSuite) TestRegisterUser() {
	user := domain.User{
		UserName: "testuser",
		Email:    "",
		Password: "password",
	}

	s.mockUserRepository.On("RegisterUser", &user).Return(nil)

	err := s.UserUsecase.RegisterUser(context.Background(), &user)

	s.Nil(err)
}

// TestLoginUser tests the LoginUser method
func (s *UserUseCasetestSuite) TestLoginUser() {
	user := domain.User{
		UserName: "testuser",
		Email:    "",
		Password: "password",
	}

	s.mockUserRepository.On("LoginUser", user).Return("my-token", nil)

	_, err := s.UserUsecase.LoginUser(context.Background(), user)

	s.Nil(err)
}

// TestForgotPassword tests the ForgotPassword method
func (s *UserUseCasetestSuite) TestForgotPassword() {
	email := "testuser@gmail.com"

	s.mockUserRepository.On("ForgotPassword", email).Return(nil)

	err := s.UserUsecase.ForgotPassword(context.Background(), email)

	s.Nil(err)
}

// TestLogoutUser tests the LogoutUser method
func (s *UserUseCasetestSuite) TestLogoutUser() {
	uid := "123"

	s.mockUserRepository.On("LogoutUser", uid).Return(nil)

	err := s.UserUsecase.LogoutUser(context.Background(), uid)

	s.Nil(err)
}

// TestPromoteDemoteUser tests the PromoteDemoteUser method
func (s *UserUseCasetestSuite) TestPromoteDemoteUser() {
	userid := "123"
	isAdmin := true

	s.mockUserRepository.On("PromoteDemoteUser", userid, isAdmin).Return(nil)

	err := s.UserUsecase.PromoteDemoteUser(context.Background(), userid, isAdmin)

	s.Nil(err)
}

// TestResetPassword tests the ResetPassword method
func (s *UserUseCasetestSuite) TestResetPassword() {
	token := "123"
	newPassword := "newpassword"

	s.mockUserRepository.On("ResetPassword", token, newPassword).Return(nil)

	err := s.UserUsecase.ResetPassword(context.Background(), token, newPassword)

	s.Nil(err)
}
