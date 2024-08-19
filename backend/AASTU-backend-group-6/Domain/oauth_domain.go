package domain

import (
	"context"

	"golang.org/x/oauth2"
)

type GoogleConfig struct {
    ClientId     string
    ClientSecret string
    REDIRECT_URI string
    // OauthSecret  string
}


type OauthConfig interface {
	InitialConfig() (*oauth2.Config , error)

}

type OauthUsecase interface {
	OauthService() (interface{})
	OauthCallback(c context.Context , query string) (interface{})

}

type URL struct {
	URL string
}