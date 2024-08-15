package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUpRoute(router *gin.Engine, collection *mongo.Collection) {
	RegisterUserRouters(collection, router)
	RegisterVerificationRoutes(collection, router)
}
