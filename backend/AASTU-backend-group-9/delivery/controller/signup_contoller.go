package controller

import (
	"blog/config"
	"blog/domain"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Env           *config.Env
}

// Signup creates a new user
func (sc *SignupController) Signup(c *gin.Context) {
	var user domain.AuthSignup
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	returnedUser, _ := sc.SignupUsecase.GetUserByEmail(c, user.Email)
	if returnedUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}
	returnedUser, _ = sc.SignupUsecase.GetUserByUsername(c, user.Username)
	if returnedUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}
	err := sc.SignupUsecase.SendOTP(c, &user, sc.Env.SMTPUsername, sc.Env.SMTPPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OTP sent"})

}

// VerifyOTP verifies the OTP
func (sc *SignupController) VerifyOTP(c *gin.Context) {
	var otp domain.OTPRequest
	if err := c.ShouldBindJSON(&otp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	otpresponse, err := sc.SignupUsecase.VerifyOTP(c, &otp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OTP verified"})
	user := domain.AuthSignup{
		Username: otpresponse.Username,
		Email:    otpresponse.Email,
		Password: otpresponse.Password,
	}
	sc.Register(c, user)
}
func (sc *SignupController) Register(c *gin.Context, user domain.AuthSignup) {
	userID, err := sc.SignupUsecase.RegisterUser(c, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Accesstoken, err := sc.SignupUsecase.CreateAccessToken(&user, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	RefreshToken, err := sc.SignupUsecase.CreateRefreshToken(&user, sc.Env.RefreshTokenSecret, sc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = sc.SignupUsecase.SaveRefreshToken(c, RefreshToken, *userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := domain.SignUpResponse{
		ID:           *userID,
		AcessToken:   Accesstoken,
		RefreshToken: RefreshToken,
	}
	c.JSON(http.StatusOK, gin.H{"data": resp})
}

func (sc *SignupController) GoogleLogin(c *gin.Context) {
	url := config.AppConfig.GoogleLoginConfig.AuthCodeURL("randomstate")
	c.Redirect(http.StatusSeeOther, url)
}

func (sc *SignupController) GoogleCallback(c *gin.Context) {
	state := c.Query("state")
	if state != "randomstate" {
		c.String(http.StatusBadRequest, "States don't Match!!")
		return
	}

	code := c.Query("code")

	googlecon := config.GoogleConfig(sc.Env)

	token, err := googlecon.Exchange(context.Background(), code)
	if err != nil {
		c.String(http.StatusBadRequest, "Code-Token Exchange Failed")
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		c.String(http.StatusBadRequest, "User Data Fetch Failed")
		return
	}

	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.String(http.StatusBadRequest, "JSON Parsing Failed")
		return
	}

	var user struct {
		Email string `json:"email"`
	}

	err = json.Unmarshal(userData, &user)
	if err != nil {
		c.String(http.StatusBadRequest, "User Email Extraction Failed")
		return
	}
	returnedUser, _ := sc.SignupUsecase.GetUserByEmail(c, user.Email)
		if returnedUser != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
			return
	}
	// Use the extracted email here
	email := user.Email
	username := email[:strings.Index(email, "@")]
	regstr := domain.AuthSignup{
		Username: username,
		Email:    email,
	}
	sc.Register(c, regstr)
}
