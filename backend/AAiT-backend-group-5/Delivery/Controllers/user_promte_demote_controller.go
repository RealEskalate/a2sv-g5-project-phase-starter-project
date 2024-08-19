package controllers

import (
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
		ctx.JSON(e.Code, e.Error())
	}
}

func (promoteDemoteController *PromoteDemoteController) DemoteUser(ctx *gin.Context) {
	userID := ctx.Param("id")

	e := promoteDemoteController.PromoteDemoteUC.DemoteUser(ctx, userID)
	if e != nil {
		ctx.JSON(e.Code, e.Error())
	}
}
