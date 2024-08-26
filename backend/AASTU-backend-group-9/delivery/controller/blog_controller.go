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

func getclaim(c *gin.Context) (*domain.JwtCustomClaims, *domain.Error) {
	claim, exists := c.Get("claim")
	if !exists {
		customError := domain.Error{
			StatusCode: http.StatusUnauthorized,
			Message:    "Unauthorized access",
		}
		return nil, &customError
	}

	userClaims, ok := claim.(domain.JwtCustomClaims)
	if !ok {
		customError := domain.Error{
			Err:        errors.New("error asserting claims"),
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
		}

		return nil, &customError
	}

	return &userClaims, nil
}

func (bc *BlogController) CreateBlog(c *gin.Context) {
	claims, err := getclaim(c)
	if err != nil {

		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}
	fmt.Println(claims)
	var req domain.BlogCreationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		customError := domain.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request data",
		}
		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}

	blog, err := bc.BlogUsecase.CreateBlog(c.Request.Context(), &req, claims)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}
	blog.AuthorName = claims.Username

	success := domain.Error{
		StatusCode: http.StatusCreated,
		Message:    "Blog created successfully",
	}

	c.JSON(success.StatusCode, success.Message)
}

func (bc *BlogController) GetBlogByID(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		customError := domain.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID format",
		}
		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}

	blog, customError := bc.BlogUsecase.GetBlogByID(c.Request.Context(), id)
	if customError != nil {
		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}
	success := domain.Error{
		StatusCode: http.StatusOK,
	}

	c.JSON(success.StatusCode, blog)
}

func (bc *BlogController) GetAllBlogs(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		err := domain.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid limit format",
		}
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}
	sortBy := c.DefaultQuery("sortBy", "likes")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		err := domain.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid page format",
		}

		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	blogs, customError := bc.BlogUsecase.GetAllBlogs(c.Request.Context(), pageInt, limit, sortBy)
	if customError != nil {
		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}

	success := domain.Error{
		StatusCode: http.StatusOK,
	}

	c.JSON(success.StatusCode, blogs)
}

func (bc *BlogController) UpdateBlog(c *gin.Context) {
	claims, er := getclaim(c)
	if er != nil {
		c.JSON(er.StatusCode, gin.H{"error": er.Message})
		return
	}

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		err := domain.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID format",
		}

		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	blog, customError := bc.BlogUsecase.GetBlogByID(c.Request.Context(), id)
	if customError != nil {
		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}
	if claims.UserID != blog.AuthorID {
		customError := domain.Error{
			StatusCode: http.StatusUnauthorized,
			Message:    "You are not authorized to update this blog",
		}
		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}

	var newBlog domain.BlogUpdateRequest
	if err := c.ShouldBindJSON(&newBlog); err != nil {
		err := domain.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request data",
		}

		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	blogs, customError := bc.BlogUsecase.UpdateBlog(c.Request.Context(), id, &newBlog)
	if customError != nil {
		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}
	success := domain.Error{
		StatusCode: http.StatusOK,
	}

	c.JSON(success.StatusCode, blogs)
}

func (bc *BlogController) DeleteBlog(c *gin.Context) {
	claims, er := getclaim(c)
	if er != nil {
		c.JSON(er.StatusCode, gin.H{"error": er.Message})
		return
	}

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		err := domain.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID format",
		}

		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	blog, customError := bc.BlogUsecase.GetBlogByID(c.Request.Context(), id)
	if customError != nil {
		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}
	if claims.UserID != blog.AuthorID {
		customError := domain.Error{
			StatusCode: http.StatusUnauthorized,
			Message:    "You are not authorized to delete this blog",
		}

		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}

	if customError := bc.BlogUsecase.DeleteBlog(c.Request.Context(), id); customError != nil {
		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}
	success := domain.Error{
		StatusCode: http.StatusOK,
		Message:    "Blog deleted successfully",
	}

	c.JSON(success.StatusCode, gin.H{"message": success.Message})
}

// Delivery/controllers/blog_controller.go
// controller/blog_controller.go
func (bc *BlogController) SearchBlogs(c *gin.Context) {
	title := c.Query("title")
	author := c.Query("author")

	// Call the use case with the search criteria
	blogs, customError := bc.BlogUsecase.SearchBlogs(c, title, author)
	if customError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": customError.Message})
		return
	}
	success := domain.Error{
		StatusCode: http.StatusOK,
	}

	c.JSON(success.StatusCode, blogs)
}

func (bc *BlogController) FilterBlogs(c *gin.Context) {
	tags := c.QueryArray("tags")
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")
	popularity := c.Query("popularity")
	blogs, customError := bc.BlogUsecase.FilterBlogs(c.Request.Context(), popularity, tags, startDate, endDate)
	if customError != nil {
		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}
	success := domain.Error{
		StatusCode: http.StatusOK,
	}
	c.JSON(success.StatusCode, blogs)
}
func (bc *BlogController) TrackView(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		err := domain.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID format",
		}

		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}
	customError := bc.BlogUsecase.TrackView(c.Request.Context(), id)
	if customError != nil {
		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}
	success := domain.Error{
		StatusCode: http.StatusOK,
		Message:    "View tracked successfully",
	}
	c.JSON(success.StatusCode, gin.H{"message": success.Message})
}

