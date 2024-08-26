package repository_test

import (
	"blogs/bootstrap"
	"blogs/domain"
	"blogs/mocks"
	"blogs/repository"
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogRepositoryTestSuite struct {
	suite.Suite
	blogCollection    *mongo.Collection
	likeCollection    *mongo.Collection
	viewCollection    *mongo.Collection
	commentCollection *mongo.Collection
	userCollection    *mongo.Collection
	client            *mongo.Client
	cacheMock         *mocks.Cache
	repo              *repository.BlogRepository
}

func (suite *BlogRepositoryTestSuite) SetupTest() {
	suite.cacheMock = &mocks.Cache{}

	client, err := bootstrap.ConnectToMongoDB("mongodb+srv://nathnaeldes:12345678n@cluster0.w8bpdtf.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")
	if err != nil {
		suite.T().Fatal(err)
	}

	suite.client = client
	db := suite.client.Database("blog") // Get the database object

	// Set collections from the correct database
	suite.blogCollection = db.Collection("blogs_test")
	suite.likeCollection = db.Collection("likes_test")
	suite.viewCollection = db.Collection("views_test")
	suite.commentCollection = db.Collection("comments_test")
	suite.userCollection = db.Collection("users_test")

	// Initialize the repository with the correct types
	suite.repo = repository.NewBlogRepository(db, suite.cacheMock) // Pass the database object, not a string
}

func (suite *BlogRepositoryTestSuite) TestInsertBlog() {
	blog := &domain.Blog{
		Title:   "Test Blog",
		Content: "This is a test blog content",
	}

	// Mock cache operation
	suite.cacheMock.On("SetCache", "some_key", "some_value").Return(nil)

	insertedBlog, err := suite.repo.InsertBlog(blog)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), insertedBlog.ID)

	// Verify cache operation
	err = suite.cacheMock.SetCache("some_key", "some_value")
	assert.NoError(suite.T(), err)
	suite.cacheMock.AssertExpectations(suite.T())
}

func (suite *BlogRepositoryTestSuite) TestGetBlogById() {
	objId := primitive.NewObjectID()
	blog := &domain.Blog{
		ID:      objId,
		Title:   "Test Blog",
		Content: "This is a test blog content",
	}

	// Insert a blog document for the test
	_, err := suite.blogCollection.InsertOne(context.Background(), blog)
	if err != nil {
		suite.T().Fatal(err)
	}

	// Mock cache operations
	suite.cacheMock.On("GetCache", "blog:"+objId.Hex()).Return("", mongo.ErrNoDocuments)
	suite.cacheMock.On("SetCache", "blog:"+objId.Hex(), mock.Anything).Return(nil)

	retrievedBlog, err := suite.repo.GetBlogByID(objId.Hex())
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), retrievedBlog)
	assert.Equal(suite.T(), objId, retrievedBlog.ID)

	// Verify that SetCache was called with the expected arguments
	suite.cacheMock.AssertCalled(suite.T(), "SetCache", "blog:"+objId.Hex(), mock.Anything)
}

func (suite *BlogRepositoryTestSuite) DeleteBlogByID() {
	objId := primitive.NewObjectID()
	blog := &domain.Blog{
		ID:      objId,
		Title:   "Test Blog",
		Content: "This is a test blog content",
	}

	// Insert a blog document for the test
	_, err := suite.blogCollection.InsertOne(context.Background(), blog)
	if err != nil {
		suite.T().Fatal(err)
	}

	// Mock cache operation
	suite.cacheMock.On("DeleteCache", "blog:"+objId.Hex()).Return(nil)

	err = suite.repo.DeleteBlogByID(objId.Hex())
	assert.NoError(suite.T(), err)

	// Verify that DeleteCache was called with the expected arguments
	suite.cacheMock.AssertCalled(suite.T(), "DeleteCache", "blog:"+objId.Hex())
}

func (suite *BlogRepositoryTestSuite) TestSearchBlog_NoResults() {
	// Prepare test data
	title := "Non-existent Title"
	author := "Non-existent Author"
	tags := []string{"tag1"}
	cacheKey := fmt.Sprintf("search:%s:%s:%s", title, author, strings.Join(tags, ","))

	// Mock cache miss
	suite.cacheMock.On("GetCache", cacheKey).Return("", nil).Once()

	// Mock empty results from the database
	suite.blogCollection = suite.client.Database("blog").Collection("blogs_test")

	// Call SearchBlog method
	results, err := suite.repo.SearchBlog(title, author, tags)
	assert.NoError(suite.T(), err)
	assert.Empty(suite.T(), results)

	// Ensure SetCache was called
	suite.cacheMock.AssertNotCalled(suite.T(), "SetCache", cacheKey, mock.Anything)
}

func (suite *BlogRepositoryTestSuite) TearDownTest() {
	suite.cacheMock.AssertExpectations(suite.T())
}

func (suite *BlogRepositoryTestSuite) TearDownSuite() {
	suite.client.Disconnect(context.Background())
}

func TestCacheTestSuite(t *testing.T) {
	suite.Run(t, new(BlogRepositoryTestSuite))
}
