package infrastructure

import (
	"fmt"

	domain "blogs/Domain"

	jwt "github.com/golang-jwt/jwt/v4"
)

func CreateToken(claims jwt.Claims, secret string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

// func CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
// 	exp := time.Now().Add(time.Hour * time.Duration(expiry))
// 	claimsRefresh := &domain.JwtCustomRefreshClaims{
// 		ID: user.ID.Hex(),
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(exp), // Convert expiration time to *jwt.NumericDate
// 		},
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
// 	rt, err := token.SignedString([]byte(secret))
// 	if err != nil {
// 		return "", err
// 	}
// 	return rt, err
// }

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractIDFromToken(requestToken string, secret string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	return claims["id"].(string), nil
}

func ExtractFromToken(requestToken string, secret string) (domain.JwtCustomClaims, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return domain.JwtCustomClaims{}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return domain.JwtCustomClaims{}, fmt.Errorf("invalid token")
	}
	return domain.JwtCustomClaims{
		ID: claims["id"].(string),
	}, nil
}
