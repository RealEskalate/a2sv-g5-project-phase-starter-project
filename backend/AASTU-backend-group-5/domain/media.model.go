package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Media struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	Path          string             `json:"path" bson:"path"`
	Uplaoded_date time.Time
}