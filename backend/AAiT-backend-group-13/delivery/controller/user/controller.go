package usercontroller

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	basecontroller "github.com/group13/blog/delivery/controller/base"
	errapi "github.com/group13/blog/delivery/errors"
	er "github.com/group13/blog/domain/errors"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	passwordreset "github.com/group13/blog/usecase/password_reset"
	usercmd "github.com/group13/blog/usecase/user/command"
	userqry "github.com/group13/blog/usecase/user/query"
	"github.com/group13/blog/usecase/user/result"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig *oauth2.Config

	oauthStateString = generateStateString(16)
)

func init() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/api/v1/auth/callback",
		ClientID:     os.Getenv("GOOGLE_ID"),
		ClientSecret: os.Getenv("GOOGLE_SECRET"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	fmt.Println("config", googleOauthConfig)
}

// UserController handles user-related HTTP requests.
type UserController struct {
	basecontroller.BaseHandler
	promoteHandler       icmd.IHandler[*usercmd.PromoteCommand, bool]
	loginHandler         icmd.IHandler[*userqry.LoginQuery, *result.LoginInResult]
	signupHandler        icmd.IHandler[*usercmd.SignUpCommand, *result.SignUpResult]
	resetPasswordHandler icmd.IHandler[*passwordreset.ResetCommand, bool]
	resetCodeSendHandler icmd.IHandler[string, time.Time]
	validateCodeHandler  icmd.IHandler[*passwordreset.ValidateCodeCommand, string]
	validateEmailHandler icmd.IHandler[string, *result.ValidateEmailResult]
	updateProfileHandler icmd.IHandler[*usercmd.UpdateProfileCommand, *result.UpdateProfileResult]
	googleSignin         icmd.IHandler[usercmd.GoogleSigninCommand, *result.LoginInResult]
	googleSignup         icmd.IHandler[usercmd.GoogleSignupCommand, bool]
}

// Config holds the configuration for creating a new UserController.
type Config struct {
	PromoteHandler       icmd.IHandler[*usercmd.PromoteCommand, bool]
	LoginHandler         icmd.IHandler[*userqry.LoginQuery, *result.LoginInResult]
	SignupHandler        icmd.IHandler[*usercmd.SignUpCommand, *result.SignUpResult]
	ResetPasswordHandler icmd.IHandler[*passwordreset.ResetCommand, bool]
	ResetCodeSendHandler icmd.IHandler[string, time.Time]
	ValidateCodeHandler  icmd.IHandler[*passwordreset.ValidateCodeCommand, string]
	ValidateEmailHandler icmd.IHandler[string, *result.ValidateEmailResult]
	UpdateProfileHandler icmd.IHandler[*usercmd.UpdateProfileCommand, *result.UpdateProfileResult]
	GoogleSignin         icmd.IHandler[usercmd.GoogleSigninCommand, *result.LoginInResult]
	GoogleSignup         icmd.IHandler[usercmd.GoogleSignupCommand, bool]
}

// New creates a new UserController with the given CQRS handlers.
func New(config Config) *UserController {
	return &UserController{
		promoteHandler:       config.PromoteHandler,
		loginHandler:         config.LoginHandler,
		signupHandler:        config.SignupHandler,
		resetPasswordHandler: config.ResetPasswordHandler,
		resetCodeSendHandler: config.ResetCodeSendHandler,
		validateCodeHandler:  config.ValidateCodeHandler,
		validateEmailHandler: config.ValidateEmailHandler,
		updateProfileHandler: config.UpdateProfileHandler,
		googleSignin:         config.GoogleSignin,
		googleSignup:         config.GoogleSignup,
	}
}

func (u UserController) RegisterPrivileged(router *gin.RouterGroup) {
	router = router.Group("/users")
	router.POST("/:username/promote", u.promte)
	router.POST("/:username/demote", u.demote)
}

func (u UserController) RegisterProtected(router *gin.RouterGroup) {
	router = router.Group("/auth")
	router.POST("/:username/logout", u.logout)
	router.PUT("/update/:id", u.updateProfile)
}
func (u UserController) RegisterPublic(router *gin.RouterGroup) {
	router = router.Group("/auth")
	router.POST("/signup", u.signUp)
	router.POST("/login", u.login)
	router.GET("/login", u.handleGoogleLogin)
	router.GET("/signup", u.handleGoogleSignup)
	router.GET("/callback", u.handleGoogleCallback)
	router.POST("/resetPasswordCode", u.forgotPassword)
	router.POST("/validateCode", u.validateCode)
	router.POST("/resetPassword", u.resetPassword)
	router.POST("/validateEmail", u.validateEmail)
}

