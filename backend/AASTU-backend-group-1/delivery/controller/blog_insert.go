package controller

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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})

		return
	}

	if blog.Title == "" {
		ctx.JSON(http.StatusBadRequest, "title is required")
		return
	}

	if blog.Content == "" {
		ctx.JSON(http.StatusBadRequest, "content is required")
		return
	}

	if len(blog.Tags) == 0 {
		blog.Tags = []string{}
	}

	blogData := &domain.Blog{
		Title:         blog.Title,
		Content:       blog.Content,
		Tags:          blog.Tags,
		CreatedAt:     time.Now(),
		LastUpdatedAt: time.Now(),
	}

	err := b.BlogUsecase.InsertBlog(blogData)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "blog created",
		"blog":    blogData,
	})
}
