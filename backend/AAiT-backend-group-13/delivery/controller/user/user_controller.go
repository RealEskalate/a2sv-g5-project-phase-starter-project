package controller

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/group13/blog/delivery/controller/user/dto"
	er "github.com/group13/blog/domain/errors"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	promotcmd "github.com/group13/blog/usecase/user/command/promote"
	usercommand "github.com/group13/blog/usecases_sof/user/command"
	"github.com/group13/blog/usecases_sof/user/result"
)

type UserController struct {
	promotHandler icmd.IHandler[*promotcmd.Command, bool]
	loginHandler  icmd.IHandler[*usercommand.LoginCommand, *result.LoginInResult]
	signupHandler icmd.IHandler[*usercommand.SignUpCommand, *result.SignUpResult]
}

type Config struct {
	PromotHandler icmd.IHandler[*promotcmd.Command, bool]
	LoginHandler  icmd.IHandler[*usercommand.LoginCommand, *result.LoginInResult]
	SignupHandler icmd.IHandler[*usercommand.SignUpCommand, *result.SignUpResult]
}

// New creates a new UserController with the given CQRS handler.
func New(config Config) *UserController {
	return &UserController{
		promotHandler: config.PromotHandler,
		loginHandler:  config.LoginHandler,
		signupHandler: config.SignupHandler,
	}
}

func (u UserController) RegisterPrivileged(router *gin.RouterGroup) {

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
	router.POST("/api/v1/auth/signup", u.SignUp)
	router.POST("/api/v1/auth/login", u.Login)

}

func (u UserController) SignUp(ctx *gin.Context) {

	var user dto.SignUpDto
	// bind input files
	if err := ctx.BindJSON(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Input"})
		log.Println("User Input cannot be bind -- user controller")
		return
	}

	command := usercommand.NewSignUpCommand(user.Username, user.FirstName, user.LastName, user.Email, user.Password)
	// pass to usercases
	_, err := u.signupHandler.Handle(&command)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return

	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{})

}

func (u UserController) Login(ctx *gin.Context) {
	var user dto.LoginDto
	// bind files to user model
	// bind input files

	if err := ctx.BindJSON(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Input"})
		log.Println("User Input cannot be bind -- user controller")
		return
	}

	command := usercommand.NewLoginCommand(user.Username, user.Password)
	// pass to usercases
	_, err := u.loginHandler.Handle(&command)

	// pass to login usercase
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		log.Println("User use case invalidated it -- user controller")
		return
	}

	log.Println("User logged in successfully -- controller")
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Logged in successfully"})

}

func (u UserController) ForgotPassword(ctx *gin.Context) {

}

func (u UserController) ResetPassword(ctx *gin.Context) {

}

func (u UserController) Logout(ctx *gin.Context) {
	// Pass to usecase
	// if err := userUsecases.ResetPassword; err != nil {
	// 	ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	// 	log.Println("User use case invalidated it -- user controller")
	// 	return
	// }

	// log.Println("User logged in successfully -- controller")
	// ctx.IndentedJSON(http.StatusNoContent, gin.H{"message": "Password Reset successful"})

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
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": er.NewUnauthorized("unauthorized")})
		return
	}

	// Type assertion to jwt.MapClaims
	jwtClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": er.NewValidation("Invalid Claims")})
		return
	}

	// Extract and parse the user_id claim as a UUID
	userIDStr, ok := jwtClaims["user_id"].(string)
	if !ok {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": er.NewValidation("Invalid Claims")})
		return
	}

	promoterId, err := uuid.Parse(userIDStr)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": er.NewValidation("Invalid Claims")})
		return
	}

	_, err = u.promotHandler.Handle(promotcmd.NewCommand(username, toAdmin, promoterId))

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	log.Println("User logged in successfully -- controller")
	ctx.IndentedJSON(http.StatusNoContent, gin.H{"message": "Password Reset successful"})
}