func (u *UserController) signUp(ctx *gin.Context) {
	var request SignUpDto
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		return
	}

	log.Printf("Started creating new account for user with username %s -- UserController", request.Username)
	command := usercmd.NewSignUpCommand(request.Username, request.FirstName, request.LastName, request.Email, request.Password)
	_, err := u.signupHandler.Handle(command)
	if err != nil {
		log.Println("User signed up failed -- UserController")
		u.Problem(ctx, errapi.FromErrDMN(err.(*er.Error)))
		return
	}

	log.Println("User signed up successfully -- UserController")
	u.Respond(ctx, http.StatusCreated, "Signed Up successfully")
}

func (u *UserController) login(ctx *gin.Context) {
	var request LoginDto
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		log.Println("User input could not be bound -- UserController")
		return
	}
	log.Printf("login in user %s -- UserController", request.Username)

	command := userqry.NewLoginQuery(request.Username, request.Password)
	res, err := u.loginHandler.Handle(command)
	if err != nil {
		u.Problem(ctx, errapi.FromErrDMN(err.(*er.Error)))
		log.Println("User use case invalidated data -- UserController")
		return
	}

	responseBody := NewLoginResponse(res)
	u.RespondWithCookies(ctx, http.StatusOK, responseBody, []*http.Cookie{
		{
			Name:     "accessToken",
			Value:    res.Token,
			Path:     "/",
			Domain:   ctx.Request.Host,
			MaxAge:   24 * 60 * 60,
			HttpOnly: true,
			Secure:   true,
		},
		{
			Name:     "refreshToken",
			Value:    res.RefreshToken,
			Path:     "/refreshToken",
			Domain:   ctx.Request.Host,
			MaxAge:   48 * 60 * 60,
			HttpOnly: true,
			Secure:   true,
		},
	})
}

func (u *UserController) resetPassword(ctx *gin.Context) {
	var request ResetPasswordDto
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		log.Println("User input could not be bound -- UserController")
		return
	}

	command := passwordreset.NewResetCommand(request.Token, request.NewPassword)
	_, err := u.resetPasswordHandler.Handle(command)
	if err != nil {
		u.Problem(ctx, errapi.FromErrDMN(err.(*er.Error)))
		log.Println("Password reset use case failed -- UserController")
		return
	}

	log.Println("Password reset successfully -- UserController")
	ctx.JSON(http.StatusOK, "Password Reset successful")
}

func (u *UserController) validateCode(ctx *gin.Context) {
	var request ValidateCodeDto
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		log.Println("User input could not be bound -- UserController")
		return
	}

	command := passwordreset.NewValidateCodeCommand(request.Code, request.Email)
	token, err := u.validateCodeHandler.Handle(command)
	if err != nil {
		u.Problem(ctx, errapi.FromErrDMN(err.(*er.Error)))
		log.Println("Code validation use case failed -- UserController")
		return
	}

	log.Println("Code validated successfully -- UserController")
	u.Respond(ctx, http.StatusOK, gin.H{"resetToken": token})
}

func (u *UserController) forgotPassword(ctx *gin.Context) {
	var request ForgotPasswordDto
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		log.Println("User input could not be bound -- UserController")
		return
	}

	_, err := u.resetCodeSendHandler.Handle(request.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		log.Println("Forgot password use case failed -- UserController")
		return
	}

	log.Println("Forgot password request processed -- UserController")
	ctx.JSON(http.StatusOK, "Password reset instructions sent")
}

func (u *UserController) logout(ctx *gin.Context) {
	u.RespondWithCookies(ctx, http.StatusNoContent, nil, []*http.Cookie{
		{
			Name:     "accessToken",
			Value:    "",
			Path:     "/",
			Domain:   ctx.Request.Host,
			MaxAge:   -1, // Delete the cookie
			HttpOnly: true,
			Secure:   true,
		},
		{
			Name:     "refreshToken",
			Value:    "",
			Path:     "/",
			Domain:   ctx.Request.Host,
			MaxAge:   -1, // Delete the cookie
			HttpOnly: true,
			Secure:   true,
		},
	})
}

func (u UserController) promte(ctx *gin.Context) {
	u.changeStatus(true, ctx)
}

func (u UserController) demote(ctx *gin.Context) {
	u.changeStatus(false, ctx)
}

