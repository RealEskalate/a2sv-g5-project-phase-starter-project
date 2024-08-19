package test

import (
	"context"
	"testing"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/RealEskalate/blogpost/mocks"
	"github.com/RealEskalate/blogpost/repository"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogRepoTestSuite struct {
	suite.Suite
	mockColl         *mocks.CollectionInterface
	mockCursor       *mocks.CursorInterface
	mockSingleResult *mocks.SingleResultInterface
	mockDeleteResult *mocks.DeleteResultInterface
	repo             *repository.BlogRepository
}

func (suite *BlogRepoTestSuite) SetupTest() {
	suite.mockColl = mocks.NewCollectionInterface(suite.T())
	suite.mockCursor = mocks.NewCursorInterface(suite.T())
	suite.mockSingleResult = mocks.NewSingleResultInterface(suite.T())
	suite.mockDeleteResult = mocks.NewDeleteResultInterface(suite.T())

	suite.repo = repository.NewBlogRepository(suite.mockColl)
}

func (suite *BlogRepoTestSuite) TestFindById() {
	suite.mockColl.On("FindOne", context.TODO(), mock.MatchedBy(func(m interface{}) bool {
		_, ok := m.(bson.M)
		return ok
	})).Return(suite.mockSingleResult)
	suite.mockSingleResult.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
		blog := args.Get(0).(*domain.Blog)
		*blog = domain.Blog{ID: primitive.NewObjectID()}
	}).Return(nil)

	id := primitive.NewObjectID()
	blog, err := suite.repo.GetOneBlogDocument(id.Hex())
	suite.NoError(err)
	suite.NotNil(blog.ID)

	suite.mockColl.AssertExpectations(suite.T())
	suite.mockSingleResult.AssertExpectations(suite.T())
}

func (suite *BlogRepoTestSuite) TestGetAllBlogs() {

	suite.mockColl.On("Find", context.TODO(), bson.M{}, mock.Anything).Return(suite.mockCursor, nil)

	suite.mockCursor.On("Next", context.TODO()).Return(true).Once()
	suite.mockCursor.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
		blog := args.Get(0).(*domain.Blog)
		*blog = domain.Blog{Title: "test title"}
	}).Return(nil)

	suite.mockCursor.On("Next", context.TODO()).Return(false).Once()

	suite.mockCursor.On("Close", context.TODO()).Return(nil)

	blogs, err := suite.repo.GetBlogDocuments(0, 10)

	suite.NoError(err)
	suite.Len(blogs, 1)
	suite.Equal("test title", blogs[0].Title)

	suite.mockColl.AssertExpectations(suite.T())
	suite.mockCursor.AssertExpectations(suite.T())
}

func (suite *BlogRepoTestSuite) TestUpdateBlogById() {
	suite.mockColl.On("UpdateOne", context.TODO(), mock.MatchedBy(func(m interface{}) bool {
		_, ok := m.(bson.M)
		return ok
	}), mock.MatchedBy(func(m interface{}) bool {
		_, ok := m.(bson.M)
		return ok
	})).Return(&mongo.UpdateResult{}, nil)

	id := primitive.NewObjectID().Hex()
	blog := domain.Blog{ID: primitive.NewObjectID(), Title: "updated title"}
	updatedBlog, err := suite.repo.UpdateBlogDocument(id, blog)
	suite.NoError(err)
	suite.Equal("updated title", updatedBlog.Title)

	suite.mockColl.AssertExpectations(suite.T())
}

func (suite *BlogRepoTestSuite) TestCreateBlog() {
	suite.mockColl.On("InsertOne", context.TODO(), mock.Anything).Return(&mongo.InsertOneResult{}, nil)

	blog := domain.Blog{Title: "test title"}
	createdBlog, err := suite.repo.CreateBlogDocument(blog)
	suite.NoError(err)
	suite.Equal("test title", createdBlog.Title)

	suite.mockColl.AssertExpectations(suite.T())
}

func (suite *BlogRepoTestSuite) TestDeleteBlogByID() {
	suite.mockColl.On("DeleteOne", context.TODO(), mock.MatchedBy(func(m interface{}) bool {
		query, ok := m.(bson.M)
		if !ok {
			return false
		}

		_, idExists := query["_id"]
		_, userExists := query["user._id"]
		return idExists && userExists
	})).Return(suite.mockDeleteResult, nil)
	suite.mockDeleteResult.On("DeletedCount").Return(int64(1))

	id := primitive.NewObjectID().Hex()
	err := suite.repo.DeleteBlogDocument(id, primitive.NewObjectID())
	suite.NoError(err)

	suite.mockColl.AssertExpectations(suite.T())
	suite.mockDeleteResult.AssertExpectations(suite.T())
}

func TestBlogRepoTestSuite(t *testing.T) {
	suite.Run(t, new(BlogRepoTestSuite))
}
