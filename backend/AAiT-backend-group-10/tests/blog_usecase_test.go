package tests

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"aait.backend.g10/domain"
	"aait.backend.g10/tests/mocks"
	"aait.backend.g10/usecases"
	"aait.backend.g10/usecases/dto"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type BlogUseCaseSuite struct {
	suite.Suite
	blogRepo    *mocks.IBlogRepository
	userRepo    *mocks.IUserRepository
	likeRepo    *mocks.LikeRepositoryInterface
	commentRepo *mocks.CommentRepositoryInterface
	aiService   *mocks.IAIService
	cacheRepo   *mocks.CacheRepoInterface
	blogUsecase usecases.BlogUseCase
}

func (suite *BlogUseCaseSuite) SetupTest() {
	suite.blogRepo = new(mocks.IBlogRepository)
	suite.userRepo = new(mocks.IUserRepository)
	suite.likeRepo = new(mocks.LikeRepositoryInterface)
	suite.commentRepo = new(mocks.CommentRepositoryInterface)
	suite.aiService = new(mocks.IAIService)
	suite.cacheRepo = new(mocks.CacheRepoInterface)
	suite.blogUsecase = *usecases.NewBlogUseCase(
		suite.blogRepo,
		suite.userRepo,
		suite.likeRepo,
		suite.commentRepo,
		suite.aiService,
		suite.cacheRepo,
	)
}

func (suite *BlogUseCaseSuite) TearDownTest() {
	suite.userRepo.AssertExpectations(suite.T())
	suite.aiService.AssertExpectations(suite.T())
	suite.cacheRepo.AssertExpectations(suite.T())
	suite.commentRepo.AssertExpectations(suite.T())
	suite.likeRepo.AssertExpectations(suite.T())
	suite.blogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseSuite) TestCreateBlog_Positive() {
	AuthorID := uuid.New()
	blog := &domain.Blog{
		Title:   "Test Blog",
		Content: "This is a test blog",
		Author:  AuthorID,
		Tags:    []string{"some", "thing"},
	}
	suite.blogRepo.On("Create", mock.AnythingOfType("*domain.Blog")).Return(nil)
	suite.userRepo.On("GetUserByID", AuthorID).Return(&domain.User{}, nil)
	suite.cacheRepo.On("Delete", "blogs:all").Return(nil)
	blogDto, err := suite.blogUsecase.CreateBlog(blog)
	suite.Nil(err)
	suite.NotNil(blogDto)
}
func (suite *BlogUseCaseSuite) TestCreateBlog_Negative() {
	AuthorID := uuid.New()
	blog := &domain.Blog{
		Title:   "Test Blog",
		Content: "This is a test blog",
		Author:  AuthorID,
		Tags:    []string{"some", "thing"},
	}
	suite.userRepo.On("GetUserByID", AuthorID).Return(&domain.User{}, nil)
	suite.blogRepo.On("Create", mock.AnythingOfType("*domain.Blog")).Return(domain.ErrBlogInsertFailed)
	blogDto, err := suite.blogUsecase.CreateBlog(blog)
	suite.Equal(err, domain.ErrBlogInsertFailed)
	suite.Nil(blogDto)
}
func (suite *BlogUseCaseSuite) TestCreateBlog_Author_NotFound_Negative() {
	AuthorID := uuid.New()
	blog := &domain.Blog{
		Title:   "Test Blog",
		Content: "This is a test blog",
		Author:  AuthorID,
		Tags:    []string{"some", "thing"},
	}
	suite.userRepo.On("GetUserByID", AuthorID).Return(nil, domain.ErrUserNotFound)
	blogDto, err := suite.blogUsecase.CreateBlog(blog)
	suite.Equal(err, domain.ErrUserNotFound)
	suite.Nil(blogDto)
}
func (suite *BlogUseCaseSuite) TestCreateBlog_CacheFailed_Positive() {
	AuthorID := uuid.New()
	blog := &domain.Blog{
		Title:   "Test Blog",
		Content: "This is a test blog",
		Author:  AuthorID,
		Tags:    []string{"some", "thing"},
	}
	suite.blogRepo.On("Create", mock.AnythingOfType("*domain.Blog")).Return(nil)
	suite.userRepo.On("GetUserByID", AuthorID).Return(&domain.User{}, nil)
	suite.cacheRepo.On("Delete", "blogs:all").Return(domain.ErrCacheDeleteFailed)
	blogDto, err := suite.blogUsecase.CreateBlog(blog)
	suite.Nil(err)
	suite.NotNil(blogDto)
}
func (suite *BlogUseCaseSuite) TestGetAllBlogs_CacheFound_Positive() {
	dummyBlogs, _ :=
		json.Marshal([]*dto.BlogDto{
			{Title: "Test Blog 1", Content: "This is a test blog 1", AuthorName: "someone", Tags: []string{"some", "thing"}, LikeCount: 1, DislikeCount: 2, CommentCount: 3},
			{Title: "Test Blog 2", Content: "This is a test blog 2", AuthorName: "someone", Tags: []string{"some", "thing"}, LikeCount: 1, DislikeCount: 2, CommentCount: 3},
		})
	suite.cacheRepo.On("Get", "blogs:all").Return(string(dummyBlogs), nil)
	blogDtos, err := suite.blogUsecase.GetAllBlogs()
	suite.Nil(err)
	suite.NotNil(blogDtos)
}

