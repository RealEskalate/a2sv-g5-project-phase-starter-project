package usecases

import (
	"blogapp/Domain"
	"blogapp/mocks"
	"context"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogUsecaseSuite struct {
	suite.Suite
	blogUsecase Domain.BlogUseCase
	repo        mocks.BlogRepository
	context     context.Context
}

func (suite *BlogUsecaseSuite) SetupTest() {
	suite.repo = *new(mocks.BlogRepository)
	suite.context = *new(context.Context)
	suite.blogUsecase = NewBlogUseCase(&suite.repo)
}
func (suite *BlogUsecaseSuite) TestCreateBlog() {
	// Create a real gin.Context for testing
	c, _ := gin.CreateTestContext(nil)
	post := &Domain.Post{}

	suite.repo.On("CreateBlog", mock.AnythingOfType("*context.timerCtx"), mock.AnythingOfType("*Domain.Post")).Return(nil, 1)

	err, number := suite.blogUsecase.CreateBlog(c, post)

	suite.Nil(err, "should be nil")
	suite.Equal(number, 1, "should equal")
	suite.repo.AssertExpectations(suite.T())
}
func (suite *BlogUsecaseSuite) TestGetPostBySlug() {
	c, _ := gin.CreateTestContext(nil)

	suite.repo.On("GetPostBySlug", mock.AnythingOfType("*context.timerCtx"), "").Return([]*Domain.Post{}, nil, 1)

	posts, err, number := suite.blogUsecase.GetPostBySlug(c, "")
	suite.IsType(posts, []*Domain.Post{}, "should equal")
	suite.Nil(err, "should be nil")
	suite.Equal(number, 1, "should equal")
}
func (suite *BlogUsecaseSuite) TestGetPostByID() {
	c, _ := gin.CreateTestContext(nil)

	suite.repo.On("GetPostByID", mock.AnythingOfType("*context.timerCtx"), primitive.ObjectID{}).Return(&Domain.Post{}, nil, 1)

	post, err, number := suite.blogUsecase.GetPostByID(c, primitive.ObjectID{})

	suite.IsType(post, &Domain.Post{}, "should equal")
	suite.Nil(err, "should be nil")
	suite.Equal(number, 1, "should equal")
}
func (suite *BlogUsecaseSuite) TestGetPostByAuthorID() {
	c, _ := gin.CreateTestContext(nil)

	suite.repo.On("GetPostByAuthorID", mock.AnythingOfType("*context.timerCtx"), primitive.ObjectID{}).Return([]*Domain.Post{}, nil, 1)

	posts, err, number := suite.blogUsecase.GetPostByAuthorID(c, primitive.ObjectID{})

	suite.IsType(posts, []*Domain.Post{}, "should equal")
	suite.Nil(err, "should be nil")
	suite.Equal(number, 1, "should equal")
}

// func (suite *BlogUsecaseSuite) TestUpdatePostByID() {
//   c, _ := gin.CreateTestContext(nil)
//   post := &Domain.Post{}
//   suite.repo.On("UpdatePostByID", mock.AnythingOfType("*context.timerCtx"), primitive.ObjectID{}, post).Return(Domain.Post{}, nil, 1)

//   err, number := suite.blogUsecase.UpdatePostByID(c, primitive.ObjectID{}, post)

//   suite.Nil(err, "should be nil")
//   suite.Equal(number, 1, "should equal")
// }

func (suite *BlogUsecaseSuite) TestGetTags() {
	c, _ := gin.CreateTestContext(nil)

	suite.repo.On("GetTags", mock.AnythingOfType("*context.timerCtx"), primitive.ObjectID{}).Return([]*Domain.Tag{}, nil, 1)

	tags, err, number := suite.blogUsecase.GetTags(c, primitive.ObjectID{})

	suite.IsType(tags, []*Domain.Tag{}, "should equal")
	suite.Nil(err, "should be nil")
	suite.Equal(number, 1, "should equal")
}

func (suite *BlogUsecaseSuite) TestGetComments() {
	c, _ := gin.CreateTestContext(nil)

	suite.repo.On("GetComments", mock.AnythingOfType("*context.timerCtx"), primitive.ObjectID{}).Return([]*Domain.Comment{}, nil, 1)

	Comments, err, number := suite.blogUsecase.GetComments(c, primitive.ObjectID{})

	suite.IsType(Comments, []*Domain.Comment{}, "should equal")
	suite.Nil(err, "should be nil")
	suite.Equal(number, 1, "should equal")
}
func (suite *BlogUsecaseSuite) TestGetAllPosts() {
	c, _ := gin.CreateTestContext(nil)

	suite.repo.On("GetAllPosts", mock.AnythingOfType("*context.timerCtx"), Domain.Filter{}).Return([]*Domain.Post{}, nil, 1, Domain.PaginationMetaData{})

	Posts, err, number, paginationMetaData := suite.blogUsecase.GetAllPosts(c, Domain.Filter{})

	suite.IsType(paginationMetaData, Domain.PaginationMetaData{}, "should equal")
	suite.IsType(Posts, []*Domain.Post{}, "should equal")
	suite.Nil(err, "should be nil")
	suite.Equal(number, 1, "should equal")
}

func (suite *BlogUsecaseSuite) TestAddTagToPost() {
	c, _ := gin.CreateTestContext(nil)

	suite.repo.On("AddTagToPost", mock.AnythingOfType("*context.timerCtx"), primitive.ObjectID{}, "").Return(nil, 1)

	err, number := suite.blogUsecase.AddTagToPost(c, primitive.ObjectID{}, "")

	suite.Nil(err, "should be nil")
	suite.Equal(number, 1, "should equal")
}
func (suite *BlogUsecaseSuite) TestLikePost() {
	c, _ := gin.CreateTestContext(nil)

	suite.repo.On("LikePost", mock.AnythingOfType("*context.timerCtx"), primitive.ObjectID{}, primitive.ObjectID{}).Return(nil, 1, "")

	err, number, s := suite.blogUsecase.LikePost(c, primitive.ObjectID{}, primitive.ObjectID{})

	suite.Nil(err, "should be nil")
	suite.Equal(number, 1, "should equal")
	suite.Equal(s, "", "should equal")
}
func (suite *BlogUsecaseSuite) TestDislikePost() {
	c, _ := gin.CreateTestContext(nil)

	suite.repo.On("DislikePost", mock.AnythingOfType("*context.timerCtx"), primitive.ObjectID{}, primitive.ObjectID{}).Return(nil, 1, "")

	err, number, s := suite.blogUsecase.DislikePost(c, primitive.ObjectID{}, primitive.ObjectID{})

	suite.Nil(err, "should be nil")
	suite.Equal(number, 1, "should equal")
	suite.Equal(s, "", "should equal")
}

func (suite *BlogUsecaseSuite) TestSearchPosts() {
	c, _ := gin.CreateTestContext(nil)

	suite.repo.On("SearchPosts", mock.AnythingOfType("*context.timerCtx"), "", Domain.Filter{}).Return([]*Domain.Post{}, nil, 1, Domain.PaginationMetaData{})

	Posts, err, number, paginationMetaData := suite.blogUsecase.SearchPosts(c, "", Domain.Filter{})

	suite.IsType(paginationMetaData, Domain.PaginationMetaData{}, "should equal")
	suite.IsType(Posts, []*Domain.Post{}, "should equal")
	suite.Nil(err, "should be nil")
	suite.Equal(number, 1, "should equal")
}

func (suite *BlogUsecaseSuite) TestDeletePost() {
	c, _ := gin.CreateTestContext(nil)

	suite.repo.On("DeletePost", mock.AnythingOfType("*context.timerCtx"), primitive.ObjectID{}).Return(nil, 1)

	err, number := suite.blogUsecase.DeletePost(c, primitive.ObjectID{})

	suite.Nil(err, "should be nil")
	suite.Equal(number, 1, "should equal")
}

func TestBloBlogUsecaseSuite(t *testing.T) {
	suite.Run(t, new(BlogUsecaseSuite))
}
