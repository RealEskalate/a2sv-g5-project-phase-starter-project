package controllers

import (
	"blogapp/Domain"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type blogController struct {
	BlogUseCase Domain.BlogUseCase
}

func NewBlogController(usecase Domain.BlogUseCase) *blogController {

	return &blogController{
		BlogUseCase: usecase,
	}
}

func (controller *blogController) CreateBlog(c *gin.Context) {
	claims, err := Getclaim(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	var newBlogPost Domain.Post
	if err := c.ShouldBindJSON(&newBlogPost); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newBlogPost.AuthorID = claims.ID
	// generate id for post
	newBlogPost.ID = primitive.NewObjectID()
	// generate empty array for comments
	newBlogPost.Comments = []*Domain.Comment{}
	// generate empty array for tags
	newBlogPost.Tags = []*Domain.Tag{}
	// generate empty array for likeDislike
	newBlogPost.LikeDislike = []*Domain.LikeDislike{}
	// generate slug
	newBlogPost.Slug = GenerateSlug(newBlogPost.Title)
	//created at and updated at
	newBlogPost.PublishedAt = time.Now()
	newBlogPost.UpdatedAt = time.Now()

	err, statusCode := controller.BlogUseCase.CreateBlog(c, &newBlogPost)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Blog created successfully",
		"blog":    newBlogPost,
	})
}
