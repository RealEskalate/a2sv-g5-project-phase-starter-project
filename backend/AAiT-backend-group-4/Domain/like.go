package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Like struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID primitive.ObjectID `bson:"user_id" json:"user_id"`
	BlogID primitive.ObjectID `bson:"blog_id" json:"blog_id"`
	Status *bool              `bson:"status" json:"status"`
}


type LikeRepository interface {
	Create(c context.Context, userID string, blogID string, status bool) error
	Like(c context.Context, userID string, blogID string) error
	Dislike(c context.Context, userId string, blogID string) error
	RemoveLike(c context.Context, id string) error
	RemoveDislike(c context.Context, id string) error
	GetLikesByUser(ctx context.Context, userID string, limit, offset int) ([]Like, error)
	GetLikesByBlog(ctx context.Context, blogID string, limit, offset int) ([]Like, error)
	GetStatus(ctx context.Context, userID primitive.ObjectID, blogID primitive.ObjectID)(bool, string , error)
	
}

type LikeUsecase interface {
	Like(c context.Context, userID string, blogID string) error
	Dislike(c context.Context, userID string, blogID string) error
	RemoveLike(ctx context.Context, likeID, blogID string) error 
	RemoveDislike(ctx context.Context, dislikeID, blogID string) error 
	GetLikesByUser(ctx context.Context, userID string, limit, page int) ([]Like, error)
	GetLikesByBlog(ctx context.Context, blogID string, limit, page int) ([]Like, error)
	GetStatus(ctx context.Context, userID string, blogID string) (bool, string, error)
	
}
