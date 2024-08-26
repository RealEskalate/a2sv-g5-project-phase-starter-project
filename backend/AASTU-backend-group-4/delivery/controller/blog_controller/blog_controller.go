package blog_controller

import (
	"blog-api/domain"
)

type BlogController struct {
	usecase domain.BlogUsecase
}

func NewBlogController(usecase domain.BlogUsecase) *BlogController {
	return &BlogController{
		usecase: usecase,
	}
}
