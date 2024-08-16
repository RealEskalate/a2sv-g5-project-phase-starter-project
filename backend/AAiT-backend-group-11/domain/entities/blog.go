package entities

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogPost struct {
    ID          primitive.ObjectID `bson:"_id,omitempty"`
    Title       string             `bson:"title"`
    Content     string             `bson:"content"`
    AuthorID    primitive.ObjectID `bson:"authorId"`
    Tags        []string           `bson:"tags"`
    CreatedAt   time.Time          `bson:"createdAt"`
    UpdatedAt   time.Time          `bson:"updatedAt"`
    ViewCount   int                `bson:"viewCount"`
    LikeCount   int                `bson:"likeCount"`
    CommentCount int               `bson:"commentCount"`
}
type Like struct {
    ID          primitive.ObjectID `bson:"_id,omitempty"`
    BlogPostID  primitive.ObjectID `bson:"blogPostId"`
    AuthorID    primitive.ObjectID `bson:"authorId"`
    CreatedAt   time.Time          `bson:"createdAt"`
}
