package controller

import (
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/internal/emailutil"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/internal/otputil"
	"github.com/gin-gonic/gin"
)

type ResetPasswordController struct {
	ResetPasswordUsecase domain.ResetPasswordUsecase
	Env          *bootstrap.Env
}

func (rc *ResetPasswordController) ForgotPassword(c *gin.Context) {
	var req domain.ForgotPasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})	
		return
	}
	_,err:=rc.ResetPasswordUsecase.GetUserByEmail(c, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if _,err:=rc.ResetPasswordUsecase.GetOTPByEmail(c, req.Email); err==nil {
		err=rc.ResetPasswordUsecase.DeleteOtp(c,req.Email)
		if err!=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}	
	}
	otp,err:=otputil.GenerateOTP()
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err=emailutil.SendOtpVerificationEmail(req.Email, otp, rc.Env)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	hashedcode,err:=bcrypt.GenerateFromPassword([]byte(otp), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	newOtp:=domain.OtpSave{
		Email: req.Email,
		Code:  string(hashedcode),
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(rc.Env.PassResetCodeExpirationMin)),
	}
	err=rc.ResetPasswordUsecase.SaveOtp(c, &newOtp)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OTP sent successfully"})
}

func (rc *ResetPasswordController) ResetPassword(c *gin.Context) {
	var req domain.ResetPasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})	
		return
	}
	user,err:=rc.ResetPasswordUsecase.GetUserByEmail(c, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	originalOtp,err:=rc.ResetPasswordUsecase.GetOTPByEmail(c, req.Email)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if time.Now().After(originalOtp.ExpiresAt){
		c.JSON(http.StatusUnauthorized, gin.H{"error": "OTP expired"})
		return
	}

	err=bcrypt.CompareHashAndPassword([]byte(originalOtp.Code), []byte(req.Code))
	if err!=nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
		return
	}

	
	err=rc.ResetPasswordUsecase.ResetPassword(c, user.ID.Hex(), &req)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}	
	err=rc.ResetPasswordUsecase.DeleteOtp(c, req.Email)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}
