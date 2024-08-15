package Domain

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Token struct {
    TokenID       primitive.ObjectID `json:"token_id" bson:"token_id"`
    AccessToken   string             `json:"access_token" bson:"access_token"`
    RefreshToken  string             `json:"refresh_token" bson:"refresh_token"`
    Username      string             `json:"username" bson:"username"`
}
