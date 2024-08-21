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

type DislikeControllerTestSuite struct {
	suite.Suite
	mockDislikeUsecase *mocks.DisLike_Usecase_interface
	dislikeController  *controller.DislikeController
}

func (suite *DislikeControllerTestSuite) SetupTest() {
	suite.mockDislikeUsecase = mocks.NewDisLike_Usecase_interface(suite.T())
	suite.dislikeController = &controller.DislikeController{
		DislikeUsecase: suite.mockDislikeUsecase,
	}
}

func (suite *DislikeControllerTestSuite) TestDeleteDisLike_ValidID() {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	validDislikeID := primitive.NewObjectID().Hex()
	c.Params = []gin.Param{{Key: "dislike_id", Value: validDislikeID}}

	suite.mockDislikeUsecase.On("DeleteDisLike", validDislikeID).Return(nil)

	suite.dislikeController.DeleteDisLike()(c)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Dislike deleted successfully!", response["message"])
}

func (suite *DislikeControllerTestSuite) TestDeleteDisLike_InvalidID() {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	invalidDislikeID := "invalid-id"
	c.Params = []gin.Param{{Key: "dislike_id", Value: invalidDislikeID}}

	suite.dislikeController.DeleteDisLike()(c)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Contains(suite.T(), response["error"], "Invalid dislike ID format")
}

func (suite *DislikeControllerTestSuite) TestDeleteDisLike_UsecaseError() {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	validDislikeID := primitive.NewObjectID().Hex()
	c.Params = []gin.Param{{Key: "dislike_id", Value: validDislikeID}}

	expectedError := errors.New("usecase error")
	suite.mockDislikeUsecase.On("DeleteDisLike", validDislikeID).Return(expectedError)

	suite.dislikeController.DeleteDisLike()(c)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Contains(suite.T(), response["error"], "Failed to delete dislike")
	assert.Contains(suite.T(), response["error"], expectedError.Error())
}

func TestDislikeControllerTestSuite(t *testing.T) {
	suite.Run(t, new(DislikeControllerTestSuite))
}
