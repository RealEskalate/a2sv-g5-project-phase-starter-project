package controller

import (
    "Blog_Starter/domain"
    "Blog_Starter/domain/mocks"
    "Blog_Starter/utils"
	"encoding/json"
    "bytes"
    "errors"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
    "bou.ke/monkey"
)

type BlogCommentControllerTestSuite struct {
    suite.Suite
    controller *BlogCommentController
    useCase    *mocks.CommentUseCase
}

func (suite *BlogCommentControllerTestSuite) SetupTest() {
    suite.useCase = new(mocks.CommentUseCase)
    suite.controller = NewBlogCommentController(suite.useCase, time.Second*2)
}

func (suite *BlogCommentControllerTestSuite) TestCreateComment_Success() {
	gin.SetMode(gin.TestMode)
    router := gin.Default()
    router.POST("/comment/:blog_id", suite.controller.CreateComment)

    commentRequest := &domain.CommentRequest{Content: "This is a comment"}
    expectedCommentRequest := &domain.CommentRequest{UserID: "user1", BlogID: "blog1", Content: "This is a comment"}
    suite.useCase.On("Create", mock.Anything, expectedCommentRequest).Return(&domain.Comment{Content: "This is a comment"}, nil)

    w := httptest.NewRecorder()
    reqBody, _ := json.Marshal(commentRequest)
    req, _ := http.NewRequest(http.MethodPost, "/comment/blog1", bytes.NewBuffer(reqBody))
    req.Header.Set("Content-Type", "application/json")

    monkey.Patch(utils.CheckUser, func(c *gin.Context) (*domain.AuthenticatedUser, error) {
        return &domain.AuthenticatedUser{UserID: "user1"}, nil
    })
    defer monkey.Unpatch(utils.CheckUser)

    router.ServeHTTP(w, req)

    suite.Equal(http.StatusOK, w.Code)
    suite.useCase.AssertExpectations(suite.T())
}

func (suite *BlogCommentControllerTestSuite) TestCreateComment_InvalidRequestFormat() {
    gin.SetMode(gin.TestMode)
    router := gin.Default()
    router.POST("/comment/:blog_id", suite.controller.CreateComment)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest(http.MethodPost, "/comment/blog1", bytes.NewBuffer([]byte("invalid json")))
    req.Header.Set("Content-Type", "application/json")

    router.ServeHTTP(w, req)

    suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *BlogCommentControllerTestSuite) TestCreateComment_Unauthorized() {
    gin.SetMode(gin.TestMode)
    router := gin.Default()
    router.POST("/comment/:blog_id", suite.controller.CreateComment)

    commentRequest := &domain.CommentRequest{Content: "This is a comment"}

    w := httptest.NewRecorder()
    reqBody, _ := json.Marshal(commentRequest)
    req, _ := http.NewRequest(http.MethodPost, "/comment/blog1", bytes.NewBuffer(reqBody))
    req.Header.Set("Content-Type", "application/json")

    monkey.Patch(utils.CheckUser, func(c *gin.Context) (*domain.AuthenticatedUser, error) {
        return nil, errors.New("unauthorized")
    })
    defer monkey.Unpatch(utils.CheckUser)

    router.ServeHTTP(w, req)

    suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *BlogCommentControllerTestSuite) TestCreateComment_InternalServerError() {
    gin.SetMode(gin.TestMode)
    router := gin.Default()
    router.POST("/comment/:blog_id", suite.controller.CreateComment)

    commentRequest := &domain.CommentRequest{Content: "This is a comment"}
	expectedCommentRequest := &domain.CommentRequest{UserID: "user1", BlogID: "blog1", Content: "This is a comment"}
    suite.useCase.On("Create", mock.Anything, expectedCommentRequest).Return(nil, errors.New("internal error"))

    w := httptest.NewRecorder()
    reqBody, _ := json.Marshal(commentRequest)
    req, _ := http.NewRequest(http.MethodPost, "/comment/blog1", bytes.NewBuffer(reqBody))
    req.Header.Set("Content-Type", "application/json")

    monkey.Patch(utils.CheckUser, func(c *gin.Context) (*domain.AuthenticatedUser, error) {
        return &domain.AuthenticatedUser{UserID: "user1"}, nil
    })
    defer monkey.Unpatch(utils.CheckUser)

    router.ServeHTTP(w, req)

    suite.Equal(http.StatusInternalServerError, w.Code)
    suite.useCase.AssertExpectations(suite.T())
}

