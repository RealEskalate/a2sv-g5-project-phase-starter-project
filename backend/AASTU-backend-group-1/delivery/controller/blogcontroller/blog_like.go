package blogcontroller

import (
	"blogs/config"
	"blogs/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (l *BlogController) AddLike(ctx *gin.Context) {
	idHex := ctx.Param("id")

	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var like struct {
		Like bool `bson:"like" json:"like"`
	}

	if err := ctx.ShouldBindJSON(&like); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	claim, ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok {
		log.Println("Error getting claims")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	newlike := domain.Like{
		BlogID: id,
		User:   claim.Username,
		Like:   like.Like,
	}

	err = l.BlogUsecase.AddLike(&newlike)
	if err != nil {
		code := config.GetStatusCode(err)
		if code == http.StatusInternalServerError {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		ctx.JSON(code, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "like added"})
}

func (l *BlogController) RemoveLike(ctx *gin.Context) {
	idHex := ctx.Param("id")

	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	claim, ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok {
		log.Println("Error getting claims")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	err = l.BlogUsecase.RemoveLike(id.Hex(), claim)
	if err != nil {
		code := config.GetStatusCode(err)

		if code == http.StatusInternalServerError {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		ctx.JSON(code, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "like removed"})
}

func (l *BlogController) GetBlogLikes(ctx *gin.Context) {
	idHex := ctx.Param("id")

	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	likes, err := l.BlogUsecase.GetBlogLikes(id.Hex())
	if err != nil {
		code := config.GetStatusCode(err)

		if code == http.StatusInternalServerError {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		ctx.JSON(code, gin.H{"error": err.Error()})
		return
	}

	if len(likes) == 0 {
		likes = []*domain.Like{}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"count": len(likes),
		"data":  likes,
	})
}
