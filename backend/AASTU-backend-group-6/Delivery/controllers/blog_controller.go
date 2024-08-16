package controllers

import (
	domain "blogs/Domain"
	"fmt"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	BlogUsecase domain.BlogUsecase
}

func NewBlogController(BlogUsecase domain.BlogUsecase) *BlogController {
	return &BlogController{
		BlogUsecase: BlogUsecase,
	}
}
func UpdateBlogController(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	var post domain.Blog
	if err := c.ShouldBind(&post); err == nil {

	}
}
