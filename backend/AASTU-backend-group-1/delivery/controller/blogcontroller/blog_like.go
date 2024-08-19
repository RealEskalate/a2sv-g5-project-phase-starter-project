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
