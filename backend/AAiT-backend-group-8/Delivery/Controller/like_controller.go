package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) LikeBlog(ctx *gin.Context) {
	blogID := ctx.Param("blogID")
	// Get token from header
	token, err := c.ExtractToken(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, _, err := c.commentUseCase.DecodeToken(token)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	likes, err := c.LikeUseCase.LikeComment(userID, blogID)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if likes {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "liked"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "unliked"})
}

func (c *Controller) GetLikes(ctx *gin.Context) {
	blogID := ctx.Param("blogID")
	likes, err := c.LikeUseCase.GetLikes(blogID)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, likes)
}
