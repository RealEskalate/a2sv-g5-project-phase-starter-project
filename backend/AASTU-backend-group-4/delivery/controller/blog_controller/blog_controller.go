package blog_controller

import (
	"blog-api/domain/blog"
)

type BlogController struct {
	usecase blog.BlogUsecase
}

func NewBlogController(usecase blog.BlogUsecase) *BlogController {
	return &BlogController{
		usecase: usecase,
	}
}
