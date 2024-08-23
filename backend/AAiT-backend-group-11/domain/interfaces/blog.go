package interfaces

import (
	"backend-starter-project/domain/dto"
	"backend-starter-project/domain/entities"
	"time"
)

type BlogRepository interface {
	CreateBlogPost( blogPost *entities.BlogPost, userId string) (*entities.BlogPost, error)
	GetBlogPostById( blogPostId string) (*entities.BlogPost, error)
	UpdateBlogPost( blogPost *entities.BlogPost) (*entities.BlogPost, error)
	DeleteBlogPost( blogPostId string) error
	GetBlogPosts( page, pageSize int, sortBy string) ([]entities.BlogPost,error)
	SearchBlogPosts( criteria string)([]entities.BlogPost, error)
	FilterBlogPosts( tags []string, startDate, endDate time.Time, sortBy string) ([]entities.BlogPost, error)
    LikeBlogPost(blogPostId, userId string) error
	DislikeBlogPost(blogPostId, userId string) error
	IncrementViewPost( postID, userID string) error
	CountBlogPosts()(int, error)
	ChangeCommentCount(blogPostId string, val int) error 

}

type BlogService interface {
	CreateBlogPost(blogPost *dto.AddBlogRequest, userId string) (*dto.AddBlogResponse, error)
	GetBlogPostById(blogPostId string, userId string) (*dto.GetBlogByIDResponse, error)
	GetBlogPosts(page, pageSize int, sortBy string) (*dto.GetBlogPostsResponse,int, error)
	UpdateBlogPost(blogPost *dto.UpdateBlogRequest, userId string) (*dto.UpdateBlogResponse, error)
	DeleteBlogPost(blogPostId,userId,role string) error
	SearchBlogPosts(criteria string) (*dto.GetBlogPostsResponse, error)
	FilterBlogPosts(dto.FilterBlogPostsRequest) (*dto.GetBlogPostsResponse, error)
}




type PopularityTrackingService interface {
    IncrementViewCount( blogPostId string) error
    LikeBlogPost( blogPostId, userId string) error
    DislikeBlogPost( blogPostId, userId string) error
    GetPopularityMetrics( blogPostId string) (map[string]int, error)
}
