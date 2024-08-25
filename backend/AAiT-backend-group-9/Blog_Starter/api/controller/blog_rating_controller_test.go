package controller

import (
    "Blog_Starter/domain"
    "Blog_Starter/domain/mocks"
    "Blog_Starter/utils"
    "context"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    "time"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/suite"
    "bou.ke/monkey"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// Test Suite
type BlogRatingControllerTestSuite struct {
    suite.Suite
    controller *BlogRatingController
    useCase    *mocks.BlogRatingUseCase
}

func (suite *BlogRatingControllerTestSuite) SetupTest() {
    suite.useCase = new(mocks.BlogRatingUseCase)
    suite.controller = NewBlogRatingController(suite.useCase, time.Second*2)
}

func (suite *BlogRatingControllerTestSuite) TestInsertRating() {
    // Mock CheckUser function
    monkey.Patch(utils.CheckUser, func(c *gin.Context) (*domain.AuthenticatedUser, error) {
        return &domain.AuthenticatedUser{UserID: "test_user"}, nil
    })
    defer monkey.Unpatch(utils.CheckUser)

    // Mock use case
    suite.useCase.On("InsertRating", mock.Anything, mock.Anything).Return(&domain.BlogRating{UserID: "test_user", BlogID: "test_blog", Rating: 5}, nil)

    // Create a request
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Request = httptest.NewRequest("POST", "/rating", strings.NewReader(`{"rating": 5}`))
    c.Params = gin.Params{gin.Param{Key: "blog_id", Value: "test_blog"}}

    // Call the controller method
    suite.controller.InsertRating(c)

    // Assertions
    suite.Equal(http.StatusOK, w.Code)
    suite.Contains(w.Body.String(), "test_user")
    suite.Contains(w.Body.String(), "test_blog")
    suite.Contains(w.Body.String(), "5")
}


func (suite *BlogRatingControllerTestSuite) TestInsertRating_InternalServerError() {
    // Mock CheckUser function
    monkey.Patch(utils.CheckUser, func(c *gin.Context) (*domain.AuthenticatedUser, error) {
        return &domain.AuthenticatedUser{UserID: "test_user"}, nil
    })
    defer monkey.Unpatch(utils.CheckUser)

    // Mock use case to return an error
    suite.useCase.On("InsertRating", mock.Anything, mock.Anything).Return(nil, context.DeadlineExceeded)

    // Create a request
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Request = httptest.NewRequest("POST", "/rating", strings.NewReader(`{"rating": 5}`))
    c.Params = gin.Params{gin.Param{Key: "blog_id", Value: "test_blog"}}

    // Call the controller method
    suite.controller.InsertRating(c)

    // Assertions
    suite.Equal(http.StatusInternalServerError, w.Code)
    suite.Contains(w.Body.String(), "internal server error")
}

func (suite *BlogRatingControllerTestSuite) TestUpdateRating() {
    // Mock use case
    suite.useCase.On("GetRatingByID", mock.Anything, "test_id").Return(&domain.BlogRating{RatingID: primitive.NewObjectID(), UserID: "test_user", BlogID: "test_blog", Rating: 5}, nil)
    suite.useCase.On("UpdateRating", mock.Anything, 4, "test_id").Return(&domain.BlogRating{RatingID: primitive.NewObjectID(), UserID: "test_user", BlogID: "test_blog", Rating: 4}, nil)

    // Create a request
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Request = httptest.NewRequest("PUT", "/rating/test_id", strings.NewReader(`{"rating": 4}`))
    c.Params = gin.Params{gin.Param{Key: "id", Value: "test_id"}}

    // Call the controller method
    suite.controller.UpdateRating(c)

    // Assertions
    suite.Equal(http.StatusOK, w.Code)
    suite.Contains(w.Body.String(), "test_user")
    suite.Contains(w.Body.String(), "test_blog")
    suite.Contains(w.Body.String(), "4")
}
func (suite *BlogRatingControllerTestSuite) TestUpdateRating_InternalServerError() {
    // Mock use case to return an error
    suite.useCase.On("GetRatingByID", mock.Anything, "test_id").Return(nil, context.DeadlineExceeded)

    // Create a request
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Request = httptest.NewRequest("PUT", "/rating/test_id", strings.NewReader(`{"rating": 4}`))
    c.Params = gin.Params{gin.Param{Key: "id", Value: "test_id"}}

    // Call the controller method
    suite.controller.UpdateRating(c)

    // Assertions
    suite.Equal(http.StatusInternalServerError, w.Code)
    suite.Contains(w.Body.String(), "internal server error")
}

func (suite *BlogRatingControllerTestSuite) TestDeleteRating() {
    // Mock use case
    suite.useCase.On("DeleteRating", mock.Anything, "test_id").Return(&domain.BlogRating{RatingID: primitive.NewObjectID(), UserID: "test_user", BlogID: "test_blog", Rating: 5}, nil)

    // Create a request
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Request = httptest.NewRequest("DELETE", "/rating/test_id", nil)
    c.Params = gin.Params{gin.Param{Key: "id", Value: "test_id"}}

    // Call the controller method
    suite.controller.DeleteRating(c)

    // Assertions
    suite.Equal(http.StatusOK, w.Code)
    suite.Contains(w.Body.String(), "test_user")
    suite.Contains(w.Body.String(), "test_blog")
    suite.Contains(w.Body.String(), "5")
}

func (suite *BlogRatingControllerTestSuite) TestDeleteRating_InternalServerError() {
    // Mock use case to return an error
    suite.useCase.On("DeleteRating", mock.Anything, "test_id").Return(nil, context.DeadlineExceeded)

    // Create a request
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Request = httptest.NewRequest("DELETE", "/rating/test_id", nil)
    c.Params = gin.Params{gin.Param{Key: "id", Value: "test_id"}}

    // Call the controller method
    suite.controller.DeleteRating(c)

    // Assertions
    suite.Equal(http.StatusInternalServerError, w.Code)
    suite.Contains(w.Body.String(), "internal server error")
}

func TestBlogRatingControllerTestSuite(t *testing.T) {
    suite.Run(t, new(BlogRatingControllerTestSuite))
}