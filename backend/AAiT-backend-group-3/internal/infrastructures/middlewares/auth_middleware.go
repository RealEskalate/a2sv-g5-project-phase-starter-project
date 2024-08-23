package middlewares

import (
	"AAIT-backend-group-3/internal/infrastructures/services"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)


type IAuthMiddleware interface{
	Authentication() gin.HandlerFunc
	RoleAuth(roles ...string) gin.HandlerFunc
}

type AuthMiddleware struct {
	jwtSvc services.IJWT
	cacheSvc services.ICacheService
}

func NewAuthMiddleware(jwtSvc services.IJWT, cacheSvc services.ICacheService) IAuthMiddleware{
	return &AuthMiddleware{
		jwtSvc: jwtSvc,
		cacheSvc: cacheSvc,
	}
}


func (mid *AuthMiddleware) Authentication() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(401, gin.H{"error": "Authorization header is required"})
            c.Abort()
            return
        }

        authSlice := strings.Split(authHeader, " ")
        if len(authSlice) != 2 || strings.ToLower(authSlice[0]) != "bearer" {
            c.JSON(401, gin.H{"error": "Invalid authorization header format"})
            c.Abort()
            return
        }

        tokenString := authSlice[1]
        isBlacklisted, err := mid.cacheSvc.IsTknBlacklisted(tokenString)
        if err != nil {
            c.JSON(500, gin.H{"error": "Internal server error"})
            c.Abort()
            return
        }
        if isBlacklisted {
            c.JSON(401, gin.H{"error": "Token is blacklisted"})
            c.Abort()
            return
        }
        token, err := mid.jwtSvc.ValidateAccessToken(tokenString)
        if err != nil || !token.Valid {
            c.JSON(401, gin.H{"error": err.Error()})
            c.Abort()
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok || !token.Valid {
            c.JSON(401, gin.H{"error": "Invalid JWT claims"})
            c.Abort()
            return
        }
		c.Set("token", tokenString)
        c.Set("claims", claims)
        c.Next()
    }
}


func (mid *AuthMiddleware) RoleAuth(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			c.JSON(401, gin.H{
				"error": "Claims not found in context",
			})
			c.Abort()
			return
		}

		claimsMap, ok := claims.(jwt.MapClaims)
		if !ok {
			c.JSON(401, gin.H{"error": "Invalid JWT claims format"})
			c.Abort()
			return
		}

		role, ok := claimsMap["role"].(string)
		if !ok || role == "" {
			c.JSON(401, gin.H{"error": "Role not found in JWT claims"})
			c.Abort()
			return
		}

		roleAuthorized := false
		for _, elem := range roles {
			if strings.EqualFold(elem, role) {
				roleAuthorized = true
				break
			}
		}

		if !roleAuthorized {
			c.JSON(403, gin.H{"error": "Your role does not have access to this resource"})
			c.Abort()
			return
		}

		c.Set("userId", claimsMap["userId"])
		c.Next()
	}
}


