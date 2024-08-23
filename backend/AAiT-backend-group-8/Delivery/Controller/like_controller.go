package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *Controller) LikeBlog(ctx *gin.Context) {
	blogID := ctx.Param("blogID")
	// Get token from header
	token, err := controller.ExtractToken(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, _, err := controller.commentUseCase.DecodeToken(token)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = controller.blogUseCase.GetBlogByID(blogID)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	likes, err := controller.LikeUseCase.LikeComment(userID, blogID)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if likes {
		err = controller.blogUseCase.UpdateBlogLikeCount(blogID, true)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "liked"})
		return
	}
	err = controller.blogUseCase.UpdateBlogLikeCount(blogID, false)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "unliked"})
}

func (controller *Controller) GetLikes(ctx *gin.Context) {
	blogID := ctx.Param("blogID")
	likes, err := controller.LikeUseCase.GetLikes(blogID)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, likes)
}
