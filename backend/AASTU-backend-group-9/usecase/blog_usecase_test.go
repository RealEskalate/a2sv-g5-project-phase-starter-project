package usecase

import (
    "blog/domain"
    "blog/domain/mocks"
    "context"
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
	 "github.com/stretchr/testify/mock"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogUsecaseTestSuite struct {
    suite.Suite
    blogUsecase    *blogUsecase
    blogRepository *mocks.BlogRepository
    popularityRepo *mocks.PopularityRepository
    commentRepo    *mocks.CommentRepository
}

func (suite *BlogUsecaseTestSuite) SetupTest() {
    suite.blogRepository = new(mocks.BlogRepository)
    suite.popularityRepo = new(mocks.PopularityRepository)
    suite.commentRepo = new(mocks.CommentRepository)
    suite.blogUsecase = &blogUsecase{
        blogRepository: suite.blogRepository,
        popularityRepo: suite.popularityRepo,
        commentRepo:    suite.commentRepo,
        contextTimeout: time.Second * 10,
    }
}

func (suite *BlogUsecaseTestSuite) TestCreateBlog() {
    ctx := context.TODO()
    req := &domain.BlogCreationRequest{
        Title:   "Test Title",
        Content: "Test Content",
        Tags:    []string{"tag1", "tag2"},
    }
    claims := &domain.JwtCustomClaims{
        UserID: primitive.NewObjectID(),
    }

    suite.blogRepository.On("CreateBlog", ctx, mock.AnythingOfType("*domain.Blog")).Return(nil)

    resp, err := suite.blogUsecase.CreateBlog(ctx, req, claims)

    assert.NoError(suite.T(), err)
    assert.NotNil(suite.T(), resp)
    suite.blogRepository.AssertExpectations(suite.T())
}

func (suite *BlogUsecaseTestSuite) TestGetBlogByID() {
    ctx := context.TODO()
    id := primitive.NewObjectID()
    blog := &domain.Blog{
        ID: id,
    }

    suite.blogRepository.On("GetBlogByID", ctx, id).Return(blog, nil)

    resp, err := suite.blogUsecase.GetBlogByID(ctx, id)

    assert.NoError(suite.T(), err)
    assert.NotNil(suite.T(), resp)
    suite.blogRepository.AssertExpectations(suite.T())
}

func (suite *BlogUsecaseTestSuite) TestGetAllBlogs() {
    ctx := context.TODO()
    page, limit := 1, 10
    sortBy := "createdAt"
    blogs := []*domain.Blog{
        {ID: primitive.NewObjectID()},
    }

    suite.blogRepository.On("GetAllBlogs", ctx, page, limit, sortBy).Return(blogs, nil)

    resp, err := suite.blogUsecase.GetAllBlogs(ctx, page, limit, sortBy)

    assert.NoError(suite.T(), err)
    assert.NotNil(suite.T(), resp)
    suite.blogRepository.AssertExpectations(suite.T())
}

func (suite *BlogUsecaseTestSuite) TestUpdateBlog() {
    ctx := context.TODO()
    id := primitive.NewObjectID()
    req := &domain.BlogUpdateRequest{
        Title:   "Updated Title",
        Content: "Updated Content",
        Tags:    []string{"tag1", "tag2"},
    }
    blog := &domain.Blog{
        ID: id,
    }

    suite.blogRepository.On("GetBlogByID", ctx, id).Return(blog, nil)
    suite.blogRepository.On("UpdateBlog", ctx, blog).Return(nil)

    resp, err := suite.blogUsecase.UpdateBlog(ctx, id, req)

    assert.NoError(suite.T(), err)
    assert.NotNil(suite.T(), resp)
    suite.blogRepository.AssertExpectations(suite.T())
}

func (suite *BlogUsecaseTestSuite) TestDeleteBlog() {
    ctx := context.TODO()
    id := primitive.NewObjectID()

    suite.blogRepository.On("DeleteBlog", ctx, id).Return(nil)

    err := suite.blogUsecase.DeleteBlog(ctx, id)

    assert.NoError(suite.T(), err)
    suite.blogRepository.AssertExpectations(suite.T())
}
func (suite *BlogUsecaseTestSuite) TestTrackView() {
    ctx := context.TODO()
    postID := primitive.NewObjectID()
    suite.popularityRepo.On("IncrementPopularity", ctx, postID, "view").Return(nil).Once()

    err := suite.blogUsecase.TrackView(ctx, postID)
    assert.NoError(suite.T(), err)
    suite.popularityRepo.AssertExpectations(suite.T())
}
func (suite *BlogUsecaseTestSuite) TestTrackLike() {
    ctx := context.TODO()
    postID := primitive.NewObjectID()
    userID := primitive.NewObjectID()
    suite.popularityRepo.On("IncrementPopularity", ctx, postID, "like").Return(nil).Once()

    err := suite.blogUsecase.TrackLike(ctx, postID, userID)
    assert.NoError(suite.T(), err)
    suite.popularityRepo.AssertExpectations(suite.T())
}
func (suite *BlogUsecaseTestSuite) TestTrackDislike() {
    ctx := context.TODO()
    postID := primitive.NewObjectID()
    userID := primitive.NewObjectID()
    suite.popularityRepo.On("IncrementPopularity", ctx, postID, "dislike").Return(nil).Once()

    err := suite.blogUsecase.TrackDislike(ctx, postID, userID)
    assert.NoError(suite.T(), err)
    suite.popularityRepo.AssertExpectations(suite.T())
}
func (suite *BlogUsecaseTestSuite) TestAddComment() {
    ctx := context.TODO()
    postID := primitive.NewObjectID()
    userID := primitive.NewObjectID()
    comment := &domain.Comment{Content: "Test comment"}
    suite.commentRepo.On("AddComment", ctx, postID, userID, comment).Return(nil).Once()
    suite.popularityRepo.On("IncrementPopularity", ctx, postID, "comment").Return(nil).Once()

    err := suite.blogUsecase.AddComment(ctx, postID, userID, comment)
    assert.NoError(suite.T(), err)
    suite.commentRepo.AssertExpectations(suite.T())
    suite.popularityRepo.AssertExpectations(suite.T())
}
func (suite *BlogUsecaseTestSuite) TestGetComments() {
    ctx := context.TODO()
    postID := primitive.NewObjectID()
    comment := &domain.Comment{Content: "Test comment"}
    suite.commentRepo.On("GetComments", ctx, postID).Return([]*domain.Comment{comment}, nil).Once()

    comments, err := suite.blogUsecase.GetComments(ctx, postID)
    assert.NoError(suite.T(), err)
    assert.NotNil(suite.T(), comments)
    suite.commentRepo.AssertExpectations(suite.T())
}
func (suite *BlogUsecaseTestSuite) TestDeleteComment() {
    ctx := context.TODO()
    postID := primitive.NewObjectID()
    commentID := primitive.NewObjectID()
    userID := primitive.NewObjectID()
    suite.commentRepo.On("DeleteComment", ctx, postID, commentID, userID).Return(nil).Once()
    suite.popularityRepo.On("DecrementPopularity", ctx, postID, "comment").Return(nil).Once()

    err := suite.blogUsecase.DeleteComment(ctx, postID, commentID, userID)
    assert.NoError(suite.T(), err)
    suite.commentRepo.AssertExpectations(suite.T())
    suite.popularityRepo.AssertExpectations(suite.T())
}
func (suite *BlogUsecaseTestSuite) TestUpdateComment() {
    ctx := context.TODO()
    postID := primitive.NewObjectID()
    commentID := primitive.NewObjectID()
    userID := primitive.NewObjectID()
    comment := &domain.Comment{Content: "Updated comment"}
    suite.commentRepo.On("UpdateComment", ctx, postID, commentID, userID, comment).Return(nil).Once()

    err := suite.blogUsecase.UpdateComment(ctx, postID, commentID, userID, comment)
    assert.NoError(suite.T(), err)
    suite.commentRepo.AssertExpectations(suite.T())
}
func TestBlogUsecaseTestSuite(t *testing.T) {
    suite.Run(t, new(BlogUsecaseTestSuite))
}