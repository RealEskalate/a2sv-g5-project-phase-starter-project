package blogcontroller

import (
	"blogs/config"
	"blogs/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (b *BlogController) AddView(ctx *gin.Context) {
	var blogs struct {
		IDs []string `json:"ids"`
	}

	if err := ctx.ShouldBindJSON(&blogs); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	var objectIDs []primitive.ObjectID
	for _, idHex := range blogs.IDs {
		id, err := primitive.ObjectIDFromHex(idHex)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id: " + idHex})
			return
		}
		objectIDs = append(objectIDs, id)
	}

	claim, ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok {
		log.Println("failed to get claims")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	err := b.BlogUsecase.AddView(objectIDs, *claim)
	if err != nil {
		code := config.GetStatusCode(err)

		if code == http.StatusInternalServerError {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, "internal server error")
			return
		}

		ctx.JSON(code, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "views have been added"})
}
