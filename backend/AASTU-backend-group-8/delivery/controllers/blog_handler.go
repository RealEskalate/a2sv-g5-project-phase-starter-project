package controllers

import (
	"fmt"
	"meleket/domain"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogController struct {
	blogUsecase domain.BlogUsecaseInterface
}

func NewBlogController(blogUsecase domain.BlogUsecaseInterface) *BlogController {
	return &BlogController{blogUsecase: blogUsecase}
}

func (bc *BlogController) CreateBlogPost(c *gin.Context) {
	var blog domain.BlogPost
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	blog.CreatedAt = time.Now()
	blog.UpdatedAt = time.Now()

	// autherID := c.MustGet("userID").(primitive.ObjectID)
	// blog.AuthorID = autherID

	createdBlog, err := bc.blogUsecase.CreateBlogPost(&blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Blog post created successfully", "id": createdBlog})
}

func (bc *BlogController) GetAllBlogPosts(c *gin.Context) {
	var paginationInfo domain.Pagination
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	sortBy := c.DefaultQuery("sortBy", "created_at")
	sortOrderStr := c.DefaultQuery("sortOrder", "-1")
	fmt.Println(pageStr)
	fmt.Println(pageSizeStr)

	limit, err := strconv.Atoi(pageSizeStr)
	paginationInfo.Limit = limit
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	page, err := strconv.Atoi(pageStr)
	paginationInfo.Page = page
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	sortOrder, err := strconv.Atoi(sortOrderStr)
	if err != nil || (sortOrder != 1 && sortOrder != -1) {
		sortOrder = -1 // Default to descending order
	}
	blogs, err := bc.blogUsecase.GetAllBlogPosts(paginationInfo, sortBy, sortOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

func (bc *BlogController) GetBlogByID(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id")) // should be updated
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	blog, err := bc.blogUsecase.GetBlogByID(id)

	fmt.Println(blog)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog post not found"})
		return
	}

	// Add veiws to the blog
	blog.Veiws += 1
	fmt.Println(blog)
	bc.blogUsecase.UpdateBlogPost(id, blog)
	blog, err = bc.blogUsecase.GetBlogByID(id)
	fmt.Println(blog)

	c.JSON(http.StatusOK, blog)

}

func (bc *BlogController) UpdateBlogPost(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id")) // should be updated
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	var blog domain.BlogPost
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	blog.UpdatedAt = time.Now()

	updatedBlog, err := bc.blogUsecase.UpdateBlogPost(id, &blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedBlog)
}

// func (bc *BlogController) SearchBlogPost(c *gin.Context) {
// 	var search domain.SearchBlogPost

// 	if err := c.ShouldBindJSON(&search); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	blogs, err := bc.blogUsecase.SearchBlogPosts(&search)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, blogs)
// }

func (bc *BlogController) DeleteBlogPost(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	if err := bc.blogUsecase.DeleteBlogPost(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog post deleted successfully"})
}
