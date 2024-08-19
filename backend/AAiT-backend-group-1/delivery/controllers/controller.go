package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogController struct {
	blogUsecase domain.BlogUseCase
}


func NewBlogController(bu domain.BlogUseCase) domain.BlogController {
	return &BlogController{
		blogUsecase: bu,
	}
}


func (bc *BlogController) CreateBlog(c *gin.Context) {
	// Implement the logic for creating a blog
	var blog domain.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := bc.blogUsecase.CreateBlog(&blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"data": blog})
}

func (bc *BlogController) GetBlog(c *gin.Context) {
	id := c.Param("id")
	blog, err := bc.blogUsecase.GetBlog(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, blog)

}

func (bc *BlogController) GetBlogs(c *gin.Context) {
	blogs, err := bc.blogUsecase.GetBlogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, blogs)
}

func (bc *BlogController) UpdateBlog(c *gin.Context) {
	var blog domain.Blog
	id := c.Param("id")
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := bc.blogUsecase.UpdateBlog(id, &blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, blog)
}

func (bc *BlogController) DeleteBlog(c *gin.Context) {
	id := c.Param("id")
	err := bc.blogUsecase.DeleteBlog(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}

func (bc *BlogController) SearchBlogs(c *gin.Context) {
	// Implement the logic for searching blogs
	title := c.Query("title")
	author := c.Query("author")
	blogs, err := bc.blogUsecase.SearchBlogs(title , author)
	if err != nil {
		c.JSON(http.StatusInternalServerError , err.Error())
	}

	c.JSON(http.StatusOK , blogs)
}

func (bc *BlogController) FilterBlogs(c *gin.Context) {
	// Implement the logic for filtering blogs
	tags := c.QueryArray("tags")
	if len(tags) == 0 {
		tags = []string{" "}
	}
	timeStr := c.Query("time")
layout := "2006-01-02" // Define the layout to match the date format
var timeValue time.Time
var err error

if timeStr == "" {
    // Set the default time value if no time query parameter is provided
    timeValue, err = time.Parse(layout, "2021-01-01")
} else {
    // Parse the provided time string
    timeValue, err = time.Parse(layout, timeStr)
}

if err != nil {
    log.Println("Invalid time format, using default value:", err)
    timeValue, _ = time.Parse(layout, "2021-01-01")
}


	popular := c.DefaultQuery("popular", "false")
	blogs, err := bc.blogUsecase.FilterBlogs(tags , timeValue , popular == "true")
	if err != nil {
		c.JSON(http.StatusInternalServerError , err.Error())
	}

	c.JSON(http.StatusOK , blogs)
}

func (bc *BlogController) LikeBlog(c *gin.Context) {
	var blogID = c.Param("id")
	userID := c.GetString("user_id")
	err := bc.blogUsecase.LikeBlog(userID , blogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError , err.Error())
		return
	}
	c.JSON(http.StatusOK , gin.H{"message" : "Blog liked successfully"})
}

func (bc *BlogController) DislikeBlog(c *gin.Context) {
	var blogID = c.Param("id")
	userID := c.GetString("user_id")
	err := bc.blogUsecase.DisLike(userID , blogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError , err.Error())
		return
	}
	c.JSON(http.StatusOK , gin.H{"message" : "Blog Disliked successfully"})
}

func (bc *BlogController) AddComment(c *gin.Context) {
	var blogID = c.Param("id")
	var comment domain.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId , _ := primitive.ObjectIDFromHex(c.GetString("user_id"))
	comment.AuthorID = userId
	err := bc.blogUsecase.AddComment(blogID , &comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": comment})

}

func (bc *BlogController) DeleteComment(c *gin.Context) {
	// Implement the logic for deleting a comment
	var commentID = c.Param("comment_id")
	var blogID = c.Param("id")
	err := bc.blogUsecase.DeleteComment(blogID , commentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError , err.Error())
		return
	}
	c.JSON(http.StatusOK , gin.H{"message" : "Comment deleted successfully"})
}

func (bc *BlogController) EditComment(c *gin.Context) {
	// Implement the logic for editing a comment
	var comment domain.Comment
	var commentID = c.Param("comment_id")
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} 
	err := bc.blogUsecase.EditComment(commentID , &comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError , err.Error())
		return
	}
	c.JSON(http.StatusOK , gin.H{"message" : "Comment edited successfully"})
}