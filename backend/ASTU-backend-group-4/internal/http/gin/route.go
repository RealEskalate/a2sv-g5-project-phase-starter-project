package gin

import (
	"github.com/gin-gonic/gin"
)

func SetUpRouter(r *gin.Engine, blogController *BlogController, chatHandler *ChatHandler, userController *UserController) {
	v1 := r.Group("/v1")
	blogsV1 := v1.Group("/blogs")
	chatV1 := v1.Group("/chats")
	authV1 := v1.Group("/auth")
	SetUpBlogRouter(blogsV1, blogController)
	SetUpChatRouter(chatV1, chatHandler)
	SetUpAuthRouter(authV1, userController)
}

func SetUpChatRouter(r *gin.RouterGroup, chatHandler *ChatHandler) {
	r.GET("/", AuthMiddleware(), chatHandler.GetChatsHandler)
	r.POST("/", AuthMiddleware(), chatHandler.CreateChatHandler)
	r.GET("/:id", AuthMiddleware(), chatHandler.GetChatHandler)
	r.DELETE("/:id", AuthMiddleware(), chatHandler.DeleteChatHandler)
	r.POST("/:id/send-message", AuthMiddleware(), chatHandler.SendMessageHandler)
}

func SetUpBlogRouter(r *gin.RouterGroup, blogController *BlogController) {
	r.Use(AuthMiddleware())

	r.POST("/", blogController.CreateBlog)
	r.GET("/", blogController.GetBlogs)
	r.GET("/:id", blogController.GetBlogByID)
	r.PUT("/:id", blogController.UpdateBlog)
	r.DELETE("/:id", blogController.DeleteBlog)
	r.GET("/search", blogController.SearchBlogs)
	r.GET("/:id/comments", blogController.GetCommentsByBlogID)
	r.POST("/:id/comments", blogController.CreateComment)
	r.DELETE("/:id/comments/:comment_id", blogController.DeleteComment)
	r.POST("/:id/likes", blogController.LikeBlog)
	r.POST("/:id/dislikes", blogController.DislikeBlog)
	r.DELETE("/:id/likes", blogController.UnLikeBlog)
	r.DELETE("/:id/dislike", blogController.UnDislikeBlog)
}

func SetUpAuthRouter(r *gin.Engine, userController *AuthController, authUsecase *auth.AuthServices) {
	r.POST("/login", AuthMiddleware(), userController.Login)
	r.POST("/register", AuthMiddleware(), userController.RegisterUser)
	r.PUT("/profile", AuthMiddleware(), userController.UpdateProfile)
	r.POST("/activate/:userID/:token", userController.ActivateUser)
	r.POST("/logout", AuthMiddleware(), userController.Logout)
}
