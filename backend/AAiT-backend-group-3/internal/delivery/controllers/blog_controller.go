package controllers

import (
	"AAIT-backend-group-3/internal/domain/models"
	"AAIT-backend-group-3/internal/usecases"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)


type BlogController struct {
	blog_usecase *usecases.BlogUsecase
}


func NewBlogController(u *usecases.BlogUsecase) *BlogController{
	return &BlogController{
		blog_usecase: u,
	}
}

// @Todo: set the AuthorID for the blog, get the author from the token
func (blogController *BlogController) CreateBlog(c *gin.Context) {
	var req models.Blog

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// get the author from the token
	claims, _ := c.Get("claims")
	authorID := claims.(map[string]interface{})["id"].(string)

	err := blogController.blog_usecase.CreateBlog(&req, authorID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create blog"})
		return
	}

	c.JSON(200, gin.H{"message": "Blog created"})
}

func (blogController *BlogController) GetBlogByID(c *gin.Context) {
	blogId := c.Param("id")

	existingBlog, err := blogController.blog_usecase.GetBlogByID(blogId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get blog"})
		return
	}

	c.JSON(200, gin.H{"blog": existingBlog})
}

func (blogController *BlogController) GetBlogs(c *gin.Context) {

	filters := c.QueryMap("filter")
	search := c.Query("search")
	page, err := strconv.Atoi(c.Query("page"))
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil{
		c.JSON(500, gin.H{"error": "Invalid limit or page number"})
		return
	}
	filterMap := make(map[string]interface{})

	for _, filter := range filters {
		keyValue := strings.SplitN(filter, ":", 2)
		if len(keyValue) == 2 {
			filterMap[keyValue[0]] = keyValue[1]
		}
	}

	blogs, err := blogController.blog_usecase.GetBlogs(filterMap, search, page, limit)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get blogs"})
		return
	}

	c.JSON(200, gin.H{"blogs": blogs})
}

func (blogController *BlogController) EditBlog(c *gin.Context) {
	blogId := c.Param("id")
	var req models.Blog

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := blogController.blog_usecase.EditBlog(blogId, &req)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to edit blog"})
		return
	}

	c.JSON(200, gin.H{"message": "Blog edited"})
}

func (blogController *BlogController) DeleteBlog(c *gin.Context) {
	blogId := c.Param("id")

	err := blogController.blog_usecase.DeleteBlog(blogId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete blog"})
		return
	}

	c.JSON(200, gin.H{"message": "Blog deleted"})
}
