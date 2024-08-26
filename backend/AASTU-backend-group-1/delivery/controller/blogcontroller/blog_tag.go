package blogcontroller

import (
	"blogs/config"
	"blogs/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (b *BlogController) GetTags(ctx *gin.Context) {

	tags,err := b.BlogUsecase.GetTags()
	if err != nil {
		code := config.GetStatusCode(err)
		log.Println(err)

		ctx.JSON(code, domain.APIResponse{
			Status:  code,
			Message: "Error getting tags",
			Error:   err.Error(),
		})
		return
	}

	if len(tags) == 0 {
		tags = []*domain.Tag{}
	}

	ctx.JSON(http.StatusOK, domain.APIResponse{
		Status:  http.StatusOK,
		Message: "Tags retrieved",
		Data:    tags,
	})
}

func (b *BlogController) InsertTag(ctx *gin.Context) {
	var tag domain.Tag
	err := ctx.BindJSON(&tag)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   "Invalid request",
		})
		return
	}

	claims, ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok {
		log.Println("Error getting claims")
		ctx.JSON(http.StatusInternalServerError, domain.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
			Error:   "Error getting claims",
		})
		return
	}



	err = b.BlogUsecase.InsertTag(&tag, claims)
	if err != nil {
		code := config.GetStatusCode(err)
		log.Println(err)

		ctx.JSON(code, domain.APIResponse{
			Status:  code,
			Message: "Error adding tag",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, domain.APIResponse{
		Status:  http.StatusCreated,
		Message: "Tag added",
	})
}


func (b *BlogController) RemoveTags(ctx *gin.Context) {
	var tag domain.Tag
	err := ctx.BindJSON(&tag)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   "Invalid request",
		})
		return
	}
    
	claims, ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok {
		log.Println("Error getting claims")
		ctx.JSON(http.StatusInternalServerError, domain.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
			Error:   "Error getting claims",
		})
		return
	}


	err = b.BlogUsecase.DeleteTag(&tag , claims)
	if err != nil {
		code := config.GetStatusCode(err)
		log.Println(err)

		ctx.JSON(code, domain.APIResponse{
			Status:  code,
			Message: "Error removing tag",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.APIResponse{
		Status:  http.StatusOK,
		Message: "Tag removed",
	})
}