package controllers_test

import (
	controllers "blogapp/Delivery/controllers"
	"blogapp/mocks"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// generate comment controller test suite

type CommentCtrlTestSuit struct {
	suite.Suite
	ctrl        *controllers.CommentController
	mockUsecase *mocks.CommentUseCase
}

// setup test suite
func (suite *CommentCtrlTestSuit) SetupTest() {
	suite.mockUsecase = new(mocks.CommentUseCase)
	suite.ctrl = controllers.NewCommentController(suite.mockUsecase)
}

// tear down
func (suite *CommentCtrlTestSuit) TearDownTest() {
	suite.mockUsecase.AssertExpectations(suite.T())
}

// test comment on post
func (suite *CommentCtrlTestSuit) TestCommentOnPost() {
	// test case success
	suite.Run("CommentOnPost success", func() {
		claim := mocks.GetSampleClaim()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("claim", claim)

		postID := primitive.NewObjectID().Hex()
		c.Params = gin.Params{
			gin.Param{Key: "id", Value: postID},
		}

		comment := mocks.GetSampleComment()
		body, err := json.Marshal(comment)
		suite.Nil(err)

		c.Request = httptest.NewRequest(http.MethodPost, "/comments/create", bytes.NewBuffer(body))

		suite.mockUsecase.On("CommentOnPost", c, mock.Anything, postID).Return(nil, http.StatusCreated).Once()

		suite.ctrl.CommentOnPost(c)
		suite.Equal(http.StatusCreated, w.Code)

	})
}

// run test
func TestCommentController(t *testing.T) {
	suite.Run(t, new(CommentCtrlTestSuit))
}
