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
	page, err1 := strconv.Atoi(c.Query("page"))
	limit, err2 := strconv.Atoi(c.Query("limit"))
	if err1 != nil {
		page = 1
	}
	if err2 != nil {
		limit = 10
	}

	filterMap := make(map[string]interface{})
	for _, filter := range filters {
		keyValue := strings.SplitN(filter, ":", 2)
		if len(keyValue) == 2 {
			filterMap[keyValue[0]] = keyValue[1]
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

	err := bc.blog_usecase.DeleteBlog(blogID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete blog"})
		return
	}

	c.JSON(200, gin.H{"message": "Blog deleted successfully"})
}

func (bc *BlogController) LikeBlog(c *gin.Context) {
	var requestBody struct {
		UserID string `json:"user_id"`
	}
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	blogID := c.Param("id")
	if err := bc.blog_usecase.LikeBlog(blogID, requestBody.UserID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Blog liked successfully"})
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
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
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
