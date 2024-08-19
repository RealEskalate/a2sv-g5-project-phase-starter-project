package controller

import (
	"context"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain"
	"github.com/gin-gonic/gin"
)

// interface for blog controllers
type blogController interface {
	GetBlogs() gin.HandlerFunc
	GetBlog() gin.HandlerFunc
	CreateBlog() gin.HandlerFunc
	UpdateBlog() gin.HandlerFunc
	DeleteBlog() gin.HandlerFunc
	// GetComments() gin.HandlerFunc
	// CreateComment() gin.HandlerFunc
	// GetComment() gin.HandlerFunc
	// UpdateComment() gin.HandlerFunc
	// DeleteComment() gin.HandlerFunc
	// CreateLike() gin.HandlerFunc
}

type BlogController struct {
	BlogUsecase domain.BlogUsecase
	Env         *bootstrap.Env
}

func (bc *BlogController) GetBlogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		blogs, err := bc.BlogUsecase.GetAllBlogs(context.Background())
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, blogs)
	}
}

func (bc *BlogController) GetBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		blogID := c.Param("id")
		blog, err := bc.BlogUsecase.GetBlogByID(context.Background(), blogID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, blog)
	}
}

func (bc *BlogController) CreateBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newBlog domain.Blog
		if err := c.ShouldBindJSON(&newBlog); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		blog, err := bc.BlogUsecase.CreateBlog(context.Background(), &newBlog)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, blog)
	}
}

func (bc *BlogController) UpdateBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		blogID := c.Param("id")
		var updatedBlog domain.Blog
		if err := c.ShouldBindJSON(&updatedBlog); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		blog, err := bc.BlogUsecase.UpdateBlog(context.Background(), blogID, &updatedBlog)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, blog)
	}
}

func (bc *BlogController) DeleteBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		blogID := c.Param("id")
		err := bc.BlogUsecase.DeleteBlog(context.TODO(), blogID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(204, nil)
	}
}
