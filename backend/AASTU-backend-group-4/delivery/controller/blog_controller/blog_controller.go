package blog_controller

import (
	domain "blog-api/domain/blog"
)

type BlogController struct {
	usecase domain.BlogUsecase
}
