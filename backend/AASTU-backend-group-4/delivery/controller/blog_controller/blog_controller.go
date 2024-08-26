package blog_controller

import (
	"blog-api/domain"
	infrastructure "blog-api/infrastructure/cloudinary"
)

type BlogController struct {
	usecase domain.BlogUsecase
	Medcont infrastructure.MediaUpload
}

func NewBlogController(usecase domain.BlogUsecase, Medcont infrastructure.MediaUpload) *BlogController {
	return &BlogController{
		usecase: usecase,
		Medcont: Medcont,
	}
}

//
