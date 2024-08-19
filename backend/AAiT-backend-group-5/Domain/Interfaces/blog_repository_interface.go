package interfaces

import (
	"context"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type BlogRepository interface {
	CreateBlog(ctx context.Context, blog *models.Blog) *models.ErrorResponse

	GetBlog(ctx context.Context, id string) (*models.Blog, *models.ErrorResponse)
	GetBlogs(ctx context.Context) ([]*models.Blog, *models.ErrorResponse)
	SearchBlogs(ctx context.Context, filter dtos.FilterBlogRequest) ([]*models.Blog, *models.ErrorResponse)

	UpdateBlog(ctx context.Context, blogID string, blog *models.Blog) *models.ErrorResponse
	DeleteBlog(ctx context.Context, id string) *models.ErrorResponse

	AddComment(ctx context.Context, blogID string, comment models.Comment) *models.ErrorResponse
	TrackPopularity(ctx context.Context, blogID string, popularity dtos.TrackPopularityRequest) *models.ErrorResponse
	IncreaseView(ctx context.Context, blogID string) *models.ErrorResponse
}
