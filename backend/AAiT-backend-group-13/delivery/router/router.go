// Package router provides functionality to set up and run the HTTP server,
// manage routes, and apply middleware based on access levels.
//
// It configures and initializes routes with varying access requirements:
// - Public routes: Accessible without authentication.
// - Protected routes: Require authentication.
// - Privileged routes: Require both authentication and admin privileges.
package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/group13/blog/delivery/common"
	authmiddleware "github.com/group13/blog/delivery/middleware/auth"
	ijwt "github.com/group13/blog/usecase/common/i_jwt"
)

// Router manages the HTTP server and its dependencies,
// including controllers and JWT authentication.
type Router struct {
	addr        string
	baseURL     string
	controllers []common.IController
	jwtService  ijwt.Service
}

// Config holds configuration settings for creating a new Router instance.
type Config struct {
	Addr        string               // Address to listen on
	BaseURL     string               // Base URL for API routes
	Controllers []common.IController // List of controllers
	JwtService  ijwt.Service         // JWT service
}

// NewRouter creates a new Router instance with the given configuration.
// It initializes the router with address, base URL, controllers, and JWT service.
func NewRouter(config Config) *Router {
	return &Router{
		addr:        config.Addr,
		baseURL:     config.BaseURL,
		controllers: config.Controllers,
		jwtService:  config.JwtService,
	}
}

// Run starts the HTTP server and sets up routes with different access levels.
//
// Routes are grouped and managed under the base URL, with the following access levels:
// - Public routes: No authentication required.
// - Protected routes: Authentication required.
// - Privileged routes: Authentication and admin privileges required.
func (r *Router) Run() error {
	router := gin.Default()

	// Setting up routes under baseURL
	api := router.Group(r.baseURL)
	{
		// Public routes (accessible without authentication)
		publicRoutes := api.Group("/v1")
		{
			for _, c := range r.controllers {
				c.RegisterPublic(publicRoutes)
			}
		}

		// Protected routes (authentication required)
		protectedRoutes := api.Group("/v1")
		protectedRoutes.Use(authmiddleware.Authoriz(r.jwtService, false))
		{
			for _, c := range r.controllers {
				c.RegisterProtected(protectedRoutes)
			}
		}

		// Privileged routes (authentication and admin privileges required)
		privilegedRoutes := api.Group("/v1")
		privilegedRoutes.Use(authmiddleware.Authoriz(r.jwtService, true))
		{
			for _, c := range r.controllers {
				c.RegisterPrivileged(privilegedRoutes)
			}
		}
	}

	log.Println("Listening on", r.addr)
	return router.Run(r.addr)
}
