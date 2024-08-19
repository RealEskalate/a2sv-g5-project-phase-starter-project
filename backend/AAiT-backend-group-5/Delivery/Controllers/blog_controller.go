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
	blogID := ctx.Param("id")

	blog, err := blogController.BlogUsecase.GetBlog(ctx, blogID)
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
	var updated_request dtos.UpdateBlogRequest

	if e := ctx.BindJSON(updated_request); e != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// get user ID from the context
	userId, exists := ctx.Get("id")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid token"})
		return
	}

	// check if the user matches the authorID of the blog
	if userId != updated_request.AuthorID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user"})
		return
	}

	// set fields for the updated blog
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

	err := blogController.BlogUsecase.UpdateBlog(ctx, updated_request.BlogID, &updated_blog)
	if err != nil {
		ctx.JSON(err.Code, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "blog update successful"})
}

func (blogController BlogController) DeleteBlog(ctx *gin.Context) {
	var delete_request dtos.DeleteBlogRequest

	if e := ctx.BindJSON(delete_request); e != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// get user ID and role from the context
	userId, exists := ctx.Get("id")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid token"})
		return
	}

	userRole, exists := ctx.Get("role")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid token"})
		return
	}

	// check if the user matches the authorID of the blog or if user is an admin
	if userId != delete_request.AuthorID && userRole != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user"})
		return
	}

	err := blogController.BlogUsecase.DeleteBlog(ctx, delete_request.BlogID)
	if err != nil {
		ctx.JSON(err.Code, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "blog delete successful"})
}

func (blogController BlogController) HandelTrackPopularity(ctx *gin.Context) {
	var request dtos.TrackPopularityRequest

	// attempt to bind payload to the request
	err := ctx.ShouldBind(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	// get userId from context
	userId := ctx.GetString("id")
	if userId == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user"})
		return
	}

	// append userId to the UserID field of the request
	request.UserID = userId

	// invoke TrackPopularity Usecase
	e := blogController.BlogUsecase.TrackPopularity(ctx, request.BlogID, request)
	if e != nil {
		ctx.JSON(e.Code, e.Error())
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "successful"})
}
