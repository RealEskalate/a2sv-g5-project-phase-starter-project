package controllers

import (
	domain "aait-backend-group4/Domain"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogController struct {
	BlogUsecase domain.BlogUsecase
}

// CreateBlog handles the creation of a new blog post.
// It parses the incoming JSON request to a domain.Blog object, sets initial values, and calls the usecase to create the blog.
func (bc *BlogController) CreateBlog(c *gin.Context) {
	var blog domain.Blog

	// Bind the JSON request body to the Blog object
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set initial values for the new blog
	blog.ID = primitive.NewObjectID()
	blog.Popularity = 0
	blog.Feedbacks = domain.Feedback{}
	blog.Created_At = time.Now()
	blog.Updated_At = time.Now()

	// Call the usecase to create the blog in the database
	if err := bc.BlogUsecase.CreateBlog(c, &blog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blog)
}

// UpdateBlog handles updating an existing blog post.
// It parses the blog ID from the URL, retrieves the updated blog data from the request body, and updates the blog.
func (bc *BlogController) UpdateBlog(c *gin.Context) {
	blogID := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	// Retrieve userID from the request headers
	userID := c.Request.Header.Get("userID")
	var blog domain.BlogUpdate

	// Bind the JSON request body to the BlogUpdate object
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the updated timestamp
	updatedTime := time.Now()
	blog.Updated_At = &updatedTime

	// Call the usecase to update the blog
	if err := bc.BlogUsecase.UpdateBlog(c, objectID, blog, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog updated successfully"})
}

// DeleteBlog handles the deletion of a blog post.
// It parses the blog ID from the URL, and userID from the request headers, and calls the usecase to delete the blog.
func (bc *BlogController) DeleteBlog(c *gin.Context) {
	blogID := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	// Retrieve userID from the request headers
	userID := c.Request.Header.Get("userID")

	// Call the usecase to delete the blog
	if err := bc.BlogUsecase.DeleteBlog(c, objectID, userID); err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Blog post not found"})
		} else if err.Error() == "unauthorized" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized to delete this blog post"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}

// FetchAll handles fetching all blog posts with pagination.
// It retrieves pagination parameters from the query and calls the usecase to get the blogs.
func (bc *BlogController) FetchAll(c *gin.Context) {
	limitParam := c.DefaultQuery("limit", "10")
	pageParam := c.DefaultQuery("page", "1")

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit <= 0 {
		limit = 10
	}

	page, err := strconv.Atoi(pageParam)
	if err != nil || page <= 0 {
		page = 1
	}

	// Call the usecase to fetch all blogs with pagination
	blogs, err := bc.BlogUsecase.FetchAll(c, limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

// FetchByPageAndPopularity handles fetching blog posts with pagination and sorted by popularity.
// It retrieves pagination parameters from the query and calls the usecase to get the blogs sorted by popularity.
func (bc *BlogController) FetchByPageAndPopularity(c *gin.Context) {
	limitParam := c.DefaultQuery("limit", "10")
	pageParam := c.DefaultQuery("page", "1")

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit <= 0 {
		limit = 10
	}

	page, err := strconv.Atoi(pageParam)
	if err != nil || page <= 0 {
		page = 1
	}

	// Call the usecase to fetch blogs sorted by popularity with pagination
	blogs, err := bc.BlogUsecase.FetchByPageAndPopularity(c, limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

// FetchByBlogAuthor handles fetching blog posts by a specific author ID with pagination.
// It retrieves the author ID and pagination parameters from the query and calls the usecase to get the blogs.
func (bc *BlogController) FetchByBlogAuthor(c *gin.Context) {
	authorID := c.Param("author_id")
	limitParam := c.DefaultQuery("limit", "10")
	pageParam := c.DefaultQuery("page", "1")

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit <= 0 {
		limit = 10
	}

	page, err := strconv.Atoi(pageParam)
	if err != nil || page <= 0 {
		page = 1
	}

	// Call the usecase to fetch blogs by author with pagination
	blogs, err := bc.BlogUsecase.FetchByBlogAuthor(c, authorID, limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

// FetchByTags handles fetching blog posts by specified tags with pagination.
// It retrieves the tags and pagination parameters from the query and calls the usecase to get the blogs.
func (bc *BlogController) FetchByTags(c *gin.Context) {
	var tags []domain.Tag
	limitParam := c.DefaultQuery("limit", "10")
	pageParam := c.DefaultQuery("page", "1")

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit <= 0 {
		limit = 10
	}

	page, err := strconv.Atoi(pageParam)
	if err != nil || page <= 0 {
		page = 1
	}

	// Bind the JSON request body to the tags slice
	if err := c.ShouldBindJSON(&tags); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the usecase to fetch blogs by tags with pagination
	blogs, err := bc.BlogUsecase.FetchByTags(c, tags, limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

// FetchByBlogTitle handles fetching blog posts by title.
// It retrieves the title from the query and calls the usecase to get the blogs matching the title.
func (bc *BlogController) FetchByBlogTitle(c *gin.Context) {
	title := c.Query("title")

	// Call the usecase to fetch blogs by title
	blogs, err := bc.BlogUsecase.FetchByBlogTitle(c, title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

// AddComment handles adding a comment to a specific blog post.
// It parses the blog ID from the URL, binds the comment from the request body, and calls the usecase to add the comment.
func (bc *BlogController) AddComment(c *gin.Context) {
	var comment domain.Comment
	id := c.Param("id")

	// Bind the JSON request body to the Comment object
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the date for the new comment
	comment.Date = time.Now()

	// Call the usecase to add the comment to the blog
	if err := bc.BlogUsecase.AddComment(c, id, comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment added successfully"})
}

// UpdateComment handles updating an existing comment in a blog post.
// It parses the blog ID from the URL, binds the updated comment from the request body, and calls the usecase to update the comment.
func (bc *BlogController) UpdateComment(c *gin.Context) {
	blogID := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	// Retrieve userID from the request headers
	userID := c.Request.Header.Get("userID")
	var updatedComment domain.Comment

	// Bind the JSON request body to the updated Comment object
	if err := c.ShouldBindJSON(&updatedComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the date for the updated comment
	updatedComment.Date = time.Now()

	// Call the usecase to update the comment in the blog
	if err := bc.BlogUsecase.UpdateComment(c, objectID, userID, updatedComment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully"})
}

// RemoveComment handles removing a comment from a blog post.
// It parses the blog ID from the URL and userID from the request headers, and calls the usecase to remove the comment.
func (bc *BlogController) RemoveComment(c *gin.Context) {
	blogID := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	// Retrieve userID from the request headers
	userID := c.Request.Header.Get("userID")

	// Call the usecase to remove the comment from the blog
	if err := bc.BlogUsecase.RemoveComment(c, objectID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment removed successfully"})
}

// SearchBlogs handles searching for blog posts based on various filters.
// It retrieves filter parameters from the query, binds them to a Filter object, and calls the usecase to search for blogs.
func (bc *BlogController) SearchBlogs(c *gin.Context) {
	var filter domain.Filter
	limitParam := c.DefaultQuery("limit", "10")
	pageParam := c.DefaultQuery("page", "1")

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit <= 0 {
		limit = 10
	}

	page, err := strconv.Atoi(pageParam)
	if err != nil || page <= 0 {
		page = 1
	}

	// Extract filter parameters from the query
	authorName := c.Query("authorName")
	if authorName != "" {
		filter.AuthorName = &authorName
	}

	tags := c.QueryArray("tags")
	if len(tags) > 0 {
		var tagObjects []domain.Tag
		for _, tag := range tags {
			tagObjects = append(tagObjects, domain.Tag(tag))
		}
		filter.Tags = &tagObjects
	}

	blogTitle := c.Query("blogTitle")
	if blogTitle != "" {
		filter.BlogTitle = &blogTitle
	}

	popularity, err := strconv.ParseFloat(c.Query("popularity"), 64)
	if err == nil {
		filter.Popularity = &popularity
	}

	sortByField := c.Query("sortByField")
	sortByOrder := c.Query("sortByOrder")
	if sortByField != "" {
		sortBy := domain.FilterParam(sortByField)
		filter.Sort_By = &sortBy
	}

	if sortByOrder != "" {
		sortBy := domain.FilterParam(sortByOrder)
		filter.Sort_By = &sortBy
	}

	// Call the usecase to search for blogs with the specified filters
	blogs, err := bc.BlogUsecase.SearchBlogs(c, filter, limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

// FetchByBlogID handles fetching a single blog post by its ID.
// It parses the blog ID from the URL and calls the usecase to get the blog post.
func (bc *BlogController) FetchByBlogID(c *gin.Context) {
	blogID := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	// Call the usecase to fetch the blog post by ID
	blog, err := bc.BlogUsecase.FetchByBlogID(c, objectID.Hex())
	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Blog post not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred"})
		}
		return
	}

	c.JSON(http.StatusOK, blog)
}
