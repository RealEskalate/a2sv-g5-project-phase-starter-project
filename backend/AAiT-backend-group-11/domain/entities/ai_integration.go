package entities

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type ContentSuggestion struct {
    ID          primitive.ObjectID `bson:"_id,omitempty"`
    UserID      primitive.ObjectID `bson:"userId"`
    BlogPostID  primitive.ObjectID `bson:"blogPostId"`
    Suggestions []string           `bson:"suggestions"`
    CreatedAt   time.Time          `bson:"createdAt"`
}
