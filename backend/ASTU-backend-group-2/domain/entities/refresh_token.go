package entities

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionRefresh = "refresh"
)

type RefreshData struct{
	Id primitive.ObjectID `json:"_id"`
	UserId string `json:"userId"`
	RefreshToken string `json:" refreshToken"`
	Revoked bool `json:"revoked"`
	Expire_date  time.Time `json:"expire_date"`
}

type RefreshTokenRequest struct {
	RefreshToken string `form:"refreshToken" binding:"required"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenUsecase interface {
	GetUserByID(c context.Context, id string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int,refreshDataId string) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int,refreshDataID string) (refreshToken string, err error) 
	DeleteRefreshData(c context.Context, id string) error 
	GetRefreshData(c context.Context, id string) (*RefreshData, error) 
	CreateRefreshData(c context.Context, refreshData RefreshData) error 
	ExtractIDFromToken(requestToken string, secret string) (string, error)
}
