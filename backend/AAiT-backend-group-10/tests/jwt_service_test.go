package tests

import (
	"testing"

	"aait.backend.g10/domain"
	"aait.backend.g10/infrastructures"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type JwtServiceSuite struct {
	suite.Suite
	service *infrastructures.JwtService
	user    *domain.User
}

func (s *JwtServiceSuite) SetupTest() {
	s.service = &infrastructures.JwtService{
		JwtSecret: "test_secret",
	}
	s.user = &domain.User{
		ID:      uuid.New(),
		Email:   "test.user@example.com",
		IsAdmin: false,
	}
}

func (s *JwtServiceSuite) TestGenerateToken_Success() {
	accessToken, refreshToken, err := s.service.GenerateToken(s.user)

	assert.Nil(s.T(), err)
	assert.NotEmpty(s.T(), accessToken)
	assert.NotEmpty(s.T(), refreshToken)

	// Validate the generated access token
	parsedToken, err := s.service.ValidateToken(accessToken)
	assert.Nil(s.T(), err)
	assert.True(s.T(), parsedToken.Valid)

	claims, ok := s.service.FindClaim(parsedToken)
	assert.True(s.T(), ok)
	assert.Equal(s.T(), s.user.Email, claims["email"])
	assert.Equal(s.T(), s.user.ID.String(), claims["id"])
	assert.Equal(s.T(), s.user.IsAdmin, claims["is_admin"])
}

func (s *JwtServiceSuite) TestValidateToken_Success() {
	accessToken, _, err := s.service.GenerateToken(s.user)
	assert.Nil(s.T(), err)

	parsedToken, err := s.service.ValidateToken(accessToken)
	assert.Nil(s.T(), err)
	assert.True(s.T(), parsedToken.Valid)
}

func (s *JwtServiceSuite) TestValidateToken_Failure() {
	// Test with an invalid token
	parsedToken, err := s.service.ValidateToken("invalid.token")
	assert.Error(s.T(), err)
	assert.Nil(s.T(), parsedToken)
	assert.Equal(s.T(), domain.ErrTokenParsingFailed, err)
}

func (s *JwtServiceSuite) TestGenerateResetToken_Success() {
	code := int64(123456)
	resetToken, err := s.service.GenerateResetToken(s.user.Email, code)

	assert.Nil(s.T(), err)
	assert.NotEmpty(s.T(), resetToken)

	// Validate the generated reset token
	parsedToken, err := s.service.ValidateToken(resetToken)
	assert.Nil(s.T(), err)
	assert.True(s.T(), parsedToken.Valid)

	claims, ok := s.service.FindClaim(parsedToken)
	assert.True(s.T(), ok)
	assert.Equal(s.T(), s.user.Email, claims["email"])
	assert.Equal(s.T(), float64(code), claims["code"])
}

func (s *JwtServiceSuite) TestCheckToken_Success() {
	accessToken, _, err := s.service.GenerateToken(s.user)
	assert.Nil(s.T(), err)

	token, err := s.service.CheckToken(accessToken)
	assert.Nil(s.T(), err)
	assert.True(s.T(), token.Valid)
}

func (s *JwtServiceSuite) TestCheckToken_Failure() {
	// Test with an invalid token
	token, err := s.service.CheckToken("invalid.token")
	assert.Error(s.T(), err)
	assert.Nil(s.T(), token)
	assert.Equal(s.T(), domain.ErrTokenParsingFailed, err)
}

func (s *JwtServiceSuite) TestFindClaim_Success() {
	accessToken, _, err := s.service.GenerateToken(s.user)
	assert.Nil(s.T(), err)

	token, err := s.service.ValidateToken(accessToken)
	assert.Nil(s.T(), err)

	claims, ok := s.service.FindClaim(token)
	assert.True(s.T(), ok)
	assert.Equal(s.T(), s.user.Email, claims["email"])
	assert.Equal(s.T(), s.user.ID.String(), claims["id"])
}

func (s *JwtServiceSuite) TestFindClaim_Failure() {
	// Test with a token that does not contain MapClaims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{})
	claims, ok := s.service.FindClaim(token)
	assert.False(s.T(), ok)
	assert.Nil(s.T(), claims)
}

func TestJwtServiceSuite(t *testing.T) {
	suite.Run(t, new(JwtServiceSuite))
}
