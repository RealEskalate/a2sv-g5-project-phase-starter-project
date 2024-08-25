package usecase_test

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"go.mongodb.org/mongo-driver/bson/primitive"

// 	"group3-blogApi/domain"
// 	"group3-blogApi/mocks"
// 	"group3-blogApi/usecase"
// )

// func TestCreateComment(t *testing.T) {
// 	mockRepo := new(mocks.CommentRepository)
// 	userMockRepo := new(mocks.UserRepository)
// 	uc := usecase.NewCommentUsecase(mockRepo, userMockRepo)

// 	t.Run("success", func(t *testing.T) {
// 		comment := &domain.Comment{
// 			PostID:  primitive.NewObjectID(),
// 			UserID:  primitive.NewObjectID(),
// 			Content: "This is a comment",
// 		}

// 		mockRepo.On("CreateComment", comment).Return(comment, nil).Once()

// 		result, err := uc.CreateComment(comment)
// 		assert.NoError(t, err)
// 		assert.Equal(t, comment, result)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("missing required fields", func(t *testing.T) {
// 		comment := &domain.Comment{}

// 		result, err := uc.CreateComment(comment)
// 		assert.Error(t, err)
// 		assert.Nil(t, result)
// 		assert.Equal(t, "missing required fields", err.Error())
// 	})
// }

// func TestUpdateComment(t *testing.T) {
// 	mockRepo := new(mocks.CommentRepository)
// 	uc := usecase.NewCommentUsecase(mockRepo)

// 	t.Run("success", func(t *testing.T) {
// 		comment := &domain.Comment{
// 			ID:      primitive.NewObjectID(),
// 			UserID:  primitive.NewObjectID(),
// 			Content: "Updated content",
// 		}

// 		existingComment := *comment

// 		mockRepo.On("GetCommentByID", comment.ID).Return(&existingComment, nil).Once()
// 		mockRepo.On("UpdateComment", comment).Return(comment, nil).Once()

// 		result, err := uc.UpdateComment(comment, "admin", comment.UserID.Hex())
// 		assert.NoError(t, err)
// 		assert.Equal(t, comment, result)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("unauthorized", func(t *testing.T) {
// 		comment := &domain.Comment{
// 			ID:      primitive.NewObjectID(),
// 			UserID:  primitive.NewObjectID(),
// 			Content: "Updated content",
// 		}

// 		existingComment := *comment

// 		mockRepo.On("GetCommentByID", comment.ID).Return(&existingComment, nil).Once()

// 		result, err := uc.UpdateComment(comment, "user", "differentUserID")
// 		assert.Error(t, err)
// 		assert.Nil(t, result)
// 		assert.Equal(t, "unauthorized", err.Error())
// 	})
// }

// func TestDeleteComment(t *testing.T) {
// 	mockRepo := new(mocks.CommentRepository)
// 	uc := usecase.NewCommentUsecase(mockRepo)

// 	t.Run("success", func(t *testing.T) {
// 		commentID := primitive.NewObjectID()
// 		comment := &domain.Comment{
// 			ID:     commentID,
// 			UserID: primitive.NewObjectID(),
// 		}

// 		mockRepo.On("GetCommentByID", commentID).Return(comment, nil).Once()
// 		mockRepo.On("DeleteComment", commentID).Return(comment, nil).Once()

// 		result, err := uc.DeleteComment(commentID.Hex(), "admin", comment.UserID.Hex())
// 		assert.NoError(t, err)
// 		assert.Equal(t, comment, result)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("unauthorized", func(t *testing.T) {
// 		commentID := primitive.NewObjectID()
// 		comment := &domain.Comment{
// 			ID:     commentID,
// 			UserID: primitive.NewObjectID(),
// 		}

// 		mockRepo.On("GetCommentByID", commentID).Return(comment, nil).Once()

// 		result, err := uc.DeleteComment(commentID.Hex(), "user", "differentUserID")
// 		assert.Error(t, err)
// 		assert.Nil(t, result)
// 		assert.Equal(t, "unauthorized", err.Error())
// 	})
// }

// // Additional test cases for other methods can be added in a similar fashion

// func TestGetCommentByID(t *testing.T) {
// 	mockRepo := new(mocks.CommentRepository)
// 	uc := usecase.NewCommentUsecase(mockRepo)

// 	t.Run("success", func(t *testing.T) {
// 		commentID := primitive.NewObjectID()
// 		expectedComment := &domain.Comment{ID: commentID}

// 		mockRepo.On("GetCommentByID", commentID).Return(expectedComment, nil).Once()

// 		result, err := uc.GetCommentByID(commentID.Hex())
// 		assert.NoError(t, err)
// 		assert.Equal(t, expectedComment, result)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("invalid comment ID", func(t *testing.T) {
// 		_, err := uc.GetCommentByID("invalidID")
// 		assert.Error(t, err)
// 		assert.Equal(t, "invalid comment ID", err.Error())
// 	})
// }

// // Similar tests can be written for the reply functions
