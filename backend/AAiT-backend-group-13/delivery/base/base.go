package basecontroller

import (
	// "errors"
	"net/http"

	"github.com/gin-gonic/gin"
	// errapi "github.com/group13/blog/domain/errors"
)

// Controller Responses
type BaseHandler struct{}


func (h *BaseHandler) RespondWithCookies(c *gin.Context, status int, v interface{}, cookies []*http.Cookie) {
	for _, cookie := range cookies {
		http.SetCookie(c.Writer, cookie)
	}
	h.Respond(c, status, v)
}

// RespondWithLocation writes a response with a Location header to the Gin context.
func (h *BaseHandler) RespondWithLocation(c *gin.Context, status int, v interface{}, resourceLocation string) {
	c.Header("Location", resourceLocation)
	h.Respond(c, status, v)
}

// Respond writes a JSON response to the Gin context.
func (h *BaseHandler) Respond(c *gin.Context, status int, v interface{}) {
	if v == nil {
		c.Status(status)
	} else {
		c.JSON(status, v)
	}
}