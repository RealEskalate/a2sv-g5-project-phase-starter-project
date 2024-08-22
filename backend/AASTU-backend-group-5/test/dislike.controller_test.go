package test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RealEskalate/blogpost/delivery/controller"
	"github.com/RealEskalate/blogpost/domain"
	"github.com/RealEskalate/blogpost/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DislikeControllerSuite struct {
	suite.Suite
	controller *controller.DislikeController
	usecase    *mocks.DisLike_Usecase_interface
}

func (suite *DislikeControllerSuite) SetupTest() {
	usecase := new(mocks.DisLike_Usecase_interface)
	suite.controller = controller.NewDislikeController(usecase)
	suite.usecase = usecase
}

// Positive Tests

func (suite *DislikeControllerSuite) TestGetDislikes() {
	postID := primitive.NewObjectID().Hex()
	expectedDislikes := []domain.DisLike{
		{UserID: primitive.NewObjectID(), PostID: primitive.NewObjectID()},
	}

	suite.usecase.On("GetDislikes", postID).Return(expectedDislikes, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "post_id", Value: postID}}

	suite.controller.GetDislikes(c)

	suite.Equal(http.StatusOK, w.Code)
	// Additional assertions for response body can be added here
}
func (suite *DislikeControllerSuite) TestCreateDislike() {
	userID := primitive.NewObjectID().Hex()
	postID := primitive.NewObjectID().Hex()

	suite.usecase.On("CreateDislike", userID, postID).Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: userID}, gin.Param{Key: "post_id", Value: postID}}

	suite.controller.CreateDislike(c)

	suite.Equal(http.StatusOK, w.Code)
}

func (suite *DislikeControllerSuite) TestToggleDislike() {
	userID := primitive.NewObjectID().Hex()
	postID := primitive.NewObjectID().Hex()

	suite.usecase.On("ToggleDislike", userID, postID).Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: userID}, gin.Param{Key: "post_id", Value: postID}}

	suite.controller.ToggleDislike(c)

	suite.Equal(http.StatusOK, w.Code)
}

func (suite *DislikeControllerSuite) TestRemoveDislike() {
	userID := primitive.NewObjectID().Hex()
	postID := primitive.NewObjectID().Hex()

	suite.usecase.On("RemoveDislike", userID, postID).Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: userID}, gin.Param{Key: "post_id", Value: postID}}

	suite.controller.RemoveDislike(c)

	suite.Equal(http.StatusOK, w.Code)
}

// Negative Tests

func (suite *DislikeControllerSuite) TestGetDislikes_Error() {
	postID := primitive.NewObjectID().Hex()

	suite.usecase.On("GetDislikes", postID).Return(nil, errors.New("error retrieving dislikes"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "post_id", Value: postID}}

	suite.controller.GetDislikes(c)

	suite.Equal(http.StatusInternalServerError, w.Code)
	suite.JSONEq(`{"error": "error retrieving dislikes"}`, w.Body.String())
}

func (suite *DislikeControllerSuite) TestCreateDislike_Error() {
	userID := primitive.NewObjectID().Hex()
	postID := primitive.NewObjectID().Hex()

	suite.usecase.On("CreateDislike", userID, postID).Return(errors.New("error creating dislike"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: userID}, gin.Param{Key: "post_id", Value: postID}}

	suite.controller.CreateDislike(c)

	suite.Equal(http.StatusInternalServerError, w.Code)
	suite.JSONEq(`{"error": "error creating dislike"}`, w.Body.String())
}

func (suite *DislikeControllerSuite) TestToggleDislike_Error() {
	userID := primitive.NewObjectID().Hex()
	postID := primitive.NewObjectID().Hex()

	suite.usecase.On("ToggleDislike", userID, postID).Return(errors.New("error toggling dislike"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: userID}, gin.Param{Key: "post_id", Value: postID}}

	suite.controller.ToggleDislike(c)

	suite.Equal(http.StatusInternalServerError, w.Code)
	suite.JSONEq(`{"error": "error toggling dislike"}`, w.Body.String())
}

func (suite *DislikeControllerSuite) TestRemoveDislike_Error() {
	userID := primitive.NewObjectID().Hex()
	postID := primitive.NewObjectID().Hex()

	suite.usecase.On("RemoveDislike", userID, postID).Return(errors.New("error removing dislike"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: userID}, gin.Param{Key: "post_id", Value: postID}}

	suite.controller.RemoveDislike(c)

	suite.Equal(http.StatusInternalServerError, w.Code)
	suite.JSONEq(`{"error": "error removing dislike"}`, w.Body.String())
}

func TestDislikeController(t *testing.T) {
	suite.Run(t, new(DislikeControllerSuite))
}
