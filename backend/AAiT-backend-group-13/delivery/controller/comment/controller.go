package commentcontroller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/group13/blog/delivery/common"
	basecontroller "github.com/group13/blog/delivery/controller/base"
	er "github.com/group13/blog/domain/errors"
	"github.com/group13/blog/domain/models"
	commentcmd "github.com/group13/blog/usecase/comment/command"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
)

type CommentController struct {
	basecontroller.BaseHandler
	addcomHandler     icmd.IHandler[*commentcmd.AddCommand, *models.Comment]
	deletecomHandler  icmd.IHandler[uuid.UUID, bool]
	getcomHandler     icmd.IHandler[uuid.UUID, *models.Comment]
	getBlogComHandler icmd.IHandler[uuid.UUID, *[]models.Comment]
}

var _ common.IController = &CommentController{}

type Config struct {
	AddcomHandler     icmd.IHandler[*commentcmd.AddCommand, *models.Comment]
	DeletecomHandler  icmd.IHandler[uuid.UUID, bool]
	GetcomHandler     icmd.IHandler[uuid.UUID, *models.Comment]
	GetBlogComHandler icmd.IHandler[uuid.UUID, *[]models.Comment]
}

// New creates a new Comment Controller with the given CQRS handler.
func New(config Config) *CommentController {
	return &CommentController{
		addcomHandler:     config.AddcomHandler,
		deletecomHandler:  config.DeletecomHandler,
		getcomHandler:     config.GetcomHandler,
		getBlogComHandler: config.GetBlogComHandler,
	}
}

// RegisterPublic registers public routes.
func (c *CommentController) RegisterPublic(route *gin.RouterGroup) {

}

func (c *CommentController) RegisterPrivileged(route *gin.RouterGroup) {}

// RegisterProtected registers protected routes.
func (c *CommentController) RegisterProtected(route *gin.RouterGroup) {
	comments := route.Group("/blogs")
	{
		comments.GET("/:id/comments", c.GetBlogComments)
		comments.GET("/:id/comments/:commentId", c.GetCommentById)
		comments.POST("/:id/comments", c.AddComment)
		comments.DELETE("/:id/comments/:commentId", c.DeleteComment)
	}
}

func (c *CommentController) GetBlogComments(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))

	if err != nil {
		c.BaseHandler.Respond(ctx, http.StatusBadRequest, gin.H{"error": "Id is invalid format"})
		return
	}

	cs, err := c.getBlogComHandler.Handle(id)

	comments := []CommentResponse{}

	for _, val := range *cs {
		comments = append(comments, *FromComment(&val))
	}

	if err != nil {
		c.BaseHandler.Respond(ctx, http.StatusNotFound, gin.H{"error": "blog not found"})
		return
	}

	c.BaseHandler.Respond(ctx, http.StatusOK, comments)

}

func (c *CommentController) GetCommentById(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("commentId"))

	if err != nil {
		c.BaseHandler.Respond(ctx, http.StatusBadRequest, gin.H{"error": "Id is invalid format"})
		return
	}

	comment, err := c.getcomHandler.Handle(id)

	if err != nil {
		c.BaseHandler.Respond(ctx, http.StatusNotFound, gin.H{"error": "blog not found"})
		return
	}

	c.BaseHandler.Respond(ctx, http.StatusOK, FromComment(comment))

}

func (c *CommentController) AddComment(ctx *gin.Context) {
	id := ctx.Param("id")
	decoded, err := uuid.Parse(id)

	if err != nil {
		c.BaseHandler.Respond(ctx, http.StatusBadRequest, gin.H{"error": "blog id is invalid format"})
		return
	}
	var coment CommentDto

	if err := ctx.ShouldBindJSON(&coment); err != nil {
		c.BaseHandler.Respond(ctx, http.StatusBadRequest, er.NewBadRequest(err.Error()))
		return
	}

	decodedUser, err := uuid.Parse(coment.UserId)

	if err != nil {
		c.BaseHandler.Respond(ctx, http.StatusBadRequest, gin.H{"error": "Id is invalid format"})
		return
	}

	command := commentcmd.NewAddCommand(coment.Content, uuid.UUID(decoded), decodedUser)
	com, err := c.addcomHandler.Handle(command)

	if err != nil {
		c.BaseHandler.Respond(ctx, http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	baseURL := fmt.Sprintf("http://%s", ctx.Request.Host)
	resourceLocation := fmt.Sprintf("%s%s/%s", baseURL, ctx.Request.URL.Path, com.ID().String())
	c.RespondWithLocation(ctx, http.StatusCreated, nil, resourceLocation)

}

func (c *CommentController) DeleteComment(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("commentId"))

	if err != nil {
		c.BaseHandler.Respond(ctx, http.StatusBadRequest, gin.H{"error": "Id is invalid format"})

	}

	r, err := c.deletecomHandler.Handle(id)

	if err != nil {
		c.BaseHandler.Respond(ctx, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !r {
		c.BaseHandler.Respond(ctx, http.StatusNotFound, gin.H{"error": "comment not found"})
		return
	}

	c.BaseHandler.Respond(ctx, http.StatusOK, gin.H{})

}
