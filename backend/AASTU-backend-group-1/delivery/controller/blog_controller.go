package controller

import (
	"blogs/domain"
)

type BlogController struct {
	BlogUsecase domain.BlogUsecase
}

func NewBlogController(bu domain.BlogUsecase) *BlogController {
	return &BlogController{
		BlogUsecase: bu,
	}
}
