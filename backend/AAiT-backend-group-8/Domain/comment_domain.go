package Domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	Body       string             `bson:"body"`
	CreatedAt  time.Time          `bson:"created_at"`
	AuthorName string             `bson:"author_name"`
	AuthorID   primitive.ObjectID `bson:"author_id"`
	BlogID     primitive.ObjectID `bson:"blog_id"`
}
