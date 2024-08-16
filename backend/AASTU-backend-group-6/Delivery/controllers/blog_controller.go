package controllers

import (
	domain "blogs/Domain"
	"net/http"

	"github.com/gin-gonic/gin"
)


func CreateBlogController(ctx *gin.Context) {
	var post domain.Post
	err := ctx.ShouldBind(&post)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		})
	}
}