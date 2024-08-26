package routers

import "github.com/gin-gonic/gin"

func (gr *MainRouter) addUserRouter(generalRouter *gin.Engine)  *gin.RouterGroup{
	userrouter := generalRouter.Group("/users")
	{
		userrouter.POST("/register", gr.handler.Register)
		userrouter.GET("/accountVerification", gr.handler.AccountVerification)
		userrouter.POST("/login", gr.handler.LoginUser)
		userrouter.GET("/forgetPassword", gr.handler.ForgetPassword)
		userrouter.POST("/resetPassword", gr.handler.ResetPassword)
		userrouter.GET("/logout", gr.handler.LogoutUser)
		userrouter.POST("/:uid/refresh", gr.handler.RefreshAccessToken)
		userrouter.GET("/", gr.handler.GetUsers)
		userrouter.GET("/:id", gr.handler.GetUserByID)
		userrouter.Use(gr.authController.AuthenticationMiddleware())
		userrouter.PUT("/changePassword", gr.handler.ChangePassword)
		userrouter.PUT("/changeEmail", gr.handler.UpdateProfiles)
		userrouter.PATCH("/promote/:username", gr.authController.ADMINMiddleware(), gr.handler.Promote)
		userrouter.PATCH("/demote/:username", gr.authController.ADMINMiddleware(), gr.handler.Demote)
		userrouter.PATCH("/promotebyemail/:email", gr.authController.ADMINMiddleware(), gr.handler.PromoteByEmail)
		userrouter.PATCH("/demotebyemail/:email", gr.authController.ADMINMiddleware(), gr.handler.DemoteByEmail)
		userrouter.DELETE("/:id", gr.authController.ADMINMiddleware(), gr.handler.DeleteUser)
	}
	return userrouter
}
