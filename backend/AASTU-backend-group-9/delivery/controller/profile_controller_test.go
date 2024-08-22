package controller_test

import (
	"blog/config"
	"blog/delivery/controller"
	"blog/domain"
	"blog/domain/mocks"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileControllerSuite struct {
	suite.Suite
	router            *gin.Engine
	ProfileUsecase    *mocks.ProfileUsecase
	ProfileController *controller.ProfileController
}

func (suite *ProfileControllerSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.ProfileUsecase = new(mocks.ProfileUsecase)
	env := &config.Env{
		SMTPUsername: "username",
		SMTPPassword: "password",
	}
	suite.ProfileController = &controller.ProfileController{
		ProfileUsecase: suite.ProfileUsecase,
		Env:            env,
	}
	suite.router = gin.Default()
	suite.router.PATCH("/update_profile", suite.ProfileController.UpdateProfile)
}

func (suite *ProfileControllerSuite) TearDownTest() {
	suite.ProfileUsecase.AssertExpectations(suite.T())
}

func (suite *ProfileControllerSuite) TestUpdateProfile() {
	suite.Run("update_user_success", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		user := domain.Profile{
			First_Name: "first",
			Last_Name:  "last",
			Bio:        "bio",
		}
		resp := domain.ProfileResponse{}
		claims := domain.JwtCustomClaims{}
		id, _ := primitive.ObjectIDFromHex(claims.Id)
		suite.ProfileUsecase.On("UpdateProfile", mock.Anything, &user, id).Return(&resp, nil).Once()
		payload, _ := json.Marshal(user)
		c.Request = httptest.NewRequest(http.MethodPatch, "/Update_profile", bytes.NewBuffer(payload))
		c.Set("claim", claims)
		suite.ProfileController.UpdateProfile(c)
		suite.Equal(200, w.Code)
		expect, err := json.Marshal( gin.H{"message": "Profile updated", "data": resp})
		suite.Nil(err)
		suite.Equal(expect, w.Body.Bytes())
	})
	suite.Run("update_user_error", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		user := domain.Profile{
			First_Name: "first",
			Last_Name:  "last",
			Bio:        "bio",
		}
		claims := domain.JwtCustomClaims{}
		id, _ := primitive.ObjectIDFromHex(claims.Id)
		suite.ProfileUsecase.On("UpdateProfile", mock.Anything, &user, id).Return(nil, errors.New("error")).Once()
		payload, _ := json.Marshal(user)
		c.Request = httptest.NewRequest(http.MethodPatch, "/Update_profile", bytes.NewBuffer(payload))
		c.Set("claim", claims)
		suite.ProfileController.UpdateProfile(c)
		suite.Equal(400, w.Code)
		expect, err := json.Marshal(gin.H{"error": "error"})
		suite.Nil(err)
		suite.Equal(expect, w.Body.Bytes())
	})
}

func TestProfileControllerSuite(t *testing.T) {
	suite.Run(t, new(ProfileControllerSuite))
}
