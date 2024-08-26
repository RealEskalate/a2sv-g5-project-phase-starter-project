package controllers_test

import (
	"blog_g2/deliveries/controllers"
	"blog_g2/domain"
	"blog_g2/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DisLikeControllerTestSuite struct {
	suite.Suite
	controller   *controllers.DisLikeController
	mockUsecase  *mocks.DisLikeUsecase
	mockContext  *gin.Context
	mockResponse *httptest.ResponseRecorder
}

func (suite *DisLikeControllerTestSuite) SetupTest() {
	suite.mockUsecase = new(mocks.DisLikeUsecase)
	suite.controller = controllers.NewDisLikeController(suite.mockUsecase)
	suite.mockResponse = httptest.NewRecorder()
	suite.mockContext, _ = gin.CreateTestContext(suite.mockResponse)
}

func (suite *DisLikeControllerTestSuite) TestCreateDisLike_Unauthorized() {
	suite.mockContext.Request = httptest.NewRequest(http.MethodPost, "/dislikes", nil)
	suite.controller.CreateDisLike(suite.mockContext)
	assert.Equal(suite.T(), http.StatusUnauthorized, suite.mockResponse.Code)
}

func (suite *DisLikeControllerTestSuite) TestCreateDisLike_Success() {
	// Setting up the context and parameters
	suite.mockContext.Params = gin.Params{gin.Param{Key: "postID", Value: "post123"}}
	suite.mockContext.Request = httptest.NewRequest(http.MethodPost, "/dislikes/post123", nil)
	suite.mockContext.Set("isadmin", true)
	suite.mockContext.Set("userid", "user456")

	// Mocking the use case
	suite.mockUsecase.On("CreateDisLike", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// Calling the method
	suite.controller.CreateDisLike(suite.mockContext)

	// Assertions
	assert.Equal(suite.T(), http.StatusOK, suite.mockResponse.Code)
	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *DisLikeControllerTestSuite) TestDeleteDisLike_Unauthorized() {
	suite.mockContext.Request = httptest.NewRequest(http.MethodDelete, "/dislikes", nil)
	suite.controller.DeleteDisLike(suite.mockContext)
	assert.Equal(suite.T(), http.StatusUnauthorized, suite.mockResponse.Code)
}

func (suite *DisLikeControllerTestSuite) TestDeleteDisLike_Success() {
	// Setting up the context and parameters
	suite.mockContext.Params = gin.Params{gin.Param{Key: "id", Value: "dislike123"}}
	suite.mockContext.Request = httptest.NewRequest(http.MethodDelete, "/dislikes/dislike123", nil)
	suite.mockContext.Set("isadmin", true)
	suite.mockContext.Set("userid", "dislike123")

	// Mocking the use case
	suite.mockUsecase.On("DeleteDisLike", mock.Anything, mock.Anything).Return(nil)

	// Calling the method
	suite.controller.DeleteDisLike(suite.mockContext)

	// Assertions
	assert.Equal(suite.T(), http.StatusOK, suite.mockResponse.Code)
	assert.JSONEq(suite.T(), `{"message":"Post undisliked successfully"}`, suite.mockResponse.Body.String())
	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *DisLikeControllerTestSuite) TestGetDisLikes_Success() {
	// Setting up the context and parameters
	suite.mockContext.Params = gin.Params{gin.Param{Key: "postID", Value: "post123"}}
	suite.mockContext.Request = httptest.NewRequest(http.MethodGet, "/dislikes/post123", nil)
	suite.mockContext.Set("isadmin", true)

	// Mocking the use case
	mockDislikes := []domain.DisLike{
		{ID: primitive.NewObjectID(), UserID: primitive.NewObjectID(), BlogID: primitive.NewObjectID()},
	}
	suite.mockUsecase.On("GetDisLikes", mock.Anything, mock.Anything).Return(mockDislikes, nil)

	// Calling the method
	suite.controller.GetDisLikes(suite.mockContext)

	// Assertions
	assert.Equal(suite.T(), http.StatusOK, suite.mockResponse.Code)
	suite.mockUsecase.AssertExpectations(suite.T())
}

func TestDisLikeControllerTestSuite(t *testing.T) {
	suite.Run(t, new(DisLikeControllerTestSuite))
}
