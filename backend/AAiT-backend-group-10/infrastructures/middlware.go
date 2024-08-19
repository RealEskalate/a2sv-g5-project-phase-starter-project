package infrastructures

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer context.Next()

		authHeader := context.GetHeader("Authorization")

		if authHeader == "" {
			context.JSON(401, gin.H{"Error": "Authoriztion header is required"})
			context.Abort()
			return
		}

		authPart := strings.Split(authHeader, "")

		if len(authPart) != 2 || strings.ToLower(authPart[0]) != "bearer" {
			context.JSON(401, gin.H{"message": "Invalid Authoriztion header"})
			context.Abort()
			return
		}

		token, err := CheckToken(authPart[1])

		if token == nil || !token.Valid {
			errMsg := "Invalid or expired token"

			if err != nil {
				errMsg = err.Error()
			}
			context.JSON(401, gin.H{"error": errMsg})
			context.Abort()
			return
		}

		claims, ok := FindClaim(token)
		if !ok {
			context.JSON(401, gin.H{"error": "Invalid token claims"})
			context.Abort()
			return
		}
		role := claims["is_admin"]
		id := claims["id"]

		context.Set("is_admin", role)
		context.Set("id", id)

	}
}

func AdminMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer context.Next()
		is_admin, exists := context.Get("is_admin")
		if !exists || is_admin != true {
			context.JSON(403, gin.H{"message": "Sorry, you are not eligible for this"})
			context.Abort()
			return
		}
	}
}
