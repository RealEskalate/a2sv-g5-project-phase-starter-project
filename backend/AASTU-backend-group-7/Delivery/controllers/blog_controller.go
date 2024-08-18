package controllers

import (
	"blogapp/Domain"

	"github.com/gin-gonic/gin"
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

	var newBlog Domain.Post
	if err := c.ShouldBindJSON(&newBlog); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newBlog.AuthorID = claims.ID
	err, statusCode := controller.BlogUseCase.CreateBlog(c, &newBlog)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Blog created successfully",
		"blog":    newBlog,
	})
}
