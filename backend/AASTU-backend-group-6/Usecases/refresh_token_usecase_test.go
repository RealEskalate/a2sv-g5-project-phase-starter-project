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

//start refresh_token_usecase tests

type RefreshTokenUsecaseTestSuite struct {
	suite.Suite
	RefreshTokenUsecaseTestSuite domain.RefreshTokenUsecase
	infra                        *mocks.Infrastructure
	mockUserRepo                 *mocks.UserRepository
	mockActiveUserRepo           *mocks.ActiveUserRepository
	contextTimeout               time.Duration
}

func (suite *RefreshTokenUsecaseTestSuite) SetupTest() {
	suite.mockUserRepo = new(mocks.UserRepository)
	suite.mockActiveUserRepo = new(mocks.ActiveUserRepository)
	suite.infra = new(mocks.Infrastructure)
	suite.contextTimeout = time.Second * 5
	suite.RefreshTokenUsecaseTestSuite = NewRefreshTokenUsecase(suite.mockUserRepo, suite.mockActiveUserRepo, suite.contextTimeout)
}

func (suite *RefreshTokenUsecaseTestSuite) TestLogout() {
	suite.Run("TestSuccess", func() {
		suite.mockActiveUserRepo.On("DeleteActiveUser", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		err := suite.RefreshTokenUsecaseTestSuite.RemoveActiveUser(context.Background(), "id", "user_agent")
		suite.NoError(err)
		suite.mockActiveUserRepo.AssertExpectations(suite.T())
		suite.mockActiveUserRepo.AssertCalled(suite.T(), "DeleteActiveUser", mock.Anything, mock.Anything, mock.Anything)
	})
	suite.Run("TestError", func() {
		suite.mockActiveUserRepo.On("DeleteActiveUser", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("some error")).Once()
		err := suite.RefreshTokenUsecaseTestSuite.RemoveActiveUser(context.Background(), "id", "user_agent")
		suite.Error(err)
		suite.mockActiveUserRepo.AssertExpectations(suite.T())
		suite.mockActiveUserRepo.AssertCalled(suite.T(), "DeleteActiveUser", mock.Anything, mock.Anything, mock.Anything)
	})
}

func (suite *RefreshTokenUsecaseTestSuite) TestCheckActiveUser() {
	suite.Run("TestSuccess", func() {
		suite.mockActiveUserRepo.On("FindActiveUser", mock.Anything, mock.Anything, mock.Anything).Return(domain.ActiveUser{}, nil).Once()
		_, err := suite.RefreshTokenUsecaseTestSuite.CheckActiveUser(context.Background(), "id", "user_agent")
		suite.NoError(err)
		suite.mockActiveUserRepo.AssertExpectations(suite.T())
		suite.mockActiveUserRepo.AssertCalled(suite.T(), "FindActiveUser", mock.Anything, mock.Anything, mock.Anything)
	})
	suite.Run("TestError", func() {
		suite.mockActiveUserRepo.On("FindActiveUser", mock.Anything, mock.Anything, mock.Anything).Return(domain.ActiveUser{}, errors.New("some error")).Once()
		_, err := suite.RefreshTokenUsecaseTestSuite.CheckActiveUser(context.Background(), "id", "user_agent")
		suite.Error(err)
		suite.mockActiveUserRepo.AssertExpectations(suite.T())
		suite.mockActiveUserRepo.AssertCalled(suite.T(), "FindActiveUser", mock.Anything, mock.Anything, mock.Anything)
	})
}

func (suite *RefreshTokenUsecaseTestSuite) TestCreateAccessToken() {
	suite.Run("TestSuccess", func() {
		suite.infra.On("CreateAccessToken", mock.Anything, mock.Anything, mock.Anything).Return("token", nil).Once()
		_, err := suite.RefreshTokenUsecaseTestSuite.CreateAccessToken(&domain.User{}, "secret", 1)
		suite.NoError(err)
		
	})

}

func (suite *RefreshTokenUsecaseTestSuite) TestCreateRefreshToken() {

	suite.Run("TestSuccess", func() {
		suite.infra.On("CreateRefreshToken", mock.Anything, mock.Anything, mock.Anything).Return("token", nil).Once()
		_, err := suite.RefreshTokenUsecaseTestSuite.CreateAccessToken(&domain.User{}, "secret", 1)
		suite.NoError(err)
		
	})
}

func (suite *RefreshTokenUsecaseTestSuite) TestExtractIDFromToken() {
	suite.Run("TestSuccess", func() {
		suite.infra.On("ExtractIDFromToken", mock.Anything, mock.Anything, mock.Anything).Return("token", nil).Once()
		_, err := suite.RefreshTokenUsecaseTestSuite.CreateAccessToken(&domain.User{}, "secret", 1)
		suite.NoError(err)
		
	})
	
}

func (suite *RefreshTokenUsecaseTestSuite) TestGetUserByID() {
	suite.Run("TestSuccess", func() {
		suite.mockUserRepo.On("FindUserByID", mock.Anything, mock.Anything).Return(domain.User{}, nil).Once()
		_, err := suite.RefreshTokenUsecaseTestSuite.GetUserByID(context.Background(), "id")
		suite.NoError(err)
		suite.mockUserRepo.AssertExpectations(suite.T())
		suite.mockUserRepo.AssertCalled(suite.T(), "FindUserByID", mock.Anything, mock.Anything)
	})
	suite.Run("TestError", func() {
		suite.mockUserRepo.On("FindUserByID", mock.Anything, mock.Anything).Return(domain.User{}, errors.New("some error")).Once()
		_, err := suite.RefreshTokenUsecaseTestSuite.GetUserByID(context.Background(), "id")
		suite.Error(err)
		suite.mockUserRepo.AssertExpectations(suite.T())
		suite.mockUserRepo.AssertCalled(suite.T(), "FindUserByID", mock.Anything, mock.Anything)
	})
}

func TestRefreshTokenUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(RefreshTokenUsecaseTestSuite))
}
