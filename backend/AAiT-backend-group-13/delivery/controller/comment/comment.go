package commentcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	basecontroller "github.com/group13/blog/delivery/base"
	common "github.com/group13/blog/delivery/common/icontroller"
	"github.com/group13/blog/delivery/controller/dto"
	er "github.com/group13/blog/domain/errors"
	"github.com/group13/blog/domain/models/comment"
	addcom "github.com/group13/blog/usecase/comment/command/add"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
)

type CommentController struct {
	basecontroller.BaseHandler
	addcomHandler     icmd.IHandler[*addcom.Command, *comment.Comment]
	deletecomHandler  icmd.IHandler[uuid.UUID, bool]
	getcomHandler     icmd.IHandler[uuid.UUID, *comment.Comment]
	getBlogComHandler icmd.IHandler[uuid.UUID, *[]comment.Comment]
}

var _ common.IController = &CommentController{}

type Config struct {
	basecontroller.BaseHandler
	AddcomHandler     icmd.IHandler[*addcom.Command, *comment.Comment]
	DeletecomHandler  icmd.IHandler[uuid.UUID, bool]
	GetcomHandler     icmd.IHandler[uuid.UUID, *comment.Comment]
	GetBlogComHandler icmd.IHandler[uuid.UUID, *[]comment.Comment]
}

// New creates a new Comment Controller with the given CQRS handler.
func New(config Config) *CommentController {
	return &CommentController{
		BaseHandler:       config.BaseHandler,
		addcomHandler:     config.AddcomHandler,
		deletecomHandler:  config.DeletecomHandler,
		getcomHandler:     config.GetcomHandler,
		getBlogComHandler: config.GetBlogComHandler,
	}
}

// RegisterPublic registers public routes.
func (c *CommentController) RegisterPublic(route *gin.RouterGroup) {}

func (c *CommentController) RegisterPrivileged(route *gin.RouterGroup) {}

// RegisterProtected registers protected routes.
func (c *CommentController) RegisterProtected(route *gin.RouterGroup) {
	comments := route.Group("/blogs")
	{
		comments.GET("/:id/comments", c.GetBlogComments)
		comments.GET("/:id/comments/:id", c.GetCommentById)
		comments.POST("/:id/comments", c.AddComment)
		comments.DELETE("/:id/comments/:id", c.DeleteComment)
	}
}

func (c *CommentController) GetBlogComments(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))

	if err != nil {
		c.BaseHandler.Respond(ctx, http.StatusBadRequest, gin.H{"error": "Id is invalid format"})
		return
	}

	comments, err := c.getBlogComHandler.Handle(id)

	if err != nil {
		c.BaseHandler.Respond(ctx, http.StatusNotFound, gin.H{"error": "blog not found"})
		return
	}

	c.BaseHandler.Respond(ctx, http.StatusOK, comments)

}

func (c *CommentController) GetCommentById(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))

	if err != nil {
		c.BaseHandler.Respond(ctx, http.StatusBadRequest, gin.H{"error": "Id is invalid format"})
		return
	}

	comment, err := c.getcomHandler.Handle(id)

	if err != nil {
		c.BaseHandler.Respond(ctx, http.StatusNotFound, gin.H{"error": "blog not found"})
		return
	}

	c.BaseHandler.Respond(ctx, http.StatusOK, comment)

}

func (c *CommentController) AddComment(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))

	if err != nil {
		c.BaseHandler.Respond(ctx, http.StatusBadRequest, gin.H{"error": "Id is invalid format"})
		return
	}
	var coment dto.CommentDto

	if err := ctx.ShouldBindJSON(&coment); err != nil {
		c.BaseHandler.Respond(ctx, http.StatusBadRequest, er.NewBadRequest(err.Error()))
		return
	}

	command := addcom.NewCommand(coment.Content, id, coment.UserId)
	com, err := c.addcomHandler.Handle(command)

	if err != nil {
		c.BaseHandler.Respond(ctx, http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.BaseHandler.Respond(ctx, http.StatusOK, com)

}

func (c *CommentController) DeleteComment(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))

	if err != nil {
		c.BaseHandler.Respond(ctx, http.StatusBadRequest, gin.H{"error": "Id is invalid format"})

	}

	_, err = c.deletecomHandler.Handle(id)

	if err != nil {
		c.BaseHandler.Respond(ctx, http.StatusNotFound, gin.H{"error": "comment not found"})
		return
	}

	c.BaseHandler.Respond(ctx, http.StatusOK, gin.H{})

}
