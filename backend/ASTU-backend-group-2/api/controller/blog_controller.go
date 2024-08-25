package controller

import (
	"context"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/api/middleware"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	custom_error "github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/errors"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/forms"
	"github.com/gin-gonic/gin"
)

// interface for blog controllers
type blogController interface {
	GetBlogs() gin.HandlerFunc
	GetBlog() gin.HandlerFunc
	CreateBlog() gin.HandlerFunc
	UpdateBlog() gin.HandlerFunc
	DeleteBlog() gin.HandlerFunc

	GetComments() gin.HandlerFunc
	CreateComment() gin.HandlerFunc
	GetComment() gin.HandlerFunc
	UpdateComment() gin.HandlerFunc
	DeleteComment() gin.HandlerFunc

	Like() gin.HandlerFunc
	DisLike() gin.HandlerFunc
}

type BlogController struct {
	BlogUsecase     entities.BlogUsecase
	CommentUsecase  entities.CommentUsecase
	UserUsecase     entities.UserUsecase
	ReactionUsecase entities.ReactionUsecase
	Env             *bootstrap.Env
}

func (bc *BlogController) GetBlogs() gin.HandlerFunc {
	return func(c *gin.Context) {

		page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
		limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
		dateFrom, _ := time.Parse(time.RFC3339, c.Query("date_from"))
		dateTo, _ := time.Parse(time.RFC3339, c.Query("date_to"))
		tags := strings.Split(c.Query("tags"), ",")
		popularityFrom, _ := strconv.Atoi(c.Query("popularity_from"))
		popularityTo, _ := strconv.Atoi(c.Query("popularity_to"))
		//in go when u split an empty string
		//Because the returned array is not empty. First element of it is an empty string ""
		if len(tags) == 1 && tags[0] == "" {
			tags = []string{}
		}
		var blogFilter entities.BlogFilter
		log.Printf("%#v\n", tags)
		blogFilter = entities.BlogFilter{
			Title:          c.Query("title"),
			Search:         c.Query("search"),
			Tags:           tags,
			DateFrom:       dateFrom,
			DateTo:         dateTo,
			Limit:          limit,
			Pages:          page,
			PopularityFrom: popularityFrom,
			PopularityTo:   popularityTo,
		}

		blogs, pagination, err := bc.BlogUsecase.GetAllBlogs(context.Background(), blogFilter)
		if err != nil {
			c.Error(err)
			return
		}

		res := entities.PaginatedResponse{
			Data:     blogs,
			MetaData: pagination,
		}

		c.JSON(http.StatusOK, res)
	}
}

func (bc *BlogController) GetBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		blogID := c.Param("id")
		blog, err := bc.BlogUsecase.GetBlogByID(context.Background(), blogID)

		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, blog)
	}
}

func (bc *BlogController) CreateBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("x-user-id").(string)
		var newBlog forms.BlogForm
		if err := c.ShouldBindJSON(&newBlog); err != nil {
			if err == io.EOF {
				c.JSON(http.StatusBadRequest, custom_error.ErrMessage(custom_error.EreInvalidRequestBody))
				return
			}
			middleware.CustomErrorResponse(c, err)
			return
		}

		user, err := bc.UserUsecase.GetUserById(context.Background(), userID)

		if err != nil {
			c.Error(err)
			return
		}

		blog, err := bc.BlogUsecase.CreateBlog(context.Background(), &newBlog, user)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(201, blog)
	}
}

func (bc *BlogController) BatchCreateBlog() gin.HandlerFunc {
	return func(c *gin.Context) {

		var newBlogs []forms.BlogForm

		if err := c.ShouldBindJSON(&newBlogs); err != nil {
			if err == io.EOF {
				c.JSON(http.StatusBadRequest, custom_error.ErrMessage(custom_error.EreInvalidRequestBody))
				return
			}
			middleware.CustomErrorResponse(c, err)
			return
		}
		userID := c.MustGet("x-user-id").(string)

		user, err := bc.UserUsecase.GetUserById(context.Background(), userID)

		if err != nil {
			c.Error(err)
			return
		}

		err = bc.BlogUsecase.BatchCreateBlog(context.Background(), &newBlogs, user)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(201, gin.H{"message": "Blogs created successfully"})
	}
}

func (bc *BlogController) UpdateBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		blogID := c.Param("id")
		userID := c.MustGet("x-user-id").(string)
		role := c.MustGet("x-user-role").(string)
		var updatedBlog forms.BlogForm

		if err := c.ShouldBindJSON(&updatedBlog); err != nil {
			if err == io.EOF {
				c.JSON(http.StatusBadRequest, custom_error.ErrMessage(custom_error.EreInvalidRequestBody))
				return
			}
			middleware.CustomErrorResponse(c, err)
			return
		}
		existingBlog, err := bc.BlogUsecase.GetBlogByID(c, blogID)
		if err != nil {
			c.Error(err)
			return
		}

		if userID != existingBlog.AuthorID && role != "admin" {
			c.JSON(http.StatusUnauthorized, custom_error.ErrorMessage{Message: "unauthorized"})
			return
		}
		blog, err := bc.BlogUsecase.UpdateBlog(context.Background(), blogID, &updatedBlog)
		if err != nil {
			c.JSON(500, custom_error.ErrorMessage{Message: err.Error()})
			return
		}
		c.JSON(200, blog)
	}
}

