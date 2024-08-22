package usercontroller

import (
	"log"
	"net/http"
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
)

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
}
func (u UserController) RegisterPublic(router *gin.RouterGroup) {
	router = router.Group("/auth")
	router.POST("/signup", u.signUp)
	router.POST("/login", u.login)
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

	u.RespondWithCookies(ctx, http.StatusOK, nil, []*http.Cookie{
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
