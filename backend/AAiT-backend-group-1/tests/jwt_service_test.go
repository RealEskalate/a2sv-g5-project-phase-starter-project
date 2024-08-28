package tests

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/infrastructure"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JWTTestSuite struct {
	suite.Suite
	jwtService    domain.JwtService
	accessSecret  string
	refreshSecret string
	verfiySecret  string
	resetSecret   string
}

func (suite *JWTTestSuite) SetupTest() {
	suite.accessSecret = "access_secret"
	suite.refreshSecret = "refresh_secret"
	suite.verfiySecret = "verfiy_secret"
	suite.resetSecret = "reset_secret"
	suite.jwtService = infrastructure.NewJWTTokenService(suite.accessSecret, suite.refreshSecret, suite.verfiySecret, suite.resetSecret)
}

func (suite *JWTTestSuite) TestGenerateAccessTokenWithPayload_Success() {
	user := domain.User{
		Username: "testUser",
		Password: "password",
		Role:     "user",
		ID:       primitive.NewObjectID(),
		Email:    "testuser@blog.com",
	}
	_, errGen := suite.jwtService.GenerateAccessTokenWithPayload(user, time.Minute)
	suite.NoError(errGen, "Error generating refresh token")
}

func (suite *JWTTestSuite) TestValidateAccessToken_Success() {
	user := domain.User{
		Username: "testUser",
		Password: "password",
		Role:     "user",
		ID:       primitive.NewObjectID(),
		Email:    "testuser@blog.com",
	}
	token, errGen := suite.jwtService.GenerateAccessTokenWithPayload(user, time.Minute)
	suite.NoError(errGen, "Error generating refresh token")

	_, errVer := suite.jwtService.ValidateAccessToken(token)
	suite.NoError(errVer, "Error should be nil")

}

func (suite *JWTTestSuite) TestGenerateAccessTokenWithPayload_Fail() {
	user := domain.User{
		Username: "testUser",
		Password: "password",
		ID:       primitive.NewObjectID(),
		Email:    "testuser@blog.com",
	}
	_, errGen := suite.jwtService.GenerateAccessTokenWithPayload(user, time.Minute)
	suite.Error(errGen, "Error generating refresh token, because user doesn't have ID")
	suite.Equal(400, errGen.StatusCode(), "Error code should be 400")
}

func (suite *JWTTestSuite) TestValidateAccessToken_Fail() {
	token := "invalid_token"
	_, errVer := suite.jwtService.ValidateAccessToken(token)
	suite.Error(errVer, "Error should not be nil")
	suite.Equal(500, errVer.StatusCode(), "Error code should be 500")
}

func (suite *JWTTestSuite) TestValidateAccessToken_TimeExp() {
	user := domain.User{
		Username: "testUser",
		Password: "password",
		Role:     "user",
		ID:       primitive.NewObjectID(),
		Email:    "testuser@blog.com",
	}
	token, errGen := suite.jwtService.GenerateAccessTokenWithPayload(user, time.Second)
	fmt.Println(token)
	suite.NoError(errGen, "Error generating refresh token")
	time.Sleep(2 * time.Second)
	_, errVer := suite.jwtService.ValidateAccessToken(token)
	suite.Error(errVer, "Error should not be nil")
	suite.Equal(500, errVer.StatusCode(), "Error code should be 500")
}

func (suite *JWTTestSuite) TestGenerateRefreshTokenWithPayload_Success() {
	user := domain.User{
		ID: primitive.NewObjectID(),
	}
	_, errGen := suite.jwtService.GenerateRefreshTokenWithPayload(user, time.Minute)
	suite.NoError(errGen, "Error generating refresh token")
}

func (suite *JWTTestSuite) TestValidateRefreshToken_Success() {
	user := domain.User{
		ID: primitive.NewObjectID(),
	}
	token, errGen := suite.jwtService.GenerateRefreshTokenWithPayload(user, time.Minute)
	suite.NoError(errGen, "Error generating refresh token")

	_, errVer := suite.jwtService.ValidateRefreshToken(token)
	suite.NoError(errVer, "Error should be nil")

}

func (suite *JWTTestSuite) TestGenerateRefreshTokenWithPayload_Fail() {
	user := domain.User{}
	_, errGen := suite.jwtService.GenerateRefreshTokenWithPayload(user, time.Minute)
	suite.Error(errGen, "Error generating refresh token, because user doesn't have ID")
	suite.Equal(400, errGen.StatusCode(), "Error code should be 400")
}

func (suite *JWTTestSuite) TestValidateRefreshToken_Fail() {
	token := "invalid_token"
	_, errVer := suite.jwtService.ValidateRefreshToken(token)
	suite.Error(errVer, "Error should not be nil")
	suite.Equal(500, errVer.StatusCode(), "Error code should be 500")
}

