package router

import (
	"blog_api/domain"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(port int, routePrefix string, db *mongo.Database) {
	router := gin.Default()

	// auth
	authRouter := router.Group("/auth")
	NewAuthRouter(db.Collection(domain.CollectionUsers), authRouter)

	// blog
	blogAuthor := router.Group("/auth")
	NewBlogRouter(db.Collection(domain.CollectionBlogs), blogAuthor)

	router.Run(fmt.Sprintf(":%v", port))
}
