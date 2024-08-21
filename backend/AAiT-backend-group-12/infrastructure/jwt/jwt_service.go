package jwt_service

import (
	"blog_api/domain"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

// JWTService provides an interface to interact with JWT tokens
type JWTService struct {
	secret string
}

// NewJWTService creates a new JWTService with the given secret
func NewJWTService(secret string) *JWTService {
	return &JWTService{secret: secret}
}

/*
Creates and signs a JWT with the username, role and tokenLifeSpan as the
payloads. Returns the signed token if there aren't any errors.
*/
func (s *JWTService) SignJWTWithPayload(username string, role string, tokenType string, tokenLifeSpan time.Duration) (string, domain.CodedError) {
	if s.secret == "" {
		return "", domain.NewError("internal server error", domain.ERR_INTERNAL_SERVER)
	}

	if tokenType != "accessToken" && tokenType != "refreshToken" {
		return "", domain.NewError("Invalid token type field", domain.ERR_INTERNAL_SERVER)
	}

	jwtSecret := []byte(s.secret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  username,
		"role":      role,
		"expiresAt": time.Now().Round(0).Add(tokenLifeSpan),
		"tokenType": tokenType,
	})
	jwtToken, signingErr := token.SignedString(jwtSecret)
	if signingErr != nil {
		return "", domain.NewError("internal server error: "+signingErr.Error(), domain.ERR_INTERNAL_SERVER)
	}

	return jwtToken, nil
}

/*
Parses the JWT token with the HMAC signing method and returns a pointer
to a jwt.Token struct if the token is valid and not tampered with.
*/
func (s *JWTService) ValidateAndParseToken(rawToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(rawToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(s.secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("error: " + err.Error())
	}

	if !token.Valid {
		return nil, fmt.Errorf("error: Invalid token,  Potentially malformed")
	}

	return token, nil
}

// GetExpiryDate returns the expiry date of the token
func (s *JWTService) GetExpiryDate(token *jwt.Token) (time.Time, domain.CodedError) {
	expiresAt, ok := token.Claims.(jwt.MapClaims)["expiresAt"]
	if !ok {
		return time.Now(), domain.NewError("Invalid token: Expiry date not found", domain.ERR_UNAUTHORIZED)
	}

	expiresAtTime, convErr := time.Parse(time.RFC3339Nano, fmt.Sprintf("%v", expiresAt))
	if convErr != nil {
		return time.Now(), domain.NewError("Error while parsing expiry date: "+convErr.Error(), domain.ERR_UNAUTHORIZED)
	}

	return expiresAtTime, nil
}

// GetUsername returns the username of the token
func (s *JWTService) GetUsername(token *jwt.Token) (string, domain.CodedError) {
	username, ok := token.Claims.(jwt.MapClaims)["username"]
	if !ok {
		return "", domain.NewError("Invalid token: Username not found", domain.ERR_UNAUTHORIZED)
	}

	return fmt.Sprintf("%v", username), nil
}

// GetRole returns the role of the token
func (s *JWTService) GetRole(token *jwt.Token) (string, domain.CodedError) {
	role, ok := token.Claims.(jwt.MapClaims)["role"]
	if !ok {
		return "", domain.NewError("Invalid token: Role not found", domain.ERR_UNAUTHORIZED)
	}

	return fmt.Sprintf("%v", role), nil
}

// GetTokenType returns the token type of the token
func (s *JWTService) GetTokenType(token *jwt.Token) (string, domain.CodedError) {
	tokenType, ok := token.Claims.(jwt.MapClaims)["tokenType"]
	if !ok {
		return "", domain.NewError("Invalid token: TokenType not found", domain.ERR_UNAUTHORIZED)
	}

	return fmt.Sprintf("%v", tokenType), nil
}
