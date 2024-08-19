package controller

import (
	"log"
	"net/http"
	"github.com/group13/blog/delivery/common"
	"github.com/gin-gonic/gin"
)

type userController struct{
	usecase usecase.UserUsecases
}

func (u userController) RegisterPrivileged (router *gin.RouterGroup){
	
}

func (u userController) RegisterPrivate (router *gin.RouterGroup){
	router = router.Group("/users")
	router.POST("/api/v1/auth/forgot-password", u.ForgotPassword)	
	router.POST("/api/v1/auth/reset-password", u.ResetPassword)	
	router.POST("POST /api/v1/users/:username/promote", u.Promte)
	router.POST("POST /api/v1/users/:username/demote", u.Demote)
	router.GET("POST /api/v1/users/:username/logout", u.Logout)

}


func (u userController) RegisterPublic  (router *gin.RouterGroup){
	router.POST("/api/v1/auth/signup", u.SignUp)
	router.POST("/api/v1/auth/login", u.Login)
	
}



func (u userController) SignUp(ctx *gin.Context) {

	var user models.User 
	// bind input files 
	if err := ctx.BindJSON(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message" :  "Invalid Input"})
		log.Println("User Input cannot be bind -- user controller")
		return
	}
	log.Println("User inputs bind successfully")
	// pass to usercases
	err := u.UserUsecase.Create(ctx, user)
	if err  != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message" : err.Error()})
		log.Println("User Usecase invalidated data -- user controller")
		return

	}
	log.Println("User singed up -- controller")
	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "Signed Up successfully"})
	
}


func (u userController) Login(ctx *gin.Context){
	var user models.User 
	// bind files to user model
	log.Println("User inputs bind successfully")

	// pass to login usercase
	if err:= userUsecases.Login; err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message" :  err.Error()})
		log.Println("User use case invalidated it -- user controller")
		return	
	}

	log.Println("User logged in successfully -- controller")
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Logged in successfully"})

}

func (u userController) ForgotPassword(ctx *gin.Context){
	var user models.User 
	// bind files to user model
	if err := ctx.BindJson(&user); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message" :  "Invalid Input"})
		log.Println("User Input cannot be bind -- user controller")
		return
	}
	log.Println("User inputs bind successfully")

	if err:= userUsecases.ForgotPassword; err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message" :  err.Error()})
		log.Println("User use case invalidated it -- user controller")
		return	
	}

	log.Println("User logged in successfully -- controller")
	ctx.IndentedJSON(http.StatusNoContent, gin.H{"message": "Logged in successfully"})	

}


func (u userController) ResetPassword(ctx *gin.Context){

	// Pass to usecase
	if err:= userUsecases.ResetPassword; err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message" :  err.Error()})
		log.Println("User use case invalidated it -- user controller")
		return	
	}

	log.Println("User logged in successfully -- controller")
	ctx.IndentedJSON(http.StatusNoContent, gin.H{"message": "Password Reset successful"})	

}



func (u userController) Logout(ctx *gin.Context){
	// Pass to usecase
	if err:= userUsecases.ResetPassword; err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message" :  err.Error()})
		log.Println("User use case invalidated it -- user controller")
		return	
	}

	log.Println("User logged in successfully -- controller")
	ctx.IndentedJSON(http.StatusNoContent, gin.H{"message": "Password Reset successful"})	

}

func (u userController) Promte(ctx *gin.Context){
	username := ctx.Param("id")
	if username == ""{
		log.Println("id not found -- user controller")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Input"})
		return
	}
	log.Println("id retrived --controller")
	// Pass to usecase
	if err:= userUsecases.Promote; err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message" :  err.Error()})
		log.Println("User use case invalidated it -- user controller")
		return	
	}

	log.Println("User logged in successfully -- controller")
	ctx.IndentedJSON(http.StatusNoContent, gin.H{"message": "Password Reset successful"})	

}


func (u userController) Demote(ctx *gin.Context){
	username := ctx.Param("id")
	if username == ""{
		log.Println("id not found -- user controller")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Input"})
		return
	}
	log.Println("id retrived --controller")
	// Pass to usecase
	if err:= userUsecases.Demote; err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message" :  err.Error()})
		log.Println("User use case invalidated it -- user controller")
		return	
	}

	log.Println("User logged in successfully -- controller")
	ctx.IndentedJSON(http.StatusNoContent, gin.H{"message": "Password Reset successful"})	

}








