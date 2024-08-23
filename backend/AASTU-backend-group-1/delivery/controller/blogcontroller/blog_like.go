package blogcontroller

import (
	"blogs/config"
	"blogs/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (l *BlogController) AddLike(ctx *gin.Context) {
	idHex := ctx.Param("id")

	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   "invalid id",
		})
		return
	}

	var like struct {
		Like bool `bson:"like" json:"like"`
	}

	if err := ctx.ShouldBindJSON(&like); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	claim, ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, domain.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
			Error:   "cannot get claims",
		})
		return
	}

	newlike := domain.Like{
		BlogID: id,
		User:   claim.Username,
		Like:   like.Like,
	}

	err = l.BlogUsecase.AddLike(&newlike)
	if err != nil {
		code := config.GetStatusCode(err)
		ctx.JSON(code, domain.APIResponse{
			Status:  code,
			Message: "Error adding like",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.APIResponse{
		Status:  http.StatusOK,
		Message: "Like added",
		Data:    newlike,
	})
}

func (l *BlogController) RemoveLike(ctx *gin.Context) {
	idHex := ctx.Param("id")

	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   "invalid id",
		})
		return
	}

	claim, ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok {
		log.Println("Error getting claims")
		ctx.JSON(http.StatusInternalServerError, domain.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
			Error:   "cannot get claims",
		})
		return
	}

	err = l.BlogUsecase.RemoveLike(id.Hex(), claim)
	if err != nil {
		code := config.GetStatusCode(err)
		ctx.JSON(code, domain.APIResponse{
			Status:  code,
			Message: "Error removing like",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (l *BlogController) GetBlogLikes(ctx *gin.Context) {
	idHex := ctx.Param("id")

	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   "invalid id",
		})
		return
	}

	likes, err := l.BlogUsecase.GetBlogLikes(id.Hex())
	if err != nil {
		code := config.GetStatusCode(err)

		if code == http.StatusInternalServerError {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, domain.APIResponse{
				Status:  http.StatusInternalServerError,
				Message: "Internal server error",
				Error:   "cannot get likes",
			})
			return
		}

		ctx.JSON(code, domain.APIResponse{
			Status:  code,
			Message: "Cannot get likes",
			Error:   err.Error(),
		})

		return
	}

	if len(likes) == 0 {
		likes = []*domain.Like{}
	}

	ctx.JSON(http.StatusOK, domain.APIResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    likes,
	})
}
