package controller

import (
	"Blog_Starter/domain"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	UserUsecase  domain.UserUsecase
	OtpUsecase domain.OtpUsecase
}


func NewLoginController(LoginUsecase domain.LoginUsecase, OtpUsecase domain.OtpUsecase, UserUsecase  domain.UserUsecase) *LoginController {
	return &LoginController{
		LoginUsecase: LoginUsecase,
		UserUsecase:  UserUsecase,
		OtpUsecase: OtpUsecase,
	}
}

func (lc *LoginController) Login(c *gin.Context){
	var request domain.UserLogin
	
	err:= c.BindJSON(&request)
	if err!= nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid request"})
		return

	}

	request.Email = strings.ToLower(request.Email)
	_,err = lc.UserUsecase.GetUserByEmail(c, request.Email)
	if err !=nil{
		c.JSON(http.StatusNotFound, gin.H{"error" : "email not found during login"})
		return
	}

	loginResponse,err := lc.LoginUsecase.Login(c, &request)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "email or password not found"})
		return
	}


	// TODO: update tokens using the right method

	c.JSON(http.StatusOK, loginResponse)

}

func(lc *LoginController) ForgotPassword(c *gin.Context){

	var request domain.ForgotPasswordRequest
	err:= c.BindJSON(&request)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad email request"})
		return
	}
	request.Email = strings.ToLower(request.Email)
	user, err := lc.UserUsecase.GetUserByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound,gin.H{"error": "user not found with given email"})
		return
	}
	// Generate a random number between 0 and 9999 (inclusive).
	randNumber := rand.Intn(10000)

	// Format the code as a 4-digit string with leading zeros.
	code := fmt.Sprintf("%04d", randNumber)

	otp := domain.Otp{
		Email:      user.Email,
		Otp:        code,
		Expiration: time.Now().Add(5 * time.Minute),
	}

	oldOtp, err := lc.OtpUsecase.GetOtpByEmail(c, request.Email)
	if err == nil {
		otp.ID = oldOtp.ID
	} else {
		otp.ID = primitive.NewObjectID()
	}

	// Save OTP to database
	err = lc.OtpUsecase.SaveOtp(c, &otp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "couldnt save otp"})
		return
	}
	// TODO send email to user 

}

func (lc *LoginController) UpdatePassword(c *gin.Context){

	var request domain.ChangePasswordRequest
	err:= c.BindJSON(&request)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request to update password"})
		return
	}
	userID := c.Param("user_id")
	err2:= lc.LoginUsecase.UpdatePassword(c,request,userID)
	if err2!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "password not updated"})
		return 
	}
	c.JSON(http.StatusOK, gin.H{"success" : "password  updated"})

}



