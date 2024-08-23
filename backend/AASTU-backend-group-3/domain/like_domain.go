package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Like struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	PostID  string `bson:"post_id,omitempty"`	
	UserID  string `bson:"user_id,omitempty"`
	Type   string `bson:"type,omitempty"`
}


type LikeRepository interface {
	
	LikeBlog(userID, blogID , Type string) error
	GetLekeExists(like Like) (bool, error)
	DisLikeBlog(userID, blogID, Type string) error
}

type LikeUsecase interface {
	LikeBlog(userID, blogID, Type string) *CustomError
	DisLikeBlog(userID, blogID, Type string) *CustomError
}



