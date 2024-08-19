package controller

import (
	"blogs/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func (l *BlogController) AddComment(ctx *gin.Context) {
    idHex := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {	
		ctx.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	

	var  comment struct {
		Content string `bson:"content" json:"content"`
		
	}


	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	claim, ok := ctx.MustGet("claims").(*domain.LoginClaims)

	if !ok {
		ctx.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
    
	
	newcomment := domain.Comment{
		BlogID : id,
		Author : claim.Username,
		Content : comment.Content,
		Date: time.Now(),
	}




	err = l.BlogUsecase.AddComment(&newcomment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.JSON(http.StatusOK, "comment added")
}
