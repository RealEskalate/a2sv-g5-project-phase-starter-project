package main

import (
	"blogApp/internal/http/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	routes.RegisterVerificationRoutes(route)

	route.Run()
}
