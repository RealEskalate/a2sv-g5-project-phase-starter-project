package routers

import (
	"AAiT-backend-group-6/mongo"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db mongo.Database, gin *gin.Engine) {
	NewBlogRouter(db, gin)

	// NewTaskRouter(db, router)

}
