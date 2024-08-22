package tests

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	mocks "github.com/aait.backend.g5.main/backend/Mocks"
	repository "github.com/aait.backend.g5.main/backend/Repository"
)


type BlogRepositorySuite struct {
	suite.Suite

	db                   *mocks.Database
	blogCollection       *mocks.Collection
	blogActionCollection *mocks.Collection
	repository           *repository.BlogMongoRepository
}

func (suite *BlogRepositorySuite) SetupTest() {
	suite.db = new(mocks.Database)
	suite.blogCollection = new(mocks.Collection)
	suite.blogActionCollection = new(mocks.Collection)
	suite.repository = &repository.BlogMongoRepository{
		BlogCollection:       suite.blogCollection,
		BlogActionCollection: suite.blogActionCollection,
	}
}

// Test CreateBlog
func (suite *BlogRepositorySuite) TestCreateBlog_Success() {
	blog := &models.Blog{
		Title:    "Test Blog",
		Content:  "Test Content",
		AuthorID: "123",
		Tags:     []string{"test", "go"},
	}

	suite.blogCollection.On("InsertOne", mock.Anything, blog).Return(nil, nil)

	result, err := suite.repository.CreateBlog(context.Background(), blog)

	suite.Nil(err)
	suite.NotNil(result)
	suite.Equal(blog.Title, result.Title)
	suite.blogCollection.AssertExpectations(suite.T())
}

func (suite *BlogRepositorySuite) TestCreateBlog_Failure() {
	blog := &models.Blog{
		Title:    "Test Blog",
		Content:  "Test Content",
		AuthorID: "123",
		Tags:     []string{"test", "go"},
	}

	suite.blogCollection.On("InsertOne", mock.Anything, blog).Return(nil, errors.New("insertion error"))

	result, err := suite.repository.CreateBlog(context.Background(), blog)

	suite.Nil(result)
	suite.NotNil(err)
	suite.Equal("Failed to create blog", err.Message)
	suite.blogCollection.AssertExpectations(suite.T())
}

// Test GetBlog
func (suite *BlogRepositorySuite) TestGetBlog_Success() {
	blog := &models.Blog{
		ID:    "123",
		Title: "Test Blog",
	}

	singleResult := &mocks.SingleResult{}
	suite.blogCollection.On("FindOne", mock.Anything, bson.M{"_id": "123"}).Return(singleResult)
	singleResult.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
		// Simulate decoding into the provided argument
		arg := args.Get(0).(*models.Blog)
		*arg = *blog
	}).Return(nil)

	result, err := suite.repository.GetBlog(context.Background(), "123")

	suite.Nil(err)
	suite.NotNil(result)
	suite.Equal(blog.Title, result.Title)
	suite.blogCollection.AssertExpectations(suite.T())
	singleResult.AssertExpectations(suite.T())
}

func (suite *BlogRepositorySuite) TestGetBlog_Failure() {
	singleResult := &mocks.SingleResult{}
	suite.blogCollection.On("FindOne", mock.Anything, bson.M{"_id": "123"}).Return(singleResult)
	singleResult.On("Decode", mock.Anything).Return(errors.New("some error"))

	result, err := suite.repository.GetBlog(context.Background(), "123")

	suite.Nil(result)
	suite.NotNil(err)
	suite.Equal("Failed to retrieve blog", err.Message)
	suite.blogCollection.AssertExpectations(suite.T())
	singleResult.AssertExpectations(suite.T())
}

// Test GetBlogs
func (suite *BlogRepositorySuite) TestGetBlogs_Success() {
	cursor := &mocks.Cursor{}
	blogs := []*models.Blog{
		{ID: "1", Title: "Blog 1"},
		{ID: "2", Title: "Blog 2"},
	}

	suite.blogCollection.On("Find", mock.Anything, bson.M{}, mock.Anything).Return(cursor, nil)
	cursor.On("Next", mock.Anything).Return(true).Times(2)
	cursor.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*models.Blog)
		if len(blogs) > 0 {
			*arg = *blogs[0]
			blogs = blogs[1:]
		}
	}).Return(nil).Times(2)
	cursor.On("Next", mock.Anything).Return(false)
	cursor.On("Err").Return(nil)
	cursor.On("Close", mock.Anything).Return(nil)

	result, err := suite.repository.GetBlogs(context.Background(), 1)

	suite.Nil(err)
	suite.NotNil(result)
	suite.Equal(2, len(result))
	suite.Equal("Blog 1", result[0].Title)
	suite.Equal("Blog 2", result[1].Title)
	suite.blogCollection.AssertExpectations(suite.T())
	cursor.AssertExpectations(suite.T())
}

