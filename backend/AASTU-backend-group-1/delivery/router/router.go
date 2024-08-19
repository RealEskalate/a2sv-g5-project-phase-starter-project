package router

import (
	"blogs/delivery/controller"
	"blogs/delivery/middleware"
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

	err := userUsecase.AddRoot()
	if err != nil {
		panic(err)
	}

	return userController
}

func publicRouter(router *gin.Engine, userController *controller.UserController) {
	router.POST("/users/register", userController.RegisterUser)
	router.POST("/users/login", userController.LoginUser)
	router.POST("/users/forgot-password", userController.ForgotPassword)
	router.GET("/users/verify", userController.VerifyUser)
	router.GET("/users/reset-password", userController.ResetPassword)
}

func protectedRouter(router *gin.Engine, userController *controller.UserController) {
	router.GET(
		"/tokens/refresh",
		middleware.AuthMiddleware("refresh"),
		userController.RefreshToken,
	)
}

func privateUserRouter(router *gin.RouterGroup, userController *controller.UserController) {
	router.PATCH("/users", userController.UpdateProfile)
	router.PATCH("/users/promote", userController.PromoteUser)
	router.POST("/users/logout", userController.LogoutUser)
}

func privateBlogRouter(router *gin.RouterGroup, blogController *controller.BlogController) {
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
	protectedRouter(router, userController)

	privateRouter := router.Group("")
	privateRouter.Use(middleware.AuthMiddleware("access"))

	privateUserRouter(privateRouter, userController)
	privateBlogRouter(privateRouter, blogController)

	return router
}
