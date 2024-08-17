package controllers

import (
	"blogapp/Domain"
	"errors"

	"github.com/gin-gonic/gin"
)

func Getclaim(c *gin.Context) (*Domain.AccessClaims, error) {
	claim, exists := c.Get("claim")
	if !exists {
		return nil, errors.New("claim not set")
	}

	userClaims, ok := claim.(*Domain.AccessClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return userClaims, nil
}
