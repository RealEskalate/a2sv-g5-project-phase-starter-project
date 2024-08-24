package controllers

import (
	"AAIT-backend-group-3/internal/domain/models"
	"AAIT-backend-group-3/internal/usecases"
	"strconv"
	"strings"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BlogControllerInterface interface {
	CreateBlog(c *gin.Context)
	GetBlogByID(c *gin.Context)
	GetBlogs(c *gin.Context)
	GetBlogsByAuthorID(c *gin.Context)
	GetBlogsByPopularity(c *gin.Context)
	GetBlogsByTags(c *gin.Context)
	UpdateBlog(c *gin.Context)
	DeleteBlog(c *gin.Context)
	LikeBlog(c *gin.Context)
	ViewBlog(c *gin.Context)
}

type BlogController struct {
	blog_usecase usecases.BlogUsecaseInterface
}

func NewBlogController(u usecases.BlogUsecaseInterface) BlogControllerInterface {
	return &BlogController{
		blog_usecase: u,
	}
}

func (bc *BlogController) CreateBlog(c *gin.Context) {
	var req models.Blog
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userId not found"})
		return
	}
	blog, err := bc.blog_usecase.CreateBlog(&req, userId.(string))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"blog_id": blog})
}

func (bc *BlogController) GetBlogByID(c *gin.Context) {
	blogID := c.Param("id")

	existingBlog, err := bc.blog_usecase.GetBlogByID(blogID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get blog"})
		return
	}
	
	c.JSON(200, gin.H{"blog": existingBlog})
}

func (bc *BlogController) GetBlogs(c *gin.Context) {
	filters := c.QueryMap("filter")
	search := c.Query("search")
	page, err1 := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, err2 := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if err1 != nil || page < 1 {
		page = 1
	}
	if err2 != nil || limit < 1 {
		limit = 10
	}

	filterMap := make(map[string]interface{})
	for key, filter := range filters {
		if key == "tags" {
			filterMap[key] = strings.Split(filter, ",")
		} else {
			filterMap[key] = filter
		}
	}

	blogs, err := bc.blog_usecase.GetBlogs(filterMap, search, page, limit)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get blogs"})
		return
	}

	c.JSON(200, gin.H{"blogs": blogs})
}

func (bc *BlogController) UpdateBlog(c *gin.Context) {
    blogID := c.Param("id")
    var req models.Blog
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    existingBlog, err := bc.blog_usecase.GetBlogByID(blogID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
        return
    }

    userID := c.GetString("userId")
    if existingBlog.AuthorID.Hex() != userID {
        c.JSON(http.StatusForbidden, gin.H{"error": "You can only update your own blog"})
        return
    }

    err = bc.blog_usecase.UpdateBlog(blogID, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Blog updated successfully"})
}

func (bc *BlogController) DeleteBlog(c *gin.Context) {
	blogID := c.Param("id")
	userID := c.GetString("userId")
	blog, err := bc.blog_usecase.GetBlogByID(blogID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get blog"})
		return
	}
	if blog.AuthorID.Hex() != userID {
		c.JSON(403, gin.H{"error": "You can only delete your own blog"})
		return
	}
	err = bc.blog_usecase.DeleteBlog(blogID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete blog"})
		return
	}
	c.JSON(200, gin.H{"message": "Blog deleted successfully"})
}

func (bc *BlogController) LikeBlog(c *gin.Context) {
    userID := c.GetString("userId")
    blogID := c.Param("id")
    
    liked, err := bc.blog_usecase.ToggleLike(blogID, userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if liked {
        c.JSON(http.StatusOK, gin.H{"message": "Blog liked successfully", "liked": true})
    } else {
        c.JSON(http.StatusOK, gin.H{"message": "Blog disliked successfully", "liked": false})
    }
}


func (bc *BlogController) ViewBlog(c *gin.Context) {
	blogID := c.Param("id")
	if err := bc.blog_usecase.ViewBlog(blogID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Blog view recorded successfully"})
}

func (bc *BlogController) GetBlogsByAuthorID(c *gin.Context) {
	authorID := c.Param("author_id")
	blogs, err := bc.blog_usecase.GetBlogsByAuthorID(authorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get blogs by author"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"blogs": blogs})
}

func (bc *BlogController) GetBlogsByPopularity(c *gin.Context) {
	limitStr := c.Query("limit")
	limit := 10 // Default value
	if limitStr != "" {
		var err error
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
			return
		}
	}
	blogs, err := bc.blog_usecase.GetBlogsByPopularity(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get popular blogs"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"blogs": blogs})
}

func (bc *BlogController) GetBlogsByTags(c *gin.Context) {
	tags := c.QueryArray("tags")
	if len(tags) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tags parameter is required"})
		return
	}
	blogs, err := bc.blog_usecase.GetBlogsByTags(tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get blogs by tags"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"blogs": blogs})
}
