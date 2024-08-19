package controllers

import (
	"blogapp/Domain"
	"blogapp/Utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	var tag = &Domain.Tag{}
	err := c.BindJSON(&tag)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tag.Slug = Utils.GenerateSlug(tag.Name)
	tag.ID = primitive.NewObjectID()
	tag.Posts = []primitive.ObjectID{}
	err,statuscode := tagController.tagUseCase.CreateTag(c,tag)
	if err != nil {
		c.JSON(statuscode, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, *tag)
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

// GetAllTags function
func (tagController *TagController) GetAllTags(c *gin.Context) {
	tags, err,statuscode := tagController.tagUseCase.GetAllTags(c)
	if err != nil {
		c.JSON(statuscode, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tags)
}

// Get tags by slug function
func (tagController *TagController) GetTagBySlug(c *gin.Context) {
	slug := c.Param("slug")
	tag, err,statuscode := tagController.tagUseCase.GetTagBySlug(c,slug)
	if err != nil {
		c.JSON(statuscode, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tag)
}

