package controllers

import (
	"blogapp/Domain"
	"blogapp/Utils"
	"time"

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
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()

	err, status := cc.commentUseCase.CommentOnPost(c, comment, objID)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Comment created successfully"})
}

func (cc *CommentController) GetCommentByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := Utils.StringToObjectId(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	comment, err, statuscode := cc.commentUseCase.GetCommentByID(c, objID)
	if err != nil {
		c.JSON(statuscode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"comment": comment})
}

func (cc *CommentController) EditComment(c *gin.Context) {
	id := c.Param("id")
	objID, err := Utils.StringToObjectId(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// get comment
	existingComment, err, statuscode := cc.commentUseCase.GetCommentByID(c, objID)
	if err != nil {
		c.JSON(statuscode, gin.H{"error": err.Error()})
		return
	}

	// check if user is author of comment
	claims, err := Getclaim(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	isAuthor, err := Utils.IsAuthorOrAdmin(*claims, existingComment.AuthorID)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	if !isAuthor {
		c.JSON(401, gin.H{"error": "You are not author of this post"})
		return
	}

	var comment = &Domain.Comment{}
	if err := c.ShouldBindJSON(comment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err, status := cc.commentUseCase.EditComment(c, objID, comment)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Comment updated successfully",
		"comment": comment,})
}
