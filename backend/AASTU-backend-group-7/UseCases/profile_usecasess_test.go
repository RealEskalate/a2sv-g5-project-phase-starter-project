package usecases_test

import (
	"blogapp/Domain"
	usecases "blogapp/UseCases"
	"blogapp/mocks"
	"context"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileUsecaseSuite struct {
	suite.Suite
	context        context.Context
	profileUsecase *usecases.ProfileUseCases
	repo           *mocks.ProfileRepository
}

func (suite *ProfileUsecaseSuite) SetupTest() {
	suite.repo = new(mocks.ProfileRepository)
	suite.profileUsecase = usecases.NewProfileUseCase(suite.repo)
	suite.context = context.Background()
}

func (suite *ProfileUsecaseSuite) TestGetProfile() {
	c, _ := gin.CreateTestContext(nil)
	id := primitive.NewObjectID()
	suite.repo.On("GetProfile", mock.Anything, mock.Anything, mock.Anything).Return(Domain.OmitedUser{}, nil, 200)
	_, err, status := suite.profileUsecase.GetProfile(c, id, Domain.AccessClaims{})
	suite.Nil(err)
	suite.Equal(200, status)
}

func (suite *ProfileUsecaseSuite) TestUpdateProfile() {
	c, _ := gin.CreateTestContext(nil)
	id := primitive.NewObjectID()
	user := Domain.User{}
	suite.repo.On("UpdateProfile", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(Domain.OmitedUser{}, nil, 200)
	_, err, status := suite.profileUsecase.UpdateProfile(c, id, user, Domain.AccessClaims{})
	suite.Nil(err)
	suite.Equal(200, status)
}

func (suite *ProfileUsecaseSuite) TestDeleteProfile() {
	c, _ := gin.CreateTestContext(nil)
	id := primitive.NewObjectID()
	suite.repo.On("DeleteProfile", mock.Anything, mock.Anything, mock.Anything).Return(nil, 200)
	err, status := suite.profileUsecase.DeleteProfile(c, id, Domain.AccessClaims{})
	suite.Nil(err)
	suite.Equal(200, status)
}

func TestProfileUsecaseSuite(t *testing.T) {
	suite.Run(t, new(ProfileUsecaseSuite))
}
