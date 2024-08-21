package usecases

import (
	domain "blogs/Domain"
	"blogs/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type LoginUsecaseTestSuite struct {
	suite.Suite
	LoginUsecaseTestSuite domain.LoginUsecase
	mockActiveUserRepo    *mocks.ActiveUserRepository
	mockUserRepo          *mocks.UserRepository
	contextTimeout        time.Duration
}

func (suite *LoginUsecaseTestSuite) SetupTest() {
	suite.mockActiveUserRepo = new(mocks.ActiveUserRepository)
	suite.mockUserRepo = new(mocks.UserRepository)
	suite.contextTimeout = time.Second * 5
	suite.LoginUsecaseTestSuite = NewLoginUsecase(suite.mockUserRepo, suite.mockActiveUserRepo, suite.contextTimeout)
}


func (suite *LoginUsecaseTestSuite) TestGetUserByEmail() {
	suite.Run("TestSuccess", func() {
		suite.mockUserRepo.On("FindUserByEmail", mock.Anything, mock.Anything).Return(domain.User{}, nil).Once()
		_, err := suite.LoginUsecaseTestSuite.GetUserByEmail(context.Background(), "email")
		suite.NoError(err)
		suite.mockUserRepo.AssertExpectations(suite.T())
		//assert if the function is called with the right arguments
		suite.mockUserRepo.AssertCalled(suite.T(), "FindUserByEmail", mock.Anything, mock.Anything)
	})
	suite.Run("TestError", func() {
		suite.mockUserRepo.On("FindUserByEmail", mock.Anything, mock.Anything).Return(domain.User{}, errors.New("some error")).Once()
		_, err := suite.LoginUsecaseTestSuite.GetUserByEmail(context.Background(), "email")
		suite.Error(err)
		suite.mockUserRepo.AssertExpectations(suite.T())
		//assert if the function is called with the right arguments
		suite.mockUserRepo.AssertCalled(suite.T(), "FindUserByEmail", mock.Anything, mock.Anything)
	})
}

func TestLogintUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(LoginUsecaseTestSuite))
}
