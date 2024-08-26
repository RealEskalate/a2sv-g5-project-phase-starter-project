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

	newBlog, uerr := c.blogUsecase.CreateBlog(username, userID, blog)
	if uerr.Message != "" {
		ctx.JSON(uerr.StatusCode, gin.H{"error": uerr.Message})
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

	newBlog, uerr := c.blogUsecase.DeleteBlog(role, userId, id)
	if uerr.Message != "" {
		ctx.JSON(uerr.StatusCode, gin.H{"error": uerr.Message})
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

	if _, uerr := c.blogUsecase.UpdateBlog(blog, ctx.GetString("role"), id); uerr.Message != "" {
		ctx.JSON(uerr.StatusCode, gin.H{"error": uerr.Message})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":      "Blog updated successfully",
		"Updated Blog": blog,
	})
}

func (c *BlogController) GetBlogByID(ctx *gin.Context) {
	id := ctx.Param("id")

	blog, uerr := c.blogUsecase.GetBlogByID(id)
	if uerr.Message != "" {
		ctx.JSON(uerr.StatusCode, gin.H{"error": uerr.Message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Blog retrieved successfully",
		"blog":    blog,
	})
}

func (c *BlogController) GetBlogs(ctx *gin.Context) {
	var page int64 = 1   // Default to page 1
	var limit int64 = 20 // Default to limit 2
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

	blogs, total, uerr := c.blogUsecase.GetBlogs(page, limit, sortBy, tag, authorName)
	if uerr.Message != "" {
		ctx.JSON(uerr.StatusCode, gin.H{"error": uerr.Message})
		return
	}

	totalPages := (total + limit - 1) / limit 

    ctx.JSON(http.StatusOK, gin.H{
        "message":       "Blogs retrieved successfully",
        "blogs":         blogs,
        "current_page":  page,
        "per_page":      limit,
        "total_records": total,
        "total_pages":   totalPages,
    })
}

func (c *BlogController) GetUserBlogs(ctx *gin.Context) {
	userID := ctx.Param("id")

	blogs, uerr := c.blogUsecase.GetUserBlogs(userID)
	if uerr.Message != "" {
		ctx.JSON(uerr.StatusCode, gin.H{"error": uerr.Message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User blogs retrieved successfully",
		"blogs":   blogs,
	})
}
