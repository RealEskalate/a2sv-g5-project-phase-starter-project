package basecontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errapi "github.com/group13/blog/delivery/errors"
)

// BaseHandler is a base struct for all HTTP request handlers that provides basic HTTP functionalities.
type BaseHandler struct{}

// Problem handles errors by writing an appropriate response to the Gin context.
// The primary usage is to hid some messages.
func (h *BaseHandler) Problem(c *gin.Context, err errapi.Error) {
	var shadowedErr errapi.Error
	switch err.StatusCode() {
	case errapi.BadRequest, errapi.Conflict, errapi.NotFound, errapi.Forbidden:
		shadowedErr = err
	case errapi.Authentication:
		shadowedErr = errapi.NewAuthentication("invalid credentials")
	default:
		shadowedErr = errapi.NewServerError("something went wrong")
	}
	h.respondError(c, shadowedErr)
}

// RespondError writes an error response to the Gin context.
func (h *BaseHandler) respondError(c *gin.Context, err errapi.Error) {
	c.JSON(err.StatusCode(), gin.H{"error": err.Error()})
}

// RespondWithCookies writes a response with cookies to the Gin context.
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
