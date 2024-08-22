package controllers_test

import (
	"blogapp/Domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	user_id = primitive.NewObjectID()
)

func mockExtractUser(c *gin.Context) (*Domain.AccessClaims, error) {
	return &Domain.AccessClaims{ID: user_id, Role: "admin"}, nil
}
