package tests

import (
	"testing"
	"time"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/mocks"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type blogUseCaseTestSuite struct{
	suite.Suite
	blogRepo *mocks.MockBlogRepository
	blogCache *mocks.MockCacheService
	blogUseCase domain.BlogUseCase
}

func (suite *blogUseCaseTestSuite) SetupSuite(){
	suite.blogRepo = new(mocks.MockBlogRepository)
	suite.blogCache = new(mocks.MockCacheService)
	suite.blogUseCase = usecases.NewBlogUseCase(suite.blogRepo, suite.blogCache)
}


func (suite *blogUseCaseTestSuite) TestCreateBlogUsecases(t *testing.T) {
	mockRepo := new(mocks.MockBlogRepository)
	mockCache := new(mocks.MockCacheService)
	useCase := usecases.NewBlogUseCase(mockRepo, mockCache)

	blog := &domain.Blog{
		Title:   "Test Blog",
		Content: "Test Content",
	}
	authorID := "authorID"

	mockRepo.On("Create", blog).Return(blog, nil)
	mockCache.On("Delete", "all_blogs").Return(nil)

	err := useCase.CreateBlog(blog, authorID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
	mockCache.AssertExpectations(t)
}

func (suite *blogUseCaseTestSuite) TestGetBlog(t *testing.T) {
	mockRepo := new(mocks.MockBlogRepository)
	mockCache := new(mocks.MockCacheService)
	useCase := usecases.NewBlogUseCase(mockRepo, mockCache)
	blogID := "blogID"
	userID := "userID"
	primitiveBlogID , err := primitive.ObjectIDFromHex(blogID)
	if err != nil {}
	blog := &domain.Blog{ID: primitiveBlogID}
	expectedDuration := time.Duration(0)
	mockCache.On("Get", blogID).Return("", nil)
	mockRepo.On("FindById", blogID).Return(blog, nil)
	mockCache.On("Set", blogID, mock.Anything, expectedDuration).Return(nil)

	result, err := useCase.GetBlog(blogID, userID)

	assert.NoError(t, err)
	assert.Equal(t, blog, result)
	mockCache.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func (suite *blogUseCaseTestSuite) TestUpdateBlogUsecase(t *testing.T) {
	mockRepo := new(mocks.MockBlogRepository)
	mockCache := new(mocks.MockCacheService)
	useCase := usecases.NewBlogUseCase(mockRepo, mockCache)

	blogID := "blogID"
	blog := &domain.Blog{
		Title: "Updated Title",
	}
	userID := "userID"

	mockRepo.On("Update", blogID, blog).Return(blog, nil)
	mockCache.On("Delete", blogID).Return(nil)
	mockCache.On("Delete", "all_blogs").Return(nil)

	err := useCase.UpdateBlog(blogID, blog, userID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
	mockCache.AssertExpectations(t)
}

func (suite *blogUseCaseTestSuite) TestDeleteBlogUsecase(t *testing.T) {
	mockRepo := new(mocks.MockBlogRepository)
	mockCache := new(mocks.MockCacheService)
	useCase := usecases.NewBlogUseCase(mockRepo, mockCache)

	blogID := "blogID"

	mockRepo.On("Delete", blogID).Return(nil)
	mockCache.On("Delete", blogID).Return(nil)
	mockCache.On("Delete", "all_blogs").Return(nil)

	err := useCase.DeleteBlog(blogID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
	mockCache.AssertExpectations(t)
}

func (suite *blogUseCaseTestSuite) TestSearchBlogsByTitle(t *testing.T) {
	mockRepo := new(mocks.MockBlogRepository)
	mockCache := new(mocks.MockCacheService)
	useCase := usecases.NewBlogUseCase(mockRepo, mockCache)

	title := "Test Title"
	pageNumber := "1"

	mockRepo.On("SearchByTitle", title, pageNumber).Return([]domain.Blog{}, nil)

	result, err := useCase.SearchBlogsByTitle(title, pageNumber)

	assert.NoError(t, err)
	assert.Equal(t, []domain.Blog{}, result)
	mockRepo.AssertExpectations(t)
}

func (suite *blogUseCaseTestSuite) TestSearchBlogsByAuthor(t *testing.T) {
	mockRepo := new(mocks.MockBlogRepository)
	mockCache := new(mocks.MockCacheService)
	useCase := usecases.NewBlogUseCase(mockRepo, mockCache)

	author := "Test Author"
	pageNumber := "1"

	mockRepo.On("SearchByAuthor", author, pageNumber).Return([]domain.Blog{}, nil)

	result, err := useCase.SearchBlogsByAuthor(author, pageNumber)

	assert.NoError(t, err)
	assert.Equal(t, []domain.Blog{}, result)
	mockRepo.AssertExpectations(t)
}

func (suite *blogUseCaseTestSuite) TestGetBlogs(t *testing.T) {
	mockRepo := new(mocks.MockBlogRepository)
	mockCache := new(mocks.MockCacheService)
	useCase := usecases.NewBlogUseCase(mockRepo, mockCache)

	pageNumber := "1"
	expectedDuration := time.Duration(0)
	mockCache.On("Get", "all_blogs").Return("", nil)
	mockRepo.On("FindAll", pageNumber).Return([]domain.Blog{}, nil)
	mockCache.On("Set", "all_blogs", mock.Anything, expectedDuration).Return(nil)

	result, err := useCase.GetBlogs(pageNumber)

	assert.NoError(t, err)
	assert.Equal(t, []domain.Blog{}, result)
	mockRepo.AssertExpectations(t)
	mockCache.AssertExpectations(t)
}

func (suite *blogUseCaseTestSuite) TestGetBlogsFromCache(t *testing.T) {
	mockRepo := new(mocks.MockBlogRepository)
	mockCache := new(mocks.MockCacheService)
	useCase := usecases.NewBlogUseCase(mockRepo, mockCache)

	pageNumber := "1"
	cachedBlogs := `[{"title":"Test Blog","content":"Test Content"}]`

	mockCache.On("Get", "all_blogs").Return(cachedBlogs, nil)

	result, err := useCase.GetBlogs(pageNumber)

	assert.NoError(t, err)
	assert.Equal(t, []domain.Blog{{Title: "Test Blog", Content: "Test Content"}}, result)
	mockRepo.AssertExpectations(t)
	mockCache.AssertExpectations(t)
}

func (suite *blogUseCaseTestSuite) TestGetBlogsFromCacheError(t *testing.T) {
	mockRepo := new(mocks.MockBlogRepository)
	mockCache := new(mocks.MockCacheService)
	useCase := usecases.NewBlogUseCase(mockRepo, mockCache)

	pageNumber := "1"
	expectedDuration := time.Duration(0)
	mockCache.On("Get", "all_blogs").Return("", nil)
	mockRepo.On("FindAll", pageNumber).Return([]domain.Blog{}, nil)
	mockCache.On("Set", "all_blogs", mock.Anything, expectedDuration).Return(nil)

	result, err := useCase.GetBlogs(pageNumber)

	assert.NoError(t, err)
	assert.Equal(t, []domain.Blog{}, result)
	mockRepo.AssertExpectations(t)
	mockCache.AssertExpectations(t)
}

func TestBlogUseCaseTestSuite(t *testing.T){
	suite.Run(t , new(blogUseCaseTestSuite))
}
