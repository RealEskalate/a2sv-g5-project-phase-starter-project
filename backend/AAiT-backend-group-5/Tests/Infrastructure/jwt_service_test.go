package infrastructure_test

import (
	"testing"

	config "github.com/aait.backend.g5.main/backend/Config"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	infrastructure "github.com/aait.backend.g5.main/backend/Infrastructure"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type JwtServiceTestSuite struct {
	suite.Suite
	jwtService interfaces.JwtService
	env        *config.Env
	ctr        *gomock.Controller
}

func (suite *JwtServiceTestSuite) SetupSuite() {
	suite.ctr = gomock.NewController(suite.T())
	suite.env = &config.Env{
		JWT_SECRET: "mysecret",
	}
	suite.jwtService = infrastructure.NewJwtService(suite.env)
}

func (suite *JwtServiceTestSuite) TearDownSuite() {
	suite.ctr.Finish()
}

func (suite *JwtServiceTestSuite) TestCreateAccessToken_Success() {
	user := models.User{
		ID:    "user1",
		Email: "user@example.com",
		Role:  "user",
	}

	token, err := suite.jwtService.CreateAccessToken(user, 60)
	suite.Nil(err)
	suite.NotEmpty(token)
}

func (suite *JwtServiceTestSuite) TestCreateRefreshToken_Success() {
	user := models.User{
		ID: "user1",
	}

	token, err := suite.jwtService.CreateRefreshToken(user, 24)
	suite.Nil(err)
	suite.NotEmpty(token)
}

func (suite *JwtServiceTestSuite) TestValidateToken_Success() {
	user := models.User{
		ID:    "user1",
		Email: "user@example.com",
		Role:  "user",
	}
	tokenStr, _ := suite.jwtService.CreateAccessToken(user, 60)

	claims, err := suite.jwtService.ValidateToken(tokenStr)
	suite.Nil(err)
	suite.NotNil(claims)
	suite.Equal(user.ID, claims.ID)
	suite.Equal(user.Email, claims.Email)
}

func (suite *JwtServiceTestSuite) TestValidateToken_InvalidToken() {
	_, err := suite.jwtService.ValidateToken("invalid_token")
	suite.NotNil(err)
}

func (suite *JwtServiceTestSuite) TestValidateAuthHeader_Success() {
	authHeader := "Bearer token_string"
	parts, err := suite.jwtService.ValidateAuthHeader(authHeader)
	suite.Nil(err)
	suite.Equal([]string{"Bearer", "token_string"}, parts)
}

func (suite *JwtServiceTestSuite) TestValidateAuthHeader_InvalidFormat() {
	_, err := suite.jwtService.ValidateAuthHeader("InvalidFormat")
	suite.NotNil(err)
}

func (suite *JwtServiceTestSuite) TestCreateURLToken_Success() {
	user := models.User{
		ID:       "user1",
		Name:     "John Doe",
		Username: "johndoe",
		Email:    "john.doe@example.com",
		Role:     "user",
	}

	token, err := suite.jwtService.CreateURLToken(user, 60)
	suite.Nil(err)
	suite.NotEmpty(token)
}

func (suite *JwtServiceTestSuite) TestValidateURLToken_Success() {
	user := models.User{
		ID:       "user1",
		Name:     "John Doe",
		Username: "johndoe",
		Email:    "john.doe@example.com",
		Role:     "user",
	}
	tokenStr, _ := suite.jwtService.CreateURLToken(user, 60)

	claims, err := suite.jwtService.ValidateURLToken(tokenStr)
	suite.Nil(err)
	suite.NotNil(claims)
	suite.Equal(user.ID, claims.ID)
	suite.Equal(user.Name, claims.Name)
}

func (suite *JwtServiceTestSuite) TestValidateURLToken_InvalidToken() {
	_, err := suite.jwtService.ValidateURLToken("invalid_token")
	suite.NotNil(err)
}

func TestJwtServiceTestSuite(t *testing.T) {
	suite.Run(t, new(JwtServiceTestSuite))
}
