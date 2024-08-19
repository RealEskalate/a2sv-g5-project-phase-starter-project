package Controller

import (
	domain "AAiT-backend-group-8/Domain"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (controller *Controller) GetComments(ctx *gin.Context) {
	blogID := ctx.Param("blogID")
	comments, err := controller.commentUseCase.GetComments(blogID)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, comments)
}

// a function to get the token from the authorization header , and decode it to get the user id
func (controller *Controller) CreateComment(ctx *gin.Context) {
	blogID := ctx.Param("blogID")
	token, err := controller.ExtractToken(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err = controller.commentUseCase.DecodeToken(token, []byte("secret"))
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var comment domain.Comment
	if err := ctx.BindJSON(&comment); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//!TODO add the name and id of the user to the comment

	if comment.Body == "" {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "comment body is required"})
		return
	}
	err = controller.commentUseCase.CreateComment(&comment, blogID)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "comment created successfully"})
}

func (controller *Controller) UpdateComment(ctx *gin.Context) {
	var comment domain.Comment
	if err := ctx.BindJSON(&comment); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if comment.Body == "" {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "comment body is required"})
		return
	}
	err := controller.commentUseCase.UpdateComment(&comment)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "comment updated successfully"})
}

func (controller *Controller) DeleteComment(ctx *gin.Context) {
	commentID := ctx.Param("commentID")
	err := controller.commentUseCase.DeleteComment(commentID)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "comment deleted successfully"})
}

func (controller *Controller) ExtractToken(ctx *gin.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return "", gin.Error{
			Err:  errors.New("authorization header not provided"),
			Type: gin.ErrorTypePublic,
		}
	}
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", gin.Error{
			Err:  errors.New("invalid authorization header format"),
			Type: gin.ErrorTypePublic,
		}
	}
	return parts[1], nil
}
