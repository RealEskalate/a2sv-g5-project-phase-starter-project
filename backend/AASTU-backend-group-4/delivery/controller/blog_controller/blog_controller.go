package blog_controller

import (
	"blog-api/domain/blog"
)

type blogController struct {
	usecase blog.BlogUsecase
}
