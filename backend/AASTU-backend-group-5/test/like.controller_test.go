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

type LikeControllerSuite struct {
	suite.Suite
	controller *controller.LikeController
	usecase    *mocks.Like_Usecase_interface
}

func (suite *LikeControllerSuite) SetupTest() {
	usecase := new(mocks.Like_Usecase_interface)
	suite.controller = controller.NewLikeController(usecase)
	suite.usecase = usecase
}

func (suite *LikeControllerSuite) TestGetLikes() {
	postID := primitive.NewObjectID().Hex()
	postObjectID, _ := primitive.ObjectIDFromHex(postID)
	expectedLikes := []domain.Like{
		{UserID: primitive.NewObjectID(), PostID: postObjectID},
	}

	suite.usecase.On("GetLikes", postID).Return(expectedLikes, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "post_id", Value: postID}}

	suite.controller.GetLikes(c)

	suite.Equal(http.StatusOK, w.Code)
	// Additional assertions for response body can be added here
}
func (suite *LikeControllerSuite) TestCreateLike() {
	userID := primitive.NewObjectID().Hex()
	postID := primitive.NewObjectID().Hex()

	suite.usecase.On("CreateLike", userID, postID).Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: userID}, gin.Param{Key: "post_id", Value: postID}}

	suite.controller.CreateLike(c)

	suite.Equal(http.StatusOK, w.Code)
}

func (suite *LikeControllerSuite) TestToggleLike() {
	userID := primitive.NewObjectID().Hex()
	postID := primitive.NewObjectID().Hex()

	suite.usecase.On("ToggleLike", userID, postID).Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: userID}, gin.Param{Key: "post_id", Value: postID}}

	suite.controller.ToggleLike(c)

	suite.Equal(http.StatusOK, w.Code)
}

func (suite *LikeControllerSuite) TestRemoveLike() {
	userID := primitive.NewObjectID().Hex()
	postID := primitive.NewObjectID().Hex()

	suite.usecase.On("RemoveLike", userID, postID).Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: userID}, gin.Param{Key: "post_id", Value: postID}}

	suite.controller.RemoveLike(c)

	suite.Equal(http.StatusOK, w.Code)
}

// Negative Tests

func (suite *LikeControllerSuite) TestGetLikes_Error() {
	postID := primitive.NewObjectID().Hex()

	suite.usecase.On("GetLikes", postID).Return(nil, errors.New("error retrieving likes"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "post_id", Value: postID}}

	suite.controller.GetLikes(c)

	suite.Equal(http.StatusInternalServerError, w.Code)
	suite.JSONEq(`{"error": "error retrieving likes"}`, w.Body.String())
}

func (suite *LikeControllerSuite) TestCreateLike_Error() {
	userID := primitive.NewObjectID().Hex()
	postID := primitive.NewObjectID().Hex()

	suite.usecase.On("CreateLike", userID, postID).Return(errors.New("error creating like"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: userID}, gin.Param{Key: "post_id", Value: postID}}

	suite.controller.CreateLike(c)

	suite.Equal(http.StatusInternalServerError, w.Code)
	suite.JSONEq(`{"error": "error creating like"}`, w.Body.String())
}

func (suite *LikeControllerSuite) TestToggleLike_Error() {
	userID := primitive.NewObjectID().Hex()
	postID := primitive.NewObjectID().Hex()

	suite.usecase.On("ToggleLike", userID, postID).Return(errors.New("error toggling like"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: userID}, gin.Param{Key: "post_id", Value: postID}}

	suite.controller.ToggleLike(c)

	suite.Equal(http.StatusInternalServerError, w.Code)
	suite.JSONEq(`{"error": "error toggling like"}`, w.Body.String())
}

func (suite *LikeControllerSuite) TestRemoveLike_Error() {
	userID := primitive.NewObjectID().Hex()
	postID := primitive.NewObjectID().Hex()

	suite.usecase.On("RemoveLike", userID, postID).Return(errors.New("error removing like"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: userID}, gin.Param{Key: "post_id", Value: postID}}

	suite.controller.RemoveLike(c)

	suite.Equal(http.StatusInternalServerError, w.Code)
	suite.JSONEq(`{"error": "error removing like"}`, w.Body.String())
}

func TestLikeControllerSuite(t *testing.T) {
	suite.Run(t, new(LikeControllerSuite))
}
