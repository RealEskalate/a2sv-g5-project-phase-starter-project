package repository

import (
	"context"
	"meleket/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LikeDislikeRepository struct {
	db *mongo.Collection
}

func NewLikeDislikeRepository(db *mongo.Collection) *LikeDislikeRepository {
	return &LikeDislikeRepository{
		db: db,
	}
}

type LikeDislikeRepositoryInterface interface {
	AddLikeDislike(likeDislike *domain.LikeDislike) error
	RemoveLikeDislike(likeDislikeID primitive.ObjectID) error
	GetLikeDislikeByBlogAndUser(blogID, userID primitive.ObjectID) (*domain.LikeDislike, error)
	GetLikesDislikesByBlog(blogID primitive.ObjectID) ([]domain.LikeDislike, error)
}

func (r *LikeDislikeRepository) AddLikeDislike(likeDislike *domain.LikeDislike) error {
	_, err := r.db.InsertOne(context.TODO(), likeDislike)
	return err
}

func (r *LikeDislikeRepository) RemoveLikeDislike(likeDislikeID primitive.ObjectID) error {
	_, err := r.db.DeleteOne(context.TODO(), bson.M{"_id": likeDislikeID})
	return err
}

func (r *LikeDislikeRepository) GetLikeDislikeByBlogAndUser(blogID, userID primitive.ObjectID) (*domain.LikeDislike, error) {
	var likeDislike domain.LikeDislike
	filter := bson.M{"blog_id": blogID, "user_id": userID}
	err := r.db.FindOne(context.TODO(), filter).Decode(&likeDislike)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &likeDislike, err
}

func (r *LikeDislikeRepository) GetLikesDislikesByBlog(blogID primitive.ObjectID) ([]domain.LikeDislike, error) {
	var likesDislikes []domain.LikeDislike
	filter := bson.M{"blog_id": blogID}
	cursor, err := r.db.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var likeDislike domain.LikeDislike
		cursor.Decode(&likeDislike)
		likesDislikes = append(likesDislikes, likeDislike)
	}
	return likesDislikes, nil
}
