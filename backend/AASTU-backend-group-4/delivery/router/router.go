package router

import (
	"blog-api/delivery/controller/blog_controller"
	"blog-api/delivery/controller/user_controller"
	"blog-api/infrastructure/auth"
	"blog-api/infrastructure/bootstrap"

	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine, bc *blog_controller.BlogController, uc *user_controller.UserController, env *bootstrap.Env) {
	// User routes
	router.POST("/signup", uc.SignUp)
	router.POST("/login", uc.Login)
	router.POST("/refresh", uc.RefreshTokens)
	router.GET("/logout", auth.JwtAuthMiddleware(env.AccessTokenSecret), uc.Logout)
	router.POST("/forgot-password", uc.ForgotPassword)
	router.POST("/reset-password", uc.ResetPassword)
	router.PUT("/updateUser", auth.JwtAuthMiddleware(env.AccessTokenSecret), uc.UpdateUser)
	router.PATCH("/user/promote-demote", auth.JwtAuthMiddleware(env.AccessTokenSecret), uc.PromoteDemote)

	router.POST("/generate", auth.JwtAuthMiddleware(env.AccessTokenSecret), bc.GenerateContent)

	// Blog routes
	r := router.Group("/blog")
	r.Use(auth.JwtAuthMiddleware(env.AccessTokenSecret))
	{
		r.POST("/create", bc.CreateBlog)
		r.GET("/", bc.GetBlogs)
		r.GET("/:id", bc.GetBlogByID)
		r.PUT("/:id", bc.UpdateBlog)
		r.DELETE("/:id", bc.DeleteBlog)
		r.GET("/search", bc.SearchBlogs)
		r.POST("/filters", bc.FilterBlog)

	}
}
