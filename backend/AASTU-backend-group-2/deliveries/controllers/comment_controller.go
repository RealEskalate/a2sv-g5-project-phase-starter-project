package controllers

import (
	"blog_g2/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	CommentUsecase domain.CommentUsecase
}

func NewCommentController(Commmgr domain.CommentUsecase) *CommentController {
	return &CommentController{
		CommentUsecase: Commmgr,
	}
}

func (cc *CommentController) CreateComment(c *gin.Context) {
	blogID := c.Param("blog_id")

	log.Println(blogID)

	var comment domain.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetString("userid")

	log.Println(blogID)

	err := cc.CommentUsecase.CreateComment(c, blogID, userID, comment)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Comment created successfully"})
}

func (cc *CommentController) GetComment(c *gin.Context) {
	blogID := c.Param("blog_id")

	comments, err := cc.CommentUsecase.GetComments(c, blogID)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (cc *CommentController) UpdateComment(c *gin.Context) {
	commentID := c.Param("comm_id")

	var comment domain.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := cc.CommentUsecase.UpdateComment(c, commentID, comment)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully"})
}

func (cc *CommentController) DeleteComment(c *gin.Context) {
	commentID := c.Param("comm_id")

	err := cc.CommentUsecase.DeleteComment(c, commentID)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
