package controllers

import (
	"meleket/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LikeDislikeController struct {
	likeDislikeUsecase *usecases.LikeDislikeUsecase
}

func NewLikeDislikeController(likeDislikeUsecase *usecases.LikeDislikeUsecase) *LikeDislikeController {
	return &LikeDislikeController{likeDislikeUsecase: likeDislikeUsecase}
}

func (c *LikeDislikeController) ToggleLikeDislike(ctx *gin.Context) {
	blogIDHex := ctx.Param("blogID")
	userIDHex := ctx.Param("userID")
	likeType := ctx.Query("type") // "like" or "dislike"

	blogID, err := primitive.ObjectIDFromHex(blogIDHex)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	userID, err := primitive.ObjectIDFromHex(userIDHex)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = c.likeDislikeUsecase.ToggleLikeDislike(blogID, userID, likeType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Action performed successfully"})
}