func (suite *BlogUseCaseSuite) TestGetAllBlogs_CacheMiss_Positive() {
	originalFunc := usecases.GetLikeAndCommentCount
	defer func() { usecases.GetLikeAndCommentCount = originalFunc }()

	usecases.GetLikeAndCommentCount = func(b *usecases.BlogUseCase, id uuid.UUID) (int, int, int, *domain.CustomError) {
		return 10, 5, 3, nil
	}
	dummyBlogs := []domain.Blog{
		{Author: uuid.New()},
		{Author: uuid.New()},
	}
	suite.cacheRepo.On("Get", "blogs:all").Return("", domain.ErrCacheNotFound)
	suite.blogRepo.On("FindAll").Return(dummyBlogs, nil)
	suite.userRepo.On("GetUserByID", mock.AnythingOfType("uuid.UUID")).Return(&domain.User{FullName: "Someone's Fullname"}, nil)
	suite.cacheRepo.On("Set", "blogs:all", mock.AnythingOfType("string"), mock.AnythingOfType("time.Duration")).Return(nil)

	blogDtos, err := suite.blogUsecase.GetAllBlogs()
	suite.Nil(err)
	suite.NotNil(blogDtos)
}

func (suite *BlogUseCaseSuite) TestGetAllBlogs_BlogNotFound_Negative() {
	originalFunc := usecases.GetLikeAndCommentCount
	defer func() { usecases.GetLikeAndCommentCount = originalFunc }()

	usecases.GetLikeAndCommentCount = func(b *usecases.BlogUseCase, id uuid.UUID) (int, int, int, *domain.CustomError) {
		return 10, 5, 3, nil
	}
	suite.cacheRepo.On("Get", "blogs:all").Return("", domain.ErrCacheNotFound)
	suite.blogRepo.On("FindAll").Return(nil, domain.ErrBlogNotFound)

	blogDtos, err := suite.blogUsecase.GetAllBlogs()
	suite.Nil(blogDtos)
	suite.Equal(err, domain.ErrBlogNotFound)
}

func (suite *BlogUseCaseSuite) TestGetAllBlogs_AuthorNotFound_Negative() {
	originalFunc := usecases.GetLikeAndCommentCount
	defer func() { usecases.GetLikeAndCommentCount = originalFunc }()

	usecases.GetLikeAndCommentCount = func(b *usecases.BlogUseCase, id uuid.UUID) (int, int, int, *domain.CustomError) {
		return 10, 5, 3, nil
	}
	dummyBlogs := []domain.Blog{
		{Author: uuid.New()},
		{Author: uuid.New()},
	}
	suite.cacheRepo.On("Get", "blogs:all").Return("", domain.ErrCacheNotFound)
	suite.blogRepo.On("FindAll").Return(dummyBlogs, nil)
	suite.userRepo.On("GetUserByID", mock.AnythingOfType("uuid.UUID")).Return(nil, domain.ErrUserNotFound)

	blogDtos, err := suite.blogUsecase.GetAllBlogs()
	suite.Nil(blogDtos)
	suite.Equal(err, domain.ErrUserNotFound)
}

func (suite *BlogUseCaseSuite) TestGetAllBlogs_LikeAndCommentNotFound_Negative() {
	originalFunc := usecases.GetLikeAndCommentCount
	defer func() { usecases.GetLikeAndCommentCount = originalFunc }()

	usecases.GetLikeAndCommentCount = func(b *usecases.BlogUseCase, id uuid.UUID) (int, int, int, *domain.CustomError) {
		return 0, 0, 0, domain.ErrLikeCountFetchFailed
	}
	dummyBlogs := []domain.Blog{
		{Author: uuid.New()},
		{Author: uuid.New()},
	}
	suite.cacheRepo.On("Get", "blogs:all").Return("", domain.ErrCacheNotFound)
	suite.blogRepo.On("FindAll").Return(dummyBlogs, nil)
	suite.userRepo.On("GetUserByID", mock.AnythingOfType("uuid.UUID")).Return(&domain.User{FullName: "Someone's Fullname"}, nil)

	blogDtos, err := suite.blogUsecase.GetAllBlogs()
	suite.Nil(blogDtos)
	suite.Equal(err, domain.ErrLikeCountFetchFailed)
}

