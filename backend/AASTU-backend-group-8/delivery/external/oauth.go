package external

import (
	"fmt"
	"meleket/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

type OauthHandler struct {
	usecases domain.UserUsecaseInterface
}

func NewOauthHandler(usecases domain.UserUsecaseInterface) *OauthHandler {
	return &OauthHandler{
		usecases: usecases,
	}
}

func (o *OauthHandler) SignInWithProvider(c *gin.Context) {
	provider := c.Param("provider")
	if provider == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Provider not specified"})
		return
	}

	q := c.Request.URL.Query()
	q.Add("provider", c.Param("provider"))
	c.Request.URL.RawQuery = q.Encode()

	req := c.Request
	res := c.Writer
	if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
		c.JSON(http.StatusOK, gothUser)
		return
	}
	gothic.BeginAuthHandler(res, req)
	return
}

func (o *OauthHandler) CallbackHandler(c *gin.Context) {

	req := c.Request
	res := c.Writer
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := o.usecases.GetUserByEmail(&user.Email); err != nil {
		err := o.usecases.Register(&domain.User{Email: user.Email, Name: user.Name})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusCreated)
		return
	}

	token, refreshToken, err := o.usecases.LoginWithProvider(&domain.User{Email: user.Email})
	fmt.Println("token: ", token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errori": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "refresh_token": refreshToken})
	return
}
