package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	b64 "encoding/base64"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/internal/tokenutil"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	SignupUsecase entities.SignupUsecase
	Env           *bootstrap.Env
}

func (sc *SignupController) VerifyEmail(c *gin.Context) {
	Verificationtoken := c.Param("token")
	decodedToken, _ := b64.URLEncoding.DecodeString(Verificationtoken)

	valid, err := tokenutil.IsAuthorized(string(decodedToken), sc.Env.VerificationTokenSecret)

	fmt.Println(string(decodedToken))

	if !valid || err != nil {
		c.JSON(http.StatusUnauthorized, entities.ErrorResponse{Message: "Invalid token"})
		return
	}

	claims, err := tokenutil.ExtractUserClaimsFromToken(string(decodedToken), sc.Env.VerificationTokenSecret)
	userID := claims["id"].(string)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entities.ErrorResponse{Message: err.Error()})
		return
	}
	user, err := sc.SignupUsecase.GetUserById(context.TODO(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entities.ErrorResponse{Message: err.Error()})
		return
	}
	if user.Active {
		c.JSON(http.StatusConflict, entities.ErrorResponse{Message: "User already activated"})
		return
	}

	err = sc.SignupUsecase.ActivateUser(context.TODO(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entities.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})

}
func (sc *SignupController) Signup(c *gin.Context) {
	var request entities.SignupRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, entities.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = sc.SignupUsecase.GetUserByEmail(c, request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, entities.ErrorResponse{Message: "User already exists with the given email"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entities.ErrorResponse{Message: err.Error()})
		return
	}

	request.Password = string(encryptedPassword)

	IsOwner, err := sc.SignupUsecase.IsOwner(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entities.ErrorResponse{Message: err.Error()})
		return
	}
	var role string
	if IsOwner {
		role = "admin"
	} else {
		role = "user"
	}

	user := entities.User{
		ID:        primitive.NewObjectID(),
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
		Active:    false,
		Bio:       request.Bio,
		// ProfileImg: request.ProfileImg,
		IsOwner:   IsOwner,
		Role:      role,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	VerificationToken, err := sc.SignupUsecase.CreateVerificationToken(&user, sc.Env.VerificationTokenSecret, sc.Env.VerificationTokenExpiryMin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entities.ErrorResponse{Message: err.Error()})
		return
	}
	user.VerToken = VerificationToken

	_, err = sc.SignupUsecase.Create(c, &user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, entities.ErrorResponse{Message: err.Error()})
		return
	}

	//send email
	encodedToken := b64.URLEncoding.EncodeToString([]byte(VerificationToken))
	err = sc.SignupUsecase.SendVerificationEmail(user.Email, encodedToken, sc.Env)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entities.ErrorResponse{Message: err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "email sent successfully, please verify your email"})

}
