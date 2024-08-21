package interfaces

import (
	"context"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"github.com/gin-gonic/gin"
)

type BlogRepository interface {
	CreateBlog(ctx context.Context, blog *models.Blog) (*models.Blog, *models.ErrorResponse)
	GetBlog(ctx context.Context, id string) (*models.Blog, *models.ErrorResponse)
	GetBlogs(ctx context.Context, page int) ([]*models.Blog, *models.ErrorResponse)
	SearchBlogs(ctx context.Context, filter dtos.FilterBlogRequest) ([]*models.Blog, *models.ErrorResponse)
	UpdateBlog(ctx context.Context, blogID string, blog *models.Blog) *models.ErrorResponse
	DeleteBlog(ctx context.Context, id string) *models.ErrorResponse
	AddComment(ctx context.Context, comment models.Comment) *models.ErrorResponse
	GetComments(ctx context.Context, blogID string) ([]models.Comment, *models.ErrorResponse)
	GetPopularity(ctx context.Context, blogID string) (*models.Popularity, *models.ErrorResponse)
	IncreaseView(ctx context.Context, blogID string) *models.ErrorResponse
}

type BlogUsecase interface {
	CreateBlog(ctx context.Context, blog *models.Blog) (*dtos.BlogResponse, *models.ErrorResponse)
	GetBlog(ctx context.Context, id string) (*dtos.BlogResponse, *models.ErrorResponse)
	GetBlogs(ctx context.Context, page int) ([]*dtos.BlogResponse, *models.ErrorResponse)
	SearchBlogs(ctx context.Context, filter dtos.FilterBlogRequest) ([]*dtos.BlogResponse, *models.ErrorResponse)
	UpdateBlog(ctx context.Context, blogID string, blog *models.Blog) *models.ErrorResponse
	DeleteBlog(ctx context.Context, deleteBlogReq dtos.DeleteBlogRequest) *models.ErrorResponse
	TrackPopularity(ctx context.Context, popularity dtos.TrackPopularityRequest) *models.ErrorResponse
	AddComment(ctx context.Context, comment models.Comment) *models.ErrorResponse
}

type BlogController interface {
	CreateBlogController(ctx *gin.Context)
	GetBlogController(ctx *gin.Context)
	GetBlogsController(ctx *gin.Context)
	SearchBlogsController(ctx *gin.Context)
	UpdateBlogController(ctx *gin.Context)
	DeleteBlogController(ctx *gin.Context)
	TrackPopularityController(ctx *gin.Context)
	AddCommentController(ctx *gin.Context)
}

type BlogPopularityActionRepository interface {
	Like(ctx context.Context, popularity dtos.TrackPopularityRequest) *models.ErrorResponse
	Dislike(ctx context.Context, popularity dtos.TrackPopularityRequest) *models.ErrorResponse
	GetBlogPopularityAction(ctx context.Context, blogID string, userID string) (*models.PopularityAction, *models.ErrorResponse)
	UndoLike(ctx context.Context, popularity dtos.TrackPopularityRequest) *models.ErrorResponse
	UndoDislike(ctx context.Context, popularity dtos.TrackPopularityRequest) *models.ErrorResponse
}

type BlogHelper interface {
	CreateSlug(blogTitle string) string
	Marshal(data interface{}) (string, *models.ErrorResponse)
	Unmarshal(dataJSON string, result interface{}) *models.ErrorResponse
	GetBlogs(ctx context.Context, data []*models.Blog) ([]*dtos.BlogResponse, *models.ErrorResponse)
	FetchComments(ctx context.Context, blogID string) ([]models.Comment, *models.ErrorResponse)
	FetchPopularity(ctx context.Context, blogID string) (*models.Popularity, *models.ErrorResponse)
	FetchFromCacheOrRepo(ctx context.Context, cacheKey string, fetchFromRepo func() (interface{}, *models.ErrorResponse)) (interface{}, *models.ErrorResponse)

	FetchFromCacheOrRepoBlogs(ctx context.Context, cacheKey string, fetchFromRepo func() (interface{}, *models.ErrorResponse)) (interface{}, *models.ErrorResponse)
}
