package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"aait.backend.g10/domain"
	"aait.backend.g10/usecases"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type BlogController struct {
	BlogUseCase usecases.IBlogUseCase
}

func NewBlogController(b usecases.IBlogUseCase) *BlogController {
	return &BlogController{
		BlogUseCase: b,
	}
}

func (b *BlogController) CreateBlog(c *gin.Context) {
	var blog domain.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		var validationErrors validator.ValidationErrors
		if errors, ok := err.(validator.ValidationErrors); ok {
			validationErrors = errors
		}
		errorMessages := make(map[string]string)
		for _, e := range validationErrors {

			field := e.Field()
			switch field {
			case "Title":
				errorMessages["title"] = "Title is required."
			case "Content":
				errorMessages["content"] = "Content is required."
			case "Author":
				errorMessages["author"] = "Author is required."
			case "Tags":
				errorMessages["tags"] = "Tags is required."
			}
		}

		if len(errorMessages) == 0 {
			errorMessages["json"] = "Invalid JSON"
		}

		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	id, err := uuid.Parse(c.MustGet("id").(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	blog.Author = id
	newBlog, cerr := b.BlogUseCase.CreateBlog(&blog)
	if cerr != nil {
		c.JSON(cerr.StatusCode, gin.H{"error": cerr.Error()})
		return
	}

	c.JSON(http.StatusCreated, newBlog)
}

func (b *BlogController) GetAllBlogs(c *gin.Context) {
	blogs, err := b.BlogUseCase.GetAllBlogs()
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

func (b *BlogController) GetBlogByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	blog, cerr := b.BlogUseCase.GetBlogByID(id)
	if cerr != nil {
		c.JSON(cerr.StatusCode, gin.H{"error": cerr.Error()})
		return
	}

	c.JSON(http.StatusOK, blog)
}

func (b *BlogController) UpdateBlog(c *gin.Context) {
	blogId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var blog domain.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		var validationErrors validator.ValidationErrors
		if errors, ok := err.(validator.ValidationErrors); ok {
			validationErrors = errors
		}

		errorMessages := make(map[string]string)
		for _, e := range validationErrors {

			field := e.Field()
			switch field {
			case "Title":
				errorMessages["title"] = "Title is required."
			case "Content":
				errorMessages["content"] = "Content is required."

			case "Author":
				errorMessages["author"] = "Author is required."

			case "Tags":
				errorMessages["tags"] = "Tags is required."
			}
		}

		if len(errorMessages) == 0 {
			errorMessages["json"] = "Invalid JSON"
		}

		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}
	requester_id, err := uuid.Parse(c.MustGet("id").(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	blog.ID = blogId
	blog.Author = requester_id
	cerr := b.BlogUseCase.UpdateBlog(&blog)
	if cerr != nil {
		c.JSON(cerr.StatusCode, gin.H{"error": cerr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog updated successfully"})
}

func (b *BlogController) DeleteBlog(c *gin.Context) {
	blogId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	requester_id, err := uuid.Parse(c.MustGet("id").(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	is_admin := c.MustGet("is_admin").(bool)

	cerr := b.BlogUseCase.DeleteBlog(blogId, requester_id, is_admin)
	if cerr != nil {
		c.JSON(cerr.StatusCode, gin.H{"error": cerr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}

func (b *BlogController) AddView(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	cerr := b.BlogUseCase.AddView(id)
	if cerr != nil {
		c.JSON(cerr.StatusCode, gin.H{"error": cerr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "View added successfully"})
}

func (b *BlogController) SearchBlogs(c *gin.Context) {
	filter := domain.BlogFilter{
		Title:  c.Query("title"),
		SortBy: c.Query("sortBy"),
		Author: c.Query("author"),
	}

	tagsParam := c.Query("tags")
	if tagsParam != "" {
		// Split the tags into a slice of strings
		tags := strings.Split(tagsParam, ",")

		// Trim spaces from each tag
		for i := range tags {
			tags[i] = strings.TrimSpace(tags[i])
		}

		filter.Tags = tags
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		page = 1 // Default to the first page if not provided or invalid
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil || limit <= 0 {
		limit = 20 // Default to 20 items per page if not provided or invalid
	}
	filter.Page = page
	filter.PageSize = limit

	blogs, totalPages, totalCount, cerr := b.BlogUseCase.SearchBlogs(filter)
	if cerr != nil {
		c.JSON(cerr.StatusCode, gin.H{"error": cerr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"blogs":       blogs,
		"totalPages":  totalPages,
		"currentPage": page,
		"totalCount":  totalCount,
	})
}
