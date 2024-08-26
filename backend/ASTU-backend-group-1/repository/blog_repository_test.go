package repository

import (
	"astu-backend-g1/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	mongomocks "github.com/sv-tools/mongoifc/mocks/mockery"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var mockBlogs = []domain.Blog{
	{BlogId: "1", Title: "title 1", Content: "content 1", AuthorId: "author 1", Date: time.Now(), Tags: []string{"tag1"}},
	{BlogId: "2", Title: "title 2", Content: "content 2", AuthorId: "author 2", Date: time.Now(), Tags: []string{"tag2"}},
	{BlogId: "3", Title: "title 3", Content: "content 3", AuthorId: "author 3", Date: time.Now(), Tags: []string{"tag3"}},
	{BlogId: "4", Title: "title 4", Content: "content 4", AuthorId: "author 4", Date: time.Now(), Tags: []string{"tag4"}},
	{BlogId: "5", Title: "title 5", Content: "content 5", AuthorId: "author 5", Date: time.Now(), Tags: []string{"tag5"}},
	{BlogId: "6", Title: "title 6", Content: "content 6", AuthorId: "author 6", Date: time.Now(), Tags: []string{"tag6"}},
	{BlogId: "7", Title: "title 7", Content: "content 7", AuthorId: "author 7", Date: time.Now(), Tags: []string{"tag7"}},
	{BlogId: "8", Title: "title 8", Content: "content 8", AuthorId: "author 8", Date: time.Now(), Tags: []string{"tag8"}},
	{BlogId: "9", Title: "title 9", Content: "content 9", AuthorId: "author 9", Date: time.Now(), Tags: []string{"tag9"}},
	{BlogId: "10", Title: "title 10", Content: "content 10", AuthorId: "author 10", Date: time.Now(), Tags: []string{"tag10"}},
}

type BlogRespositoryTestSuite struct {
	suite.Suite
	client *mongomocks.Client
	bcoll           *mongomocks.Collection
	ccoll           *mongomocks.Collection
	rcoll           *mongomocks.Collection
	BlogRepository domain.BlogRepository
}

func (suite *BlogRespositoryTestSuite) SetupSuite() {
	suite.client = &mongomocks.Client{}
	suite.bcoll = &mongomocks.Collection{}
	suite.ccoll = &mongomocks.Collection{}
	suite.rcoll = &mongomocks.Collection{}
	suite.BlogRepository = NewBlogRepository(suite.client,suite.bcoll,suite.ccoll,suite.rcoll)
}

func (suite *BlogRespositoryTestSuite) TearDownSuite() {
	suite.bcoll.AssertExpectations(suite.T())
	suite.rcoll.AssertExpectations(suite.T())
	suite.ccoll.AssertExpectations(suite.T())
}

func (suite *BlogRespositoryTestSuite) TestCreate() {
	assert := assert.New(suite.T())
	suite.bcoll.On("InsertOne", mock.Anything, mock.Anything, mock.Anything).Return(&mongo.InsertOneResult{
		InsertedID: primitive.NewObjectID(),
	}, nil)
	result, err := suite.BlogRepository.CreateBlog(mockBlogs[0])
	assert.NoError(err)
	assert.Equal(result, mockBlogs[0])
}

func (suite *BlogRespositoryTestSuite) TestGet() {
	assert := assert.New(suite.T())
	suite.T().Parallel()
	suite.T().Run("Getting all Blogs", func(t *testing.T) {
		cur := &mongomocks.Cursor{}
		for i, Blog := range mockBlogs {
			cur.On("Next", mock.Anything).Return(i < len(mockBlogs)).Once()
			cur.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
				arg := args.Get(0).(*domain.Blog)
				*arg = Blog
			}).Return(nil).Once()
		}
		cur.On("Next", mock.Anything).Return(false).Once()
		suite.bcoll.On("Find", mock.Anything, mock.Anything, mock.Anything).Return(cur, nil)
		defer cur.AssertExpectations(suite.T())
		result, err := suite.BlogRepository.GetBlog(domain.BlogFilterOption{})
		assert.NoError(err)
		assert.Equal(mockBlogs, result)
	})
	// todo this is for a single result
	suite.T().Run("Getting by Blogname", func(t *testing.T) {
		cur := &mongomocks.Cursor{}
		singleResult := &mongomocks.SingleResult{}
		singleResult.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*domain.Blog)
			*arg = mockBlogs[0]
		}).Return(nil)
		suite.bcoll.On("FindOne", mock.Anything, mock.Anything, mock.Anything).Return(singleResult)
		defer cur.AssertExpectations(suite.T())
		result, err := suite.BlogRepository.GetBlog(domain.BlogFilterOption{
			Filter: domain.BlogFilters{
				Title: mockBlogs[0].Title,
			},
		})
		assert.NoError(err)
		assert.Equal(mockBlogs, result[0])
	})
	suite.T().Run("Getting by Email", func(t *testing.T) {
		cur := &mongomocks.Cursor{}
		singleResult := &mongomocks.SingleResult{}
		singleResult.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*domain.Blog)
			*arg = mockBlogs[0]
		}).Return(nil)
		suite.bcoll.On("Find", mock.Anything, mock.Anything, mock.Anything).Return(singleResult)
		defer cur.AssertExpectations(suite.T())
		result, err := suite.BlogRepository.GetBlog(domain.BlogFilterOption{
			Filter: domain.BlogFilters{
                Tags: mockBlogs[0].Tags,
            },
		})
		assert.NoError(err)
		assert.Equal(mockBlogs, result[0])
	})
	suite.T().Run("Getting by Id", func(t *testing.T) {
		cur := &mongomocks.Cursor{}
		singleResult := &mongomocks.SingleResult{}
		singleResult.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*domain.Blog)
			*arg = mockBlogs[0]
		}).Return(nil)
		suite.bcoll.On("FindOne", mock.Anything, mock.Anything, mock.Anything).Return(singleResult)
		defer cur.AssertExpectations(suite.T())
		result, err := suite.BlogRepository.GetBlog(domain.BlogFilterOption{
			Filter: domain.BlogFilters{
                AuthorId: mockBlogs[0].AuthorId,
            },
		})
		assert.NoError(err)
		assert.Equal(mockBlogs, result[0])
	})
}

// func TestBlogRepository(t *testing.T) {
// 	suite.Run(t, new(BlogRespositoryTestSuite))
// }