package blog_usecase

import (
	"time"

	"blog-api/domain"
)

type BlogUsecase struct {
	blogRepo       domain.BlogRepository
	commentRepo    domain.CommentRepository
	likeRepo       domain.LikeRepository
	contextTimeout time.Duration
}

func NewBlogUsecase(blogRepository domain.BlogRepository, commentRepository domain.CommentRepository, likeRepository domain.LikeRepository, timeout time.Duration) *BlogUsecase {
	return &BlogUsecase{
		blogRepo:       blogRepository,
		commentRepo:    commentRepository,
		likeRepo:       likeRepository,
		contextTimeout: timeout,
	}
}
