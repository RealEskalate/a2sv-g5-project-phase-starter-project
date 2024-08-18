package controllers

import (
	"blogapp/Domain"
	"blogapp/Utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentController struct {
	commentUseCase Domain.CommentUseCase
}

func NewCommentController(usecase Domain.CommentUseCase) *CommentController {
	return &CommentController{
		commentUseCase: usecase,
	}
}

func (cc *CommentController) CommentOnPost(c *gin.Context) {
	claim, err := Getclaim(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	objID, err := Utils.StringToObjectId(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var comment = &Domain.Comment{}
	if err := c.ShouldBindJSON(comment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	comment.ID = primitive.NewObjectID()
	comment.PostID = objID
	comment.AuthorID = claim.ID

	err, status := cc.commentUseCase.CommentOnPost(c, comment, objID)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Comment created successfully"})
}
