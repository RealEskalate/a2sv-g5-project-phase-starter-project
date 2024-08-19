package controllers

import (
	"blog_api/domain"
	"net/http"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BlogController struct {
	blogUseCase domain.BlogUseCaseInterface
}

var validate = validator.New()

func NewBlogController(bu domain.BlogUseCaseInterface) *BlogController {
	return &BlogController{
		blogUseCase: bu,
	}
}

// CreateBlogHandler handles the HTTP request for creating a new blog post.
func (bc *BlogController) CreateBlogHandler(c *gin.Context) {
	var blog domain.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{"error": err.Error()})
		return
	}
	err := validate.Struct(blog)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{"error": err.Error()})
		return
	}

	newErr := bc.blogUseCase.CreateBlogPost(c, &blog)
	if newErr != nil {
		c.JSON(GetHTTPErrorCode(newErr), domain.Response{"error": newErr.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.Response{"message": "blog created successfully"})
}

// UpdateBlogHandler handles the HTTP request to update a blog post.
func (bc *BlogController) UpdateBlogHandler(c *gin.Context) {
	blogId := c.Param("id")
	var blog domain.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{"error": err.Error()})
		return
	}
	err := bc.blogUseCase.EditBlogPost(c, blogId, &blog)
	if err != nil {
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, domain.Response{"message": "created successfuly"})
}

// DeleteBlogHandler handles the HTTP DELETE request to delete a blog post.
func (bc *BlogController) DeleteBlogHandler(c *gin.Context) {
	blogId := c.Param("id")

	err := bc.blogUseCase.DeleteBlogPost(c, blogId)
	if err != nil {
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.Response{"message": "deleted successfuly"})
}

// GetBlogHandler handles the HTTP GET request to retrieve a list of blog posts based on filters.
func (bc *BlogController) GetBlogHandler(c *gin.Context) {
	var filters domain.BlogFilterOptions
	if err := c.ShouldBindQuery(&filters); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{"error": "Invalid query parameters"})
		return
	}

	blogs, total, err := bc.blogUseCase.GetBlogPosts(c, filters)
	if err != nil {
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"total": total, "blogs": blogs})
}

// GetBlogByIDHandler handles the HTTP GET request to retrieve a single blog post by its ID.
func (bc *BlogController) GetBlogByIDHandler(c *gin.Context) {
	blogId := c.Param("id")

	blog, err := bc.blogUseCase.GetBlogPostByID(c, blogId)
	if err != nil {
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blog)
}

// TrackBlogPopularityHandler handles the HTTP POST request to track the popularity of a blog post.
func (bc *BlogController) TrackBlogPopularityHandler(c *gin.Context) {
	var requestBody struct {
		BlogID   string `json:"blogID" validate:"required"`
		Action   string `json:"action" validate:"required,oneof=like dislike"`
		Username string `json:"username" validate:"required"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{"error": "Invalid input data"})
		return
	}

	err := bc.blogUseCase.TrackBlogPopularity(c, requestBody.BlogID, requestBody.Action, requestBody.Username)
	if err != nil {
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.Response{"message": "Action applied successfully"})
}

// GenerateContentHandler handles requests to generate content
func (bc *BlogController) GenerateContentHandler(c *gin.Context) {
	var req struct {
		Topics []string `json:"topics"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error decoding request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	content, err := bc.blogUseCase.GenerateBlogContent(req.Topics)
	if err != nil {
		log.Printf("Error generating blog content: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate content"})
		return
	}

	response := gin.H{"content": content}
	c.JSON(http.StatusOK, response)
}

// ReviewContentHandler handles requests to review content
func (bc *BlogController) ReviewContentHandler(c *gin.Context) {
	var req struct {
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error decoding request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	suggestions, err := bc.blogUseCase.ReviewBlogContent(req.Content)
	if err != nil {
		log.Printf("Error generating suggestions: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to review content"})
		return
	}

	response := gin.H{"suggestions": suggestions}
	c.JSON(http.StatusOK, response)
}

// GenerateTopicHandler handles requests to generate topics
func (bc *BlogController) GenerateTopicHandler(c *gin.Context) {
	var req struct {
		Keywords []string `json:"keywords"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error decoding request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	topics, err := bc.blogUseCase.GenerateTrendingTopics(req.Keywords)
	if err != nil {
		log.Printf("Error generating trending topics: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate topics"})
		return
	}

	response := gin.H{"topics": topics}
	c.JSON(http.StatusOK, response)
}
