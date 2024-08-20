package controllers

import (
	"meleket/domain"
	"net/http"

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

	// autherID := c.MustGet("userID").(primitive.ObjectID)
	// blog.AuthorID = autherID

	createdBlog, err := bc.blogUsecase.CreateBlogPost(&blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Blog post created successfully", "blog": createdBlog})
}

func (bc *BlogController) GetAllBlogPosts(c *gin.Context) {
	blogs, err := bc.blogUsecase.GetAllBlogPosts()
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
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog post not found"})
		return
	}

	c.JSON(http.StatusOK, blog)

}

func (bc *BlogController) UpdateBlogPost(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	var blog domain.BlogPost
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the usecase to update the blog post
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
