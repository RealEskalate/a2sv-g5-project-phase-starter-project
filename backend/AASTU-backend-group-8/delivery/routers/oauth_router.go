package routers

import (
	"meleket/delivery/external"
	"meleket/usecases"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func NewOauthRouter(r *gin.Engine, userUsecase *usecases.UserUsecase) {
	goth.UseProviders(
		google.New(os.Getenv("OAUTH_CLIENT_ID"), os.Getenv("OAUTH_CLIENT_SECRET"), os.Getenv("OAUTH_CALLBACK_URL")),
	)

	oauthHandler := external.NewOauthHandler(userUsecase)

	gothic.Store = sessions.NewCookieStore([]byte("secret"))
	r.GET("/auth/:provider", oauthHandler.SignInWithProvider)
	r.GET("/auth/:provider/callback", oauthHandler.CallbackHandler)
}
