package controller

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogController struct {
    blogService interfaces.BlogService
}

func NewBlogController(blogService interfaces.BlogService) *BlogController {
    return &BlogController{
        blogService: blogService,
    }
}

func (bc *BlogController) CreateBlogPost(c *gin.Context) {
    
	var blogPost entities.BlogPost
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}
	userIdStr, ok := userId.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while finding user"})
		return
	}
    if err := c.ShouldBindJSON(&blogPost); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    createdBlogPost, err := bc.blogService.CreateBlogPost(&blogPost, userIdStr)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Blog post created successfully", "blogPost": createdBlogPost})
}


func (bc *BlogController) GetBlogPosts(c *gin.Context) {
    // Parse query parameters for pagination
    pageStr := c.DefaultQuery("page", "1")
    pageSizeStr := c.DefaultQuery("pageSize", "20")
    sortBy := c.DefaultQuery("sortBy", "createdAt")

    page, err := strconv.Atoi(pageStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
        return
    }

    pageSize, err := strconv.Atoi(pageSizeStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
        return
    }

    blogPosts, totalPosts, err := bc.blogService.GetBlogPosts(page, pageSize, sortBy)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Calculate pagination metadata
    totalPages := (totalPosts + pageSize - 1) / pageSize

    // Return the response with blog posts and pagination metadata
    c.JSON(http.StatusOK, gin.H{
        "blogPosts": blogPosts,
        "pagination": gin.H{
            "currentPage": page,
            "pageSize":    pageSize,
            "totalPages":  totalPages,
            "totalPosts":  totalPosts,
        },
    })
}

func (bc *BlogController) UpdateBlogPost(c *gin.Context) {
	// Parse the blog post ID from the URL
	blogPostId := c.Param("id")
	userId := c.GetString("userId")
	objID, err := primitive.ObjectIDFromHex(blogPostId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog post ID"})
		return
	}

	// Bind the incoming JSON to the blogPost entity
	var blogPost entities.BlogPost
	if err := c.ShouldBindJSON(&blogPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Set the ID to the object ID from the URL
	blogPost.ID = objID

	// Update the blog post
	updatedBlogPost, err := bc.blogService.UpdateBlogPost(&blogPost,userId)
	if err != nil {
		if errors.Is(err, errors.New("unauthorized: only the author can update this post")) {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to update this blog post"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the updated blog post as confirmation
	c.JSON(http.StatusOK, gin.H{"message": "Blog post updated successfully", "blogPost": updatedBlogPost})
}



func (bc *BlogController) DeleteBlogPost(c *gin.Context) {
	// Parse the blog post ID from the URL
	blogPostId := c.Param("id")
	userId:= c.GetString("userId")
	role := c.GetString("role")	

	// Delete the blog post
	err := bc.blogService.DeleteBlogPost(blogPostId, userId,role)
	if err != nil {
		if errors.Is(err, errors.New("unauthorized: only the author or an admin can delete this post")) {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to delete this blog post"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "Blog post deleted successfully"})
}


func (bc *BlogController) SearchBlogPosts(c *gin.Context) {
	criteria := c.Query("criteria")
	tags := c.QueryArray("tags")

	startDateStr := c.Query("startDate")
	endDateStr := c.Query("endDate")

	var startDate, endDate time.Time
	var err error

	if startDateStr != "" {
		startDate, err = time.Parse(time.RFC3339, startDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format"})
			return
		}
	}

	if endDateStr != "" {
		endDate, err = time.Parse(time.RFC3339, endDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format"})
			return
		}
	}

	blogPosts, err := bc.blogService.SearchBlogPosts(criteria, tags, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogPosts)
}