func (suite *BlogUseCaseSuite) TestGetAllBlogs_CacheSetFailed_Positive() {
	originalFunc := usecases.GetLikeAndCommentCount
	defer func() { usecases.GetLikeAndCommentCount = originalFunc }()

	usecases.GetLikeAndCommentCount = func(b *usecases.BlogUseCase, id uuid.UUID) (int, int, int, *domain.CustomError) {
		return 10, 5, 3, nil
	}
	dummyBlogs := []domain.Blog{
		{Author: uuid.New()},
		{Author: uuid.New()},
	}
	suite.cacheRepo.On("Get", "blogs:all").Return("", domain.ErrCacheNotFound)
	suite.blogRepo.On("FindAll").Return(dummyBlogs, nil)
	suite.userRepo.On("GetUserByID", mock.AnythingOfType("uuid.UUID")).Return(&domain.User{FullName: "Someone's Fullname"}, nil)
	suite.cacheRepo.On("Set", "blogs:all", mock.AnythingOfType("string"), mock.AnythingOfType("time.Duration")).Return(domain.ErrCacheSetFailed)

	blogDtos, err := suite.blogUsecase.GetAllBlogs()
	suite.Nil(err)
	suite.NotNil(blogDtos)
}

func (suite *BlogUseCaseSuite) TestGetBlogByID_CacheFound_Positive() {
	originalFunc := usecases.GetLikeAndCommentCount
	defer func() { usecases.GetLikeAndCommentCount = originalFunc }()

	usecases.GetLikeAndCommentCount = func(b *usecases.BlogUseCase, id uuid.UUID) (int, int, int, *domain.CustomError) {
		return 10, 5, 3, nil
	}
	dummyBlog := &dto.BlogDto{
		ID:           uuid.New(),
		LikeCount:    1,
		DislikeCount: 2,
		CommentCount: 3,
	}
	dummyBlogJson, _ := json.Marshal(dummyBlog)
	suite.cacheRepo.On("Get", "blog:"+dummyBlog.ID.String()).Return(string(dummyBlogJson), nil)
	blogDto, err := suite.blogUsecase.GetBlogByID(dummyBlog.ID)
	suite.Nil(err)
	suite.NotNil(blogDto)
}

func (suite *BlogUseCaseSuite) TestGetBlogByID_CacheMiss_Positive() {
	originalFunc := usecases.GetLikeAndCommentCount
	defer func() { usecases.GetLikeAndCommentCount = originalFunc }()

	usecases.GetLikeAndCommentCount = func(b *usecases.BlogUseCase, id uuid.UUID) (int, int, int, *domain.CustomError) {
		return 10, 5, 3, nil
	}
	dummyBlog := domain.Blog{
		ID:     uuid.New(),
		Author: uuid.New(),
	}
	suite.cacheRepo.On("Get", "blog:"+dummyBlog.ID.String()).Return("", domain.ErrCacheNotFound)
	suite.blogRepo.On("FindByID", dummyBlog.ID).Return(&dummyBlog, nil)
	suite.userRepo.On("GetUserByID", dummyBlog.Author).Return(&domain.User{FullName: "Someone's Fullname"}, nil)
	suite.cacheRepo.On("Set", "blog:"+dummyBlog.ID.String(), mock.AnythingOfType("string"), mock.AnythingOfType("time.Duration")).Return(nil)
	blogDto, err := suite.blogUsecase.GetBlogByID(dummyBlog.ID)
	suite.Nil(err)
	suite.NotNil(blogDto)
}
func (suite *BlogUseCaseSuite) TestGetBlogByID_CacheMiss_BlogNotFound_Negative() {
	originalFunc := usecases.GetLikeAndCommentCount
	defer func() { usecases.GetLikeAndCommentCount = originalFunc }()

	usecases.GetLikeAndCommentCount = func(b *usecases.BlogUseCase, id uuid.UUID) (int, int, int, *domain.CustomError) {
		return 10, 5, 3, nil
	}
	dummyBlog := domain.Blog{
		ID: uuid.New(),
	}
	suite.cacheRepo.On("Get", "blog:"+dummyBlog.ID.String()).Return("", domain.ErrCacheNotFound)
	suite.blogRepo.On("FindByID", dummyBlog.ID).Return(nil, domain.ErrBlogNotFound)

	blogDto, err := suite.blogUsecase.GetBlogByID(dummyBlog.ID)
	suite.Nil(blogDto)
	suite.NotNil(err, domain.ErrBlogNotFound)
}
func (suite *BlogUseCaseSuite) TestGetBlogByID_AuthorNotFound_Negative() {
	originalFunc := usecases.GetLikeAndCommentCount
	defer func() { usecases.GetLikeAndCommentCount = originalFunc }()

	usecases.GetLikeAndCommentCount = func(b *usecases.BlogUseCase, id uuid.UUID) (int, int, int, *domain.CustomError) {
		return 10, 5, 3, nil
	}
	dummyBlog := domain.Blog{
		ID:     uuid.New(),
		Author: uuid.New(),
	}
	suite.cacheRepo.On("Get", "blog:"+dummyBlog.ID.String()).Return("", domain.ErrCacheNotFound)
	suite.blogRepo.On("FindByID", dummyBlog.ID).Return(&dummyBlog, nil)
	suite.userRepo.On("GetUserByID", dummyBlog.Author).Return(nil, domain.ErrUserNotFound)

	blogDto, err := suite.blogUsecase.GetBlogByID(dummyBlog.ID)
	suite.Nil(blogDto)
	suite.NotNil(err, domain.ErrUserNotFound)
}