func (suite *BlogCommentControllerTestSuite) TestDeleteComment_Success() {
    gin.SetMode(gin.TestMode)
    router := gin.Default()
    router.DELETE("/comment/:id", suite.controller.DeleteCommment)

    suite.useCase.On("Delete", mock.Anything, "comment1").Return(&domain.Comment{CommentID: primitive.NewObjectID()}, nil)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest(http.MethodDelete, "/comment/comment1", nil)

    router.ServeHTTP(w, req)

    suite.Equal(http.StatusOK, w.Code)
    suite.useCase.AssertExpectations(suite.T())
}

func (suite *BlogCommentControllerTestSuite) TestDeleteComment_InternalServerError() {
    gin.SetMode(gin.TestMode)
    router := gin.Default()
    router.DELETE("/comment/:id", suite.controller.DeleteCommment)

    suite.useCase.On("Delete", mock.Anything, "comment1").Return(nil, errors.New("internal error"))

    w := httptest.NewRecorder()
    req, _ := http.NewRequest(http.MethodDelete, "/comment/comment1", nil)

    router.ServeHTTP(w, req)

    suite.Equal(http.StatusInternalServerError, w.Code)
    suite.useCase.AssertExpectations(suite.T())
}

func (suite *BlogCommentControllerTestSuite) TestUpdateComment_Success() {
    gin.SetMode(gin.TestMode)
    router := gin.Default()
    router.PUT("/comment/:id", suite.controller.UpdateComment)

    commentRequest := &domain.CommentRequest{Content: "Updated content"}
    suite.useCase.On("Update", mock.Anything, "Updated content", "comment1").Return(&domain.Comment{Content: "Updated content"}, nil)

    w := httptest.NewRecorder()
    reqBody, _ := json.Marshal(commentRequest)
    req, _ := http.NewRequest(http.MethodPut, "/comment/comment1", bytes.NewBuffer(reqBody))
    req.Header.Set("Content-Type", "application/json")

    router.ServeHTTP(w, req)

    suite.Equal(http.StatusOK, w.Code)
    suite.useCase.AssertExpectations(suite.T())
}

func (suite *BlogCommentControllerTestSuite) TestUpdateComment_InvalidRequestFormat() {
    gin.SetMode(gin.TestMode)
    router := gin.Default()
    router.PUT("/comment/:id", suite.controller.UpdateComment)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest(http.MethodPut, "/comment/comment1", bytes.NewBuffer([]byte("invalid json")))
    req.Header.Set("Content-Type", "application/json")

    router.ServeHTTP(w, req)

    suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *BlogCommentControllerTestSuite) TestUpdateComment_InternalServerError() {
    gin.SetMode(gin.TestMode)
    router := gin.Default()
    router.PUT("/comment/:id", suite.controller.UpdateComment)

    commentRequest := &domain.CommentRequest{Content: "Updated content"}
    suite.useCase.On("Update", mock.Anything, "Updated content", "comment1").Return(nil, errors.New("internal error"))

    w := httptest.NewRecorder()
    reqBody, _ := json.Marshal(commentRequest)
    req, _ := http.NewRequest(http.MethodPut, "/comment/comment1", bytes.NewBuffer(reqBody))
    req.Header.Set("Content-Type", "application/json")

    router.ServeHTTP(w, req)

    suite.Equal(http.StatusInternalServerError, w.Code)
    suite.useCase.AssertExpectations(suite.T())
}

func TestBlogCommentControllerTestSuite(t *testing.T) {
    suite.Run(t, new(BlogCommentControllerTestSuite))
}