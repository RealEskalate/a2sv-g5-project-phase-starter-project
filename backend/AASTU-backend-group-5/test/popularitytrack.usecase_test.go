package test

import (
	"errors"
	"testing"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/RealEskalate/blogpost/mocks"
	"github.com/RealEskalate/blogpost/usecase"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogPopularityUsecaseTestSuite struct {
	suite.Suite
	mockRepo              *mocks.BlogPopularityRepository
	blogPopularityUsecase domain.BlogPopularityUsecase
}

func (suite *BlogPopularityUsecaseTestSuite) SetupTest() {
	suite.mockRepo = new(mocks.BlogPopularityRepository)
	suite.blogPopularityUsecase = usecase.NewBlogPopularityUsecase(suite.mockRepo)
}

func (suite *BlogPopularityUsecaseTestSuite) TestGetSortedPopularBlogs_ValidInput() {
	sortBy := []domain.SortBy{domain.SortByLikeCount, domain.SortByCommentCount}
	sortOrder := []domain.SortOrder{domain.SortOrderDescending, domain.SortOrderAscending}

	expectedBlogs := []domain.Blog{
		{
			ID:           primitive.NewObjectID(),
			Title:        "Blog 1",
			Content:      "Content of blog 1",
			LikeCount:    100,
			CommentCount: 20,
			DisLikeCount: 5,
		},
		{
			ID:           primitive.NewObjectID(),
			Title:        "Blog 2",
			Content:      "Content of blog 2",
			LikeCount:    80,
			CommentCount: 15,
			DisLikeCount: 3,
		},
	}

	suite.mockRepo.On("GetPopularBlogs", sortBy, sortOrder).Return(expectedBlogs, nil)

	blogs, err := suite.blogPopularityUsecase.GetSortedPopularBlogs(sortBy, sortOrder)

	suite.NoError(err)
	suite.Equal(expectedBlogs, blogs)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *BlogPopularityUsecaseTestSuite) TestGetSortedPopularBlogs_EmptyInput() {
	sortBy := []domain.SortBy{}
	sortOrder := []domain.SortOrder{}

	expectedBlogs := []domain.Blog{}

	suite.mockRepo.On("GetPopularBlogs", sortBy, sortOrder).Return(expectedBlogs, nil)

	blogs, err := suite.blogPopularityUsecase.GetSortedPopularBlogs(sortBy, sortOrder)

	suite.NoError(err)
	suite.Empty(blogs)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *BlogPopularityUsecaseTestSuite) TestGetSortedPopularBlogs_RepoError() {
	sortBy := []domain.SortBy{domain.SortByLikeCount}
	sortOrder := []domain.SortOrder{domain.SortOrderDescending}

	expectedError := errors.New("repository error")

	suite.mockRepo.On("GetPopularBlogs", sortBy, sortOrder).Return(nil, expectedError)

	blogs, err := suite.blogPopularityUsecase.GetSortedPopularBlogs(sortBy, sortOrder)

	suite.Error(err)
	suite.Nil(blogs)
	suite.Equal(expectedError, err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func TestBlogPopularityUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(BlogPopularityUsecaseTestSuite))
}
