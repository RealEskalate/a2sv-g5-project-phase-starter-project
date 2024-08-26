package usecase_test

// import (
// 	"errors"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"go.mongodb.org/mongo-driver/bson/primitive"

// 	"group3-blogApi/domain"
// 	"group3-blogApi/mocks"
// 	"group3-blogApi/usecase"
// )

// func TestCreateBlog(t *testing.T) {
// 	mockBlogRepo := new(mocks.BlogRepository)
// 	mockID := primitive.NewObjectID()
// 	mockBlog := domain.Blog{
// 		ID:       mockID,
// 		Title:    "Test Blog",
// 		Content:  "This is a test blog",
// 		AuthorID: "user123",
// 	}
// 	username := "testuser"
// 	userID := "user123"

// 	t.Run("success", func(t *testing.T) {
// 		mockBlogRepo.On("CreateBlog", username, userID, mockBlog).Return(mockBlog, nil).Once()

// 		uc := usecase.NewBlogUsecase(mockBlogRepo)
// 		createdBlog, err := uc.CreateBlog(username, userID, mockBlog)

// 		assert.NoError(t, err)
// 		assert.Equal(t, mockBlog, createdBlog)
// 		mockBlogRepo.AssertExpectations(t)
// 	})

// 	t.Run("error", func(t *testing.T) {
// 		mockBlogRepo.On("CreateBlog", username, userID, mockBlog).Return(domain.Blog{}, errors.New("unexpected error")).Once()

// 		uc := usecase.NewBlogUsecase(mockBlogRepo)
// 		_, err := uc.CreateBlog(username, userID, mockBlog)

// 		assert.Error(t, err)
// 		assert.Equal(t, "unexpected error", err.Error())
// 		mockBlogRepo.AssertExpectations(t)
// 	})
// }

// func TestDeleteBlog(t *testing.T) {
// 	mockBlogRepo := new(mocks.BlogRepository)
// 	mockID := primitive.NewObjectID() // Use primitive.ObjectID
// 	mockBlog := domain.Blog{
// 		ID:       mockID,
// 		AuthorID: "user123",
// 	}

// 	t.Run("success with admin", func(t *testing.T) {
// 		mockBlogRepo.On("DeleteBlog", mockID.Hex()).Return(mockBlog, nil).Once()

// 		uc := usecase.NewBlogUsecase(mockBlogRepo)
// 		deletedBlog, err := uc.DeleteBlog("admin", "user123", mockID.Hex())

// 		assert.NoError(t, err)
// 		assert.Equal(t, mockBlog, deletedBlog)
// 		mockBlogRepo.AssertExpectations(t)
// 	})

// 	t.Run("success with user", func(t *testing.T) {
// 		mockBlogRepo.On("GetBlogByID", mockID.Hex()).Return(mockBlog, nil).Once()
// 		mockBlogRepo.On("DeleteBlog", mockID.Hex()).Return(mockBlog, nil).Once()

// 		uc := usecase.NewBlogUsecase(mockBlogRepo)
// 		deletedBlog, err := uc.DeleteBlog("user", "user123", mockID.Hex())

// 		assert.NoError(t, err)
// 		assert.Equal(t, mockBlog, deletedBlog)
// 		mockBlogRepo.AssertExpectations(t)
// 	})

// 	t.Run("unauthorized user", func(t *testing.T) {
// 		unauthorizedBlog := domain.Blog{
// 			ID:       mockID,
// 			AuthorID: "anotherUser",
// 		}
// 		mockBlogRepo.On("GetBlogByID", mockID.Hex()).Return(unauthorizedBlog, nil).Once()

// 		uc := usecase.NewBlogUsecase(mockBlogRepo)
// 		_, err := uc.DeleteBlog("user", "user123", mockID.Hex())

// 		assert.Error(t, err)
// 		assert.Equal(t, "unauthorized to delete blog", err.Error())
// 		mockBlogRepo.AssertExpectations(t)
// 	})

// 	t.Run("error when getting blog by ID", func(t *testing.T) {
// 		mockBlogRepo.On("GetBlogByID", mockID.Hex()).Return(domain.Blog{}, errors.New("unexpected error")).Once()

// 		uc := usecase.NewBlogUsecase(mockBlogRepo)
// 		_, err := uc.DeleteBlog("user", "user123", mockID.Hex())

// 		assert.Error(t, err)
// 		assert.Equal(t, "unexpected error", err.Error())
// 		mockBlogRepo.AssertExpectations(t)
// 	})
// }

// func TestUpdateBlog(t *testing.T) {
// 	mockBlogRepo := new(mocks.BlogRepository)
// 	mockID := primitive.NewObjectID() // Use primitive.ObjectID
// 	mockBlog := domain.Blog{
// 		ID:       mockID,
// 		AuthorID: "user123",
// 		Title:    "Updated Title",
// 		Content:  "Updated Content",
// 	}

// 	t.Run("success", func(t *testing.T) {
// 		mockBlogRepo.On("GetBlogByID", mockID.Hex()).Return(mockBlog, nil).Once()
// 		mockBlogRepo.On("UpdateBlog", mockBlog, mockID.Hex()).Return(mockBlog, nil).Once()

