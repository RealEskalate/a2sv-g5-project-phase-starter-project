package blog_usecase

import (
	"blog-api/domain"
	"context"
)

func (bu *BlogUsecase) GetBlogs(ctx context.Context, page, limit int, sortBy string) ([]domain.Blog, int, error) {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	// Get blog posts from the repository
	blogs, err := bu.blogRepo.GetPaginatedBlogs(ctx, page, limit, sortBy)
	if err != nil {
		return nil, 0, err
	}

	// Get total number of posts for pagination metadata
	totalPosts, err := bu.blogRepo.GetTotalBlogs(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Fetch additional popularity metrics
	for i := range blogs {
		blogID := blogs[i].ID
		blogs[i].Likes, err = bu.likeRepo.GetLikesCount(ctx, blogID)
		if err != nil {
			return nil, 0, err
		}
		blogs[i].Comments, err = bu.commentRepo.GetCommentsCount(ctx, blogID)
		if err != nil {
			return nil, 0, err
		}
	}

	return blogs, totalPosts, nil
}
