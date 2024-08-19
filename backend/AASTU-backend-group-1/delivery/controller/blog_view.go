package controller

import (
	"blogs/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (l *BlogController) AddView(ctx *gin.Context) {
	var blogs struct {
		IDs []string `json:"ids"`
	}

	if err := ctx.ShouldBindJSON(&blogs); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		ctx.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
    
    
	err := l.BlogUsecase.AddView(objectIDs,*claim)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.JSON(http.StatusOK, "view added")


}
	

