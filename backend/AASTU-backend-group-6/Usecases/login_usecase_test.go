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

func (suite *LoginUsecaseTestSuite) TestSaveActiveUser() {
	suite.Run("TestSuccess", func() {
		suite.mockActiveUserRepo.On("SaveActiveUser", mock.Anything, mock.Anything).Return(nil).Once()
		err := suite.LoginUsecaseTestSuite.SaveAsActiveUser(domain.ActiveUser{}, "refreshtoken", context.Background())
		suite.NoError(err)
		suite.mockActiveUserRepo.AssertExpectations(suite.T())
		//assert if the function is called with the right arguments
		suite.mockActiveUserRepo.AssertCalled(suite.T(), "SaveActiveUser", mock.Anything, mock.Anything)

	})
	suite.Run("TestError", func() {
		suite.mockActiveUserRepo.On("SaveActiveUser", mock.Anything, mock.Anything).Return(errors.New("some error")).Once()
		err := suite.LoginUsecaseTestSuite.SaveAsActiveUser(domain.ActiveUser{}, "refreshtoken", context.Background())
		suite.Error(err)
		suite.mockActiveUserRepo.AssertExpectations(suite.T())
		//assert if the function is called with the right arguments
		suite.mockActiveUserRepo.AssertCalled(suite.T(), "SaveActiveUser", mock.Anything, mock.Anything)
	})
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

func (suite *LoginUsecaseTestSuite) TestCreateAccessToken() {
	suite.Run("TestSuccess", func() {
		suite.mockUserRepo.On("CreateAccessToken", mock.Anything, mock.Anything, mock.Anything).Return("token", nil).Once()
		_, err := suite.LoginUsecaseTestSuite.CreateAccessToken(&domain.User{}, "secret", 1)
		suite.NoError(err)
	
		suite.mockUserRepo.AssertExpectations(suite.T())
		//assert if the function is called with the right arguments
		suite.mockUserRepo.AssertCalled(suite.T(), "CreateAccessToken", mock.Anything, mock.Anything, mock.Anything)
	})
	suite.Run("TestError", func() {
		suite.mockUserRepo.On("CreateAccessToken", mock.Anything, mock.Anything, mock.Anything).Return("", errors.New("some error")).Once()
		_, err := suite.LoginUsecaseTestSuite.CreateAccessToken(&domain.User{}, "secret", 1)
		suite.Error(err)
		suite.mockUserRepo.AssertExpectations(suite.T())
		//assert if the function is called with the right arguments
		suite.mockUserRepo.AssertCalled(suite.T(), "CreateAccessToken", mock.Anything, mock.Anything, mock.Anything)
	})
}

func (suite *LoginUsecaseTestSuite) TestCreateRefreshToken() {
	suite.Run("TestSuccess", func() {
		suite.mockUserRepo.On("CreateRefreshToken", mock.Anything, mock.Anything, mock.Anything).Return("token", nil).Once()
		_, err := suite.LoginUsecaseTestSuite.CreateRefreshToken(&domain.User{}, "secret", 1)
		suite.NoError(err)
		suite.mockUserRepo.AssertExpectations(suite.T())
		//assert if the function is called with the right arguments
		suite.mockUserRepo.AssertCalled(suite.T(), "CreateRefreshToken", mock.Anything, mock.Anything, mock.Anything)
	})
	suite.Run("TestError", func() {
		suite.mockUserRepo.On("CreateRefreshToken", mock.Anything, mock.Anything, mock.Anything).Return("", errors.New("some error")).Once()
		_, err := suite.LoginUsecaseTestSuite.CreateRefreshToken(&domain.User{}, "secret", 1)
		suite.Error(err)
		suite.mockUserRepo.AssertExpectations(suite.T())
		//assert if the function is called with the right arguments
		suite.mockUserRepo.AssertCalled(suite.T(), "CreateRefreshToken", mock.Anything, mock.Anything, mock.Anything)
	})
}

func TestLogintUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(LogoutUsecaseTestSuite))
}
