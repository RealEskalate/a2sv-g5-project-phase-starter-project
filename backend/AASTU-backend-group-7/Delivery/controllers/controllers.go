package controllers

import (
	"blogapp/Domain"
	"errors"
	"strings"

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

// genreate slug from title
func GenerateSlug(title string) string {
	slug := title
	slug = strings.ToLower(slug)
	slug = strings.Replace(slug, " ", "-", -1)
	return slug
}
