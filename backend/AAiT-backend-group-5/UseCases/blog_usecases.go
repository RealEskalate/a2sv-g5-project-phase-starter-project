package usecases

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"

	config "github.com/aait.backend.g5.main/backend/Config"
	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type blogUsecase struct {
	repository interfaces.BlogRepository
	redisCache interfaces.RedisCache
	env        config.Env
	cacheTTL   time.Duration
	helper     interfaces.BlogHelper
}

func NewblogUsecase(
	repository interfaces.BlogRepository,
	redisCache interfaces.RedisCache,
	env config.Env,
	cacheTTL time.Duration,
	helper interfaces.BlogHelper) interfaces.BlogUsecase {

	return &blogUsecase{
		repository: repository,
		redisCache: redisCache,
		env:        env,
		cacheTTL:   cacheTTL,
		helper:     helper,
	}
}

// Common cache handler for fetching data with caching logic
func (b *blogUsecase) fetchFromCacheOrRepo(ctx context.Context, cacheKey string, fetchFromRepo func() (interface{}, *models.ErrorResponse)) (interface{}, *models.ErrorResponse) {
	cachedData, err := b.redisCache.Get(ctx, cacheKey)

	if err == redis.Nil {
		data, repoErr := fetchFromRepo()
		if repoErr != nil {
			return nil, repoErr
		}

		dataJSON, _ := b.helper.Marshal(data)
		b.redisCache.Set(ctx, cacheKey, string(dataJSON), b.cacheTTL)

		return data, nil
	} else if err != nil {
		return nil, models.InternalServerError("Error while fetching data from cache")
	}

	var result interface{}
	if err := b.helper.Unmarshal(cachedData, &result); err != nil {
		return nil, models.InternalServerError("Error while unmarshalling data from cache")
	}

	return result, nil
}

func (b *blogUsecase) CreateBlog(ctx context.Context, blog *models.Blog) *models.ErrorResponse {
	slug := b.helper.CreateSlug(blog.Title)
	blog.Slug = slug
	return b.repository.CreateBlog(ctx, blog)
}

func (b *blogUsecase) GetBlog(ctx context.Context, id string) (*models.Blog, *models.ErrorResponse) {
	data, err := b.fetchFromCacheOrRepo(ctx, id, func() (interface{}, *models.ErrorResponse) {
		return b.repository.GetBlog(ctx, id)
	})

	if err != nil {
		return nil, err
	}

	return data.(*models.Blog), nil
}

func (b *blogUsecase) GetBlogs(ctx context.Context) ([]*models.Blog, *models.ErrorResponse) {
	data, err := b.fetchFromCacheOrRepo(ctx, b.env.REDIS_BLOG_KEY, func() (interface{}, *models.ErrorResponse) {
		return b.repository.GetBlogs(ctx)
	})

	if err != nil {
		return nil, err
	}

	return data.([]*models.Blog), nil
}

func (b *blogUsecase) SearchBlogs(ctx context.Context, filter dtos.FilterBlogRequest) ([]*models.Blog, *models.ErrorResponse) {
	searchJson, err := json.Marshal(filter)

	if err != nil {
		return nil, models.InternalServerError("Error while marshalling search filter")
	}

	cachedBlogs, sErr := b.redisCache.Get(ctx, string(searchJson))

	if sErr == redis.Nil {
		blogs, nErr := b.repository.SearchBlogs(ctx, filter)

		if nErr != nil {
			return nil, nErr
		}

		blogsJson, _ := b.helper.Marshal(blogs)
		b.redisCache.Set(ctx, string(searchJson), string(blogsJson), b.cacheTTL)

		return blogs, nil
	} else if sErr != nil {
		return nil, models.InternalServerError("Error while fetching search query from cache")
	}

	var blogs []*models.Blog
	if err := b.helper.Unmarshal(cachedBlogs, &blogs); err != nil {
		return nil, models.InternalServerError("Error while unmarshalling search query from cache")
	}

	return blogs, nil
}

func (b *blogUsecase) UpdateBlog(ctx context.Context, blogID string, blog *models.Blog) *models.ErrorResponse {

	if err := b.repository.UpdateBlog(ctx, blogID, blog); err != nil {
		return err
	}

	if err := b.redisCache.Delete(ctx, blogID); err != nil {
		return models.InternalServerError("Error while deleting blog from cache")
	}

	updatedBlog, _ := b.repository.GetBlog(ctx, blogID)
	dataJSON, _ := b.helper.Marshal(updatedBlog)
	b.redisCache.Set(ctx, blogID, string(dataJSON), b.cacheTTL)

	return nil
}

func (b *blogUsecase) DeleteBlog(ctx context.Context, id string) *models.ErrorResponse {
	if err := b.repository.DeleteBlog(ctx, id); err != nil {
		return err
	}

	if err := b.redisCache.Delete(ctx, b.env.REDIS_BLOG_KEY); err != nil {
		return models.InternalServerError("Error while deleting blog from cache")
	}

	blogs, _ := b.repository.GetBlogs(ctx)
	blogsJSON, _ := b.helper.Marshal(blogs)
	b.redisCache.Set(ctx, b.env.REDIS_BLOG_KEY, string(blogsJSON), b.cacheTTL)

	return nil
}

func (b *blogUsecase) TrackPopularity(ctx context.Context, blogID string, popularity dtos.TrackPopularityRequest) *models.ErrorResponse {
	return b.repository.TrackPopularity(ctx, blogID, popularity)
}
