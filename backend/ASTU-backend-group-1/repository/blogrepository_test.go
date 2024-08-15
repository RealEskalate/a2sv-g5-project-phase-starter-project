package repository

import (
	"astu-backend-g1/domain"
	"time"
	// "astu-backend-g1/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	mongomocks "github.com/sv-tools/mongoifc/mocks/mockery"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogRespositoryTestSuite struct {
	suite.Suite
	coll           *mongomocks.Collection
	BlogRepository domain.BlogRepository
}

func (suite *BlogRespositoryTestSuite) SetupSuite() {
	suite.coll = &mongomocks.Collection{}
	suite.BlogRepository = NewBlogRepository(suite.coll)
}

// func (suite *BlogRespositoryTestSuite) TearDownSuite() {
//   suite.coll.AssertExpectations(suite.T())
// }

func (suite *BlogRespositoryTestSuite) TestCreate() {
	assert := assert.New(suite.T())
	suite.coll.On("InsertOne", mock.Anything, mock.Anything, mock.Anything).Return(&mongo.InsertOneResult{
		InsertedID: primitive.NewObjectID(),
	}, nil)
	expectedBlog := domain.Blog{
		ID:       "test_id",
        Title:    "Test Title",
        Content:  "Test Content",
        AuthorID: "test_author_id",
        Date:     time.Now(),
        Tags:     []string{"test_tag1", "test_tag2"},
		
	}
	result, err := suite.BlogRepository.Create(expectedBlog)
	assert.NoError(err)
	assert.Equal(result, expectedBlog)
	suite.coll.On("InsertOne", mock.Anything, mock.Anything, mock.Anything).Return(&mongo.InsertOneResult{
		InsertedID: primitive.NewObjectID(),
	}, nil)
	expectedBlog = domain.Blog{
		ID:       "",
        Title:    "Test Title",
        Content:  "Test Content",
        AuthorID: "test_author_id",
        Date:     time.Now(),
        Tags:     []string{"test_tag1", "test_tag2"},
		
	}
	result, err = suite.BlogRepository.Create(expectedBlog)
	assert.Error(err)
	assert.Equal(result, domain.Blog{})
}

func TestBlogRepository(t *testing.T) {
	suite.Run(t, new(BlogRespositoryTestSuite))
}
