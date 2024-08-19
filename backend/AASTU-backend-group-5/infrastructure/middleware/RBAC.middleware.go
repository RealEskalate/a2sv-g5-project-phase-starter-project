package middleware

import (
	"net/http"
	"strings"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RoleBasedAuth(protected bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		//get token with claims
		auth := c.GetHeader("Authorization")

		var claims = domain.UserClaims{}
		authSplit := strings.Split(auth, " ")
		_, err := jwt.ParseWithClaims(authSplit[1], &claims, func(token *jwt.Token) (interface{}, error) {
			return "", nil
		})

		if err != nil {
			c.IndentedJSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		user := domain.UserClaims{
			ID: claims.ID, 
			Name: claims.Name, 
			Email: claims.Email, 
			IsAdmin: claims.IsAdmin}

		if claims.IsAdmin {
			c.Set("filter", bson.M{})
		} else {
			if protected {
				c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "you need to be an admin"})
				c.Abort()
				return
			}

			path := c.Request.URL.Path
			idx := c.Param("id")
			objid, _ := primitive.ObjectIDFromHex(idx)
        	if strings.Contains(path, "user") && idx != "" && objid != user.ID {
				c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
				c.Abort()
				return
			}
			c.Set("filter", bson.M{"user._id": claims.ID})
		}

		c.Set("user", user)
		c.Next()
	}

}