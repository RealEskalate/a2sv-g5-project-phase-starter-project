package blog_usecase

import (
	"blog-api/domain"
	"context"
)

func (bu *BlogUsecase) GetBlogs(ctx context.Context, page, limit int, sortBy string) ([]domain.Blog, int, error) {
	// Get blog posts from the repository
	blogs, err := bu.blogRepo.GetPaginatedBlogs(context.Background(), page, limit, sortBy)
	if err != nil {
		return nil, 0, err
	}

	// Get total number of posts for pagination metadata
	totalPosts, err := bu.blogRepo.GetTotalBlogs(context.Background())
	if err != nil {
		return nil, 0, err
	}

	// Fetch additional popularity metrics
	for i := range blogs {
		blogID := blogs[i].ID
		blogs[i].Likes, err = bu.likeRepo.GetLikesCount(context.TODO(), blogID)
		if err != nil {
			return nil, 0, err
		}
		blogs[i].Comments, err = bu.commentRepo.GetCommentsCount(context.TODO(), blogID)
		if err != nil {
			return nil, 0, err
		}
	}

	return blogs, totalPosts, nil
}
