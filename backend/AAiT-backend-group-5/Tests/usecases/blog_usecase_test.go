package usecases_test

import (
	"context"
	// "encoding/json"
	"strconv"
	"testing"
	"time"

	// "github.com/go-redis/redis/v8"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	config "github.com/aait.backend.g5.main/backend/Config"
	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	mocks "github.com/aait.backend.g5.main/backend/Mocks"
	usecases "github.com/aait.backend.g5.main/backend/UseCases"
)

type BlogUsecaseTestSuite struct {
	suite.Suite
	repositoryMock   *mocks.MockBlogRepository
	userRepoMock     *mocks.MockUserRepository
	popularityMock   *mocks.MockBlogPopularityActionRepository
	commentMock      *mocks.MockBlogCommentRepository
	cacheServiceMock *mocks.MockRedisCache
	blogHelperMock   *mocks.MockBlogHelper
	blogUsecase      interfaces.BlogUsecase
	ctrl             *gomock.Controller
}

func (suite *BlogUsecaseTestSuite) SetupSuite() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.repositoryMock = mocks.NewMockBlogRepository(suite.ctrl)
	suite.userRepoMock = mocks.NewMockUserRepository(suite.ctrl)
	suite.popularityMock = mocks.NewMockBlogPopularityActionRepository(suite.ctrl)
	suite.commentMock = mocks.NewMockBlogCommentRepository(suite.ctrl)
	suite.cacheServiceMock = mocks.NewMockRedisCache(suite.ctrl)
	suite.blogHelperMock = mocks.NewMockBlogHelper(suite.ctrl)

	suite.blogUsecase = usecases.NewblogUsecase(
		suite.repositoryMock,
		suite.cacheServiceMock,
		config.Env{},
		time.Minute,
		suite.blogHelperMock,
		suite.userRepoMock,
		suite.popularityMock,
		suite.commentMock,
	)
}

func (suite *BlogUsecaseTestSuite) TearDownSuite() {
	suite.ctrl.Finish()
}

func (suite *BlogUsecaseTestSuite) TestCreateBlog_Success() {
	ctx := context.Background()
	blog := &models.Blog{Title: "New Blog", Slug: "new-blog"}
	slug := "new-blog"
	blogResponse := &dtos.BlogResponse{
		Blog:       *blog,
		Comments:   []models.Comment{},
		Popularity: models.Popularity{},
	}

	suite.blogHelperMock.
		EXPECT().
		CreateSlug(blog.Title).
		Return(slug)

	suite.repositoryMock.
		EXPECT().
		CreateBlog(ctx, blog).
		Return(blog, nil)

	suite.cacheServiceMock.
		EXPECT().
		InvalidateAllBlogCaches(ctx).
		Return(nil)

	result, err := suite.blogUsecase.CreateBlog(ctx, blog)
	suite.Nil(err)
	suite.Equal(blogResponse, result)
}

func (suite *BlogUsecaseTestSuite) TestGetBlog_Success() {
	ctx := context.Background()
	blogID := "1"
	blog := &models.Blog{ID: blogID}
	blogResponse := &dtos.BlogResponse{
		Blog:       *blog,
		Comments:   []models.Comment{},
		Popularity: models.Popularity{},
	}

	suite.blogHelperMock.
		EXPECT().
		FetchFromCacheOrRepo(ctx, blogID, gomock.Any()).
		Return(blog, nil)

	suite.repositoryMock.
		EXPECT().
		IncreaseView(ctx, blog.ID).
		Return(nil)

	suite.blogHelperMock.
		EXPECT().
		FetchComments(ctx, blog.ID).
		Return([]models.Comment{}, nil)

	suite.blogHelperMock.
		EXPECT().
		FetchPopularity(ctx, blog.ID).
		Return(&models.Popularity{}, nil)

	result, err := suite.blogUsecase.GetBlog(ctx, blogID)
	suite.Nil(err)
	suite.Equal(blogResponse, result)
}

func (suite *BlogUsecaseTestSuite) TestGetBlogs_Success() {
	ctx := context.Background()
	page := 1
	blogs := []*models.Blog{{}}
	blogResponse := []*dtos.BlogResponse{{}}

	suite.blogHelperMock.
		EXPECT().
		FetchFromCacheOrRepoBlogs(ctx, strconv.Itoa(page), gomock.Any()).
		Return(blogs, nil)

	suite.blogHelperMock.
		EXPECT().
		GetBlogs(ctx, blogs).
		Return(blogResponse, nil)

	result, err := suite.blogUsecase.GetBlogs(ctx, page)
	suite.Nil(err)
	suite.Equal(blogResponse, result)
}

