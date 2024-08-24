package tests

import (
	"blog_api/domain"
	"blog_api/mocks"
	"blog_api/usecase"
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type BlogUseCaseTestSuite struct {
	suite.Suite
	blogUseCase      *usecase.BlogUseCase
	mockBlogRepo     *mocks.BlogRepositoryInterface
	mockAIService    *mocks.AIServicesInterface
	mockCacheRepo    *mocks.CacheRepositoryInterface
	mockENV          domain.EnvironmentVariables
}

func (suite *BlogUseCaseTestSuite) SetupTest() {
	suite.mockBlogRepo = new(mocks.BlogRepositoryInterface)
	suite.mockAIService = new(mocks.AIServicesInterface)
	suite.mockCacheRepo = new(mocks.CacheRepositoryInterface)
	suite.mockENV = domain.EnvironmentVariables{
		CACHE_EXPIRATION: 60,
	}

	suite.blogUseCase = usecase.NewBlogUseCase(
		suite.mockBlogRepo,
		time.Second*5,
		suite.mockAIService,
		suite.mockCacheRepo,
		suite.mockENV,
	)
}

func (suite *BlogUseCaseTestSuite) TestCreateBlogPost_Success() {
	newBlog := &domain.NewBlog{
		Title:   "Test Title",
		Content: "This is the content of the blog",
		Tags:    []string{"tag1", "tag2"},
	}

	suite.mockBlogRepo.On("InsertBlogPost", mock.Anything, mock.Anything).Return(nil)

	err := suite.blogUseCase.CreateBlogPost(context.TODO(), newBlog, "testUser")

	assert.Nil(suite.T(), err)
	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestCreateBlogPost_Failure() {
	newBlog := &domain.NewBlog{
		Title:   "Test Title",
		Content: "This is the content of the blog",
		Tags:    []string{"tag1", "tag2"},
	}

	suite.mockBlogRepo.On("InsertBlogPost", mock.Anything, mock.Anything).Return(domain.NewError("error inserting blog post", domain.ERR_INTERNAL_SERVER))

	err := suite.blogUseCase.CreateBlogPost(context.TODO(), newBlog, "testUser")

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "error inserting blog post", err.Error())
	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestDeleteBlogPost_Success() {
	blog := &domain.Blog{
		Username: "testUser",
	}

	suite.mockBlogRepo.On("FetchBlogPostByID", mock.Anything, "testBlogID", false).Return(blog, nil)
	suite.mockBlogRepo.On("DeleteBlogPost", mock.Anything, "testBlogID").Return(nil)

	err := suite.blogUseCase.DeleteBlogPost(context.TODO(), "testBlogID", "testUser")

	assert.Nil(suite.T(), err)
	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestDeleteBlogPost_Unauthorized() {
	blog := &domain.Blog{
		Username: "anotherUser",
	}

	suite.mockBlogRepo.On("FetchBlogPostByID", mock.Anything, "testBlogID", false).Return(blog, nil)

	err := suite.blogUseCase.DeleteBlogPost(context.TODO(), "testBlogID", "testUser")

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), domain.ERR_FORBIDDEN, err.Error())
	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestEditBlogPost_Success() {
	blog := &domain.Blog{
		Username: "testUser",
	}
	updatedBlog := &domain.NewBlog{
		Title:   "Updated Title",
		Content: "Updated Content",
	}

	suite.mockBlogRepo.On("FetchBlogPostByID", mock.Anything, "testBlogID", false).Return(blog, nil)
	suite.mockBlogRepo.On("UpdateBlogPost", mock.Anything, "testBlogID", updatedBlog).Return(nil)

	err := suite.blogUseCase.EditBlogPost(context.TODO(), "testBlogID", updatedBlog, "testUser")

	assert.Nil(suite.T(), err)
	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestEditBlogPost_Unauthorized() {
	blog := &domain.Blog{
		Username: "anotherUser",
	}

	updatedBlog := &domain.NewBlog{
		Title:   "Updated Title",
		Content: "Updated Content",
	}

	suite.mockBlogRepo.On("FetchBlogPostByID", mock.Anything, "testBlogID", false).Return(blog, nil)

	err := suite.blogUseCase.EditBlogPost(context.TODO(), "testBlogID", updatedBlog, "testUser")

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), domain.ERR_FORBIDDEN, err.Error())
	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestGetBlogPosts() {
	// Set up the test inputs
	filters := domain.BlogFilterOptions{
        Title:         "Test Title",
        Tags:          []string{"tag1", "tag2"},
        DateFrom:      time.Date(2024, time.July, 22, 15, 12, 15, 419908000, time.Local),
        DateTo:        time.Date(2024, time.August, 22, 15, 12, 15, 421436200, time.Local),
        SortBy:        "created_at",
        SortDirection: "desc",
        Page:          1,
        PostsPerPage:  10,
    }

	blogs := []domain.Blog{
		{
			ID:    "1",
			Title: "First Blog",
		},
		{
			ID:    "2",
			Title: "Second Blog",
		},
	}

	// Marshal the blogs into JSON for caching
	cachedBlogs, _ := json.Marshal(blogs)

	// // Mock the cache miss and database fetch
	suite.mockCacheRepo.On("IsCached", mock.Anything).Return(false).Once() // Simulate cache miss
	suite.mockBlogRepo.On("FetchBlogPosts", mock.Anything, filters).Return(blogs, len(blogs), nil).Once()
	suite.mockCacheRepo.On("CacheData", mock.Anything, string(cachedBlogs), time.Hour).Return(nil).Once()

	// Call the method under test
	resultBlogs, total, err := suite.blogUseCase.GetBlogPosts(context.Background(), filters)

	// Assertions
	suite.NoError(err)
	suite.Equal(blogs, resultBlogs)
	suite.Equal(len(blogs), total)

	// Ensure all the expectations were met
	suite.mockBlogRepo.AssertExpectations(suite.T())
	suite.mockCacheRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestGetBlogPostByID() {

	blogID := "12345"
	blog := &domain.Blog{
		ID:    blogID,
		Title: "Sample Blog",
	}

	cachedBlog, _ := json.Marshal(blog)

	suite.mockCacheRepo.On("IsCached", "blog:"+blogID).Return(false).Once() // Simulate cache miss
	suite.mockBlogRepo.On("FetchBlogPostByID", mock.Anything, blogID, true).Return(blog, nil).Once()
	suite.mockCacheRepo.On("CacheData", "blog:"+blogID, string(cachedBlog), time.Hour).Return(nil).Once()

	resultBlog, err := suite.blogUseCase.GetBlogPostByID(context.Background(), blogID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), blog, resultBlog)

	// Ensure all the expectations were met
	suite.mockCacheRepo.AssertExpectations(suite.T())
	suite.mockCacheRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestTrackBlogPopularity() {
    blogID := "12345"
    action := "like"
    state := true
    username := "user123"

    suite.mockBlogRepo.On("TrackBlogPopularity", mock.Anything, blogID, action, state, username).Return(nil).Once()

    err := suite.blogUseCase.TrackBlogPopularity(context.Background(), blogID, action, state, username)

    suite.NoError(err)

    // Ensure all the expectations were met
    suite.mockBlogRepo.AssertExpectations(suite.T())
}
func (suite *BlogUseCaseTestSuite) TestGetBlogPosts_CacheFail() {
    filters := domain.BlogFilterOptions{
        Title:         "Test Title",
        Tags:          []string{"tag1", "tag2"},
        DateFrom:      time.Date(2024, time.July, 22, 15, 29, 42, 744490500, time.Local),
        DateTo:        time.Date(2024, time.August, 22, 15, 29, 42, 746932300, time.Local),
        SortBy:        "created_at",
        SortDirection: "desc",
        Page:          1,
        PostsPerPage:  10,
    }

    blogs := []domain.Blog{
        {
            ID:    "1",
            Title: "First Blog",
        },
        {
            ID:    "2",
            Title: "Second Blog",
        },
    }

    cachedBlogs, _ := json.Marshal(blogs)

    // Simulate a cache miss
    suite.mockCacheRepo.On("IsCached", mock.Anything).Return(false).Once()
    // Simulate fetching data from the repository
    suite.mockBlogRepo.On("FetchBlogPosts", mock.Anything, filters).Return(blogs, len(blogs), nil).Once()
    // Simulate a failure in caching the data
    suite.mockCacheRepo.On("CacheData", mock.Anything, string(cachedBlogs), time.Hour).Return(domain.NewError("cache failure", domain.ERR_INTERNAL_SERVER)).Once()

    resultBlogs, total, err := suite.blogUseCase.GetBlogPosts(context.Background(), filters)

    assert.NoError(suite.T(), err)
    assert.Equal(suite.T(), blogs, resultBlogs)
    assert.Equal(suite.T(), len(blogs), total)

    // The test should pass as the main functionality isn't affected by caching failure.
    suite.mockBlogRepo.AssertExpectations(suite.T())
    suite.mockCacheRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestGetBlogPosts_FetchFail() {
    filters := domain.BlogFilterOptions{
        Title:         "Test Title",
        Tags:          []string{"tag1", "tag2"},
        DateFrom:      time.Date(2024, time.July, 22, 15, 42, 48, 7799700, time.Local),
        DateTo:        time.Date(2024, time.August, 22, 15, 42, 48, 9354300, time.Local),
        SortBy:        "created_at",
        SortDirection: "desc",
        Page:          1,
        PostsPerPage:  10,
    }

    // Simulate a cache miss
    suite.mockCacheRepo.On("IsCached", mock.Anything).Return(false).Once()
    // Simulate fetching data from the repository and failing
    suite.mockBlogRepo.On("FetchBlogPosts", mock.Anything, filters).Return(nil, 0,  domain.NewError("cache failure", domain.ERR_INTERNAL_SERVER)).Once()

    resultBlogs, total, err := suite.blogUseCase.GetBlogPosts(context.Background(), filters)

    assert.Error(suite.T(), err)
    assert.Nil(suite.T(), resultBlogs)
    assert.Equal(suite.T(), 0, total)

    suite.mockBlogRepo.AssertExpectations(suite.T())
    suite.mockCacheRepo.AssertExpectations(suite.T())
}


func (suite *BlogUseCaseTestSuite) TestGetBlogPostByID_CacheRetrieveFail() {

    blogID := "12345"
    blog := &domain.Blog{
        ID:    blogID,
        Title: "Sample Blog",
    }

    // Simulate cache hit but failure in retrieving data
    suite.mockCacheRepo.On("IsCached", "blog:"+blogID).Return(true).Once()
    suite.mockCacheRepo.On("GetCacheData", "blog:"+blogID).Return("", domain.NewError("cache failure", domain.ERR_INTERNAL_SERVER)).Once()
    // Simulate fetching data from the repository
    suite.mockBlogRepo.On("FetchBlogPostByID", mock.Anything, blogID, true).Return(blog, nil).Once()
	suite.mockCacheRepo.On("CacheData", "blog:"+blogID, mock.Anything, mock.Anything).Return(nil).Once()

    resultBlog, err := suite.blogUseCase.GetBlogPostByID(context.Background(), blogID)

    assert.NoError(suite.T(), err)
    assert.Equal(suite.T(), blog, resultBlog)

    suite.mockBlogRepo.AssertExpectations(suite.T())
    suite.mockCacheRepo.AssertExpectations(suite.T())
}


func (suite *BlogUseCaseTestSuite) TestGetBlogPostByID_FetchFail() {

    blogID := "12345"

    // Simulate a cache miss
    suite.mockCacheRepo.On("IsCached", "blog:"+blogID).Return(false).Once()
    // Simulate fetching data from the repository fails
    suite.mockBlogRepo.On("FetchBlogPostByID", mock.Anything, blogID, true).Return(nil, domain.NewError("cache failure", domain.ERR_INTERNAL_SERVER)).Once()

    resultBlog, err := suite.blogUseCase.GetBlogPostByID(context.Background(), blogID)

    assert.Error(suite.T(), err)
    assert.Nil(suite.T(), resultBlog)

    suite.mockBlogRepo.AssertExpectations(suite.T())
    suite.mockCacheRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestTrackBlogPopularity_Fail() {

    blogID := "12345"
    action := "like"
    state := true
    username := "user123"

    // Simulate failure in tracking blog popularity
    suite.mockBlogRepo.On("TrackBlogPopularity", mock.Anything, blogID, action, state, username).Return(domain.NewError("failed to track popularity", domain.ERR_INTERNAL_SERVER)).Once()

    err := suite.blogUseCase.TrackBlogPopularity(context.Background(), blogID, action, state, username)

    assert.Error(suite.T(), err)
    suite.mockBlogRepo.AssertExpectations(suite.T())
}



func (suite *BlogUseCaseTestSuite) TestGenerateBlogContent() {
    topics := []string{"AI", "Technology"}
    expectedContent := "Generated content for AI and Technology"

    suite.mockAIService.On("GenerateContent", topics).Return(expectedContent, nil)

    content, err := suite.blogUseCase.GenerateBlogContent(topics)

    suite.NoError(err)
    suite.Equal(expectedContent, content)

    suite.mockAIService.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestReviewBlogContent() {
    blogContent := "This is a blog post"
    expectedSuggestions := "Suggestions for improvement"

    suite.mockAIService.On("ReviewContent", blogContent).Return(expectedSuggestions, nil)

    suggestions, err := suite.blogUseCase.ReviewBlogContent(blogContent)

    suite.NoError(err)
    suite.Equal(expectedSuggestions, suggestions)

    suite.mockAIService.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestGenerateTrendingTopics() {
    keywords := []string{"AI", "Tech"}
    expectedTopics := []string{"AI in 2024", "Future of Tech"}

    suite.mockAIService.On("GenerateTrendingTopics", keywords).Return(expectedTopics, nil)

    topics, err := suite.blogUseCase.GenerateTrendingTopics(keywords)

    suite.NoError(err)
    suite.Equal(expectedTopics, topics)

    suite.mockAIService.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestAddComment() {
    ctx := context.Background()
    blogID := "blog123"
    newComment := &domain.NewComment{Content: "Great post!"}
    userName := "user123"

    suite.mockBlogRepo.On("CreateComment", mock.Anything, mock.Anything, blogID, userName).Return(nil)

    err := suite.blogUseCase.AddComment(ctx, blogID, newComment, userName)

    suite.NoError(err)
    suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestDeleteComment() {
    ctx := context.Background()
    blogID := "blog123"
    commentID := "comment123"
    userName := "user123"

    suite.mockBlogRepo.On("DeleteComment", mock.Anything, commentID, blogID, userName).Return(nil)

    err := suite.blogUseCase.DeleteComment(ctx, blogID, commentID, userName)

    suite.NoError(err)
    suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestUpdateComment() {
    ctx := context.Background()
    blogID := "blog123"
    commentID := "comment123"
    updatedComment := &domain.NewComment{Content: "Updated comment"}
    userName := "user123"

    suite.mockBlogRepo.On("UpdateComment", mock.Anything, updatedComment, commentID, blogID, userName).Return(nil)

    err := suite.blogUseCase.UpdateComment(ctx, blogID, commentID, updatedComment, userName)

    suite.NoError(err)
    suite.mockBlogRepo.AssertExpectations(suite.T())
}

func TestBlogUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(BlogUseCaseTestSuite))
}
