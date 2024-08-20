package blogcontroller

import (
	"blogs/domain"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (l *BlogController) AddLike(ctx *gin.Context) {
	idHex := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	var like struct {
		Like   bool               `bson:"like" json:"like"`

	}

	if err := ctx.ShouldBindJSON(&like); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	claim, ok := ctx.MustGet("claims").(*domain.LoginClaims)

	if !ok {
		ctx.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	newlike := domain.Like{
		BlogID:  id,
		User:    claim.Username,
		Like:    like.Like,

	}

	err = l.BlogUsecase.AddLike(&newlike)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.JSON(http.StatusOK, "like added")
}

func (l *BlogController) RemoveLike(ctx *gin.Context) {
	idHex := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	claim, ok := ctx.MustGet("claims").(*domain.LoginClaims)

	if !ok {
		ctx.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	err = l.BlogUsecase.RemoveLike(id.Hex(), claim)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.JSON(http.StatusOK, "like removed")
}

func (l *BlogController) GetBlogLikes(ctx *gin.Context) {
	idHex := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	likes, err := l.BlogUsecase.GetBlogLikes(id.Hex())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.JSON(http.StatusOK, likes)
}
