package usecases_test

import (
	usecases "blogapp/UseCases"
	"blogapp/mocks"
	"context"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RefreshUsecaseSuite struct {
	suite.Suite
	context        context.Context
	refreshUsecase *usecases.RefreshUseCase
	repo           *mocks.RefreshRepository
}

func (suite *RefreshUsecaseSuite) SetupTest() {
	suite.repo = new(mocks.RefreshRepository)
	suite.refreshUsecase = usecases.NewRefreshUseCase(suite.repo)
	suite.context = context.Background()
}

func (suite *RefreshUsecaseSuite) TestDeleteToken() {
	c, _ := gin.CreateTestContext(nil)
	id := primitive.NewObjectID()
	suite.repo.On("DeleteToken", mock.Anything, mock.Anything).Return(nil, 200)
	err, status := suite.refreshUsecase.DeleteToken(c, id)
	suite.Nil(err)
	suite.Equal(200, status)

}

func (suite *RefreshUsecaseSuite) TestFindToken() {
	c, _ := gin.CreateTestContext(nil)
	id := primitive.NewObjectID()
	suite.repo.On("FindToken", mock.Anything, mock.Anything).Return("", nil, 200)
	_, err, status := suite.refreshUsecase.FindToken(c, id)
	suite.Nil(err)
	suite.Equal(200, status)
}

func (suite *RefreshUsecaseSuite) TestStoreToken() {
	c, _ := gin.CreateTestContext(nil)
	id := primitive.NewObjectID()
	suite.repo.On("StoreToken", mock.Anything, mock.Anything, mock.Anything).Return(nil, 200)
	err, status := suite.refreshUsecase.StoreToken(c, id, "")
	suite.Nil(err)
	suite.Equal(200, status)
}

func TestRefreshUsecaseSuite(t *testing.T) {
	suite.Run(t, new(RefreshUsecaseSuite))
}
