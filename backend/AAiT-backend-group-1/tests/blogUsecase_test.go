package tests

import (
	"testing"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/mocks"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogUseCaseTestSuite struct {
	suite.Suite
	mockBlogRepo *mocks.MockBlogRepository
	mockCache    *mocks.MockCacheService
	blogUseCase  domain.BlogUseCase
}

func (suite *BlogUseCaseTestSuite) SetupTest() {
	suite.mockBlogRepo = new(mocks.MockBlogRepository)
	suite.mockCache = new(mocks.MockCacheService)
	suite.blogUseCase = usecases.NewBlogUseCase(suite.mockBlogRepo, suite.mockCache)
}

func (suite *BlogUseCaseTestSuite) TestCreateBlog() {
	authorID := primitive.NewObjectID().Hex()
	blog := &domain.Blog{
		Title:   "Sample Title",
		Content: "Sample Content",
		Tags:    []string{"Go", "Testing"},
	}

	suite.mockBlogRepo.On("Create", blog).Return(blog, nil)
	suite.mockCache.On("Delete", "all_blogs").Return(nil)

	err := suite.blogUseCase.CreateBlog(blog, authorID)

	assert.NoError(suite.T(), err)
	suite.mockBlogRepo.AssertExpectations(suite.T())
	suite.mockCache.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestGetBlog() {
	blogID := primitive.NewObjectID().Hex()
	blog := &domain.Blog{
		ID:      primitive.NewObjectID(),
		Title:   "Sample Title",
		Content: "Sample Content",
	}

	cacheBlogJson := `{"id":"` + blog.ID.Hex() + `","title":"Sample Title","content":"Sample Content"}`

	suite.mockCache.On("Get", blogID).Return(cacheBlogJson, nil)

	result, err := suite.blogUseCase.GetBlog(blogID, "user123")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), blog.Title, result.Title)

	suite.mockCache.AssertCalled(suite.T(), "Get", blogID)

	suite.mockCache.On("Get", blogID).Return("", nil)
	suite.mockBlogRepo.On("FindById", blogID).Return(blog, nil)
	suite.mockCache.On("Set", blogID, mock.Anything, mock.Anything).Return(nil)

	result, err = suite.blogUseCase.GetBlog(blogID, "user123")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), blog.Title, result.Title)

	// suite.mockBlogRepo.AssertExpectations(suite.T())
	// suite.mockCache.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestUpdateBlog() {
	blogID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID().Hex()
	authorID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		suite.T().Errorf("Error converting userID to ObjectID: %v", err)
		return
	}
	blog := &domain.Blog{
		ID:       primitive.NewObjectID(),
		Title:    "Updated Title",
		Content:  "Updated Content",
		AuthorID: authorID,
	}

	suite.mockBlogRepo.On("FindById", blogID).Return(blog, nil)
	suite.mockBlogRepo.On("Update", blogID, blog).Return(blog, nil)
	suite.mockCache.On("Delete", blogID).Return(nil)
	suite.mockCache.On("Delete", "all_blogs").Return(nil)

	err = suite.blogUseCase.UpdateBlog(blogID, blog, userID)

	assert.NoError(suite.T(), err)
	suite.mockBlogRepo.AssertExpectations(suite.T())
	suite.mockCache.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestDeleteBlog() {
	blogID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID().Hex()
	primUserID, _ := primitive.ObjectIDFromHex(userID)
	blog := &domain.Blog{
		ID:       primitive.NewObjectID(),
		AuthorID: primUserID,
	}

	suite.mockBlogRepo.On("FindById", blogID).Return(blog, nil)
	suite.mockBlogRepo.On("Delete", blogID).Return(nil)
	suite.mockCache.On("Delete", blogID).Return(nil)
	suite.mockCache.On("Delete", "all_blogs").Return(nil)

	err := suite.blogUseCase.DeleteBlog(blogID, userID)

	assert.NoError(suite.T(), err)
	suite.mockBlogRepo.AssertExpectations(suite.T())
	suite.mockCache.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestLikeBlog() {
	blogID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID().Hex()

	suite.mockBlogRepo.On("Like", blogID, userID).Return(nil)
	suite.mockCache.On("Delete", blogID).Return(nil)
	suite.mockCache.On("Increment", "blog:like_count:"+blogID).Return(nil)

	err := suite.blogUseCase.LikeBlog(userID, blogID)

	assert.NoError(suite.T(), err)
	suite.mockBlogRepo.AssertExpectations(suite.T())
	suite.mockCache.AssertExpectations(suite.T())
}

func TestBlogUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(BlogUseCaseTestSuite))
}
