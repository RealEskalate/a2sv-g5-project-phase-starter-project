package routers

import (
	"astu-backend-g1/delivery/controllers"
	"astu-backend-g1/infrastructure"

	_ "astu-backend-g1/delivery/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type MainRouter struct {
	blogController controllers.BlogController
	authController infrastructure.GeneralAuthorizationController
	aiController   controllers.AIController
	handler        controllers.UserController
}

func NewMainRouter(uc controllers.UserController, bc controllers.BlogController, authc infrastructure.GeneralAuthorizationController, ac controllers.AIController) *MainRouter {
	return &MainRouter{
		blogController: bc,
		authController: authc,
		handler:        uc,
		aiController:   ac,
	}
}

func (gr *MainRouter) GinBlogRouter() {

	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	blogRouter := gr.addBlogRouter(router)
	_ = gr.addUserRouter(router)
	// conf := config.Config{}
	// prompts := infrastructure.Prompts{}
	// _ = gr.AddAIRoutes(router, conf, prompts)
	commentRouter := gr.addCommentRouter(blogRouter)
	_ = gr.addReplyRouter(commentRouter)

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Welcome to Blog API get"})
		ctx.Abort()
	})
	router.POST("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Welcome to Blog API create"})
		ctx.Abort()
	})
	router.DELETE("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Welcome to Blog API delete"})
		ctx.Abort()
	})
	router.PATCH("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Welcome to Blog API patch"})
		ctx.Abort()
	})
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"message": "Not such route"})
		ctx.Abort()
	})
	router.Run(":8000")
}
