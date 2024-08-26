package controller

import (
	"blog/domain"
	"errors"
	"fmt"
	"net/http"

	// "os/user"
	"strconv"

	"blog/config"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogController struct {
	BlogUsecase domain.BlogUsecase

	Env *config.Env
}

func getclaim(c *gin.Context) (*domain.JwtCustomClaims, error) {
	claim, exists := c.Get("claim")
	if !exists {
		return nil, errors.New("claim not set")
	}

	userClaims, ok := claim.(domain.JwtCustomClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return &userClaims, nil
}

func (bc *BlogController) CreateBlog(c *gin.Context) {
	claims, err := getclaim(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(claims)
	var req domain.BlogCreationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blog, err := bc.BlogUsecase.CreateBlog(c.Request.Context(), &req, claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	blog.AuthorName = claims.Username

	c.JSON(http.StatusCreated, blog)
}

func (bc *BlogController) GetBlogByID(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	blog, err := bc.BlogUsecase.GetBlogByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blog)
}

func (bc *BlogController) GetAllBlogs(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit format"})
		return
	}
	sortBy := c.DefaultQuery("sortBy", "likes")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page format"})
		return
	}

	blogs, err := bc.BlogUsecase.GetAllBlogs(c.Request.Context(), pageInt, limit, sortBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

func (bc *BlogController) UpdateBlog(c *gin.Context) {
	claims, err := getclaim(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	blog, err := bc.BlogUsecase.GetBlogByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if claims.UserID != blog.AuthorID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to update this blog"})
		return
	}

	var newBlog domain.BlogUpdateRequest
	if err := c.ShouldBindJSON(&newBlog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blogs, err := bc.BlogUsecase.UpdateBlog(c.Request.Context(), id, &newBlog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

func (bc *BlogController) DeleteBlog(c *gin.Context) {
	claims, err := getclaim(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	blog, err := bc.BlogUsecase.GetBlogByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}
	if claims.UserID != blog.AuthorID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to update this blog"})
		return
	}

	if err := bc.BlogUsecase.DeleteBlog(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}

// Delivery/controllers/blog_controller.go
// controller/blog_controller.go
func (bc *BlogController) SearchBlogs(c *gin.Context) {
	title := c.Query("title")
	author := c.Query("author")

	// Call the use case with the search criteria
	blogs, err := bc.BlogUsecase.SearchBlogs(c, title, author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

// func (bc *BlogController) FilterBlogsByTags(c *gin.Context) {
// 	tags := c.QueryArray("tags")
// 	blogs, err := bc.BlogUsecase.FilterBlogsByTags(c.Request.Context(), tags)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, blogs)
// }

// func (bc *BlogController) FilterBlogsByDate(c *gin.Context) {
// 	date := c.Query("date")
// 	blogs, err := bc.BlogUsecase.FilterBlogsByDate(c.Request.Context(), date)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, blogs)
// }

func (bc *BlogController) FilterBlogs(c *gin.Context) {
	tags := c.QueryArray("tags")
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")
	popularity := c.Query("popularity")
	blogs, err := bc.BlogUsecase.FilterBlogs(c.Request.Context(), popularity, tags, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blogs)
}
func (bc *BlogController) TrackView(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	err := bc.BlogUsecase.TrackView(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "View tracked successfully"})
}

func (bc *BlogController) TrackLike(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	claims, err := getclaim(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	userID := claims.UserID
	err = bc.BlogUsecase.TrackLike(c.Request.Context(), id, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Like tracked successfully"})
}

func (bc *BlogController) TrackDislike(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	claims, err := getclaim(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	userID := claims.UserID
	err = bc.BlogUsecase.TrackDislike(c.Request.Context(), id, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Dislike tracked successfully"})
}

func (bc *BlogController) AddComment(c *gin.Context) {
	post_id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	claims, err := getclaim(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	var comment domain.Comment
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	userID := claims.UserID
	err = bc.BlogUsecase.AddComment(c.Request.Context(), post_id, userID, &comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment added successfully"})
}

func (bc *BlogController) AddReply(c *gin.Context) {

	post_id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	comment_id, _ := primitive.ObjectIDFromHex(c.Param("comment_id"))
	claims, err := getclaim(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	var reply domain.Comment
	if err := c.BindJSON(&reply); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	userID := claims.UserID
	err = bc.BlogUsecase.AddReply(c.Request.Context(), post_id, comment_id, userID, &reply)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Reply added successfully"})
}


func (bc *BlogController) TrackCommentPopularity(c *gin.Context) {
	post_id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	comment_id, _ := primitive.ObjectIDFromHex(c.Param("comment_id"))
	metric := c.Query("metric")
	claims, err := getclaim(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	userID := claims.UserID
	err = bc.BlogUsecase.TrackCommentPopularity(c.Request.Context(), post_id, comment_id, userID, metric)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Like tracked successfully"})
}


func (bc *BlogController) GetComments(c *gin.Context) {
	// Convert the post ID from the URL parameter to an ObjectID
	postID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	// Retrieve the comments from the use case
	comments, err := bc.BlogUsecase.GetComments(c.Request.Context(), postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Prepare the response, transforming each comment into the response structure
	response := make([]domain.ResponseComment, len(comments))
	for i, comment := range comments {
		response[i] = domain.ResponseComment{
			AuthorID: comment.AuthorID,
			Comments: comment.Content,
		}
	}

	// Send the response with a status code of 200 (OK)
	c.JSON(http.StatusOK, response)
}

func (bc *BlogController) DeleteComment(c *gin.Context) {
	// Convert the post ID from the URL parameter to an ObjectID
	postID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	// Convert the comment ID from the URL parameter to an ObjectID
	commentID, err := primitive.ObjectIDFromHex(c.Param("comment_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	// Retrieve the claims (user data)
	claims, err := getclaim(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}

	// Extract the userID from claims
	userID := claims.UserID

	// Call the use case to delete the comment
	err = bc.BlogUsecase.DeleteComment(c.Request.Context(), postID, commentID, userID)
	if err != nil {
		if err.Error() == "you are not authorized to delete this comment" {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Send a success response
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

func (bc *BlogController) UpdateComment(c *gin.Context) {
	post_id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	comment_id, _ := primitive.ObjectIDFromHex(c.Param("comment_id"))
	claims, err := getclaim(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	userID := claims.UserID
	var comment domain.Comment
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	err = bc.BlogUsecase.UpdateComment(c.Request.Context(), post_id, comment_id, userID, &comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully"})
}

