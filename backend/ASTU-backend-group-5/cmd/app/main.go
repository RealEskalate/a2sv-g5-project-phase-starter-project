package main

import (
	"blogApp/internal/http/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.RegisterVerificationRoutes(router)

	router.Run(fmt.Sprintf(":%d", 8080))
}
