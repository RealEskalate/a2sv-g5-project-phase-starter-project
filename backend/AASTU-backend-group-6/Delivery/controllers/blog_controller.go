package controllers

import (
	domain "blogs/Domain"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	BlogUsecase domain.BlogUsecase
}

// CommentOnBlog implements domain.BlogUsecase.
func (b BlogController) CommentOnBlog(c *gin.Context) {
	panic("unimplemented")
}

// CreateBlog implements domain.BlogUsecase.
func (b BlogController) CreateBlog(c *gin.Context) {
	panic("unimplemented")
}

// DeleteBlogByID implements domain.BlogUsecase.
func (b BlogController) DeleteBlogByID(c *gin.Context) {
	panic("unimplemented")
}

// FilterBlogsByTag implements domain.BlogUsecase.
func (b BlogController) FilterBlogsByTag(c *gin.Context) {
	panic("unimplemented")
}

// GetBlogByID implements domain.BlogUsecase.
func (b BlogController) GetBlogByID(c *gin.Context) {
	c.JSON(200, gin.H{"des blognal": "blog"})
}

// GetBlogs implements domain.BlogUsecase.
func (b BlogController) GetBlogs(c *gin.Context) {
	c.JSON(200, gin.H{"des blogs": "blogs"})
}

// GetMyBlogByID implements domain.BlogUsecase.
func (b BlogController) GetMyBlogByID(c *gin.Context) {
	panic("unimplemented")
}

// GetMyBlogs implements domain.BlogUsecase.
func (b BlogController) GetMyBlogs(c *gin.Context) {
	panic("unimplemented")
}

// SearchBlogByTitleAndAuthor implements domain.BlogUsecase.
func (b BlogController) SearchBlogByTitleAndAuthor(c *gin.Context) {
	panic("unimplemented")
}

// UpdateBlogByID implements domain.BlogUsecase.
func (b BlogController) UpdateBlogByID(c *gin.Context) {
	panic("unimplemented")
}

func NewBlogController(BlogUsecase domain.BlogUsecase) BlogController {
	return BlogController{
		BlogUsecase: BlogUsecase,
	}
}
