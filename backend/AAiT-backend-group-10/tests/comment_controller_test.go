package tests

import (
	"bytes"
	"encoding/json"
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

type CommentControllerSuite struct {
	suite.Suite
	controller  *controllers.CommentController
	mockUseCase *mocks.CommentUsecaseInterface
}

func (s *CommentControllerSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	s.mockUseCase = new(mocks.CommentUsecaseInterface)
	s.controller = &controllers.CommentController{
		CommentUsecase: s.mockUseCase,
	}
}

func (s *CommentControllerSuite) TearDownTest() {
	s.mockUseCase.AssertExpectations(s.T())
}

func (s *CommentControllerSuite) TestGetComments_Success() {
	blogID := uuid.New()
	mockComments := []*dto.CommentDto{
		{ID: uuid.New(), Comment: "Test Comment 1", BlogID: blogID},
		{ID: uuid.New(), Comment: "Test Comment 2", BlogID: blogID},
	}

	s.mockUseCase.On("GetComments", blogID).Return(mockComments, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "blog_id", Value: blogID.String()}}

	s.controller.GetComments(c)

	assert.Equal(s.T(), http.StatusOK, w.Code)
	var response map[string][]*dto.CommentDto
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), mockComments, response["comments"])
}

func (s *CommentControllerSuite) TestGetComments_InvalidID() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "blog_id", Value: "invalid-uuid"}}

	s.controller.GetComments(c)

	assert.Equal(s.T(), http.StatusBadRequest, w.Code)
}

func (s *CommentControllerSuite) TestGetComments_Failure() {
	blogID := uuid.New()
	s.mockUseCase.On("GetComments", blogID).Return(nil, &domain.CustomError{
		StatusCode: http.StatusInternalServerError,
		Message:    "failed to get comments",
	})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "blog_id", Value: blogID.String()}}

	s.controller.GetComments(c)

	assert.Equal(s.T(), http.StatusInternalServerError, w.Code)
	assert.Equal(s.T(), `{"error":"failed to get comments"}`, w.Body.String())
}