func (suite *BlogUseCaseSuite) TestGetBlogByID_LikeAndCommentFetchFailed_Negative() {
	originalFunc := usecases.GetLikeAndCommentCount
	defer func() { usecases.GetLikeAndCommentCount = originalFunc }()

	usecases.GetLikeAndCommentCount = func(b *usecases.BlogUseCase, id uuid.UUID) (int, int, int, *domain.CustomError) {
		return 0, 0, 0, domain.ErrLikeCountFetchFailed
	}
	dummyBlog := domain.Blog{
		ID:     uuid.New(),
		Author: uuid.New(),
	}
	suite.cacheRepo.On("Get", "blog:"+dummyBlog.ID.String()).Return("", domain.ErrCacheNotFound)
	suite.blogRepo.On("FindByID", dummyBlog.ID).Return(&dummyBlog, nil)
	suite.userRepo.On("GetUserByID", dummyBlog.Author).Return(&domain.User{FullName: "Someone's Fullname"}, nil)

	blogDto, err := suite.blogUsecase.GetBlogByID(dummyBlog.ID)
	suite.Nil(blogDto)
	suite.Equal(err, domain.ErrLikeCountFetchFailed)
}
func (suite *BlogUseCaseSuite) TestGetBlogByID_CacheSetFailed_Positive() {
	originalFunc := usecases.GetLikeAndCommentCount
	defer func() { usecases.GetLikeAndCommentCount = originalFunc }()

	usecases.GetLikeAndCommentCount = func(b *usecases.BlogUseCase, id uuid.UUID) (int, int, int, *domain.CustomError) {
		return 10, 5, 3, nil
	}
	dummyBlog := domain.Blog{
		ID:     uuid.New(),
		Author: uuid.New(),
	}
	suite.cacheRepo.On("Get", "blog:"+dummyBlog.ID.String()).Return("", domain.ErrCacheNotFound)
	suite.blogRepo.On("FindByID", dummyBlog.ID).Return(&dummyBlog, nil)
	suite.userRepo.On("GetUserByID", dummyBlog.Author).Return(&domain.User{FullName: "Someone's Fullname"}, nil)
	suite.cacheRepo.On("Set", "blog:"+dummyBlog.ID.String(), mock.AnythingOfType("string"), mock.AnythingOfType("time.Duration")).Return(domain.ErrCacheSetFailed)
	blogDto, err := suite.blogUsecase.GetBlogByID(dummyBlog.ID)
	suite.Nil(err)
	suite.NotNil(blogDto)
}

func (suite *BlogUseCaseSuite) TestUpdateBlog_Positive() {
	blog := &domain.Blog{
		ID:        uuid.New(),
		Author:    uuid.New(),
		UpdatedAt: time.Now(),
	}
	suite.blogRepo.On("FindByID", blog.ID).Return(blog, nil)
	suite.blogRepo.On("Update", blog).Return(nil)
	suite.cacheRepo.On("Delete", "blogs:all").Return(nil)
	suite.cacheRepo.On("Delete", "blog:"+blog.ID.String()).Return(nil)
	err := suite.blogUsecase.UpdateBlog(blog)
	suite.Nil(err)
}

func (suite *BlogUseCaseSuite) TestUpdateBlog_BlogNotFound_Negative() {
	blog := &domain.Blog{}
	suite.blogRepo.On("FindByID", mock.AnythingOfType("uuid.UUID")).Return(nil, domain.ErrBlogNotFound)
	err := suite.blogUsecase.UpdateBlog(blog)
	suite.Equal(err, domain.ErrBlogNotFound)
}
func (suite *BlogUseCaseSuite) TestUpdateBlog_UnAuthorized_Negative() {
	BlogID := uuid.New()
	blog := &domain.Blog{
		ID:     BlogID,
		Author: uuid.New(),
	}
	updatedBlog := &domain.Blog{
		ID:     BlogID,
		Author: uuid.New(),
	}
	suite.blogRepo.On("FindByID", updatedBlog.ID).Return(blog, nil)
	err := suite.blogUsecase.UpdateBlog(updatedBlog)
	suite.Equal(err, domain.ErrUnAuthorized)
}
func (suite *BlogUseCaseSuite) TestUpdateBlog_UpdateFailed_Negative() {
	blog := &domain.Blog{
		ID:        uuid.New(),
		Author:    uuid.New(),
		UpdatedAt: time.Now(),
	}
	suite.blogRepo.On("FindByID", blog.ID).Return(blog, nil)
	suite.blogRepo.On("Update", blog).Return(domain.ErrBlogUpdateFailed)
	err := suite.blogUsecase.UpdateBlog(blog)
	suite.Equal(err, domain.ErrBlogUpdateFailed)
}

