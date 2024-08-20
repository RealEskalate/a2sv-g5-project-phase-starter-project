package controller



import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	basecontroller "github.com/group13/blog/delivery/base"
	dto "github.com/group13/blog/delivery/controller/user/dto"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	logincommand "github.com/group13/blog/usecase/user/command/login"
	signupcommand "github.com/group13/blog/usecase/user/command/signup"
	resetcodevalidate "github.com/group13/blog/usecase/password_reset/code_validator" // Add this line
	promotcmd "github.com/group13/blog/usecase/user/command/promote"
	"github.com/group13/blog/usecase/user/result"
	forgotpassword "github.com/group13/blog/usecase/password_reset/reset"
)

type UserController struct { 
	basecontroller.BaseHandler
	promotHandler icmd.IHandler[*promotcmd.Command, bool]
	loginHandler  icmd.IHandler[*logincommand.LoginCommand, *result.LoginInResult]
	signupHandler icmd.IHandler[*signupcommand.SignUpCommand, *result.SignUpResult]
	resetPasswordhandler icmd.IHandler[*resetcodevalidate.Command, bool]
	forgotPasswordHandler icmd.IHandler[*forgotpassword.Command, bool]
	validateEmailHandler icmd.IHandler[string, *result.ValidateEmailResult]
}

type Config struct {
	basecontroller.BaseHandler
	PromotHandler icmd.IHandler[*promotcmd.Command, bool]
	LoginHandler  icmd.IHandler[*logincommand.LoginCommand, *result.LoginInResult]
	SignupHandler icmd.IHandler[*signupcommand.SignUpCommand, *result.SignUpResult]
	ResetPasswordHandler icmd.IHandler[*resetcodevalidate.Command, bool]
	ForgotPasswordHandler icmd.IHandler[*forgotpassword.Command, bool]
	validateEmailHander icmd.IHandler[string, *result.ValidateEmailResult]
}

// New creates a new UserController with the given CQRS handler.
func New(config Config) *UserController {
	return &UserController{
		BaseHandler:   config.BaseHandler,
		promotHandler: config.PromotHandler,
		loginHandler:  config.LoginHandler,
		signupHandler: config.SignupHandler,
		resetPasswordhandler: config.ResetPasswordHandler,
		forgotPasswordHandler: config.ForgotPasswordHandler,
		validateEmailHandler: config.validateEmailHander,
	}
}

func (u UserController) RegisterPrivileged(router *gin.RouterGroup) {
	router = router.Group("/admin")
	router.POST("/api/v1/auth/promote", u.Promte)
	router.POST("/api/v1/auth/demote", u.Demote)
}

func (u UserController) RegisterPrivate(router *gin.RouterGroup) {
	router = router.Group("/users")
	router.POST("/api/v1/auth/forgot-password", u.ForgotPassword)
	router.POST("/api/v1/auth/reset-password", u.ResetPassword)
	router.POST("POST /api/v1/users/:username/promote", u.Promte)
	router.POST("POST /api/v1/users/:username/demote", u.Demote)
	router.GET("POST /api/v1/users/:username/logout", u.Logout)

}

func (u UserController) RegisterPublic(router *gin.RouterGroup) {
	router = router.Group("/")
	router.POST("/api/v1/auth/signup", u.SignUp)
	router.POST("/api/v1/auth/login", u.Login)

}


func (u UserController) SignUp(ctx *gin.Context) {

	var user dto.SignUpDto
	// bind input files
	if err := ctx.BindJSON(&user); err != nil {

		u.BaseHandler.Respond(ctx, http.StatusBadRequest, gin.H{"message": "Invalid Input"})
		log.Println("User Input cannot be bind -- user controller")
		return
	}

	log.Println("User inputs bind successfully")

	command := signupcommand.NewSignUpCommand(user.Username, user.FirstName, user.LastName, user.Email, user.Password)
	// pass to usercases
	_, err := u.signupHandler.Handle(&command)

	if err != nil {
		u.BaseHandler.Respond(ctx, http.StatusInternalServerError, gin.H{"message": err.Error()})
		log.Println("User Usecase invalidated data -- user controller")
		return

	}

	log.Println("User singed up -- controller")
	u.BaseHandler.Respond(ctx, http.StatusCreated, gin.H{"message": "Signed Up successfully"})

}

func (u UserController) Login(ctx *gin.Context) {
	var user dto.LoginDto
	// bind files to user model
	// bind input files

	if err := ctx.BindJSON(&user); err != nil {
		u.BaseHandler.Respond(ctx, http.StatusBadRequest, gin.H{"message": "Invalid Input"})
		log.Println("User Input cannot be bind -- user controller")
		return
	}

	command := logincommand.NewLoginCommand(user.Username, user.Password)
	// pass to usercases
	res, err := u.loginHandler.Handle(&command)

	// pass to login usercase
	if err != nil {
		u.BaseHandler.Respond(ctx, http.StatusInternalServerError, gin.H{"message": err.Error()})
		log.Println("User use case invalidated it -- user controller")
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
			Value:    res.Refreshtoekn,
			Path:    "/",
			Domain:   ctx.Request.Host,
			MaxAge:   48 * 60 * 60,
			HttpOnly: true,
			Secure:   true,
	
		},
	})
}
	


