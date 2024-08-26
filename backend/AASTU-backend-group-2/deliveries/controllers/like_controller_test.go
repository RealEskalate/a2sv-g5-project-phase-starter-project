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

type LikeControllerTestSuite struct {
	suite.Suite
	controller   *controllers.LikeController
	mockUsecase  *mocks.LikeUsecase
	mockContext  *gin.Context
	mockResponse *httptest.ResponseRecorder
}

func (suite *LikeControllerTestSuite) SetupTest() {
	suite.mockUsecase = new(mocks.LikeUsecase)
	suite.controller = controllers.NewLikeController(suite.mockUsecase)
	suite.mockResponse = httptest.NewRecorder()
	suite.mockContext, _ = gin.CreateTestContext(suite.mockResponse)
}

func (suite *LikeControllerTestSuite) TestCreateLike_Unauthorized() {
	suite.mockContext.Request = httptest.NewRequest(http.MethodPost, "/likes", nil)
	suite.controller.CreateLike(suite.mockContext)
	assert.Equal(suite.T(), http.StatusUnauthorized, suite.mockResponse.Code)
}

func (suite *LikeControllerTestSuite) TestCreateLike_Success() {
	// Setting up the context and parameters
	suite.mockContext.Params = gin.Params{gin.Param{Key: "postID", Value: "post123"}}
	suite.mockContext.Request = httptest.NewRequest(http.MethodPost, "/likes/post123", nil)
	suite.mockContext.Set("isadmin", true)
	suite.mockContext.Set("userid", "user456")

	// Mocking the use case
	suite.mockUsecase.On("CreateLike", mock.Anything, "user456", "post123").Return(nil)

	// Calling the method
	suite.controller.CreateLike(suite.mockContext)

	// Assertions
	assert.Equal(suite.T(), http.StatusOK, suite.mockResponse.Code)
	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *LikeControllerTestSuite) TestDeleteLike_Unauthorized() {
	suite.mockContext.Request = httptest.NewRequest(http.MethodDelete, "/likes", nil)
	suite.controller.DeleteLike(suite.mockContext)
	assert.Equal(suite.T(), http.StatusUnauthorized, suite.mockResponse.Code)
}

func (suite *LikeControllerTestSuite) TestDeleteLike_Success() {
	// Setting up the context and parameters
	suite.mockContext.Params = gin.Params{gin.Param{Key: "id", Value: "like123"}}
	suite.mockContext.Request = httptest.NewRequest(http.MethodDelete, "/likes/like123", nil)
	suite.mockContext.Set("isadmin", true)
	suite.mockContext.Set("userid", "like123")

	// Mocking the use case
	suite.mockUsecase.On("DeleteLike", mock.Anything, "like123").Return(nil)

	// Calling the method
	suite.controller.DeleteLike(suite.mockContext)

	// Assertions
	assert.Equal(suite.T(), http.StatusOK, suite.mockResponse.Code)
	assert.JSONEq(suite.T(), `{"message":"Post unliked successfully"}`, suite.mockResponse.Body.String())
	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *LikeControllerTestSuite) TestGetLikes_Success() {
	// Setting up the context and parameters
	suite.mockContext.Params = gin.Params{gin.Param{Key: "postID", Value: "post123"}}
	suite.mockContext.Request = httptest.NewRequest(http.MethodGet, "/likes/post123", nil)
	suite.mockContext.Set("isadmin", true)

	// Mocking the use case with a sample Like object
	mockLikes := []domain.Like{
		{
			ID:     primitive.NewObjectID(),
			UserID: primitive.NewObjectID(),
			BlogID: primitive.NewObjectID(),
		},
	}
	suite.mockUsecase.On("GetLikes", mock.Anything, "post123").Return(mockLikes, nil)

	// Calling the method
	suite.controller.GetLikes(suite.mockContext)

	// Constructing the expected JSON response
	// expectedResponse := `{
	// 	"likes": [
	// 		{
	// 			"_id": "` + mockLikes[0].ID.Hex() + `",
	// 			"user_id": "` + mockLikes[0].UserID.Hex() + `",
	// 			"post_id": "` + mockLikes[0].BlogID.Hex() + `"
	// 		}
	// 	]
	// }`

	// Assertions
	assert.Equal(suite.T(), http.StatusOK, suite.mockResponse.Code)
	// assert.JSONEq(suite.T(), expectedResponse, suite.mockResponse.Body.String())

	suite.mockUsecase.AssertExpectations(suite.T())
}

func TestLikeControllerTestSuite(t *testing.T) {
	suite.Run(t, new(LikeControllerTestSuite))
}