func (u UserController) changeStatus(toAdmin bool, ctx *gin.Context) {
	username := ctx.Param("id")
	claims, exists := ctx.Get("userClaims")
	if !exists {
		u.Problem(ctx, errapi.NewAuthentication("authentication required"))
		return
	}

	// Type assertion to jwt.MapClaims
	jwtClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		u.Problem(ctx, errapi.NewAuthentication("authentication required"))
		return
	}

	// Extract and parse the user_id claim as a UUID
	userIDStr, ok := jwtClaims["user_id"].(string)
	if !ok {
		u.Problem(ctx, errapi.NewAuthentication("authentication required"))
		return
	}

	promoterId, err := uuid.Parse(userIDStr)
	if err != nil {
		u.Problem(ctx, errapi.NewAuthentication("authentication required"))
		return
	}

	_, err = u.promoteHandler.Handle(usercmd.NewPromoteCommand(username, toAdmin, promoterId))
	if err != nil {
		u.Problem(ctx, errapi.FromErrDMN(err.(*er.Error)))
		return
	}

	log.Println("User logged in successfully -- controller")
	u.BaseHandler.Respond(ctx, http.StatusNoContent, gin.H{"message": "User status changed successfully"})

}

func (u UserController) validateEmail(ctx *gin.Context) {

	encryptedValue := ctx.Query("secret")

	// Pass the secret to the use case
	_, err := u.validateEmailHandler.Handle(encryptedValue)
	if err != nil {
		u.Problem(ctx, errapi.FromErrDMN(err.(*er.Error)))
		log.Println("User use case invalidated it -- user controller")
		return
	}
	log.Println("Email validated successfully -- controller")
	u.BaseHandler.Respond(ctx, http.StatusOK, gin.H{"message": "Email validated successfully"})
}

func (u UserController) updateProfile(ctx *gin.Context) {
	var request UpdateProfileDto
	userid := ctx.Param("id")
	if userid == "" {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		log.Println("User input could not be bound -- UserController")
		return
	}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		log.Println("User input could not be bound -- UserController")
		return
	}

	command := usercmd.NewUpdateProfileCommand(request.Username, request.FirstName, request.LastName, request.Email, request.Password, userid)
	_, err := u.updateProfileHandler.Handle(command)
	if err != nil {
		u.Problem(ctx, errapi.FromErrDMN(err.(*er.Error)))
		log.Println("User use case invalidated data -- UserController")
		return
	}

	log.Println("User profile updated successfully -- UserController")
	u.Respond(ctx, http.StatusOK, "Profile updated successfully")
}
func (u UserController) handleGoogleLogin(ctx *gin.Context) {
	state := oauthStateString + "_login"
	url := googleOauthConfig.AuthCodeURL(state)
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func (u UserController) handleGoogleSignup(ctx *gin.Context) {
	state := oauthStateString + "_signup"
	url := googleOauthConfig.AuthCodeURL(state)
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func (u UserController) handleGoogleCallback(ctx *gin.Context) {
	state := ctx.Query("state")
	code := ctx.Query("code")

	userInfo, err := getUserInfo(state, code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if state == oauthStateString+"_login" {

		command := usercmd.NewGoogleSigninCommand(userInfo.Email, userInfo.VerifiedEmail)
		res, err := u.googleSignin.Handle(*command)

		if err != nil {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}
		u.RespondWithCookies(ctx, http.StatusOK, res, []*http.Cookie{
			{
				Name:     "accessToken",
				Value:    res.Token,
				Path:     "/",
				Domain:   ctx.Request.Host,
				MaxAge:   24 * 60 * 60,
				HttpOnly: true,
				Secure:   true,
			},
			{
				Name:     "refreshToken",
				Value:    res.RefreshToken,
				Path:     "/",
				Domain:   ctx.Request.Host,
				MaxAge:   48 * 60 * 60,
				HttpOnly: true,
				Secure:   true,
			},
		})
	} else if state == oauthStateString+"_signup" {
		command := usercmd.NewGoogleSignupCommand(userInfo.GivenName, userInfo.FamilyName, userInfo.Email, userInfo.VerifiedEmail)

		registered, err := u.googleSignup.Handle(*command)
		if err != nil {
			ctx.JSON(http.StatusConflict, err.Error())
			return
		}
		if registered {
			u.BaseHandler.Respond(ctx, http.StatusCreated, "Signed Up Successfully")
			return
		}
	} else {
		ctx.JSON(http.StatusBadRequest, "Invalid state")
	}
}

func getUserInfo(state, code string) (*UserInfo, error) {
	if state != oauthStateString+"_login" && state != oauthStateString+"_signup" {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	var userInfo UserInfo
	if err := json.Unmarshal(contents, &userInfo); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user info: %s", err.Error())
	}

	return &userInfo, nil
}

func generateStateString(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		log.Fatalf("Unable to generate random state: %v", err)
	}
	// URL-safe base64 encoding
	return base64.URLEncoding.EncodeToString(b)
}
