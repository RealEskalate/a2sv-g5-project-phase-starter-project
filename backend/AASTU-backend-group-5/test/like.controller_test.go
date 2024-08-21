package test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RealEskalate/blogpost/delivery/controller"
	"github.com/RealEskalate/blogpost/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LikeControllerTestSuite struct {
	suite.Suite
	mockLikeUsecase *mocks.Like_Usecase_interface
	likeController  *controller.LikeController
}

func (suite *LikeControllerTestSuite) SetupTest() {
	suite.mockLikeUsecase = mocks.NewLike_Usecase_interface(suite.T())
	suite.likeController = &controller.LikeController{
		LikeUsecase: suite.mockLikeUsecase,
	}
}

func (suite *LikeControllerTestSuite) TestDeleteLike_ValidID() {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	validLikeID := primitive.NewObjectID().Hex()
	c.Params = []gin.Param{{Key: "like_id", Value: validLikeID}}

	suite.mockLikeUsecase.On("DeleteLike", validLikeID).Return(nil)

	suite.likeController.DeleteLike()(c)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Like deleted successfully!", response["message"])
}

func (suite *LikeControllerTestSuite) TestDeleteLike_InvalidID() {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	invalidLikeID := "invalid-id"
	c.Params = []gin.Param{{Key: "like_id", Value: invalidLikeID}}

	suite.likeController.DeleteLike()(c)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Contains(suite.T(), response["error"], "Invalid like ID format")
}

func (suite *LikeControllerTestSuite) TestDeleteLike_UsecaseError() {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	validLikeID := primitive.NewObjectID().Hex()
	c.Params = []gin.Param{{Key: "like_id", Value: validLikeID}}

	expectedError := errors.New("usecase error")
	suite.mockLikeUsecase.On("DeleteLike", validLikeID).Return(expectedError)

	suite.likeController.DeleteLike()(c)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Contains(suite.T(), response["error"], "Failed to delete like")
	assert.Contains(suite.T(), response["error"], expectedError.Error())
}

func TestLikeControllerTestSuite(t *testing.T) {
	suite.Run(t, new(LikeControllerTestSuite))
}
