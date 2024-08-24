package tests

import (
	jwt_service "blog_api/infrastructure/jwt"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/suite"
)

const KEY = "secret"

type JWTServiceTestSuite struct {
	suite.Suite
	jwtService       *jwt_service.JWTService
	jwtService_empty *jwt_service.JWTService
}

func (suite *JWTServiceTestSuite) SetupSuite() {
	suite.jwtService = jwt_service.NewJWTService(KEY)
	suite.jwtService_empty = jwt_service.NewJWTService("")
}

func (suite *JWTServiceTestSuite) TestSignWithJWTPayload_Positive() {
	username := "suser"
	role := "admin"

	token, err := suite.jwtService.SignJWTWithPayload(username, role, "accessToken", time.Minute)

	suite.NoError(err, "no error when given valid inputs")
	suite.True(strings.HasPrefix(token, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9."))
}

func (suite *JWTServiceTestSuite) TestSignWithJWTPayload_Negative_InvalidRole() {
	username := "suser"
	role := "admin"

	token, err := suite.jwtService.SignJWTWithPayload(username, role, "invalid", time.Minute)

	suite.Error(err, "error when given valid inputs")
	suite.Equal(token, "")
}

func (suite *JWTServiceTestSuite) TestSignWithJWTPayload_Negative_NoKey() {
	username := "suser"
	role := "admin"

	token, err := suite.jwtService_empty.SignJWTWithPayload(username, role, "invalid", time.Minute)

	suite.Error(err, "error when given valid inputs")
	suite.Equal(token, "")
}

func (suite *JWTServiceTestSuite) TestValidateAndParseToken_Positive() {
	username := "suser"
	role := "admin"

	token, _ := suite.jwtService.SignJWTWithPayload(username, role, "accessToken", time.Minute)
	jwtToken, err := suite.jwtService.ValidateAndParseToken(token)

	suite.NoError(err, "no error when given the same secret key")
	suite.Equal(username, jwtToken.Claims.(jwt.MapClaims)["username"])
	suite.Equal(role, jwtToken.Claims.(jwt.MapClaims)["role"])
}

func (suite *JWTServiceTestSuite) TestValidateAndParseToken_MalformedToken() {
	username := "suser"
	role := "admin"

	token, _ := suite.jwtService.SignJWTWithPayload(username, role, "accessToken", time.Minute)
	_, err := suite.jwtService.ValidateAndParseToken(token + "a")

	suite.Error(err, "error when given a malformed token")
}

func (suite *JWTServiceTestSuite) TestValidateAndParseToken_Different() {
	username := "suser"
	role := "admin"

	token, _ := suite.jwtService.SignJWTWithPayload(username, role, "accessToken", time.Minute)
	_, err := suite.jwtService_empty.ValidateAndParseToken(token)

	suite.Error(err, "error when given different keys")
}

func (suite *JWTServiceTestSuite) TestGetExpiryDate_Positive() {
	t := time.Now().Round(0).Add(time.Minute).Format(time.RFC3339)
	token := &jwt.Token{
		Claims: jwt.MapClaims{
			"expiresAt": t,
		},
	}

	expiryDate, err := suite.jwtService.GetExpiryDate(token)

	suite.NoError(err, "no error when given a valid token")
	suite.Equal(t, expiryDate.Format(time.RFC3339))
}

func (suite *JWTServiceTestSuite) TestGetExpiryDate_Negative_ClaimNotFound() {
	token := &jwt.Token{
		Claims: jwt.MapClaims{},
	}

	_, err := suite.jwtService.GetExpiryDate(token)

	suite.Error(err, "no error when given a valid token")
}

func (suite *JWTServiceTestSuite) TestGetUsername_Positive() {
	content := "username"
	token := &jwt.Token{
		Claims: jwt.MapClaims{
			"username": content,
		},
	}

	foundUsername, err := suite.jwtService.GetUsername(token)

	suite.NoError(err, "no error when given a valid token")
	suite.Equal(content, foundUsername)
}

func (suite *JWTServiceTestSuite) TestGetUsername_Negative_ClaimNotFound() {
	token := &jwt.Token{
		Claims: jwt.MapClaims{},
	}

	_, err := suite.jwtService.GetUsername(token)

	suite.Error(err, "no error when given a valid token")
}

func (suite *JWTServiceTestSuite) TestGetTokenType_Positive() {
	content := "tokenType"
	token := &jwt.Token{
		Claims: jwt.MapClaims{
			"tokenType": content,
		},
	}

	foundTokentype, err := suite.jwtService.GetTokenType(token)

	suite.NoError(err, "no error when given a valid token")
	suite.Equal(content, foundTokentype)
}

func (suite *JWTServiceTestSuite) TestGetTokentype_Negative_ClaimNotFound() {
	token := &jwt.Token{
		Claims: jwt.MapClaims{},
	}

	_, err := suite.jwtService.GetTokenType(token)

	suite.Error(err, "no error when given a valid token")
}

func (suite *JWTServiceTestSuite) TestGetRole_Positive() {
	content := "role"
	token := &jwt.Token{
		Claims: jwt.MapClaims{
			"role": content,
		},
	}

	foundRole, err := suite.jwtService.GetRole(token)

	suite.NoError(err, "no error when given a valid token")
	suite.Equal(content, foundRole)
}

func (suite *JWTServiceTestSuite) TestGetRole_Negative_ClaimNotFound() {
	token := &jwt.Token{
		Claims: jwt.MapClaims{},
	}

	_, err := suite.jwtService.GetRole(token)

	suite.Error(err, "no error when given a valid token")
}

func TestJWTService(t *testing.T) {
	suite.Run(t, new(JWTServiceTestSuite))
}
