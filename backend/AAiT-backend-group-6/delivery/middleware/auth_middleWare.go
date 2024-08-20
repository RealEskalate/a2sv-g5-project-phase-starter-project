package middleware

import (
	"AAiT-backend-group-6/utils"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func JwtAuthMiddleware(secret string) gin.HandlerFunc{

    return func(c *gin.Context){
      authHeader := c.GetHeader("Authorization")
      if authHeader == ""{
        c.JSON(401, gin.H{"error": "Authorization header is required"})
        c.Abort()
        return
      }

      authParts := strings.Split(authHeader, " ")
      if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer"{
        c.JSON(401, gin.H{"error": "invalid authorization header"})
        c.Abort()
        return
      }

      claims, err := utils.ValidateToken(authParts[1], secret)

      if err != nil{
        c.JSON(401, gin.H{"error": "Invalid JWT"})
        c.Abort()
        return
      }
      c.Set("email", claims.Email)
	  c.Set("username", claims.Username)
	  c.Set("user_id",claims.User_id)
	  c.Set("user_type", claims.User_type)
	  c.Next()
    }
  }

func AuthRole(reqRole string) gin.HandlerFunc{
  return func(c *gin.Context){
    userType := c.GetString("user_type")
    if userType != reqRole{ 
      fmt.Println("usertype",userType, "reqrole: ",reqRole)
      c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user"})
      c.Abort()
      return
    }
    c.Next()
  } 
}