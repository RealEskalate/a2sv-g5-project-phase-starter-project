package controllers

import (
	"blogapp/Domain"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type blogController struct {
	AuthUseCase Domain.BlogUseCase
}

func Getclaim(c *gin.Context) (*Domain.AccessClaims, error) {
	claim, exists := c.Get("claim")
	if !exists {
		return nil, errors.New("claim not set")
	}

	userClaims, ok := claim.(*Domain.AccessClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return userClaims, nil
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
