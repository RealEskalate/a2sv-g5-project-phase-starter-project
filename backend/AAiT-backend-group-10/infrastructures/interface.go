package infrastructures

import (
	"os"

	"aait.backend.g10/domain"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type Infranstructure struct {
	JWTSecret string
}

type InfrastructureInterface interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password string, hash string) bool
	GenerateToken(user *domain.User) (string, string, error)
	ValidateToken(token string) (*jwt.Token, error)
	GenerateResetToken(email string) (string, error)
	SendResetEmail(user *domain.User, resetToken string) error
}

func NewInfrastructure() InfrastructureInterface {
	godotenv.Load()
	jwtSecret := os.Getenv("jwtSecret")
	return &Infranstructure{
		JWTSecret: jwtSecret,
	}
}
