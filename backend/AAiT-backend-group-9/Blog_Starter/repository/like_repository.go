package repository

import (
	"Blog_Starter/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LikeRepository struct {
	database   *mongo.Database
	collection string
}

// GetByID implements domain.LikeRepository.
func (l *LikeRepository) GetByID(c context.Context, userID string, blogID string) (*domain.Like, error) {
	collection:= l.database.Collection(l.collection)
	var like domain.Like
	err := collection.FindOne(c, bson.M{"user_id": userID, "blog_id": blogID}).Decode(&like)
	if err != nil {
		return nil,err


}
return &like, nil

}

// LikeBlog implements domain.LikeRepository.
func (l *LikeRepository) LikeBlog(c context.Context, like *domain.Like) (*domain.Like, error) {
	collection:= l.database.Collection(l.collection)
	_,err:= collection.InsertOne(c,like)
	if err!=nil{
		return nil,err
	}
	return like,nil
}

// UnlikeBlog implements domain.LikeRepository.
func (l *LikeRepository) UnlikeBlog(c context.Context, likeID string) (*domain.Like, error) {
	collection := l.database.Collection(l.collection)
	var like domain.Like
	err:= collection.FindOne(c, bson.M{"_id": likeID}).Decode(&like)
	_, err2 := collection.DeleteOne(c, bson.M{"_id": like.LikeID})
	if err!=nil{
		return nil,err2
	}
	
	return &like,nil

	
	
}

func NewLikeRepository(db *mongo.Database, collection string) domain.LikeRepository {
	return &LikeRepository{
		database:   db,
		collection: collection,
	}
}
