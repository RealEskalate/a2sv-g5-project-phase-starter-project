package blog_usecase

import (
	"time"

	"blog-api/domain"
)

type BlogUsecase struct {
	blogRepo       domain.BlogRepository
	commentRepo    domain.CommentRepository
	likeRepo       domain.LikeRepository
	genAIService   domain.AIContentGenerator
	contextTimeout time.Duration
}

func NewBlogUsecase(blogRepo domain.BlogRepository, commentRepo domain.CommentRepository, likeRepo domain.LikeRepository, aiService domain.AIContentGenerator, timeout time.Duration) domain.BlogUsecase {
	return &BlogUsecase{
		blogRepo:       blogRepo,
		commentRepo:    commentRepo,
		likeRepo:       likeRepo,
		genAIService:   aiService,
		contextTimeout: timeout,
	}
}
