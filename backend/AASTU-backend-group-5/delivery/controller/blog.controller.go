package controller

import (
	"net/http"
	"strconv"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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
		var blog domain.PostBlog
		var blg domain.Blog
		if err := c.BindJSON(&blog); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request. Ensure the blog data is correct: " + err.Error()})
			return
		}

		iuser, _ := c.Get("user")

		user := domain.User{}
		if iuser != nil {
			user = iuser.(domain.User)
		}
		blg.ID = primitive.NewObjectID()
		blog.Owner = user
		createdBlog, err := bc.BlogUsecase.CreateBlog(blog)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to add blog post: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Blog post added successfully!", "blog": createdBlog})
	}
}

func (bc *BlogController) GetMyBlogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		blogs, err := bc.BlogUsecase.GetBlogs(0, 0)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Failed to retrieve your blog posts. Please try again later: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Your blog posts retrieved successfully!", "blogs": blogs})
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

		id, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format. Please provide a valid blog ID."})
			return
		}

		blog, err := bc.BlogUsecase.GetOneBlog(id.Hex())
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Blog post not found: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Blog post retrieved successfully!", "blog": blog})
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
		if err := c.BindJSON(&blog); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request: " + err.Error()})
			return
		}

		blog.ID = id

		updatedBlog, err := bc.BlogUsecase.UpdateBlog(id.Hex(), blog)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update blog post: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Blog post updated successfully!", "blog": updatedBlog})
	}
}
func (bc *BlogController) DeleteBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")

		id, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format. Please provide a valid blog ID."})
			return
		}

		err = bc.BlogUsecase.DeleteBlog(id.Hex())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete blog post: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Blog post deleted successfully!"})
	}
}

func (bc *BlogController) FilterBlogs() gin.HandlerFunc {
	return func(c *gin.Context) {

		title := c.Query("title")
		author := c.Query("author")

		tags := c.QueryArray("tags")
		startDate := c.Query("start_date")
		endDate := c.Query("end_date")
		popularity := c.Query("popularity")

		filters := make(map[string]interface{})

		if title != "" {
			filters["title"] = title
		}
		if author != "" {
			filters["owner.username"] = author
		}

		if len(tags) > 0 {
			filters["tag"] = tags
		}
		if startDate != "" {
			filters["created_at"] = bson.M{"$gte": startDate}
		}
		if endDate != "" {
			filters["created_at"] = bson.M{"$lte": endDate}
		}

		if popularity != "" {
			popularityValue, err := strconv.Atoi(popularity)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid popularity value."})
				return
			}
			filters["like_count"] = bson.M{"$gte": popularityValue}
		}

		blogs, err := bc.BlogUsecase.FilterBlog(filters)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to filter blog posts. Please try again: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Filtered blog posts retrieved successfully!", "blogs": blogs})
	}
}
