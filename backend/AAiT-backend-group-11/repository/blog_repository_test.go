package repository_test

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/mongo/mocks"
	"backend-starter-project/repository"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestCreateBlogPost(t *testing.T)  {
	// Initialize the mock collection and context
	mockCollection := new(mocks.Collection)
	mockContext := context.Background()

	// Create the blogRepository instance with the mock collection
	repo := repository.NewBlogRepository(mockCollection, mockContext)

	// Define the user ID and create the ObjectID from it
	userID := "60d5f3c4a6e6b5b0d4a3d66f"
	userObjectID, _ := primitive.ObjectIDFromHex(userID)

	blogPost := &entities.BlogPost{
		Title:    "Test Blog",
		Content:  "This is a test blog post.",
		AuthorID: userObjectID,
	}

	// Setup the expected behavior of the mock
	mockCollection.On("InsertOne", mockContext, blogPost).Return(&mongo.InsertOneResult{}, nil)

	// Call the CreateBlogPost method
	result, err := repo.CreateBlogPost(blogPost, userID)

	// Assert that the result is not nil and that there is no error
	assert.NoError(t, err)
	assert.NotNil(t, result)
	
	// Assert that the mock expectations were met
	mockCollection.AssertExpectations(t)
}

func TestGetBlogPostById(t *testing.T)  {
	// Initialize the mock collection and context
	mockCollection := new(mocks.Collection)
	mockContext := context.Background()

	// Create the blogRepository instance with the mock collection
	repo := repository.NewBlogRepository(mockCollection, mockContext)

	// Define the blog post ID and create the ObjectID from it
	blogPostID := "60d5f3c4a6e6b5b0d4a3d66f"
	blogPostObjectID, _ := primitive.ObjectIDFromHex(blogPostID)

	expectedBlog := &entities.BlogPost{
		ID:       blogPostObjectID,
		Title:    "Test Blog",
		Content:  "This is a test blog post.",
		AuthorID: blogPostObjectID,
	}

	mockSingleResult := new(mocks.SingleResult)
	// Setup the expected behavior of the mock
	mockCollection.On("FindOne", mockContext, bson.M{"_id": blogPostObjectID}).Return(mockSingleResult, nil)
	mockSingleResult.On("Decode", mock.AnythingOfType("*entities.BlogPost")).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*entities.BlogPost)
		*arg = *expectedBlog
	}).Return(nil)

	// Call the GetBlogPostById method
	result, err := repo.GetBlogPostById(blogPostID)

	// Assert that the result is not nil and that there is no error
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Assert that the mock expectations were met
	mockCollection.AssertExpectations(t)
}

func TestUpdateBlogPost(t *testing.T)  {
	mockCollection := new(mocks.Collection)
	mockContext := context.Background()

	repo := repository.NewBlogRepository(mockCollection, mockContext)

	blogPostID := primitive.NewObjectID()

	blogPost := &entities.BlogPost{
		ID:      blogPostID,
		Title:   "Updated Title",
		Content: "Updated Content",
		Tags:    []string{"go", "mongodb"},
	}

	filter := bson.M{"_id": blogPostID}

	update := bson.M{
		"$set": bson.M{
			"title":    blogPost.Title,
			"content":  blogPost.Content,
			"tags":     blogPost.Tags,
		},
	}

	mockCollection.On("UpdateOne", mockContext, filter, update).Return(&mongo.UpdateResult{
		ModifiedCount: 1,
	}, nil)

	// Call the UpdateBlogPost method
	result, err := repo.UpdateBlogPost(blogPost)

	assert.NoError(t, err)

	assert.Equal(t, blogPost.Title, result.Title)
	assert.Equal(t, blogPost.Content, result.Content)
	assert.Equal(t, blogPost.Tags, result.Tags)
	assert.WithinDuration(t, time.Now(), result.UpdatedAt, time.Second)

	// Ensure the UpdateOne method was called as expected
	mockCollection.AssertExpectations(t)
}

