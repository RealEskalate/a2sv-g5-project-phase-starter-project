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
)

func TestAddComment(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Create a mock collection
		mockCollection := new(mocks.Collection)
		// Create a mock context
		mockCtx := context.Background()

		// Instantiate the repository with the mock collection
		repo := repository.NewCommentRepository(mockCollection, mockCtx)

		// Create a sample comment to be added
		comment := &entities.Comment{
			Content:   "This is a test comment",
			AuthorID:  primitive.NewObjectID(),
			ID:    primitive.NewObjectID(),
		}

		// Mock the InsertOne method to return a successful insertion
		mockCollection.On("InsertOne", mockCtx, comment).Return(nil, nil)

		// Call the AddComment method
		result, err := repo.AddComment(comment)

		// Assert that there was no error and the comment was returned
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, comment.Content, result.Content)
		assert.WithinDuration(t, time.Now(), result.CreatedAt, time.Second)

		// Verify that InsertOne was called with the correct arguments
		mockCollection.AssertCalled(t, "InsertOne", mockCtx, comment)
		mockCollection.AssertExpectations(t)
	})

	t.Run("InsertOne Error", func(t *testing.T) {
		// Create a mock collection
		mockCollection := new(mocks.Collection)
		// Create a mock context
		mockCtx := context.Background()

		// Instantiate the repository with the mock collection
		repo := repository.NewCommentRepository(mockCollection, mockCtx)

		// Create a sample comment to be added
		comment := &entities.Comment{
			Content:   "This is a test comment",
			AuthorID:  primitive.NewObjectID(),
			ID:    primitive.NewObjectID(),
		}

		// Mock the InsertOne method to return an error
		mockCollection.On("InsertOne", mockCtx, comment).Return(nil, assert.AnError)

		// Call the AddComment method
		result, err := repo.AddComment(comment)

		// Assert that there was an error and no comment was returned
		assert.Error(t, err)
		assert.Nil(t, result)

		// Verify that InsertOne was called with the correct arguments
		mockCollection.AssertCalled(t, "InsertOne", mockCtx, comment)
		mockCollection.AssertExpectations(t)
	})
}

func TestDeleteComment(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Create a mock collection
		mockCollection := new(mocks.Collection)
		// Create a mock context
		mockCtx := context.Background()

		// Instantiate the repository with the mock collection
		repo := repository.NewCommentRepository(mockCollection, mockCtx)

		// Valid ObjectID for the comment
		commentID := "507f1f77bcf86cd799439011"
		objID, _ := primitive.ObjectIDFromHex(commentID)

		mockDeleteCount := int64(1)
		// Mock the DeleteOne method to return a successful deletion result
		mockCollection.On("DeleteOne", mockCtx, bson.M{"_id": objID}).Return(mockDeleteCount, nil)

		// Call the DeleteComment method
		err := repo.DeleteComment(commentID)

		// Assert that there was no error
		assert.NoError(t, err)

		// Verify that DeleteOne was called with the correct arguments
		mockCollection.AssertCalled(t, "DeleteOne", mockCtx, bson.M{"_id": objID})
		mockCollection.AssertExpectations(t)
	})

	t.Run("Invalid ObjectID Error", func(t *testing.T) {
		// Create a mock collection
		mockCollection := new(mocks.Collection)
		// Create a mock context
		mockCtx := context.Background()

		// Instantiate the repository with the mock collection
		repo := repository.NewCommentRepository(mockCollection, mockCtx)

		// Invalid ObjectID for the comment
		invalidCommentID := "invalidID"

		// Call the DeleteComment method with an invalid ID
		err := repo.DeleteComment(invalidCommentID)

		// Assert that there was an error
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "the provided hex string is not a valid ObjectID")

		// Verify that DeleteOne was never called
		mockCollection.AssertNotCalled(t, "DeleteOne", mock.Anything, mock.Anything)
	})

	t.Run("DeleteOne Error", func(t *testing.T) {
		// Create a mock collection
		mockCollection := new(mocks.Collection)
		// Create a mock context
		mockCtx := context.Background()

		// Instantiate the repository with the mock collection
		repo := repository.NewCommentRepository(mockCollection, mockCtx)

		// Valid ObjectID for the comment
		commentID := "507f1f77bcf86cd799439011"
		objID, _ := primitive.ObjectIDFromHex(commentID)

		mockDeleteCount := int64(1)
		// Mock the DeleteOne method to return an error
		mockCollection.On("DeleteOne", mockCtx, bson.M{"_id": objID}).Return(mockDeleteCount, assert.AnError)

		// Call the DeleteComment method
		err := repo.DeleteComment(commentID)

		// Assert that there was an error
		assert.Error(t, err)
		assert.Equal(t, assert.AnError, err)

		// Verify that DeleteOne was called with the correct arguments
		mockCollection.AssertCalled(t, "DeleteOne", mockCtx, bson.M{"_id": objID})
		mockCollection.AssertExpectations(t)
	})
}

