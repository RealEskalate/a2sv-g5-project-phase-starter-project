package routes

import "github.com/gin-gonic/gin"

func SetUp(router *gin.Engine){
	aiRoute := router.Group("")
	NewAiRequestRoute(aiRoute)
}