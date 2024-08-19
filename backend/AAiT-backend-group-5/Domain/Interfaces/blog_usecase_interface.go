package interfaces

import (
	"context"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type BlogUsecase interface {
	CreateBlog(ctx context.Context, blog *models.Blog) *models.ErrorResponse
	GetBlog(ctx context.Context, id string) (*models.Blog, *models.ErrorResponse)
	GetBlogs(ctx context.Context) ([]*models.Blog, *models.ErrorResponse)
	SearchBlogs(ctx context.Context, filter dtos.FilterBlogRequest) ([]*models.Blog, *models.ErrorResponse)
	UpdateBlog(ctx context.Context, blogID string,blog *models.Blog) *models.ErrorResponse
	DeleteBlog(ctx context.Context, id string) *models.ErrorResponse
	TrackPopularity(ctx context.Context, blogID string, popularity dtos.TrackPopularityRequest) *models.ErrorResponse
 }
 