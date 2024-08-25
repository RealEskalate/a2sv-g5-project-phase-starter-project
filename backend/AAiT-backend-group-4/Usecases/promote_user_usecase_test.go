package usecases_test

import (
	domain "aait-backend-group4/Domain"
	usecases "aait-backend-group4/Usecases"
	"aait-backend-group4/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PromoteUsecaseTestSuite struct {
	suite.Suite
	mockUserRepo   *mocks.UserRepository
	promoteUsecase domain.UserUsecase
}

func (suite *PromoteUsecaseTestSuite) SetupTest() {
	suite.mockUserRepo = new(mocks.UserRepository)
	suite.promoteUsecase = usecases.NewUserProfileUsecase(suite.mockUserRepo, 2*time.Second)
}

func (suite *PromoteUsecaseTestSuite) TestPromote_Success() {
	ctx := context.Background()
	userID := primitive.NewObjectID().Hex()

	promotedUser := domain.User{
		ID:        primitive.NewObjectID(),
		Username:  "testuser",
		Email:     "test@example.com",
		User_Role: "admin", // Assuming promotion changes role to admin
	}

	suite.mockUserRepo.On("Promote", mock.Anything, userID).Return(promotedUser, nil)

	resultUser, err := suite.promoteUsecase.Promote(ctx, userID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), promotedUser, resultUser)
	assert.Equal(suite.T(), "admin", resultUser.User_Role)

	suite.mockUserRepo.AssertExpectations(suite.T())
}

func (suite *PromoteUsecaseTestSuite) TestPromote_Error() {
	ctx := context.Background()
	userID := primitive.NewObjectID().Hex()

	suite.mockUserRepo.On("Promote", mock.Anything, userID).Return(domain.User{}, errors.New("error"))

	resultUser, err := suite.promoteUsecase.Promote(ctx, userID)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), domain.User{}, resultUser)

	suite.mockUserRepo.AssertExpectations(suite.T())
}

func TestPromoteUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(PromoteUsecaseTestSuite))
}
