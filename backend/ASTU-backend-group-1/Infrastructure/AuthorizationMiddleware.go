package infrastructure

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthMiddleware struct {
	collection *mongo.Collection
}

func NewAuthMiddleware(collection *mongo.Collection) GeneralAuthorizer {
	return &AuthMiddleware{
		collection: collection,
	}
}

func (am *AuthMiddleware) AUTH(tokenString, secretKey string) jwt.Claims {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, http.ErrAbortHandler
		}
		return []byte(secretKey), nil
	})
	if err != nil || !token.Valid {
		return nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && claims.Valid() == nil {
		if claims["exp"].(int64) < time.Now().Unix() {
			// todo: check if  there is un expired refresh token in the token collection
			var refreshToken jwt.Token
			err := am.collection.FindOne(context.TODO(), bson.M{"_id": claims["ID"]}).Decode(&refreshToken)
			if err != nil {
				return nil
			}
			refreshClaim, ok := refreshToken.Claims.(jwt.MapClaims)
			if ok {
				if refreshClaim["exp"].(int64) < time.Now().Unix() {

					am.collection.DeleteOne(context.TODO(), bson.M{"_id": claims["ID"]})
					return nil
				}
				claims["exp"] = time.Now().Add(1 * time.Minute).Unix()
				return claims
			}
		}
		return claims
	}
	return nil
}

func (am *AuthMiddleware) AdminAuth(anyClaim any) bool {
	claims := anyClaim.(jwt.MapClaims)
	StringIsAdmin := strings.ToLower(claims["IsAdmin"].(string))
	return StringIsAdmin == "true"
}

//	func (am *AuthMiddleware) OwnerAuth(anyClaim any) bool {
//		claims := anyClaim.(jwt.MapClaims)
//		isOwner := claims["IsOwner"].(string)
//		return isOwner == "true"
//	}
func (am *AuthMiddleware) UserAuth(anyClaim any) bool {
	claims := anyClaim.(jwt.MapClaims)
	IsAdmin := strings.ToLower(claims["IsAdmin"].(string))
	return IsAdmin == "false"
}
