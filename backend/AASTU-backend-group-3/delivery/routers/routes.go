// Delivery/routers/router.go
package routers

import (
    "group3-blogApi/infrastracture"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    rateLimiter := infrastracture.NewRateLimitterMiddleware()

    // Add rate limiter middleware
    router.Use(rateLimiter.RateLimitter())

    // Add CORS middleware
    router.Use(cors.Default())

    // auth routes
    SetUpAuth(router)

    // public routes
    SetUpPublicRoutes(router)

    // user routes
    SetUpUser(router)

    // Admin routes
    SetUpAdmin(router)

    // Blog routes
    SetUpBlog(router)

    // Like and Dislike routes
    SetUpLike(router)

    // Comment routes
    SetUpComment(router)

    // Ai routes
    SetUpAi(router)

    return router
}