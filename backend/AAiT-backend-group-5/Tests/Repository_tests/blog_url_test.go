package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive" 

	"go.mongodb.org/mongo-driver/mongo"

	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	mocks "github.com/aait.backend.g5.main/backend/Mocks"
	repository "github.com/aait.backend.g5.main/backend/Repository"
)

type BlogUrlRepositorySuite struct {
	suite.Suite

	Urlcollection *mocks.Collection
	repository            *repository.URL_Repo
}

func (suite *BlogUrlRepositorySuite) SetupTest() {
	suite.Urlcollection = &mocks.Collection{}
	suite.repository = &repository.URL_Repo{
		Collection: suite.Urlcollection,
	}
}
// Test SaveURL Success Case
func (suite *BlogUrlRepositorySuite) TestSaveURL_Success() {
	url := models.URL{Token: "https://example.com"}

	// Mocking InsertOne to return a successful response
	suite.Urlcollection.On("InsertOne", mock.Anything, mock.Anything).Return(nil, nil)

	err := suite.repository.SaveURL(url, context.Background())

	suite.Nil(err)
	suite.Urlcollection.AssertExpectations(suite.T())
}

// Test SaveURL Failure Case
func (suite *BlogUrlRepositorySuite) TestSaveURL_Failure() {
	url := models.URL{Token: "https://example.com"}

	// Mocking InsertOne to return an error
	suite.Urlcollection.On("InsertOne", mock.Anything, mock.Anything).Return(nil, errors.New("insertion error"))

	err := suite.repository.SaveURL(url, context.Background())

	suite.NotNil(err)
	suite.Equal("insertion error", err.Message) // Assuming err.Message contains the error message in your models
	suite.Urlcollection.AssertExpectations(suite.T())
}

// Test GetURL Success Case
func (suite *BlogUrlRepositorySuite) TestGetURL_Success() {
	url := models.URL{
		ID:           primitive.NewObjectID(),
		ShortURLCode: "abc123",
		Token:        "https://example.com",
	}

	// Mocking FindOne to return a successful SingleResult with the expected URL
	singleResult := &mocks.SingleResult{}
	suite.Urlcollection.On("FindOne", mock.Anything, bson.D{{Key: "short_url", Value: "abc123"}}).Return(singleResult)
	singleResult.On("Decode", mock.AnythingOfType("*models.URL")).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*models.URL)
		*arg = url
	}).Return(nil)

	result, err := suite.repository.GetURL("abc123", context.Background())

	suite.Nil(err)
	suite.NotNil(result)
	suite.Equal(url.ShortURLCode, result.ShortURLCode)
	suite.Equal(url.Token, result.Token)
	suite.Urlcollection.AssertExpectations(suite.T())
	singleResult.AssertExpectations(suite.T())
}

// Test GetURL Failure Case (Document Not Found)
func (suite *BlogUrlRepositorySuite) TestGetURL_Failure_NotFound() {
	// Mocking FindOne to return a SingleResult that triggers a not found error
	singleResult := &mocks.SingleResult{}
	suite.Urlcollection.On("FindOne", mock.Anything, bson.D{{Key: "short_url", Value: "abc123"}}).Return(singleResult)
	singleResult.On("Decode", mock.AnythingOfType("*models.URL")).Return(mongo.ErrNoDocuments)

	result, err := suite.repository.GetURL("abc123", context.Background())

	suite.NotNil(err)
	suite.Nil(result)
	suite.Equal("mongo: no documents in result", err.Message) // Assuming err.Message contains the error message in your models
	suite.Urlcollection.AssertExpectations(suite.T())
	singleResult.AssertExpectations(suite.T())
}

// Test DeleteURL Success Case
func (suite *BlogUrlRepositorySuite) TestDeleteURL_Success() {
	// Mocking DeleteOne to return a successful response
	id := primitive.NewObjectID().Hex()
	suite.Urlcollection.On("DeleteOne", mock.Anything, bson.D{{Key: "_id", Value: id}}).Return(int64(1), nil)

	err := suite.repository.DeleteURL(id, context.Background())

	suite.Nil(err)
	suite.Urlcollection.AssertExpectations(suite.T())
}

// Test DeleteURL Failure Case
func (suite *BlogUrlRepositorySuite) TestDeleteURL_Failure() {
	// Mocking DeleteOne to return an error
	id := primitive.NewObjectID().Hex()
	suite.Urlcollection.On("DeleteOne", mock.Anything, bson.D{{Key: "_id", Value: id}}).Return(int64(0), errors.New("deletion error"))

	err := suite.repository.DeleteURL(id, context.Background())

	suite.NotNil(err)
	suite.Equal("deletion error", err.Message) // Assuming err.Message contains the error message in your models
	suite.Urlcollection.AssertExpectations(suite.T())
}

// Running the test suite
func TestBlogUrlRepositorySuite(t *testing.T) {
	suite.Run(t, new(BlogUrlRepositorySuite))
}