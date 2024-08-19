package Domain

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type Blog struct {
    Id           primitive.ObjectID `bson:"_id,omitempty"`
    Title        string             `bson:"title"`
    Body         string             `bson:"body"`
    Tags         []string           `bson:"tags"`
    CreatedAt    time.Time          `bson:"created_at"`
    LastUpdated  time.Time          `bson:"last_updated"`
    AuthorName   string             `bson:"author_name"`
    AuthorID     primitive.ObjectID `bson:"author_id"`
    ViewCount    int                `bson:"view_count"`
    LikeCount    int                `bson:"like_count"`
    CommentCount int                `bson:"comment_count"`
}
