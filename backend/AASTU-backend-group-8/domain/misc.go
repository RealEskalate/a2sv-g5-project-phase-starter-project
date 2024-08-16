package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tag struct {
    // ID        uint      `json:"id" gorm:"primaryKey"`
    ID        primitive.ObjectID    `json:"id"  bson:"_id,omitempty"`
    Name      string    `json:"name" validate:"required"`
}

type Pagination struct {
    Page    int `json:"page" validate:"min=1"`
    Limit   int `json:"limit" validate:"min=1,max=100"`
}

type BlogFilter struct {
    Tags []string `json:"tags"`
    DateRange struct {
        From time.Time `json:"from"`
        To   time.Time `json:"to"`
    } `json:"date_range"`
}

type RefreshToken struct {
	// username     string    `bson:"token" json:"token"`
	UserID    string    `bson:"user_id" json:"user_id"`
	ExpiresAt time.Time `bson:"expires_at" json:"expires_at"`
}

// // IsExpired checks if the refresh token is expired
// func (r *RefreshToken) IsExpired() bool {
// 	return time.Now().After(r.ExpiresAt)
// }


//todo:
// the AI model
// the token models