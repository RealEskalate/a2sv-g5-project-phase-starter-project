package router

import (
	"astu-backend-g1/delivery/controllers"

	"github.com/gin-gonic/gin"
)

type UserRoute struct {
	handler controllers.UserController
	//middelware gin.HandlerFunc
}

func NewUserRoute(router *gin.Engine, handler controllers.UserController) *UserRoute{
	return &UserRoute{
		handler: handler,
	}
}

func (r *UserRoute) UserRoutes() *gin.RouterGroup{
	ro:=gin.Default()
	userrouter := ro.Group("/user")
	userrouter.POST("/register", r.handler.Register)
	userrouter.GET("/verify?email=:email&pwd=:pwd", r.handler.AccountVerification)
	userrouter.POST("/login", r.handler.LoginUser)
	userrouter.GET("/forgetPassword", r.handler.ForgetPassword)
	userrouter.POST("/resetPassword?email=:email&pwd=:pwd", r.handler.ResetPassword)
	return userrouter
}