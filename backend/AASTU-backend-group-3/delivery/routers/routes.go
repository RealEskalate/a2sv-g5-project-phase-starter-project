// Delivery/routers/router.go
package routers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()


    // auth routes
    // SetUpAuth(router)

    // // public routes
    // SetPublicRoutes(router)
    // // user routes
    // SetUpUser(router)
    // // Admin routes
    // SetUpAdmin(router)

    // // oauth routes
    SetupOauthRouter(router)

    //public routes
    SetUpPublicRoutes(router)

    // // user routes
    SetUpUser(router)

    // // Admin routes
    SetUpAdmin(router)



    // Blog routes
    SetUpBlog(router)

    // Like and Dislike routes
    SetUpLike(router)

    // Comment routes
    SetUpComment(router)
   


    // // Ai routes
    SetUpAi(router)
    return router
}

