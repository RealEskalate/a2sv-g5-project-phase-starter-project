// Delivery/routers/router.go
package routers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    // Public Routes

    SetUpAuth(router)
    SetPublicRoutes(router)
    SetUpUser(router)
    SetUpAdmin(router)
    SetUpAi(router)
   


  
    return router
}

