package blogcontroller

import (
	"blogs/config"
	"blogs/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (b *BlogController) UpdateBlogByID(ctx *gin.Context) {
	idHex := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var blogUpdate struct {
		Title   string   `json:"title"`
		Content string   `json:"content"`
		Tags    []string `json:"tags"`
	}

	if err := ctx.ShouldBindJSON(&blogUpdate); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if blogUpdate.Title == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "title cannot be empty"})
		return
	}

	if blogUpdate.Content == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "content cannot be empty"})
		return
	}

	claim, ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok {
		log.Println("Error getting claims")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	blog := &domain.Blog{
		Title:   blogUpdate.Title,
		Content: blogUpdate.Content,
		Tags:    blogUpdate.Tags,
	}

	newBlog, err := b.BlogUsecase.UpdateBlogByID(id.Hex(), blog, claim)
	if err != nil {
		code := config.GetStatusCode(err)

		if code == http.StatusInternalServerError {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		ctx.JSON(code, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, newBlog)
}
