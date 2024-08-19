package controller

import (
	"Blog_Starter/domain"
	"Blog_Starter/utils"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	blogUseCase domain.BlogUseCase
	ctx         context.Context
}

func NewBlogController(blogUseCase domain.BlogUseCase, ctx context.Context) *BlogController {
	return &BlogController{
		blogUseCase: blogUseCase,
		ctx:         ctx,
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
	blogModel, err := bc.blogUseCase.CreateBlog(bc.ctx, &blog)
	if err != nil {
		// Check for specific errors and return appropriate status codes
		if err.Error() == "content length should be greater than 10" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusCreated, blogModel)
}

// GetBlogByID godoc
func (bc *BlogController) GetBlogByID(c *gin.Context) {
	// implementation create a context and pass to the usecase not the gin context
	blogID := c.Param("blog_id")
	blog, err := bc.blogUseCase.GetBlogByID(bc.ctx, blogID)
	if err != nil {
		// Check for specific errors and return appropriate status codes
		if err.Error() == "invalid blog id" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else if err.Error() == "blog not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, blog)
}

// GetAllBlog godoc
func (bc *BlogController) GetAllBlog(c *gin.Context) {
	skipStr := c.Query("skip")
	limitStr := c.Query("limit")
	skip, _ := strconv.ParseInt(skipStr, 10, 64)
	limit, _ := strconv.ParseInt(limitStr, 10, 64)
	sortBy := c.Query("sort_by")
	// implementation
	blogs, paginationMetadata, err := bc.blogUseCase.GetAllBlog(bc.ctx, skip, limit, sortBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"blogs": blogs, "metadata": paginationMetadata})
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
	user, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	blog.UserID = user.UserID
	// call the useCase getBlogByID to check whether the blog exists or not

	blogModel, err := bc.blogUseCase.UpdateBlog(bc.ctx, &blog, blogID)
	if err != nil {
		// Check for specific errors and return appropriate status codes
		if err.Error() == "blog not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else if err.Error() == "user is not the owner of the blog" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, blogModel)
}

// DeleteBlog godoc
func (bc *BlogController) DeleteBlog(c *gin.Context) {
	// implementation
	blogID := c.Param("blog_id")
	user, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	err = bc.blogUseCase.DeleteBlog(bc.ctx, blogID, user.UserID)
	if err != nil {
		// Check for specific errors and return appropriate status codes
		if err.Error() == "blog not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else if err.Error() == "user is not the owner of the blog" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
