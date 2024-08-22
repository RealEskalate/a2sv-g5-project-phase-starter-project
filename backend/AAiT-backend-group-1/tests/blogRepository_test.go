package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateBlog(t *testing.T) {
	mockRepo := new(mocks.MockBlogRepository)
	mockBlog := &domain.Blog{
		ID:      primitive.NewObjectID(),
		Title:     "Test Blog",
		Content:   "Test Content",
		AuthorID:  primitive.NewObjectID(),
		CreatedAt: time.Now(),
	}

	mockRepo.On("Create", mockBlog).Return(mockBlog, nil)

	result, err := mockRepo.Create(mockBlog)

	assert.Nil(t, err)
	assert.Equal(t, mockBlog, result)

	mockRepo.AssertExpectations(t)
}

func TestFindById(t *testing.T) {
	mockRepo := new(mocks.MockBlogRepository)
	mockBlog := &domain.Blog{
		ID:      primitive.NewObjectID(),
		Title:     "Test Blog",
		Content:   "Test Content",
		AuthorID:  primitive.NewObjectID(),
		CreatedAt: time.Now(),
	}

	mockRepo.On("FindById", "1").Return(mockBlog, nil)

	result, err := mockRepo.FindById("1")

	assert.Nil(t, err)
	assert.Equal(t, mockBlog, result)

	mockRepo.AssertExpectations(t)
}

func TestFindById_Error(t *testing.T) {
	mockRepo := new(mocks.MockBlogRepository)
	mockError := &domain.CustomError{
		Code:    404,
		Message: "Blog not found",
	}

	mockRepo.On("FindById", "8").Return(nil, mockError)

	result, err := mockRepo.FindById("8")
	fmt.Println(result, err)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, 404, err.StatusCode())
	assert.Equal(t, "Blog not found", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestUpdateBlog(t *testing.T) {
	mockRepo := new(mocks.MockBlogRepository)
	mockBlog := &domain.Blog{
		ID:        primitive.NewObjectID(),
		Title:     "Updated Blog",
		Content:   "Updated Content",
		AuthorID:  primitive.NewObjectID(),
		CreatedAt: time.Now(),
	}

	mockRepo.On("Update", "1", mockBlog).Return(mockBlog, nil)

	result, err := mockRepo.Update("1", mockBlog)

	assert.Nil(t, err)
	assert.Equal(t, mockBlog, result)

	mockRepo.AssertExpectations(t)
}

func TestDeleteBlog(t *testing.T) {
	mockRepo := new(mocks.MockBlogRepository)
	mockRepo.On("Delete", "1").Return(nil)

	err := mockRepo.Delete("1")

	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteBlog_Error(t *testing.T) {
	mockRepo := new(mocks.MockBlogRepository)
	mockError := &domain.CustomError{
		Code:    500,
		Message: "Internal Server Error",
	}

	mockRepo.On("Delete", "1").Return(mockError)

	err := mockRepo.Delete("1")

	assert.NotNil(t, err)
	assert.Equal(t, 500, err.StatusCode())
	assert.Equal(t, "Internal Server Error", err.Error())

	mockRepo.AssertExpectations(t)
}
