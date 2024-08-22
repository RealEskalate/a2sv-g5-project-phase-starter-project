package controllers

import (
	"blog_api/domain"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)
var validate = validator.New()

func (bc *BlogController) HandleCreateComment(c *gin.Context) {
	var comment domain.NewComment
	if err :=c.ShouldBindJSON(&comment); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	newErr := validate.Struct(comment)
	if newErr != nil {
		c.JSON(http.StatusBadRequest, domain.Response{"error": newErr.Error()})
		return
	}

	userName, exists := c.Keys["username"]
	if !exists{
		c.JSON(http.StatusForbidden, gin.H{"message": "coudn't find the username field"})
		return 
	}
	created_By := userName.(string)
	blogID := c.Param("blogId")
	err := bc.blogUseCase.AddComment(c, blogID, &comment, created_By)
	if err != nil{
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"message": "created successfully" })

}


func (bc *BlogController) HandleUpdateComment(c *gin.Context) {
	var comment domain.NewComment
	if err :=c.ShouldBindJSON(&comment); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	newErr := validate.Struct(comment)
	if newErr != nil {
		c.JSON(http.StatusBadRequest, domain.Response{"error": newErr.Error()})
		return
	}

	userName, exists := c.Keys["username"]
	if !exists{
		c.JSON(http.StatusForbidden, gin.H{"message": "coudn't find the username field"})
		return 
	}
	updateBy := userName.(string)
	blogID := c.Param("blogId")
	commentID := c.Param("commentId")
	err := bc.blogUseCase.UpdateComment(c, blogID, commentID, &comment, updateBy)

	if err != nil{
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "updated successfully"})
}

func (bc *BlogController) HandleDeleteComment(c *gin.Context) {
	userName, exists := c.Keys["username"]
	if !exists{
		c.JSON(http.StatusForbidden, gin.H{"message": "coudn't find the username field"})
		return 
	}

	deleteBy := userName.(string)
	blogID := c.Param("blogId")
	commentID := c.Param("commentId")
	err := bc.blogUseCase.DeleteComment(c, blogID, commentID, deleteBy)
	if err != nil{
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Deleted successfully"})
}