func TestDeleteBlogPost(t *testing.T)  {
	mockCollection := new(mocks.Collection)
	mockContext := context.Background()

	repo := repository.NewBlogRepository(mockCollection, mockContext)

	blogPostID := primitive.NewObjectID()

	filter := bson.M{"_id": blogPostID}

	mockDeleteCount := int64(1)
	mockCollection.On("DeleteOne", mockContext, filter).Return(mockDeleteCount, nil)

	// Call the DeleteBlogPost method
	err := repo.DeleteBlogPost(blogPostID.Hex())

	assert.NoError(t, err)

	// Ensure the DeleteOne method was called as expected
	mockCollection.AssertExpectations(t)
}

func TestGetBlogPosts(t *testing.T)  {
	mockCollection := new(mocks.Collection)
	mockContext := context.Background()

	repo := repository.NewBlogRepository(mockCollection, mockContext)

	expectedBlogPosts := []entities.BlogPost{
		{
			ID:      primitive.NewObjectID(),
			Title:   "Test Blog 1",
			Content: "This is a test blog post.",
			Tags:    []string{"go", "mongodb"},
		},
		{
			ID:      primitive.NewObjectID(),
			Title:   "Test Blog 2",
			Content: "This is another test blog post.",
			Tags:    []string{"go", "mongodb"},
		},
	}

	mockCursor := new(mocks.Cursor)
	mockCursor.On("Next", mockContext).Return(true).Once()
	mockCursor.On("Decode", mock.AnythingOfType("*entities.BlogPost")).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*entities.BlogPost)
		*arg = expectedBlogPosts[0]
	}).Return(nil).Once()

	mockCursor.On("Next", mockContext).Return(true).Once()
	mockCursor.On("Decode", mock.AnythingOfType("*entities.BlogPost")).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*entities.BlogPost)
		*arg = expectedBlogPosts[1]
	}).Return(nil).Once()

	mockCursor.On("Next", mockContext).Return(false).Once()
	mockCursor.On("Err").Return(nil)
	mockCursor.On("Close", mockContext).Return(nil)

	// Setup the expected behavior of the mock collection's Find method
	mockCollection.On("Find", mockContext, bson.M{}, mock.Anything).Return(mockCursor, nil)

	// Call the GetBlogPosts method with sample pagination and sorting parameters
	page, pageSize, sortBy := 1, 10, "likes"
	actualBlogPosts, err := repo.GetBlogPosts(page, pageSize, sortBy)

	// Assert that no error occurred
	assert.NoError(t, err)

	// Assert that the returned blogPosts match the expected ones
	assert.Equal(t, expectedBlogPosts, actualBlogPosts)

	// Ensure the Find method was called as expected
	mockCollection.AssertExpectations(t)
	mockCursor.AssertExpectations(t)
}

func TestCountBlogPosts(t *testing.T){
	mockCollection := new(mocks.Collection)
	mockContext := context.Background()

	repo := repository.NewBlogRepository(mockCollection, mockContext)

	mockCollection.On("CountDocuments", mockContext, bson.M{}).Return(int64(2), nil)

	count, err := repo.CountBlogPosts()

	assert.NoError(t, err)
	assert.Equal(t, int64(2), int64(count))

	mockCollection.AssertExpectations(t)
}

func TestChangeCommentCount(t *testing.T){
	mockCollection := new(mocks.Collection)
	mockContext := context.Background()

	repo := repository.NewBlogRepository(mockCollection, mockContext)

	blogPostID := primitive.NewObjectID()

	filter := bson.M{"_id": blogPostID}

	update := bson.M{
		"$inc": bson.M{
			"commentCount": 1,
		},
	}

	mockCollection.On("FindOneAndUpdate", mockContext, filter, update).Return(&mongo.SingleResult{}, nil)

	// Call the ChangeCommentCount method
	err := repo.ChangeCommentCount(blogPostID.Hex(), 1)

	assert.NoError(t, err)

	// Ensure the UpdateOne method was called as expected
	mockCollection.AssertExpectations(t)
}