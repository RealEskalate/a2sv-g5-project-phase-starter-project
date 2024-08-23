package gin

import (
	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/auth"
	blogDomain "github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/blog"
	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/chat"
	"github.com/gin-gonic/gin"
)

func SetUpChatRouter(r *gin.Engine, chatHandler *ChatHandler, chatUsecase *chat.ChatUsecase) {
	r.GET("/", AuthMiddleware(), chatHandler.GetChatsHandler)
	r.POST("/", AuthMiddleware(), chatHandler.CreateChatHandler)
	r.POST("/", AuthMiddleware(), chatHandler.GenerateChatTitleHandler)
	r.GET("/:id", AuthMiddleware(), chatHandler.GetChatHandler)
	r.DELETE("/:id", AuthMiddleware(), chatHandler.DeleteChatHandler)
	r.POST("/:id/send-message", AuthMiddleware(), chatHandler.SendMessageHandler)
}

func SetUpBlogRouter(r *gin.Engine, blogController *BlogController, blogUseCase *blogDomain.BlogUseCase) {
	r.POST("/", AuthMiddleware(), blogController.CreateBlog)
	r.GET("/", AuthMiddleware(), blogController.GetBlogs)
	r.GET("/:id", AuthMiddleware(), blogController.GetBlogByID)
	r.PUT("/:id", AuthMiddleware(), blogController.UpdateBlog)
	r.DELETE("/:id", AuthMiddleware(), blogController.DeleteBlog)
	r.GET("/search", AuthMiddleware(), blogController.SearchBlogs)
	r.GET("/:id/comments", AuthMiddleware(), blogController.GetCommentsByBlogID)
	r.POST("/:id/comments", AuthMiddleware(), blogController.CreateComment)
	r.DELETE("/:id/comments/:comment_id", AuthMiddleware(), blogController.DeleteComment)
	r.POST("/:id/like", AuthMiddleware(), blogController.LikeBlog)
	r.POST("/:id/dislike", AuthMiddleware(), blogController.DislikeBlog)
	r.DELETE("/:id/unlike", AuthMiddleware(), blogController.UnLikeBlog)
	r.DELETE("/:id/undislike", AuthMiddleware(), blogController.UnDislikeBlog)
}

func SetUpAuthRouter(r *gin.Engine, userController *AuthController, authUsecase *auth.AuthServices) {
	r.POST("/login", AuthMiddleware(), userController.Login)
	r.POST("/register", AuthMiddleware(), userController.RegisterUser)
	r.PUT("/profile", AuthMiddleware(), userController.UpdateProfile)
	r.POST("/activate/:userID/:token", AuthMiddleware(), userController.ActivateUser)
	r.POST("/logout", AuthMiddleware(), userController.Logout)
}
