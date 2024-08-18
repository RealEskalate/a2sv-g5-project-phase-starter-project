package route

import (
	"backend-starter-project/bootstrap"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


func Setup(env *bootstrap.Env, db *mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")

	NewAuthRouter(env, db, publicRouter)
	
}