func (suite *JWTTestSuite) TestValidateRefreshToken_TimeExp() {
	user := domain.User{
		Username: "testUser",
		Password: "password",
		Role:     "user",
		ID:       primitive.NewObjectID(),
		Email:    "testuser@blog.com",
	}
	token, errGen := suite.jwtService.GenerateRefreshTokenWithPayload(user, time.Second)
	fmt.Println(token)
	suite.NoError(errGen, "Error generating refresh token")
	time.Sleep(2 * time.Second)
	_, errVer := suite.jwtService.ValidateRefreshToken(token)
	suite.Error(errVer, "Error should not be nil")
	suite.Equal(500, errVer.StatusCode(), "Error code should be 500")
}

func (suite *JWTTestSuite) TestGenerateVerificationToken_Success() {
	user := domain.User{
		Username: "testUser",
		Password: "password",
		Role:     "user",
		Email:    "testuser@blog.com",
	}
	_, errGen := suite.jwtService.GenerateVerificationToken(user, time.Minute)
	suite.NoError(errGen, "Error should be nil")
}

func (suite *JWTTestSuite) TestValidateVerificationToken_Success() {
	user := domain.User{
		Username: "testUser",
		Password: "password",
		Role:     "user",
		Email:    "user@blog.com",
	}
	token, errGen := suite.jwtService.GenerateVerificationToken(user, time.Minute)
	suite.NoError(errGen, "Error generating refresh token")

	_, errVer := suite.jwtService.ValidateVerificationToken(token)
	suite.NoError(errVer, "Error should be nil")
}

func (suite *JWTTestSuite) TestGenerateVerificationToken_Fail() {
	user := domain.User{
		Username: "testUser",
		Password: "password",
		Email:    "testuser@blog.com",
	}
	_, errGen := suite.jwtService.GenerateVerificationToken(user, time.Minute)
	suite.Error(errGen, "Error should not be nil because user doesn't have role")
	suite.Equal(400, errGen.StatusCode(), "Error code should be 400")

}

func (suite *JWTTestSuite) TestValidateVerificationToken_Fail() {
	token := "invalid_token"
	_, errVer := suite.jwtService.ValidateVerificationToken(token)
	suite.Error(errVer, "Error should not be nil")
	suite.Equal(http.StatusBadRequest, errVer.StatusCode(), "Error code should be 500")
}

func (suite *JWTTestSuite) TestValidateVerificationToken_TimeExp() {
	user := domain.User{
		Username: "testUser",
		Password: "password",
		Role:     "user",
		Email:    "user@blog.com",
	}
	token, errGen := suite.jwtService.GenerateVerificationToken(user, time.Second)
	fmt.Println(token)
	suite.NoError(errGen, "Error generating refresh token")
	time.Sleep(2 * time.Second)
	_, errVer := suite.jwtService.ValidateVerificationToken(token)
	suite.Error(errVer, "Error should not be nil")
	suite.Equal(http.StatusBadRequest, errVer.StatusCode(), "Error code should be 500")
}

func (suite *JWTTestSuite) TestGenerateResetPasswordToken_Success() {
	email := "user@blog.com"
	_, errGen := suite.jwtService.GenerateResetToken(email, time.Minute)
	suite.NoError(errGen, "Error should be nil")
}

func (suite *JWTTestSuite) TestValidateResetPasswordToken_Success() {
	email := "user@blog.com"
	token, errGen := suite.jwtService.GenerateResetToken(email, time.Minute)
	suite.NoError(errGen, "Error generating refresh token")

	_, errVer := suite.jwtService.ValidateResetToken(token)
	suite.NoError(errVer, "Error should be nil")
}

func (suite *JWTTestSuite) TestGenerateResetPasswordToken_Fail() {
	email := ""
	_, errGen := suite.jwtService.GenerateResetToken(email, time.Minute)
	suite.Error(errGen, "Error should not be nil because email is empty")
	suite.Equal(http.StatusBadRequest, errGen.StatusCode(), "Error code should be 400")
}

func (suite *JWTTestSuite) TestValidateResetPasswordToken_Fail() {
	token := "invalid_token"
	_, errVer := suite.jwtService.ValidateResetToken(token)
	suite.Error(errVer, "Error should not be nil")
	suite.Equal(http.StatusInternalServerError, errVer.StatusCode(), "Error code should be 500")
}

func (suite *JWTTestSuite) TestValidateResetPasswordToken_TimeExp() {
	email := "user@blog.com"
	token, errGen := suite.jwtService.GenerateResetToken(email, time.Second)
	suite.NoError(errGen, "Error generating refresh token")
	time.Sleep(2 * time.Second)
	_, errVer := suite.jwtService.ValidateResetToken(token)
	suite.Error(errVer, "Error should not be nil")
	suite.Equal(http.StatusInternalServerError, errVer.StatusCode(), "Error code should be 500")
}

func TestJWTTestSuite(t *testing.T) {
	suite.Run(t, new(JWTTestSuite))
}
