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
func (c *Controller) RegisterPublic(route *gin.RouterGroup) {

}

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
		c.Problem(ctx, errapi.NewBadRequest(err.Error()))
		return
	}

	cmd := blogcmd.NewAddCommand(request.Title, request.Content, request.Tags, request.UserId)
	blog, err := c.addHandler.Handle(cmd)
	if err != nil {
		c.Problem(ctx, errapi.FromErrDMN(err.(*er.Error)))
		return
	}

	baseURL := fmt.Sprintf("http://%s", ctx.Request.Host)
	resourceLocation := fmt.Sprintf("%s%s/%s", baseURL, ctx.Request.URL.Path, blog.ID().String())
	c.RespondWithLocation(ctx, http.StatusCreated, nil, resourceLocation)
}

func (c *Controller) updateBlog(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		c.Problem(ctx, errapi.NewBadRequest(err.Error()))
		return
	}

	var request BlogUpdateDto
	if err := ctx.ShouldBindJSON(&request); err != nil {
		c.Problem(ctx, errapi.NewBadRequest(err.Error()))
		return
	}

	cmd := blogcmd.NewUpdateCommand(id, request.Title, request.Content, request.Tags)
	_, err = c.updateHandler.Handle(cmd)
	if err != nil {
		c.Problem(ctx, errapi.FromErrDMN(err.(*er.Error)))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *Controller) deleteBlog(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		c.Problem(ctx, errapi.NewBadRequest(err.Error()))
		return
	}

	_, err = c.deleteHandler.Handle(id)
	if err != nil {
		c.Problem(ctx, errapi.FromErrDMN(err.(*er.Error)))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *Controller) getBlogs(ctx *gin.Context) {
	cursor, limit, authorId := extractBlogQueryParams(ctx)

	log.Printf("Fetching blogs for query cursor = %s, limit = %d, authorId = %v -- BlogController", cursor, limit, authorId)
	lastSeenIdStr, err := decodeCursor(cursor)
	if err != nil {
		log.Printf("Error decoding cursor %v", err)
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

	var nextCursor string
	if len(blogs) > 0 {
		nextCursor = buildCursor(blogs[len(blogs)-1])
	}
	response := []*BlogResponse{}
	for _, blog := range blogs {
		response = append(response, FromBlog(blog))
	}
	ctx.JSON(http.StatusOK, gin.H{"blogs": response, "cursor": nextCursor})
}

func (c *Controller) getBlogById(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		c.Problem(ctx, errapi.NewBadRequest(err.Error()))
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
	authorId := parseUUID(ctx.Query("authorId"))

	return cursor, limit, authorId
}

func decodeCursor(cursor string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return "", fmt.Errorf("failed to decode cursor: %w", err)
	}
	return string(decodedBytes), nil
}

func buildCursor(lastBlog *models.Blog) string {
	if lastBlog == nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString([]byte(lastBlog.ID().String()))
}

// parseIntOrDefault parses a string as an integer, returning a default value on failure.
func parseIntOrDefault(value string, defaultValue int) int {
	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return parsedValue
}

// parseUUID parses a string as a UUID, returning uuid.Nil on failure.
func parseUUID(value string) uuid.UUID {
	parsedValue, err := uuid.Parse(value)
	if err != nil {
		return uuid.Nil
	}
	return parsedValue
}
