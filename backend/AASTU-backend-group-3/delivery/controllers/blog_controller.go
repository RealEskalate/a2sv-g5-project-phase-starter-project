package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"group3-blogApi/domain"
)

type BlogController struct {
	blogUsecase domain.BlogUsecase
}

func NewBlogController(blogUsecase domain.BlogUsecase) *BlogController {
	return &BlogController{
		blogUsecase: blogUsecase,
	}
}

func (c *BlogController) CreateBlog(ctx *gin.Context) {
	var blog domain.Blog
	if err := ctx.ShouldBindJSON(&blog); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := ctx.GetString("user_id")
	username := ctx.GetString("username")

	newBlog, err := c.blogUsecase.CreateBlog(username, userID, blog)
	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":      "Blog created successfully",
		"Created_blog": newBlog,
	})
}

func (c *BlogController) DeleteBlog(ctx *gin.Context) {
	id := ctx.Param("id")
	role := ctx.GetString("role")
	userId := ctx.GetString("user_id")

	newBlog, err := c.blogUsecase.DeleteBlog(role, userId, id)
	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":      "Blog deleted successfully",
		"Deleted Blog": newBlog,
	})
}

func (c *BlogController) UpdateBlog(ctx *gin.Context) {
	id := ctx.Param("id")

	var blog domain.Blog
	if err := ctx.ShouldBindJSON(&blog); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	blog.AuthorID = ctx.GetString("user_id")

	if _, err := c.blogUsecase.UpdateBlog(blog, ctx.GetString("role"), id); err != nil {
		ctx.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":      "Blog updated successfully",
		"Updated Blog": blog,
	})
}

func (c *BlogController) GetBlogByID(ctx *gin.Context) {
	id := ctx.Param("id")

	blog, err := c.blogUsecase.GetBlogByID(id)
	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Blog retrieved successfully",
		"blog":    blog,
	})
}

func (c *BlogController) GetBlogs(ctx *gin.Context) {
	var page int64 = 1   // Default to page 1
	var limit int64 = 2 // Default to limit 2
	var sortBy string
	var tag string
	var authorName string

	if p := ctx.Query("page"); p != "" {
		if parsedPage, err := strconv.ParseInt(p, 10, 64); err == nil {
			page = parsedPage
		}
	}

	if l := ctx.Query("limit"); l != "" {
		if parsedLimit, err := strconv.ParseInt(l, 10, 64); err == nil {
			limit = parsedLimit
		}
	}

	sortBy = ctx.Query("sortBy")
	tag = ctx.Query("tag")
	authorName = ctx.Query("authorName")

	blogs, err := c.blogUsecase.GetBlogs(page, limit, sortBy, tag, authorName)
	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Blogs retrieved successfully",
		"blogs":   blogs,
	})
}

func (c *BlogController) GetUserBlogs(ctx *gin.Context) {
	userID := ctx.Param("id")

	blogs, err := c.blogUsecase.GetUserBlogs(userID)
	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User blogs retrieved successfully",
		"blogs":   blogs,
	})
}
