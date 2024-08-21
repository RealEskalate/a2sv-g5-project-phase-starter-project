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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LogoutUsecaseTestSuite struct {
	suite.Suite
	LogoutUsecaseTestSuite domain.LogoutUsecase
	mockActiveUserRepo     *mocks.ActiveUserRepository
	contextTimeout         time.Duration
}

func (suite *LogoutUsecaseTestSuite) SetupTest() {
	suite.mockActiveUserRepo = new(mocks.ActiveUserRepository)
	suite.contextTimeout = time.Second * 5
	suite.LogoutUsecaseTestSuite = NewLogoutUsecase(suite.mockActiveUserRepo, suite.contextTimeout)
}
func (suite *LogoutUsecaseTestSuite) TestLogout() {
	// Setup

	suite.Run("sucess", func() {

		id := "testID"
		userAgent := "testUserAgent"
		ctx := context.Background()

		// Mock
		suite.mockActiveUserRepo.On("DeleteActiveUser", id, userAgent, mock.Anything).Return(nil).Once()

		// Execute the method
		err := suite.LogoutUsecaseTestSuite.Logout(ctx, id, userAgent)

		// Assertions
		suite.NoError(err)
		suite.mockActiveUserRepo.AssertCalled(suite.T(), "DeleteActiveUser", id, userAgent, mock.Anything)
	})
	suite.Run("Error_test", func() {

		id := "testID"
		userAgent := "testUserAgent"
		ctx := context.Background()

		// Mock to return an error
		suite.mockActiveUserRepo.On("DeleteActiveUser", id, userAgent, mock.Anything).Return(errors.New("some error")).Once()

		// Execute the method
		err := suite.LogoutUsecaseTestSuite.Logout(ctx, id, userAgent)

		// Assertions
		suite.Error(err)
		suite.EqualError(err, "some error")
		suite.mockActiveUserRepo.AssertCalled(suite.T(), "DeleteActiveUser", id, userAgent, mock.Anything)
	})

}

func (suite *LogoutUsecaseTestSuite) TestCheckActiveUser() {
	// Setup
	suite.Run("success", func() {

		id := "testID"
		userAgent := "testUserAgent"
		expectedActiveUser := domain.ActiveUser{
			ID:        primitive.NewObjectID(),
			UserAgent: userAgent,
		}

		// Mock
		suite.mockActiveUserRepo.On("FindActiveUser", id, userAgent, mock.Anything).Return(expectedActiveUser, nil).Once()

		// Execute the method
		activeUser, err := suite.LogoutUsecaseTestSuite.CheckActiveUser(context.Background(), id, userAgent)

		// Assertions
		suite.NoError(err)
		suite.Equal(expectedActiveUser, activeUser)
	})
	suite.Run("Error", func() {

		id := "testID"
		userAgent := "testUserAgent"
		expectedActiveUser := domain.ActiveUser{}

		// Mock
		suite.mockActiveUserRepo.On("FindActiveUser", id, userAgent, mock.Anything).Return(domain.ActiveUser{}, errors.New("user not found")).Once()

		// Execute the method
		activeUser, err := suite.LogoutUsecaseTestSuite.CheckActiveUser(context.Background(), id, userAgent)

		// Assertions
		suite.Error(err)
		suite.Equal(expectedActiveUser, activeUser)
	})
}
func TestLogoutUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(LogoutUsecaseTestSuite))
}
