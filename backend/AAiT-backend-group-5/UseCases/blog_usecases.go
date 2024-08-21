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
	repository   interfaces.BlogRepository
	userRepo     interfaces.UserRepository
	popularity   interfaces.BlogPopularityActionRepository
	cacheService interfaces.RedisCache
	env          config.Env
	cacheTTL     time.Duration
	helper       interfaces.BlogHelper
}

func NewblogUsecase(
	repository interfaces.BlogRepository,
	cacheService interfaces.RedisCache,
	env config.Env,
	cacheTTL time.Duration,
	helper interfaces.BlogHelper,
	userRepo interfaces.UserRepository,
	popularity interfaces.BlogPopularityActionRepository,
) interfaces.BlogUsecase {

	return &blogUsecase{
		repository:   repository,
		cacheService: cacheService,
		env:          env,
		cacheTTL:     cacheTTL,
		helper:       helper,
		userRepo:     userRepo,
		popularity:   popularity,
	}
}

func (b *blogUsecase) getBlogs(ctx context.Context, data []*models.Blog) ([]*dtos.BlogResponse, *models.ErrorResponse) {
	var response []*dtos.BlogResponse

	for _, blog := range data {
		blogComments, commentErr := b.fetchComments(ctx, blog.ID)
		if commentErr != nil {
			return nil, commentErr
		}

		blogPopularity, popularityErr := b.fetchPopularity(ctx, blog.ID)
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

func (b *blogUsecase) fetchComments(ctx context.Context, blogID string) ([]models.Comment, *models.ErrorResponse) {
	comments, err := b.repository.GetComments(ctx, blogID)
	if err != nil {
		if err.Code == 404 {
			return []models.Comment{}, nil
		}
		return nil, err
	}
	return comments, nil
}

func (b *blogUsecase) fetchPopularity(ctx context.Context, blogID string) (*models.Popularity, *models.ErrorResponse) {
	popularity, err := b.repository.GetPopularity(ctx, blogID)
	if err != nil {
		if err.Code == 404 {
			return &models.Popularity{}, nil
		}
		return nil, err
	}
	return popularity, nil
}

func (b *blogUsecase) fetchFromCacheOrRepo(ctx context.Context, cacheKey string, fetchFromRepo func() (interface{}, *models.ErrorResponse)) (interface{}, *models.ErrorResponse) {
	cachedData, err := b.cacheService.Get(ctx, cacheKey)

	if err == redis.Nil {
		// Fetch from repository if cache is empty
		data, repoErr := fetchFromRepo()
		if repoErr != nil {
			return nil, repoErr
		}

		// Marshal and store in cache
		dataJSON, marshalErr := b.helper.Marshal(data)
		if marshalErr != nil {
			return nil, models.InternalServerError("Error while marshalling data")
		}
		b.cacheService.Set(ctx, cacheKey, string(dataJSON), b.cacheTTL)

		return data, nil
	} else if err != nil {
		return nil, models.InternalServerError("Error while fetching data from cache")
	}

	// Unmarshal directly into the expected struct
	var blog models.Blog
	if unmarshalErr := json.Unmarshal([]byte(cachedData), &blog); unmarshalErr != nil {
		return nil, models.InternalServerError("Error while unmarshalling data from cache")
	}

	return &blog, nil
}


func (b *blogUsecase) CreateBlog(ctx context.Context, blog *models.Blog) (*dtos.BlogResponse, *models.ErrorResponse) {
	slug := b.helper.CreateSlug(blog.Title)
	blog.Slug = slug

	newBlog, err := b.repository.CreateBlog(ctx, blog)

	if err != nil {
		return nil, err
	}

	if err := b.cacheService.Delete(ctx, b.env.REDIS_BLOG_KEY); err != nil {
		return nil, models.InternalServerError("Error while Invalidating blog from cache")
	}

	return &dtos.BlogResponse{
		Blog:       *newBlog,
		Comments:   []models.Comment{},
		Popularity: models.Popularity{},
	}, nil
}

func (b *blogUsecase) GetBlog(ctx context.Context, id string) (*dtos.BlogResponse, *models.ErrorResponse) {
	data, err := b.fetchFromCacheOrRepo(ctx, id, func() (interface{}, *models.ErrorResponse) {
		return b.repository.GetBlog(ctx, id)
	})

	if err != nil {
		return nil, err
	}

	blog := *data.(*models.Blog)

	blogComments, commentErr := b.fetchComments(ctx, blog.ID)
	if commentErr != nil {
		return nil, commentErr
	}

	blogPopularity, popularityErr := b.fetchPopularity(ctx, blog.ID)
	if popularityErr != nil {
		return nil, popularityErr
	}

	return &dtos.BlogResponse{
		Blog:       blog,
		Comments:   blogComments,
		Popularity: *blogPopularity,
	}, nil
}

func (b *blogUsecase) GetBlogs(ctx context.Context) ([]*dtos.BlogResponse, *models.ErrorResponse) {
	data, err := b.fetchFromCacheOrRepo(ctx, b.env.REDIS_BLOG_KEY, func() (interface{}, *models.ErrorResponse) {
		return b.repository.GetBlogs(ctx)
	})

	if err != nil {
		return nil, err
	}

	blogs, ok := data.([]*models.Blog)
	if !ok {
		return nil, &models.ErrorResponse{Message: "Data type mismatch: expected []*models.Blog"}
	}

	return b.getBlogs(ctx, blogs)
}

func (b *blogUsecase) SearchBlogs(ctx context.Context, filter dtos.FilterBlogRequest) ([]*dtos.BlogResponse, *models.ErrorResponse) {
	searchJson, err := json.Marshal(filter)

	if err != nil {
		return nil, models.InternalServerError("Error while marshalling search filter")
	}

	cachedBlogs, sErr := b.cacheService.Get(ctx, string(searchJson))

	if sErr == redis.Nil {
		blogs, nErr := b.repository.SearchBlogs(ctx, filter)

		resBlogs, rErr := b.getBlogs(ctx, blogs)

		if nErr != nil {
			return nil, nErr
		}

		if rErr != nil {
			return nil, rErr
		}

		blogsJson, _ := b.helper.Marshal(resBlogs)
		b.cacheService.Set(ctx, string(searchJson), string(blogsJson), b.cacheTTL)

		return resBlogs, nil
	} else if sErr != nil {
		return nil, models.InternalServerError("Error while fetching search query from cache")
	}

	var blogs []*models.Blog
	if err := b.helper.Unmarshal(cachedBlogs, &blogs); err != nil {
		return nil, models.InternalServerError("Error while unmarshalling search query from cache")
	}

	return b.getBlogs(ctx, blogs)
}

func (b *blogUsecase) UpdateBlog(ctx context.Context, blogID string, blog *models.Blog) *models.ErrorResponse {

	if err := b.repository.UpdateBlog(ctx, blogID, blog); err != nil {
		return err
	}

	if err := b.cacheService.Delete(ctx, blogID); err != nil {
		return models.InternalServerError("Error while deleting blog from cache")
	}

	updatedBlog, _ := b.repository.GetBlog(ctx, blogID)
	dataJSON, _ := b.helper.Marshal(updatedBlog)
	b.cacheService.Set(ctx, blogID, string(dataJSON), b.cacheTTL)

	return nil
}

func (b *blogUsecase) DeleteBlog(ctx context.Context, deleteBlogReq dtos.DeleteBlogRequest) *models.ErrorResponse {

	blog, err := b.repository.GetBlog(ctx, deleteBlogReq.BlogID)
	user, uErr := b.userRepo.GetUserByID(ctx, deleteBlogReq.AuthorID)
	if err != nil {
		return err
	}

	if uErr != nil {
		return uErr
	}

	if blog.AuthorID != deleteBlogReq.AuthorID || user.Role != "admin" {
		return models.Unauthorized("You are not authorized to delete this blog")
	}

	id := deleteBlogReq.BlogID
	if err := b.repository.DeleteBlog(ctx, id); err != nil {
		return err
	}

	if err := b.cacheService.Delete(ctx, b.env.REDIS_BLOG_KEY); err != nil {
		return models.InternalServerError("Error while deleting blog from cache")
	}

	blogs, _ := b.repository.GetBlogs(ctx)
	blogsJSON, _ := b.helper.Marshal(blogs)
	b.cacheService.Set(ctx, b.env.REDIS_BLOG_KEY, string(blogsJSON), b.cacheTTL)

	return nil
}

func (b *blogUsecase) TrackPopularity(ctx context.Context, popularity dtos.TrackPopularityRequest) *models.ErrorResponse {

	existingAction, err := b.popularity.GetBlogPopularityAction(ctx, popularity.BlogID, popularity.UserID)

	if err != nil {
		return err
	}

	if existingAction != nil {
		if existingAction.Action == popularity.Action {
			switch existingAction.Action {
			case models.Like:
				err := b.popularity.UndoLike(ctx, popularity)
				if err != nil {
					return err
				}
			case models.Dislike:
				err := b.popularity.UndoDislike(ctx, popularity)
				if err != nil {
					return err
				}
			}
			return nil
		}

		switch existingAction.Action {
		case models.Like:
			if popularity.Action == models.Dislike {
				err := b.popularity.UndoLike(ctx, popularity)
				if err != nil {
					return err
				}
				err = b.popularity.Dislike(ctx, popularity)
				if err != nil {
					return err
				}
			}
		case models.Dislike:
			if popularity.Action == models.Like {
				err := b.popularity.UndoDislike(ctx, popularity)
				if err != nil {
					return err
				}
				err = b.popularity.Like(ctx, popularity)
				if err != nil {
					return err
				}
			}
		}
	} else {
		if popularity.Action == models.Like {
			return b.popularity.Like(ctx, popularity)
		}
		return b.popularity.Dislike(ctx, popularity)
	}

	return nil
}

func (b *blogUsecase) AddComment(ctx context.Context, comment models.Comment) *models.ErrorResponse {
	return b.repository.AddComment(ctx, comment)
}
