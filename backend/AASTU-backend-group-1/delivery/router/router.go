package router

import (
	"blogs/delivery/controller"
	"blogs/repository"
	"blogs/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func getBlogController(database *mongo.Database) *controller.BlogController {
	blogRepository := repository.NewBlogRepository(database)
	blogUsecase := usecase.NewBlogUsecase(blogRepository)
	blogController := controller.NewBlogController(blogUsecase)
	return blogController
}

func getUserController(database *mongo.Database) *controller.UserController {
	userRepository := repository.NewUserRepository(database)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	return userController
}

func publicRouter(router *gin.Engine, userController *controller.UserController) {
	router.POST("/users/register", func(ctx *gin.Context) {})
	router.POST("/users/login", func(ctx *gin.Context) {})
	router.POST("/users/reset-password", func(ctx *gin.Context) {})
	router.POST("/users/forgot-password", func(ctx *gin.Context) {})
}

func privateUserRouter(router *gin.Engine, userController *controller.UserController) {
	router.POST("/users/logout", func(ctx *gin.Context) {})
	router.POST("/tokens/refresh", func(ctx *gin.Context) {})

	router.POST("/users/:id/blogs", func(ctx *gin.Context) {})
	router.GET("/users/:id/blogs", func(ctx *gin.Context) {})
}

func privateBlogRouter(router *gin.Engine, blogController *controller.BlogController) {
	router.POST("/blogs", blogController.InsertBlog)
	router.GET("/blogs", blogController.GetBlogs)
	router.GET("/blogs/:id", func(ctx *gin.Context) {})
	router.PUT("/blogs/:id", func(ctx *gin.Context) {})
	router.DELETE("/blogs/:id", func(ctx *gin.Context) {})

	router.POST("/blogs/:id/comments", func(ctx *gin.Context) {})
	router.GET("/blogs/:id/comments", func(ctx *gin.Context) {})

	router.POST("/blogs/:id/likes", func(ctx *gin.Context) {})
	router.GET("/blogs/:id/likes", func(ctx *gin.Context) {})
}

func SetupRouter(mongoClient *mongo.Client) *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default())

	database := mongoClient.Database("blog")
	blogController := getBlogController(database)
	userController := getUserController(database)

	publicRouter(router, userController)

	// router.Use(middleware.AuthMiddleware(mongoClient))

	privateUserRouter(router, userController)
	privateBlogRouter(router, blogController)

	return router
}
