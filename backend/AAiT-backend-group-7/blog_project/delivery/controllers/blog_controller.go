package controllers

import (
	"blog_project/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	BlogUsecase domain.IBlogUsecase
}

func NewBlogController(blogUsecase domain.IBlogUsecase) domain.IBlogController {
	return &BlogController{BlogUsecase: blogUsecase}
}

func (bc *BlogController) GetAllBlogs(c *gin.Context) {
	sortOrder := c.DefaultQuery("sort", "DESC") // Default to "DESC" if not specified
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid page number"})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid limit number"})
		return
	}

	blogs, err := bc.BlogUsecase.GetAllBlogs(c, sortOrder, page, limit)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, blogs)
}

func (bc *BlogController) CreateBlog(c *gin.Context) {
	var blog domain.Blog
	if err := c.BindJSON(&blog); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newBlog, err := bc.BlogUsecase.CreateBlog(c, blog)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, newBlog)
}

func (bc *BlogController) UpdateBlog(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid blog ID"})
		return
	}

	var blog domain.Blog
	if err := c.BindJSON(&blog); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	updatedBlog, err := bc.BlogUsecase.UpdateBlog(c, id, blog)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, updatedBlog)
}

func (bc *BlogController) DeleteBlog(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid blog ID"})
		return
	}

	if err := bc.BlogUsecase.DeleteBlog(c, id); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Blog deleted successfully"})
}

func (bc *BlogController) LikeBlog(c *gin.Context) {
	blogID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid blog ID"})
		return
	}

	likedBlog, err := bc.BlogUsecase.LikeBlog(c, blogID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, likedBlog)
}

func (bc *BlogController) DislikeBlog(c *gin.Context) {
	blogID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid blog ID"})
		return
	}

	dislikedBlog, err := bc.BlogUsecase.DislikeBlog(c, blogID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, dislikedBlog)
}

func (bc *BlogController) AddComment(c *gin.Context) {
	blogID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid blog ID"})
		return
	}

	var comment domain.Comment
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	commentedBlog, err := bc.BlogUsecase.AddComment(c, blogID, comment.Content)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, commentedBlog)
}

func (bc *BlogController) Search(c *gin.Context) {
	author := c.Query("author")
	tags := c.QueryArray("tags")
	title := c.Query("title")

	blogs, err := bc.BlogUsecase.Search(c, author, tags, title)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, blogs)
}

type BlogContent struct {
	Content string `json:"content"`
}

func (bc *BlogController) AiRecommendation(c *gin.Context) {
	var content BlogContent
	if err := c.ShouldBindJSON(&content); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	recommendation, err := bc.BlogUsecase.AiRecommendation(c, content.Content)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, recommendation)
}