func (suite *BlogRepositorySuite) TestGetBlogs_Failure() {
	suite.blogCollection.On("Find", mock.Anything, bson.M{}, mock.Anything).Return(nil, errors.New("find error"))

	result, err := suite.repository.GetBlogs(context.Background(), 1)

	suite.Nil(result)
	suite.NotNil(err)
	suite.Equal("Failed to retrieve blogs", err.Message)
	suite.blogCollection.AssertExpectations(suite.T())
}

// Test UpdateBlog
func (suite *BlogRepositorySuite) TestUpdateBlog_Success() {
	blog := &models.Blog{
		ID:      "123",
		Title:   "Updated Blog",
		Content: "Updated Content",
		// Keep other fields as needed
	}

	update := bson.M{"$set": bson.M{
		"title":      blog.Title,
		"content":    blog.Content,
		"updated_at": time.Now().Truncate(time.Minute),
	}}

	suite.blogCollection.On("UpdateOne", mock.Anything, bson.M{"_id": blog.ID}, update).Return(nil, nil)

	err := suite.repository.UpdateBlog(context.Background(), blog.ID, blog)

	suite.Nil(err)
	suite.blogCollection.AssertExpectations(suite.T())
}

func (suite *BlogRepositorySuite) TestUpdateBlog_Failure() {
	blog := &models.Blog{
		Title:   "Updated Blog",
		Content: "Updated Content",
	}

	suite.blogCollection.On("UpdateOne", mock.Anything, bson.M{"_id": blog.ID}, mock.Anything).Return(nil, errors.New("update error"))

	err := suite.repository.UpdateBlog(context.Background(), blog.ID, blog)

	suite.NotNil(err)
	suite.Equal("Failed to update blog", err.Message)
	suite.blogCollection.AssertExpectations(suite.T())
}

// Test DeleteBlog
func (suite *BlogRepositorySuite) TestDeleteBlog_Success() {
	suite.blogCollection.On("DeleteOne", mock.Anything, bson.M{"_id": "123"}).Return(int64(1), nil)

	err := suite.repository.DeleteBlog(context.Background(), "123")

	suite.Nil(err)
	suite.blogCollection.AssertExpectations(suite.T())
}

func (suite *BlogRepositorySuite) TestDeleteBlog_Failure() {
	suite.blogCollection.On("DeleteOne", mock.Anything, bson.M{"_id": "123"}).Return(int64(0), errors.New("deletion error"))

	err := suite.repository.DeleteBlog(context.Background(), "123")

	suite.NotNil(err)
	suite.Equal("Failed to delete blog", err.Message)
	suite.blogCollection.AssertExpectations(suite.T())
}

// Test IncreaseView
func (suite *BlogRepositorySuite) TestIncreaseView_Success() {
	filter := bson.M{"blog_id": "123"}
	update := bson.M{"$inc": bson.M{"view_count": 1}}
	option := options.Update().SetUpsert(true)

	suite.blogActionCollection.On("UpdateOne", mock.Anything, filter, update, option).Return(nil, nil)

	err := suite.repository.IncreaseView(context.Background(), "123")

	suite.Nil(err)
	suite.blogActionCollection.AssertExpectations(suite.T())
}

func (suite *BlogRepositorySuite) TestIncreaseView_Failure() {
	suite.blogActionCollection.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("update error"))

	err := suite.repository.IncreaseView(context.Background(), "123")

	suite.NotNil(err)
	suite.Equal("Failed to increase view count", err.Message)
	suite.blogActionCollection.AssertExpectations(suite.T())
}

// Test GetPopularity
func (suite *BlogRepositorySuite) TestGetPopularity_Success() {
	popularity := &models.Popularity{
		BlogID: "123",
	}

	// Create a mock SingleResult instance
	singleResult := new(mocks.SingleResult)

	// Set up the mock expectations
	suite.blogActionCollection.On("FindOne", mock.Anything, bson.M{"blog_id": "123"}, mock.Anything).Return(singleResult)
	singleResult.On("Decode", mock.AnythingOfType("*models.Popularity")).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*models.Popularity)
		*arg = *popularity
	}).Return(nil)

	result, err := suite.repository.GetPopularity(context.Background(), "123")

	suite.Nil(err)
	suite.NotNil(result)
	suite.Equal(popularity.BlogID, result.BlogID)
	suite.blogActionCollection.AssertExpectations(suite.T())
	singleResult.AssertExpectations(suite.T())
}

