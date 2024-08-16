package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


func SetUpRoute(router *gin.Engine, collection *mongo.Collection) {
	RegisterUserRoutes(collection, router)
	RegisterVerificationRoutes(collection, router)
	RegisterAdminUserRoutes(collection, router)
}
