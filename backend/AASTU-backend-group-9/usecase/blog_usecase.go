package usecase

import (
	"context"
	"errors"
	"blog/domain"

	"time"

	
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "github.com/gin-gonic/gin"
)

type blogUsecase struct  {
	blogRepository domain.BlogRepository
	contextTimeout time.Duration
}

func NewBlogUsecase(blogRepository domain.BlogRepository, timeout time.Duration) domain.BlogUsecase {
	return &blogUsecase{
		blogRepository: blogRepository,
		contextTimeout: timeout,
	}
}

func (bu *blogUsecase) CreateBlog(ctx context.Context, req *domain.BlogCreationRequest, claims *domain.JwtCustomClaims) (*domain.BlogResponse, error) {
	


	blog := &domain.Blog{
		ID:        primitive.NewObjectID(),
		Title:     req.Title,
		Content:   req.Content,
		AuthorID:  claims.UserID,
		Tags:      req.Tags,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := bu.blogRepository.CreateBlog(ctx, blog); err != nil {
		return nil, err
	}

	return &domain.BlogResponse{
		ID:        blog.ID,
		Title:     blog.Title,
		Content:   blog.Content,
		Tags:      blog.Tags,
		CreatedAt: blog.CreatedAt,
		UpdatedAt: blog.UpdatedAt,
		AuthorID: blog.AuthorID,

	}, nil
}

func (bu *blogUsecase) GetBlogByID(ctx context.Context, id primitive.ObjectID) (*domain.Blog, error) {
	blog, err := bu.blogRepository.GetBlogByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return blog, nil
}

func (bu *blogUsecase) GetAllBlogs(ctx context.Context, page int, limit int, sortBy string) ([]*domain.BlogResponse, error) {
	blogs, err := bu.blogRepository.GetAllBlogs(ctx, page, limit, sortBy)
	if err != nil {
		return nil, err
	}

	var blogResponses []*domain.BlogResponse
	for _, blog := range blogs {
		blogResponses = append(blogResponses, &domain.BlogResponse{
			ID:        blog.ID,
			Title:     blog.Title,
			Content:   blog.Content,
			Tags:      blog.Tags,
			CreatedAt: blog.CreatedAt,
			UpdatedAt: blog.UpdatedAt,
			AuthorID: blog.AuthorID,
		})
	}

	return blogResponses, nil
}

func (bu *blogUsecase) UpdateBlog(ctx context.Context, id primitive.ObjectID, req *domain.BlogUpdateRequest) (*domain.BlogResponse, error) {
	blog, err := bu.blogRepository.GetBlogByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Title != "" {
		blog.Title = req.Title
	}
	if req.Content != "" {
		blog.Content = req.Content
	}
	if req.Tags != nil {
		blog.Tags = req.Tags
	}
	blog.UpdatedAt = time.Now()

	if err := bu.blogRepository.UpdateBlog(ctx, blog); err != nil {
		return nil, err
	}

	return &domain.BlogResponse{
		ID:        blog.ID,
		Title:     blog.Title,
		Content:   blog.Content,
		Tags:      blog.Tags,
		CreatedAt: blog.CreatedAt,
		UpdatedAt: blog.UpdatedAt,
		AuthorID: blog.AuthorID,
	}, nil
}

func (bu *blogUsecase) DeleteBlog(ctx context.Context, id primitive.ObjectID) error {
	return bu.blogRepository.DeleteBlog(ctx, id)
}

// Usecases/blog_usecase.go
// Usecases/blog_usecase.go
// domain/usecase.go
// controller/blog_controller.go
// usecase/blog_usecase.go
func (bu *blogUsecase) SearchBlogs(ctx context.Context, query string, filters *domain.BlogFilters) ([]*domain.BlogResponse, error) {
    // Implement the search functionality here
    blogs, err := bu.blogRepository.SearchBlogs(ctx, query, filters)
    if err != nil {
        return nil, err
    }

    var blogResponses []*domain.BlogResponse
    for _, blog := range blogs {
        blogResponses = append(blogResponses, &domain.BlogResponse{
            ID:        blog.ID,
            Title:     blog.Title,
            Content:   blog.Content,
            Tags:      blog.Tags,
            CreatedAt: blog.CreatedAt,
            UpdatedAt: blog.UpdatedAt,
            AuthorID:  blog.AuthorID,
        })
    }

    return blogResponses, nil
}

func (bu *blogUsecase) FilterBlogsByTags(ctx context.Context, tags []string) ([]*domain.Blog, error) {
	blogs, err := bu.blogRepository.FilterBlogsByTags(ctx, tags)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (bu *blogUsecase) FilterBlogsByDate(ctx context.Context, date string) ([]*domain.Blog, error) {
	blogs, err := bu.blogRepository.FilterBlogsByDate(ctx, date)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (bu *blogUsecase) FilterBlogsByPopularity(ctx context.Context, popularity string) ([]*domain.Blog, error) {
	blogs, err := bu.blogRepository.FilterBlogsByPopularity(ctx, popularity)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}



func (bu *blogUsecase) TrackView(ctx context.Context, id primitive.ObjectID) error {
    return bu.blogRepository.IncrementViews(ctx, id)
}

func (bu *blogUsecase) TrackLike(ctx context.Context, id primitive.ObjectID, userID string) error {
    liked, err := bu.blogRepository.HasUserLiked(ctx, id, userID)
    if err != nil {
        return err
    }
    if liked {
        return errors.New("user has already liked this post")
    }
    return bu.blogRepository.IncrementLikes(ctx, id)
}

func (bu *blogUsecase) TrackDislike(ctx context.Context, id primitive.ObjectID, userID string) error {
    disliked, err := bu.blogRepository.HasUserDisliked(ctx, id, userID)
    if err != nil {
        return err
    }
    if disliked {
        return errors.New("user has already disliked this post")
    }
    return bu.blogRepository.IncrementDislikes(ctx, id)
}

func (bu *blogUsecase) AddComment(ctx context.Context, id primitive.ObjectID, comment *domain.Comment) error {
    return bu.blogRepository.AddComment(ctx, id, comment)

}