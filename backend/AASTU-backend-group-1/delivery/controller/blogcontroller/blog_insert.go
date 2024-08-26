package blogcontroller

import (
	"blogs/domain"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (b *BlogController) InsertBlog(ctx *gin.Context) {
	var blog struct {
		Title   string   `json:"title"`
		Content string   `json:"content"`
		Tags    []string `json:"tags"`
	}

	if err := ctx.ShouldBindJSON(&blog); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	
	if blog.Title == "" {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   "title is required",
		})
		return
	}

	if blog.Content == "" {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   "content is required",
		})
		return
	}

	if len(blog.Tags) == 0 {
		blog.Tags = []string{}
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

	blogData := &domain.Blog{
		Title:         blog.Title,
		Content:       blog.Content,
		Author:        claim.Username,
		Tags:          blog.Tags,
		CreatedAt:     time.Now(),
		LastUpdatedAt: time.Now(),
		ViewsCount:    0,
		LikesCount:    0,
		CommentsCount: 0,
	}

	newblog, err := b.BlogUsecase.InsertBlog(blogData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, domain.APIResponse{
		Status:  http.StatusCreated,
		Message: "Blog created",
		Data:    newblog,
	})
}
