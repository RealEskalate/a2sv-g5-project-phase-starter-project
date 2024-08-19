package controller

import (
	"Blog_Starter/domain"
	"context"
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

func (lc *LoginController) Login(c *gin.Context){
	var request domain.UserLogin
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	err:= c.BindJSON(&request)
	if err!= nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : "cannot bind login request"})
		return

	}
	request.Email = strings.ToLower(request.Email)
	_,err2:= lc.UserUsecase.GetUserByEmail(ctx, request.Email)
	if err2!=nil{
		c.JSON(http.StatusNotFound, gin.H{"error" : "email not found during login"})
		return
	}

	loginResponse,err := lc.LoginUsecase.Login(ctx, &request)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "email or password not found"})
		return
	}

	// TODO: update tokens using the right method
	c.JSON(http.StatusOK, loginResponse)
	



}

func(lc *LoginController) ForgotPassword(c gin.Context){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	var request domain.ForgotPasswordRequest
	err:= c.BindJSON(&request)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad email request"})
		return
	}
	request.Email = strings.ToLower(request.Email)
	user, err := lc.UserUsecase.GetUserByEmail(ctx, request.Email)
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

	oldOtp, err := lc.OtpUsecase.GetOtpByEmail(ctx, request.Email)
	if err == nil {
		otp.ID = oldOtp.ID
	} else {
		otp.ID = primitive.NewObjectID()
	}

	// Save OTP to database
	err = lc.OtpUsecase.SaveOtp(ctx, &otp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "couldnt save otp"})
		return
	}
	// TODO send email to user 

}

func (lc *LoginController) UpdatePassword(c gin.Context){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 100)
	defer cancel()
	var request domain.ChangePasswordRequest
	err:= c.BindJSON(&request)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request to update password"})
		return
	}
	userID := c.Param("user_id")
	err2:= lc.LoginUsecase.UpdatePassword(ctx,request,userID)
	if err2!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "password not updated"})
		return 
	}
	c.JSON(http.StatusOK, gin.H{"success" : "password  updated"})

}



