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

type BlogPopularityUseCaseTestSuite struct {
	suite.Suite
	mockRepo              *mocks.BlogPopularityRepository
	blogPopularityUseCase domain.BlogPopularityUsecase
}

func (suite *BlogPopularityUseCaseTestSuite) SetupTest() {
	suite.mockRepo = mocks.NewBlogPopularityRepository(suite.T())
	suite.blogPopularityUseCase = usecase.NewBlogPopularityUsecase(suite.mockRepo)
}

func (suite *BlogPopularityUseCaseTestSuite) TestGetSortedPopularBlogsByComments() {
	sortBy := "comments"
	sortOrder := 1

	expectedBlogs := []domain.Blog{
		{
			ID:           primitive.NewObjectID(),
			Title:        "Second Blog",
			Content:      "This is the second blog",
			LikeCount:    50,
			CommentCount: 20,
			DisLikeCount: 2,
		},
		{
			ID:           primitive.NewObjectID(),
			Title:        "Third Blog",
			Content:      "This is the third blog",
			LikeCount:    75,
			CommentCount: 15,
			DisLikeCount: 3,
		},
	}

	suite.mockRepo.On("GetPopularBlogs", sortBy, sortOrder).Return(expectedBlogs, nil)

	sortedBlogs, err := suite.blogPopularityUseCase.GetSortedPopularBlogs(sortBy, sortOrder)

	suite.NoError(err)
	suite.Len(sortedBlogs, len(expectedBlogs))
	suite.Equal(expectedBlogs, sortedBlogs)

	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *BlogPopularityUseCaseTestSuite) TestGetSortedPopularBlogsByDislikes() {
	sortBy := "dislikes"
	sortOrder := -1

	expectedBlogs := []domain.Blog{
		{
			ID:           primitive.NewObjectID(),
			Title:        "Fourth Blog",
			Content:      "This is the fourth blog",
			LikeCount:    30,
			CommentCount: 5,
			DisLikeCount: 10,
		},
	}

	suite.mockRepo.On("GetPopularBlogs", sortBy, sortOrder).
		Return(expectedBlogs, nil)

	blogs, err := suite.blogPopularityUseCase.GetSortedPopularBlogs(sortBy, sortOrder)

	suite.NoError(err)
	suite.Len(blogs, len(expectedBlogs))
	suite.ElementsMatch(expectedBlogs, blogs)

	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *BlogPopularityUseCaseTestSuite) TestGetSortedPopularBlogsEmptyResult() {
	sortBy := "likes"
	sortOrder := 1

	expectedBlogs := []domain.Blog{}

	suite.mockRepo.On("GetPopularBlogs", sortBy, sortOrder).Return(expectedBlogs, nil)

	blogs, err := suite.blogPopularityUseCase.GetSortedPopularBlogs(sortBy, sortOrder)

	suite.NoError(err)
	suite.Empty(blogs)

	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *BlogPopularityUseCaseTestSuite) TestGetSortedPopularBlogsError() {
	sortBy := "invalid"
	sortOrder := 1

	expectedError := errors.New("invalid sort parameter")

	suite.mockRepo.On("GetPopularBlogs", sortBy, sortOrder).Return(nil, expectedError)

	blogs, err := suite.blogPopularityUseCase.GetSortedPopularBlogs(sortBy, sortOrder)

	suite.Error(err)
	suite.Equal(expectedError, err)
	suite.Nil(blogs)

	suite.mockRepo.AssertExpectations(suite.T())
}

func TestBlogPopularityUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(BlogPopularityUseCaseTestSuite))
}