// 		uc := usecase.NewBlogUsecase(mockBlogRepo)
// 		updatedBlog, err := uc.UpdateBlog(mockBlog, "user", mockID.Hex())

// 		assert.NoError(t, err)
// 		assert.Equal(t, mockBlog, updatedBlog)
// 		mockBlogRepo.AssertExpectations(t)
// 	})

// 	t.Run("unauthorized", func(t *testing.T) {
// 		mockBlogRepo.On("GetBlogByID", mockID.Hex()).Return(domain.Blog{
// 			ID:       mockID,
// 			AuthorID: "anotherUser",
// 		}, nil).Once()

// 		uc := usecase.NewBlogUsecase(mockBlogRepo)
// 		_, err := uc.UpdateBlog(mockBlog, "user", mockID.Hex())

// 		assert.Error(t, err)
// 		assert.Equal(t, "unauthorized to update blog", err.Error())
// 		mockBlogRepo.AssertExpectations(t)
// 	})

// 	t.Run("error when getting blog by ID", func(t *testing.T) {
// 		mockBlogRepo.On("GetBlogByID", mockID.Hex()).Return(domain.Blog{}, errors.New("unexpected error")).Once()

// 		uc := usecase.NewBlogUsecase(mockBlogRepo)
// 		_, err := uc.UpdateBlog(mockBlog, "user", mockID.Hex())

// 		assert.Error(t, err)
// 		assert.Equal(t, "unexpected error", err.Error())
// 		mockBlogRepo.AssertExpectations(t)
// 	})

// 	t.Run("error when updating blog", func(t *testing.T) {
// 		mockBlogRepo.On("GetBlogByID", mockID.Hex()).Return(mockBlog, nil).Once()
// 		mockBlogRepo.On("UpdateBlog", mockBlog, mockID.Hex()).Return(domain.Blog{}, errors.New("unexpected error")).Once()

// 		uc := usecase.NewBlogUsecase(mockBlogRepo)
// 		_, err := uc.UpdateBlog(mockBlog, "user", mockID.Hex())

// 		assert.Error(t, err)
// 		assert.Equal(t, "unexpected error", err.Error())
// 		mockBlogRepo.AssertExpectations(t)
// 	})
// }

// func TestGetBlogByID(t *testing.T) {
// 	mockBlogRepo := new(mocks.BlogRepository)
// 	mockID := primitive.NewObjectID() // Use primitive.ObjectID
// 	mockBlog := domain.Blog{
// 		ID:       mockID,
// 		AuthorID: "user123",
// 		Title:    "Test Blog",
// 		Content:  "This is a test blog",
// 	}

// 	t.Run("success", func(t *testing.T) {
// 		mockBlogRepo.On("GetBlogByID", mockID.Hex()).Return(mockBlog, nil).Once()

// 		uc := usecase.NewBlogUsecase(mockBlogRepo)
// 		blog, err := uc.GetBlogByID(mockID.Hex())

// 		assert.NoError(t, err)
// 		assert.Equal(t, mockBlog, blog)
// 		mockBlogRepo.AssertExpectations(t)
// 	})

// 	t.Run("error", func(t *testing.T) {
// 		mockBlogRepo.On("GetBlogByID", mockID.Hex()).Return(domain.Blog{}, errors.New("unexpected error")).Once()

// 		uc := usecase.NewBlogUsecase(mockBlogRepo)
// 		_, err := uc.GetBlogByID(mockID.Hex())

// 		assert.Error(t, err)
// 		assert.Equal(t, "unexpected error", err.Error())
// 		mockBlogRepo.AssertExpectations(t)
// 	})
// }

// func TestGetBlogs(t *testing.T) {
// 	mockBlogRepo := new(mocks.BlogRepository)
// 	mockID1 := primitive.NewObjectID() // Use primitive.ObjectID
// 	mockID2 := primitive.NewObjectID() // Use primitive.ObjectID
// 	mockBlogs := []domain.Blog{
// 		{
// 			ID:       mockID1,
// 			AuthorID: "user123",
// 			Title:    "Test Blog 1",
// 			Content:  "This is a test blog 1",
// 		},
// 		{
// 			ID:       mockID2,
// 			AuthorID: "user124",
// 			Title:    "Test Blog 2",
// 			Content:  "This is a test blog 2",
// 		},
// 	}

// 	t.Run("success", func(t *testing.T) {
// 		mockBlogRepo.On("GetBlogs", int64(1), int64(10), "date", "", "").Return(mockBlogs, nil).Once()

// 		uc := usecase.NewBlogUsecase(mockBlogRepo)
// 		blogs, err := uc.GetBlogs(1, 10, "date", "", "")

// 		assert.NoError(t, err)
// 		assert.Equal(t, mockBlogs, blogs)
// 		mockBlogRepo.AssertExpectations(t)
// 	})

// 	t.Run("error", func(t *testing.T) {
// 		mockBlogRepo.On("GetBlogs", int64(1), int64(10), "date", "", "").Return(nil, errors.New("unexpected error")).Once()

// 		uc := usecase.NewBlogUsecase(mockBlogRepo)
// 		_, err := uc.GetBlogs(1, 10, "date", "", "")

// 		assert.Error(t, err)
// 		assert.Equal(t, "unexpected error", err.Error())
// 		mockBlogRepo.AssertExpectations(t)
// 	})
// }
