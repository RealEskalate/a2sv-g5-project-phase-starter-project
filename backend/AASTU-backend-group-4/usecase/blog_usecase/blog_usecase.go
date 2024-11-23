package blog_usecase

import (
	"time"

	"blog-api/domain"
	"blog-api/infrastructure/bootstrap"
)

type BlogUsecase struct {
	blogRepo       domain.BlogRepository
	commentRepo    domain.CommentRepository
	likeRepo       domain.LikeRepository
	genAIService   domain.AIContentGenerator
	Env            bootstrap.Env
	contextTimeout time.Duration
}

func NewBlogUsecase(blogRepo domain.BlogRepository, commentRepo domain.CommentRepository, likeRepo domain.LikeRepository, aiService domain.AIContentGenerator, Env bootstrap.Env, timeout time.Duration) domain.BlogUsecase {
	return &BlogUsecase{
		blogRepo:       blogRepo,
		commentRepo:    commentRepo,
		likeRepo:       likeRepo,
		genAIService:   aiService,
		Env:            Env,
		contextTimeout: timeout,
	}
}
