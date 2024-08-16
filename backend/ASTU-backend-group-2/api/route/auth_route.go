package route

import (
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/api/controller"
	"github.com/gin-gonic/gin"
)

func AuthHandlers(r *gin.RouterGroup, ctrl controller.AuthHandlers) {

	r.POST("login", ctrl.Login())
	r.POST("logout", ctrl.Logout())
	r.POST("signup", ctrl.Signup())
	r.POST("forgot-password", ctrl.ForgotPassword())

	r.GET("verify-email", ctrl.VerifyEmail())
}