func (suite *BlogUseCaseSuite) TestUpdateBlog_CacheFailed_Positive() {
	blog := &domain.Blog{
		ID:        uuid.New(),
		Author:    uuid.New(),
		UpdatedAt: time.Now(),
	}
	suite.blogRepo.On("FindByID", blog.ID).Return(blog, nil)
	suite.blogRepo.On("Update", blog).Return(nil)
	suite.cacheRepo.On("Delete", "blogs:all").Return(domain.ErrCacheDeleteFailed)
	suite.cacheRepo.On("Delete", "blog:"+blog.ID.String()).Return(domain.ErrCacheSetFailed)
	err := suite.blogUsecase.UpdateBlog(blog)
	suite.Nil(err)
}

func (suite *BlogUseCaseSuite) TestDeleteBlog_Positive() {
	blog := &domain.Blog{
		ID:     uuid.New(),
		Author: uuid.New(),
	}
	suite.blogRepo.On("FindByID", blog.ID).Return(blog, nil)
	suite.blogRepo.On("Delete", blog.ID).Return(nil)
	suite.commentRepo.On("DeleteCommentsByBlog", blog.ID).Return(nil)
	suite.likeRepo.On("DeleteLikesByBlog", blog.ID).Return(nil)
	suite.cacheRepo.On("Delete", "blogs:all").Return(nil)
	suite.cacheRepo.On("Delete", "blog:"+blog.ID.String()).Return(nil)

	err := suite.blogUsecase.DeleteBlog(blog.ID, blog.Author, false)
	suite.Nil(err)
}

func (suite *BlogUseCaseSuite) TestDeleteBlog_ByAdmin_Positive() {
	blog := &domain.Blog{
		ID: uuid.New(),
	}
	suite.blogRepo.On("FindByID", blog.ID).Return(blog, nil)
	suite.blogRepo.On("Delete", blog.ID).Return(nil)
	suite.commentRepo.On("DeleteCommentsByBlog", blog.ID).Return(nil)
	suite.likeRepo.On("DeleteLikesByBlog", blog.ID).Return(nil)
	suite.cacheRepo.On("Delete", "blogs:all").Return(nil)
	suite.cacheRepo.On("Delete", "blog:"+blog.ID.String()).Return(nil)

	err := suite.blogUsecase.DeleteBlog(blog.ID, blog.Author, true)
	suite.Nil(err)
}

func (suite *BlogUseCaseSuite) TestDeleteBlog_Blog_NotFound_Negative() {
	suite.blogRepo.On("FindByID", mock.AnythingOfType("uuid.UUID")).Return(nil, domain.ErrBlogNotFound)
	err := suite.blogUsecase.DeleteBlog(uuid.New(), uuid.New(), false)
	suite.Equal(err, domain.ErrBlogNotFound)
}

func (suite *BlogUseCaseSuite) TestDeleteBlog__UnAuthorized_Negative() {
	blog := &domain.Blog{
		ID:     uuid.New(),
		Author: uuid.New(),
	}
	suite.blogRepo.On("FindByID", blog.ID).Return(blog, nil)

	err := suite.blogUsecase.DeleteBlog(blog.ID, uuid.New(), false)
	suite.Equal(err, domain.ErrUnAuthorized)
}

