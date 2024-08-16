package controller

import "github.com/gin-gonic/gin"

// interface for blog controllers
type ProfileController interface {
	GetProfile() gin.HandlerFunc
	UpdateProfile() gin.HandlerFunc
	DeleteProfile() gin.HandlerFunc
	PromoteUser() gin.HandlerFunc
	DemoteUser() gin.HandlerFunc
}
