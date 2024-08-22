package utils

import (
	"AAiT-backend-group-6/domain"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)


type UserClaim struct{
	User_id			string			
	Username		string
	Email			string
	User_type		string
	jwt.StandardClaims
}

func ValidateToken(signedToken string, secret_key string) (claims *UserClaim, err error){
	token, msg := jwt.ParseWithClaims(
		signedToken, 
		&UserClaim{}, 
		func(t *jwt.Token) (interface{}, error) {
			return []byte(secret_key), nil
		},
	)

	if msg != nil || !token.Valid{
		err = msg
		return
	}

	claims, ok:= token.Claims.(*UserClaim)
	if !ok{
		err = errors.New("the token is invalid")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix(){
		err = errors.New("token is expired")
		return
	}
	return claims, err
}

func GenerateAccessToken(user *domain.User, expiry int, secret_key string) (signedAccessToken string, err error){
	claims := &UserClaim{
		User_id: user.ID.Hex(),
		Username: user.Username,
		Email: user.Email,
		User_type: user.User_type,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(expiry)).Unix(),
		},
	}

	signedAccessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret_key))
    if err != nil {
        return "", err
    }

    return
}

func GenerateRefreshToken(user *domain.User, expiry int, secret_key string) (signedRefreshToken string, err error){
	refreshClaims := &UserClaim{
		User_id: user.ID.Hex(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(expiry)).Unix(),
		},
	}

    signedRefreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secret_key))
    if err != nil {
        return "", err
    }

    return
}