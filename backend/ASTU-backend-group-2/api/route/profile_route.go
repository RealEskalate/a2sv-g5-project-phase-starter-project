package route

import (
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/api/controller"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/api/middleware"
	"github.com/gin-gonic/gin"
)

// ProfileHandlers is a function that defines all the routes for the profile
func ProfileHandlers(r *gin.RouterGroup, ctrl controller.ProfileController) {

	// only authenticated users can access
	r.Use(middleware.UserMiddleware())
	r.GET("/profiles/:id", ctrl.GetProfile())
	r.PUT("/profiles/:id", ctrl.UpdateProfile())
	r.PATCH("/profiles/:id", ctrl.UpdateProfile())
	r.DELETE("/profiles/:id", ctrl.DeleteProfile())

	// promote/demote user to admin
	r.Use(middleware.AdminMiddleware())
	r.POST("/profiles/:id/promote", ctrl.PromoteUser())
	r.POST("/profiles/:id/demote", ctrl.DemoteUser())

}
