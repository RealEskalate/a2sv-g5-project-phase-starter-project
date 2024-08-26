package blogusecase

import (
	"blogs/domain"
)

type BlogUsecase struct {
	BlogRepo domain.BlogRepository
	TagRepo domain.TagRepository
}

func NewBlogUsecase(br domain.BlogRepository, tr domain.TagRepository) *BlogUsecase {
	
	return &BlogUsecase{
		BlogRepo: br,
		TagRepo: tr,
	}
}
