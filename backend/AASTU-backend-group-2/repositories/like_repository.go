package repositories

import (
	"blog_g2/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LikeRepository struct {
	client         *mongo.Client
	database       *mongo.Database
	likecollection *mongo.Collection
	blogcollection *mongo.Collection
}

func NewLikeRepository(mongoClient *mongo.Client) domain.LikeRepository {
	return &LikeRepository{
		client:         mongoClient,
		database:       mongoClient.Database("Blog-manager"),
		likecollection: mongoClient.Database("Blog-manager").Collection("Likes"),
		blogcollection: mongoClient.Database("Blog-manager").Collection("Blogs"),
	}

}

func (lrep *LikeRepository) GetLikes(post_id string) ([]domain.Like, error) {

	var likes []domain.Like
	postid, _ := primitive.ObjectIDFromHex(post_id)
	err := lrep.likecollection.FindOne(context.TODO(), bson.M{"_id": postid}).Decode(&likes)
	if err != nil {
		return []domain.Like{}, err
	}

	return likes, nil
}

func (lrep *LikeRepository) CreateLike(user_id string, post_id string) error {

	var like domain.Like
	like.UserID, _ = primitive.ObjectIDFromHex(user_id)
	like.PostID, _ = primitive.ObjectIDFromHex(post_id)

	_, err := lrep.likecollection.InsertOne(context.TODO(), like)

	if err != nil {
		return err
	}
	_, err = lrep.blogcollection.UpdateOne(context.TODO(), bson.M{"_id": like.PostID}, bson.M{"$inc": bson.M{"likes": 1}})
	if err != nil {
		return err
	}
	return nil
}

func (lrep *LikeRepository) DeleteLike(like_id string) error {

	likeID, _ := primitive.ObjectIDFromHex(like_id)
	var like domain.Like
	err := lrep.likecollection.FindOne(context.TODO(), bson.M{"_id": likeID}).Decode(&like)
	if err != nil {
		return err
	}
	_, err = lrep.blogcollection.UpdateOne(context.TODO(), bson.M{"_id": like.PostID}, bson.M{"$inc": bson.M{"likes": -1}})
	if err != nil {
		return err
	}
	_, err = lrep.likecollection.DeleteOne(context.TODO(), bson.M{"_id": likeID})
	if err != nil {
		return err
	}

	return nil
}
