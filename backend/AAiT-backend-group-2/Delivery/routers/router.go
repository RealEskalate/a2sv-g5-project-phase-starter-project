package routers





import (
	domain "AAiT-backend-group-2/Domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(db *mongo.Database, gin *gin.Engine, configs *domain.Config) {
	publicRouter := gin.Group("/api/")

	NewUserRouter(db, publicRouter, configs)
	NewAuthRouter(db, publicRouter, configs)
	NewBlogRouter(db, publicRouter, configs)
}