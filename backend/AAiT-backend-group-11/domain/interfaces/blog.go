package interfaces

import (
	"backend-starter-project/domain/entities"
	"time"
)

type BlogRepository interface {
	CreateBlogPost(blogPost *entities.BlogPost) (*entities.BlogPost, error)
	GetBlogPostById(blogPostId string) (*entities.BlogPost, error)
	UpdateBlogPost(blogPost *entities.BlogPost) (*entities.BlogPost, error)
	DeleteBlogPost(blogPostId string) error
	GetBlogPosts(page, pageSize int, sortBy string) ([]entities.BlogPost, error)
	SearchBlogPosts(criteria string) ([]entities.BlogPost, error)
	FilterBlogPosts(tags []string, dateRange []time.Time, sortBy string) ([]entities.BlogPost, error)
}

type BlogService interface {
	CreateBlogPost(blogPost *entities.BlogPost) (*entities.BlogPost, error)
	GetBlogPostById(blogPostId string) (*entities.BlogPost, error)
	UpdateBlogPost(blogPost *entities.BlogPost) (*entities.BlogPost, error)
	DeleteBlogPost(blogPostId string) error
	GetBlogPosts(page, pageSize int, sortBy string) ([]entities.BlogPost, error)
	SearchBlogPosts(criteria string) ([]entities.BlogPost, error)
	FilterBlogPosts(tags []string, dateRange []time.Time, sortBy string) ([]entities.BlogPost, error)
}




type PopularityTrackingService interface {
    IncrementViewCount(blogPostId string) error
    LikeBlogPost(blogPostId, userId string) error
    DislikeBlogPost(blogPostId, userId string) error
    GetPopularityMetrics(blogPostId string) (map[string]int, error)
}
