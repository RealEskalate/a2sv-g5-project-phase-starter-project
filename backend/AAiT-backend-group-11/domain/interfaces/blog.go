package interfaces

import (
	"backend-starter-project/domain/entities"
	"context"
	"time"
)

type BlogRepository interface {
	CreateBlogPost(c context.Context, blogPost *entities.BlogPost) (*entities.BlogPost, error)
	GetBlogPostById(c context.Context, blogPostId string) (*entities.BlogPost, error)
	UpdateBlogPost(c context.Context, blogPost *entities.BlogPost) (*entities.BlogPost, error)
	DeleteBlogPost(c context.Context, blogPostId string) error
	GetBlogPosts(c context.Context, page, pageSize int, sortBy string) ([]entities.BlogPost, error)
	SearchBlogPosts(c context.Context, criteria string) ([]entities.BlogPost, error)
	FilterBlogPosts(c context.Context, tags []string, dateRange []time.Time, sortBy string) ([]entities.BlogPost, error)
    LikeBlogPost(c context.Context,blogPostId, userId string) error
	DislikeBlogPost(c context.Context,blogPostId, userId string) error
	ViewPost(c context.Context, postID, userID string) error
}

type BlogService interface {
	CreateBlogPost(c context.Context, blogPost *entities.BlogPost) (*entities.BlogPost, error)
	GetBlogPostById(c context.Context, blogPostId string) (*entities.BlogPost, error)
	UpdateBlogPost(c context.Context, blogPost *entities.BlogPost) (*entities.BlogPost, error)
	DeleteBlogPost(c context.Context, blogPostId string) error
	GetBlogPosts(c context.Context, page, pageSize int, sortBy string) ([]entities.BlogPost, error)
	SearchBlogPosts(c context.Context, criteria string) ([]entities.BlogPost, error)
	FilterBlogPosts(c context.Context, tags []string, dateRange []time.Time, sortBy string) ([]entities.BlogPost, error)
}




type PopularityTrackingService interface {
    IncrementViewCount(c context.Context, blogPostId string) error
    LikeBlogPost(c context.Context, blogPostId, userId string) error
    DislikeBlogPost(c context.Context, blogPostId, userId string) error
    GetPopularityMetrics(c context.Context, blogPostId string) (map[string]int, error)
}
