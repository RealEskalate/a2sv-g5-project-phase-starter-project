package controllers_test

import (
	controllers "blogapp/Delivery/controllers"
	"blogapp/Domain"
	jwtservice "blogapp/Infrastructure/jwt_service"
	"blogapp/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// generate tag controller test suite

type RefreshCtrlTestSuit struct {
	suite.Suite
	ctrl        *controllers.RefreshController
	mockUsecase *mocks.RefreshUseCase
}

// setup test suite
func (suite *RefreshCtrlTestSuit) SetupTest() {
	suite.mockUsecase = new(mocks.RefreshUseCase)
	suite.ctrl = controllers.NewRefreshController(suite.mockUsecase)
}

// tear down
func (suite *RefreshCtrlTestSuit) TearDownTest() {
	suite.mockUsecase.AssertExpectations(suite.T())
}

// test refresh
func (suite *RefreshCtrlTestSuit) TestRefresh() {
	// test case success
	suite.Run("Refresh success", func() {
		claim := mocks.GetSampleClaim()
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("claim", claim)
		c.Request = httptest.NewRequest(http.MethodPost, "/refresh", nil)
		// find token
		suite.mockUsecase.On("FindToken", c, claim.ID).Return("token", nil, http.StatusOK).Once()
		// verify token
		jwtservice.VerifyRefreshToken = func(tokenString string, userid primitive.ObjectID) error {
			return nil
		}

		// create access token
		jwtservice.CreateAccessToken = func(user Domain.User) (string, error) {
			return "newtoken", nil
		}

		suite.ctrl.Refresh(c)
		suite.Equal(http.StatusOK, w.Code)
	})

	// test case fail
	suite.Run("Refresh fail", func() {
		claim := mocks.GetSampleClaim()
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("claim", claim)
		c.Request = httptest.NewRequest(http.MethodPost, "/refresh", nil)
		// find token
		suite.mockUsecase.On("FindToken", c, claim.ID).Return("", nil, http.StatusUnauthorized).Once()
		suite.ctrl.Refresh(c)
		suite.Equal(http.StatusUnauthorized, w.Code)
	})

}

// run test suite
func TestRefreshCtrlTestSuit(t *testing.T) {
	suite.Run(t, new(RefreshCtrlTestSuit))
}
