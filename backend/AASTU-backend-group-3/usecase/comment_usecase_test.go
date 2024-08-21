package usecase

import (
	"group3-blogApi/domain"
	"group3-blogApi/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateComment(t *testing.T) {
	mockCommentRepo := new(mocks.CommentRepository)
	mockComment := &domain.Comment{
		PostID:  primitive.NewObjectID(),
		UserID:  primitive.NewObjectID(),
		Content: "This is a test comment",
	}

	t.Run("success", func(t *testing.T) {
		mockCommentRepo.On("CreateComment", mockComment).Return(mockComment, nil).Once()

		u := NewCommentUsecase(mockCommentRepo)
		result, err := u.CreateComment(mockComment)

		assert.NoError(t, err)
		assert.Equal(t, mockComment, result)
		mockCommentRepo.AssertExpectations(t)
	})

	t.Run("missing required fields", func(t *testing.T) {
		invalidComment := &domain.Comment{}

		u := NewCommentUsecase(mockCommentRepo)
		_, err := u.CreateComment(invalidComment)

		assert.Error(t, err)
		assert.Equal(t, "missing required fields", err.Error())
	})
}

func TestUpdateComment(t *testing.T) {
	mockCommentRepo := new(mocks.CommentRepository)
	mockComment := &domain.Comment{
		ID:      primitive.NewObjectID(),
		Content: "Updated comment",
	}

	t.Run("success", func(t *testing.T) {
		mockCommentRepo.On("UpdateComment", mockComment).Return(mockComment, nil).Once()

		u := NewCommentUsecase(mockCommentRepo)
		result, err := u.UpdateComment(mockComment)

		assert.NoError(t, err)
		assert.Equal(t, mockComment, result)
		mockCommentRepo.AssertExpectations(t)
	})

	t.Run("invalid comment ID", func(t *testing.T) {
		invalidComment := &domain.Comment{}

		u := NewCommentUsecase(mockCommentRepo)
		_, err := u.UpdateComment(invalidComment)

		assert.Error(t, err)
		assert.Equal(t, "invalid comment ID", err.Error())
	})
}

func TestDeleteComment(t *testing.T) {
	mockCommentRepo := new(mocks.CommentRepository)
	commentID := primitive.NewObjectID().Hex()

	t.Run("success", func(t *testing.T) {
		mockComment := &domain.Comment{
			ID: primitive.NewObjectID(),
		}
		mockCommentRepo.On("DeleteComment", mock.AnythingOfType("primitive.ObjectID")).Return(mockComment, nil).Once()

		u := NewCommentUsecase(mockCommentRepo)
		result, err := u.DeleteComment(commentID)

		assert.NoError(t, err)
		assert.Equal(t, mockComment, result)
		mockCommentRepo.AssertExpectations(t)
	})

	t.Run("invalid comment ID", func(t *testing.T) {
		u := NewCommentUsecase(mockCommentRepo)
		_, err := u.DeleteComment("invalid-id")

		assert.Error(t, err)
		assert.Equal(t, "invalid comment ID", err.Error())
	})
}

func TestGetCommentByID(t *testing.T) {
	mockCommentRepo := new(mocks.CommentRepository)
	commentID := primitive.NewObjectID().Hex()

	t.Run("success", func(t *testing.T) {
		mockComment := &domain.Comment{
			ID: primitive.NewObjectID(),
		}
		mockCommentRepo.On("GetCommentByID", mock.AnythingOfType("primitive.ObjectID")).Return(mockComment, nil).Once()

		u := NewCommentUsecase(mockCommentRepo)
		result, err := u.GetCommentByID(commentID)

		assert.NoError(t, err)
		assert.Equal(t, mockComment, result)
		mockCommentRepo.AssertExpectations(t)
	})

	t.Run("invalid comment ID", func(t *testing.T) {
		u := NewCommentUsecase(mockCommentRepo)
		_, err := u.GetCommentByID("invalid-id")

		assert.Error(t, err)
		assert.Equal(t, "invalid comment ID", err.Error())
	})
}

func TestGetComments(t *testing.T) {
	mockCommentRepo := new(mocks.CommentRepository)
	postID := primitive.NewObjectID().Hex()
	mockComments := []*domain.Comment{
		{
			ID: primitive.NewObjectID(),
		},
	}

	t.Run("success", func(t *testing.T) {
		mockCommentRepo.On("GetCommentsByPostID", postID, int64(1), int64(10)).Return(mockComments, nil).Once()

		u := NewCommentUsecase(mockCommentRepo)
		result, err := u.GetComments(postID, 1, 10)

		assert.NoError(t, err)
		assert.Equal(t, convertComments(mockComments), result)
		mockCommentRepo.AssertExpectations(t)
	})

	t.Run("invalid pagination parameters", func(t *testing.T) {
		u := NewCommentUsecase(mockCommentRepo)
		_, err := u.GetComments(postID, 0, 10)

		assert.Error(t, err)
		assert.Equal(t, "invalid pagination parameters", err.Error())
	})
}

func TestLikeComment(t *testing.T) {
	mockCommentRepo := new(mocks.CommentRepository)
	commentID := primitive.NewObjectID().Hex()
	userID := "user123"

	t.Run("success", func(t *testing.T) {
		mockCommentRepo.On("LikeComment", mock.AnythingOfType("primitive.ObjectID"), userID).Return(nil).Once()

		u := NewCommentUsecase(mockCommentRepo)
		err := u.LikeComment(commentID, userID)

		assert.NoError(t, err)
		mockCommentRepo.AssertExpectations(t)
	})

	t.Run("invalid comment ID", func(t *testing.T) {
		u := NewCommentUsecase(mockCommentRepo)
		err := u.LikeComment("invalid-id", userID)

		assert.Error(t, err)
		assert.Equal(t, "invalid comment ID", err.Error())
	})
}
