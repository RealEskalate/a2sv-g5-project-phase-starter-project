package infrastructure

import (
	"astu-backend-g1/domain"
	"os"
	"time"

	"github.com/golang-jwt/jwt"

	"golang.org/x/crypto/bcrypt"
)

// Genratetoken generates a JWT token for the given user and password.
func Genratetoken(user *domain.User, pwd string) (string, error) {
	var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

	// User login logic
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd)) != nil {
		return "Invalid username or password", nil
	}



	expirationTime := time.Now().Add(24 * 7 * time.Hour)
	claims := &domain.Claims{
		ID:      user.ID,
		Email:   user.Email,
		IsAdmin: user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtToken, err := token.SignedString(jwtSecret)
	return jwtToken, err
}
