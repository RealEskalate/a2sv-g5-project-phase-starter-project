package blog

import (
	"blogApp/internal/ai"
	"blogApp/internal/domain"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *BlogHandler) GetAiBlog(c *gin.Context) {
	userClaims, err := GetClaims(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	userId := userClaims.UserID

	var body struct {
		Query string `json:"query"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(body)

	blog, err := ai.GetAiBlog(userId, body.Query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blog)
}

func (h *BlogHandler) ModerateBlog(c *gin.Context) {
	blog := new(domain.Blog)
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	grade, message, err := ai.ModerateBlog(blog.Content, blog.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"grade": grade, "message": message})

}
