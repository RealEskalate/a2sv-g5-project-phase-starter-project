package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"group3-blogApi/mocks"
)

func TestLikeUsecase_LikeBlog(t *testing.T) {
	mockLikeRepo := new(mocks.LikeRepository)
	mockLikeUsecase := NewLikeUsecase(mockLikeRepo)

	userID := "user123"
	blogID := "blog123"
	Type := "like"

	t.Run("success", func(t *testing.T) {
		mockLikeRepo.On("LikeBlog", userID, blogID, Type).Return(nil).Once()

		err := mockLikeUsecase.LikeBlog(userID, blogID, Type)
		assert.NoError(t, err)

		mockLikeRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockLikeRepo.On("LikeBlog", userID, blogID, Type).Return(assert.AnError).Once()

		err := mockLikeUsecase.LikeBlog(userID, blogID, Type)
		assert.Error(t, err)

		mockLikeRepo.AssertExpectations(t)
	})
}

func TestLikeUsecase_DisLikeBlog(t *testing.T) {
	mockLikeRepo := new(mocks.LikeRepository)
	mockLikeUsecase := NewLikeUsecase(mockLikeRepo)

	userID := "user123"
	blogID := "blog123"
	Type := "dislike"

	t.Run("success", func(t *testing.T) {
		mockLikeRepo.On("DisLikeBlog", userID, blogID, Type).Return(nil).Once()

		err := mockLikeUsecase.DisLikeBlog(userID, blogID, Type)
		assert.NoError(t, err)

		mockLikeRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockLikeRepo.On("DisLikeBlog", userID, blogID, Type).Return(assert.AnError).Once()

		err := mockLikeUsecase.DisLikeBlog(userID, blogID, Type)
		assert.Error(t, err)

		mockLikeRepo.AssertExpectations(t)
	})
}
