package controllers

import (
	"AAIT-backend-group-3/internal/domain/models"
	"AAIT-backend-group-3/internal/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommentControllerInterface interface {
	CreateComment(c *gin.Context)
	GetCommentByID(c *gin.Context)
	EditComment(c *gin.Context)
	DeleteComment(c *gin.Context)
	GetCommentsByIDList(c *gin.Context)
	GetCommentByAuthorID(c *gin.Context)
	GetCommentByBlogID(c *gin.Context)
}

type CommentController struct {
	commentUsecase usecases.CommentUsecaseInterface
}

func NewCommentController(commentUsecase usecases.CommentUsecaseInterface) CommentControllerInterface {
	return &CommentController{
		commentUsecase: commentUsecase,
	}
}

func (cc *CommentController) CreateComment(c *gin.Context) {
	var comment models.Comment
	blogID := c.Param("blogID")

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	commentID, err := cc.commentUsecase.CreateComment(&comment, blogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"comment_id": commentID})
}

func (cc *CommentController) GetCommentByID(c *gin.Context) {
	commentID := c.Param("commentID")

	comment, err := cc.commentUsecase.GetCommentByID(commentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (cc *CommentController) EditComment(c *gin.Context) {
	commentID := c.Param("commentID")
	var updatedComment models.Comment

	if err := c.ShouldBindJSON(&updatedComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := cc.commentUsecase.EditComment(commentID, &updatedComment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment updated"})
}

func (cc *CommentController) DeleteComment(c *gin.Context) {
	commentID := c.Param("commentID")  
	blogID := c.Query("blogID") 
	

	if commentID == "" || blogID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "commentID or blogID cannot be empty"})
		return
	}
	comment, err := cc.commentUsecase.GetCommentByID(commentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	authorID := comment.AuthorID.Hex()
	userID := c.GetString("userId")
	if authorID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own comment"})
		return
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = cc.commentUsecase.DeleteComment(blogID, commentID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted"})
}
}

func (cc *CommentController) GetCommentsByIDList(c *gin.Context) {
	var commentIDs []string

	if err := c.ShouldBindJSON(&commentIDs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	comments, err := cc.commentUsecase.GetCommentsByIDList(commentIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (cc *CommentController) GetCommentByAuthorID(c *gin.Context) {
	authorID := c.Param("authorID")

	comments, err := cc.commentUsecase.GetCommentByAuthorID(authorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}


func (cc *CommentController) GetCommentByBlogID(c *gin.Context) {
	blogID := c.Param("blogID")
	comments, err := cc.commentUsecase.GetCommentByBlogID(blogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}