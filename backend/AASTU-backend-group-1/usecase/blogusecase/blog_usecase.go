package blogusecase

import (
	"blogs/domain"
)

type BlogUsecase struct {
	BlogRepo domain.BlogRepository
}

func NewBlogUsecase(br domain.BlogRepository) domain.BlogUsecase {
	return &BlogUsecase{
		BlogRepo: br,
	}
}
