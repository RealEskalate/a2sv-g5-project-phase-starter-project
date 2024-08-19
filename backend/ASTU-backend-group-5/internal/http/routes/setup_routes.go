package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUpRoute(router *gin.Engine, blogsCollection, commentsCollection, likesCollection, viewsCollection, tagsCollection, userCollection *mongo.Collection) {
	RegisterUserRoutes(userCollection, router)
	RegisterVerificationRoutes(userCollection, router)
	RegisterAdminUserRoutes(userCollection, router)
	RegisterBlogRoutes(blogsCollection, commentsCollection, likesCollection, viewsCollection, tagsCollection, router)
	
}
