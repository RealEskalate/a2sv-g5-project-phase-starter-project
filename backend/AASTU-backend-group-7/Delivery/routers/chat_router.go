package routers

import (
	"blogapp/Delivery/controllers"
	"blogapp/Infrastructure/auth_middleware"
)

func ChatRouter() {
	chatRouter := Router.Group("/chat", auth_middleware.AuthMiddleware())
	{

		chatController := controllers.NewChatController()

		chatRouter.POST("", chatController.GetChatCompletion)
		chatRouter.POST("/tags", chatController.GetChatCompletionByTags)
		chatRouter.POST("/enhance", chatController.GetChatCompletionEnhancements)
		chatRouter.POST("/generate", chatController.HandleGeneratePostRequest)

	}
}
