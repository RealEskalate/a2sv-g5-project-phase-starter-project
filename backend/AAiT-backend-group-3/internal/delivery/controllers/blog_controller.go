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
	claims, _ := c.Get("claims")
	authorID := claims.(map[string]interface{})["id"].(string)

	blog, err := bc.blog_usecase.CreateBlog(&req, authorID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create blog"})
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
	if err1 != nil || err2 != nil {
		c.JSON(400, gin.H{"error": "Invalid limit or page number"})
		return
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

    userID := c.GetString("userID")
    if existingBlog.AuthorID.Hex() != userID {
        c.JSON(http.StatusForbidden, gin.H{"error": "You can only update your own blog"})
        return
    }


    err = bc.blog_usecase.UpdateBlog(blogID, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit blog"})
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