func (s *CommentControllerSuite) TestAddComment_Success() {
	commenterID := uuid.New()
	comment := domain.Comment{
		Comment: "Test Comment",
		BlogID:  uuid.New(),
		CommenterID: commenterID,
	}

	s.mockUseCase.On("AddComment", comment).Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/comments", bytes.NewBufferString(`{
		"comment": "Test Comment",
		"blog_id": "`+comment.BlogID.String()+`"
	}`))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("id", commenterID.String())

	s.controller.AddComment(c)

	assert.Equal(s.T(), http.StatusCreated, w.Code)
	assert.Equal(s.T(), `{"message":"Comment added successfully"}`, w.Body.String())
}

func (s *CommentControllerSuite) TestAddComment_InvalidJSON() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/comments", bytes.NewBufferString(`{invalid-json}`))
	c.Request.Header.Set("Content-Type", "application/json")

	s.controller.AddComment(c)

	assert.Equal(s.T(), http.StatusBadRequest, w.Code)
	assert.Contains(s.T(), w.Body.String(), "invalid character")
}

func (s *CommentControllerSuite) TestAddComment_Failure() {
	commenterID := uuid.New()
	comment := domain.Comment{
		Comment: "Test Comment",
		BlogID:  uuid.New(),
		CommenterID: commenterID,
	}

	s.mockUseCase.On("AddComment", comment).Return(&domain.CustomError{
		StatusCode: http.StatusInternalServerError,
		Message:    "failed to add comment",
	})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/comments", bytes.NewBufferString(`{
		"comment": "Test Comment",
		"blog_id": "`+comment.BlogID.String()+`"
	}`))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("id", commenterID.String())

	s.controller.AddComment(c)

	assert.Equal(s.T(), http.StatusInternalServerError, w.Code)
	assert.Equal(s.T(), `{"error":"failed to add comment"}`, w.Body.String())
}


func (s *CommentControllerSuite) TestUpdateComment_Success() {
	commentID := uuid.New()
	requesterID := uuid.New()
	comment := domain.Comment{
		ID:      commentID,
		Comment: "Updated Comment",
		BlogID:  uuid.New(),
		CommenterID: requesterID,
	}

	s.mockUseCase.On("UpdateComment", requesterID, comment).Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPut, "/comments/"+commentID.String(), bytes.NewBufferString(`{
		"comment":"Updated Comment",
		"blog_id":"`+comment.BlogID.String()+`"
	}`))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = gin.Params{{Key: "id", Value: commentID.String()}}
	c.Set("id", requesterID.String())

	s.controller.UpdateComment(c)

	assert.Equal(s.T(), http.StatusOK, w.Code)
	assert.Contains(s.T(), w.Body.String(), "Comment updated successfully")
}

func (s *CommentControllerSuite) TestUpdateComment_InvalidID() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "invalid-uuid"}}
	c.Request = httptest.NewRequest(http.MethodPut, "/comments/invalid-uuid", nil)

	s.controller.UpdateComment(c)

	assert.Equal(s.T(), http.StatusBadRequest, w.Code)
}

func (s *CommentControllerSuite) TestUpdateComment_InvalidJSON() {
	commentID := uuid.New()
	requesterID := uuid.New()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: commentID.String()}}
	c.Request = httptest.NewRequest(http.MethodPut, "/comments/"+commentID.String(), bytes.NewBufferString(`{invalid-json}`))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("id", requesterID.String())

	s.controller.UpdateComment(c)

	assert.Equal(s.T(), http.StatusBadRequest, w.Code)
	assert.Contains(s.T(), w.Body.String(), "invalid character")
}

func (s *CommentControllerSuite) TestUpdateComment_Failure() {
	commentID := uuid.New()
	requesterID := uuid.New()
	comment := domain.Comment{
		ID:      commentID,
		Comment: "Updated Comment",
		BlogID:  uuid.New(),
		CommenterID: requesterID,
	}

	s.mockUseCase.On("UpdateComment", requesterID, comment).Return(&domain.CustomError{
		StatusCode: http.StatusInternalServerError,
		Message:    "failed to update comment",
	})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPut, "/comments/"+commentID.String(), bytes.NewBufferString(`{
		"comment":"Updated Comment",
		"blog_id":"`+comment.BlogID.String()+`"
	}`))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = gin.Params{{Key: "id", Value: commentID.String()}}
	c.Set("id", requesterID.String())

	s.controller.UpdateComment(c)

	assert.Equal(s.T(), http.StatusInternalServerError, w.Code)
	assert.Equal(s.T(), `{"error":"failed to update comment"}`, w.Body.String())
}


func (s *CommentControllerSuite) TestDeleteComment_Success() {
	commentID := uuid.New()
	requesterID := uuid.New()
	requesterRole := true // Admin role

	s.mockUseCase.On("DeleteComment", commentID, requesterID, requesterRole).Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: commentID.String()}}
	c.Set("id", requesterID.String())
	c.Set("is_admin", requesterRole)

	s.controller.DeleteComment(c)

	assert.Equal(s.T(), http.StatusOK, w.Code)
	assert.Contains(s.T(), w.Body.String(), "Comment deleted successfully")
}

func (s *CommentControllerSuite) TestDeleteComment_InvalidID() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "invalid-uuid"}}
	c.Request = httptest.NewRequest(http.MethodDelete, "/comments/invalid-uuid", nil)

	s.controller.DeleteComment(c)

	assert.Equal(s.T(), http.StatusBadRequest, w.Code)
}

func (s *CommentControllerSuite) TestDeleteComment_Failure() {
	commentID := uuid.New()
	requesterID := uuid.New()
	requesterRole := false // Not an admin

	s.mockUseCase.On("DeleteComment", commentID, requesterID, requesterRole).Return(&domain.CustomError{
		StatusCode: http.StatusInternalServerError,
		Message:    "failed to delete comment",
	})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: commentID.String()}}
	c.Set("id", requesterID.String())
	c.Set("is_admin", requesterRole)

	s.controller.DeleteComment(c)

	assert.Equal(s.T(), http.StatusInternalServerError, w.Code)
	assert.Equal(s.T(), `{"error":"failed to delete comment"}`, w.Body.String())
}

func TestCommentControllerSuite(t *testing.T) {
	suite.Run(t, new(CommentControllerSuite))
}
