package test

import (
	"errors"
	"testing"
	"time"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/RealEskalate/blogpost/mocks"
	"github.com/RealEskalate/blogpost/usecase"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogUseCaseSuite struct {
	suite.Suite
	usecase *usecase.BlogUseCase
	repo    *mocks.Blog_Repository_interface
}

func (suite *BlogUseCaseSuite) SetupTest() {
	repo := new(mocks.Blog_Repository_interface)
	suite.usecase = usecase.NewBlogUsecase(repo)
	suite.repo = repo
}

func (suite *BlogUseCaseSuite) TestCreateBlog() {
	postBlog := domain.PostBlog{
		Title:   "Test Title",
		Content: "Test Content",
		Tag:     []string{"Test Tag"},
	}

	mockedID := primitive.NewObjectID()
	mockedTime := time.Now()

	// Create an expected blog with the mocked ID and time
	expectedBlog := domain.Blog{
		ID:        mockedID,
		Title:     postBlog.Title,
		Content:   postBlog.Content,
		CreatedAt: mockedTime,
		UpdatedAt: mockedTime,
		Tag:       postBlog.Tag,
	}

	suite.repo.On("CreateBlogDocument", mock.AnythingOfType("domain.Blog")).Return(expectedBlog, nil)

	createdBlog, err := suite.usecase.CreateBlog(postBlog)

	suite.NoError(err)
	suite.Equal(expectedBlog.Title, createdBlog.Title)
	suite.Equal(expectedBlog.Content, createdBlog.Content)
	suite.Equal(expectedBlog.Tag, createdBlog.Tag)
	suite.repo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseSuite) TestCreateBlog_Error() {
	postBlog := domain.PostBlog{
		Title:   "Test Title",
		Content: "Test Content",
		Tag:     []string{"Test Tag"},
	}
	expectedError := errors.New("creation failed")

	suite.repo.On("CreateBlogDocument", mock.AnythingOfType("domain.Blog")).Return(domain.Blog{}, expectedError)

	createdBlog, err := suite.usecase.CreateBlog(postBlog)

	suite.Error(err)
	suite.Equal(domain.Blog{}, createdBlog)
	suite.repo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseSuite) TestGetBlogs() {
	expectedBlogs := []domain.Blog{
		{
			ID:        primitive.NewObjectID(),
			Title:     "Test Title",
			Content:   "Test Content",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Tag:       []string{"Test Tag"},
		},
	}

	suite.repo.On("GetBlogDocuments", 1, 10).Return(expectedBlogs, nil)

	blogs, err := suite.usecase.GetBlogs(10, 1)

	suite.NoError(err)
	suite.Len(blogs, len(expectedBlogs))
	suite.repo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseSuite) TestGetBlogs_Error() {
	suite.repo.On("GetBlogDocuments", 1, 10).Return(nil, errors.New("failed to retrieve blogs"))

	blogs, err := suite.usecase.GetBlogs(10, 1)

	suite.Error(err)
	suite.Empty(blogs)
	suite.repo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseSuite) TestGetOneBlog() {
	expectedBlog := domain.Blog{
		ID:        primitive.NewObjectID(),
		Title:     "Test Title",
		Content:   "Test Content",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Tag:       []string{"Test Tag"},
	}

	suite.repo.On("GetOneBlogDocument", expectedBlog.ID.Hex()).Return(expectedBlog, nil)

	blog, err := suite.usecase.GetOneBlog(expectedBlog.ID.Hex())

	suite.NoError(err)
	suite.Equal(expectedBlog.Title, blog.Title)
	suite.repo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseSuite) TestGetOneBlog_Error() {
	id := primitive.NewObjectID().Hex()
	suite.repo.On("GetOneBlogDocument", id).Return(domain.Blog{}, errors.New("blog not found"))

	blog, err := suite.usecase.GetOneBlog(id)

	suite.Error(err)
	suite.Equal(domain.Blog{}, blog)
	suite.repo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseSuite) TestUpdateBlog() {
	id := primitive.NewObjectID().Hex()
	updateBlog := domain.Blog{
		Title:   "Updated Title",
		Content: "Updated Content",
	}

	expectedBlog := domain.Blog{
		ID:        primitive.NewObjectID(),
		Title:     "Updated Title",
		Content:   "Updated Content",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Tag:       []string{"Test Tag"},
	}

	suite.repo.On("UpdateBlogDocument", id, updateBlog).Return(expectedBlog, nil)

	blog, err := suite.usecase.UpdateBlog(id, updateBlog)

	suite.NoError(err)
	suite.Equal(expectedBlog.Title, blog.Title)
	suite.repo.AssertExpectations(suite.T())
}
func (suite *BlogUseCaseSuite) TestUpdateBlog_Error() {
	id := primitive.NewObjectID().Hex()
	updateBlog := domain.Blog{
		Title:   "Updated Title",
		Content: "Updated Content",
	}

	suite.repo.On("UpdateBlogDocument", id, updateBlog).Return(domain.Blog{}, errors.New("update failed"))

	blog, err := suite.usecase.UpdateBlog(id, updateBlog)

	suite.Error(err)
	suite.Equal(domain.Blog{}, blog)
	suite.repo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseSuite) TestDeleteBlog() {
	id := primitive.NewObjectID().Hex()

	suite.repo.On("DeleteBlogDocument", id).Return(nil)

	err := suite.usecase.DeleteBlog(id)

	suite.NoError(err)
	suite.repo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseSuite) TestDeleteBlog_Error() {
	id := primitive.NewObjectID().Hex()

	suite.repo.On("DeleteBlogDocument", id).Return(errors.New("delete failed"))

	err := suite.usecase.DeleteBlog(id)

	suite.Error(err)
	suite.repo.AssertExpectations(suite.T())
}

func TestBlogUseCaseSuite(t *testing.T) {
	suite.Run(t, new(BlogUseCaseSuite))
}
