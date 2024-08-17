package controllers

import (
	"blogapp/Domain"
	"fmt"

	"github.com/gin-gonic/gin"
)

type blogController struct {
	AuthUseCase Domain.BlogUseCase
}

func NewBlogController(usecase Domain.BlogUseCase) *blogController {

	return &blogController{
		AuthUseCase: usecase,
	}
}

func (controller *blogController) CreateBlog(c *gin.Context) {
	claims, err := Getclaim(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(claims.ID)

	// var newBlog Domain.Blog
	// if err := c.ShouldBindJSON(&newBlog); err != nil {
	// 	c.JSON(400, gin.H{"error": err.Error()})
	// 	return
	// }

	// newBlog.UserID = claims.ID

	// if err := controller.AuthUseCase.CreateBlog(&newBlog); err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error()})
	// 	return
	// }

	// c.JSON(200, gin.H{"message": "Blog created successfully"})
}
