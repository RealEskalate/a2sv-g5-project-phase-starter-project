package controllers

import (
	"blogapp/Domain"
	"blogapp/Utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TagController struct
type TagController struct {
	tagUseCase Domain.TagUseCase
}

func NewTagsController(usecase Domain.TagUseCase) *TagController {
	return &TagController{
		tagUseCase: usecase,
	}
}

// CreateTag function
func (tagController *TagController) CreateTag(c *gin.Context) {
	var tag Domain.Tag
	err := c.BindJSON(&tag)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err,statuscode := tagController.tagUseCase.CreateTag(c,&tag)
	if err != nil {
		c.JSON(statuscode, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, tag)
}

// DeleteTag function
func (tagController *TagController) DeleteTag(c *gin.Context) {
	id, err := Utils.StringToObjectId(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err,statuscode := tagController.tagUseCase.DeleteTag(c,id)
	if err != nil {
		c.JSON(statuscode, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tag deleted successfully"})
}
