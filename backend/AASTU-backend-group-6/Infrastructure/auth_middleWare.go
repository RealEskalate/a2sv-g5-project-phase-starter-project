package infrastructure

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Auth struct{
	env Config
}
type AuthInterface interface{
	AuthenticationMiddleware() gin.HandlerFunc
}

func NewAuthMiddleware (env Config)*Auth{
	return &Auth{
		env : env,
	}
}
func (authenticate *Auth) AuthenticationMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")
		if authHeader == ""{
			c.JSON(http.StatusUnauthorized, gin.H{
				"message" : "Authorization header required",
			})
			c.Abort()
			return
		}
		auth := strings.Split(authHeader, " ")
		if len(auth) != 2 || strings.ToLower(auth[0]) != "bearer"{
			c.JSON(http.StatusUnauthorized, gin.H{
				"message" : "Invalid authorization header.",
			})
			c.Abort()
			return
		}
		claims, err := ExtractFromToken(auth[1], authenticate.env.AccessTokenSecret)
		if err != nil{
			c.JSON(http.StatusUnauthorized, gin.H{
				"message" : err,
			})
			c.Abort()
			return
		}
		c.Set("user_name", claims.Name)
		c.Set("user_id" , claims.ID)
		c.Set("role", claims.Role)
		fmt.Println(claims.Name, claims.ID, claims.Role, "dani")
		c.Next()
		// Can check the expiration time of the token if it is valid or not
	}
}
