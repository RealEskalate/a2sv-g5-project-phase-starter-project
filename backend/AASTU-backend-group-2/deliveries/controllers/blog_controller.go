package controllers

import (
	"blog_g2/domain"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	Blogusecase domain.BlogUsecase
}

// Blog-controller constructor
func NewBlogController(Blogmgr domain.BlogUsecase) *BlogController {
	return &BlogController{
		Blogusecase: Blogmgr,
	}

}

func (controller *BlogController) CreateBlog(c *gin.Context) {

}
func (controller *BlogController) RetrieveBlog(c *gin.Context) {

}
func (controller *BlogController) UpdateBlog(c *gin.Context) {

}
func (controller *BlogController) DeleteBlog(c *gin.Context) {

}
func (controller *BlogController) SearchBlog(c *gin.Context) {

}
func (controller *BlogController) FilterBlog(c *gin.Context) {

}
