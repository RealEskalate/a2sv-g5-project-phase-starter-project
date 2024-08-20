package blogcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/group13/blog/delivery/controller/user/dto"
	er "github.com/group13/blog/domain/errors"
	blogmodel "github.com/group13/blog/domain/models/blog"
	addcmd "github.com/group13/blog/usecase/blog/command/add"
	updatecmd "github.com/group13/blog/usecase/blog/command/update"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
)

// Controller handles HTTP requests related to tasks.
type Controller struct {
	addHandler    icmd.IHandler[*addcmd.Command, *blogmodel.Blog]
	updateHandler icmd.IHandler[*updatecmd.Command, *blogmodel.Blog]
	deleteHandler icmd.IHandler[uuid.UUID, bool]
}

type Config struct {
	AddHandler    icmd.IHandler[*addcmd.Command, *blogmodel.Blog]
	UpdateHandler icmd.IHandler[*updatecmd.Command, *blogmodel.Blog]
	DeleteHandler icmd.IHandler[uuid.UUID, bool]
}

// New creates a new TaskController with the given CQRS handlers and task repository.
func New(config Config) *Controller {
	return &Controller{
		addHandler:    config.AddHandler,
		updateHandler: config.UpdateHandler,
		deleteHandler: config.DeleteHandler,
	}
}

// RegisterPublic registers public routes.
func (c *Controller) RegisterPublic(route *gin.RouterGroup) {}

// RegisterPrivileged registers privileged routes.
func (c *Controller) RegisterPrivileged(route *gin.RouterGroup) {
	blogs := route.Group("/blogs")
	{
		blogs.POST("", c.AddBlog)
		blogs.PUT("/:id", c.UpdateBlog)
		blogs.DELETE("/:id", c.DeleteBlog)
	}
}

func (c *Controller) AddBlog(ctx *gin.Context) {
	var request dto.BlogDto

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, er.NewBadRequest(err.Error()))
		return
	}

	cmd := addcmd.NewCommand(request.Title, request.Content, request.Tags, request.UserId)
	blog, err := c.addHandler.Handle(cmd)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusCreated, blog)
}

func (c *Controller) UpdateBlog(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, er.NewBadRequest("Invalid Id Format"))
		return
	}

	var request dto.BlogUpdateDto

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}

	cmd := updatecmd.NewCommand(id, request.Title, request.Content, request.Tags)
	_, err = c.updateHandler.Handle(cmd)

	if err != nil {
		if err == er.BlogNotFound {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "blog not found"})
			return
		} else {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "server error"})

		}
		return
	}

	ctx.IndentedJSON(http.StatusNoContent, gin.H{})
}

func (c *Controller) DeleteBlog(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, er.NewBadRequest("Invalid Id Format"))
		return
	}

	_, err = c.deleteHandler.Handle(id)

	if err != nil {
		if err == er.BlogNotFound {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "blog not found"})
			return
		} else {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		}

		return
	}

	ctx.IndentedJSON(http.StatusNoContent, gin.H{})
}