func (suite *BlogUseCaseSuite) TestDeleteBlog_CommentDeletionFailed_Negative() {
	blog := &domain.Blog{
		ID: uuid.New(),
	}
	suite.blogRepo.On("FindByID", blog.ID).Return(blog, nil)
	suite.commentRepo.On("DeleteCommentsByBlog", blog.ID).Return(domain.ErrCommentFetchFailed)
	err := suite.blogUsecase.DeleteBlog(blog.ID, uuid.New(), true)
	suite.Equal(err, domain.ErrCommentFetchFailed)
}
func (suite *BlogUseCaseSuite) TestDeleteBlog_LikeDeletionFailed_Negative() {
	blog := &domain.Blog{
		ID: uuid.New(),
	}
	suite.blogRepo.On("FindByID", blog.ID).Return(blog, nil)
	suite.commentRepo.On("DeleteCommentsByBlog", blog.ID).Return(nil)
	suite.likeRepo.On("DeleteLikesByBlog", blog.ID).Return(domain.ErrLikeCountFetchFailed)
	err := suite.blogUsecase.DeleteBlog(blog.ID, uuid.New(), true)
	suite.Equal(err, domain.ErrLikeCountFetchFailed)
}
func (suite *BlogUseCaseSuite) TestDeleteBlog_CacheFailed_Positive() {
	blog := &domain.Blog{
		ID: uuid.New(),
	}
	suite.blogRepo.On("FindByID", blog.ID).Return(blog, nil)
	suite.blogRepo.On("Delete", blog.ID).Return(nil)
	suite.commentRepo.On("DeleteCommentsByBlog", blog.ID).Return(nil)
	suite.likeRepo.On("DeleteLikesByBlog", blog.ID).Return(nil)
	suite.cacheRepo.On("Delete", "blogs:all").Return(domain.ErrCacheSetFailed)
	suite.cacheRepo.On("Delete", "blog:"+blog.ID.String()).Return(domain.ErrCacheSetFailed)

	err := suite.blogUsecase.DeleteBlog(blog.ID, blog.Author, true)
	suite.Nil(err)
}
func (suite *BlogUseCaseSuite) TestDeleteBlog_UpdateFailed_Negative() {
	blog := &domain.Blog{
		ID: uuid.New(),
	}
	suite.blogRepo.On("FindByID", blog.ID).Return(blog, nil)
	suite.blogRepo.On("Delete", blog.ID).Return(domain.ErrBlogUpdateFailed)
	suite.commentRepo.On("DeleteCommentsByBlog", blog.ID).Return(nil)
	suite.likeRepo.On("DeleteLikesByBlog", blog.ID).Return(nil)
	suite.cacheRepo.On("Delete", "blogs:all").Return(nil)
	suite.cacheRepo.On("Delete", "blog:"+blog.ID.String()).Return(nil)

	err := suite.blogUsecase.DeleteBlog(blog.ID, blog.Author, true)
	suite.Equal(err, domain.ErrBlogUpdateFailed)
}

func (suite *BlogUseCaseSuite) TestAddView_Positive() {
	id := uuid.New()
	suite.blogRepo.On("AddView", id).Return(nil)
	err := suite.blogUsecase.AddView(id)
	suite.Nil(err)

}
func (suite *BlogUseCaseSuite) TestAddView_Negative() {
	id := uuid.New()
	suite.blogRepo.On("AddView", id).Return(domain.NewCustomError("BlogViewFailed", 404))
	err := suite.blogUsecase.AddView(id)
	suite.Equal(err, domain.NewCustomError("BlogViewFailed", 404))

}

