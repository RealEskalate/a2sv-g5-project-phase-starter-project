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

type UserUsecaseSuite struct {
	suite.Suite
	context     context.Context
	userUsecase *usecases.UserUseCases
	repo        *mocks.UserRepository
}

func (suite *UserUsecaseSuite) SetupTest() {
	suite.repo = new(mocks.UserRepository)
	suite.userUsecase = usecases.NewUserUseCase(suite.repo)
	suite.context = context.Background()
}

func (suite *UserUsecaseSuite) TestGetUsers() {
	c, _ := gin.CreateTestContext(nil)
	suite.repo.On("GetUsers", mock.Anything).Return([]*Domain.OmitedUser{}, nil, 200)
	_, err, status := suite.userUsecase.GetUsers(c)
	suite.Nil(err)
	suite.Equal(200, status)
}

func (suite *UserUsecaseSuite) TestGetUsersById() {
	c, _ := gin.CreateTestContext(nil)
	id := primitive.NewObjectID()
	current_user := Domain.AccessClaims{}
	suite.repo.On("GetUsersById", mock.Anything, mock.Anything, mock.Anything).Return(Domain.OmitedUser{}, nil, 200)
	_, err, status := suite.userUsecase.GetUsersById(c, id, current_user)
	suite.Nil(err)
	suite.Equal(200, status)
}

func (suite *UserUsecaseSuite) TestCreateUser() {
	c, _ := gin.CreateTestContext(nil)
	user := Domain.User{}
	suite.repo.On("CreateUser", mock.Anything, mock.Anything).Return(Domain.OmitedUser{}, nil, 200)
	_, err, status := suite.userUsecase.CreateUser(c, &user)
	suite.Nil(err)
	suite.Equal(200, status)
}

func (suite *UserUsecaseSuite) TestUpdateUsersById() {
	c, _ := gin.CreateTestContext(nil)
	id := primitive.NewObjectID()
	user := Domain.User{}
	current_user := Domain.AccessClaims{}
	suite.repo.On("UpdateUsersById", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(Domain.OmitedUser{}, nil, 200)
	_, err, status := suite.userUsecase.UpdateUsersById(c, id, user, current_user)
	suite.Nil(err)
	suite.Equal(200, status)
}

func (suite *UserUsecaseSuite) TestDeleteUsersById() {
	c, _ := gin.CreateTestContext(nil)
	id := primitive.NewObjectID()
	current_user := Domain.AccessClaims{}
	suite.repo.On("DeleteUsersById", mock.Anything, mock.Anything, mock.Anything).Return(nil, 200)
	err, status := suite.userUsecase.DeleteUsersById(c, id, current_user)
	suite.Nil(err)
	suite.Equal(200, status)
}

func (suite *UserUsecaseSuite) TestPromoteUser() {
	c, _ := gin.CreateTestContext(nil)
	id := primitive.NewObjectID()
	current_user := Domain.AccessClaims{}
	suite.repo.On("PromoteUser", mock.Anything, mock.Anything, mock.Anything).Return(Domain.OmitedUser{}, nil, 200)
	_, err, status := suite.userUsecase.PromoteUser(c, id, current_user)
	suite.Nil(err)
	suite.Equal(200, status)
}

func (suite *UserUsecaseSuite) TestDemoteUser() {
	c, _ := gin.CreateTestContext(nil)
	id := primitive.NewObjectID()
	current_user := Domain.AccessClaims{}
	suite.repo.On("DemoteUser", mock.Anything, mock.Anything, mock.Anything).Return(Domain.OmitedUser{}, nil, 200)
	_, err, status := suite.userUsecase.DemoteUser(c, id, current_user)
	suite.Nil(err)
	suite.Equal(200, status)
}

func TestUserUsecaseSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseSuite))
}
