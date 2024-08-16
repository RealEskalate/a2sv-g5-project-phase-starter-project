package usecases

import (
	"blogapp/Domain"
	"time"
)

type blogUsecase struct {
	AuthRepository Domain.BlogRepository
	contextTimeout time.Duration
}

func NewBlogUseCase(repo Domain.BlogRepository) *blogUsecase {
	return &blogUsecase{
		AuthRepository: repo,
		contextTimeout: time.Second * 10,
	}
}