func (suite *BlogUseCaseSuite) TestSearchBlogs_CacheFound_Positive() {
	filter := domain.BlogFilter{
		Title:    "some title",
		Author:   "someone",
		Tags:     []string{"some", "thing"},
		SortBy:   "some",
		Page:     1,
		PageSize: 10,
	}
	dummyBlogs, _ := json.Marshal([]*dto.BlogDto{
		{Title: "Test Blog 1", Content: "This is a test blog 1", AuthorName: "someone", Tags: []string{"some", "thing"}, LikeCount: 1, DislikeCount: 2, CommentCount: 3},
		{Title: "Test Blog 2", Content: "This is a test blog 2", AuthorName: "someone", Tags: []string{"some", "thing"}, LikeCount: 1, DislikeCount: 2, CommentCount: 3},
	})
	cacheKey := fmt.Sprintf("blogs:search:%s:%s:%s:%s:%d:%d", filter.Title, filter.Author, filter.SortBy, filter.Tags, filter.Page, filter.PageSize)
	suite.cacheRepo.On("Get", cacheKey).Return(string(dummyBlogs), nil)
	blogDtos, total, pages, err := suite.blogUsecase.SearchBlogs(filter)
	suite.Nil(err)
	suite.NotNil(blogDtos)
	suite.NotNil(total)
	suite.NotNil(pages)
}
func (suite *BlogUseCaseSuite) TestSearchBlogs_CacheMiss_Positive() {
	originalFunc := usecases.GetLikeAndCommentCount
	defer func() { usecases.GetLikeAndCommentCount = originalFunc }()

	usecases.GetLikeAndCommentCount = func(b *usecases.BlogUseCase, id uuid.UUID) (int, int, int, *domain.CustomError) {
		return 10, 5, 3, nil
	}
	authorID1 := uuid.New()
	filter := domain.BlogFilter{
		Title:     "some title",
		Author:    "someone",
		Tags:      []string{"some", "thing"},
		Page:      1,
		SortBy:    "recent",
		PageSize:  10,
		AuthorIds: []uuid.UUID{authorID1},
	}
	dummyBlogs := []domain.Blog{
		{Author: authorID1},
	}
	cacheKey := fmt.Sprintf("blogs:search:%s:%s:%s:%s:%d:%d", filter.Title, filter.Author, filter.SortBy, filter.Tags, filter.Page, filter.PageSize)
	suite.cacheRepo.On("Get", cacheKey).Return("", domain.ErrCacheNotFound)
	suite.userRepo.On("GetAllUsersWithName", filter.Author).Return([]uuid.UUID{authorID1}, nil)
	suite.blogRepo.On("Search", filter).Return(dummyBlogs, 1, nil)
	suite.userRepo.On("GetUserByID", filter.AuthorIds[0]).Return(&domain.User{}, nil)
	suite.cacheRepo.On("Set", cacheKey, mock.AnythingOfType("string"), mock.AnythingOfType("time.Duration")).Return(nil)

	blogDtos, total, pages, err := suite.blogUsecase.SearchBlogs(filter)
	suite.Nil(err)
	suite.NotNil(blogDtos)
	suite.NotNil(total)
	suite.NotNil(pages)

}
func (suite *BlogUseCaseSuite) TestSearchBlogs_NameNotFound_Negative() {
	originalFunc := usecases.GetLikeAndCommentCount
	defer func() { usecases.GetLikeAndCommentCount = originalFunc }()

	usecases.GetLikeAndCommentCount = func(b *usecases.BlogUseCase, id uuid.UUID) (int, int, int, *domain.CustomError) {
		return 10, 5, 3, nil
	}
	authorID1 := uuid.New()
	filter := domain.BlogFilter{
		Title:     "some title",
		Author:    "someone",
		Tags:      []string{"some", "thing"},
		Page:      1,
		SortBy:    "recent",
		PageSize:  10,
		AuthorIds: []uuid.UUID{authorID1},
	}
	cacheKey := fmt.Sprintf("blogs:search:%s:%s:%s:%s:%d:%d", filter.Title, filter.Author, filter.SortBy, filter.Tags, filter.Page, filter.PageSize)
	suite.cacheRepo.On("Get", cacheKey).Return("", domain.ErrCacheNotFound)
	suite.userRepo.On("GetAllUsersWithName", filter.Author).Return(nil, domain.ErrUserNotFound)

	blogDtos, total, pages, err := suite.blogUsecase.SearchBlogs(filter)
	suite.Equal(err, domain.ErrUserNotFound)
	suite.Nil(blogDtos)
	suite.Equal(total, 0)
	suite.Equal(pages, 0)

}
func (suite *BlogUseCaseSuite) TestSearchBlogs_SearchFailed_Negative() {
	originalFunc := usecases.GetLikeAndCommentCount
	defer func() { usecases.GetLikeAndCommentCount = originalFunc }()

	usecases.GetLikeAndCommentCount = func(b *usecases.BlogUseCase, id uuid.UUID) (int, int, int, *domain.CustomError) {
		return 10, 5, 3, nil
	}
	authorID1 := uuid.New()
	filter := domain.BlogFilter{
		Title:     "some title",
		Author:    "someone",
		Tags:      []string{"some", "thing"},
		Page:      1,
		SortBy:    "recent",
		PageSize:  10,
		AuthorIds: []uuid.UUID{authorID1},
	}
	cacheKey := fmt.Sprintf("blogs:search:%s:%s:%s:%s:%d:%d", filter.Title, filter.Author, filter.SortBy, filter.Tags, filter.Page, filter.PageSize)
	suite.cacheRepo.On("Get", cacheKey).Return("", domain.ErrCacheNotFound)
	suite.userRepo.On("GetAllUsersWithName", filter.Author).Return([]uuid.UUID{authorID1}, nil)
	suite.blogRepo.On("Search", filter).Return(nil, 0, domain.ErrBlogNotFound)

	blogDtos, total, pages, err := suite.blogUsecase.SearchBlogs(filter)
	suite.Nil(blogDtos)
	suite.Equal(err, domain.ErrBlogNotFound)
	suite.Equal(total, 0)
	suite.Equal(pages, 0)

}

