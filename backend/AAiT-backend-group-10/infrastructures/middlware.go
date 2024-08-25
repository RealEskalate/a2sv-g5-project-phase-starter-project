package infrastructures

import (
	"net/http"
	"strings"

	"aait.backend.g10/usecases/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AuthMiddleware(JwtService interfaces.IJwtService) gin.HandlerFunc {
	return func(context *gin.Context) {
		defer context.Next()

		authHeader := context.GetHeader("Authorization")

		if authHeader == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"Error": "Authorization header is required"})
			context.Abort()
			return
		}

		authPart := strings.Split(authHeader, " ")

		if len(authPart) != 2 || strings.ToLower(authPart[0]) != "bearer" {
			context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Authorization header"})
			context.Abort()
			return
		}

		token, err := JwtService.ValidateToken(authPart[1])

		if token == nil || !token.Valid {
			errMsg := "Invalid or expired token"

			context.JSON(http.StatusUnauthorized, gin.H{"error": errMsg})
			context.Abort()
			return
		}
		if err != nil {
			context.JSON(401, gin.H{"error": err.Message})
			context.Abort()
			return
		}

		claims, ok := JwtService.FindClaim(token)
		if !ok {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			context.Abort()
			return
		}
		role := claims["is_admin"]
		id, error := uuid.Parse(claims["id"].(string))
		if error != nil {
			context.JSON(401, gin.H{"error": "Invalid token claims"})
			context.Abort()
			return
		}

		if role == nil || id == uuid.Nil {
			context.JSON(401, gin.H{"error": "Invalid token claims"})
			context.Abort()
			return
		}
		context.Set("is_admin", role)
		context.Set("id", id)
		context.Set("token", authPart[1])
		context.Next()

	}
}

func AdminMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer context.Next()
		is_admin, exists := context.Get("is_admin")
		if !exists || is_admin != true {
			context.JSON(http.StatusForbidden, gin.H{"message": "Sorry, you must be an admin"})
			context.Abort()
			return
		}
	}
}
