package repository_test

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/mongo/mocks"
	"backend-starter-project/repository"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
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