func (u UserController) ForgotPassword(ctx *gin.Context) {
	var request dto.ForgotPasswordDto
	// bind input files
	if err := ctx.BindJSON(&request); err != nil {
		u.BaseHandler.Respond(ctx, http.StatusBadRequest, gin.H{"message": "Invalid Input"})
		log.Println("User Input cannot be bind -- user controller")
		return
	}

	// pass to usecase
	command := forgotpassword.NewCommand(request.Id, request.Token, request.NewPassword)
	_, err := u.forgotPasswordHandler.Handle(command)

	if err != nil {
		u.BaseHandler.Respond(ctx, http.StatusInternalServerError, gin.H{"message": err.Error()})
		log.Println("User use case invalidated it -- user controller")
		return
	}	
	log.Println("Password reset successfully -- controller")
	u.BaseHandler.Respond(ctx, http.StatusOK, gin.H{"message": "Password Reset successful"})
	
}

func (u UserController) ResetPassword(ctx *gin.Context) {
	var request dto.ResetPasswordDto 
	// bind input files
	if err := ctx.BindJSON(&request); err != nil {
		u.BaseHandler.Respond(ctx, http.StatusBadRequest, gin.H{"message": "Invalid Input"})
		log.Println("User Input cannot be bind -- user controller")
		return
	}

	// pass to usecase
	command := resetcodevalidate.NewCommand(request.Code, request.Id )
	_, err := u.resetPasswordhandler.Handle(command)

	if err != nil {
		u.BaseHandler.Respond(ctx, http.StatusInternalServerError, gin.H{"message": err.Error()})
		log.Println("User use case invalidated it -- user controller")
		return
	}	
	log.Println("User logged in successfully -- controller")
	u.BaseHandler.Respond(ctx, http.StatusOK, gin.H{"message": "Password Reset successful"})
	

}

func (u UserController) Logout(ctx *gin.Context) {
	u.RespondWithCookies(ctx, http.StatusNoContent, nil, []*http.Cookie{
		{
			Name:     "",
			Value:    "",
			Path:     "/",
			Domain:   ctx.Request.Host,
			MaxAge:   -1, // Delete the cookie
			HttpOnly: true,
			Secure:   true,
		},
	})
}


func (u UserController) Promte(ctx *gin.Context) {
	u.ChangeStatus(true, ctx)
}

func (u UserController) Demote(ctx *gin.Context) {
	u.ChangeStatus(false, ctx)
}

func (u UserController) ChangeStatus(toAdmin bool, ctx *gin.Context) {
	username := ctx.Param("id")
	claims, exists := ctx.Get("userClaims")
	if !exists {
		u.BaseHandler.Respond(ctx, http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		
		return
	}

	// Type assertion to jwt.MapClaims
	jwtClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		u.BaseHandler.Respond(ctx, http.StatusBadRequest, gin.H{"message": "Invalid Input"})
		return
	}

	// Extract and parse the user_id claim as a UUID
	userIDStr, ok := jwtClaims["user_id"].(string)
	if !ok {
		u.BaseHandler.Respond(ctx, http.StatusBadRequest, gin.H{"message": "Invalid Input"})
		return
	}

	promoterId, err := uuid.Parse(userIDStr)

	if err != nil {
		u.BaseHandler.Respond(ctx, http.StatusBadRequest, gin.H{"message": "Invalid Input"})
		return
	}

	_, err = u.promotHandler.Handle(promotcmd.NewCommand(username, toAdmin, promoterId))

	if err != nil {
	u.BaseHandler.Respond(ctx, http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	log.Println("User logged in successfully -- controller")
	u.BaseHandler.Respond(ctx, http.StatusNoContent, gin.H{"message": "User status changed successfully"})
	
}



func (u UserController) ValidateEmail(ctx *gin.Context) {

	encryptedValue := ctx.Query("secret")

	// Pass the secret to the use case
	_, err := u.validateEmailHandler.Handle(encryptedValue)

	if err != nil {
		u.BaseHandler.Respond(ctx, http.StatusInternalServerError, gin.H{"message": err.Error()})
		log.Println("User use case invalidated it -- user controller")
		return
	}   
	log.Println("Email validated successfully -- controller")
	u.BaseHandler.Respond(ctx, http.StatusOK, gin.H{"message": "Email validated successfully"})
}