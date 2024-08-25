package middleware

import (
	"errors"
	"net/http"

	custom_error "github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/errors"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/validators"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/go-playground/validator/v10"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err == nil {
			return
		}

		status := custom_error.MapErrorToStatusCode(err)
		c.Status(status)
		error_message := custom_error.ErrMessage(err)
		msg := render.JSON{Data: error_message}
		msg.Render(c.Writer)
	}

}

func CustomErrorResponse(c *gin.Context, err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]custom_error.ValidationErrorResponse, len(ve))
		for i, fe := range ve {
			out[i] = custom_error.ValidationErrorResponse{Field: fe.Field(), Message: validators.ValidationMessage(fe)}
		}
		c.JSON(http.StatusBadRequest, custom_error.ErrValidation(out))
		return
	}

	c.JSON(http.StatusBadRequest, custom_error.ErrMessage(err))
}
