package router

import (
	"blog_api/domain"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(port int, routePrefix string, db *mongo.Database, redisClient *redis.Client) {
	router := gin.Default()

	// auth
	authRouter := router.Group("/api/" + routePrefix + "/auth")
	NewAuthRouter(db.Collection(domain.CollectionUsers), authRouter, redisClient)
	oauthRouter := router.Group("/")
	NewOAuthRouter(oauthRouter)

	// blog
	blogAuthor := router.Group("/api/" + routePrefix + "/blogs")
	NewBlogRouter(db.Collection(domain.CollectionBlogs), blogAuthor)

	router.Run(fmt.Sprintf(":%v", port))
}
