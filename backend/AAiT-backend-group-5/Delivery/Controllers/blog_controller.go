package controllers

import (
	"errors"
	"net/http"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogController struct {
	BlogUsecase interfaces.BlogUsecase
	JwtService  interfaces.JwtService
}

func (blogController BlogController) GetBlogs(ctx *gin.Context) {
	blogs, err := blogController.BlogUsecase.GetBlogs(ctx)
	if err != nil {
		ctx.JSON(err.Code, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}

func (blogController BlogController) GetBlog(ctx *gin.Context) {
	blog, err := blogController.BlogUsecase.GetBlog(ctx)
	if err != nil {
		ctx.JSON(err.Code, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

func (blogController BlogController) CreateBlog(ctx *gin.Context) {
	var new_blog *models.Blog

	if e := ctx.BindJSON(new_blog); e != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	err := blogController.BlogUsecase.CreateBlog(ctx, new_blog)
	if err != nil {
		ctx.JSON(err.Code, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "blog created successfully"})
}

func (blogController BlogController) SearchBlogs(ctx *gin.Context) {
	var updated_request dtos.FilterBlogRequest

	if e := ctx.BindJSON(&updated_request); e != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	blogs, err := blogController.BlogUsecase.SearchBlogs(ctx, updated_request)
	if err != nil {
		ctx.JSON(err.Code, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}

func (blogController BlogController) UpdateBlog(ctx *gin.Context) {
	var updated_request *dtos.UpdateBlogRequest

	if e := ctx.BindJSON(updated_request); e != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// Use the JWT service to get the claims from the token
	authHeader := ctx.GetHeader("Authorization")
	claims, err := blogController.JwtService.GetClaims(authHeader)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.New("invalid token"))
		return
	}

	// Get user role and user ID from the claims
	userRole := claims.Role
	userId, err := primitive.ObjectIDFromHex(claims.Id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.New("invalid user id"))
		return
	}

	// update only allowed for blog author or admin
	if userId == updated_request.AuthorID || userRole == "admin" {

		blogID, e := primitive.ObjectIDFromHex(updated_request.BlogID)
		if e != nil {
			ctx.JSON(http.StatusInternalServerError, errors.New("internal server error"))
			return
		}

		updated_blog := models.Blog{
			ID:      blogID,
			Title:   updated_request.Title,
			Content: updated_request.Content,
			Tags:    updated_request.Tags,
		}

		err := blogController.BlogUsecase.UpdateBlog(ctx, &updated_blog)
		if err != nil {
			ctx.JSON(err.Code, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "blog update successful"})
		return
	}

	ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "update not allowed"})
}

func (blogController BlogController) DeleteBlog(ctx *gin.Context) {
	var delete_request *dtos.DeleteBlogRequest

	if e := ctx.BindJSON(delete_request); e != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// Use the JWT service to get the claims from the token
	authHeader := ctx.GetHeader("Authorization")
	claims, err := blogController.JwtService.GetClaims(authHeader)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.New("invalid token"))
		return
	}

	// Get user role and user ID from the claims
	userRole := claims.Role
	userId, err := primitive.ObjectIDFromHex(claims.Id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.New("invalid user id"))
		return
	}

	// update only allowed for blog author or admin
	if userId == delete_request.AuthorID || userRole == "admin" {

		err := blogController.BlogUsecase.DeleteBlog(ctx, delete_request.BlogID)
		if err != nil {
			ctx.JSON(err.Code, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "blog delete successful"})
		return
	}

	ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "delete not allowed"})
}
