package usecases

import (
	"context"
	"time"

	"strconv"

	config "github.com/aait.backend.g5.main/backend/Config"
	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type blogUsecase struct {
	repository   interfaces.BlogRepository
	userRepo     interfaces.UserRepository
	popularity   interfaces.BlogPopularityActionRepository
	comment      interfaces.BlogCommentRepository
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
	comment interfaces.BlogCommentRepository,
) interfaces.BlogUsecase {

	return &blogUsecase{
		repository:   repository,
		cacheService: cacheService,
		env:          env,
		cacheTTL:     cacheTTL,
		helper:       helper,
		userRepo:     userRepo,
		popularity:   popularity,
		comment:      comment,
	}
}

func (b *blogUsecase) CreateBlog(ctx context.Context, blog *models.Blog) (*dtos.BlogResponse, *models.ErrorResponse) {
	slug := b.helper.CreateSlug(blog.Title)
	blog.Slug = slug

	newBlog, err := b.repository.CreateBlog(ctx, blog)

	if err != nil {
		return nil, err
	}
	if err := b.cacheService.InvalidateAllBlogCaches(ctx); err != nil {
		return nil, models.InternalServerError("Error while invalidating all blog caches")
	}
	return &dtos.BlogResponse{
		Blog:       *newBlog,
		Comments:   []models.Comment{},
		Popularity: models.Popularity{},
	}, nil
}

func (b *blogUsecase) GetBlog(ctx context.Context, id string) (*dtos.BlogResponse, *models.ErrorResponse) {
	data, err := b.helper.FetchFromCacheOrRepo(ctx, id, func() (interface{}, *models.ErrorResponse) {
		return b.repository.GetBlog(ctx, id)
	})

	if err != nil {
		return nil, err
	}

	blog := *data.(*models.Blog)
	if err := b.repository.IncreaseView(ctx, blog.ID); err != nil {
		return nil, err
	}

	blogComments, commentErr := b.helper.FetchComments(ctx, blog.ID)
	if commentErr != nil {
		return nil, commentErr
	}

	blogPopularity, popularityErr := b.helper.FetchPopularity(ctx, blog.ID)
	if popularityErr != nil {
		return nil, popularityErr
	}

	return &dtos.BlogResponse{
		Blog:       blog,
		Comments:   blogComments,
		Popularity: *blogPopularity,
	}, nil
}

func (b *blogUsecase) GetBlogs(ctx context.Context, page int) ([]*dtos.BlogResponse, *models.ErrorResponse) {

	data, err := b.helper.FetchFromCacheOrRepoBlogs(ctx, strconv.Itoa(page), func() (interface{}, *models.ErrorResponse) {
		return b.repository.GetBlogs(ctx, page)
	})

	if err != nil {
		return nil, err
	}

	blogs, ok := data.([]*models.Blog)
	if !ok {
		return nil, &models.ErrorResponse{Message: "Data type mismatch: expected []*models.Blog"}
	}

	return b.helper.GetBlogs(ctx, blogs)
}

func (b *blogUsecase) SearchBlogs(ctx context.Context, filter dtos.FilterBlogRequest) ([]*dtos.BlogResponse, *models.ErrorResponse) {

	if filter.AuthorName != "" {

		author, Err := b.userRepo.GetUserByName(ctx, filter.AuthorName)

		if Err != nil {
			return nil, Err
		}

		filter.AuthorID = author.ID

	}

	blogs, nErr := b.repository.SearchBlogs(ctx, filter)

	if nErr != nil {
		return nil, nErr
	}

	return b.helper.GetBlogs(ctx, blogs)
}

func (b *blogUsecase) UpdateBlog(ctx context.Context, blogID string, blog *models.Blog) *models.ErrorResponse {

	if _, err := b.repository.GetBlog(ctx, blogID); err != nil {
		return err
	}

	if err := b.repository.UpdateBlog(ctx, blogID, blog); err != nil {
		return err
	}

	if err := b.cacheService.InvalidateAllBlogCaches(ctx); err != nil {
		return models.InternalServerError("Error while invalidating all blog caches")
	}

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

	if blog.AuthorID != deleteBlogReq.AuthorID && user.Role != models.RoleAdmin {
		return models.Unauthorized("You are not authorized to delete this blog")
	}

	id := deleteBlogReq.BlogID
	if err := b.comment.DeleteComments(ctx, id); err != nil {
		return err
	}

	if err := b.repository.DeleteBlog(ctx, id); err != nil {
		return err
	}

	if err := b.cacheService.InvalidateAllBlogCaches(ctx); err != nil {
		return models.InternalServerError("Error while invalidating all blog caches")
	}

	return nil
}

func (b *blogUsecase) TrackPopularity(ctx context.Context, popularity dtos.TrackPopularityRequest) *models.ErrorResponse {

	if popularity.Action != models.Like && popularity.Action != models.Dislike {
		return models.BadRequest("Invalid action")
	}

	existingAction, err := b.popularity.GetBlogPopularityAction(ctx, popularity.BlogID, popularity.UserID)

	if err != nil && err.Code != 404 {
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
