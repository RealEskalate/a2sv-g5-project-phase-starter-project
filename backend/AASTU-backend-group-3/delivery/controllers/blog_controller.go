package controllers

import (
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"group3-blogApi/domain"
	"strconv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"fmt"
)

type BlogController struct {
	blogUsecase domain.BlogUsecase
}

func NewBlogController(blogUsecase domain.BlogUsecase) *BlogController {
	return &BlogController{
		blogUsecase: blogUsecase,
	}
}

// Public Routes

// GetBlogs retrieves all blogs with pagination and sorting
func (c *BlogController) GetBlogs(ctx *gin.Context) {
    offset, err := strconv.ParseInt(ctx.Query("offset"), 10, 64)
    if err != nil {
        offset = 0
    }

    limit, err := strconv.ParseInt(ctx.Query("limit"), 10, 64)
    if err != nil {
        limit = 10
    }

    sortBy := ctx.Query("sortBy")
    if sortBy == "" {
        sortBy = "created_at" // default sort field
    }

    blogs, err := c.blogUsecase.GetBlogs(ctx, offset, limit, sortBy)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get blogs"})
        return
    }

    ctx.JSON(http.StatusOK, blogs)
}

// GetBlog retrieves a single blog by its ID
func (bc *BlogController) GetBlog(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        fmt.Println("Invalid ObjectID:", err)
    } else {
        fmt.Println("Valid ObjectID:", objID)
    }
	

	blog, err := bc.blogUsecase.GetBlogByID(context.Background(), id)
	if err != nil {
		if err == domain.ErrBlogNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, blog)
}

// GetUserBlogs retrieves all blogs by a specific user
func (bc *BlogController) GetUserBlogs(c *gin.Context) {
	userID := c.Param("id")

	//need change           -----------------------------------------------


	_, _, sortBy := parsePaginationAndSortingParams(c)

	filters := map[string]interface{}{
		"user_id": userID,
	}

	blogs, err := bc.blogUsecase.FilterBlogs(context.Background(), filters, sortBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

// SearchBlogs allows users to search for blogs
func (bc *BlogController) SearchBlogs(c *gin.Context) {
	query := c.Param("q")
	filters := parseFilters(c)

	blogs, err := bc.blogUsecase.SearchBlogs(context.Background(), query, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

// FilterBlogs allows users to filter blogs based on criteria
func (bc *BlogController) FilterBlogs(c *gin.Context) {
	filters := parseFilters(c)
	sortBy := c.Query("sort_by")

	blogs, err := bc.blogUsecase.FilterBlogs(context.Background(), filters, sortBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

// Admin Routes

// CreateBlog allows an admin to create a new blog
func (c *BlogController) CreateBlog(ctx *gin.Context) {
    // Extract the token from the context (assuming JWT middleware is used)
    user, exists := ctx.Get("user_id")
    if !exists {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    // Assuming `user` is a struct that contains the user's ID
    userId := user.(string) // or user.(User).ID if it's a custom struct

    var blog domain.Blog
    if err := ctx.ShouldBindJSON(&blog); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	blog.ID = primitive.NewObjectID()

    // Set the AuthorID from the token
    blog.AuthorID = userId

    // Call the service or repository layer to create the blog
    blogID, err := c.blogUsecase.CreateBlog(ctx, blog)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create blog"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"id": blogID})
}


// UpdateBlog allows an admin to update a blog
func (bc *BlogController) UpdateBlog(c *gin.Context) {
	id := c.Param("id")

	var blog domain.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blog.ID,_ = convertStringToObjectID(id)
	if err := bc.blogUsecase.UpdateBlog(context.Background(), blog); err != nil {
		if err == domain.ErrBlogNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog updated successfully"})
}

// DeleteBlog allows an admin to delete a blog
func (bc *BlogController) DeleteBlog(c *gin.Context) {
	id := c.Param("id")

	if err := bc.blogUsecase.DeleteBlog(context.Background(), id); err != nil {
		if err == domain.ErrBlogNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}

// UpdateBlogVisibility allows an admin to update the visibility of a blog
func (bc *BlogController) UpdateBlogVisibility(c *gin.Context) {
	id := c.Param("id")

	visibility := c.Query("visibility")
	if visibility == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Visibility parameter is required"})
		return
	}

	blog, err := bc.blogUsecase.GetBlogByID(context.Background(), id)
	if err != nil {
		if err == domain.ErrBlogNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	blog.Visibility = visibility
	if err := bc.blogUsecase.UpdateBlog(context.Background(), *blog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog visibility updated successfully"})
}

// Helper functions to parse pagination, sorting, and filtering parameters

func parsePaginationAndSortingParams(c *gin.Context) (int64, int64, string) {
	offset, _ := strconv.ParseInt(c.Query("offset"), 10, 64)
	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
	sortBy := c.Query("sort_by")
	return offset, limit, sortBy
}

func parseFilters(c *gin.Context) map[string]interface{} {
	filters := make(map[string]interface{})
	for key, values := range c.Request.URL.Query() {
		if key != "q" && key != "sort_by" && key != "offset" && key != "limit" {
			filters[key] = values[0]
		}
	}
	return filters
}


func convertStringToObjectID(idStr string) (primitive.ObjectID, error) {
    objectID, err := primitive.ObjectIDFromHex(idStr)
    if err != nil {
        return primitive.NilObjectID, fmt.Errorf("invalid ObjectID: %w", err)
    }
    return objectID, nil
}
