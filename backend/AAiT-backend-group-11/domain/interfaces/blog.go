package interfaces

import (
	"backend-starter-project/domain/entities"
	"time"
)

type BlogRepository interface {
	CreateBlogPost(c context.Context, blogPost *entities.BlogPost) (*entities.BlogPost, error)
	GetBlogPostById(c context.Context, blogPostId string) (*entities.BlogPost, error)
	UpdateBlogPost(c context.Context, blogPost *entities.BlogPost) (*entities.BlogPost, error)
	DeleteBlogPost(c context.Context, blogPostId string) error
	GetBlogPosts(c context.Context, page, pageSize int, sortBy string) ([]entities.BlogPost,error)
	SearchBlogPosts(c context.Context, criteria string, tags []string, startDate, endDate time.Time)([]entities.BlogPost, error)
	FilterBlogPosts(c context.Context, tags []string, dateRange []time.Time, sortBy string) ([]entities.BlogPost, error)
    LikeBlogPost(c context.Context,blogPostId, userId string) error
	DislikeBlogPost(c context.Context,blogPostId, userId string) error
	IncrementViewPost(c context.Context, postID, userID string) error
	CountBlogPosts(c context.Context) (int, error)

}

type BlogService interface {
	CreateBlogPost(c context.Context, blogPost *entities.BlogPost) (*entities.BlogPost, error)
	GetBlogPostById(c context.Context, blogPostId string) (*entities.BlogPost, error)
	UpdateBlogPost(c context.Context, blogPost *entities.BlogPost) (*entities.BlogPost, error)
	DeleteBlogPost(c context.Context, blogPostId string) error
	GetBlogPosts(c context.Context, page, pageSize int, sortBy string) ([]entities.BlogPost,int, error)
	SearchBlogPosts(c context.Context, criteria string, tags []string, startDate, endDate time.Time) ([]entities.BlogPost, error)
	FilterBlogPosts(c context.Context, tags []string, dateRange []time.Time, sortBy string) ([]entities.BlogPost, error)
}




type PopularityTrackingService interface {
    IncrementViewCount(blogPostId string) error
    LikeBlogPost(blogPostId, userId string) error
    DislikeBlogPost(blogPostId, userId string) error
    GetPopularityMetrics(blogPostId string) (map[string]int, error)
}