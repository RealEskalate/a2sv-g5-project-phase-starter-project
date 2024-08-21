package router

import (
	"blog_api/domain"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

// SetupRouter sets up the router with the given port, route prefix, database, and redis client
func SetupRouter(port int, routePrefix string, db *mongo.Database, redisClient *redis.Client) {
	router := gin.Default()

	// serve static assets
	router.Static("/assets", "./local")

	// auth
	authRouter := router.Group("/api/" + routePrefix + "/auth")
	NewAuthRouter(db.Collection(domain.CollectionUsers), authRouter, redisClient)

	// oauth
	oauthRouter := router.Group("/")
	NewOAuthRouter(oauthRouter)

	// blog
	blogAuthor := router.Group("/api/" + routePrefix + "/blogs")
	NewBlogRouter(db.Collection(domain.CollectionBlogs), redisClient, blogAuthor)

	router.Run(fmt.Sprintf(":%v", port))
}