func (bc *BlogController) TrackLike(c *gin.Context) {
	id, er := primitive.ObjectIDFromHex(c.Param("id"))

	if er != nil {
		er := domain.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID format",
		}

		c.JSON(er.StatusCode, gin.H{"error": er.Message})
		return

	}
	claims, err := getclaim(c)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}
	userID := claims.UserID
	customError := bc.BlogUsecase.TrackLike(c.Request.Context(), id, userID)
	if customError != nil {
		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}
	success := domain.Error{
		StatusCode: http.StatusOK,
		Message:    "Like tracked successfully",
	}
	c.JSON(success.StatusCode, gin.H{"message": success.Message})
}

func (bc *BlogController) TrackDislike(c *gin.Context) {
	id, er := primitive.ObjectIDFromHex(c.Param("id"))
	if er != nil {
		err := domain.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID format",
		}

		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}
	claims, err := getclaim(c)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}
	userID := claims.UserID
	customError := bc.BlogUsecase.TrackDislike(c.Request.Context(), id, userID)
	if customError != nil {
		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}
	success := domain.Error{
		StatusCode: http.StatusOK,
		Message:    "Dislike tracked successfully",
	}
	c.JSON(success.StatusCode, gin.H{"message": success.Message})

}

func (bc *BlogController) AddComment(c *gin.Context) {
	post_id, er := primitive.ObjectIDFromHex(c.Param("id"))
	if er != nil {
		err := domain.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID format",
		}

		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	claims, err := getclaim(c)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}
	var comment domain.Comment
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	userID := claims.UserID
	customError := bc.BlogUsecase.AddComment(c.Request.Context(), post_id, userID, &comment)
	if customError != nil {
		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}
	success := domain.Error{
		StatusCode: http.StatusOK,
		Message:    "Comment added successfully",
	}
	c.JSON(success.StatusCode, gin.H{"message": success.Message})
}

func (bc *BlogController) AddReply(c *gin.Context) {

	post_id, er := primitive.ObjectIDFromHex(c.Param("id"))
	if er != nil {
		err := domain.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID format",
		}

		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}
	comment_id, _ := primitive.ObjectIDFromHex(c.Param("comment_id"))
	claims, err := getclaim(c)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}
	var reply domain.Comment
	if err := c.BindJSON(&reply); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	userID := claims.UserID
	customError := bc.BlogUsecase.AddReply(c.Request.Context(), post_id, comment_id, userID, &reply)
	if customError != nil {
		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}
	success := domain.Error{
		StatusCode: http.StatusOK,
		Message:    "Reply added successfully",
	}

	c.JSON(success.StatusCode, success.Message)
}

func (bc *BlogController) TrackCommentPopularity(c *gin.Context) {
	post_id, er := primitive.ObjectIDFromHex(c.Param("id"))

	if er != nil {
		err := domain.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID format",
		}

		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return

	}
	comment_id, _ := primitive.ObjectIDFromHex(c.Param("comment_id"))
	metric := c.Query("metric")
	claims, err := getclaim(c)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}
	userID := claims.UserID
	customError := bc.BlogUsecase.TrackCommentPopularity(c.Request.Context(), post_id, comment_id, userID, metric)
	if customError != nil {
		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}
	success := domain.Error{
		StatusCode: http.StatusOK,
		Message:    "Comment popularity tracked successfully",
	}
	c.JSON(success.StatusCode, gin.H{"message": success.Message})

}

func (bc *BlogController) GetComments(c *gin.Context) {
	// Convert the post ID from the URL parameter to an ObjectID
	postID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		err := domain.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID format",
		}

		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	// Retrieve the comments from the use case
	comments, customError := bc.BlogUsecase.GetComments(c.Request.Context(), postID)
	if customError != nil {
		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
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
	success := domain.Error{
		StatusCode: http.StatusOK,
	}
	c.JSON(success.StatusCode, response)
}

func (bc *BlogController) DeleteComment(c *gin.Context) {
	// Convert the post ID from the URL parameter to an ObjectID
	postID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		err := domain.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID format",
		}

		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	// Convert the comment ID from the URL parameter to an ObjectID
	commentID, err := primitive.ObjectIDFromHex(c.Param("comment_id"))
	if err != nil {
		err := domain.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID format",
		}

		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	// Retrieve the claims (user data)
	claims, er := getclaim(c)
	if er != nil {
		c.JSON(er.StatusCode, gin.H{"error": er.Message})
		return
	}

	// Extract the userID from claims
	userID := claims.UserID

	// Call the use case to delete the comment
	customError := bc.BlogUsecase.DeleteComment(c.Request.Context(), postID, commentID, userID)
	if customError != nil {
		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}

	// Send a success response
	success := domain.Error{
		StatusCode: http.StatusOK,
		Message:    "Comment deleted successfully",
	}
	c.JSON(success.StatusCode, success.Message)
}

func (bc *BlogController) UpdateComment(c *gin.Context) {
	post_id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		err := domain.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID format",
		}

		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	// Convert the comment ID from the URL parameter to an ObjectID
	comment_id, err := primitive.ObjectIDFromHex(c.Param("comment_id"))
	if err != nil {
		err := domain.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID format",
		}

		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}
	claims, er := getclaim(c)
	if er != nil {
		c.JSON(er.StatusCode, gin.H{"error": er.Message})
		return
	}
	userID := claims.UserID
	var comment domain.Comment
	if err := c.BindJSON(&comment); err != nil {
		err := domain.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request",
		}

		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}
	customError := bc.BlogUsecase.UpdateComment(c.Request.Context(), post_id, comment_id, userID, &comment)
	if customError != nil {
		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}

	success := domain.Error{
		StatusCode: http.StatusOK,
		Message:    "Comment updated successfully",
	}
	c.JSON(success.StatusCode, gin.H{"message": success.Message})
}
