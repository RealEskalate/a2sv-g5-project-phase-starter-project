package repositories

import (
	"blog_g2/domain"
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LikeRepository struct {
	client            *mongo.Client
	database          *mongo.Database
	likecollection    *mongo.Collection
	blogcollection    *mongo.Collection
	dislikecollection *mongo.Collection
}

func NewLikeRepository(mongoClient *mongo.Client) domain.LikeRepository {
	return &LikeRepository{
		client:            mongoClient,
		database:          mongoClient.Database("Blog-manager"),
		likecollection:    mongoClient.Database("Blog-manager").Collection("Likes"),
		blogcollection:    mongoClient.Database("Blog-manager").Collection("Blogs"),
		dislikecollection: mongoClient.Database("Blog-manager").Collection("Dislikes"),
	}

}

func (lrep *LikeRepository) GetLikes(post_id string) ([]domain.Like, error) {

	var likes []domain.Like
	postid, _ := primitive.ObjectIDFromHex(post_id)
	log.Println("postid", postid)
	cursor, err := lrep.likecollection.Find(context.TODO(), bson.M{"post_id": postid})

	if err != nil {
		return []domain.Like{}, err
	}

	for cursor.Next(context.Background()) {
		var like domain.Like
		cursor.Decode(&like)
		likes = append(likes, like)
	}

	return likes, nil
}

func (lrep *LikeRepository) CreateLike(user_id string, post_id string) error {

	var like domain.Like
	like.UserID, _ = primitive.ObjectIDFromHex(user_id)
	like.BlogID, _ = primitive.ObjectIDFromHex(post_id)
	like.ID = primitive.NewObjectID()

	// check if user has already liked the post
	var checklike domain.Like
	erro := lrep.likecollection.FindOne(context.TODO(), bson.M{"user_id": like.UserID, "post_id": like.BlogID}).Decode(&checklike)
	if erro == nil {
		return errors.New("User has already liked this post")
	}

	lrep.deleteDisLike(user_id)

	_, err := lrep.likecollection.InsertOne(context.TODO(), like)

	if err != nil {
		return err
	}
	_, err = lrep.blogcollection.UpdateOne(context.TODO(), bson.M{"_id": like.BlogID}, bson.M{"$inc": bson.M{"likes": 1}})
	if err != nil {
		return err
	}
	return nil
}

func (lrep *LikeRepository) DeleteLike(like_id string) error {

	likeID, _ := primitive.ObjectIDFromHex(like_id)
	var like domain.Like
	err := lrep.likecollection.FindOne(context.TODO(), bson.M{"user_id": likeID}).Decode(&like)
	if err != nil {
		return err
	}
	_, err = lrep.blogcollection.UpdateOne(context.TODO(), bson.M{"_id": like.BlogID}, bson.M{"$inc": bson.M{"likes": -1}})

	if err != nil {
		return err
	}
	_, err = lrep.likecollection.DeleteOne(context.TODO(), bson.M{"user_id": likeID})
	if err != nil {
		return err
	}

	return nil
}

func (lrep *LikeRepository) deleteDisLike(dislike_id string) error {
	dislikeID, _ := primitive.ObjectIDFromHex(dislike_id)
	var dislike domain.DisLike

	err := lrep.dislikecollection.FindOne(context.TODO(), bson.M{"user_id": dislikeID}).Decode(&dislike)
	if err != nil {
		return err
	}
	_, err = lrep.blogcollection.UpdateOne(context.TODO(), bson.M{"_id": dislike.BlogID}, bson.M{"$inc": bson.M{"dislikes": -1}})
	if err != nil {
		return err
	}
	_, err = lrep.dislikecollection.DeleteOne(context.TODO(), bson.M{"user_id": dislikeID})
	if err != nil {
		return err
	}

	return nil
}
