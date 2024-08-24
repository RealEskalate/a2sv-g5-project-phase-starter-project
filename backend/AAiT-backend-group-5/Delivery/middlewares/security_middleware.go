package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func NewSecureMiddleware() gin.HandlerFunc {
	secureMiddleware := secure.New(secure.Options{
		FrameDeny:          true,
		ContentTypeNosniff: true,
		BrowserXssFilter:   true,
		SSLRedirect:        true,
		SSLProxyHeaders:    map[string]string{"X-Forwarded-Proto": "https"},
		IsDevelopment:      true,
		ReferrerPolicy:     "strict-origin-when-cross-origin",
	})

	secureFunc := func() gin.HandlerFunc {
		return func(c *gin.Context) {
			err := secureMiddleware.Process(c.Writer, c.Request)

			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}

			if status := c.Writer.Status(); status > 300 && status < 399 {
				c.Abort()
			}
		}
	}()

	return secureFunc
}
