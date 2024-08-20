// Delivery/routers/router.go
package routers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

  
    // auth routes
    SetUpAuth(router)

    // // oauth routes
    SetupOauthRouter(router)

    //public routes
    SetUpPublicRoutes(router)

    // // user routes
    SetUpUser(router)

    // // Admin routes
    // SetUpAdmin(router)

    // // Ai routes
    // SetUpAi(router)
   


  
    return router
}

