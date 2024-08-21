package controller

import (
	"context"
	"net/http"
	"time"

	b64 "encoding/base64"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/internal/tokenutil"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Env           *bootstrap.Env
}
func (sc *SignupController)VerifyEmail(c *gin.Context){
	Verificationtoken:=c.Param("Verificationtoken")
	decodedToken, _ :=b64.URLEncoding.DecodeString(Verificationtoken)
	if valid,err:=tokenutil.IsAuthorized(string(decodedToken),sc.Env.VerificationTokenSecret);!valid || err!=nil{
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid token"})
		return 
	}

	userID,err:=tokenutil.ExtractIDFromToken(string(decodedToken),sc.Env.VerificationTokenSecret)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	vuser,err:=sc.SignupUsecase.GetUserById(context.TODO(),userID)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	if vuser.VerToken!=string(decodedToken){
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid token"})
		return
	}
	err=sc.SignupUsecase.ActivateUser(context.TODO(),userID)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})

}
func (sc *SignupController) Signup(c *gin.Context) {
	var request domain.SignupRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = sc.SignupUsecase.GetUserByEmail(c, request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "User already exists with the given email"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	request.Password = string(encryptedPassword)

	IsOwner,err:=sc.SignupUsecase.IsOwner(c)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	var role string 
	if IsOwner{
		role="admin"
	}else{
		role="user"
	}

	user := domain.User{
		ID:        primitive.NewObjectID(),
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
		Active:    false,
		Bio:       request.Bio,
		ProfileImg: request.ProfileImg,
		IsOwner:   IsOwner,
		Role:      role,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	

	accessToken, err := sc.SignupUsecase.CreateAccessToken(&user, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := sc.SignupUsecase.CreateRefreshToken(&user, sc.Env.RefreshTokenSecret, sc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	VerificationToken, err := sc.SignupUsecase.CreateVerificationToken(&user, sc.Env.VerificationTokenSecret, sc.Env.VerificationTokenExpiryMin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	user.VerToken = VerificationToken
	user.Tokens = append(user.Tokens, refreshToken)

	_, err = sc.SignupUsecase.Create(c, &user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	//send email
	encodedToken:=b64.URLEncoding.EncodeToString([]byte(VerificationToken))
	err=sc.SignupUsecase.SendVerificationEmail(user.Email,encodedToken,sc.Env)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}

	signupResponse := domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, signupResponse)
}
