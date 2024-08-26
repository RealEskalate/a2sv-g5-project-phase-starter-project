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

	autherID := c.MustGet("userID").(primitive.ObjectID)
	fmt.Println("Author ID: ", autherID)
	blog.AuthorID = autherID

	BlogID, err := bc.blogUsecase.CreateBlogPost(&blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Blog post created successfully", "id": BlogID})

}

func (bc *BlogController) GetAllBlogPosts(c *gin.Context) {
	var filter domain.BlogFilter
	filter.Title = c.Query("title")
	filter.AuthorID = c.Query("author_id")
	filter.Tags = c.QueryArray("tags")
	filter.Search = c.Query("search")

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

	blogs, err := bc.blogUsecase.GetAllBlogPosts(paginationInfo, sortBy, sortOrder, filter)
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

// 	c.JSON(http.StatusOK, blogPost)
// }

// Handler to like a blog post
func (bc *BlogController) LikeBlogPost(c *gin.Context) {
	blogIDParam := c.Param("id")
	blogID, err := primitive.ObjectIDFromHex(blogIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	userID := c.MustGet("userID").(primitive.ObjectID) // Already a valid ObjectID
	// Assume userID is obtained from JWT middleware

	err = bc.blogUsecase.LikeBlogPost(blogID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog post liked successfully"})
}

func (bc *BlogController) DislikeBlogPost(c *gin.Context) {
	blogIDParam := c.Param("id")
	blogID, err := primitive.ObjectIDFromHex(blogIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	userID := c.MustGet("userID").(primitive.ObjectID) // Already a valid ObjectID
	// Assume userID is obtained from JWT middleware

	err = bc.blogUsecase.DislikeBlogPost(blogID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog post disliked successfully"})
}

// Handler to add a comment to a blog post
func (bc *BlogController) AddComment(c *gin.Context) {
	blogIDParam := c.Param("id")
	blogID, err := primitive.ObjectIDFromHex(blogIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	var comment domain.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment data"})
		return
	}

	comment.UserID = c.MustGet("userID").(primitive.ObjectID) // Assume userID is obtained from JWT middleware
	comment.CreatedAt = time.Now()

	err = bc.blogUsecase.AddCommentToBlogPost(blogID, comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to add comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment added successfully"})
}
