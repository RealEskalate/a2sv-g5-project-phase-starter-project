package controllers

import (
	"meleket/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CommentController handles comment-related requests
type CommentController struct {
	commentUsecase domain.CommentUsecaseInterface
}

// NewCommentController creates a new instance of CommentController
func NewCommentController(commentUsecase domain.CommentUsecaseInterface) *CommentController {
	return &CommentController{commentUsecase: commentUsecase}
}

// AddComment adds a new comment to a blog post
// AddComment adds a new comment or reply to a blog post or another comment// AddComment adds a new comment to a blog post
func (cc *CommentController) AddComment(c *gin.Context) {
	var comment domain.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract blog ID from the URL and check for validity
	blogID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	// Assign the blog ID to the comment
	comment.BlogID = blogID

	// Generate a new ObjectID for the comment itself
	comment.ID = primitive.NewObjectID()

	// Extract user ID from the claims (usually comes from middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Cast the extracted userID to ObjectID
	comment.UserID = userID.(primitive.ObjectID)

	// Set the current time for the comment creation
	comment.CreatedAt = time.Now()

	// Call the use case to add the comment to the database
	if err := cc.commentUsecase.AddComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the successfully added comment
	c.JSON(http.StatusCreated, gin.H{"message": "Comment added successfully", "comment": comment})
}

// GetCommentsByBlogID retrieves all comments for a specific blog post
func (cc *CommentController) GetCommentsByBlogID(c *gin.Context) {
	blogID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	comments, err := cc.commentUsecase.GetCommentsByBlogID(blogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

// UpdateComment updates an existing comment
func (cc *CommentController) UpdateComment(c *gin.Context) {
	commentID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	var content struct {
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&content); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.commentUsecase.UpdateComment(commentID, content.Content); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully"})
}

// DeleteComment deletes a comment by its ID
func (cc *CommentController) DeleteComment(c *gin.Context) {
	commentID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	if err := cc.commentUsecase.DeleteComment(commentID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

// AddReply adds a reply to an existing comment
// AddReply adds a reply to an existing comment
func (cc *CommentController) AddReply(c *gin.Context) {
	commentID, err := primitive.ObjectIDFromHex(c.Param("id")) // Comment ID as parameter
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	var reply domain.Reply
	if err := c.ShouldBindJSON(&reply); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign the comment ID to the reply's CommentId field
	reply.CommentId = commentID

	// Generate a new ObjectID for the reply itself
	reply.ID = primitive.NewObjectID()

	// Set reply creation time
	reply.CreatedAt = time.Now()

	// Extract user ID from claims and assign to reply's UserID field
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	reply.UserID = userID.(primitive.ObjectID)

	// Call the usecase to add the reply to the comment
	if err := cc.commentUsecase.AddReply(commentID, reply); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Reply added successfully", "reply": reply})
}

// UpdateReply updates an existing reply within a comment
func (cc *CommentController) UpdateReply(c *gin.Context) {
	commentID, err := primitive.ObjectIDFromHex(c.Param("commentID")) // Comment ID as parameter
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	replyID, err := primitive.ObjectIDFromHex(c.Param("replyID")) // Reply ID as parameter
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reply ID"})
		return
	}

	var content struct {
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&content); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.commentUsecase.UpdateReply(commentID, replyID, content.Content); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reply updated successfully"})
}

// DeleteReply deletes a reply from a comment
func (cc *CommentController) DeleteReply(c *gin.Context) {
	commentID, err := primitive.ObjectIDFromHex(c.Param("commentID")) // Comment ID as parameter
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	replyID, err := primitive.ObjectIDFromHex(c.Param("replyID")) // Reply ID as parameter
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reply ID"})
		return
	}

	if err := cc.commentUsecase.DeleteReply(commentID, replyID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reply deleted successfully"})
}
