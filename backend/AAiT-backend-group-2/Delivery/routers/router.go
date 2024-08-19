package routers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(db *mongo.Database, gin *gin.Engine, jwtSecret string) {
	publicRouter := gin.Group("/api/")

	NewUserRouter(db, publicRouter, jwtSecret)
}