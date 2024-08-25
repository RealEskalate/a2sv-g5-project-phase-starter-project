package controller

import (
    "Blog_Starter/domain"
    "Blog_Starter/domain/mocks"
    "Blog_Starter/utils"
    "bytes"
    "encoding/json"
    "errors"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"
	"go.mongodb.org/mongo-driver/bson/primitive"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/suite"
    "bou.ke/monkey"
)

type LikeControllerTestSuite struct {
    suite.Suite
    controller *LikeController
    useCase    *mocks.LikeUseCase
}

func (suite *LikeControllerTestSuite) SetupTest() {
    suite.useCase = new(mocks.LikeUseCase)
    suite.controller = NewLikeController(suite.useCase, time.Second*2)
}

func (suite *LikeControllerTestSuite) TestLikeBlog_Success() {
    gin.SetMode(gin.TestMode)
    router := gin.Default()
    router.POST("/like/:blog_id", suite.controller.LikeBlog)

    like := &domain.Like{UserID: "user1", BlogID: "blog1"}
    suite.useCase.On("LikeBlog", mock.Anything, like).Return(like, nil)

    w := httptest.NewRecorder()
    reqBody, _ := json.Marshal(like)
    req, _ := http.NewRequest(http.MethodPost, "/like/blog1", bytes.NewBuffer(reqBody))
    req.Header.Set("Content-Type", "application/json")

    monkey.Patch(utils.CheckUser, func(c *gin.Context) (*domain.AuthenticatedUser, error) {
        return &domain.AuthenticatedUser{UserID: "user1"}, nil
    })
    defer monkey.Unpatch(utils.CheckUser)

    router.ServeHTTP(w, req)

    suite.Equal(http.StatusOK, w.Code)
    suite.useCase.AssertExpectations(suite.T())
}

func (suite *LikeControllerTestSuite) TestLikeBlog_Unauthorized() {
    gin.SetMode(gin.TestMode)
    router := gin.Default()
    router.POST("/like/:blog_id", suite.controller.LikeBlog)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest(http.MethodPost, "/like/blog1", nil)
    req.Header.Set("Content-Type", "application/json")

    monkey.Patch(utils.CheckUser, func(c *gin.Context) (*domain.AuthenticatedUser, error) {
        return nil, errors.New("unauthorized")
    })
    defer monkey.Unpatch(utils.CheckUser)

    router.ServeHTTP(w, req)

    suite.Equal(http.StatusUnauthorized, w.Code)
}

func (suite *LikeControllerTestSuite) TestLikeBlog_InternalServerError() {
    gin.SetMode(gin.TestMode)
    router := gin.Default()
    router.POST("/like/:blog_id", suite.controller.LikeBlog)

    like := &domain.Like{UserID: "user1", BlogID: "blog1"}
    suite.useCase.On("LikeBlog", mock.Anything, like).Return(nil, errors.New("internal error"))

    w := httptest.NewRecorder()
    reqBody, _ := json.Marshal(like)
    req, _ := http.NewRequest(http.MethodPost, "/like/blog1", bytes.NewBuffer(reqBody))
    req.Header.Set("Content-Type", "application/json")

    monkey.Patch(utils.CheckUser, func(c *gin.Context) (*domain.AuthenticatedUser, error) {
        return &domain.AuthenticatedUser{UserID: "user1"}, nil
    })
    defer monkey.Unpatch(utils.CheckUser)

    router.ServeHTTP(w, req)

    suite.Equal(http.StatusInternalServerError, w.Code)
    suite.useCase.AssertExpectations(suite.T())
}

func (suite *LikeControllerTestSuite) TestUnlikeBlog_Success() {
    gin.SetMode(gin.TestMode)
    router := gin.Default()
    router.DELETE("/unlike/:id", suite.controller.UnlikeBlog)

    like := &domain.Like{LikeID: primitive.NewObjectID()}
    suite.useCase.On("UnlikeBlog", mock.Anything, "like1").Return(like, nil)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest(http.MethodDelete, "/unlike/like1", nil)

    router.ServeHTTP(w, req)

    suite.Equal(http.StatusOK, w.Code)
    suite.useCase.AssertExpectations(suite.T())
}

func (suite *LikeControllerTestSuite) TestUnlikeBlog_InternalServerError() {
    gin.SetMode(gin.TestMode)
    router := gin.Default()
    router.DELETE("/unlike/:id", suite.controller.UnlikeBlog)

    suite.useCase.On("UnlikeBlog", mock.Anything, "like1").Return(nil, errors.New("internal error"))

    w := httptest.NewRecorder()
    req, _ := http.NewRequest(http.MethodDelete, "/unlike/like1", nil)

    router.ServeHTTP(w, req)

    suite.Equal(http.StatusInternalServerError, w.Code)
    suite.useCase.AssertExpectations(suite.T())
}

func (suite *LikeControllerTestSuite) TestGetByID_Success() {
    gin.SetMode(gin.TestMode)
    router := gin.Default()
    router.GET("/like/:blog_id", suite.controller.GetByID)

    like := &domain.Like{UserID: "user1", BlogID: "blog1"}
    suite.useCase.On("GetByID", mock.Anything, "user1", "blog1").Return(like, nil)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest(http.MethodGet, "/like/blog1", nil)

    monkey.Patch(utils.CheckUser, func(c *gin.Context) (*domain.AuthenticatedUser, error) {
        return &domain.AuthenticatedUser{UserID: "user1"}, nil
    })
    defer monkey.Unpatch(utils.CheckUser)

    router.ServeHTTP(w, req)

    suite.Equal(http.StatusOK, w.Code)
    suite.useCase.AssertExpectations(suite.T())
}

func (suite *LikeControllerTestSuite) TestGetByID_Unauthorized() {
    gin.SetMode(gin.TestMode)
    router := gin.Default()
    router.GET("/like/:blog_id", suite.controller.GetByID)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest(http.MethodGet, "/like/blog1", nil)

    monkey.Patch(utils.CheckUser, func(c *gin.Context) (*domain.AuthenticatedUser, error) {
        return nil, errors.New("unauthorized")
    })
    defer monkey.Unpatch(utils.CheckUser)

    router.ServeHTTP(w, req)

    suite.Equal(http.StatusUnauthorized, w.Code)
}

func (suite *LikeControllerTestSuite) TestGetByID_InternalServerError() {
    gin.SetMode(gin.TestMode)
    router := gin.Default()
    router.GET("/like/:blog_id", suite.controller.GetByID)

    suite.useCase.On("GetByID", mock.Anything, "user1", "blog1").Return(nil, errors.New("internal error"))

    w := httptest.NewRecorder()
    req, _ := http.NewRequest(http.MethodGet, "/like/blog1", nil)

    monkey.Patch(utils.CheckUser, func(c *gin.Context) (*domain.AuthenticatedUser, error) {
        return &domain.AuthenticatedUser{UserID: "user1"}, nil
    })
    defer monkey.Unpatch(utils.CheckUser)

    router.ServeHTTP(w, req)

    suite.Equal(http.StatusInternalServerError, w.Code)
    suite.useCase.AssertExpectations(suite.T())
}

func TestLikeControllerTestSuite(t *testing.T) {
    suite.Run(t, new(LikeControllerTestSuite))
}