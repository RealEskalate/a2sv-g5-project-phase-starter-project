package blogcontroller

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/group13/blog/delivery/common"
	basecontroller "github.com/group13/blog/delivery/controller/base"
	errapi "github.com/group13/blog/delivery/errors"
	er "github.com/group13/blog/domain/errors"
	"github.com/group13/blog/domain/models"
	blogcmd "github.com/group13/blog/usecase/blog/command"
	blogqry "github.com/group13/blog/usecase/blog/query"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	iqry "github.com/group13/blog/usecase/common/cqrs/query"
)

// Controller handles HTTP requests related to blogs.
type Controller struct {
	basecontroller.BaseHandler
	addHandler         icmd.IHandler[*blogcmd.AddCommand, *models.Blog]
	updateHandler      icmd.IHandler[*blogcmd.UpdateCommand, *models.Blog]
	deleteHandler      icmd.IHandler[uuid.UUID, bool]
	getMultipleHandler icmd.IHandler[*blogqry.GetMultipleQuery, []*models.Blog]
	getHandler         iqry.IHandler[uuid.UUID, *models.Blog]
}

var _ common.IController = &Controller{}

// Config holds the dependencies required by the Controller.
type Config struct {
	AddHandler         icmd.IHandler[*blogcmd.AddCommand, *models.Blog]
	UpdateHandler      icmd.IHandler[*blogcmd.UpdateCommand, *models.Blog]
	DeleteHandler      icmd.IHandler[uuid.UUID, bool]
	GetMultipleHandler icmd.IHandler[*blogqry.GetMultipleQuery, []*models.Blog]
	GetHandler         iqry.IHandler[uuid.UUID, *models.Blog]
}

// New creates a new blog Controller with the given dependencies.
func New(config Config) *Controller {
	return &Controller{
		addHandler:         config.AddHandler,
		updateHandler:      config.UpdateHandler,
		deleteHandler:      config.DeleteHandler,
		getMultipleHandler: config.GetMultipleHandler,
		getHandler:         config.GetHandler,
	}
}

// RegisterPublic registers public routes.
func (c *Controller) RegisterPublic(route *gin.RouterGroup) {}

// RegisterPrivileged registers privileged routes.
func (c *Controller) RegisterPrivileged(route *gin.RouterGroup) {}

// RegisterProtected registers protected routes.
func (c *Controller) RegisterProtected(route *gin.RouterGroup) {
	blogs := route.Group("/blogs")
	{
		blogs.POST("", c.addBlog)
		blogs.PUT("/:id", c.updateBlog)
		blogs.DELETE("/:id", c.deleteBlog)
		blogs.GET("/", c.getBlogs)
		blogs.GET("/:id", c.getBlogById)
	}
}

func (c *Controller) addBlog(ctx *gin.Context) {
	var request BlogDto

	if err := ctx.ShouldBindJSON(&request); err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, er.NewValidation(err.Error()))
		return
	}

	username := uuid.New()
	cmd := blogcmd.NewAddCommand(request.Title, request.Content, request.Tags, username)
	blog, err := c.addHandler.Handle(cmd)
	if err != nil {
		c.respondWithError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.IndentedJSON(http.StatusCreated, blog)
}

func (c *Controller) updateBlog(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, er.NewValidation("Invalid Id Format"))
		return
	}

	var request BlogUpdateDto
	if err := ctx.ShouldBindJSON(&request); err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, er.NewValidation("Invalid Request"))
		return
	}

	cmd := blogcmd.NewUpdateCommand(id, request.Title, request.Content, request.Tags)
	_, err = c.updateHandler.Handle(cmd)
	if err != nil {
		if err == er.BlogNotFound {
			c.respondWithError(ctx, http.StatusNotFound, err)
		} else {
			c.respondWithError(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *Controller) deleteBlog(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, er.NewValidation("Invalid Id Format"))
		return
	}

	_, err = c.deleteHandler.Handle(id)
	if err != nil {
		if err == er.BlogNotFound {
			c.respondWithError(ctx, http.StatusNotFound, err)
		} else {
			c.respondWithError(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *Controller) getBlogs(ctx *gin.Context) {
	cursor, limit, authorId := extractBlogQueryParams(ctx)

	lastSeenIdStr, err := decodeCursor(cursor)
	if err != nil {
		lastSeenIdStr = uuid.Nil.String()
	}

	lastSeenId, err := uuid.Parse(lastSeenIdStr)
	if err != nil {
		lastSeenId = uuid.Nil
	}

	blogs, err := c.getMultipleHandler.Handle(blogqry.NewGetMultipleQuery(authorId, limit, &lastSeenId))
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}

func (c *Controller) getBlogById(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, er.NewValidation("Invalid Id Format"))
		return
	}

	log.Printf("Serving Get blog by id for %v -- Blogcontroller", id)
	blog, err := c.getHandler.Handle(id)
	if err != nil {
		log.Printf("Fetched Blog unsuccessfull %s -- BlogController", err.Error())
		c.Problem(ctx, errapi.FromErrDMN(err.(*er.Error)))
	}
	log.Printf("Fetched Blog for %v successfull -- BlogController", id)

	c.Respond(ctx, http.StatusOK, FromBlog(blog))
}

// extractBlogQueryParams extracts and validates query parameters for blog queries.
func extractBlogQueryParams(ctx *gin.Context) (string, int, uuid.UUID) {
	cursor := ctx.Query("cursor")
	limit := parseIntOrDefault(ctx.Query("limit"), 10)
	authorId := parseUUIDOrNil(ctx.Query("authorId"))

	return cursor, limit, authorId
}

func decodeCursor(cursor string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return "", fmt.Errorf("failed to decode cursor: %w", err)
	}
	return string(decodedBytes), nil
}

func (c *Controller) BuildCursor(lastBlog *models.Blog) string {
	if lastBlog == nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString([]byte(lastBlog.ID().String()))
}

// respondWithError sends an error response with the specified status code.
func (c *Controller) respondWithError(ctx *gin.Context, code int, err error) {
	ctx.IndentedJSON(code, gin.H{"error": err.Error()})
}

// parseIntOrDefault parses a string as an integer, returning a default value on failure.
func parseIntOrDefault(value string, defaultValue int) int {
	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return parsedValue
}

// parseUUIDOrNil parses a string as a UUID, returning uuid.Nil on failure.
func parseUUIDOrNil(value string) uuid.UUID {
	parsedValue, err := uuid.Parse(value)
	if err != nil {
		return uuid.Nil
	}
	return parsedValue
}
