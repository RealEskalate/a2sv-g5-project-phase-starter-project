package route

import (
	"mime"
	"path/filepath"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/gin-gonic/gin"
)

func NewPublicFileRouter(env *bootstrap.Env, group *gin.RouterGroup) {
	group.Use(func(c *gin.Context) {
		if c.Request.Method == "GET" {
			ext := filepath.Ext(c.Request.URL.Path)
			mimeType := mime.TypeByExtension(ext)
			if mimeType != "" {
				c.Header("Content-Type", mimeType)
			}
		}
		c.Next()
	})
	group.Static("/images", "../static/images")
}
