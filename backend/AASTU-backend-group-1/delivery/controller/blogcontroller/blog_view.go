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
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	var objectIDs []primitive.ObjectID
	for _, idHex := range blogs.IDs {
		id, err := primitive.ObjectIDFromHex(idHex)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, domain.APIResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid request",
				Error:   "invalid id",
			})
			return
		}
		objectIDs = append(objectIDs, id)
	}

	claim, ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok {
		log.Println("failed to get claims")
		ctx.JSON(http.StatusInternalServerError, domain.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
			Error:   "cannot get claims",
		})
		return
	}

	err := b.BlogUsecase.AddView(objectIDs, *claim)
	if err != nil {
		code := config.GetStatusCode(err)
		ctx.JSON(code, domain.APIResponse{
			Status:  code,
			Message: "Failed to add views",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, domain.APIResponse{
		Status:  http.StatusCreated,
		Message: "Views added",
	})
}
