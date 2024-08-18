package controllers

import (
	"blogapp/Domain"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// is user author of post or admin
func IsAuthorOrAdmin(c *gin.Context, authorID primitive.ObjectID) (bool, error) {
	claim, err := Getclaim(c)
	if err != nil {
		return false, err
	}

	if claim.Role == "admin" {
		return true, nil
	}

	if claim.ID == authorID {
		return true, nil
	}

	return false, nil
}


