package controller

import (
	"blogs/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserController) PromoteUser(ctx *gin.Context) {
	claims, ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok {
		log.Println("Error getting claims")
		ctx.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}

	var request struct {
		Username string `json:"username"`
		Promoted bool   `json:"promoted"`
	}

	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if request.Username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}

	err = u.UserUsecase.PromoteUser(request.Username, request.Promoted, claims)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	promotionWord := "demoted"
	if request.Promoted {
		promotionWord = "promoted"
	}

	ctx.JSON(200, gin.H{
		"message": "User " + request.Username + " has been " + promotionWord,
	})
}
