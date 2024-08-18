package controllers

import (
	"blog_api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BlogController struct {
	blogUseCase domain.BlogUseCaseInterface
}

var validate = validator.New()
func (bc *BlogController) CreateBlogHandler(c *gin.Context){
	var blog domain.Blog
	if err := c.ShouldBindJSON(&blog); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	err := validate.Struct(blog)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	err = bc.blogUseCase.CreateBlogPost(c, &blog)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "blog created successfully"})

}

func (bc *BlogController) UpdateBlogHandler (c *gin.Context){
	blogId := c.Param("id") 
	var blog domain.Blog
	if err := c.ShouldBindJSON(&blog); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	err := bc.blogUseCase.EditBlogPost(c, blogId, &blog)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	c.JSON(http.StatusCreated, gin.H{"message": "created successfuly"})
}

func (bc *BlogController) DeleteBlogHandler (c *gin.Context){
	blogId := c.Param("id") 

	err := bc.blogUseCase.DeleteBlogPost(c, blogId)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to delete data"})
		return 
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted successfuly"})
}

func (bc *BlogController) GetBlogHandler (c *gin.Context){
	// robel implement this
}

func (bc *BlogController) GetBlogByIDHandler (c *gin.Context){
	// robel implement this
}