func (bc *BlogController) DeleteBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		blogID := c.Param("id")
		userID := c.MustGet("x-user-id").(string)
		role := c.MustGet("x-user-role").(string)

		existingBlog, err := bc.BlogUsecase.GetBlogByID(c, blogID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		if userID != existingBlog.AuthorID && role != "admin" {
			c.JSON(http.StatusUnauthorized, custom_error.ErrorMessage{Message: "unauthorized"})
			return
		}
		err = bc.BlogUsecase.DeleteBlog(context.TODO(), blogID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(204, nil)
	}
}
func (bc *BlogController) GetByTags() gin.HandlerFunc {
	return func(c *gin.Context) {
		tags, _ := c.GetQueryArray("tags")
		limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.Query("page"), 10, 64)

		blogs, pagination, err := bc.BlogUsecase.GetByTags(context.TODO(), tags, limit, page)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"blogs": blogs, "pageination": pagination})
	}
}

func (bc *BlogController) GetbyPopularity() gin.HandlerFunc {
	return func(c *gin.Context) {
		limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
		blogs, pagination, err := bc.BlogUsecase.GetByPopularity(context.Background(), limit, page)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"blogs": blogs, "pageination": pagination})
	}
}
func (bc *BlogController) SortByDate() gin.HandlerFunc {
	return func(c *gin.Context) {
		limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
		blogs, pagination, err := bc.BlogUsecase.SortByDate(context.Background(), limit, page)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"blogs": blogs, "pageination": pagination})
	}
}
func (bc *BlogController) GetComments() gin.HandlerFunc {
	return func(c *gin.Context) {
		blogID := c.Param("id")
		limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.Query("page"), 10, 64)

		comments, pageination, err := bc.CommentUsecase.GetComments(c, blogID, limit, page)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"comments": comments, "metadata": pageination})
	}
}
func (bc *BlogController) CreateComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("x-user-id").(string)
		blogID := c.Param("id")
		var commentIn forms.CommentForm

		if err := c.BindJSON(&commentIn); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		comment, err := bc.CommentUsecase.CreateComment(c, userID, blogID, &commentIn)

		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusCreated, gin.H{"comment": comment})

	}
}
func (bc *BlogController) GetComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		commentID := c.Param("comment_id")

		comment, err := bc.CommentUsecase.GetComment(c, commentID)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"comment": comment})
	}
}
func (bc *BlogController) UpdateComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		commentID := c.Param("comment_id")
		userID := c.MustGet("x-user-id").(string)
		role := c.MustGet("x-user-role").(string)
		var commentUpd entities.CommentUpdate

		if err := c.BindJSON(&commentUpd); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		existingComment, err := bc.CommentUsecase.GetComment(c, commentID)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if existingComment.UserID.Hex() != userID && role != "admin" {
			c.JSON(http.StatusUnauthorized, custom_error.ErrorMessage{Message: "unauthorized"})
			return
		}

		comment, err := bc.CommentUsecase.UpdateComment(c.Request.Context(), commentID, &commentUpd)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if existingComment.UserID.Hex() != userID && role != "admin" {
			c.JSON(http.StatusUnauthorized, custom_error.ErrorMessage{Message: "unauthorizrd"})
			return
		}

		comment, err = bc.CommentUsecase.UpdateComment(c.Request.Context(), commentID, &commentUpd)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"comment": comment})
	}
}

func (bc *BlogController) DeleteComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		commentID := c.Param("comment_id")
		userID := c.MustGet("x-user-id").(string)
		role := c.MustGet("x-user-role").(string)

		comment, err := bc.CommentUsecase.GetComment(c, commentID)
		if err != nil {
			c.Error(err)
			return
		}
		if comment.UserID.Hex() != userID && role != "admin" {
			c.JSON(http.StatusUnauthorized, custom_error.ErrorMessage{Message: "unauthorized"})
			return
		}
		err = bc.CommentUsecase.DeleteComment(c, commentID)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusNoContent, gin.H{})
	}
}
func (bc *BlogController) Like() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("x-user-id").(string)
		BlogID := c.Param("id")

		err := bc.ReactionUsecase.ToggleLike(c, BlogID, userID)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}
func (bc *BlogController) Dislike() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("x-user-id").(string)
		BlogID := c.Param("id")

		log.Println("[ctrl] disliking blog", BlogID)
		err := bc.ReactionUsecase.ToggleDislike(c, BlogID, userID)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}
