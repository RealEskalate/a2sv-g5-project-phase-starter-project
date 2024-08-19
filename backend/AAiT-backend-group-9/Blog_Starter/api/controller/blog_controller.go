package controller

import (
	"Blog_Starter/domain"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	blogUseCase domain.BlogUseCase
	c           context.Context
}

func NewBlogController(blogUseCase domain.BlogUseCase, c context.Context) *BlogController {
	return &BlogController{
		blogUseCase: blogUseCase,
		c:           c,
	}
}

// CreateBlog godoc
func (bc *BlogController) CreateBlog(c *gin.Context) {
	// implementation
	var blog domain.BlogCreate
	err := c.ShouldBindJSON(&blog)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	blogModel, err := bc.blogUseCase.CreateBlog(bc.c, &blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, blogModel)
}

// GetBlogByID godoc
func (bc *BlogController) GetBlogByID(c *gin.Context) {
	// implementation create a context and pass to the usecase not the gin context
	blogID := c.Param("blog_id")
	blog, err := bc.blogUseCase.GetBlogByID(bc.c, blogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blog)

}

// GetAllBlog godoc
func (bc *BlogController) GetAllBlog(c *gin.Context) {
	// implementation
	blogs, err := bc.blogUseCase.GetAllBlog(bc.c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blogs)
}

// UpdateBlog godoc
func (bc *BlogController) UpdateBlog(c *gin.Context) {
	// implementation
	blogID := c.Param("blog_id")
	var blog domain.BlogUpdate
	err := c.ShouldBindJSON(&blog)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	blogModel, err := bc.blogUseCase.UpdateBlog(bc.c, &blog, blogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blogModel)
}

// DeleteBlog godoc
func (bc *BlogController) DeleteBlog(c *gin.Context) {
	// implementation
	blogID := c.Param("blog_id")
	err := bc.blogUseCase.DeleteBlog(bc.c, blogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
