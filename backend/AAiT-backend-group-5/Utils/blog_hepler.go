package utils

import (
	"context"
	"encoding/json"
	"regexp"
	"strings"
	"time"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"github.com/go-redis/redis/v8"
	"github.com/mitchellh/mapstructure"
)

type blogHelper struct {
	repository   interfaces.BlogRepository
	cacheService interfaces.RedisCache
	cacheTTL     time.Duration
}

func NewBlogHelper(repository interfaces.BlogRepository, cacheService interfaces.RedisCache, cacheTTL time.Duration) interfaces.BlogHelper {
	return &blogHelper{
		repository:   repository,
		cacheService: cacheService,
	}
}

func (blog *blogHelper) CreateSlug(blogTitle string) string {
	slug := strings.ToLower(blogTitle)
	reg := regexp.MustCompile("[^a-z0-9]+")
	slug = reg.ReplaceAllString(slug, "-")

	slug = strings.Trim(slug, "-")

	return slug
}

func (blog *blogHelper) Marshal(data interface{}) (string, *models.ErrorResponse) {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return "", models.InternalServerError("Error while marshalling data: " + err.Error())
	}
	return string(dataJSON), nil
}

func (blog *blogHelper) Unmarshal(dataJSON string, result interface{}) *models.ErrorResponse {
	if err := json.Unmarshal([]byte(dataJSON), result); err != nil {
		return models.InternalServerError("Error while unmarshalling data: " + err.Error())
	}
	return nil
}

func (b *blogHelper) GetBlogs(ctx context.Context, data []*models.Blog) ([]*dtos.BlogResponse, *models.ErrorResponse) {
	var response []*dtos.BlogResponse

	for _, blog := range data {
		blogComments, commentErr := b.FetchComments(ctx, blog.ID)
		if commentErr != nil {
			return nil, commentErr
		}

		blogPopularity, popularityErr := b.FetchPopularity(ctx, blog.ID)
		if popularityErr != nil {
			return nil, popularityErr
		}

		response = append(response, &dtos.BlogResponse{
			Blog:       *blog,
			Comments:   blogComments,
			Popularity: *blogPopularity,
		})
	}

	return response, nil
}

func (b *blogHelper) FetchComments(ctx context.Context, blogID string) ([]models.Comment, *models.ErrorResponse) {
	comments, err := b.repository.GetComments(ctx, blogID)
	if err != nil {
		if err.Code == 404 {
			return []models.Comment{}, nil
		}
		return nil, err
	}
	return comments, nil
}

func (b *blogHelper) FetchPopularity(ctx context.Context, blogID string) (*models.Popularity, *models.ErrorResponse) {
	popularity, err := b.repository.GetPopularity(ctx, blogID)
	if err != nil {
		if err.Code == 404 {
			return &models.Popularity{}, nil
		}
		return nil, err
	}
	return popularity, nil
}

func (b *blogHelper) FetchFromCacheOrRepo(ctx context.Context, cacheKey string, fetchFromRepo func() (interface{}, *models.ErrorResponse)) (interface{}, *models.ErrorResponse) {
	cachedData, err := b.cacheService.Get(ctx, cacheKey)

	if err == redis.Nil {
		data, repoErr := fetchFromRepo()
		if repoErr != nil {
			return nil, repoErr
		}

		dataJSON, marshalErr := b.Marshal(data)
		if marshalErr != nil {
			return nil, models.InternalServerError("Error while marshalling data")
		}
		b.cacheService.Set(ctx, cacheKey, string(dataJSON), b.cacheTTL)

		return data, nil
	} else if err != nil {
		return nil, models.InternalServerError("Error while fetching data from cache")
	}

	var blog models.Blog
	if unmarshalErr := json.Unmarshal([]byte(cachedData), &blog); unmarshalErr != nil {
		return nil, models.InternalServerError("Error while unmarshalling data from cache")
	}

	return &blog, nil
}

func (b *blogHelper) FetchFromCacheOrRepoBlogs(ctx context.Context, cacheKey string, fetchFromRepo func() (interface{}, *models.ErrorResponse)) (interface{}, *models.ErrorResponse) {
	cachedData, err := b.cacheService.Get(ctx, cacheKey)

	if err == redis.Nil {
		data, repoErr := fetchFromRepo()

		if repoErr != nil {
			return nil, repoErr
		}

		dataJSON, marshalErr := b.Marshal(data)
		if marshalErr != nil {
			return nil, models.InternalServerError("Error while marshalling data")
		}
		b.cacheService.Set(ctx, cacheKey, string(dataJSON), b.cacheTTL)

		return data, nil
	} else if err != nil {
		return nil, models.InternalServerError("Error while fetching data from cache")
	}

	var result []map[string]interface{}
	if unmarshalErr := json.Unmarshal([]byte(cachedData), &result); unmarshalErr != nil {
		return nil, models.InternalServerError("Error while unmarshalling data from cache")
	}

	var blogs []*models.Blog
	for _, item := range result {
		var blog models.Blog
		if err := mapstructure.Decode(item, &blog); err != nil {
			return nil, models.InternalServerError("Error while converting cached data to Blog struct")
		}
		blogs = append(blogs, &blog)
	}

	return blogs, nil
}
