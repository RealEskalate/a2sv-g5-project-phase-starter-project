package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type URL struct {
	ID           primitive.ObjectID `bson:"_id"`
	ShortURLCode string             `bson:"short_url"`
	Token        string             `bson:"long_url"`
}