func TestGetCommentsByBlogPostId(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Create a mock collection
		mockCollection := new(mocks.Collection)
		// Create a mock context
		mockCtx := context.Background()

		// Create a mock cursor
		mockCursor := new(mocks.Cursor)

		// Instantiate the repository with the mock collection
		repo := repository.NewCommentRepository(mockCollection, mockCtx)

		// Valid ObjectID for the blog post
		blogPostID := "507f1f77bcf86cd799439011"
		objID, _ := primitive.ObjectIDFromHex(blogPostID)

		// Mock the Find method to return the mock cursor
		mockCollection.On("Find", mockCtx, bson.M{"blogPostId": objID}).Return(mockCursor, nil)

		// Mock the cursor's Next and Decode methods
		mockCursor.On("Next", mockCtx).Return(true).Once() // Simulate a single comment
		mockCursor.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*entities.Comment)
			*arg = entities.Comment{ID: objID, Content: "Sample Comment"}
		}).Return(nil)

		mockCursor.On("Next", mockCtx).Return(false) // No more comments
		mockCursor.On("Err").Return(nil)
		mockCursor.On("Close", mockCtx).Return(nil)

		// Call the GetCommentsByBlogPostId method
		comments, err := repo.GetCommentsByBlogPostId(blogPostID)

		// Assert that there was no error and the comment was retrieved correctly
		assert.NoError(t, err)
		assert.Len(t, comments, 1)
		assert.Equal(t, "Sample Comment", comments[0].Content)

		// Verify that the mock methods were called with the correct arguments
		mockCollection.AssertCalled(t, "Find", mockCtx, bson.M{"blogPostId": objID})
		mockCursor.AssertExpectations(t)
	})

	t.Run("Invalid ObjectID Error", func(t *testing.T) {
		// Create a mock collection
		mockCollection := new(mocks.Collection)
		// Create a mock context
		mockCtx := context.Background()

		// Instantiate the repository with the mock collection
		repo := repository.NewCommentRepository(mockCollection, mockCtx)

		// Invalid ObjectID for the blog post
		invalidBlogPostID := "invalidID"

		// Call the GetCommentsByBlogPostId method with an invalid ID
		comments, err := repo.GetCommentsByBlogPostId(invalidBlogPostID)

		// Assert that there was an error and no comments were retrieved
		assert.Error(t, err)
		assert.Nil(t, comments)
		assert.Contains(t, err.Error(), "the provided hex string is not a valid ObjectID")

		// Verify that Find was never called
		mockCollection.AssertNotCalled(t, "Find", mock.Anything, mock.Anything)
	})

	t.Run("Find Error", func(t *testing.T) {
		// Create a mock collection
		mockCollection := new(mocks.Collection)
		// Create a mock context
		mockCtx := context.Background()

		// Instantiate the repository with the mock collection
		repo := repository.NewCommentRepository(mockCollection, mockCtx)

		// Valid ObjectID for the blog post
		blogPostID := "507f1f77bcf86cd799439011"
		objID, _ := primitive.ObjectIDFromHex(blogPostID)

		// Mock the Find method to return an error
		mockCollection.On("Find", mockCtx, bson.M{"blogPostId": objID}).Return(nil, assert.AnError)

		// Call the GetCommentsByBlogPostId method
		comments, err := repo.GetCommentsByBlogPostId(blogPostID)

		// Assert that there was an error and no comments were retrieved
		assert.Error(t, err)
		assert.Nil(t, comments)
		assert.Equal(t, assert.AnError, err)

		// Verify that Find was called with the correct arguments
		mockCollection.AssertCalled(t, "Find", mockCtx, bson.M{"blogPostId": objID})
	})
}

func TestGetCommentByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Create a mock collection
		mockCollection := new(mocks.Collection)
		// Create a mock context
		mockCtx := context.Background()

		// Instantiate the repository with the mock collection
		repo := repository.NewCommentRepository(mockCollection, mockCtx)

		// Valid ObjectID for the comment
		commentID := "507f1f77bcf86cd799439011"
		objID, _ := primitive.ObjectIDFromHex(commentID)

		mockSingleResult := new(mocks.SingleResult)
		// Mock the FindOne method to return the comment
		mockCollection.On("FindOne", mockCtx, bson.M{"_id": objID}).Return(mockSingleResult)
		mockSingleResult.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*entities.Comment)
			*arg = entities.Comment{ID: objID, Content: "Sample Comment"}
		}).Return(nil)

		// Call the GetCommentByID method
		comment, err := repo.GetCommentById(commentID)

		// Assert that there was no error and the comment was retrieved correctly
		assert.NoError(t, err)
		assert.NotNil(t, comment)

		// Verify that FindOne was called with the correct arguments
		mockCollection.AssertCalled(t, "FindOne", mockCtx, bson.M{"_id": objID})
		mockCollection.AssertExpectations(t)
	})

	t.Run("Invalid ObjectID Error", func(t *testing.T) {
		// Create a mock collection
		mockCollection := new(mocks.Collection)
		// Create a mock context
		mockCtx := context.Background()

		// Instantiate the repository with the mock collection
		repo := repository.NewCommentRepository(mockCollection, mockCtx)

		// Invalid ObjectID for the comment
		invalidCommentID := "invalidID"

		// Call the GetCommentByID method with an invalid ID
		comment, err := repo.GetCommentById(invalidCommentID)

		// Assert that there was an error and no comment was retrieved
		assert.Error(t, err)
		assert.Nil(t, comment)
		assert.Contains(t, err.Error(), "the provided hex string is not a valid ObjectID")

		// Verify that FindOne was never called
		mockCollection.AssertNotCalled(t, "FindOne", mock.Anything, mock.Anything)
	})

	t.Run("FindOne Error", func(t *testing.T) {
		// Create a mock collection
		mockCollection := new(mocks.Collection)
		// Create a mock context
		mockCtx := context.Background()

		// Instantiate the repository with the mock collection
		repo := repository.NewCommentRepository(mockCollection, mockCtx)

		// Valid ObjectID for the comment
		commentID := "507f1f77bcf86cd799439011"

		objectID, _ := primitive.ObjectIDFromHex(commentID)

		// Create a mock single result
		mockSingleResult := new(mocks.SingleResult)

		// Mock the FindOne method to return an error
		mockCollection.On("FindOne", mockCtx, bson.M{"_id": objectID}).Return(mockSingleResult)
		mockSingleResult.On("Decode", mock.Anything).Return(assert.AnError)

		// Call the GetCommentByID method
		comment, err := repo.GetCommentById(commentID)

		// Assert that there was an error and no comment was retrieved
		assert.Error(t, err)
		assert.Nil(t, comment)
		assert.Equal(t, assert.AnError, err)

		// Verify that FindOne was called with the correct arguments
		mockCollection.AssertCalled(t, "FindOne", mockCtx, bson.M{"_id": objectID})
	})
}