func (suite *BlogUsecaseTestSuite) TestSearchBlogs_Success() {
	ctx := context.Background()
	filter := dtos.FilterBlogRequest{}
	blogs := []*models.Blog{{}}
	blogResponse := []*dtos.BlogResponse{{}}

	suite.repositoryMock.
		EXPECT().
		SearchBlogs(ctx, filter).
		Return(blogs, nil)

	suite.blogHelperMock.
		EXPECT().
		GetBlogs(ctx, blogs).
		Return(blogResponse, nil)

	result, err := suite.blogUsecase.SearchBlogs(ctx, filter)
	suite.Nil(err)
	suite.Equal(blogResponse, result)
}

func (suite *BlogUsecaseTestSuite) TestUpdateBlog_Success() {
	ctx := context.Background()
	blogID := "1"
	blog := &models.Blog{ID: blogID}

	suite.repositoryMock.
		EXPECT().
		GetBlog(ctx, blogID).
		Return(blog, nil)

	suite.repositoryMock.
		EXPECT().
		UpdateBlog(ctx, blogID, blog).
		Return(nil)

	suite.cacheServiceMock.
		EXPECT().
		InvalidateAllBlogCaches(ctx).
		Return(nil)

	err := suite.blogUsecase.UpdateBlog(ctx, blogID, blog)
	suite.Nil(err)
}

func (suite *BlogUsecaseTestSuite) TestDeleteBlog_Success() {
	ctx := context.Background()
	deleteBlogReq := dtos.DeleteBlogRequest{
		BlogID:   "1",
		AuthorID: "author1",
	}
	blog := &models.Blog{ID: deleteBlogReq.BlogID, AuthorID: deleteBlogReq.AuthorID}
	user := &models.User{ID: deleteBlogReq.AuthorID, Role: models.RoleUser}

	suite.repositoryMock.
		EXPECT().
		GetBlog(ctx, deleteBlogReq.BlogID).
		Return(blog, nil)

	suite.userRepoMock.
		EXPECT().
		GetUserByID(ctx, deleteBlogReq.AuthorID).
		Return(user, nil)

	suite.commentMock.
		EXPECT().
		DeleteComments(ctx, deleteBlogReq.BlogID).
		Return(nil)

	suite.repositoryMock.
		EXPECT().
		DeleteBlog(ctx, deleteBlogReq.BlogID).
		Return(nil)

	suite.cacheServiceMock.
		EXPECT().
		InvalidateAllBlogCaches(ctx).
		Return(nil)

	err := suite.blogUsecase.DeleteBlog(ctx, deleteBlogReq)
	suite.Nil(err)
}

func (suite *BlogUsecaseTestSuite) TestTrackPopularity_Success() {
	ctx := context.Background()
	popularity := dtos.TrackPopularityRequest{BlogID: "1", UserID: "user1", Action: models.Like}

	suite.popularityMock.
		EXPECT().
		GetBlogPopularityAction(ctx, popularity.BlogID, popularity.UserID).
		Return(nil, nil)

	suite.popularityMock.
		EXPECT().
		Like(ctx, popularity).
		Return(nil)

	err := suite.blogUsecase.TrackPopularity(ctx, popularity)
	suite.Nil(err)
}

func (suite *BlogUsecaseTestSuite) TestTrackPopularity_UndoAction() {
	ctx := context.Background()
	popularity := dtos.TrackPopularityRequest{BlogID: "1", UserID: "user1", Action: models.Like}
	existingAction := &models.PopularityAction{Action: models.Like}

	suite.popularityMock.
		EXPECT().
		GetBlogPopularityAction(ctx, popularity.BlogID, popularity.UserID).
		Return(existingAction, nil)

	suite.popularityMock.
		EXPECT().
		UndoLike(ctx, popularity).
		Return(nil)

	err := suite.blogUsecase.TrackPopularity(ctx, popularity)
	suite.Nil(err)
}

func TestBlogUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(BlogUsecaseTestSuite))
}
