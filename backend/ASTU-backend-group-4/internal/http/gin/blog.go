package gin

import "github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/blog"

type BlogController struct {
	blogUseCase blog.BlogUseCase
}

func NewBlogController(blogUseCase blog.BlogUseCase) *BlogController {
	return &BlogController{
		blogUseCase: blogUseCase,
	}
}