func (suite *BlogUseCaseSuite) TestSearchBlogs_UserID_NotFound_Negative() {
	originalFunc := usecases.GetLikeAndCommentCount
	defer func() { usecases.GetLikeAndCommentCount = originalFunc }()

	usecases.GetLikeAndCommentCount = func(b *usecases.BlogUseCase, id uuid.UUID) (int, int, int, *domain.CustomError) {
		return 10, 5, 3, nil
	}
	authorID1 := uuid.New()
	filter := domain.BlogFilter{
		Title:     "some title",
		Author:    "someone",
		Tags:      []string{"some", "thing"},
		Page:      1,
		SortBy:    "recent",
		PageSize:  10,
		AuthorIds: []uuid.UUID{authorID1},
	}
	dummyBlogs := []domain.Blog{
		{Author: authorID1},
	}
	cacheKey := fmt.Sprintf("blogs:search:%s:%s:%s:%s:%d:%d", filter.Title, filter.Author, filter.SortBy, filter.Tags, filter.Page, filter.PageSize)
	suite.cacheRepo.On("Get", cacheKey).Return("", domain.ErrCacheNotFound)
	suite.userRepo.On("GetAllUsersWithName", filter.Author).Return([]uuid.UUID{authorID1}, nil)
	suite.blogRepo.On("Search", filter).Return(dummyBlogs, 1, nil)
	suite.userRepo.On("GetUserByID", filter.AuthorIds[0]).Return(nil, domain.ErrUserNotFound)


	blogDtos, total, pages, err := suite.blogUsecase.SearchBlogs(filter)
	suite.Equal(err, domain.ErrUserNotFound)
	suite.Nil(blogDtos)
	suite.Equal(total, 0)
	suite.Equal(pages, 0)

}
func (suite *BlogUseCaseSuite) TestSearchBlogs_CacheSetFailed_Positive() {
	originalFunc := usecases.GetLikeAndCommentCount
	defer func() { usecases.GetLikeAndCommentCount = originalFunc }()

	usecases.GetLikeAndCommentCount = func(b *usecases.BlogUseCase, id uuid.UUID) (int, int, int, *domain.CustomError) {
		return 10, 5, 3, nil
	}
	authorID1 := uuid.New()
	filter := domain.BlogFilter{
		Title:     "some title",
		Author:    "someone",
		Tags:      []string{"some", "thing"},
		Page:      1,
		SortBy:    "recent",
		PageSize:  10,
		AuthorIds: []uuid.UUID{authorID1},
	}
	dummyBlogs := []domain.Blog{
		{Author: authorID1},
	}
	cacheKey := fmt.Sprintf("blogs:search:%s:%s:%s:%s:%d:%d", filter.Title, filter.Author, filter.SortBy, filter.Tags, filter.Page, filter.PageSize)
	suite.cacheRepo.On("Get", cacheKey).Return("", domain.ErrCacheNotFound)
	suite.userRepo.On("GetAllUsersWithName", filter.Author).Return([]uuid.UUID{authorID1}, nil)
	suite.blogRepo.On("Search", filter).Return(dummyBlogs, 1, nil)
	suite.userRepo.On("GetUserByID", filter.AuthorIds[0]).Return(&domain.User{}, nil)
	suite.cacheRepo.On("Set", cacheKey, mock.AnythingOfType("string"), mock.AnythingOfType("time.Duration")).Return(domain.ErrCacheSetFailed)

	blogDtos, total, pages, err := suite.blogUsecase.SearchBlogs(filter)
	suite.Nil(err)
	suite.NotNil(blogDtos)
	suite.NotNil(total)
	suite.NotNil(pages)

}

func(suite *BlogUseCaseSuite) TestGenerateBlogContent_Positive() {
	topic := "Some topic"
	keywords := []string{"key", "words"}
	generated_content := &domain.BlogContentResponse{
		SuggestedContent: "Some generated content",
	}
	suite.aiService.On("GenerateContent", topic, keywords).Return(generated_content, nil)
	content, err := suite.aiService.GenerateContent(topic, keywords)
	suite.Nil(err)
	suite.Equal(content, generated_content)
}

func(suite *BlogUseCaseSuite) TestGenerateBlogContent_Negative() {
	topic := "Some topic"
	keywords := []string{"key", "words"}
	suite.aiService.On("GenerateContent", topic, keywords).Return(nil, domain.NewCustomError("AIError", 404))
	content, err := suite.aiService.GenerateContent(topic, keywords)
	suite.Nil(content)
	suite.Equal(err, domain.NewCustomError("AIError", 404))
}

func(suite *BlogUseCaseSuite) TestSuggestImprovements_Positive() {
	content := "Some content"
	suggestion_response := &domain.SuggestionResponse{
		Suggestions: "Some suggestion",
	}
	suite.aiService.On("SuggestImprovements", content).Return(suggestion_response, nil)
	response, err := suite.aiService.SuggestImprovements(content)
	suite.Nil(err)
	suite.Equal(suggestion_response, response)
}
func(suite *BlogUseCaseSuite) TestSuggestImprovements_Negative() {
	content := "Some content"
	suite.aiService.On("SuggestImprovements", content).Return(nil, domain.NewCustomError("AIError", 404))
	response, err := suite.aiService.SuggestImprovements(content)
	suite.Nil(response)
	suite.Equal(err, domain.NewCustomError("AIError", 404))
}


func TestBlogCaseSuite(t *testing.T) {
	suite.Run(t, new(BlogUseCaseSuite))
}
