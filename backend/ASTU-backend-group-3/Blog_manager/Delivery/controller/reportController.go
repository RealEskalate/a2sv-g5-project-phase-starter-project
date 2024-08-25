package controller

import (
	"ASTU-backend-group-3/Blog_manager/Domain"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "net/http"
    "time"
)

func (h *BlogController) BlogReport(c *gin.Context) {
	blogID := c.Param("id")
	var input struct {
		Content string   `json:"content" binding:"required"`
		Tag   string `json:"tag" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username := c.GetString("username")
	userID, err := h.userUsecase.GiveId(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot get a user"})
		return

	}

	report := &Domain.Report{
		Id:        primitive.NewObjectID().Hex(),
		Content:   input.Content,
		BlogId:    blogID,
		UserId:    userID,
		CreatedAt: time.Now().Format(time.RFC3339),
		Tag:       input.Tag,
	}


	err = h.blogUsecase.ReportBlog(blogID, *report)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog reported successfully"})
}
 