func (suite *BlogRepositorySuite) TestGetPopularity_Failure() {
	// Create a mock SingleResult instance
	singleResult := new(mocks.SingleResult)

	// Set up the mock expectations
	suite.blogActionCollection.On("FindOne", mock.Anything, bson.M{"blog_id": "123"}).Return(singleResult)
	singleResult.On("Decode", mock.AnythingOfType("*models.Popularity")).Return(errors.New("some error"))

	result, err := suite.repository.GetPopularity(context.Background(), "123")

	suite.Nil(result)
	suite.NotNil(err)
	suite.Equal("Failed to retrieve popularity information", err.Message)
	suite.blogActionCollection.AssertExpectations(suite.T())
}

// Test SearchBlogsByPopularity
func (suite *BlogRepositorySuite) TestSearchBlogsByPopularity_Success() {
	// Mocking cursor and expected popularity data
	cursor := &mocks.Cursor{}
	popularity := []*models.Popularity{
		{BlogID: "1", ViewCount: 30, LikeCount: 15, DislikeCount: 5},
		{BlogID: "2", ViewCount: 25, LikeCount: 10, DislikeCount: 5},
	}

	// Filter that matches the test logic
	filter := dtos.FilterBlogRequest{
		LikeCount:    10,
		DislikeCount: 5,
		ViewCount:    20,
	}

	// Expected BSON filter based on the filter criteria
	bfilter := bson.M{
		"dislike_count": bson.M{"$gte": filter.DislikeCount},
		"like_count":    bson.M{"$gte": filter.LikeCount},
		"view_count":    bson.M{"$gte": filter.ViewCount},
	}

	// Mocking the collection's Find call
	suite.blogActionCollection.On("Find", mock.Anything, bfilter, mock.Anything).Return(cursor, nil)

	// Setting up cursor behavior
	cursor.On("Next", mock.Anything).Return(true).Times(len(popularity)) // Return true for each document
	cursor.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*models.Popularity)
		*arg = *popularity[0]           // Assign the first element
		popularity = popularity[1:]     // Remove the assigned element for subsequent iterations
	}).Return(nil).Times(len(popularity)) // Repeat for the number of documents

	// Close the cursor at the end
	cursor.On("Next", mock.Anything).Return(false) // No more documents
	cursor.On("Err").Return(nil)                   // No error on cursor
	cursor.On("Close", mock.Anything).Return(nil)  // Close the cursor

	// Simulate an input blogs slice for testing
	blogsSlice := map[string]*models.Blog{
		"1": {ID: "1", Title: "Blog 1"},
		"2": {ID: "2", Title: "Blog 2"},
	}

	// Call the repository method
	result, err := suite.repository.SearchBlogsByPopularity(context.Background(), filter, blogsSlice)

	// Assertions
	suite.Nil(err)
	suite.NotNil(result)
	suite.Equal(2, len(result))                      // Should return 2 blogs
	suite.Equal("Blog 1", result[0].Title)           // First blog should match the title in blogsSlice
	suite.Equal("Blog 2", result[1].Title)           // Second blog should match the title in blogsSlice
	suite.blogActionCollection.AssertExpectations(suite.T()) // Ensure mocks were called as expected
}


func (suite *BlogRepositorySuite) TestSearchBlogsByPopularity_Failure() {
	filter := dtos.FilterBlogRequest{
		LikeCount:    10,
		DislikeCount: 5,
		ViewCount:    20,
	}

	suite.blogActionCollection.On("Find", mock.Anything, mock.Anything).Return(nil, errors.New("find error"))

	result, err := suite.repository.SearchBlogsByPopularity(context.Background(), filter, map[string]*models.Blog{})

	suite.Nil(result)
	suite.NotNil(err)
	suite.Equal("Failed to search blogs by popularity", err.Message)
	suite.blogActionCollection.AssertExpectations(suite.T())
}

func TestBlogRepositorySuite(t *testing.T) {
	suite.Run(t, new(BlogRepositorySuite))
}
