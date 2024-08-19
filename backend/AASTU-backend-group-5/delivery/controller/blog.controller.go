package controller

import (
	"net/http"
	"strconv"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogController struct {
	BlogUsecase domain.Blog_Usecase_interface
	UserUsecase domain.User_Usecase_interface
}

func NewBlogController(blogUsecase domain.Blog_Usecase_interface, userUsecase domain.User_Usecase_interface) *BlogController {
	return &BlogController{
		BlogUsecase: blogUsecase,
		UserUsecase: userUsecase,
	}
}

func (bc *BlogController) CreateBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		var blog domain.Blog
		if err := c.BindJSON(&blog); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request. Ensure the blog data is correct: " + err.Error()})
			return
		}

		blog.ID = primitive.NewObjectID()

		claims, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. Please log in to add a blog post."})
			return
		}

		userClaims := claims.(*domain.Claims)
		createdByID, err := primitive.ObjectIDFromHex(userClaims.UserID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID. Please try again."})
			return
		}
		blog.ID = createdByID

		createdBlog, err := bc.BlogUsecase.CreateBlog(blog)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to add blog post. Please ensure all required fields are filled: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Blog post added successfully!", "blog": createdBlog})
	}
}

func (bc *BlogController) GetMyBlogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. Please log in to view your blog posts."})
			return
		}

		userClaims, ok := claims.(*domain.Claims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user claims. Please try again."})
			return
		}

		userId, _ := primitive.ObjectIDFromHex(userClaims.UserID)

		blogs, err := bc.BlogUsecase.GetBlogs(0, 0) // Assuming limit and page_number are 0 for the current user
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Failed to retrieve your blog posts. Please try again later: " + err.Error()})
			return
		}

		var userBlogs []domain.Blog
		for _, blog := range blogs {
			if blog.ID == userId {
				userBlogs = append(userBlogs, blog)
			}
		}

		c.JSON(http.StatusOK, gin.H{"message": "Your blog posts retrieved successfully!", "blogs": userBlogs})
	}
}

func (bc *BlogController) GetAllBlogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
		pageNumber, _ := strconv.Atoi(c.DefaultQuery("page_number", "1"))

		blogs, err := bc.BlogUsecase.GetBlogs(limit, pageNumber)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Failed to retrieve blog posts. Please try again later: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "All blog posts retrieved successfully!", "blogs": blogs})
	}
}

func (bc *BlogController) GetOneBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")

		_, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format. Please provide a valid blog ID."})
			return
		}

		blogs, err := bc.BlogUsecase.GetOneBlog(idStr)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Blog post not found. Please ensure the blog ID is correct: " + err.Error()})
			return
		}

		if len(blogs) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Blog post not found."})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Blog post retrieved successfully!", "blog": blogs[0]})
	}
}

func (bc *BlogController) UpdateBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")

		id, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format. Please provide a valid blog ID."})
			return
		}

		var blog domain.Blog
		blog.ID = id
		if err := c.BindJSON(&blog); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request. Ensure the blog data is correct: " + err.Error()})
			return
		}

		updatedBlog, err := bc.BlogUsecase.UpdateBlog(idStr, blog)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update blog post. Please ensure all required fields are filled: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Blog post updated successfully!", "blog": updatedBlog})
	}
}

func (bc *BlogController) DeleteBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")

		if err := bc.BlogUsecase.DeleteBlog(idStr); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete blog post. Please try again: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Blog post deleted successfully!"})
	}
}

func (bc *BlogController) FilterBlogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		filters := make(map[string]string)
		for key, values := range c.Request.URL.Query() {
			if len(values) > 0 {
				filters[key] = values[0]
			}
		}

		blogs, err := bc.BlogUsecase.FilterBlog(filters)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to filter blog posts. Please try again: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Filtered blog posts retrieved successfully!", "blogs": blogs})
	}
}
