package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"aait.backend.g10/delivery/controllers"
	"aait.backend.g10/domain"
	"aait.backend.g10/tests/mocks"
	"aait.backend.g10/usecases/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LikeControllerSuite struct {
	suite.Suite
	controller  *controllers.LikeController
	mockUseCase *mocks.LikeUsecaseInterface
}

func (s *LikeControllerSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	s.mockUseCase = new(mocks.LikeUsecaseInterface)
	s.controller = &controllers.LikeController{
		LikeUseCase: s.mockUseCase,
	}
}

func (s *LikeControllerSuite) TearDownTest() {
	s.mockUseCase.AssertExpectations(s.T())
}

func (s *LikeControllerSuite) TestLikeBlog_Success() {
	isLike := true
	like := domain.Like{
		BlogID:    uuid.New(),
		ReacterID: uuid.New(),
		IsLike:    &isLike,
	}

	s.mockUseCase.On("LikeBlog", like).Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/like", bytes.NewBufferString(`{
		"blog_id": "`+like.BlogID.String()+`",
		"is_like": true
	}`))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("id", like.ReacterID.String())

	s.controller.LikeBlog(c)

	assert.Equal(s.T(), http.StatusCreated, w.Code)
	assert.Equal(s.T(), `{"message":"Like added successfully"}`, w.Body.String())
}

func (s *LikeControllerSuite) TestLikeBlog_InvalidJSON() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/like", bytes.NewBufferString(`{invalid-json}`))
	c.Request.Header.Set("Content-Type", "application/json")

	s.controller.LikeBlog(c)

	assert.Equal(s.T(), http.StatusBadRequest, w.Code)
	assert.Contains(s.T(), w.Body.String(), "invalid character")
}

func (s *LikeControllerSuite) TestLikeBlog_InvalidReacterID() {
	like := domain.Like{
		BlogID: uuid.New(),
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/like", bytes.NewBufferString(`{
		"blog_id": "`+like.BlogID.String()+`",
		"is_like": true
	}`))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("id", "invalid-uuid")

	s.controller.LikeBlog(c)

	assert.Equal(s.T(), http.StatusBadRequest, w.Code)
	assert.Contains(s.T(), w.Body.String(), "Invalid ID")
}

func (s *LikeControllerSuite) TestLikeBlog_Failure() {
	isLike := true
	like := domain.Like{
		BlogID:    uuid.New(),
		ReacterID: uuid.New(),
		IsLike:    &isLike,
	}

	s.mockUseCase.On("LikeBlog", like).Return(&domain.CustomError{
		StatusCode: http.StatusInternalServerError,
		Message:    "failed to like blog",
	})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/like", bytes.NewBufferString(`{
		"blog_id": "`+like.BlogID.String()+`",
		"is_like": true
	}`))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("id", like.ReacterID.String())

	s.controller.LikeBlog(c)

	assert.Equal(s.T(), http.StatusInternalServerError, w.Code)
	assert.Equal(s.T(), `{"error":"failed to like blog"}`, w.Body.String())
}

func (s *LikeControllerSuite) TestDeleteLike_Success() {
	likeDto := dto.UnlikeDto{
		BlogID:    uuid.New(),
		ReacterID: uuid.New(),
	}

	s.mockUseCase.On("DeleteLike", likeDto).Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/unlike", bytes.NewBufferString(`{
		"blog_id": "`+likeDto.BlogID.String()+`"
	}`))

	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("id", likeDto.ReacterID.String())

	s.controller.DeleteLike(c)

	assert.Equal(s.T(), http.StatusOK, w.Code)
	assert.Equal(s.T(), `{"message":"Like deleted successfully"}`, w.Body.String())
}

func (s *LikeControllerSuite) TestDeleteLike_InvalidJSON() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/unlike", bytes.NewBufferString(`{invalid-json}`))
	c.Request.Header.Set("Content-Type", "application/json")

	s.controller.DeleteLike(c)

	assert.Equal(s.T(), http.StatusBadRequest, w.Code)
	assert.Contains(s.T(), w.Body.String(), "invalid character")
}

func (s *LikeControllerSuite) TestDeleteLike_InvalidReacterID() {
	likeDto := dto.UnlikeDto{
		BlogID: uuid.New(),
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/unlike", bytes.NewBufferString(`{
		"blog_id": "`+likeDto.BlogID.String()+`"
	}`))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("id", "invalid-uuid")

	s.controller.DeleteLike(c)

	assert.Equal(s.T(), http.StatusBadRequest, w.Code)
	assert.Contains(s.T(), w.Body.String(), "Invalid ID")
}

func (s *LikeControllerSuite) TestDeleteLike_Failure() {
	likeDto := dto.UnlikeDto{
		BlogID:    uuid.New(),
		ReacterID: uuid.New(),
	}

	s.mockUseCase.On("DeleteLike", likeDto).Return(&domain.CustomError{
		StatusCode: http.StatusInternalServerError,
		Message:    "failed to delete like",
	})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/unlike", bytes.NewBufferString(`{
		"blog_id": "`+likeDto.BlogID.String()+`"
	}`))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("id", likeDto.ReacterID.String())

	s.controller.DeleteLike(c)

	assert.Equal(s.T(), http.StatusInternalServerError, w.Code)
	assert.Equal(s.T(), `{"error":"failed to delete like"}`, w.Body.String())
}

func TestLikeControllerSuite(t *testing.T) {
	suite.Run(t, new(LikeControllerSuite))
}
