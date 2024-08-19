package interfaces

import (
	"backend-starter-project/domain/entities"
	"context"
	"time"
)

type BlogRepository interface {
	CreateBlogPost( blogPost *entities.BlogPost, userId string) (*entities.BlogPost, error)
	GetBlogPostById( blogPostId string) (*entities.BlogPost, error)
	UpdateBlogPost( blogPost *entities.BlogPost) (*entities.BlogPost, error)
	DeleteBlogPost( blogPostId string) error
	GetBlogPosts( page, pageSize int, sortBy string) ([]entities.BlogPost,error)
	SearchBlogPosts( criteria string, tags []string, startDate, endDate time.Time)([]entities.BlogPost, error)
	FilterBlogPosts( tags []string, dateRange []time.Time, sortBy string) ([]entities.BlogPost, error)
    LikeBlogPost(blogPostId, userId string) error
	DislikeBlogPost(blogPostId, userId string) error
	IncrementViewPost( postID, userID string) error
	CountBlogPosts() (int, error)

}

type BlogService interface {
	CreateBlogPost(blogPost *entities.BlogPost, userId string) (*entities.BlogPost, error)
	GetBlogPostById(blogPostId string) (*entities.BlogPost, error)
	UpdateBlogPost(blogPost *entities.BlogPost) (*entities.BlogPost, error)
	DeleteBlogPost(blogPostId string) error
	GetBlogPosts(page, pageSize int, sortBy string) ([]entities.BlogPost,int, error)
	SearchBlogPosts(criteria string, tags []string, startDate, endDate time.Time) ([]entities.BlogPost, error)
	FilterBlogPosts(tags []string, dateRange []time.Time, sortBy string) ([]entities.BlogPost, error)
}




type PopularityTrackingService interface {
    IncrementViewCount(c context.Context, blogPostId string) error
    LikeBlogPost(c context.Context, blogPostId, userId string) error
    DislikeBlogPost(c context.Context, blogPostId, userId string) error
    GetPopularityMetrics(c context.Context, blogPostId string) (map[string]int, error)
}
