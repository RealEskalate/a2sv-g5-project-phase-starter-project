package controllers

import (
	"net/http"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	"github.com/gin-gonic/gin"
)

type PromoteDemoteController struct {
	PromoteDemoteUC interfaces.PromoteDemoteUserUsecase
}

func NewPromoteDemoteController(promoteDemoteUC interfaces.PromoteDemoteUserUsecase) *PromoteDemoteController {
	return &PromoteDemoteController{
		PromoteDemoteUC: promoteDemoteUC,
	}
}

func (promoteDemoteController *PromoteDemoteController) PromoteUser(ctx *gin.Context) {
	userID := ctx.Param("id")

	e := promoteDemoteController.PromoteDemoteUC.PromoteUser(ctx, userID)
	if e != nil {
		ctx.IndentedJSON(e.Code, gin.H{"error": e.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user successfully promoted"})
}

func (promoteDemoteController *PromoteDemoteController) DemoteUser(ctx *gin.Context) {
	userID := ctx.Param("id")

	e := promoteDemoteController.PromoteDemoteUC.DemoteUser(ctx, userID)
	if e != nil {
		ctx.IndentedJSON(e.Code, gin.H{"error": e.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user successfully demoted"})
}
