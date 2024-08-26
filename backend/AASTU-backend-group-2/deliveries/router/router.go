package router

import (
	"blog_g2/deliveries/controllers"
	"blog_g2/infrastructure/middleware"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetRouter(router *gin.Engine, com *controllers.CommentController, c *controllers.BlogController, cu *controllers.UserController, oc *controllers.OAuthController, client *mongo.Client, lc *controllers.LikeController, dsl *controllers.DisLikeController) {

	router.POST("/user/register", cu.RegisterUser)
	router.POST("/user/verify-email", cu.VerifyEmail)
	router.POST("/user/login", cu.LoginUser)
	router.GET("/logout", middleware.AuthMiddleware(client), cu.LogoutUser)

	router.POST("/forget-password", cu.ForgotPassword)
	router.POST("/reset-password", cu.ResetPassword)

	// Google OAuth Routes
	router.GET("/auth/google", oc.HandleGoogleLogin)
	router.GET("/oauth2/callback", oc.HandleGoogleCallback)

	router.POST("/generate", middleware.AuthMiddleware(client), c.GeneratePost)
	router.POST("/file", middleware.AuthMiddleware(client), c.FileUpload)

	r := router.Group("/blog")
	r.Use(middleware.AuthMiddleware(client))
	{
		r.POST("/", c.CreateBlog)
		r.GET("/", c.RetrieveBlog)
		r.PUT("/:id", c.UpdateBlog)
		r.DELETE("/:id", c.DeleteBlog)
		r.GET("/search", c.SearchBlog)
		r.GET("/filter", c.FilterBlog)

		r.POST("/comment/:blog_id", com.CreateComment)
		r.GET("/comment/:blog_id", com.GetComment)
		r.PUT("/comment/:blog_id/:id", com.UpdateComment)
		r.DELETE("/comment/:blog_id/:id", com.DeleteComment)

		r.POST("/like/:postID", lc.CreateLike)
		r.DELETE("/like/:postID", lc.DeleteLike)
		r.GET("/like/:postID", lc.GetLikes)

		r.POST("/dislike/:postID", dsl.CreateDisLike)
		r.DELETE("/dislike/:postID", dsl.DeleteDisLike)
		r.GET("/dislike/:postID", dsl.GetDisLikes)

	}

}
