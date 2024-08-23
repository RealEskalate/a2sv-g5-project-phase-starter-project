package controller

import (
	"Blog_Starter/domain"
	"Blog_Starter/utils"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogController struct {
	blogUseCase 		domain.BlogUseCase
}

func NewBlogController(blogUseCase domain.BlogUseCase) *BlogController {
	return &BlogController{
		blogUseCase: blogUseCase,
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

	blogModel, err := bc.blogUseCase.CreateBlog(c, &blog)
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
	blog, err := bc.blogUseCase.GetBlogByID(c, blogID)
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
	blogs, paginationMetadata, err := bc.blogUseCase.GetAllBlog(c, skip, limit, sortBy)
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
	//call the useCase getBlogByID to check whether the blog exists or not
	blogModel, err := bc.blogUseCase.UpdateBlog(c, &blog, blogID)
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

	err = bc.blogUseCase.DeleteBlog(c, blogID, user.UserID, user.Role)
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

func (bc *BlogController) FilterBlog(c *gin.Context) {
	var filterReq domain.BlogFilterRequest
	if err := c.BindJSON(&filterReq); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error" : "invalid request format"})
		return
	}

	filtrationResponse, err := bc.blogUseCase.FilterBlogs(c, &filterReq)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error" : "no matches found"})
			return
		}
		if err.Error() == "invalid request format" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error" : "internal server error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"filtered_blogs" : filtrationResponse})
}

func (bc *BlogController) SearchBlog(c *gin.Context) {
	author := c.Query("author")
	title := c.Query("title")
	if title == "" && author == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error" : "invalid request format"})
		return
	}

	var searchRequest domain.BlogSearchRequest
	searchRequest.Author = author
	searchRequest.Title = title
	searchResult, err := bc.blogUseCase.SearchBlogs(c, &searchRequest)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error" : "no matches found"})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error" : "internal server error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"search_result" : searchResult})
		
}