package Router

import (
	controller "AAiT-backend-group-8/Delivery/Controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(controller *controller.Controller) *gin.Engine {
	router := gin.Default()

	router.POST("/register", controller.RegisterUser)
	router.GET("/verify", controller.VerifyEmail)
	router.POST("/login", controller.Login)
	router.POST("/refresh", controller.RefreshToken)
	router.POST("/comment/:blogID", controller.CreateComment)
	router.GET("/comment/:blogID", controller.GetComments)
	router.PATCH("/comment/:commentID", controller.UpdateComment)
	router.DELETE("/comment/:commentID", controller.DeleteComment)	
    router.POST("/forgot-password", controller.ForgotPassword)
	router.GET("/store-token", controller.StoreToken) 
	router.POST("/reset-password", controller.ResetPassword)

	return router
}
